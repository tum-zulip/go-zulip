package api

// CreateMessageReminderResponse struct for CreateMessageReminderResponse
type CreateMessageReminderResponse struct {
	Response

	// Unique Id of the scheduled message reminder.
	ReminderId int64 `json:"reminder_id,omitempty"`
}

// GetRemindersResponse struct for GetRemindersResponse
type GetRemindersResponse struct {
	Response

	// Returns all of the current user's undelivered reminders, ordered by `scheduled_delivery_timestamp` (ascending).
	Reminders []Reminder `json:"reminders,omitempty"`
}
