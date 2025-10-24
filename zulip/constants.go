package zulip

import "github.com/tum-zulip/go-zulip/zulip/internal/utils"

// The avatar data source type for the current user.
//   - AvatarSourceGravatar = Gravatar
//   - AvatarSourceUploaded = Uploaded by user
type AvatarSource string

const (
	AvatarSourceGravatar AvatarSource = "G"
	AvatarSourceUploaded AvatarSource = "U"
)

func (e *AvatarSource) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedAvatarSourceValues)
}

var allowedAvatarSourceValues = []AvatarSource{AvatarSourceGravatar, AvatarSourceUploaded}

// Unread count badge (appears in desktop sidebar and browser tab)
//   - BadgeCountAllUnreadMessages = All unread messages
//   - BadgeCountDMsMentionsAndFollowedTopics = DMs, mentions, and followed topics
//   - BadgeCountDMsAndMentions = DMs and mentions
//   - BadgeCountNone = None
//
// **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
type BadgeCount int

const (
	BadgeCountAllUnreadMessages            BadgeCount = 1
	BadgeCountDMsMentionsAndFollowedTopics BadgeCount = 2
	BadgeCountDMsAndMentions               BadgeCount = 3
	BadgeCountNone                         BadgeCount = 4
)

func (e *BadgeCount) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedBadgeCountValues)
}

var allowedBadgeCountValues = []BadgeCount{BadgeCountAllUnreadMessages, BadgeCountDMsMentionsAndFollowedTopics, BadgeCountDMsAndMentions, BadgeCountNone}

// BotType types Zulip supports
//   - BotTypeGeneric = Generic bot
//   - BotTypeIncomingWebhook = Incoming webhook
//   - BotTypeOutgoingWebhook = Outgoing webhook
//   - BotTypeEmbedded = Embedded bot
//
// These can be used to identify the type of a bot user.
type BotType int

const (
	BotTypeGeneric         BotType = 1
	BotTypeIncomingWebhook BotType = 2
	BotTypeOutgoingWebhook BotType = 3
	BotTypeEmbedded        BotType = 4
)

func (e *BotType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedBotTypeValues)
}

var allowedBotTypeValues = []BotType{BotTypeGeneric, BotTypeIncomingWebhook, BotTypeOutgoingWebhook, BotTypeEmbedded}

// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.
//   - ChannelDefaultViewTopTopicInChannel = Top topic in the channel
//   - ChannelDefaultViewChannelFeed = Channel feed
//   - ChannelDefaultViewListOfTopics = List of topics
//   - ChannelDefaultViewTopUnreadTopicInChannel = Top unread topic in channel
//
// **Changes**: The "Top unread topic in channel" is new in Zulip 11.0 (feature level 401).  The "List of topics" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the "Channel feed" behavior.
type ChannelDefaultView int

const (
	ChannelDefaultViewTopTopicInChannel       ChannelDefaultView = 1
	ChannelDefaultViewChannelFeed             ChannelDefaultView = 2
	ChannelDefaultViewListOfTopics            ChannelDefaultView = 3
	ChannelDefaultViewTopUnreadTopicInChannel ChannelDefaultView = 4
)

func (e *ChannelDefaultView) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedChannelDefaultViewValues)
}

var allowedChannelDefaultViewValues = []ChannelDefaultView{ChannelDefaultViewTopTopicInChannel, ChannelDefaultViewChannelFeed, ChannelDefaultViewListOfTopics, ChannelDefaultViewTopUnreadTopicInChannel}

// Controls which [color theme] to use.
//   - ColorSchemeAutomatic = Automatic
//   - ColorSchemeDark = Dark theme
//   - ColorSchemeLight = Light theme
//
// Automatic detection is implementing using the standard `prefers-color-scheme` media query.
//
// [color theme]: https://zulip.com/help/dark-theme
type ColorScheme int

const (
	ColorSchemeAutomatic ColorScheme = 1
	ColorSchemeDark      ColorScheme = 2
	ColorSchemeLight     ColorScheme = 3
)

func (e *ColorScheme) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedColorSchemeValues)
}

var allowedColorSchemeValues = []ColorScheme{ColorSchemeAutomatic, ColorSchemeDark, ColorSchemeLight}

