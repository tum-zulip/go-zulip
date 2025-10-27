package reminders

import "github.com/tum-zulip/go-zulip/zulip"

// CreateMessageReminderResponse struct for CreateMessageReminderResponse
type CreateMessageReminderResponse struct {
	zulip.Response

	// Unique ID of the scheduled message reminder.
	ReminderID int64 `json:"reminder_id,omitempty"`
}

// GetRemindersResponse struct for GetRemindersResponse
type GetRemindersResponse struct {
	zulip.Response

	// Returns all of the current user's undelivered reminders, ordered by `scheduled_delivery_timestamp` (ascending).
	Reminders []zulip.Reminder `json:"reminders,omitempty"`
}
