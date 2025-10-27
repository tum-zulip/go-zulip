package events

import "github.com/tum-zulip/go-zulip/zulip"

// ScheduledMessagesAddEvent Event sent to a user's clients when scheduled messages are created.
//
// **Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesAddEvent struct {
	event

	// An array of objects containing details of the newly created scheduled messages.
	ScheduledMessages []zulip.ScheduledMessage `json:"scheduled_messages,omitempty"`
}

// ScheduledMessagesUpdateEvent Event sent to a user's clients when a scheduled message is edited.
//
// **Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesUpdateEvent struct {
	event

	ScheduledMessage zulip.ScheduledMessage `json:"scheduled_message,omitempty"`
}

// ScheduledMessagesRemoveEvent Event sent to a user's clients when a scheduled message is deleted.
//
// **Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesRemoveEvent struct {
	event

	// The ID of the scheduled message that was deleted.
	ScheduledMessageID int64 `json:"scheduled_message_id,omitempty"`
}
