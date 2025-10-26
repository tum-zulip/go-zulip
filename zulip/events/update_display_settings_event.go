package events

// UpdateDisplaySettingsEvent Event sent to clients that have requested the `update_display_settings` event type and did not include `user_settings_object` in their `client_capabilities` when registering the event queue.
//
// **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.
type UpdateDisplaySettingsEvent struct {
	event

	// The Zulip API email of the user.
	User string `json:"user,omitempty"`
	// Name of the changed display setting.
	SettingName string      `json:"setting_name,omitempty"`
	Setting     UpdateValue `json:"setting,omitempty"`
	// Present only if the setting to be changed is `default_language`. Contains the name of the new default language in English.
	LanguageName *string `json:"language_name,omitempty"`
}
