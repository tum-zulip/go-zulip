package messages_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_AddReaction(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.AddReaction(ctx, msg.MessageID).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_CheckMessagesMatchNarrow(t *testing.T) {
	channelName, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.CheckMessagesMatchNarrow(ctx).
			MsgIDs([]int64{msg.MessageID}).
			Narrow(
				z.Where(z.ChannelNameIs(channelName)).
					And(z.TopicIs(msg.Topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)

		key := strconv.Itoa(int(msg.MessageID))
		assert.Contains(t, resp.Messages, key)
	})
}

func Test_DeleteMessage(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.DeleteMessage(ctx, msg.MessageID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_GetFileTemporaryURL(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		upload := UploadFileForTest(ctx, t, apiClient)

		realmID, filename := parseUploadedFilePath(t, upload.URL)

		resp, _, err := apiClient.GetFileTemporaryURL(ctx, realmID, filename).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.NotEmpty(t, resp.URL)
	})
}

func Test_GetMessage(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.GetMessage(ctx, msg.MessageID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, msg.MessageID, resp.Message.ID)
		require.WithinDuration(t, time.Now(), resp.Message.Timestamp, 3*time.Minute)
		// TODO: require.WithinDuration(t, time.Now(), resp.Message.LastEditTimestamp, 3*time.Minute)
		// TODO: require.WithinDuration(t, time.Now(), resp.Message.LastMovedTimestamp, 3*time.Minute)
	})
}

func Test_GetMessageHistory(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		updateContent := fmt.Sprintf("updated %s", UniqueName("message"))
		_, _, err := apiClient.UpdateMessage(ctx, msg.MessageID).
			Content(updateContent).
			Execute()
		require.NoError(t, err)

		resp, _, err := apiClient.GetMessageHistory(ctx, msg.MessageID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.NotEmpty(t, resp.MessageHistory)
	})
}

func Test_GetMessages(t *testing.T) {
	channelName, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.GetMessages(ctx).
			Anchor(strconv.Itoa(int(msg.MessageID))).
			IncludeAnchor(true).
			NumBefore(0).
			NumAfter(0).
			Narrow(z.Where(z.ChannelNameIs(channelName)).
				And(z.TopicIs(msg.Topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.NotEmpty(t, resp.Messages)
	})
}

func Test_GetReadReceipts(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.GetReadReceipts(ctx, msg.MessageID).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_MarkAllAsRead(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp, _, err := apiClient.MarkAllAsRead(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_MarkChannelAsRead(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.MarkChannelAsRead(ctx).
			ChannelID(msg.ChannelID).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_MarkTopicAsRead(t *testing.T) {
	otherClient := GetOtherNormalClient(t)
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, otherClient, channelID)

		// mark the topic as read using the same client that created the message,
		// since the topic may not be visible to other clients immediately.
		resp, _, err := apiClient.MarkTopicAsRead(ctx).
			ChannelID(msg.ChannelID).
			TopicName(msg.Topic).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_RemoveReaction(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		_, _, err := apiClient.AddReaction(ctx, msg.MessageID).
			EmojiName("smile").
			Execute()
		require.NoError(t, err)

		resp, _, err := apiClient.RemoveReaction(ctx, msg.MessageID).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_RenderMessage(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		content := "**bold** normal _italic_"
		resp, _, err := apiClient.RenderMessage(ctx).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Contains(t, resp.Rendered, "<strong>bold</strong>")
	})
}

func Test_ReportMessage(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		// TODO: Fix this test
		t.Skip()

		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.ReportMessage(ctx, msg.MessageID).
			ReportType("spam").
			Description("reported by automated tests").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_SendMessage(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		topic := UniqueName("topic")
		content := fmt.Sprintf("message sent via client %s", UniqueName("content"))

		resp, _, err := apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(channelID)).
			Topic(topic).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Positive(t, resp.ID)
	})
}

func Test_UpdateMessage(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)
		newContent := fmt.Sprintf("edited %s", UniqueName("content"))

		resp, _, err := apiClient.UpdateMessage(ctx, msg.MessageID).
			Content(newContent).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_UpdateMessageFlags(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageID}).
			Op("add").
			Flag("starred").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Contains(t, resp.Messages, msg.MessageID)
	})
}

