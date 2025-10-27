package zulip

import (
	"encoding/json"
	"time"
)

// Attachment Dictionary containing details of a file uploaded by a user.
type Attachment struct {
	// The unique ID for the attachment.
	ID int64 `json:"id,omitempty"`
	// Name of the uploaded file.
	Name string `json:"name,omitempty"`
	// A representation of the path of the file within the repository of user-uploaded files. If the `path_id` of a file is `{realm_id}/ab/cdef/temp_file.py`, its URL will be: `{server_url}/user_uploads/{realm_id}/ab/cdef/temp_file.py`.
	PathID string `json:"path_id,omitempty"`
	// Size of the file in bytes.
	Size int64 `json:"size,omitempty"`
	// Time when the attachment was uploaded as a UNIX timestamp multiplied by 1000 (matching the format of getTime() in JavaScript).
	//
	// **Changes**: Changed in Zulip 3.0 (feature level 22). This field was previously a floating point number.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Contains basic details on any Zulip messages that have been sent referencing this [uploaded file]. This includes messages sent by any user in the Zulip organization who sent a message containing a link to the uploaded file.
	//
	// [uploaded file]: https://zulip.com/api/upload-file
	Messages []AttachmentMessages `json:"messages,omitempty"`
}

// AttachmentMessages struct for AttachmentMessages
type AttachmentMessages struct {
	// Time when the message was sent as a UNIX timestamp multiplied by 1000 (matching the format of getTime() in JavaScript).
	//
	// **Changes**: Changed in Zulip 3.0 (feature level 22). This field was previously strangely called `name` and was a floating point number.
	DateSent time.Time `json:"date_sent,omitempty"`
	// The unique message ID. Messages should always be displayed sorted by ID.
	ID int64 `json:"id,omitempty"`
}

type attachmentJSON struct {
	ID         int64                `json:"id,omitempty"`
	Name       string               `json:"name,omitempty"`
	PathID     string               `json:"path_id,omitempty"`
	Size       int64                `json:"size,omitempty"`
	CreateTime int64                `json:"create_time,omitempty"`
	Messages   []AttachmentMessages `json:"messages,omitempty"`
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	aux := attachmentJSON{
		ID:         o.ID,
		Name:       o.Name,
		PathID:     o.PathID,
		Size:       o.Size,
		Messages:   o.Messages,
		CreateTime: o.CreateTime.UnixMilli(),
	}

	return json.Marshal(aux)
}

func (o *Attachment) UnmarshalJSON(data []byte) error {
	var aux attachmentJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ID = aux.ID
	o.Name = aux.Name
	o.PathID = aux.PathID
	o.Size = aux.Size
	o.Messages = aux.Messages
	o.CreateTime = time.UnixMilli(aux.CreateTime)

	return nil
}

type attachmentMessagesJSON struct {
	DateSent int64 `json:"date_sent,omitempty"`
	ID       int64 `json:"id,omitempty"`
}

func (o AttachmentMessages) MarshalJSON() ([]byte, error) {
	aux := attachmentMessagesJSON{
		ID:       o.ID,
		DateSent: o.DateSent.UnixMilli(),
	}

	return json.Marshal(aux)
}

func (o *AttachmentMessages) UnmarshalJSON(data []byte) error {
	var aux attachmentMessagesJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ID = aux.ID
	o.DateSent = time.UnixMilli(aux.DateSent)

	return nil
}
