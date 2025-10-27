package events

import "github.com/tum-zulip/go-zulip/zulip"

// DeleteMessageEvent Event sent when a message has been deleted.  Sent to all users who currently are subscribed to the messages' recipient. May also be sent to additional users who had access to it, including, in particular, an administrator user deleting messages in a stream that they are not subscribed to.  This means that clients can assume that they will always receive an event of this type for deletions that the client itself initiated.  This event is also sent when the user loses access to a message, such as when it is [moved to a channel] that the user does not [have permission to access].
//
// **Changes**: Before Zulip 9.0 (feature level 274), this event was only sent to subscribers of the message's recipient.  Before Zulip 5.0 (feature level 77), events for direct messages contained additional `sender_id` and `recipient_id` fields.
//
// [moved to a channel]: https://zulip.com/help/move-content-to-another-channel
// [have permission to access]: https://zulip.com/help/channel-permissions
type DeleteMessageEvent struct {
	event
	// Only present for clients that support the `bulk_message_deletion` [client capability].  A sorted list containing the IDs of the newly deleted messages.
	//
	// **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	MessageIDs []int64 `json:"message_ids,omitempty"`
	// Only present for clients that do not support the `bulk_message_deletion`
	//
	// [client capability].  The ID of the newly deleted message.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	MessageID *int64 `json:"message_id,omitempty"`

	// The type of message. Either `RecipientTypeStream` or `RecipientTypePrivate`.
	MessageType zulip.RecipientType `json:"message_type,omitempty"`
	// Only present if `message_type` is `"stream"`.  The ID of the channel to which the message was sent.
	ChannelID *int64 `json:"stream_id,omitempty"`
	// Only present if `message_type` is `"stream"`.  The topic to which the message was sent.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name was empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
}
