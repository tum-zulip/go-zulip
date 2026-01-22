package realtimeevents_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/tum-zulip/go-zulip/internal/test_fixtures"
	"github.com/tum-zulip/go-zulip/zulip"
	realtimeevents "github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

func setupEventMockServer() *httptest.Server {
	var pollCount atomic.Int64
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testfixtures.RegisterQueueResponse))
	})

	mux.HandleFunc("/api/v1/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Write([]byte(testfixtures.DeleteQueueResponse))
			return
		}
		// Simulate realistic event distribution: 90% empty/heartbeat, 10% messages
		count := pollCount.Add(1)
		switch count % 10 {
		case 0:
			w.Write([]byte(testfixtures.EventsResponseMessage))
		case 5:
			w.Write([]byte(testfixtures.EventsResponseHeartbeat))
		default:
			w.Write([]byte(testfixtures.EventsResponseEmpty))
		}
	})

	return httptest.NewServer(mux)
}

func createBenchmarkClient(b *testing.B, serverURL string) client.Client {
	b.Helper()
	insecure := true
	c, _ := client.NewClient(&zulip.RC{
		Site:     serverURL,
		Email:    "bench@example.com",
		APIKey:   "key",
		Insecure: &insecure,
	},
		client.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 5,
				IdleConnTimeout:     30 * time.Second,
			},
		}),
		client.WithLogger(slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{Level: slog.LevelError}))),
		client.SkipWarnOnInsecureTLS(),
	)
	return c
}

func runConcurrent(b *testing.B, concurrency int, fn func(workerID, iterations int)) {
	b.Helper()
	b.ResetTimer()
	var wg sync.WaitGroup
	for i := range concurrency {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			iter := b.N / concurrency
			if id == 0 {
				iter += b.N % concurrency
			}
			fn(id, iter)
		}(i)
	}
	wg.Wait()
}

func BenchmarkRegisterQueue(b *testing.B) {
	server := setupEventMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	b.ResetTimer()
	for range b.N {
		c.RegisterQueue(ctx).Execute()
	}
}

//nolint:funlen
func BenchmarkGetEvents(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		server := setupEventMockServer()
		defer server.Close()
		c := createBenchmarkClient(b, server.URL)
		ctx := context.Background()

		resp, _, _ := c.RegisterQueue(ctx).Execute()
		queueID := *resp.QueueID
		lastEventID := resp.LastEventID

		b.ResetTimer()
		for range b.N {
			resp, _, _ := c.GetEvents(ctx).
				QueueID(queueID).
				LastEventID(lastEventID).
				DontBlock(true).
				Execute()
			if len(resp.Events) > 0 {
				lastEventID = resp.Events[len(resp.Events)-1].GetID()
			}
		}
	})

	b.Run("Concurrent-10", func(b *testing.B) {
		server := setupEventMockServer()
		defer server.Close()
		c := createBenchmarkClient(b, server.URL)
		ctx := context.Background()
		resp, _, _ := c.RegisterQueue(ctx).Execute()
		queueID := *resp.QueueID

		runConcurrent(b, 10, func(_, iter int) {
			lastEventID := int64(0)
			for range iter {
				resp, _, _ := c.GetEvents(ctx).
					QueueID(queueID).
					LastEventID(lastEventID).
					DontBlock(true).
					Execute()
				if len(resp.Events) > 0 {
					lastEventID = resp.Events[len(resp.Events)-1].GetID()
				}
			}
		})
	})

	b.Run("Concurrent-25", func(b *testing.B) {
		server := setupEventMockServer()
		defer server.Close()
		c := createBenchmarkClient(b, server.URL)
		ctx := context.Background()
		resp, _, _ := c.RegisterQueue(ctx).Execute()
		queueID := *resp.QueueID

		runConcurrent(b, 25, func(_, iter int) {
			lastEventID := int64(0)
			for range iter {
				resp, _, _ := c.GetEvents(ctx).
					QueueID(queueID).
					LastEventID(lastEventID).
					DontBlock(true).
					Execute()
				if len(resp.Events) > 0 {
					lastEventID = resp.Events[len(resp.Events)-1].GetID()
				}
			}
		})
	})
}

