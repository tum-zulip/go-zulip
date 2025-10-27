package events

import "github.com/tum-zulip/go-zulip/zulip"

// ReactionEvent Event sent when a reaction is added to or removed from a message. Sent to all users who were recipients of the message.
type ReactionEvent struct {
	event

	zulip.EmojiReaction

	// The ID of the message to which a reaction was added or removed.
	MessageID int64 `json:"message_id,omitempty"`

	// Deprecated
	User interface{} `json:"user,omitempty"`
}
