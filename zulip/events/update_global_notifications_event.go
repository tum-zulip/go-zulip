package events

// UpdateGlobalNotificationsEvent Event sent to a user's clients when that user's [notification settings] have changed with an additional rule that it is only sent to clients that did not include `user_settings_object` in their `client_capabilities` when registering the event queue.
//
// **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.
//
// [notification settings]: https://zulip.com/api/update-settings
type UpdateGlobalNotificationsEvent struct {
	event

	// The Zulip API email of the user.
	User string `json:"user,omitempty"`
	// Name of the changed notification setting.
	NotificationName string      `json:"notification_name,omitempty"`
	Setting          UpdateValue `json:"setting,omitempty"`
}
