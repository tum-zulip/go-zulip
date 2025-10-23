package drafts

import "github.com/tum-zulip/go-zulip/zulip"

// CreateDraftsResponse struct for CreateDraftsResponse
type CreateDraftsResponse struct {
	zulip.Response
	// An array of the Ids for the drafts that were just created in the same order as they were submitted.
	Ids []int64 `json:"ids,omitempty"`
}

// CreateSavedSnippetResponse struct for CreateSavedSnippetResponse
type CreateSavedSnippetResponse struct {
	zulip.Response

	// The unique Id of the saved snippet created.
	SavedSnippetId int64 `json:"saved_snippet_id,omitempty"`
}

// GetSavedSnippetsResponse struct for GetSavedSnippetsResponse
type GetSavedSnippetsResponse struct {
	zulip.Response

	// An array of dictionaries containing data on all of the current user's saved snippets.
	SavedSnippets []zulip.SavedSnippet `json:"saved_snippets,omitempty"`
}

// GetDraftsResponse struct for GetDraftsResponse
type GetDraftsResponse struct {
	zulip.Response

	// The number of drafts the user currently has. Also the number of drafts returned under "drafts".
	Count int64 `json:"count,omitempty"`
	// Returns all of the current user's drafts, in order of last edit time (with the most recently edited draft appearing first).
	Drafts []zulip.Draft `json:"drafts,omitempty"`
}
