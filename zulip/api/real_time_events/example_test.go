package real_time_events_test

import (
	"context"
	"fmt"

	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/api/real_time_events"
	"github.com/tum-zulip/go-zulip/zulip/client"
	"github.com/tum-zulip/go-zulip/zulip/events"
)

// In this example, we use the event queue to build an echo bot
// that replies to every message sent to it by echoing the message content.
// This test serves as documentation for event-driven bot usage.
//
// The core bot logic shows how to:
// 1. Register an event queue for message events
// 2. Connect to the event queue
// 3. Listen for message events in a goroutine
// 4. Respond to each message by echoing it back
func Example() {
	rc, _ := z.NewZulipRCFromFile("~/.zuliprc")
	client, _ := client.NewClient(rc)

	ctx := context.Background()

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

	queueID := *registerResp.QueueId
	lastEventID := registerResp.LastEventId

	// Create an event queue handler for processing events
	queue := real_time_events.NewEventQueue(client, nil)

	// Connect to the event queue
	channel, err := queue.Connect(ctx, queueID, lastEventID)
	if err != nil {
		panic("failed to connect to event queue")
	}
	defer queue.Close()

	// Process events
	for event := range channel {
		// Check if this is a message event
		msgEvent, ok := event.(events.MessageEvent)
		if !ok {
			continue
		}

		// Echo the message back
		client.SendMessage(ctx).
			To(z.UserAsRecipient(msgEvent.Message.SenderId)).
			Content(fmt.Sprintf("Echo: %s", msgEvent.Message.Content)).
			Execute()
	}
}
