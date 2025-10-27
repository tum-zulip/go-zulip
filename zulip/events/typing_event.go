package events

import "github.com/tum-zulip/go-zulip/zulip"

// TypingStartEvent Event sent when a user starts typing a message.  Sent to all clients for users who would receive the message being typed, with the additional rule that typing notifications for channel messages are only sent to clients that included `stream_typing_notifications` in their [client capabilities] when registering the event queue.  See [POST /typing] endpoint for more details.
//
// **Changes**: Typing notifications for channel messages are new in Zulip 4.0 (feature level 58).
//
// [client capabilities]: https://zulip.com/api/register-queue#parameter-client_capabilities
// [POST /typing]: https://zulip.com/api/set-typing-status
type TypingEvent struct {
	event
	// Type of message being composed. Must be `RecipientTypeStream` or `RecipientTypeDirect`.
	//
	// **Changes**: In Zulip 8.0 (feature level 215), replaced the value `RecipientTypePrivate` with `RecipientTypeDirect`.  New in Zulip 4.0 (feature level 58). Previously, all typing notifications were implicitly direct messages.
	MessageType zulip.RecipientType `json:"message_type,omitempty"`
	Sender      UserIdentifier      `json:"sender,omitempty"`
	// Only present if `message_type` is `RecipientTypeDirect`.  Array of dictionaries describing the set of users who would be recipients of the message being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.
	Recipients []UserIdentifier `json:"recipients,omitempty"`
	// Only present if `message_type` is `RecipientTypeStream`.  The unique ID of the channel to which message is being typed.
	//
	// **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
	ChannelID *int64 `json:"stream_id,omitempty"`
	// Only present if `message_type` is `RecipientTypeStream`.  Topic within the channel where the message is being typed.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
}

// UserIdentifier Object containing the user ID and Zulip API email of a recipient.
type UserIdentifier struct {
	// The ID of the user.
	UserID int64 `json:"user_id,omitempty"`
	// The Zulip API email address for the user.
	Email string `json:"email,omitempty"`
}
