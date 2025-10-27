package scheduled_messages_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/scheduled_messages"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_CreateScheduledMessage(t *testing.T) {

	otherUserID := GetUserID(t, GetOtherNormalClient(t))

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		createScheduledMessage(t, apiClient, z.UserAsRecipient(otherUserID))
	})
}

func Test_DeleteScheduledMessage(t *testing.T) {

	otherUserID := GetUserID(t, GetOtherNormalClient(t))

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, z.UserAsRecipient(otherUserID))

		resp, httpResp, err := apiClient.DeleteScheduledMessage(ctx, msg.ScheduledMessageID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func Test_GetScheduledMessages(t *testing.T) {

	otherUserID := GetUserID(t, GetOtherNormalClient(t))

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, z.UserAsRecipient(otherUserID))

		resp, httpResp, err := apiClient.GetScheduledMessages(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)
		assert.GreaterOrEqual(t, len(resp.ScheduledMessages), 1)
		found := false
		for _, m := range resp.ScheduledMessages {
			if m.ScheduledMessageID == msg.ScheduledMessageID {
				found = true
				require.WithinDuration(t, time.Now().Add(1*time.Hour), m.ScheduledDeliveryTimestamp, 3*time.Minute)
				break
			}
		}
		assert.True(t, found, "created scheduled message not found in list")
	})
}

func Test_UpdateScheduledMessage(t *testing.T) {

	otherUserID := GetUserID(t, GetOtherNormalClient(t))

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, z.UserAsRecipient(otherUserID))

		resp, httpResp, err := apiClient.
			UpdateScheduledMessage(ctx, msg.ScheduledMessageID).
			Content(UniqueName("Updated scheduled message")).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpResp.StatusCode)

	})
}

func createScheduledMessage(t *testing.T, apiClient client.Client, to z.Recipient) *scheduled_messages.CreateScheduledMessageResponse {
	ctx := context.Background()

	resp, httpResp, err := apiClient.CreateScheduledMessage(ctx).
		Content(UniqueName("")).
		To(to).
		ScheduledDeliveryTimestamp(time.Now().Add(1 * time.Hour)).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, httpResp.StatusCode)
	return resp
}
