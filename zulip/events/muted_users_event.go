package events

import "github.com/tum-zulip/go-zulip/zulip"

// MutedUsersEvent Event sent to a user's clients when that user's set of configured [muted users] have changed.
//
// **Changes**: New in Zulip 4.0 (feature level 48).
//
// [muted users]: https://zulip.com/api/mute-user
type MutedUsersEvent struct {
	event
	// A list of dictionaries where each dictionary describes a muted user.
	MutedUsers []zulip.MutedUser `json:"muted_users,omitempty"`
}
