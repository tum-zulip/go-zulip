package client_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/tum-zulip/go-zulip/internal/test_fixtures"
	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
)

func setupAPIMockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testfixtures.SendMessageResponse))
	})
	mux.HandleFunc("/api/v1/streams", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testfixtures.GetChannelsResponse))
	})
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testfixtures.GetUsersResponse))
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

func BenchmarkSendMessage(b *testing.B) {
	server := setupAPIMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	b.Run("Sequential", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			c.SendMessage(ctx).
				RecipientType(zulip.RecipientTypeChannel).
				To(zulip.ChannelAsRecipient(1)).
				Topic("benchmark").
				Content("test message").
				Execute()
		}
	})

	b.Run("Concurrent-10", func(b *testing.B) {
		runConcurrent(b, 10, func(_, iter int) {
			for range iter {
				c.SendMessage(ctx).
					RecipientType(zulip.RecipientTypeChannel).
					To(zulip.ChannelAsRecipient(1)).
					Topic("benchmark").
					Content("test message").
					Execute()
			}
		})
	})

	b.Run("Concurrent-25", func(b *testing.B) {
		runConcurrent(b, 25, func(_, iter int) {
			for range iter {
				c.SendMessage(ctx).
					RecipientType(zulip.RecipientTypeChannel).
					To(zulip.ChannelAsRecipient(1)).
					Topic("benchmark").
					Content("test message").
					Execute()
			}
		})
	})
}

func BenchmarkGetChannels(b *testing.B) {
	server := setupAPIMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	b.Run("Sequential", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			c.GetChannels(ctx).Execute()
		}
	})

	b.Run("Concurrent-10", func(b *testing.B) {
		runConcurrent(b, 10, func(_, iter int) {
			for range iter {
				c.GetChannels(ctx).Execute()
			}
		})
	})

	b.Run("Concurrent-25", func(b *testing.B) {
		runConcurrent(b, 25, func(_, iter int) {
			for range iter {
				c.GetChannels(ctx).Execute()
			}
		})
	})
}

func BenchmarkGetUsers(b *testing.B) {
	server := setupAPIMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	b.Run("Sequential", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			c.GetUsers(ctx).Execute()
		}
	})

	b.Run("Concurrent-10", func(b *testing.B) {
		runConcurrent(b, 10, func(_, iter int) {
			for range iter {
				c.GetUsers(ctx).Execute()
			}
		})
	})

	b.Run("Concurrent-25", func(b *testing.B) {
		runConcurrent(b, 25, func(_, iter int) {
			for range iter {
				c.GetUsers(ctx).Execute()
			}
		})
	})
}

//nolint:funlen
func BenchmarkMixedWorkload(b *testing.B) {
	server := setupAPIMockServer()
	defer server.Close()
	c := createBenchmarkClient(b, server.URL)
	ctx := context.Background()

	b.Run("Sequential", func(b *testing.B) {
		b.ResetTimer()
		for i := range b.N {
			switch i % 5 {
			case 0, 1, 2:
				c.SendMessage(ctx).
					RecipientType(zulip.RecipientTypeChannel).
					To(zulip.ChannelAsRecipient(1)).
					Topic("benchmark").
					Content("test").
					Execute()
			case 3:
				c.GetChannels(ctx).Execute()
			case 4:
				c.GetUsers(ctx).Execute()
			}
		}
	})

	b.Run("Concurrent-10", func(b *testing.B) {
		runConcurrent(b, 10, func(_, iter int) {
			for j := range iter {
				switch j % 5 {
				case 0, 1, 2:
					c.SendMessage(ctx).
						RecipientType(zulip.RecipientTypeChannel).
						To(zulip.ChannelAsRecipient(1)).
						Topic("benchmark").
						Content("test").
						Execute()
				case 3:
					c.GetChannels(ctx).Execute()
				case 4:
					c.GetUsers(ctx).Execute()
				}
			}
		})
	})

	b.Run("Concurrent-25", func(b *testing.B) {
		runConcurrent(b, 25, func(_, iter int) {
			for j := range iter {
				switch j % 5 {
				case 0, 1, 2:
					c.SendMessage(ctx).
						RecipientType(zulip.RecipientTypeChannel).
						To(zulip.ChannelAsRecipient(1)).
						Topic("benchmark").
						Content("test").
						Execute()
				case 3:
					c.GetChannels(ctx).Execute()
				case 4:
					c.GetUsers(ctx).Execute()
				}
			}
		})
	})
}

func BenchmarkClientInitialization(b *testing.B) {
	server := setupAPIMockServer()
	defer server.Close()
	insecure := true
	rc := &zulip.RC{
		Site:     server.URL,
		Email:    "bench@example.com",
		APIKey:   "key",
		Insecure: &insecure,
	}

	b.Run("WithDefaults", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			client.NewClient(rc, client.SkipWarnOnInsecureTLS())
		}
	})

	b.Run("WithLogger", func(b *testing.B) {
		logger := slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{Level: slog.LevelError}))
		b.ResetTimer()
		for range b.N {
			client.NewClient(rc, client.WithLogger(logger), client.SkipWarnOnInsecureTLS())
		}
	})

	b.Run("WithStatistics", func(b *testing.B) {
		logger := slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{Level: slog.LevelError}))
		b.ResetTimer()
		for range b.N {
			client.NewClient(rc, client.WithLogger(logger), client.EnableStatistics(), client.SkipWarnOnInsecureTLS())
		}
	})
}