func Test_UpdateMessageFlagsForNarrow(t *testing.T) {
	channelName, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		msg := CreateChannelMessage(t, apiClient, channelID)

		resp, _, err := apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageID))).
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
		assert.GreaterOrEqual(t, resp.ProcessedCount, int64(1))
	})
}

func Test_UploadFile(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		resp := UploadFileForTest(ctx, t, apiClient)

		assert.NotEmpty(t, resp.URI)
		if GetFeatureLevel(t) >= 272 {
			assert.NotEmpty(t, resp.URL)
		}
		if GetFeatureLevel(t) >= 285 {
			assert.NotEmpty(t, resp.Filename)
		}
	})
}

func Test_AddReaction_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.AddReaction(ctx, 99999999).
			EmojiName("smile").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with valid message but non-existent emoji
		msg := CreateChannelMessage(t, apiClient, channelID)
		_, _, err = apiClient.AddReaction(ctx, msg.MessageID).
			EmojiName("nonexistentemoji1234567890").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func Test_CheckMessagesMatchNarrow_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with empty msgIDs list - should fail
		_, _, err := apiClient.CheckMessagesMatchNarrow(ctx).
			MsgIDs([]int64{}).
			Narrow(z.Where(z.ChannelNameIs("nonexistent-channel"))).
			Execute()

		require.ErrorAs(t, err, &z.BadNarrowError{})
	})
}

func Test_DeleteMessage_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.DeleteMessage(ctx, 99999999).Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with negative messageID
		_, _, err = apiClient.DeleteMessage(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_GetFileTemporaryURL_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent realmID and invalid filename
		_, _, err := apiClient.GetFileTemporaryURL(ctx, 99999, "nonexistent/file.txt").Execute()
		require.Error(t, err)

		// Test with negative realmID
		_, _, err = apiClient.GetFileTemporaryURL(ctx, -1, "file.txt").Execute()
		require.Error(t, err)
	})
}

func Test_GetMessage_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.GetMessage(ctx, 99999999).Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with negative messageID
		_, _, err = apiClient.GetMessage(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_GetMessageHistory_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.GetMessageHistory(ctx, 99999999).Execute()
		require.Error(t, err)

		// Test with a fresh message (one without edit history)
		msg := CreateChannelMessage(t, apiClient, channelID)
		resp, _, err := apiClient.GetMessageHistory(ctx, msg.MessageID).Execute()
		// Should succeed but have minimal history
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_GetMessages_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent anchor messageID
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
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.GetReadReceipts(ctx, 99999999).Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with negative messageID
		_, _, err = apiClient.GetReadReceipts(ctx, -1).Execute()
		require.Error(t, err)
	})
}

func Test_MarkAllAsRead_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// MarkAllAsRead should generally succeed even if no messages to mark
		resp, _, err := apiClient.MarkAllAsRead(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)

		// Calling it again should also succeed (idempotency test)
		resp, _, err = apiClient.MarkAllAsRead(ctx).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_MarkChannelAsRead_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent channelID
		_, _, err := apiClient.MarkChannelAsRead(ctx).
			ChannelID(99999999).
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})

		// Test with negative channelID
		_, _, err = apiClient.MarkChannelAsRead(ctx).
			ChannelID(-1).
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func Test_MarkTopicAsRead_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent channelID
		_, _, err := apiClient.MarkTopicAsRead(ctx).
			ChannelID(99999999).
			TopicName("nonexistent-topic").
			Execute()
		require.Error(t, err)

		// Test with valid channelID but non-existent topic
		_, _, err = apiClient.MarkTopicAsRead(ctx).
			ChannelID(channelID).
			TopicName("this-topic-definitely-does-not-exist-xyz123").
			Execute()
		// May succeed or fail depending on server behavior
		_ = err
	})
}