// See the [Custom profile fields] article for details on what each type means.
//   - CustomFieldTypeShortText = Short text
//   - CustomFieldTypeLongText = Long text
//   - CustomFieldTypeListOfOptions = List of options
//   - CustomFieldTypeDatePicker = Date picker
//   - CustomFieldTypeLink = Link
//   - CustomFieldTypePersonPicker = Person picker
//   - CustomFieldTypeExternalAccount = External account
//   - CustomFieldTypePronouns = Pronouns
type CustomFieldType int

const (
	CustomFieldTypeShortText       CustomFieldType = 1 // Short text
	CustomFieldTypeLongText        CustomFieldType = 2 // Long text
	CustomFieldTypeListOfOptions   CustomFieldType = 3 // List of options
	CustomFieldTypeDatePicker      CustomFieldType = 4 // Date picker
	CustomFieldTypeLink            CustomFieldType = 5 // Link
	CustomFieldTypePersonPicker    CustomFieldType = 6 // Person picker
	CustomFieldTypeExternalAccount CustomFieldType = 7 // External account
	CustomFieldTypePronouns        CustomFieldType = 8 // Pronouns
)

func (e *CustomFieldType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedCustomFieldTypeValues)
}

var allowedCustomFieldTypeValues = []CustomFieldType{CustomFieldTypeShortText, CustomFieldTypeLongText, CustomFieldTypeListOfOptions, CustomFieldTypeDatePicker, CustomFieldTypeLink, CustomFieldTypePersonPicker, CustomFieldTypeExternalAccount, CustomFieldTypePronouns}

// Whether to [hide inactive channels] in the left sidebar.
//   - DemoteInactiveChannelsAutomatic = Automatic
//   - DemoteInactiveChannelsAlways = Always
//   - DemoteInactiveChannelsNever = Never
//
// [hide inactive channels]: https://zulip.com/help/manage-inactive-channels
type DemoteInactiveChannels int

const (
	DemoteInactiveChannelsAutomatic DemoteInactiveChannels = 1
	DemoteInactiveChannelsAlways    DemoteInactiveChannels = 2
	DemoteInactiveChannelsNever     DemoteInactiveChannels = 3
)

func (e *DemoteInactiveChannels) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedDemoteInactiveChannelsValues)
}

var allowedDemoteInactiveChannelsValues = []DemoteInactiveChannels{DemoteInactiveChannelsAutomatic, DemoteInactiveChannelsAlways, DemoteInactiveChannelsNever}

// The [policy] for [which other users] in this organization can see the user's real email address.
//   - EmailVisibilityEveryone = Everyone
//   - EmailVisibilityMembersOnly = Members only
//   - EmailVisibilityAdministratorsOnly = Administrators only
//   - EmailVisibilityNobody = Nobody
//   - EmailVisibilityModeratorsOnly = Moderators only
//
// **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
//
// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
// [which other users]: https://zulip.com/help/configure-email-visibility
type EmailVisibility int

func (e *EmailVisibility) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedEmailVisibilityValues)
}

const (
	EmailVisibilityEveryone           EmailVisibility = 1
	EmailVisibilityMembersOnly        EmailVisibility = 2
	EmailVisibilityAdministratorsOnly EmailVisibility = 3
	EmailVisibilityNobody             EmailVisibility = 4
	EmailVisibilityModeratorsOnly     EmailVisibility = 5
)

var allowedEmailVisibilityValues = []EmailVisibility{EmailVisibilityEveryone, EmailVisibilityMembersOnly, EmailVisibilityAdministratorsOnly, EmailVisibilityNobody, EmailVisibilityModeratorsOnly}

// The user's configured [emoji set], used to display emoji to the user everywhere they appear in the UI.
//   - EmojisetGoogle = Google modern
//   - EmojisetGoogleBlob = Google classic
//   - EmojisetTwitter = Twitter
//   - EmojisetText = Plain text
//
// [emoji set]: https://zulip.com/help/emoji-and-emoticons#use-emoticons
type Emojiset string

const (
	EmojisetGoogle     Emojiset = "google"
	EmojisetGoogleBlob Emojiset = "google-blob"
	EmojisetTwitter    Emojiset = "twitter"
	EmojisetText       Emojiset = "text"
)

