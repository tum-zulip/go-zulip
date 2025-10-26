package events

import "github.com/tum-zulip/go-zulip/zulip"

// TypingEditMessageStartEvent Event sent when a user starts editing a message. Event sent when a user starts typing in a textarea to edit the content of a message. See the [edit message typing notifications endpoint].  Clients requesting `typing_edit_message` event type that have `receives_typing_notifications` enabled will receive this event if they would have been notified if the message's content edit were to be saved (E.g., because they were a direct message recipient or are a subscribe to the channel).
//
// **Changes**: New in Zulip 10.0 (feature level 351). Previously, typing notifications were not available when editing messages.
//
// [edit message typing notifications endpoint]: https://zulip.com/api/set-typing-status-for-message-edit
type TypingEditMessageEvent struct {
	event

	// The Id of the user who is typing the edit of the message.  Clients should be careful to display this user as the person who is typing, not that of the sender of the message, in case a collaborative editing feature be might be added in the future.
	SenderId int64 `json:"sender_id,omitempty"`
	// Indicates the message id of the message that is being edited.
	MessageId int64          `json:"message_id,omitempty"`
	Recipient *RecipientData `json:"recipient,omitempty"`
}

// Recipient Object containing details about recipient of message edit typing notification.
type RecipientData struct {
	// Type of message being composed. Must be `RecipientTypeChannel` or `RecipientTypeDirect`.
	Type zulip.RecipientType `json:"type,omitempty"`
	// Only present if `type` is `RecipientTypeChannel`.  The unique Id of the channel to which message is being edited.
	ChannelId *int64 `json:"channel_id,omitempty"`
	// Only present if `type` is `RecipientTypeChannel`.  Topic within the channel where the message is being edited.
	Topic *string `json:"topic,omitempty"`
	// Present only if `type` is `RecipientTypeDirect`.  The user Ids of every recipient of this direct message.
	UserIds []int64 `json:"user_ids,omitempty"`
}
