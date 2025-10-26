package zulip_test

import (
	"context"

	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
)

func Example() {
	// Load configuration from zuliprc file and create a client
	rc, _ := z.NewZulipRCFromFile("~/.zuliprc")
	client, _ := client.NewClient(rc)

	ctx := context.Background()

	// Get own user id
	userResp, _, _ := client.GetOwnUser(ctx).Execute()
	userId := userResp.UserId

	// Create a channel
	// older versions need to use the more verbose Subscribe() endpoint
	channelResp, _, _ := client.CreateChannel(ctx).
		Name("zulip-community").
		Description("A channel to greet the Zulip community").
		Subscribers([]int64{userId}).
		Execute()

	// Send a message greeting the Zulip community
	client.SendMessage(ctx).
		To(z.ChannelAsRecipient(channelResp.Id)).
		Topic("introductions").
		Content("Hello Zulip community! 👋").
		Execute()
}