func (e *Emojiset) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedEmojisetValues)
}

var allowedEmojisetValues = []Emojiset{EmojisetGoogle, EmojisetGoogleBlob, EmojisetTwitter, EmojisetText}

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
	return utils.UnmarshalStringEnum(data, e, allowedEventOpValues)
}

var allowedEventOpValues = []EventOp{EventOpAdd, EventOpRemove, EventOpUpdate, EventOpOnboardingSteps, EventOpDeactivated, EventOpCreate, EventOpDelete, EventOpChange, EventOpReorder, EventOpStart, EventOpStop, EventOpUpdateDict, EventOpPeerAdd, EventOpPeerRemove, EventOpAddMembers, EventOpAddSubgroups, EventOpRemoveMembers, EventOpRemoveSubgroups}

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
	return utils.UnmarshalStringEnum(data, e, allowedEventTypeValues)
}

var allowedEventTypeValues = []EventType{EventTypeAlertWords, EventTypeAttachment, EventTypeChannelFolder, EventTypeCustomProfileFields, EventTypeDefaultChannelGroups, EventTypeDefaultChannels, EventTypeDeleteMessage, EventTypeDrafts, EventTypeHasZoomToken, EventTypeHeartbeat, EventTypeInvitesChanged, EventTypeMessage, EventTypeMutedTopics, EventTypeMutedUsers, EventTypeNavigationView, EventTypeOnboardingSteps, EventTypePresence, EventTypePushDevice, EventTypeReaction, EventTypeRealm, EventTypeRealmBot, EventTypeRealmDomains, EventTypeRealmEmoji, EventTypeRealmExport, EventTypeRealmExportConsent, EventTypeRealmFilters, EventTypeRealmLinkifiers, EventTypeRealmPlaygrounds, EventTypeRealmUser, EventTypeRealmUserSettingsDefaults, EventTypeReminders, EventTypeRestart, EventTypeSavedSnippets, EventTypeScheduledMessages, EventTypeChannel, EventTypeSubmessage, EventTypeSubscription, EventTypeTyping, EventTypeTypingEditMessage, EventTypeUpdateDisplaySettings, EventTypeUpdateGlobalNotifications, EventTypeUpdateMessage, EventTypeUpdateMessageFlags, EventTypeUserGroup, EventTypeUserSettings, EventTypeUserStatus, EventTypeUserTopic, EventTypeWebReloadClient, EventTypeUnknown, EventTypeInvalid}

// Whether the data export is a public or a standard data export.
//   - ExportTypePublicData = Public data export.
//   - ExportTypeStandardData = Standard data export.
//
// **Changes**: New in Zulip 10.0 (feature level 304).
// Previously, the export type was not included in these objects because only public data exports could be created or listed via the API or UI.
type ExportType int

const (
	ExportTypePublicData   ExportType = 1
	ExportTypeStandardData ExportType = 2
)

func (e *ExportType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedExportTypeValues)
}

var allowedExportTypeValues = []ExportType{ExportTypePublicData, ExportTypeStandardData}

// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.
//   - HomeViewRecentTopics = Recent conversations view
//   - HomeViewInbox = Inbox view
//   - HomeViewAllMessages = Combined feed view
//
// **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
//
// [home view]: https://zulip.com/help/configure-home-view
type HomeView string

const (
	HomeViewRecentTopics HomeView = "recent_topics"
	HomeViewInbox        HomeView = "inbox"
	HomeViewAllMessages  HomeView = "all_messages"
)

func (e *HomeView) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedHomeViewValues)
}

var allowedHomeViewValues = []HomeView{HomeViewRecentTopics, HomeViewInbox, HomeViewAllMessages}

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
	return utils.UnmarshalIntEnum(data, e, allowedMarkReadOnScrollPolicyValues)
}

var allowedMarkReadOnScrollPolicyValues = []MarkReadOnScrollPolicy{MarkReadOnScrollPolicyAlways, MarkReadOnScrollPolicyOnlyInConversationViews, MarkReadOnScrollPolicyNever}

type MessageRetentionDays string

