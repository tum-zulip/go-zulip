package events

import "github.com/tum-zulip/go-zulip/zulip/internal/utils"

// EventOp represents the type of operation that occurred in a model event.
//   - EventOpAdd
//   - EventOpRemove
//   - EventOpUpdate
//   - EventOpOnboardingSteps
//   - EventOpDeactivated
//   - EventOpCreate
//   - EventOpDelete
//   - EventOpChange
//   - EventOpReorder
//   - EventOpStart
//   - EventOpStop
//   - EventOpUpdateDict
//   - EventOpPeerAdd
//   - EventOpPeerRemove
//   - EventOpAddMembers
//   - EventOpAddSubgroups
//   - EventOpRemoveMembers
//   - EventOpRemoveSubgroups
type EventOp string

const (
	EventOpAdd             EventOp = "add"
	EventOpRemove          EventOp = "remove"
	EventOpUpdate          EventOp = "update"
	EventOpOnboardingSteps EventOp = "onboarding_steps"
	EventOpDeactivated     EventOp = "deactivated"
	EventOpCreate          EventOp = "create"
	EventOpDelete          EventOp = "delete"
	EventOpChange          EventOp = "change"
	EventOpReorder         EventOp = "reorder"
	EventOpStart           EventOp = "start"
	EventOpStop            EventOp = "stop"
	EventOpUpdateDict      EventOp = "update_dict"
	EventOpPeerAdd         EventOp = "peer_add"
	EventOpPeerRemove      EventOp = "peer_remove"
	EventOpAddMembers      EventOp = "add_members"
	EventOpAddSubgroups    EventOp = "add_subgroups"
	EventOpRemoveMembers   EventOp = "remove_members"
	EventOpRemoveSubgroups EventOp = "remove_subgroups"
)

func (e *EventOp) UnmarshalJSON(data []byte) error {
	allowedEventOpValues := []EventOp{
		EventOpAdd,
		EventOpRemove,
		EventOpUpdate,
		EventOpOnboardingSteps,
		EventOpDeactivated,
		EventOpCreate,
		EventOpDelete,
		EventOpChange,
		EventOpReorder,
		EventOpStart,
		EventOpStop,
		EventOpUpdateDict,
		EventOpPeerAdd,
		EventOpPeerRemove,
		EventOpAddMembers,
		EventOpAddSubgroups,
		EventOpRemoveMembers,
		EventOpRemoveSubgroups,
	}

	return utils.UnmarshalStringEnum(data, e, allowedEventOpValues)
}

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

	EventTypeUnknown EventType = "unknown" // used when utils.Unmarshaling an unknown event type
	EventTypeInvalid EventType = "invalid" // used when an invalid value is supplied
)

