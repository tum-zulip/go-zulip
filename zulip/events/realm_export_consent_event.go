package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmExportConsentEvent Event sent to administrators when the [data export consent] status for a user changes, whether due to a user changing their consent preferences or a user being created or reactivated (since user creation/activation events do not contain these data).
//
// **Changes**: New in Zulip 10.0 (feature level 312). Previously, there was not event available to administrators with these data.
//
// [data export consent]: https://zulip.com/help/export-your-organization#configure-whether-administrators-can-export-your-private-data
type RealmExportConsentEvent struct {
	event

	zulip.ExportConsent
}