const (
	MessageRetentionDaysRealmDefault MessageRetentionDays = "realm_default"
	MessageRetentionDaysUnlimited    MessageRetentionDays = "unlimited"
	MessageRetentionDaysForever      MessageRetentionDays = "forever" // Deprecated: use MessageRetentionDaysUnlimited
)

func (e *MessageRetentionDays) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedMessageRetentionDaysValues)
}

var allowedMessageRetentionDaysValues = []MessageRetentionDays{MessageRetentionDaysRealmDefault, MessageRetentionDaysUnlimited, MessageRetentionDaysForever}

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
	return utils.UnmarshalIntEnum(data, e, allowedNameInEmailNotificationsPolicyValues)
}

var allowedNameInEmailNotificationsPolicyValues = []NameInEmailNotificationsPolicy{NameInEmailNotificationsPolicyAutomatic, NameInEmailNotificationsPolicyAlways, NameInEmailNotificationsPolicyNever}

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
	return utils.UnmarshalIntEnum(data, e, allowedOrgTypeValues)
}

var allowedOrgTypeValues = []OrgType{OrgTypeUnspecified, OrgTypeBusiness, OrgTypeOpenSource, OrgTypeEducationNonProfit, OrgTypeEducationForProfit, OrgTypeResearch, OrgTypeEventOrConference, OrgTypeNonProfitRegistered, OrgTypeGovernment, OrgTypePoliticalGroup, OrgTypeCommunity, OrgTypePersonal, OrgTypeOther}

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
	return utils.UnmarshalIntEnum(data, e, allowedPlanTypeValues)
}

var allowedPlanTypeValues = []PlanType{PlanTypeSelfHosted, PlanTypeLimited, PlanTypeStandard, PlanTypeStandardFree}

// PresenceStatus types Zulip supports
//   - PresenceStatusActive = active
//   - PresenceStatusIdle = idle
//
// These can be used to represent a user's presence status.
type PresenceStatus string

const (
	PresenceStatusActive PresenceStatus = "active"
	PresenceStatusIdle   PresenceStatus = "idle"
)

func (e *PresenceStatus) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedPresenceStatusValues)
}

var allowedPresenceStatusValues = []PresenceStatus{PresenceStatusActive, PresenceStatusIdle}

// A string indicating the type of emoji.
//
// Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
type ReactionType string

const (
	// In this namespace, `emoji_code` will be the Id of the uploaded [custom emoji].
	//
	// [custom emoji]: https://zulip.com/help/custom-emoji
	ReactionTypeRealmEmoji      ReactionType = "realm_emoji"
	ReactionTypeUnicodeEmoji    ReactionType = "unicode_emoji"     // In this namespace, `emoji_code` will be a  dash-separated hex encoding of the sequence of Unicode codepoints that define this emoji in the Unicode specification.
	ReactionTypeZulipExtraEmoji ReactionType = "zulip_extra_emoji" // These are special emoji included with Zulip. In this namespace, `emoji_code` will be the name of the emoji (e.g. "zulip").
	ReactionTypeEmpty           ReactionType = ""                  // For users who set a status without selecting an emoji.
)

func (e *ReactionType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedReactionTypeValues)
}

var allowedReactionTypeValues = []ReactionType{ReactionTypeUnicodeEmoji, ReactionTypeRealmEmoji, ReactionTypeZulipExtraEmoji, ReactionTypeEmpty}

// RecipientType - Enum for RecipientType
//   - RecipientTypeEmpty = An unaddressed draft.
//   - RecipientTypeDirect = A direct message draft (one-on-one or group).
//   - RecipientTypePrivate = Legacy value, maps to RecipientTypeDirect.
//   - RecipientTypeChannel = A channel message draft.
//   - RecipientTypeStream = Legacy value, maps to RecipientTypeChannel.
//
// When utils.Unmarshaling, `RecipientTypeStream` will always be converted to `RecipientTypeChannel`, and `RecipientTypePrivate` to `RecipientTypeDirect`, but when marshaling, the values will be as-is, because some API endpoints still expect the legacy values.
type RecipientType string

const (
	RecipientTypeEmpty   RecipientType = ""        //  An unaddressed draft.
	RecipientTypeDirect  RecipientType = "direct"  //  A direct message draft (one-on-one or group).
	RecipientTypePrivate RecipientType = "private" //  Legacy value, maps to RecipientTypeDirect.
	RecipientTypeChannel RecipientType = "channel" //  A channel message draft.
	RecipientTypeStream  RecipientType = "stream"  //  Legacy value, maps to RecipientTypeChannel.
)

