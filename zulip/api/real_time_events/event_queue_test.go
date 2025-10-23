package real_time_events_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func TestEventQueue(t *testing.T) {
	t.Run("Connect requires queue ID", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		q := real_time_events.NewEventQueue(apiClient, nil)

		events, err := q.Connect(context.Background(), "", 0)
		require.Error(t, err)
		require.Nil(t, events)
	}))

	t.Run("Polls events and updates state", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.RegisterQueue(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, httpResp.StatusCode, 200)
		require.NotNil(t, resp.QueueId)

		q := real_time_events.NewEventQueue(apiClient, real_time_events.NewEventQueueLoggingErrorHandler())

		events, err := q.Connect(context.Background(), *resp.QueueId, resp.LastEventId)
		require.NoError(t, err)
		require.NotNil(t, events)
		require.Equal(t, *resp.QueueId, q.QueueId())
		require.Equal(t, resp.LastEventId, q.LastEventId())

		go func() {
			for i := 0; i < 2; i++ {
				time.Sleep(200 * time.Millisecond)
				apiClient.SetTypingStatus(ctx).
					Type_(z.RecipientTypeDirect).
					To(z.UserAsRecipient(GetUserId(t, apiClient))).
					Op(z.TypingStatusOpStart).
					Execute()
			}
		}()

		e1 := <-events
		require.NotNil(t, e1)

		e2 := <-events
		require.NotNil(t, e2)

		require.Greater(t, e2.GetId(), e1.GetId())

		require.NoError(t, q.Close())
		for range events {
		}
	}))

	t.Run("Close without Connect", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		q := real_time_events.NewEventQueue(apiClient, nil)

		require.Error(t, q.Close())
	}))
}
