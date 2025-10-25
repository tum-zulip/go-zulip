package real_time_events_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_DeleteQueue(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		queueId, _ := registerMessageEventQueue(t, apiClient)

		resp, httpResp, err := apiClient.DeleteQueue(ctx).
			QueueId(queueId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetEvents(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		queueId, lastEventId := registerMessageEventQueue(t, apiClient)
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := apiClient.SetTypingStatus(ctx).
				Op(z.TypingStatusOpStart).
				To(z.UserAsRecipient(GetUserId(t, apiClient))).
				Execute()

			require.NoError(t, err)
		}()

		resp, httpResp, err := apiClient.GetEvents(ctx).
			QueueId(queueId).
			LastEventId(lastEventId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_RegisterQueue(t *testing.T) {

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		registerMessageEventQueue(t, apiClient)
	})

}

func registerMessageEventQueue(t *testing.T, apiClient client.Client) (string, int64) {
	t.Helper()

	ctx := context.Background()

	resp, httpResp, err := apiClient.RegisterQueue(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)

	require.NotEmpty(t, resp.QueueId)
	// TODO: janez require.NotNil(t, resp.Realm)

	// TODO: require.WithinDuration(t, time.Now(), resp.Realm.DateCreated, 59*time.Minute)
	// TODO: require.WithinDuration(t, time.Now(), resp.Realm.ServerGeneration, 59*time.Minute)
	// TODO: also test resp.Realm.PushNotificationsEndabledEndTimestamp

	return *resp.QueueId, resp.LastEventId
}