func (v RecipientType) IsDirectMessage() bool {
	return v == RecipientTypeDirect || v == RecipientTypePrivate
}

func (v RecipientType) IsChannelMessage() bool {
	return v == RecipientTypeChannel || v == RecipientTypeStream
}

func (v RecipientType) ToLegacy() RecipientType {
	if v == RecipientTypeChannel {
		return RecipientTypeStream
	}
	return v
}

func (e *RecipientType) UnmarshalJSON(data []byte) error {
	err := utils.UnmarshalStringEnum(data, e, allowedRecipientTypeValues)
	if err != nil {
		return err
	}
	if *e == RecipientTypePrivate {
		*e = RecipientTypeDirect
	}
	if *e == RecipientTypeStream {
		*e = RecipientTypeChannel
	}
	return nil
}

var allowedRecipientTypeValues = []RecipientType{RecipientTypeEmpty, RecipientTypeDirect, RecipientTypePrivate, RecipientTypeChannel, RecipientTypeStream}

// Controls whether the resolved-topic notices are marked as read.
//   - ResolvedTopicNoticeAutoReadPolicyAlways = Always mark resolved-topic notices as read.
//   - ResolvedTopicNoticeAutoReadPolicyExceptFollowed = Mark resolved-topic notices as read in topics not followed by the user.
//   - ResolvedTopicNoticeAutoReadPolicyNever = Never mark resolved-topic notices as read.
//
// **Changes**: New in Zulip 11.0 (feature level 385).
type ResolvedTopicNoticeAutoReadPolicy string

const (
	ResolvedTopicNoticeAutoReadPolicyAlways         ResolvedTopicNoticeAutoReadPolicy = "always"
	ResolvedTopicNoticeAutoReadPolicyExceptFollowed ResolvedTopicNoticeAutoReadPolicy = "except_followed"
	ResolvedTopicNoticeAutoReadPolicyNever          ResolvedTopicNoticeAutoReadPolicy = "never"
)

func (e *ResolvedTopicNoticeAutoReadPolicy) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedResolvedTopicNoticeAutoReadPolicyValues)
}

var allowedResolvedTopicNoticeAutoReadPolicyValues = []ResolvedTopicNoticeAutoReadPolicy{ResolvedTopicNoticeAutoReadPolicyAlways, ResolvedTopicNoticeAutoReadPolicyExceptFollowed, ResolvedTopicNoticeAutoReadPolicyNever}

// [Organization-level role] of the user. Possible values are:
//   - RoleOwner = Organization owner
//   - RoleAdmin = Organization administrator
//   - RoleModerator = Organization moderator
//   - RoleMember = Member
//   - RoleGuest = Guest
//
// **Changes**: New in Zulip 4.0 (feature level 59).
//
// [Organization-level role]: https://zulip.com/api/roles-and-permissions
type Role int

// [Organization-level role] of the user. Possible values are:
//   - RoleOwner = Organization owner
//   - RoleAdmin = Organization administrator
//   - RoleModerator = Organization moderator
//   - RoleMember = Member
//   - RoleGuest = Guest
//
// **Changes**: New in Zulip 4.0 (feature level 59).
//
// [Organization-level role]: https://zulip.com/api/roles-and-permissions
const (
	RoleOwner     Role = 100
	RoleAdmin     Role = 200
	RoleModerator Role = 300
	RoleMember    Role = 400
	RoleGuest     Role = 600
)

func (e *Role) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedRoleValues)
}

var allowedRoleValues = []Role{RoleMember, RoleGuest, RoleModerator, RoleAdmin, RoleOwner}

// One of the channel properties described below:
//   - SubscriptionPropertyColor
//   - SubscriptionPropertyIsMuted
//   - SubscriptionPropertyPinToTop
//   - SubscriptionPropertyDesktopNotifications
//   - SubscriptionPropertyAudibleNotifications
//   - SubscriptionPropertyPushNotifications
//   - SubscriptionPropertyEmailNotifications
//   - SubscriptionPropertyWildcardMentionsNotify
//   - SubscriptionPropertyInHomeView
//
// [muted]: https://zulip.com/help/mute-a-channel
// [subscription update events]: https://zulip.com/api/get-events#subscription-update
type SubscriptionProperty string

