package events

// SubmessageEvent Event sent when a submessage is added to a message.  Submessages are an **experimental** API used for widgets such as the `/poll` widget in Zulip.
type SubmessageEvent struct {
	event

	// The type of the message.
	MsgType string `json:"msg_type,omitempty"`
	// The new content of the submessage.
	Content string `json:"content,omitempty"`
	// The ID of the message to which the submessage has been added.
	MessageID int64 `json:"message_id,omitempty"`
	// The ID of the user who sent the message.
	SenderID int64 `json:"sender_id,omitempty"`
	// The ID of the submessage.
	SubmessageID int64 `json:"submessage_id,omitempty"`
}
