package realtimeevents_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_DeleteQueue(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		queueID, _ := registerMessageEventQueue(t, apiClient)

		resp, _, err := apiClient.DeleteQueue(ctx).
			QueueID(queueID).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_GetEvents(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		queueID, lastEventID := registerMessageEventQueue(t, apiClient)

		var typingErr error
		wait := make(chan struct{})
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, typingErr = apiClient.SetTypingStatus(ctx).
				Op(z.TypingStatusOpStart).
				To(z.UserAsRecipient(GetUserID(t, apiClient))).
				Execute()

			close(wait)
		}()

		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(queueID).
			LastEventID(lastEventID).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)

		<-wait
		require.NoError(t, typingErr)
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

	resp, _, err := apiClient.RegisterQueue(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotEmpty(t, resp.QueueID)
	// TODO: janez require.NotNil(t, resp.Realm)

	// TODO: require.WithinDuration(t, time.Now(), resp.Realm.DateCreated, 59*time.Minute)
	// TODO: require.WithinDuration(t, time.Now(), resp.Realm.ServerGeneration, 59*time.Minute)
	// TODO: also test resp.Realm.PushNotificationsEndabledEndTimestamp

	return *resp.QueueID, resp.LastEventID
}
