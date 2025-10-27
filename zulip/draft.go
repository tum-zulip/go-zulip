package zulip

import (
	"encoding/json"
	"reflect"
	"time"
)

// Draft A dictionary for representing a message draft.
type Draft struct {
	// The unique ID of the draft. It will only used whenever the drafts are fetched. This field should not be specified when the draft is being created or edited.
	ID *int64 `json:"id,omitempty"`
	// The type of the draft. Either unaddressed (empty string), `RecipientTypeStream`, or `RecipientTypePrivate` (for one-on-one and group direct messages).
	Type RecipientType `json:"type"`
	// An array of the tentative target audience IDs. For channel messages, this should contain exactly 1 ID, the ID of the target channel. For direct messages, this should be an array of target user IDs. For unaddressed drafts, this is ignored, and clients should send an empty array.
	To Recipient `json:"to"`
	// For channel message drafts, the tentative topic name. For direct or unaddressed messages, this will be ignored and should ideally be the empty string. Should not contain null bytes.
	Topic string `json:"topic"`
	// The body of the draft. Should not contain null bytes.
	Content string `json:"content"`
	// A Unix timestamp (seconds only) representing when the draft was last edited. When creating a draft, this key need not be present and it will be filled in automatically by the server.
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type draftJSON struct {
	ID        *int64        `json:"id,omitempty"`
	Type      RecipientType `json:"type"`
	To        []int64       `json:"to"`
	Topic     string        `json:"topic"`
	Content   string        `json:"content"`
	Timestamp int64         `json:"timestamp,omitempty"`
}

func (o Draft) MarshalJSON() ([]byte, error) {
	aux := draftJSON{
		ID:        o.ID,
		Type:      o.Type,
		To:        o.To.asArray(),
		Topic:     o.Topic,
		Content:   o.Content,
		Timestamp: o.Timestamp.Unix(),
	}

	return json.Marshal(aux)
}

func (o *Draft) UnmarshalJSON(data []byte) error {
	var aux draftJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ID = aux.ID
	o.Type = aux.Type

	if aux.To == nil {
		aux.To = []int64{}
	}
	switch o.Type {
	case RecipientTypeEmpty:
		o.To = Recipient{}
	case RecipientTypeChannel, RecipientTypeStream:
		if len(aux.To) != 1 {
			return &json.UnsupportedValueError{
				Value: reflect.ValueOf(aux.To),
				Str:   "expected exactly one channel ID for channel recipient",
			}
		}
		o.To = ChannelAsRecipient(aux.To[0])
	case RecipientTypeDirect, RecipientTypePrivate:
		o.To = UsersAsRecipient(aux.To)
	default:
		return &json.UnsupportedValueError{Value: reflect.ValueOf(o.Type), Str: "unknown recipient type"}
	}

	o.Topic = aux.Topic
	o.Content = aux.Content
	o.Timestamp = time.Unix(aux.Timestamp, 0).UTC()

	return nil
}
