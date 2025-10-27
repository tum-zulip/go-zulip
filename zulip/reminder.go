package zulip

import (
	"encoding/json"
	"time"
)

// Reminder Object containing details of the scheduled message.
type Reminder struct {
	// The unique ID of the reminder, which can be used to delete the reminder.  This is different from the unique ID that the message would have after being sent.
	ReminderID int64 `json:"reminder_id"`
	// The type of the reminder. Always set to `RecipientTypePrivate`.
	Type RecipientType `json:"type"`
	// Contains the ID of the user who scheduled the reminder, and to which the reminder will be sent.
	To []int64 `json:"to"`
	// The content/body of the reminder, in [Zulip-flavored Markdown] format.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	Content string `json:"content"`
	// The content/body of the reminder rendered in HTML.
	RenderedContent string `json:"rendered_content"`
	// The UNIX timestamp for when the message will be sent by the server, in UTC seconds.
	ScheduledDeliveryTimestamp time.Time `json:"scheduled_delivery_timestamp"`
	// Whether the server has tried to send the reminder and it failed to successfully send.  Clients that support unscheduling reminders should display scheduled messages with `"failed": true` with an indicator that the server failed to send the message at the scheduled time, so that the user is aware of the failure and can get the content of the scheduled message.
	Failed bool `json:"failed"`
	// The ID of the message that the reminder is created for.
	ReminderTargetMessageID int64 `json:"reminder_target_message_id"`
}

type reminderJSON struct {
	ReminderID                 int64         `json:"reminder_id"`
	Type                       RecipientType `json:"type"`
	To                         []int64       `json:"to"`
	Content                    string        `json:"content"`
	RenderedContent            string        `json:"rendered_content"`
	ScheduledDeliveryTimestamp int64         `json:"scheduled_delivery_timestamp"`
	Failed                     bool          `json:"failed"`
	ReminderTargetMessageID    int64         `json:"reminder_target_message_id"`
}

func (o Reminder) MarshalJSON() ([]byte, error) {
	aux := reminderJSON{
		ReminderID:                 o.ReminderID,
		Type:                       o.Type,
		To:                         o.To,
		Content:                    o.Content,
		RenderedContent:            o.RenderedContent,
		ScheduledDeliveryTimestamp: o.ScheduledDeliveryTimestamp.Unix(),
		Failed:                     o.Failed,
		ReminderTargetMessageID:    o.ReminderTargetMessageID,
	}

	return json.Marshal(aux)
}

func (o *Reminder) UnmarshalJSON(data []byte) error {
	var aux reminderJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ReminderID = aux.ReminderID
	o.Type = aux.Type
	o.To = aux.To
	o.Content = aux.Content
	o.RenderedContent = aux.RenderedContent
	o.ScheduledDeliveryTimestamp = time.Unix(aux.ScheduledDeliveryTimestamp, 0).UTC()
	o.Failed = aux.Failed
	o.ReminderTargetMessageID = aux.ReminderTargetMessageID

	return nil
}