const (
	// Whether the channel is [muted].
	//
	// **Changes**: As of Zulip 6.0 (feature level 139), updating either "is_muted" or "in_home_view" generates two [subscription update events], one for each property, that are sent to clients. Prior to this feature level, updating either property only generated a subscription update event for "in_home_view". Prior to Zulip 2.1.0, this feature was represented by the more confusingly named "in_home_view" (with the opposite value: `in_home_view=!is_muted`); for backwards-compatibility, modern Zulip still accepts that property.
	//
	// [muted]: https://zulip.com/help/mute-a-channel
	// [subscription update events]: https://zulip.com/api/get-events#subscription-update
	SubscriptionPropertyIsMuted                SubscriptionProperty = "is_muted"
	SubscriptionPropertyColor                  SubscriptionProperty = "color"                    // The hex value of the user's display color for the channel.
	SubscriptionPropertyPinToTop               SubscriptionProperty = "pin_to_top"               // Whether to pin the channel at the top of the channel list.
	SubscriptionPropertyDesktopNotifications   SubscriptionProperty = "desktop_notifications"    // Whether to show desktop notifications for all messages sent to the channel.
	SubscriptionPropertyAudibleNotifications   SubscriptionProperty = "audible_notifications"    // Whether to play a sound notification for all messages sent to the channel.
	SubscriptionPropertyPushNotifications      SubscriptionProperty = "push_notifications"       // Whether to trigger a mobile push notification for all messages sent to the channel.
	SubscriptionPropertyEmailNotifications     SubscriptionProperty = "email_notifications"      // Whether to trigger an email notification for all messages sent to the channel.
	SubscriptionPropertyWildcardMentionsNotify SubscriptionProperty = "wildcard_mentions_notify" // Whether wildcard mentions trigger notifications as though they were personal mentions in this channel.
	SubscriptionPropertyInHomeView             SubscriptionProperty = "in_home_view"             // Legacy name for `is_muted`, provided for backwards-compatibility with older Zulip server versions.
)

func (e *SubscriptionProperty) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedSubscriptionPropertyValues)
}

var allowedSubscriptionPropertyValues = []SubscriptionProperty{SubscriptionPropertyColor, SubscriptionPropertyIsMuted, SubscriptionPropertyPinToTop, SubscriptionPropertyDesktopNotifications, SubscriptionPropertyAudibleNotifications, SubscriptionPropertyPushNotifications, SubscriptionPropertyEmailNotifications, SubscriptionPropertyWildcardMentionsNotify, SubscriptionPropertyInHomeView}

// Topic interaction types Zulip supports
//   - TopicInteractionyTopicsTheUserParticipatesIn = Topics the user participates in
//   - TopicInteractionyTopicsTheUserSendsAMessageTo = Topics the user sends a message to
//   - TopicInteractionyTopicsTheUserStarts = Topics the user starts
//   - TopicInteractionyNever = Never
//
// These can be used to set which topics to unmute or follow automatically.
type TopicInteraction int

const (
	TopicInteractionyTopicsTheUserParticipatesIn  TopicInteraction = 1
	TopicInteractionyTopicsTheUserSendsAMessageTo TopicInteraction = 2
	TopicInteractionyTopicsTheUserStarts          TopicInteraction = 3
	TopicInteractionyNever                        TopicInteraction = 4
)

func (e *TopicInteraction) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedTopicInteractionValues)
}

var allowedTopicInteractionValues = []TopicInteraction{TopicInteractionyTopicsTheUserParticipatesIn, TopicInteractionyTopicsTheUserSendsAMessageTo, TopicInteractionyTopicsTheUserStarts, TopicInteractionyNever}

// Whether [named topics] and the empty topic (i.e., ["general chat" topic] are enabled in this channel.
//   - TopicsPolicyInherit
//   - TopicsPolicyAllowEmptyTopic
//   - TopicsPolicyDisableEmptyTopic
//   - TopicsPolicyEmptyTopicOnly
//
// **Changes**: In Zulip 11.0 (feature level 404), the `"empty_topic_only"` option was added.  New in Zulip 11.0 (feature level 392).
//
// [named topics]: https://zulip.com/help/introduction-to-topics
// ["general chat" topic]: https://zulip.com/help/general-chat-topic
type TopicsPolicy string

