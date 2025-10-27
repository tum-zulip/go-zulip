package events

import "github.com/tum-zulip/go-zulip/zulip"

// UserStatusEvent Event sent to all users who can access the modified user when the status of a user changes.
//
// **Changes**: Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
type UserStatusEvent struct {
	event

	zulip.UserStatus

	// The ID of the user whose status changed.
	UserID int64 `json:"user_id,omitempty"`
}
