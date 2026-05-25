package realtimeevents_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	. "github.com/tum-zulip/go-zulip/internal/test_utils"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
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

	resp, _, err := apiClient.RegisterQueue(ctx).
		FetchEventTypes([]events.EventType{events.EventTypeRealm, events.EventTypeMessage}).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	require.NotEmpty(t, resp.QueueID)
	require.NotNil(t, resp.RealmDateCreated)
	require.NotNil(t, resp.ServerGeneration)

	realmDateCreated := time.Unix(*resp.RealmDateCreated, 0).UTC()
	serverGeneration := time.Unix(*resp.ServerGeneration, 0).UTC()

	require.WithinDuration(t, time.Now(), realmDateCreated, 59*time.Minute)
	require.WithinDuration(t, time.Now(), serverGeneration, 59*time.Minute)

	// PushNotificationsEnabledEndTimestamp is optional and may be nil
	if resp.PushNotificationsEnabledEndTimestamp != nil {
		pushNotifTimestamp := time.Unix(*resp.PushNotificationsEnabledEndTimestamp, 0).UTC()
		require.True(t, pushNotifTimestamp.After(time.Now()))
	}

	return *resp.QueueID, resp.LastEventID
}
