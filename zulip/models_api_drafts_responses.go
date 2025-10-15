package zulip

// CreateDraftsResponse struct for CreateDraftsResponse
type CreateDraftsResponse struct {
	Response
	// An array of the Ids for the drafts that were just created in the same order as they were submitted.
	Ids []int64 `json:"ids,omitempty"`
}

// CreateSavedSnippetResponse struct for CreateSavedSnippetResponse
type CreateSavedSnippetResponse struct {
	Response

	// The unique Id of the saved snippet created.
	SavedSnippetId int64 `json:"saved_snippet_id,omitempty"`
}

// GetSavedSnippetsResponse struct for GetSavedSnippetsResponse
type GetSavedSnippetsResponse struct {
	Response

	// An array of dictionaries containing data on all of the current user's saved snippets.
	SavedSnippets []SavedSnippet `json:"saved_snippets,omitempty"`
}

// GetDraftsResponse struct for GetDraftsResponse
type GetDraftsResponse struct {
	Response

	// The number of drafts the user currently has. Also the number of drafts returned under \"drafts\".
	Count int64 `json:"count,omitempty"`
	// Returns all of the current user's drafts, in order of last edit time (with the most recently edited draft appearing first).
	Drafts []Draft `json:"drafts,omitempty"`
}

// CreateInviteLinkResponse struct for CreateInviteLinkResponse
type CreateInviteLinkResponse struct {
	Response

	// The URL of the [reusable invitation link](zulip.com/help/invite-new-users#create-a-reusable-invitation-link that was created by this request.
	InviteLink string `json:"invite_link,omitempty"`
}

// GetInvitesResponse struct for GetInvitesResponse
type GetInvitesResponse struct {
	Response

	// An array of objects, each representing a single unexpired [invitation](zulip.com/help/invite-new-users.
	Invites []Invite `json:"invites,omitempty"`
}
