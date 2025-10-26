package events

// SubmessageEvent Event sent when a submessage is added to a message.  Submessages are an **experimental** API used for widgets such as the `/poll` widget in Zulip.
type SubmessageEvent struct {
	event
	// The type of the message.
	MsgType string `json:"msg_type,omitempty"`
	// The new content of the submessage.
	Content string `json:"content,omitempty"`
	// The Id of the message to which the submessage has been added.
	MessageId int64 `json:"message_id,omitempty"`
	// The Id of the user who sent the message.
	SenderId int64 `json:"sender_id,omitempty"`
	// The Id of the submessage.
	SubmessageId int64 `json:"submessage_id,omitempty"`
}
