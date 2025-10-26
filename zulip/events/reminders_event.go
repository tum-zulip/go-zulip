package events

import "github.com/tum-zulip/go-zulip/zulip"

// RemindersAddEvent Event sent to a user's clients when a reminder is scheduled.
//
// **Changes**: New in Zulip 11.0 (feature level 399).
type RemindersAddEvent struct {
	event

	// An array of objects containing details of the newly created reminders.
	Reminders []zulip.Reminder `json:"reminders,omitempty"`
}

// RemindersRemoveEvent Event sent to a user's clients when a reminder is deleted.
//
// **Changes**: New in Zulip 11.0 (feature level 399).
type RemindersRemoveEvent struct {
	event

	// The Id of the reminder that was deleted.
	ReminderId int64 `json:"reminder_id,omitempty"`
}
