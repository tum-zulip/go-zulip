package zulip_test

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func Test_RealTimeEventsAPIService(t *testing.T) {
	t.Parallel()

	t.Run("DeleteQueue", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		queueId, _ := registerMessageEventQueue(t, apiClient)

		resp, httpRes, err := apiClient.DeleteQueue(ctx).
			QueueId(queueId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetEvents", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		ctx := context.Background()

		queueId, lastEventId := registerMessageEventQueue(t, apiClient)
		go func() {
			time.Sleep(2 * time.Second)
			_, _, err := apiClient.SetTypingStatus(ctx).
				Op(zulip.TypingStatusOpStart).
				To(zulip.UserAsRecipient(getOwnUserId(t, apiClient))).
				Execute()

			require.NoError(t, err)
		}()

		resp, httpRes, err := apiClient.GetEvents(ctx).
			QueueId(queueId).
			LastEventId(lastEventId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("RegisterQueue", runForAllClients(t, func(t *testing.T, apiClient zulip.Client) {
		registerMessageEventQueue(t, apiClient)
	}))

}

func registerMessageEventQueue(t *testing.T, apiClient zulip.Client) (string, int64) {
	t.Helper()

	ctx := context.Background()

	resp, httpRes, err := apiClient.RegisterQueue(ctx).Execute()

	slog.Error("debug", "resp", resp, "httpRes", httpRes, "err", err)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	require.NotEmpty(t, resp.QueueId)

	return *resp.QueueId, resp.LastEventId
}
