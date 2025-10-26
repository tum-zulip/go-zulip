package zulip_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	z "github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/client"
	. "github.com/tum-zulip/go-zulip/zulip/internal/test_utils"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

// TestDocExampleBasicAPI demonstrates how to use the go-zulip client to create a channel
// and send a message greeting the Zulip community. This test serves as documentation
// for basic API usage and how to use the fluent builder pattern.
//
// This example shows:
// - Loading configuration from ~/.zuliprc
// - Creating a client
// - Using the builder pattern to create a channel
// - Using the builder pattern to send a message to the channel
func TestDocExampleBasicAPI(t *testing.T) {
	RequireFeatureLevel(t, 417)

	// During testing, we create a temporary zuliprc from the dev-only zulip endpoint.
	// Use path to your own zuliprc file in real usage.
	path := CreateIntegrationTestRC(t, TestAdminUsername)

	// 1. Load configuration from zuliprc file
	rc, err := zuliprc.NewZulipRCFromFile(path)
	require.NoError(t, err)

	// 2. Create a client
	client, err := client.NewClient(rc)
	require.NoError(t, err)
	require.NotNil(t, client)

	ctx := context.Background()

	// 3. get own user id
	userResp, _, err := client.GetOwnUser(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, userResp)
	userId := userResp.UserId
	require.Greater(t, userId, int64(0))

	// 4. Create a channel
	// older versions need to use the more verbose Subscribe() endpoint
	channelResp, _, err := client.CreateChannel(ctx).
		Name("zulip-community").
		Description("A channel to greet the Zulip community").
		Subscribers([]int64{userId}).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, channelResp)

	// Send a message greeting the Zulip community
	messageResp, _, err := client.SendMessage(ctx).
		To(z.ChannelAsRecipient(channelResp.Id)).
		Topic("introductions").
		Content("Hello Zulip community! 👋").
		Execute()

	require.NoError(t, err)
	require.NotNil(t, messageResp)
	require.Greater(t, messageResp.Id, int64(0))
}
