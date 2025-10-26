package events

// UserSettingsUpdateEvent Event sent to a user's clients when that user's settings have changed.
//
// **Changes**: New in Zulip 5.0 (feature level 89), replacing the previous `update_display_settings` and `update_global_notifications` event types, which are still present for backwards compatibility reasons.
type UserSettingsUpdateEvent struct {
	event
	// Name of the changed setting.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
	// Present only if the setting to be changed is `default_language`. Contains the name of the new default language in English.
	LanguageName *string `json:"language_name,omitempty"`
}
