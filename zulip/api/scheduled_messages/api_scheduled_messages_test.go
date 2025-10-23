package scheduled_messages_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
)

func Test_ScheduledMessagesAPIService(t *testing.T) {
	t.Parallel()

	otherUserId := getOwnUserId(t, GetOtherNormalClient(t))

	t.Run("CreateScheduledMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		createScheduledMessage(t, apiClient, []int64{otherUserId})
	}))

	t.Run("DeleteScheduledMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, []int64{otherUserId})

		resp, httpRes, err := apiClient.DeleteScheduledMessage(ctx, msg.ScheduledMessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetScheduledMessages", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, []int64{otherUserId})

		resp, httpRes, err := apiClient.GetScheduledMessages(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.GreaterOrEqual(t, len(resp.ScheduledMessages), 1)
		found := false
		for _, m := range resp.ScheduledMessages {
			if m.ScheduledMessageId == msg.ScheduledMessageId {
				found = true
				break
			}
		}
		assert.True(t, found, "created scheduled message not found in list")
	}))

	t.Run("UpdateScheduledMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, []int64{otherUserId})

		resp, httpRes, err := apiClient.
			UpdateScheduledMessage(ctx, msg.ScheduledMessageId).
			Content(uniqueName("Updated scheduled message")).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))
}

func createScheduledMessage(t *testing.T, apiClient z.Client, to []int64) *z.CreateScheduledMessageResponse {
	ctx := context.Background()

	resp, httpRes, err := apiClient.CreateScheduledMessage(ctx).
		Content(uniqueName("This is a scheduled message")).
		To(z.UsersAsRecipient(to)).
		ScheduledDeliveryTimestamp(time.Now().Add(1 * time.Hour)).
		RecipientType(z.RecipientTypeDirect).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, httpRes.StatusCode)
	return resp
}
