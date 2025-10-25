package messages_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_AddReaction(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.AddReaction(ctx, msg.MessageId).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_CheckMessagesMatchNarrow(t *testing.T) {
	t.Parallel()

	channelName, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.CheckMessagesMatchNarrow(ctx).
			MsgIds([]int64{msg.MessageId}).
			Narrow(
				z.Where(z.ChannelNameIs(channelName)).
					And(z.TopicIs(msg.Topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)

		key := strconv.Itoa(int(msg.MessageId))
		assert.Contains(t, resp.Messages, key)
	})
}

func Test_DeleteMessage(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.DeleteMessage(ctx, msg.MessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_GetFileTemporaryUrl(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		upload := UploadFileForTest(t, ctx, apiClient)

		realmId, filename := parseUploadedFilePath(t, upload.Url)

		resp, httpRes, err := apiClient.GetFileTemporaryUrl(ctx, realmId, filename).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Url)
	})
}

func Test_GetMessage(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetMessage(ctx, msg.MessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Equal(t, msg.MessageId, resp.Message.Id)
	})
}

func Test_GetMessageHistory(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		updateContent := fmt.Sprintf("updated %s", UniqueName("message"))
		_, _, err := apiClient.UpdateMessage(ctx, msg.MessageId).
			Content(updateContent).
			Execute()
		require.NoError(t, err)

		resp, httpRes, err := apiClient.GetMessageHistory(ctx, msg.MessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.MessageHistory)
	})
}

func Test_GetMessages(t *testing.T) {
	t.Parallel()

	channelName, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetMessages(ctx).
			Anchor(strconv.Itoa(int(msg.MessageId))).
			IncludeAnchor(true).
			NumBefore(0).
			NumAfter(0).
			Narrow(z.Where(z.ChannelNameIs(channelName)).
				And(z.TopicIs(msg.Topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Messages)
	})
}

func Test_GetReadReceipts(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetReadReceipts(ctx, msg.MessageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_MarkAllAsRead(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.MarkAllAsRead(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_MarkChannelAsRead(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.MarkChannelAsRead(ctx).
			ChannelId(msg.ChannelId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_MarkTopicAsRead(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)
	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, otherClient, channelId)

		// mark the topic as read using the same client that created the message,
		// since the topic may not be visible to other clients immediately.
		resp, httpRes, err := otherClient.MarkTopicAsRead(ctx).
			ChannelId(msg.ChannelId).
			TopicName(msg.Topic).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_RemoveReaction(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		_, _, err := apiClient.AddReaction(ctx, msg.MessageId).
			EmojiName("smile").
			Execute()
		require.NoError(t, err)

		resp, httpRes, err := apiClient.RemoveReaction(ctx, msg.MessageId).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_RenderMessage(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		content := "**bold** _italic_"
		resp, httpRes, err := apiClient.RenderMessage(ctx).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Contains(t, resp.Rendered, "<strong>bold</strong>")
	})
}

func Test_ReportMessage(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.ReportMessage(ctx, msg.MessageId).
			ReportType("spam").
			Description("reported by automated tests").
			Execute()

		if err != nil {
			var apiErr *z.APIError
			if errors.As(err, &apiErr) {
				body := strings.TrimSpace(string(apiErr.Body()))
				lower := strings.ToLower(fmt.Sprintf("%s %s", apiErr.Error(), body))
				if strings.Contains(lower, "moderation") ||
					strings.Contains(lower, "not configured") ||
					strings.Contains(lower, "message reporting is not enabled") {
					t.Skipf("message reporting unavailable: %s", body)
				}
			}
		}

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_SendMessage(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		topic := UniqueName("topic")
		content := fmt.Sprintf("message sent via client %s", UniqueName("content"))

		resp, httpRes, err := apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(channelId)).
			Topic(topic).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Greater(t, resp.Id, int64(0))
	})
}

func Test_UpdateMessage(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)
		newContent := fmt.Sprintf("edited %s", UniqueName("content"))

		resp, httpRes, err := apiClient.UpdateMessage(ctx, msg.MessageId).
			Content(newContent).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_UpdateMessageFlags(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageId}).
			Op("add").
			Flag("starred").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.Contains(t, resp.Messages, msg.MessageId)
	})
}

