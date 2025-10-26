package events

import "github.com/tum-zulip/go-zulip/zulip"

// SavedSnippetsEvent Event containing details of a newly created saved snippet.
//
// **Changes**: New in Zulip 10.0 (feature level 297).
type SavedSnippetsEvent struct {
	event

	SavedSnippet zulip.SavedSnippet `json:"saved_snippet,omitempty"`
}

// SavedSnippetsRemoveEvent Event containing the Id of a deleted saved snippet.
//
// **Changes**: New in Zulip 10.0 (feature level 297).
type SavedSnippetsRemoveEvent struct {
	event

	// The Id of the saved snippet that was just deleted.
	//
	// **Changes**: New in Zulip 10.0 (feature level 297).
	SavedSnippetId int64 `json:"saved_snippet_id,omitempty"`
}
