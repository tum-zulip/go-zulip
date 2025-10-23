package invites

import "github.com/tum-zulip/go-zulip/zulip"

// CreateInviteLinkResponse struct for CreateInviteLinkResponse
type CreateInviteLinkResponse struct {
	zulip.Response

	// The URL of the [reusable invitation link] that was created by this request.
	//
	// [reusable invitation link]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
	InviteLink string `json:"invite_link,omitempty"`
}

// GetInvitesResponse struct for GetInvitesResponse
type GetInvitesResponse struct {
	zulip.Response

	// An array of objects, each representing a single unexpired [invitation].
	//
	// [invitation]: https://zulip.com/help/invite-new-users
	Invites []zulip.Invite `json:"invites,omitempty"`
}
