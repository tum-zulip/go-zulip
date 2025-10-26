# Go API client for go-zulip

> **WARNING** This is still under active development and the go-api is not stable (yet)

## Getting Started
To use the go-zulip client in your Go project, you can install it via:

```bash
go get github.com/tum-zulip/go-zulip/zulip
```

Then, you can create a new client and start making API calls:

```go
import (
    "context"
    z "github.com/tum-zulip/go-zulip/zulip"
)

func main() {
	// Load configuration from zuliprc file and create a client
	rc, _ := z.NewZulipRCFromFile(".zuliprc")
	client, _ := client.NewClient(rc)

	ctx := context.Background()

	// Get own user id
	userResp, _, _ := client.GetOwnUser(ctx).Execute()
	userId := userResp.UserId

	// Create a channel
	// Older versions need to use the more verbose Subscribe() endpoint
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
```
