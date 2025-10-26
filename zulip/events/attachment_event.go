package events

import "github.com/tum-zulip/go-zulip/zulip"

// AttachmentAddEvent Event sent to a user's clients when the user uploads a new file in a Zulip message. Useful to implement live update in UI showing all files the current user has uploaded.
type AttachmentAddEvent struct {
	event

	Attachment zulip.Attachment `json:"attachment,omitempty"`
	// The total size of all files uploaded by in the organization, in bytes.
	UploadSpaceUsed int64 `json:"upload_space_used,omitempty"`
}

// AttachmentUpdateEvent Event sent to a user's clients when details of a file that user uploaded are changed. Most updates will be changes in the list of messages that reference the uploaded file.
type AttachmentUpdateEvent struct{ AttachmentAddEvent }

// AttachmentRemoveEvent Event sent to a user's clients when the user deletes a file they had uploaded. Useful primarily for UI showing all the files the current user has uploaded.
type AttachmentRemoveEvent struct {
	event
	Attachment AttachmentId `json:"attachment,omitempty"`
	// The total size of all files uploaded by in the organization, in bytes.
	UploadSpaceUsed int64 `json:"upload_space_used,omitempty"`
}

type AttachmentId struct {
	// The Id of the deleted attachment.
	Id int64 `json:"id,omitempty"`
}
