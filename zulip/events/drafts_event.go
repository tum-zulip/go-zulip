package events

import "github.com/tum-zulip/go-zulip/zulip"

// DraftsAddEvent Event containing details of newly created drafts.
type DraftsAddEvent struct {
	event

	// An array containing objects for the newly created drafts.
	Drafts []zulip.Draft `json:"drafts,omitempty"`
}

// DraftsUpdateEvent Event containing details for an edited draft.
type DraftsUpdateEvent struct {
	event

	Draft zulip.Draft `json:"draft,omitempty"`
}

// DraftsRemoveEvent Event containing the Id of a deleted draft.
type DraftsRemoveEvent struct {
	event

	// The Id of the draft that was just deleted.
	DraftId int64 `json:"draft_id,omitempty"`
}
