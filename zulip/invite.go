package zulip

import (
	"encoding/json"
	"time"
)

// Invite A dictionary containing details about an [invitation].
//
// [invitation]: https://zulip.com/help/invite-new-users
type Invite struct {
	// The ID of the invitation.  Note that email invitations and reusable invitation links are stored in different database tables on the server, so each ID is guaranteed to be unique in combination with the boolean value of `is_multiuse`, e.g. there can only be one invitation with `id: 1` and `is_multiuse: true`.
	ID int64 `json:"id,omitempty"`
	// The [user ID] of the user who created the invitation.
	//
	// **Changes**: New in Zulip 3.0 (feature level 22), replacing the `ref` field which contained the Zulip display email address of the user who created the invitation.
	//
	// [user ID]: https://zulip.com/api/get-user
	InvitedByUserID int64 `json:"invited_by_user_id,omitempty"`
	// The UNIX timestamp for when the invitation was created, in UTC seconds.
	Invited time.Time `json:"invited,omitempty"`
	// The UNIX timestamp for when the invitation will expire, in UTC seconds. If `null`, the invitation never expires.
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	// The [organization-level role] of the user that is created when the invitation is accepted. Possible values are:
	//   - RoleOwner
	//   - RoleAdmin
	//   - RoleModerator
	//   - RoleMember
	//   - RoleGuest
	//
	// [organization-level role]: https://zulip.com/api/roles-and-permissions
	InvitedAs Role `json:"invited_as,omitempty"`
	// The email address the invitation was sent to. This will not be present when `is_multiuse` is `true` (i.e. the invitation is a reusable invitation link).
	Email string `json:"email,omitempty"`
	// A boolean indicating whether the referrer has opted to receive a direct message from [notification bot] when a user account is created using this invitation.
	//
	// **Changes**: New in Zulip 9.0 (feature level 267). Previously, referrers always received such direct messages.
	//
	// [notification bot]: https://zulip.com/help/configure-automated-notices
	NotifyReferrerOnJoin bool `json:"notify_referrer_on_join,omitempty"`
	// The URL of the reusable invitation link. This will not be present when `is_multiuse` is `false` (i.e. the invitation is an email invitation).
	LinkUrl string `json:"link_url,omitempty"`
	// A boolean specifying whether the [invitation] is a reusable invitation link or an email invitation.
	//
	// [invitation]: https://zulip.com/help/invite-new-users
	IsMultiuse bool `json:"is_multiuse,omitempty"`
}

type inviteJSON struct {
	ID                   int64  `json:"id,omitempty"`
	InvitedByUserID      int64  `json:"invited_by_user_id,omitempty"`
	Invited              int64  `json:"invited,omitempty"`
	ExpiryDate           *int64 `json:"expiry_date,omitempty"`
	InvitedAs            Role   `json:"invited_as,omitempty"`
	Email                string `json:"email,omitempty"`
	NotifyReferrerOnJoin bool   `json:"notify_referrer_on_join,omitempty"`
	LinkUrl              string `json:"link_url,omitempty"`
	IsMultiuse           bool   `json:"is_multiuse,omitempty"`
}

func (o Invite) MarshalJSON() ([]byte, error) {
	aux := inviteJSON{
		ID:                   o.ID,
		InvitedByUserID:      o.InvitedByUserID,
		Invited:              o.Invited.Unix(),
		InvitedAs:            o.InvitedAs,
		Email:                o.Email,
		NotifyReferrerOnJoin: o.NotifyReferrerOnJoin,
		LinkUrl:              o.LinkUrl,
		IsMultiuse:           o.IsMultiuse,
	}

	if o.ExpiryDate != nil {
		ts := o.ExpiryDate.Unix()
		aux.ExpiryDate = &ts
	}

	return json.Marshal(aux)
}

func (o *Invite) UnmarshalJSON(data []byte) error {
	var aux inviteJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.ID = aux.ID
	o.InvitedByUserID = aux.InvitedByUserID
	o.Invited = time.Unix(aux.Invited, 0).UTC()
	if aux.ExpiryDate != nil {
		t := time.Unix(*aux.ExpiryDate, 0).UTC()
		o.ExpiryDate = &t
	} else {
		o.ExpiryDate = nil
	}
	o.InvitedAs = aux.InvitedAs
	o.Email = aux.Email
	o.NotifyReferrerOnJoin = aux.NotifyReferrerOnJoin
	o.LinkUrl = aux.LinkUrl
	o.IsMultiuse = aux.IsMultiuse

	return nil
}
