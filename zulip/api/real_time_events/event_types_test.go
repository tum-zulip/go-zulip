package realtimeevents_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/tum-zulip/go-zulip/internal/test_utils"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

func Test_MessageEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()
		content := "Test message content"

		// Register queue for message events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		userID := GetUserID(t, apiClient)
		senderID := GetUserID(t, otherClient)

		var sendErr error
		wait := make(chan struct{})

		// Trigger message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)

			_, _, sendErr = otherClient.SendMessage(context.Background()).
				To(z.UserAsRecipient(userID)).
				Content(content).
				Execute()
			close(wait)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		msgEvent, ok := event.(events.MessageEvent)
		require.True(t, ok, "event is not of type MessageEvent")
		assert.Equal(t, senderID, msgEvent.Message.SenderID)
		assert.Contains(t, msgEvent.Message.Content, content)

		<-wait
		require.NoError(t, sendErr)
	})
}

func Test_ReactionEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageID := sendDirectMessage(t, otherClient, GetUserID(t, apiClient), "Message to react to")

		// Register queue for reaction events with narrow to this specific message
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeReaction}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		emojiName := "smile"

		var addErr error
		wait := make(chan struct{})

		// Trigger reaction event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, addErr = otherClient.AddReaction(ctx, messageID).
				EmojiName(emojiName).
				Execute()
			close(wait)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]

		reactionEvent, ok := event.(events.ReactionEvent)
		require.True(t, ok, "event is not of type ReactionEvent")

		assert.Equal(t, messageID, reactionEvent.MessageID)
		assert.Equal(t, emojiName, reactionEvent.EmojiName)
		assert.Equal(t, events.EventOpAdd, reactionEvent.Op)

		<-wait
		require.NoError(t, addErr)
	})
}

func Test_UpdateMessageEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageID := sendDirectMessage(t, apiClient, GetUserID(t, otherClient), "Original content")

		// Register queue for update_message events with narrow
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeUpdateMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		newContent := "Updated content"

		var updateErr error
		wait := make(chan struct{})

		// Trigger update message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, updateErr = apiClient.UpdateMessage(ctx, messageID).
				Content(newContent).
				Execute()
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, events.UpdateMessageEvent{}, event)
		updateEvent, ok := event.(events.UpdateMessageEvent)
		require.True(t, ok, "event is not of type UpdateMessageEvent")

		assert.Equal(t, messageID, updateEvent.MessageID)
		assert.NotNil(t, updateEvent.Content)
		assert.Contains(t, *updateEvent.Content, newContent)

		<-wait
		require.NoError(t, updateErr)
	})
}

func Test_DeleteMessageEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageID := sendDirectMessage(t, apiClient, GetUserID(t, otherClient), "Message to delete")

		// Register queue for delete_message events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeDeleteMessage}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		var delErr error
		wait := make(chan struct{})

		// Trigger delete message event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, delErr = apiClient.DeleteMessage(ctx, messageID).Execute()
			close(wait)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		deleteEvent, ok := event.(events.DeleteMessageEvent)
		require.True(t, ok, "event is not of type DeleteMessageEvent")

		// Check if message ID is in the list
		if deleteEvent.MessageIDs != nil {
			assert.Contains(t, deleteEvent.MessageIDs, messageID)
		} else if deleteEvent.MessageID != nil {
			assert.Equal(t, messageID, *deleteEvent.MessageID)
		}

		<-wait
		require.NoError(t, delErr)
	})
}

func Test_TypingEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Register queue for typing events with narrow to direct messages
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeTyping}).
			Narrow(z.Where(z.IsDirectMessage())).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		// Trigger typing event in goroutine
		var typingErr error
		wait := make(chan struct{})

		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, typingErr = otherClient.SetTypingStatus(ctx).
				Op(z.TypingStatusOpStart).
				To(z.UserAsRecipient(GetUserID(t, apiClient))).
				Execute()
			close(wait)
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		typingEvent, ok := event.(events.TypingEvent)
		require.True(t, ok, "event is not of type TypingEvent")
		assert.Equal(t, events.EventOpStart, typingEvent.Op)

		<-wait
		require.NoError(t, typingErr)
	})
}

func Test_UpdateMessageFlagsEvent(t *testing.T) {
	otherClient := GetOtherNormalClient(t)

	RunForAllClients(t, func(t *testing.T, apiClient client.Client) {
		ctx := context.Background()

		// Setup: Send a direct message first
		messageID := sendDirectMessage(t, otherClient, GetUserID(t, apiClient), "Message to star")

		// Register queue for update_message_flags events
		queueResp, _, err := apiClient.RegisterQueue(ctx).
			EventTypes([]events.EventType{events.EventTypeUpdateMessageFlags}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, queueResp.QueueID)

		var starErr error
		wait := make(chan struct{})

		// Trigger flag update event in goroutine
		go func() {
			time.Sleep(200 * time.Millisecond)
			_, _, starErr = apiClient.UpdateMessageFlags(ctx).
				Messages([]int64{messageID}).
				Op("add").
				Flag("starred").
				Execute()
		}()

		// Get events
		resp, _, err := apiClient.GetEvents(ctx).
			QueueID(*queueResp.QueueID).
			LastEventID(queueResp.LastEventID).
			Execute()
		require.NoError(t, err)
		require.Len(t, resp.Events, 1)

		event := resp.Events[0]
		require.IsType(t, events.UpdateMessageFlagsAddEvent{}, event)
		flagEvent, eventOk := event.(events.UpdateMessageFlagsAddEvent)
		require.True(t, eventOk, "event is not of type UpdateMessageFlagsAddEvent")
		assert.Equal(t, events.EventOpAdd, flagEvent.Op)
		assert.Equal(t, "starred", flagEvent.Flag)
		assert.Contains(t, flagEvent.Messages, messageID)

		<-wait
		require.NoError(t, starErr)
	})
}

// Helper function to send a direct message and return message ID.
func sendDirectMessage(t *testing.T, apiClient client.Client, toUserID int64, content string) int64 {
	t.Helper()

	ctx := context.Background()
	resp, _, err := apiClient.SendMessage(ctx).
		To(z.UserAsRecipient(toUserID)).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp.ID
}
