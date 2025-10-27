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

func Test_ConnectRequiresQueueID(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		q := real_time_events.NewEventQueue(apiClient)

		events, err := q.Connect(context.Background(), "", 0)
		require.Error(t, err)
		require.Nil(t, events)
	})
}

func Test_PollsEventsAndUpdatesState(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpResp, err := apiClient.RegisterQueue(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, httpResp.StatusCode, 200)
		require.NotNil(t, resp.QueueID)

		q := real_time_events.NewEventQueue(apiClient)

		events, err := q.Connect(context.Background(), *resp.QueueID, resp.LastEventID)
		require.NoError(t, err)
		require.NotNil(t, events)
		require.Equal(t, *resp.QueueID, q.QueueID())
		require.Equal(t, resp.LastEventID, q.LastEventID())

		errs := make([]error, 0)
		wait := make(chan struct{})

		go func() {
			for i := 0; i < 2; i++ {
				time.Sleep(200 * time.Millisecond)
				_, _, typingErr := apiClient.SetTypingStatus(ctx).
					To(z.UserAsRecipient(GetUserID(t, apiClient))).
					Op(z.TypingStatusOpStart).
					Execute()

				errs = append(errs, typingErr)
			}
			close(wait)
		}()

		e1 := <-events
		require.NotNil(t, e1)

		e2 := <-events
		require.NotNil(t, e2)

		require.Greater(t, e2.GetID(), e1.GetID())

		require.NoError(t, q.Close())
		for range events {
		}
		<-wait

		for _, err := range errs {
			require.NoError(t, err)
		}
	})
}

func Test_CloseWithoutConnect(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		q := real_time_events.NewEventQueue(apiClient, nil)

		require.Error(t, q.Close())
	})
}