const (
	// Messages can be sent to named topics in this channel, and the [organization-level `realm_topics_policy`] is used for whether messages can be sent to the empty topic in this channel.
	//
	// [organization-level `realm_topics_policy`]: https://zulip.com/help/require-topics#set-the-default-general-chat-topic-configuration
	TopicsPolicyInherit TopicsPolicy = "inherit"
	// Messages can be sent to both named topics and the empty topic in this channel.
	TopicsPolicyAllowEmptyTopic TopicsPolicy = "allow_empty_topic"
	// Messages can be sent to named topics in this channel, but the empty topic is disabled.
	TopicsPolicyDisableEmptyTopic TopicsPolicy = "disable_empty_topic"
	// Messages can be sent to the empty topic in this channel, but named topics are disabled.
	//
	// See ["general chat" channels]. The `"empty_topic_only"` policy can only be set if all existing messages in the channel are already in the empty topic.  When creating a new channel, if the `topics_policy` is not specified, the `"inherit"` option will be set.
	//
	// ["general chat" channels]: https://zulip.com/help/general-chat-channels
	TopicsPolicyEmptyTopicOnly TopicsPolicy = "empty_topic_only"
)

func (e *TopicsPolicy) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedTopicsPolicyValues)
}

var allowedTopicsPolicyValues = []TopicsPolicy{TopicsPolicyAllowEmptyTopic, TopicsPolicyDisableEmptyTopic, TopicsPolicyEmptyTopicOnly, TopicsPolicyInherit}

// TypingStatusOp represents whether the user has started or stopped typing.
//   - TypingStatusOpStart = start
//   - TypingStatusOpStop = stop
type TypingStatusOp string

const (
	TypingStatusOpStart TypingStatusOp = "start"
	TypingStatusOpStop  TypingStatusOp = "stop"
)

func (e *TypingStatusOp) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedTypingStatusOpValues)
}

var allowedTypingStatusOpValues = []TypingStatusOp{TypingStatusOpStart, TypingStatusOpStop}

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.
//   - UnreadsCountDisplayAllChannels = All channels
//   - UnreadsCountDisplayUnmutedChannelsAndTopics = Unmuted channels and topics
//   - UnreadsCountDisplayNoChannels = No channels
//
// **Changes**: New in Zulip 8.0 (feature level 210).
type UnreadsCountDisplay int

const (
	UnreadsCountDisplayAllChannels              UnreadsCountDisplay = 1
	UnreadsCountDisplayUnmutedChannelsAndTopics UnreadsCountDisplay = 2
	UnreadsCountDisplayNoChannels               UnreadsCountDisplay = 3
)

func (e *UnreadsCountDisplay) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedUnreadsCountDisplayValues)
}

var allowedUnreadsCountDisplayValues = []UnreadsCountDisplay{UnreadsCountDisplayAllChannels, UnreadsCountDisplayUnmutedChannelsAndTopics, UnreadsCountDisplayNoChannels}

// The style selected by the user for the right sidebar user list.
//   - UserListStyleCompact = Compact
//   - UserListStyleWithStatus = With status
//   - UserListStyleWithAvatarAndStatus = With avatar and status
//
// **Changes**: New in Zulip 6.0 (feature level 141).
type UserListStyle int

const (
	UserListStyleCompact             UserListStyle = 1
	UserListStyleWithStatus          UserListStyle = 2
	UserListStyleWithAvatarAndStatus UserListStyle = 3
)

func (e *UserListStyle) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedUserListStyleValues)
}

var allowedUserListStyleValues = []UserListStyle{UserListStyleCompact, UserListStyleWithStatus, UserListStyleWithAvatarAndStatus}

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
	return utils.UnmarshalIntEnum(data, e, allowedVideoChatProviderValues)
}

var allowedVideoChatProviderValues = []VideoChatProvider{VideoChatProviderNone, VideoChatProviderJitsiMeet, VideoChatProviderZoomUserOAuthIntegration, VideoChatProviderBigBlueButton, VideoChatProviderZoomServerToServerOAuth}

