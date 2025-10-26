package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmExportEvent Event sent to the user who requested a [data export] when the status of the data export changes.
//
// [data export]: https://zulip.com/help/export-your-organization
type RealmExportEvent struct {
	event
	// An array of dictionaries where each dictionary contains details about a data export of the organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 304), `export_type` parameter was not present as only public data export was supported via API.
	Exports []zulip.RealmExport `json:"exports,omitempty"`
}
