package zulip

type EventType string

const (
	EventTypeAlertWords          EventType = "alert_words"
	EventTypeAttachment          EventType = "attachment"
	EventTypeChannelFolder       EventType = "channel_folder"
	EventTypeCustomProfileFields EventType = "custom_profile_fields"
	EventTypeDefaultStreamGroups EventType = "default_stream_groups"
	EventTypeDefaultStreams      EventType = "default_streams"
	EventTypeDeleteMessage       EventType = "delete_message"
	EventTypeDrafts              EventType = "drafts"
	EventTypeHasZoomToken        EventType = "has_zoom_token"
	EventTypeHeartbeat           EventType = "heartbeat"
	EventTypeInvitesChanged      EventType = "invites_changed"
	EventTypeMessage             EventType = "message"
	EventTypeMutedTopics         EventType = "muted_topics"
	EventTypeMutedUsers          EventType = "muted_users"
	EventTypeNavigationView      EventType = "navigation_view"
	EventTypeOnboardingSteps     EventType = "onboarding_steps"
	EventTypePresence            EventType = "presence"
	EventTypePushDevice          EventType = "push_device"
	EventTypeReaction            EventType = "reaction"
	EventTypeRealm               EventType = "realm"
	EventTypeRealmBot            EventType = "realm_bot"
	EventTypeRealmDomains        EventType = "realm_domains"
	EventTypeRealmEmoji          EventType = "realm_emoji"
	EventTypeRealmExport         EventType = "realm_export"
	EventTypeRealmExportConsent  EventType = "realm_export_consent"
	// Deprecated
	EventTypeRealmFilters              EventType = "realm_filters"
	EventTypeRealmLinkifiers           EventType = "realm_linkifiers"
	EventTypeRealmPlaygrounds          EventType = "realm_playgrounds"
	EventTypeRealmUser                 EventType = "realm_user"
	EventTypeRealmUserSettingsDefaults EventType = "realm_user_settings_defaults"
	EventTypeReminders                 EventType = "reminders"
	EventTypeRestart                   EventType = "restart"
	EventTypeSavedSnippets             EventType = "saved_snippets"
	EventTypeScheduledMessages         EventType = "scheduled_messages"
	EventTypeStream                    EventType = "stream"
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
)

var AllowedEventTypeEnumValues = []EventType{
	EventTypeAlertWords,
	EventTypeAttachment,
	EventTypeChannelFolder,
	EventTypeCustomProfileFields,
	EventTypeDefaultStreamGroups,
	EventTypeDefaultStreams,
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
	EventTypeStream,
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
}

func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedEventTypeEnumValues,
			Value:   v,
			VarName: "EventType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventType) IsValid() bool {
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
