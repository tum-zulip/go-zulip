package events

import "github.com/tum-zulip/go-zulip/zulip"

// UpdateMessageFlagsAddEvent Event sent to a user when [message flags] are added to messages.  This can reflect a direct user action, or can be the indirect consequence of another action. Whatever the cause, if there's a change in the set of message flags that the user has for a message, then an `update_message_flags` event will be sent with the change. Note that this applies when the user already had access to the message, and continues to have access to it. When a message newly appears or disappears, a [`message`] or [`delete_message`] event is sent instead.  Some examples of actions that trigger an `update_message_flags` event:  - The `"starred"` flag is added when the user chooses to [star a   message]. - The `"read"` flag is added when the user marks messages as read by   scrolling through them, or uses [Mark all messages as read] on a conversation. - The `"read"` flag is added when the user [mutes] a   message's sender. - The `"read"` flag is added after the user unsubscribes from a channel,   or messages are moved to a not-subscribed channel, provided the user   can still access the messages at all. Note a   [`delete_message`] event is sent in the case where the   user can no longer access the messages.  In some cases, a change in message flags that's caused by another change may happen a short while after the original change, rather than simultaneously. For example, when messages that were unread are moved to a channel where the user is not subscribed, the resulting change in message flags (and the corresponding `update_message_flags` event with flag `"read"`) may happen later than the message move itself. The delay in that example is typically at most a few hundred milliseconds and can in rare cases be minutes or longer.
//
// [message flags]: https://zulip.com/api/update-message-flags#available-flags
// [`message`]: https://zulip.com/api/get-events#message
// [`delete_message`]: https://zulip.com/api/get-events#delete_message
// [Mark all messages as read]: https://zulip.com/help/marking-messages-as-read#mark-messages-in-multiple-topics-and-channels-as-read
// [star a   message]: https://zulip.com/help/star-a-message
// [mutes]: https://zulip.com/help/mute-a-user
type UpdateMessageFlagsAddEvent struct {
	event

	// Old name for the `op` field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the `op` field.
	// Deprecated
	Operation *string `json:"operation,omitempty"`
	// The [flag] that was added/removed.
	//
	// [flag]: https://zulip.com/api/update-message-flags#available-flags
	Flag string `json:"flag,omitempty"`
	// Array containing the Ids of all messages to which the flag was added/removed.
	Messages []int64 `json:"messages,omitempty"`
	// Whether the specified flag was added to all messages. This field is only relevant for the `"read"` flag, and will be `false` for all other flags.  When `true` for the `"read"` flag, then the `messages` array will be empty.
	All bool `json:"all,omitempty"`
}

// UpdateMessageFlagsRemoveEvent Event sent to a user when [message flags] are removed from messages.  See the description for the [`update_message_flags` op: `add`] event for more details about these events.
// [message flags]: https://zulip.com/api/update-message-flags#available-flags
// [`update_message_flags` op: `add`]: https://zulip.com/api/get-events#update_message_flags-add
type UpdateMessageFlagsRemoveEvent struct {
	event

	// Old name for the `op` field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the `op` field.
	// Deprecated
	Operation *string `json:"operation,omitempty"`
	// The [flag] to be removed.
	//
	// [flag]: https://zulip.com/api/update-message-flags#available-flags
	Flag string `json:"flag,omitempty"`
	// Array containing the Ids of the messages from which the flag was removed.
	Messages []int64 `json:"messages,omitempty"`
	// Will be `false` for all specified flags.  **Deprecated** and will be removed in a future release.
	// Deprecated
	All *bool `json:"all,omitempty"`

	// Only present if the specified `flag` is `"read"`.  A set of data structures describing the messages that are being marked as unread with additional details to allow clients to update the `unread_msgs` data structure for these messages (which may not be otherwise known to the client).
	//
	// **Changes**: New in Zulip 5.0 (feature level 121). Previously, marking already read messages as unread was not supported by the Zulip API.
	MessageDetails map[string]MessageDetail `json:"message_details,omitempty"`
}

// UpdateMessageFlagsRemoveEventMessageDetailsValue `{message_id}`: Object containing details about the message with the specified Id.
type MessageDetail struct {
	// The type of this message. Either `RecipientTypeStream` or `RecipientTypePrivate`.
	Type zulip.RecipientType `json:"type"`
	// A flag which indicates whether the message contains a mention of the user.  Present only if the message mentions the current user.
	Mentioned *bool `json:"mentioned,omitempty"`
	// Present only if `type` is `RecipientTypePrivate`.  The user Ids of every recipient of this direct message, excluding yourself. Will be the empty list for a message you had sent to only yourself.
	UserIds []int64 `json:"user_ids,omitempty"`
	// Present only if `type` is `RecipientTypeStream`.  The Id of the channel where the message was sent.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// Present only if `type` is `RecipientTypeStream`.  Name of the topic where the message was sent.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
	// **Deprecated** internal implementation detail. Clients should ignore this field as it will be removed in the future.
	// Deprecated
	UnmutedStreamMsg *bool `json:"unmuted_stream_msg,omitempty"`
}