func (e *EventType) UnmarshalJSON(data []byte) error {
	allowedEventTypeValues := []EventType{
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

	return utils.UnmarshalStringEnum(data, e, allowedEventTypeValues)
}

// The [organization type] for the realm.
//   - OrgTypeUnspecified = Unspecified
//   - OrgTypeBusiness = Business
//   - OrgTypeOpenSource = Open-source project
//   - OrgTypeEducationNonProfit = Education (non-profit)
//   - OrgTypeEducationForProfit = Education (for-profit)
//   - OrgTypeResearch = Research
//   - OrgTypeEventOrConference = Event or conference
//   - OrgTypeNonProfitRegistered = Non-profit (registered)
//   - OrgTypeGovernment = Government
//   - OrgTypePoliticalGroup = Political group
//   - OrgTypeCommunity = Community
//   - OrgTypePersonal = Personal
//   - OrgTypeOther = Other
//
// **Changes**: New in Zulip 6.0 (feature level 128).
//
// [organization type]: https://zulip.com/help/organization-type
type OrgType int

const (
	OrgTypeUnspecified         OrgType = 0
	OrgTypeBusiness            OrgType = 10
	OrgTypeOpenSource          OrgType = 20
	OrgTypeEducationNonProfit  OrgType = 30
	OrgTypeEducationForProfit  OrgType = 35
	OrgTypeResearch            OrgType = 40
	OrgTypeEventOrConference   OrgType = 50
	OrgTypeNonProfitRegistered OrgType = 60
	OrgTypeGovernment          OrgType = 70
	OrgTypePoliticalGroup      OrgType = 80
	OrgTypeCommunity           OrgType = 90
	OrgTypePersonal            OrgType = 100
	OrgTypeOther               OrgType = 1000
)

func (e *OrgType) UnmarshalJSON(data []byte) error {
	allowedOrgTypeValues := []OrgType{
		OrgTypeUnspecified,
		OrgTypeBusiness,
		OrgTypeOpenSource,
		OrgTypeEducationNonProfit,
		OrgTypeEducationForProfit,
		OrgTypeResearch,
		OrgTypeEventOrConference,
		OrgTypeNonProfitRegistered,
		OrgTypeGovernment,
		OrgTypePoliticalGroup,
		OrgTypeCommunity,
		OrgTypePersonal,
		OrgTypeOther,
	}

	return utils.UnmarshalIntEnum(data, e, allowedOrgTypeValues)
}

// The plan type of the organization.
//   - PlanTypeSelfHosted = Self-hosted organization (SELF_HOSTED)
//   - PlanTypeLimited = Zulip Cloud free plan (LIMITED)
//   - PlanTypeStandard = Zulip Cloud Standard plan (STANDARD)
//   - PlanTypeStandardFree = Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)
type PlanType int

const (
	PlanTypeSelfHosted   PlanType = 1
	PlanTypeLimited      PlanType = 2
	PlanTypeStandard     PlanType = 3
	PlanTypeStandardFree PlanType = 4
)

func (e *PlanType) UnmarshalJSON(data []byte) error {
	allowedPlanTypeValues := []PlanType{PlanTypeSelfHosted, PlanTypeLimited, PlanTypeStandard, PlanTypeStandardFree}

	return utils.UnmarshalIntEnum(data, e, allowedPlanTypeValues)
}

// Whether to [include organization name in subject of message notification emails].
//   - NameInEmailNotificationsPolicyAutomatic
//   - NameInEmailNotificationsPolicyAlways
//   - NameInEmailNotificationsPolicyNever
//
// **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
//
// [include organization name in subject of message notification emails]: https://zulip.com/help/email-notifications#include-organization-name-in-subject-line
type NameInEmailNotificationsPolicy int

const (
	NameInEmailNotificationsPolicyAutomatic NameInEmailNotificationsPolicy = 1
	NameInEmailNotificationsPolicyAlways    NameInEmailNotificationsPolicy = 2
	NameInEmailNotificationsPolicyNever     NameInEmailNotificationsPolicy = 3
)

func (e *NameInEmailNotificationsPolicy) UnmarshalJSON(data []byte) error {
	allowedNameInEmailNotificationsPolicyValues := []NameInEmailNotificationsPolicy{
		NameInEmailNotificationsPolicyAutomatic,
		NameInEmailNotificationsPolicyAlways,
		NameInEmailNotificationsPolicyNever,
	}

	return utils.UnmarshalIntEnum(data, e, allowedNameInEmailNotificationsPolicyValues)
}

// Whether or not to mark messages as read when the user scrolls through their feed.
//   - MarkReadOnScrollPolicyAlways = Always
//   - MarkReadOnScrollPolicyOnlyInConversationViews = Only in conversation views
//   - MarkReadOnScrollPolicyNever = Never
//
// **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the "Always" setting when marking messages as read.
type MarkReadOnScrollPolicy int

const (
	MarkReadOnScrollPolicyAlways                  MarkReadOnScrollPolicy = 1
	MarkReadOnScrollPolicyOnlyInConversationViews MarkReadOnScrollPolicy = 2
	MarkReadOnScrollPolicyNever                   MarkReadOnScrollPolicy = 3
)

func (e *MarkReadOnScrollPolicy) UnmarshalJSON(data []byte) error {
	allowedMarkReadOnScrollPolicyValues := []MarkReadOnScrollPolicy{
		MarkReadOnScrollPolicyAlways,
		MarkReadOnScrollPolicyOnlyInConversationViews,
		MarkReadOnScrollPolicyNever,
	}

	return utils.UnmarshalIntEnum(data, e, allowedMarkReadOnScrollPolicyValues)
}

// The configured [video call provider] for the organization.
//   - VideoChatProviderNone = None
//   - VideoChatProviderJitsiMeet = Jitsi Meet
//   - VideoChatProviderZoomUserOAuthIntegration = Zoom (User OAuth integration)
//   - VideoChatProviderBigBlueButton = BigBlueButton
//   - VideoChatProviderZoomServerToServerOAuth = Zoom (Server to Server OAuth integration)
//
// Note that only one of the [Zoom integrations] can be configured on a Zulip server.
// **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.
//
// [Zoom integrations]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom
// [video call provider]: https://zulip.com/help/configure-call-provider
type VideoChatProvider int

const (
	VideoChatProviderNone                     VideoChatProvider = 0
	VideoChatProviderJitsiMeet                VideoChatProvider = 1
	VideoChatProviderZoomUserOAuthIntegration VideoChatProvider = 3
	VideoChatProviderBigBlueButton            VideoChatProvider = 4
	VideoChatProviderZoomServerToServerOAuth  VideoChatProvider = 5
)

func (e *VideoChatProvider) UnmarshalJSON(data []byte) error {
	allowedVideoChatProviderValues := []VideoChatProvider{
		VideoChatProviderNone,
		VideoChatProviderJitsiMeet,
		VideoChatProviderZoomUserOAuthIntegration,
		VideoChatProviderBigBlueButton,
		VideoChatProviderZoomServerToServerOAuth,
	}

	return utils.UnmarshalIntEnum(data, e, allowedVideoChatProviderValues)
}
