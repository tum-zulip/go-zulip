package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmDomainsAddEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsAddEvent struct {
	event

	RealmDomain *zulip.RealmDomain `json:"realm_domain,omitempty"`
}

// RealmDomainsChangeEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsChangeEvent struct {
	event

	RealmDomain zulip.RealmDomain `json:"realm_domain,omitempty"`
}

// RealmDomainsRemoveEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsRemoveEvent struct {
	event

	// The domain to be removed.
	Domain string `json:"domain,omitempty"`
}
