package realtimeevents_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	. "github.com/tum-zulip/go-zulip/internal/test_utils"
	z "github.com/tum-zulip/go-zulip/zulip"
	realtimeevents "github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

func Test_RealmSettingsDefaults(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		queue := realtimeevents.NewEventQueue(apiClient)
		settings := queue.GetRealmSettings()

		require.Equal(t, int64(10000), settings.ServerTypingStartedWaitPeriodMilliseconds)
		require.Equal(t, int64(5000), settings.ServerTypingStoppedWaitPeriodMilliseconds)
		require.Equal(t, int64(15000), settings.ServerTypingStartedExpiryPeriodMilliseconds)
	})
}

func Test_UpdateRealmSettingsFromResponse(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, _, err := apiClient.RegisterQueue(ctx).
			FetchEventTypes([]events.EventType{events.EventTypeRealm}).
			Execute()
		require.NoError(t, err)

		queue := realtimeevents.NewEventQueue(apiClient)
		queue.UpdateRealmSettings(resp)

		settings := queue.GetRealmSettings()
		require.NotNil(t, settings)
		require.Positive(t, settings.ServerTypingStartedWaitPeriodMilliseconds)
		require.Positive(t, settings.ServerTypingStoppedWaitPeriodMilliseconds)
		require.Positive(t, settings.ServerTypingStartedExpiryPeriodMilliseconds)
	})
}

func Test_RealmSettingsConcurrentAccess(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, _, err := apiClient.RegisterQueue(ctx).
			FetchEventTypes([]events.EventType{events.EventTypeRealm}).
			Execute()
		require.NoError(t, err)

		queue := realtimeevents.NewEventQueue(apiClient)

		var wg sync.WaitGroup

		// Concurrent readers
		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 100 {
					settings := queue.GetRealmSettings()
					if settings == nil || settings.ServerTypingStartedWaitPeriodMilliseconds <= 0 {
						t.Error("invalid realm settings from concurrent read")
						return
					}
				}
			}()
		}

		// Concurrent writer
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				queue.UpdateRealmSettings(resp)
			}
		}()

		wg.Wait()
	})
}

func Test_StartTypingSendsNotification(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeTyping}).
			Execute()
		require.NoError(t, err)

		queue := realtimeevents.NewEventQueue(apiClient)
		queue.UpdateRealmSettings(resp)

		evChan, err := queue.Connect(ctx, *resp.QueueID, resp.LastEventID)
		require.NoError(t, err)
		defer func() { require.NoError(t, queue.Close()) }()

		notifier, err := queue.StartTyping(ctx, apiClient, z.UserAsRecipient(GetUserID(t, apiClient)))
		require.NoError(t, err)
		require.NotNil(t, notifier)

		// Should receive typing start event
		select {
		case e := <-evChan:
			require.NotNil(t, e)
			_, ok := e.(events.TypingEvent)
			require.True(t, ok, "expected TypingEvent, got %T", e)
		case <-time.After(5 * time.Second):
			t.Fatal("timeout waiting for typing event")
		}

		require.NoError(t, notifier.Close())
	})
}
