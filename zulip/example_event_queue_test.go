package zulip_test

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

// TestExampleEchoBot demonstrates how to use the event queue to build an echo bot
// that replies to every message sent to it by echoing the message content.
// This test serves as documentation for event-driven bot usage.
//
// Note: The test infrastructure here is intentionally hidden from documentation examples.
// The core bot logic shows how to:
// 1. Register an event queue for message events
// 2. Connect to the event queue
// 3. Listen for message events in a goroutine
// 4. Respond to each message by echoing it back
func TestExampleEchoBot(t *testing.T) {

	// During testing, we create a temporary zuliprc from the dev-only zulip endpoint.
	// Use path to your own zuliprc file in real usage.
	path := CreateIntegrationTestRC(t, TestAdminUsername)

	rc, err := zuliprc.NewZulipRCFromFile(path)
	require.NoError(t, err)

	client, err := client.NewClient(rc)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	// get own user id
	userResp, _, err := client.GetOwnUser(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, userResp)
	userId := userResp.UserId
	require.Greater(t, userId, int64(0))

	// Register an event queue to listen for message events
	// Use filtering to only receive direct messages to the bot and reduce server and client load
	registerResp, _, err := client.RegisterQueue(ctx).
		// only request minimal initial data
		FetchEventTypes([]events.EventType{events.EventTypeMessage}).
		// // only subscribe to message events
		EventTypes([]events.EventType{events.EventTypeMessage}).
		// // only for direct messages to the bot
		Narrow(z.Where(z.IsDirectMessage())).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, registerResp)
	require.NotNil(t, registerResp.QueueId)

	queueID := *registerResp.QueueId
	lastEventID := registerResp.LastEventId

	// Create an event queue handler with logging
	queue := real_time_events.NewEventQueue(client, nil)

	// Connect to the event queue
	channel, err := queue.Connect(ctx, queueID, lastEventID)
	require.NoError(t, err)
	require.NotNil(t, channel)

	defer func() {
		err := queue.Close()

		require.NoError(t, err)
	}()

	// send direct message from different user to trigger echo
	go func() {
		time.Sleep(1 * time.Second) // slight delay to ensure bot is listening
		sendMessage(t, "Hello, Echo Bot!", userId)
	}()

	// Process events
	timeout := time.After(5 * time.Second) // Example timeout
	for {
		select {
		case <-timeout:
			t.Fatalf("test timed out waiting for echo response")

		case event := <-channel:
			require.NotNil(t, event)

			// Check if this is a message event
			msgEvent, ok := event.(events.MessageEvent)
			require.True(t, ok, "expected MessageEvent, got %T", event)

			// Echo the message back
			_, _, err := client.SendMessage(ctx).
				To(z.UserAsRecipient(msgEvent.Message.SenderId)).
				Content(fmt.Sprintf("Echo: %s", msgEvent.Message.Content)).
				Execute()

			require.NoError(t, err)

			// Exit after processing one message for test purposes
			return
		}
	}
}

func sendMessage(t *testing.T, content string, userId int64) {
	other := GetOtherNormalClient(t)

	slog.Error("sending message to echo bot")
	resp, _, err := other.SendMessage(context.Background()).
		To(z.UserAsRecipient(userId)).
		Content(content).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
}