// Controls visibility policy
//   - VisibilityPolicyNone = None. Removes the visibility policy previously set for the topic.
//   - VisibilityPolicyMuted = Muted. [Mutes the topic] in a channel.
//   - VisibilityPolicyUnmuted = Unmuted. [Unmutes the topic] in a muted channel.
//   - VisibilityPolicyFollowed = Followed. [Follows the topic].
//
// [Mutes the topic]: https://zulip.com/help/mute-a-topic
// [Follows the topic]: https://zulip.com/help/follow-a-topic
type VisibilityPolicy int

const (
	VisibilityPolicyNone     VisibilityPolicy = 0
	VisibilityPolicyMuted    VisibilityPolicy = 1
	VisibilityPolicyUnmuted  VisibilityPolicy = 2
	VisibilityPolicyFollowed VisibilityPolicy = 3
)

func (e *VisibilityPolicy) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalIntEnum(data, e, allowedVisibilityPolicyValues)
}

var allowedVisibilityPolicyValues = []VisibilityPolicy{VisibilityPolicyNone, VisibilityPolicyMuted, VisibilityPolicyUnmuted, VisibilityPolicyFollowed}

// Controls how animated images should be played in the message feed in the web/desktop application.
//   - WebAnimateImagePreviewsAlways = Always play the animated images in the message feed.
//   - WebAnimateImagePreviewsOnHover = Play the animated images on hover over them in the message feed.
//   - WebAnimateImagePreviewsNever = Never play animated images in the message feed.
//
// **Changes**: New in Zulip 9.0 (feature level 275).
type WebAnimateImagePreviews string

const (
	WebAnimateImagePreviewsAlways  WebAnimateImagePreviews = "always"
	WebAnimateImagePreviewsOnHover WebAnimateImagePreviews = "on_hover"
	WebAnimateImagePreviewsNever   WebAnimateImagePreviews = "never"
)

func (e *WebAnimateImagePreviews) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedWebAnimateImagePreviewsValues)
}

var allowedWebAnimateImagePreviewsValues = []WebAnimateImagePreviews{WebAnimateImagePreviewsAlways, WebAnimateImagePreviewsOnHover, WebAnimateImagePreviewsNever}

const (
	// NarrowOperatorChannel filters messages by channel ID or name.
	NarrowOperatorChannel NarrowOperator = "channel"
	// NarrowOperatorTopic filters messages by topic name within a channel.
	NarrowOperatorTopic NarrowOperator = "topic"
	// NarrowOperatorStream is a legacy alias for NarrowOperatorChannel.
	NarrowOperatorStream NarrowOperator = "stream"
	// NarrowOperatorId filters messages by message ID.
	NarrowOperatorId NarrowOperator = "id"
	// NarrowOperatorDm filters direct messages between two users.
	NarrowOperatorDm NarrowOperator = "dm"
	// NarrowOperatorDmIncluding filters messages in a group DM including specific users.
	NarrowOperatorDmIncluding NarrowOperator = "dm_including"
	// NarrowOperatorSender filters messages sent by a user (by email or ID).
	NarrowOperatorSender NarrowOperator = "sender"
	// NarrowOperatorWith filters direct messages with a specific user.
	NarrowOperatorWith NarrowOperator = "with"
	// NarrowOperatorHas filters messages that have a specific property (e.g., "reactions").
	NarrowOperatorHas NarrowOperator = "has"
	// NarrowOperatorIs filters messages by status (e.g., "muted", "starred", "alerted").
	NarrowOperatorIs NarrowOperator = "is"
	// NarrowOperatorSearch filters messages using full-text search.
	NarrowOperatorSearch NarrowOperator = "search"
)

func (e *NarrowOperator) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalStringEnum(data, e, allowedNarrowOperatorValues)
}

var allowedNarrowOperatorValues = []NarrowOperator{NarrowOperatorChannel, NarrowOperatorTopic, NarrowOperatorStream, NarrowOperatorId, NarrowOperatorDm, NarrowOperatorDmIncluding, NarrowOperatorSender, NarrowOperatorWith, NarrowOperatorHas, NarrowOperatorIs, NarrowOperatorSearch}