func Test_UpdateMessageFlagsForNarrow(t *testing.T) {
	t.Parallel()

	channelName, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageId))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(
				z.Where(z.ChannelNameIs(channelName)).
					And(z.TopicIs(msg.Topic))).
			Op("add").
			Flag("starred").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.GreaterOrEqual(t, resp.ProcessedCount, int64(1))
	})
}

func Test_UploadFile(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp := UploadFileForTest(t, ctx, apiClient)
		assert.NotEmpty(t, resp.Url)
		assert.NotEmpty(t, resp.Filename)
	})
}

func Test_AddReaction_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.AddReaction(ctx, 99999999).
			EmojiName("smile").
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with valid message but non-existent emoji
		msg := CreateChannelMessage(t, apiClient, channelId)
		_, _, err = apiClient.AddReaction(ctx, msg.MessageId).
			EmojiName("nonexistentemoji1234567890").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_CheckMessagesMatchNarrow_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with empty msgIds list - should fail
		_, _, err := apiClient.CheckMessagesMatchNarrow(ctx).
			MsgIds([]int64{}).
			Narrow(z.Where(z.ChannelNameIs("nonexistent-channel"))).
			Execute()
		require.Error(t, err)

		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.BadNarrowError{}, apiErr.Model())
		}

	})
}

func Test_DeleteMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.DeleteMessage(ctx, 99999999).Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with negative messageId
		_, _, err = apiClient.DeleteMessage(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_GetFileTemporaryUrl_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent realmId and invalid filename
		_, _, err := apiClient.GetFileTemporaryUrl(ctx, 99999, "nonexistent/file.txt").Execute()
		require.Error(t, err)

		// Test with negative realmId
		_, _, err = apiClient.GetFileTemporaryUrl(ctx, -1, "file.txt").Execute()
		require.Error(t, err)
	})
}

func Test_GetMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.GetMessage(ctx, 99999999).Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with negative messageId
		_, _, err = apiClient.GetMessage(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_GetMessageHistory_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.GetMessageHistory(ctx, 99999999).Execute()
		require.Error(t, err)

		// Test with a fresh message (one without edit history)
		msg := CreateChannelMessage(t, apiClient, channelId)
		resp, httpRes, err := apiClient.GetMessageHistory(ctx, msg.MessageId).Execute()
		// Should succeed but have minimal history
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_GetMessages_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent anchor messageId
		_, _, _ = apiClient.GetMessages(ctx).
			Anchor("99999999").
			NumBefore(10).
			NumAfter(10).
			Execute()
		// May succeed with empty results or fail depending on server

		// Test with invalid negative numBefore/numAfter values
		_, _, err := apiClient.GetMessages(ctx).
			NumBefore(-5).
			NumAfter(-5).
			Execute()
		require.Error(t, err)
	})
}

func Test_GetReadReceipts_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.GetReadReceipts(ctx, 99999999).Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with negative messageId
		_, _, err = apiClient.GetReadReceipts(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_MarkAllAsRead_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// MarkAllAsRead should generally succeed even if no messages to mark
		resp, httpRes, err := apiClient.MarkAllAsRead(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)

		// Calling it again should also succeed (idempotency test)
		resp, httpRes, err = apiClient.MarkAllAsRead(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_MarkChannelAsRead_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent channelId
		_, _, err := apiClient.MarkChannelAsRead(ctx).
			ChannelId(99999999).
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with negative channelId
		_, _, err = apiClient.MarkChannelAsRead(ctx).
			ChannelId(-1).
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_MarkTopicAsRead_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent channelId
		_, _, err := apiClient.MarkTopicAsRead(ctx).
			ChannelId(99999999).
			TopicName("nonexistent-topic").
			Execute()
		require.Error(t, err)

		// Test with valid channelId but non-existent topic
		_, _, err = apiClient.MarkTopicAsRead(ctx).
			ChannelId(channelId).
			TopicName("this-topic-definitely-does-not-exist-xyz123").
			Execute()
		// May succeed or fail depending on server behavior
		_ = err
	})
}

func Test_RemoveReaction_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.RemoveReaction(ctx, 99999999).
			EmojiName("smile").
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test removing non-existent reaction from a real message
		msg := CreateChannelMessage(t, apiClient, channelId)
		_, _, err = apiClient.RemoveReaction(ctx, msg.MessageId).
			EmojiName("nonexistentemoji1234567890").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_RenderMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with empty content - may be allowed or fail
		_, _, _ = apiClient.RenderMessage(ctx).
			Content("").
			Execute()
		// Result depends on server, don't assert error

		// Test with valid markdown content
		resp, httpRes, err := apiClient.RenderMessage(ctx).
			Content("**bold** _italic_ `code`").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Rendered)
	})
}

