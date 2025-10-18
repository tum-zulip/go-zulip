package zulip

import (
	"encoding/json"
	"time"
)

// ScheduledMessage struct for ScheduledMessage
type ScheduledMessage struct {
	// The unique Id of the scheduled message, which can be used to modify or delete the scheduled message.  This is different from the unique Id that the message will have after it is sent.
	ScheduledMessageId int64 `json:"scheduled_message_id"`
	// The type of the scheduled message. Either `\"stream\"` or `\"private\"`.
	Type RecipientType `json:"type"`
	To   Recipient     `json:"to"`
	// Only present if `type` is `\"stream\"`.  The topic for the channel message.
	Topic *string `json:"topic,omitempty"`
	// The content/body of the scheduled message, in [Zulip-flavored Markdown](https://zulip.com/help/format-your-message-using-markdown) format.  See [Markdown message formatting](https://zulip.com/api/message-formatting) for details on Zulip's HTML format.
	Content string `json:"content"`
	// The content/body of the scheduled message rendered in HTML.
	RenderedContent string `json:"rendered_content"`
	// The UNIX timestamp for when the message will be sent by the server, in UTC seconds.
	ScheduledDeliveryTimestamp time.Time `json:"scheduled_delivery_timestamp"`
	// Whether the server has tried to send the scheduled message and it failed to successfully send.  Clients that support unscheduling and editing scheduled messages should display scheduled messages with `\"failed\": true` with an indicator that the server failed to send the message at the scheduled time, so that the user is aware of the failure and can get the content of the scheduled message.  **Changes**: New in Zulip 7.0 (feature level 181).
	Failed bool `json:"failed"`
}

type scheduledMessageJSON struct {
	ScheduledMessageId         int64         `json:"scheduled_message_id"`
	Type                       RecipientType `json:"type"`
	To                         Recipient     `json:"to"`
	Topic                      *string       `json:"topic,omitempty"`
	Content                    string        `json:"content"`
	RenderedContent            string        `json:"rendered_content"`
	ScheduledDeliveryTimestamp int64         `json:"scheduled_delivery_timestamp"`
	Failed                     bool          `json:"failed"`
}

func (o ScheduledMessage) MarshalJSON() ([]byte, error) {
	aux := scheduledMessageJSON{
		ScheduledMessageId:         o.ScheduledMessageId,
		Type:                       o.Type,
		To:                         o.To,
		Topic:                      o.Topic,
		Content:                    o.Content,
		RenderedContent:            o.RenderedContent,
		ScheduledDeliveryTimestamp: o.ScheduledDeliveryTimestamp.Unix(),
		Failed:                     o.Failed,
	}

	return json.Marshal(aux)
}

func (o *ScheduledMessage) UnmarshalJSON(data []byte) error {
	var aux scheduledMessageJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ScheduledMessageId = aux.ScheduledMessageId
	o.Type = aux.Type
	o.To = aux.To
	o.Topic = aux.Topic
	o.Content = aux.Content
	o.RenderedContent = aux.RenderedContent
	o.ScheduledDeliveryTimestamp = time.Unix(aux.ScheduledDeliveryTimestamp, 0).UTC()
	o.Failed = aux.Failed

	return nil
}
