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

func Test_ScheduledMessagesAPIService(t *testing.T) {
	t.Parallel()

	otherUserId := GetUserId(t, GetOtherNormalClient(t))

	t.Run("CreateScheduledMessage", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		createScheduledMessage(t, apiClient, []int64{otherUserId})
	}))

	t.Run("DeleteScheduledMessage", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, []int64{otherUserId})

		resp, httpRes, err := apiClient.DeleteScheduledMessage(ctx, msg.ScheduledMessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))

	t.Run("GetScheduledMessages", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
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

	t.Run("UpdateScheduledMessage", RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := createScheduledMessage(t, apiClient, []int64{otherUserId})

		resp, httpRes, err := apiClient.
			UpdateScheduledMessage(ctx, msg.ScheduledMessageId).
			Content(UniqueName("Updated scheduled message")).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}))
}

func createScheduledMessage(t *testing.T, apiClient client.Client, to []int64) *scheduled_messages.CreateScheduledMessageResponse {
	ctx := context.Background()

	resp, httpRes, err := apiClient.CreateScheduledMessage(ctx).
		Content(UniqueName("This is a scheduled message")).
		To(z.UsersAsRecipient(to)).
		ScheduledDeliveryTimestamp(time.Now().Add(1 * time.Hour)).
		RecipientType(z.RecipientTypeDirect).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, httpRes.StatusCode)
	return resp
}