func Test_RemoveReaction_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.RemoveReaction(ctx, 99999999).
			EmojiName("smile").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test removing non-existent reaction from a real message
		msg := CreateChannelMessage(t, apiClient, channelID)
		_, _, err = apiClient.RemoveReaction(ctx, msg.MessageID).
			EmojiName("nonexistentemoji1234567890").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func Test_RenderMessage_Invalid_Input(t *testing.T) {
	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with empty content - may be allowed or fail
		_, _, _ = apiClient.RenderMessage(ctx).
			Content("").
			Execute()
		// Result depends on server, don't assert error

		// Test with valid markdown content
		resp, _, err := apiClient.RenderMessage(ctx).
			Content("**bold** _italic_ `code`").
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.NotEmpty(t, resp.Rendered)
	})
}

func Test_ReportMessage_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.ReportMessage(ctx, 99999999).
			ReportType("spam").
			Description("test report").
			Execute()
		require.Error(t, err)

		// Test with valid message but invalid report type
		msg := CreateChannelMessage(t, apiClient, channelID)
		_, _, err = apiClient.ReportMessage(ctx, msg.MessageID).
			ReportType("invalid-report-type-xyz").
			Description("test report").
			Execute()
		require.Error(t, err)
	})
}

func Test_SendMessage_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent userID
		_, _, err := apiClient.SendMessage(ctx).
			To(z.UserAsRecipient(99999999)).
			Content("test message").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with non-existent channelID
		_, _, err = apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(99999999)).
			Topic("test-topic").
			Content("test message").
			Execute()
		require.ErrorAs(t, err, &z.NonExistingChannelIDError{})

		// Test with empty content
		_, _, err = apiClient.SendMessage(ctx).
			To(z.ChannelAsRecipient(channelID)).
			Topic("test-topic").
			Content("").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func Test_UpdateMessage_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.UpdateMessage(ctx, 99999999).
			Content("updated content").
			Execute()
		require.ErrorAs(t, err, &z.CodedError{})

		// Create another message for valid update test
		msg2 := CreateChannelMessage(t, apiClient, channelID)
		newContent := fmt.Sprintf("successfully updated %s", UniqueName("msg"))
		resp, _, err := apiClient.UpdateMessage(ctx, msg2.MessageID).
			Content(newContent).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}

func Test_UpdateMessageFlags_Invalid_Input(t *testing.T) {
	_, channelID := GetChannelWithAllClients(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Test with non-existent messageID
		_, _, err := apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{99999999}).
			Op("add").
			Flag("starred").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})

		// Test with invalid op value
		msg := CreateChannelMessage(t, apiClient, channelID)
		_, _, err = apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageID}).
			Op("invalid-op").
			Flag("starred").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})

		// Test with invalid flag value
		_, _, err = apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.MessageID}).
			Op("add").
			Flag("invalid-flag-xyz").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func Test_UpdateMessageFlagsForNarrow_Invalid_Input(t *testing.T) {
	channelName, channelID := GetChannelWithAllClients(t)

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
		require.ErrorAs(t, err, &z.CodedError{})

		// Test with invalid op value
		msg := CreateChannelMessage(t, apiClient, channelID)
		_, _, err = apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageID))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(z.Where(z.ChannelNameIs(channelName))).
			Op("invalid-op").
			Flag("starred").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})

		// Test with invalid flag value
		_, _, err = apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.MessageID))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(z.Where(z.ChannelNameIs(channelName))).
			Op("add").
			Flag("invalid-flag-xyz").
			Execute()

		require.ErrorAs(t, err, &z.CodedError{})
	})
}

func parseUploadedFilePath(t *testing.T, uploadPath string) (int64, string) {
	t.Helper()

	require.NotEmpty(t, uploadPath)

	trimmed := strings.TrimLeft(uploadPath, "/")
	parts := strings.Split(trimmed, "/")
	require.GreaterOrEqual(t, len(parts), 3, "unexpected upload path: %s", uploadPath)
	require.Equal(t, "user_uploads", parts[0], "unexpected upload prefix: %s", uploadPath)

	realmID, err := strconv.ParseInt(parts[1], 10, 64)
	require.NoError(t, err)

	filename := strings.Join(parts[2:], "/")
	require.NotEmpty(t, filename)

	return realmID, filename
}
