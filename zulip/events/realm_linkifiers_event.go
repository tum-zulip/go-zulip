package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmLinkifiersEvent Event sent to all users in a Zulip organization when the set of configured [linkifiers] for the organization has changed.  Processing this event is important for doing Markdown local echo correctly.  Clients will not receive this event unless the event queue is registered with the client capability `{"linkifier_url_template": true}`. See [`POST /register`], the `linkifier_url_template` client capability was not required. The requirement was added because linkifiers were updated to contain a URL template instead of a URL format string, which was not a backwards-compatible change.  New in Zulip 4.0 (feature level 54), replacing the deprecated `realm_filters` event type.
//
// **Changes**: Before Zulip 7.0 (feature level 176
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
// [`POST /register`]: https://zulip.com/api/register-queue#parameter-client_capabilities for how client capabilities can be specified.
type RealmLinkifiersEvent struct {
	event

	// An ordered array of dictionaries where each dictionary contains details about a single linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [`PATCH /realm/linkifiers`].
	//
	// [`PATCH /realm/linkifiers`]: https://zulip.com/api/reorder-linkifiers
	RealmLinkifiers []zulip.RealmLinkifiers `json:"realm_linkifiers,omitempty"`
}
