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
)

func Test_MessagesAPIService(t *testing.T) {

	otherClient := GetOtherNormalClient(t)

	channelName, channelId := createChannelWithAllClients(t)

	t.Parallel()

	t.Run("AddReaction", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.AddReaction(ctx, msg.messageId).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("CheckMessagesMatchNarrow", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.CheckMessagesMatchNarrow(ctx).
			MsgIds([]int64{msg.messageId}).
			Narrow(
				z.Where(z.ChannelNameIs(channelName)).
					And(z.TopicIs(msg.topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)

		key := strconv.Itoa(int(msg.messageId))
		assert.Contains(t, resp.Messages, key)
	}))

	t.Run("DeleteMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.DeleteMessage(ctx, msg.messageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("GetFileTemporaryUrl", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		upload := uploadFileForTest(t, ctx, apiClient)

		realmId, filename := parseUploadedFilePath(t, upload.Url)

		resp, httpRes, err := apiClient.GetFileTemporaryUrl(ctx, realmId, filename).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Url)
	}))

	t.Run("GetMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetMessage(ctx, msg.messageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Equal(t, msg.messageId, resp.Message.Id)
	}))

	t.Run("GetMessageHistory", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		updateContent := fmt.Sprintf("updated %s", uniqueName("message"))
		_, _, err := apiClient.UpdateMessage(ctx, msg.messageId).
			Content(updateContent).
			Execute()
		require.NoError(t, err)

		resp, httpRes, err := apiClient.GetMessageHistory(ctx, msg.messageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.MessageHistory)
	}))

	t.Run("GetMessages", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetMessages(ctx).
			Anchor(strconv.Itoa(int(msg.messageId))).
			IncludeAnchor(true).
			NumBefore(0).
			NumAfter(0).
			Narrow(z.Where(z.ChannelNameIs(channelName)).
				And(z.TopicIs(msg.topic))).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.NotEmpty(t, resp.Messages)
	}))

	t.Run("GetReadReceipts", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.GetReadReceipts(ctx, msg.messageId).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("MarkAllAsRead", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp, httpRes, err := apiClient.MarkAllAsRead(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("MarkChannelAsRead", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.MarkChannelAsRead(ctx).
			ChannelId(msg.channelId).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("MarkTopicAsRead", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, otherClient, channelId)

		// mark the topic as read using the same client that created the message,
		// since the topic may not be visible to other clients immediately.
		resp, httpRes, err := otherClient.MarkTopicAsRead(ctx).
			ChannelId(msg.channelId).
			TopicName(msg.topic).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("RemoveReaction", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		_, _, err := apiClient.AddReaction(ctx, msg.messageId).
			EmojiName("smile").
			Execute()
		require.NoError(t, err)

		resp, httpRes, err := apiClient.RemoveReaction(ctx, msg.messageId).
			EmojiName("smile").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("RenderMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		content := "**bold** _italic_"
		resp, httpRes, err := apiClient.RenderMessage(ctx).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Contains(t, resp.Rendered, "<strong>bold</strong>")
	}))

	t.Run("ReportMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.ReportMessage(ctx, msg.messageId).
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
		requireStatusOK(t, httpRes)
	}))

	t.Run("SendMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		topic := uniqueName("topic")
		content := fmt.Sprintf("message sent via client %s", uniqueName("content"))

		resp, httpRes, err := apiClient.SendMessage(ctx).
			RecipientType(z.RecipientTypeChannel).
			To(z.ChannelAsRecipient(channelId)).
			Topic(topic).
			Content(content).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Greater(t, resp.Id, int64(0))
	}))

	t.Run("UpdateMessage", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)
		newContent := fmt.Sprintf("edited %s", uniqueName("content"))

		resp, httpRes, err := apiClient.UpdateMessage(ctx, msg.messageId).
			Content(newContent).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
	}))

	t.Run("UpdateMessageFlags", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateMessageFlags(ctx).
			Messages([]int64{msg.messageId}).
			Op("add").
			Flag("starred").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.Contains(t, resp.Messages, msg.messageId)
	}))

	t.Run("UpdateMessageFlagsForNarrow", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		msg := createChannelMessage(t, apiClient, channelId)

		resp, httpRes, err := apiClient.UpdateMessageFlagsForNarrow(ctx).
			Anchor(strconv.Itoa(int(msg.messageId))).
			NumBefore(0).
			NumAfter(0).
			IncludeAnchor(true).
			Narrow(
				z.Where(z.ChannelNameIs(channelName)).
					And(z.TopicIs(msg.topic))).
			Op("add").
			Flag("starred").
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		requireStatusOK(t, httpRes)
		assert.GreaterOrEqual(t, resp.ProcessedCount, int64(1))
	}))

	t.Run("UploadFile", runForAllClients(t, func(t *testing.T, apiClient z.Client) {
		ctx := context.Background()

		resp := uploadFileForTest(t, ctx, apiClient)
		assert.NotEmpty(t, resp.Url)
		assert.NotEmpty(t, resp.Filename)
	}))
}

type channelMessage struct {
	channelId int64
	topic     string
	messageId int64
}

func createChannelMessage(t *testing.T, apiClient z.Client, channelId int64) channelMessage {
	t.Helper()

	topic := uniqueName("topic")
	content := fmt.Sprintf("automated test message %s", uniqueName("content"))
	messageId := sendChannelMessage(t, apiClient, channelId, topic, content)
	assert.Greater(t, messageId, int64(0))

	return channelMessage{
		channelId: channelId,
		topic:     topic,
		messageId: messageId,
	}
}

func uploadFileForTest(t *testing.T, ctx context.Context, apiClient z.Client) *z.UploadFileResponse {
	t.Helper()

	tmp, err := os.CreateTemp("", "z.upload-*.txt")
	require.NoError(t, err)
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	_, err = tmp.WriteString("uploaded from automated test")
	require.NoError(t, err)
	_, err = tmp.Seek(0, 0)
	require.NoError(t, err)

	resp, httpRes, err := apiClient.UploadFile(ctx).
		Filename(tmp).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)

	return resp
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

func createDirectMessage(t *testing.T, apiClient z.Client, to int64) int64 {
	t.Helper()

	content := uniqueName("content")
	resp, httpRes, err := apiClient.SendMessage(context.Background()).
		RecipientType(z.RecipientTypePrivate).
		To(z.UserAsRecipient(to)).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	requireStatusOK(t, httpRes)
	require.Greater(t, resp.Id, int64(0))

	return resp.Id
}