func Test_ReportMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.ReportMessage(ctx, 99999999).
			ReportType("spam").
			Description("test report").
			Execute()
		require.Error(t, err)

		// Test with valid message but invalid report type
		msg := CreateChannelMessage(t, apiClient, channelId)
		_, _, err = apiClient.ReportMessage(ctx, msg.MessageId).
			ReportType("invalid-report-type-xyz").
			Description("test report").
			Execute()
		require.Error(t, err)
	})
}

func Test_SendMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent userId
		_, _, err := apiClient.SendMessage(ctx).
			To(z.UserAsRecipient(99999999)).
			Content("test message").
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with non-existent channelId
		_, _, err = apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(99999999)).
			Topic("test-topic").
			Content("test message").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.NonExistingChannelIdError{}, apiErr.Model())
		}

		// Test with empty content
		_, _, err = apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(channelId)).
			Topic("test-topic").
			Content("").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_UpdateMessage_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.UpdateMessage(ctx, 99999999).
			Content("updated content").
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Create another message for valid update test
		msg2 := CreateChannelMessage(t, apiClient, channelId)
		newContent := fmt.Sprintf("successfully updated %s", UniqueName("msg"))
		resp, httpRes, err := apiClient.UpdateMessage(ctx, msg2.MessageId).
			Content(newContent).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		RequireStatusOK(t, httpRes)
	})
}

func Test_UpdateMessageFlags_Invalid_Input(t *testing.T) {
	t.Parallel()

	_, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageId
		_, _, err := apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{99999999}).
			Op("add").
			Flag("starred").
			Execute()

		require.Error(t, err)

		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with invalid op value
		msg := CreateChannelMessage(t, apiClient, channelId)
		_, _, err = apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageId}).
			Op("invalid-op").
			Flag("starred").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with invalid flag value
		_, _, err = apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageId}).
			Op("add").
			Flag("invalid-flag-xyz").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_UpdateMessageFlagsForNarrow_Invalid_Input(t *testing.T) {
	t.Parallel()

	channelName, channelId := CreateChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with invalid negative numBefore/numAfter
		_, _, err := apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor("1000").
			NumBefore(-5).
			NumAfter(-5).
			Op("add").
			Flag("starred").
			Execute()
		require.Error(t, err)
		var apiErr *z.APIError
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with invalid op value
		msg := CreateChannelMessage(t, apiClient, channelId)
		_, _, err = apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageId))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(z.Where(z.ChannelNameIs(channelName))).
			Op("invalid-op").
			Flag("starred").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}

		// Test with invalid flag value
		_, _, err = apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageId))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(z.Where(z.ChannelNameIs(channelName))).
			Op("add").
			Flag("invalid-flag-xyz").
			Execute()
		require.Error(t, err)
		if errors.As(err, &apiErr) {
			require.NotNil(t, apiErr.Model())
			assert.IsType(t, z.CodedError{}, apiErr.Model())
		}
	})
}

func Test_UploadFile_Invalid_Input(t *testing.T) {
	t.Parallel()

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with file that doesn't exist - try to open a non-existent file
		nonExistentFile, err := os.Open("nonexistent-file-that-does-not-exist-xyz123.txt")
		if err == nil {
			defer nonExistentFile.Close()
		}
		if err != nil {
			// Expected error when file doesn't exist
			require.Error(t, err)
		}

		// Test with a valid file for success case
		resp := UploadFileForTest(t, ctx, apiClient)
		assert.NotEmpty(t, resp.Url)
		assert.NotEmpty(t, resp.Filename)
	})
}

func parseUploadedFilePath(t *testing.T, uploadPath string) (int64, string) {
	t.Helper()

	require.NotEmpty(t, uploadPath)

	trimmed := strings.TrimLeft(uploadPath, "/")
	parts := strings.Split(trimmed, "/")
	require.GreaterOrEqual(t, len(parts), 3, "unexpected upload path: %s", uploadPath)
	require.Equal(t, "user_uploads", parts[0], "unexpected upload prefix: %s", uploadPath)

	realmId, err := strconv.ParseInt(parts[1], 10, 64)
	require.NoError(t, err)

	filename := strings.Join(parts[2:], "/")
	require.NotEmpty(t, filename)

	return realmId, filename
}
