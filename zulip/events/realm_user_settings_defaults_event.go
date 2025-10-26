package events

// RealmUserSettingsDefaultsEvent Event sent to all users in a Zulip organization when the [default settings for new users] of the organization (realm) have changed.
// See [PATCH /realm/user_settings_defaults] for details on possible properties.
// **Changes**: New in Zulip 5.0 (feature level 95).
//
// [default settings for new users]: https://zulip.com/help/configure-default-new-user-settings
// [PATCH /realm/user_settings_defaults]: https://zulip.com/api/update-realm-user-settings-defaults
type RealmUserSettingsDefaultsEvent struct {
	event

	// The name of the property that was changed.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}
