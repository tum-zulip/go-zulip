package reminders_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/tum-zulip/go-zulip/internal/test_utils"
	"github.com/tum-zulip/go-zulip/zulip/client"
)

func Test_CreateMessageReminder(t *testing.T) {
	RequireFeatureLevel(t, 381)

	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		createMessageReminder(t, apiClient, channelID)
	})
}

func Test_DeleteReminder(t *testing.T) {
	RequireFeatureLevel(t, 399)

	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		reminderID := createMessageReminder(t, apiClient, channelID)

		resp, _, err := apiClient.DeleteReminder(ctx, reminderID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_GetReminders(t *testing.T) {
	RequireFeatureLevel(t, 399)

	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		reminderID := createMessageReminder(t, apiClient, channelID)

		resp, _, err := apiClient.GetReminders(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.GreaterOrEqual(t, len(resp.Reminders), 1)
		found := false
		for _, r := range resp.Reminders {
			if r.ReminderID == reminderID {
				require.WithinDuration(t, time.Now().Add(1*time.Hour), r.ScheduledDeliveryTimestamp, 3*time.Minute)
				found = true
			}
		}
		assert.True(t, found, "Created reminder not found in list of reminders")
	})
}

func createMessageReminder(t *testing.T, apiClient client.Client, channelID int64) int64 {
	msg := CreateChannelMessage(t, apiClient, channelID)

	note := "This is a reminder note"

	resp, _, err := apiClient.CreateMessageReminder(context.Background()).
		MessageID(msg.MessageID).
		Note(UniqueName(note)).
		ScheduledDeliveryTimestamp(time.Now().Add(1 * time.Hour)).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	return resp.ReminderID
}
