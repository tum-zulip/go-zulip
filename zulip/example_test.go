package zulip_test

import (
	"context"
	"log/slog"

	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
)

// Example demonstrates how to create a channel and send a message using the Zulip API client.
//
//nolint:testableexamples // .zuliprc is required for this example to work
func Example() {
	// Load configuration from zuliprc file and create a client
	rc, _ := z.NewZulipRCFromFile("~/.zuliprc")
	client, _ := client.NewClient(rc)

	ctx := context.Background()

	// Get own user id
	userResp, _, _ := client.GetOwnUser(ctx).Execute()
	userID := userResp.UserID

	// Create a channel
	// Older versions need to use the more verbose Subscribe() endpoint
	channelResp, _, err := client.CreateChannel(ctx).
		Name("zulip-community").
		Description("A channel to greet the Zulip community").
		Subscribers([]int64{userID}).
		Execute()
	if err != nil {
		slog.Error("failed to create channel", "error", err)
		return
	}

	// Send a message greeting the Zulip community
	_, _, err = client.SendMessage(ctx).
		To(z.ChannelAsRecipient(channelResp.ID)).
		Topic("introductions").
		Content("Hello Zulip community! 👋").
		Execute()
	if err != nil {
		slog.Error("failed to send message", "error", err)
	}
}