func BenchmarkEventQueueListener(b *testing.B) {
	runListener := func(b *testing.B, consumers int) {
		server := setupEventMockServer()
		defer server.Close()
		c := createBenchmarkClient(b, server.URL)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		regResp, _, _ := c.RegisterQueue(ctx).Execute()
		eq := realtimeevents.NewEventQueue(c,
			realtimeevents.WithLogger(slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{Level: slog.LevelError}))),
			realtimeevents.WithEventQueueChannelBufferSize(100),
		)

		eventChan, _ := eq.Connect(ctx, *regResp.QueueID, -1)
		defer eq.Close()

		var wg sync.WaitGroup
		var processed atomic.Int64
		b.ResetTimer()

		for range consumers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range eventChan {
					if processed.Add(1) >= int64(b.N) {
						cancel()
						return
					}
				}
			}()
		}
		wg.Wait()
		b.StopTimer()
	}

	b.Run("SingleConsumer", func(b *testing.B) { runListener(b, 1) })
	b.Run("Concurrent-5", func(b *testing.B) { runListener(b, 5) })
	b.Run("Concurrent-10", func(b *testing.B) { runListener(b, 10) })
}

func BenchmarkLongLivedQueue(b *testing.B) {
	server := setupEventMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	resp, _, _ := c.RegisterQueue(ctx).Execute()
	queueID := *resp.QueueID
	lastEventID := resp.LastEventID

	b.ResetTimer()
	for range b.N {
		resp, _, _ := c.GetEvents(ctx).
			QueueID(queueID).
			LastEventID(lastEventID).
			DontBlock(true).
			Execute()
		if len(resp.Events) > 0 {
			lastEventID = resp.Events[len(resp.Events)-1].GetID()
		}
	}
}

func BenchmarkMultipleQueues(b *testing.B) {
	runMultiple := func(b *testing.B, numQueues, pollsPerQueue int) {
		server := setupEventMockServer()
		defer server.Close()
		ctx := context.Background()

		type queueInfo struct {
			client      client.Client
			queueID     string
			lastEventID int64
		}

		queues := make([]queueInfo, numQueues)
		for i := range numQueues {
			c := createBenchmarkClient(b, server.URL)
			resp, _, _ := c.RegisterQueue(ctx).Execute()
			queues[i] = queueInfo{c, *resp.QueueID, resp.LastEventID}
		}

		b.ResetTimer()
		var wg sync.WaitGroup
		for qi := range numQueues {
			wg.Add(1)
			go func(q queueInfo) {
				defer wg.Done()
				lastEventID := q.lastEventID
				for range b.N / numQueues {
					for range pollsPerQueue {
						resp, _, _ := q.client.GetEvents(ctx).
							QueueID(q.queueID).
							LastEventID(lastEventID).
							DontBlock(true).
							Execute()
						if len(resp.Events) > 0 {
							lastEventID = resp.Events[len(resp.Events)-1].GetID()
						}
					}
				}
			}(queues[qi])
		}
		wg.Wait()
	}

	b.Run("Queues-3-Polls-10", func(b *testing.B) { runMultiple(b, 3, 10) })
	b.Run("Queues-5-Polls-10", func(b *testing.B) { runMultiple(b, 5, 10) })
	b.Run("Queues-10-Polls-5", func(b *testing.B) { runMultiple(b, 10, 5) })
}

func BenchmarkEventProcessing(b *testing.B) {
	server := setupEventMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	resp, _, _ := c.RegisterQueue(ctx).Execute()
	queueID := *resp.QueueID

	b.Run("Heartbeat", func(b *testing.B) {
		b.ResetTimer()
		for i := range b.N {
			resp, _, _ := c.GetEvents(ctx).
				QueueID(queueID).
				LastEventID(int64(i)).
				DontBlock(true).
				Execute()
			for _, event := range resp.Events {
				_ = event.GetType()
				_ = event.GetID()
			}
		}
	})

	b.Run("Message", func(b *testing.B) {
		b.ResetTimer()
		for i := range b.N {
			resp, _, _ := c.GetEvents(ctx).
				QueueID(queueID).
				LastEventID(int64(i * 10)).
				DontBlock(true).
				Execute()
			for _, event := range resp.Events {
				if event.GetType() == events.EventTypeMessage {
					_, _ = event.GetOp()
				}
			}
		}
	})
}

func BenchmarkDeleteQueue(b *testing.B) {
	server := setupEventMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	queueIDs := make([]string, b.N)
	for i := range b.N {
		resp, _, _ := c.RegisterQueue(ctx).Execute()
		queueIDs[i] = *resp.QueueID
	}

	b.ResetTimer()
	for i := range b.N {
		c.DeleteQueue(ctx).QueueID(queueIDs[i]).Execute()
	}
}
