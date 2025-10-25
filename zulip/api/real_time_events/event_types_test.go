package real_time_events_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
)

func Test_MessageEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		content := "Test message content"

		// Register queue for message events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		// Trigger message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			sendDirectMessage(t, otherClient, GetUserId(t, apiClient), content)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.MessageEvent{}, event)
		msgEvent := event.(z.MessageEvent)
		assert.Contains(t, msgEvent.Message.Content, content)

		// Cleanup
		_, _, _ = apiClient.DeleteQueue(ctx).QueueId(*queueResp.QueueId).Execute()
	})
}

func Test_ReactionEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageId := sendDirectMessage(t, otherClient, GetUserId(t, apiClient), "Message to react to")

		// Register queue for reaction events with narrow to this specific message
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeReaction}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		emojiName := "smile"

		// Trigger reaction event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := otherClient.AddReaction(ctx, messageId).
				EmojiName(emojiName).
				Execute()
			require.NoError(t, err)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.ReactionEvent{}, event)
		reactionEvent := event.(z.ReactionEvent)
		assert.Equal(t, messageId, reactionEvent.MessageId)
		assert.Equal(t, emojiName, reactionEvent.EmojiName)
		assert.Equal(t, z.EventOpAdd, reactionEvent.Op)

	})
}

func Test_UpdateMessageEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageId := sendDirectMessage(t, apiClient, GetUserId(t, otherClient), "Original content")

		// Register queue for update_message events with narrow
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeUpdateMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		newContent := "Updated content"

		// Trigger update message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := apiClient.UpdateMessage(ctx, messageId).
				Content(newContent).
				Execute()
			require.NoError(t, err)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.UpdateMessageEvent{}, event)
		updateEvent := event.(z.UpdateMessageEvent)
		assert.Equal(t, messageId, updateEvent.MessageId)
		assert.NotNil(t, updateEvent.Content)
		assert.Contains(t, *updateEvent.Content, newContent)

	})
}

func Test_DeleteMessageEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageId := sendDirectMessage(t, apiClient, GetUserId(t, otherClient), "Message to delete")

		// Register queue for delete_message events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeDeleteMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		// Trigger delete message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := apiClient.DeleteMessage(ctx, messageId).Execute()
			require.NoError(t, err)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.DeleteMessageEvent{}, event)
		deleteEvent := event.(z.DeleteMessageEvent)
		// Check if message ID is in the list
		if deleteEvent.MessageIds != nil {
			assert.Contains(t, deleteEvent.MessageIds, messageId)
		} else if deleteEvent.MessageId != nil {
			assert.Equal(t, messageId, *deleteEvent.MessageId)
		}

	})
}

func Test_TypingEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Register queue for typing events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeTyping}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		// Trigger typing event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := otherClient.SetTypingStatus(ctx).
				Op(z.TypingStatusOpStart).
				To(z.UserAsRecipient(GetUserId(t, apiClient))).
				Execute()
			require.NoError(t, err)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.TypingEvent{}, event)
		typingEvent := event.(z.TypingEvent)
		assert.Equal(t, z.EventOpStart, typingEvent.Op)

	})
}

func Test_UpdateMessageFlagsEvent(t *testing.T) {
	t.Parallel()

	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageId := sendDirectMessage(t, otherClient, GetUserId(t, apiClient), "Message to star")

		// Register queue for update_message_flags events
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]z.EventType{z.EventTypeUpdateMessageFlags}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueId)

		// Trigger flag update event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, err := apiClient.UpdateMessageFlags(ctx).
				Messages([]int64{messageId}).
				Op("add").
				Flag("starred").
				Execute()
			require.NoError(t, err)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueId(*queueResp.QueueId).
			LastEventId(queueResp.LastEventId).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, z.UpdateMessageFlagsAddEvent{}, event)
		flagEvent := event.(z.UpdateMessageFlagsAddEvent)
		assert.Equal(t, z.EventOpAdd, flagEvent.Op)
		assert.Equal(t, "starred", flagEvent.Flag)
		assert.Contains(t, flagEvent.Messages, messageId)
	})
}

// Helper function to send a direct message and return message ID
func sendDirectMessage(t *testing.T, apiClient client.Client, toUserId int64, content string) int64 {
	t.Helper()

	ctx := context.Background()
	resp, _, err := apiClient.SendMessage(ctx).
		To(z.UserAsRecipient(toUserId)).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp.Id
}
