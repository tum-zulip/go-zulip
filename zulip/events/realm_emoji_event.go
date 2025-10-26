package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmEmojiUpdateEvent Event sent to all users in a Zulip organization when a [custom emoji] has been updated, typically when a new emoji has been added or an old one has been deactivated. The event contains all custom emoji configured for the organization, not just the updated custom emoji.
//
// [custom emoji]: https://zulip.com/help/custom-emoji
type RealmEmojiUpdateEvent struct {
	event

	// An object in which each key describes a realm emoji.
	RealmEmoji map[string]zulip.RealmEmoji `json:"realm_emoji,omitempty"`
}
