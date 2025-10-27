package scheduledmessages

import "github.com/tum-zulip/go-zulip/zulip"

// CreateScheduledMessageResponse struct for CreateScheduledMessageResponse.
type CreateScheduledMessageResponse struct {
	zulip.Response

	// The unique ID of the scheduled message.  This is different from the unique ID that the message will have after it is sent.
	ScheduledMessageID int64 `json:"scheduled_message_id,omitempty"`
}

// GetScheduledMessagesResponse struct for GetScheduledMessagesResponse.
type GetScheduledMessagesResponse struct {
	zulip.Response

	// Returns all of the current user's undelivered scheduled messages, ordered by `scheduled_delivery_timestamp` (ascending).
	ScheduledMessages []zulip.ScheduledMessage `json:"scheduled_messages,omitempty"`
}
