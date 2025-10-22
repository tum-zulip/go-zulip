package zulip_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func Test_RealTimeEventsAPIService(t *testing.T) {

	t.Run("DeleteQueue", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		queueId, _ := registerMessageEventQueue(t, apiClient)

		resp, httpRes, err := apiClient.DeleteQueue(ctx).
			QueueId(queueId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetEvents", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		queueId, lastEventId := registerMessageEventQueue(t, apiClient)
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := apiClient.SetTypingStatus(ctx).
				Op(z.TypingStatusOpStart).
				To(z.UserAsRecipient(getOwnUserId(t, apiClient))).
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

	t.Run("RegisterQueue", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		registerMessageEventQueue(t, apiClient)
	}))

}

func registerMessageEventQueue(t *testing.T, apiClient z.Client) (string, int64) {
	t.Helper()

	ctx := context.Background()

	resp, httpRes, err := apiClient.RegisterQueue(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	require.NotEmpty(t, resp.QueueId)

	return *resp.QueueId, resp.LastEventId
}
