package zulip

// RealmExport Object containing details about a [data export].
//
// [data export]: https://zulip.com/help/export-your-organization
type RealmExport struct {
	// The ID of the data export.
	ID int64 `json:"id,omitempty"`
	// The ID of the user who created the data export.
	ActingUserID int64 `json:"acting_user_id,omitempty"`
	// The UNIX timestamp of when the data export was started.
	ExportTime float32 `json:"export_time,omitempty"`
	// The UNIX timestamp of when the data export was deleted.  Will be `null` if the data export has not been deleted.
	DeletedTimestamp *float32 `json:"deleted_timestamp,omitempty"`
	// The UNIX timestamp of when the data export failed.  Will be `null` if the data export succeeded, or if it's still being generated.
	FailedTimestamp *float32 `json:"failed_timestamp,omitempty"`
	// The URL to download the generated data export.  Will be `null` if the data export failed, or if it's still being generated.
	ExportUrl *string `json:"export_url,omitempty"`
	// Whether the data export is pending, which indicates it is still being generated, or if it succeeded, failed or was deleted before being generated.  Depending on the size of the organization, it can take anywhere from seconds to an hour to generate the data export.
	Pending bool `json:"pending,omitempty"`
	// Whether the data export is a public or a standard data export.
	//   - ExportTypePublicData
	//   - ExportTypeStandardData
	//
	// **Changes**: New in Zulip 10.0 (feature level 304). Previously, the export type was not included in these objects because only public data exports could be created or listed via the API or UI.
	ExportType ExportType `json:"export_type,omitempty"`
}

// ExportConsent struct for ExportConsent
type ExportConsent struct {
	// The user ID.
	UserID int64 `json:"user_id,omitempty"`
	// Whether the user has consented for their private data export.
	Consented bool `json:"consented,omitempty"`
}
