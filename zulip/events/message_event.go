package events

import "github.com/tum-zulip/go-zulip/zulip"

// MessageEvent Event type for messages.
// **Changes**: In Zulip 3.1 (feature level 26), the `sender_short_name` field was removed from message objects.
type MessageEvent struct {
	event

	Message zulip.Message `json:"message,omitempty"`
	// The user's [message flags] for the message.  Clients should inspect the flags field rather than assuming that new messages are unread; [muted users], messages sent by the current user, and more subtle scenarios can result in a new message that the server has already marked as read for the user.
	//
	// **Changes**: In Zulip 8.0 (feature level 224), the `wildcard_mentioned` flag was deprecated in favor of the `stream_wildcard_mentioned` and `topic_wildcard_mentioned` flags. The `wildcard_mentioned` flag exists for backwards compatibility with older clients and equals `stream_wildcard_mentioned || topic_wildcard_mentioned`. Clients supporting older server versions should treat this field as a previous name for the `stream_wildcard_mentioned` flag as topic wildcard mentions were not available prior to this feature level.
	//
	// [message flags]: https://zulip.com/api/update-message-flags#available-flags
	//
	// [muted users]: https://zulip.com/api/mute-user
	Flags []string `json:"flags,omitempty"`
}
