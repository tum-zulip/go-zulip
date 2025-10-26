package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmPlaygroundsEvent Event sent to all users in a Zulip organization when the set of configured [code playgrounds] for the organization has changed.
//
// **Changes**: New in Zulip 4.0 (feature level 49).
//
// [code playgrounds]: https://zulip.com/help/code-blocks#code-playgrounds
type RealmPlaygroundsEvent struct {
	event
	// An array of dictionaries where each dictionary contains data about a single playground entry.
	RealmPlaygrounds []zulip.RealmPlayground `json:"realm_playgrounds,omitempty"`
}
