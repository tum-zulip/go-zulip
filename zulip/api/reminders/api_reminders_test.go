package reminders_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_CreateMessageReminder(t *testing.T) {

	_, channelId := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		createMessageReminder(t, apiClient, channelId)
	})
}

func Test_DeleteReminder(t *testing.T) {

	_, channelId := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		reminderId := createMessageReminder(t, apiClient, channelId)

		resp, httpResp, err := apiClient.DeleteReminder(ctx, reminderId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetReminders(t *testing.T) {

	_, channelId := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		reminderId := createMessageReminder(t, apiClient, channelId)

		resp, httpResp, err := apiClient.GetReminders(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		assert.GreaterOrEqual(t, len(resp.Reminders), 1)
		found := false
		for _, r := range resp.Reminders {
			if r.ReminderId == reminderId {
				require.WithinDuration(t, time.Now().Add(1*time.Hour), r.ScheduledDeliveryTimestamp, 3*time.Minute)
				found = true
			}
		}
		assert.True(t, found, "Created reminder not found in list of reminders")
	})
}

func createMessageReminder(t *testing.T, apiClient client.Client, channelId int64) int64 {
	msg := CreateChannelMessage(t, apiClient, channelId)

	note := "This is a reminder note"

	resp, httpResp, err := apiClient.CreateMessageReminder(context.Background()).
		MessageId(msg.MessageId).
		Note(UniqueName(note)).
		ScheduledDeliveryTimestamp(time.Now().Add(1 * time.Hour)).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpResp.StatusCode)

	return resp.ReminderId
}
