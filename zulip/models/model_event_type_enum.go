package models

// EventType represents the type of event that occurred in a model event.
//   - EventTypeAlertWords
//   - EventTypeAttachment
//   - EventTypeChannelFolder
//   - EventTypeCustomProfileFields
//   - EventTypeDefaultChannelGroups
//   - EventTypeDefaultChannels
//   - EventTypeDeleteMessage
//   - EventTypeDrafts
//   - EventTypeHasZoomToken
//   - EventTypeHeartbeat
//   - EventTypeInvitesChanged
//   - EventTypeMessage
//   - EventTypeMutedTopics
//   - EventTypeMutedUsers
//   - EventTypeNavigationView
//   - EventTypeOnboardingSteps
//   - EventTypePresence
//   - EventTypePushDevice
//   - EventTypeReaction
//   - EventTypeRealm
//   - EventTypeRealmBot
//   - EventTypeRealmDomains
//   - EventTypeRealmEmoji
//   - EventTypeRealmExport
//   - EventTypeRealmExportConsent
//   - EventTypeRealmFilters
//   - EventTypeRealmLinkifiers
//   - EventTypeRealmPlaygrounds
//   - EventTypeRealmUser
//   - EventTypeRealmUserSettingsDefaults
//   - EventTypeReminders
//   - EventTypeRestart
//   - EventTypeSavedSnippets
//   - EventTypeScheduledMessages
//   - EventTypeChannel
//   - EventTypeSubmessage
//   - EventTypeSubscription
//   - EventTypeTyping
//   - EventTypeTypingEditMessage
//   - EventTypeUpdateDisplaySettings
//   - EventTypeUpdateGlobalNotifications
//   - EventTypeUpdateMessage
//   - EventTypeUpdateMessageFlags
//   - EventTypeUserGroup
//   - EventTypeUserSettings
//   - EventTypeUserStatus
//   - EventTypeUserTopic
//   - EventTypeWebReloadClient
type EventType string

const (
	EventTypeAlertWords                EventType = "alert_words"
	EventTypeAttachment                EventType = "attachment"
	EventTypeChannelFolder             EventType = "channel_folder"
	EventTypeCustomProfileFields       EventType = "custom_profile_fields"
	EventTypeDefaultChannelGroups      EventType = "default_stream_groups"
	EventTypeDefaultChannels           EventType = "default_streams"
	EventTypeDeleteMessage             EventType = "delete_message"
	EventTypeDrafts                    EventType = "drafts"
	EventTypeHasZoomToken              EventType = "has_zoom_token"
	EventTypeHeartbeat                 EventType = "heartbeat"
	EventTypeInvitesChanged            EventType = "invites_changed"
	EventTypeMessage                   EventType = "message"
	EventTypeMutedTopics               EventType = "muted_topics"
	EventTypeMutedUsers                EventType = "muted_users"
	EventTypeNavigationView            EventType = "navigation_view"
	EventTypeOnboardingSteps           EventType = "onboarding_steps"
	EventTypePresence                  EventType = "presence"
	EventTypePushDevice                EventType = "push_device"
	EventTypeReaction                  EventType = "reaction"
	EventTypeRealm                     EventType = "realm"
	EventTypeRealmBot                  EventType = "realm_bot"
	EventTypeRealmDomains              EventType = "realm_domains"
	EventTypeRealmEmoji                EventType = "realm_emoji"
	EventTypeRealmExport               EventType = "realm_export"
	EventTypeRealmExportConsent        EventType = "realm_export_consent"
	EventTypeRealmFilters              EventType = "realm_filters" // Deprecated
	EventTypeRealmLinkifiers           EventType = "realm_linkifiers"
	EventTypeRealmPlaygrounds          EventType = "realm_playgrounds"
	EventTypeRealmUser                 EventType = "realm_user"
	EventTypeRealmUserSettingsDefaults EventType = "realm_user_settings_defaults"
	EventTypeReminders                 EventType = "reminders"
	EventTypeRestart                   EventType = "restart"
	EventTypeSavedSnippets             EventType = "saved_snippets"
	EventTypeScheduledMessages         EventType = "scheduled_messages"
	EventTypeChannel                   EventType = "stream"
	EventTypeSubmessage                EventType = "submessage"
	EventTypeSubscription              EventType = "subscription"
	EventTypeTyping                    EventType = "typing"
	EventTypeTypingEditMessage         EventType = "typing_edit_message"
	EventTypeUpdateDisplaySettings     EventType = "update_display_settings"
	EventTypeUpdateGlobalNotifications EventType = "update_global_notifications"
	EventTypeUpdateMessage             EventType = "update_message"
	EventTypeUpdateMessageFlags        EventType = "update_message_flags"
	EventTypeUserGroup                 EventType = "user_group"
	EventTypeUserSettings              EventType = "user_settings"
	EventTypeUserStatus                EventType = "user_status"
	EventTypeUserTopic                 EventType = "user_topic"
	EventTypeWebReloadClient           EventType = "web_reload_client"
	// Special non-zulip event types used for error handling
	EventTypeUnknown EventType = "unknown" // used when unmarshaling an unknown event type
	EventTypeInvalid EventType = "invalid" // used when an invalid value is supplied
)

var allowedEventTypeEnumValues = []EventType{
	EventTypeAlertWords,
	EventTypeAttachment,
	EventTypeChannelFolder,
	EventTypeCustomProfileFields,
	EventTypeDefaultChannelGroups,
	EventTypeDefaultChannels,
	EventTypeDeleteMessage,
	EventTypeDrafts,
	EventTypeHasZoomToken,
	EventTypeHeartbeat,
	EventTypeInvitesChanged,
	EventTypeMessage,
	EventTypeMutedTopics,
	EventTypeMutedUsers,
	EventTypeNavigationView,
	EventTypeOnboardingSteps,
	EventTypePresence,
	EventTypePushDevice,
	EventTypeReaction,
	EventTypeRealm,
	EventTypeRealmBot,
	EventTypeRealmDomains,
	EventTypeRealmEmoji,
	EventTypeRealmExport,
	EventTypeRealmExportConsent,
	EventTypeRealmFilters,
	EventTypeRealmLinkifiers,
	EventTypeRealmPlaygrounds,
	EventTypeRealmUser,
	EventTypeRealmUserSettingsDefaults,
	EventTypeReminders,
	EventTypeRestart,
	EventTypeSavedSnippets,
	EventTypeScheduledMessages,
	EventTypeChannel,
	EventTypeSubmessage,
	EventTypeSubscription,
	EventTypeTyping,
	EventTypeTypingEditMessage,
	EventTypeUpdateDisplaySettings,
	EventTypeUpdateGlobalNotifications,
	EventTypeUpdateMessage,
	EventTypeUpdateMessageFlags,
	EventTypeUserGroup,
	EventTypeUserSettings,
	EventTypeUserStatus,
	EventTypeUserTopic,
	EventTypeWebReloadClient,
	EventTypeUnknown,
	EventTypeInvalid,
}

func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    allowedEventTypeEnumValues,
			Value:   v,
			VarName: "EventType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventType) IsValid() bool {
	for _, existing := range allowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
