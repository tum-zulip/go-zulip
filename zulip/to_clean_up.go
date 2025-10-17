package zulip

import "time"

// SubscriptionRemoveEvent7Data An object containing the properties that have changed.  **Changes**: In Zulip 10.0 (feature level 316), `edit_topic_policy` property was removed and replaced by `can_move_messages_between_topics_group` realm setting.  In Zulip 7.0 (feature level 183), the `community_topic_editing_limit_seconds` property was removed. It was documented as potentially returned as a changed property in this event, but in fact it was only ever returned in the [`POST /register`](zulip.com/api/register-queue) response.  Before Zulip 6.0 (feature level 150), on changing any of `allow_message_editing`, `message_content_edit_limit_seconds`, or `edit_topic_policy` settings, this object included all the three settings irrespective of which of these settings were changed. Now, a separate event is sent for each changed setting.
type SubscriptionRemoveEvent7Data struct {
	// Whether this organization's [message edit policy][config-message-editing] allows editing the content of messages.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  [config-message-editing]: /help/restrict-message-editing-and-deletion
	AllowMessageEditing bool `json:"allow_message_editing,omitempty"`
	// Dictionary of authentication method keys mapped to dictionaries that describe the properties of the named authentication method for the organization - its enabled status and availability for use by the organization.  Clients should use this to implement server-settings UI to change which methods are enabled for the organization. For authentication UI itself, clients should use the pre-authentication metadata returned by [`GET /server_settings`](zulip.com/api/get-server-settings.  **Changes**: In Zulip 9.0 (feature level 243), the values in this dictionary were changed. Previously, the values were a simple boolean indicating whether the backend is enabled or not.
	AuthenticationMethods               map[string]RealmAuthenticationMethod `json:"authentication_methods,omitempty"`
	CanAccessAllUsersGroup              GroupSettingValue                    `json:"can_access_all_users_group,omitempty"`
	CanCreateGroups                     GroupSettingValue                    `json:"can_create_groups,omitempty"`
	CanCreateBotsGroup                  GroupSettingValue                    `json:"can_create_bots_group,omitempty"`
	CanCreateWriteOnlyBotsGroup         GroupSettingValue                    `json:"can_create_write_only_bots_group,omitempty"`
	CanCreatePublicChannelGroup         GroupSettingValue                    `json:"can_create_public_channel_group,omitempty"`
	CanCreatePrivateChannelGroup        GroupSettingValue                    `json:"can_create_private_channel_group,omitempty"`
	CanCreateWebPublicChannelGroup      GroupSettingValue                    `json:"can_create_web_public_channel_group,omitempty"`
	CanAddCustomEmojiGroup              GroupSettingValue                    `json:"can_add_custom_emoji_group,omitempty"`
	CanAddSubscribersGroup              GroupSettingValue                    `json:"can_add_subscribers_group,omitempty"`
	CanDeleteAnyMessageGroup            GroupSettingValue                    `json:"can_delete_any_message_group,omitempty"`
	CanDeleteOwnMessageGroup            GroupSettingValue                    `json:"can_delete_own_message_group,omitempty"`
	CanSetDeleteMessagePolicyGroup      GroupSettingValue                    `json:"can_set_delete_message_policy_group,omitempty"`
	CanSetTopicsPolicyGroup             GroupSettingValue                    `json:"can_set_topics_policy_group,omitempty"`
	CanInviteUsersGroup                 GroupSettingValue                    `json:"can_invite_users_group,omitempty"`
	CanMentionManyUsersGroup            GroupSettingValue                    `json:"can_mention_many_users_group,omitempty"`
	CanMoveMessagesBetweenChannelsGroup GroupSettingValue                    `json:"can_move_messages_between_channels_group,omitempty"`
	CanMoveMessagesBetweenTopicsGroup   GroupSettingValue                    `json:"can_move_messages_between_topics_group,omitempty"`
	CanResolveTopicsGroup               GroupSettingValue                    `json:"can_resolve_topics_group,omitempty"`
	CanManageAllGroups                  GroupSettingValue                    `json:"can_manage_all_groups,omitempty"`
	CanManageBillingGroup               GroupSettingValue                    `json:"can_manage_billing_group,omitempty"`
	CanSummarizeTopicsGroup             GroupSettingValue                    `json:"can_summarize_topics_group,omitempty"`
	CreateMultiuseInviteGroup           GroupSettingValue                    `json:"create_multiuse_invite_group,omitempty"`
	// The default pygments language code to be used for code blocks in this organization. If an empty string, no default has been set.  **Changes**: Prior to Zulip 8.0 (feature level 195), a server bug meant that both `null` and an empty string could represent that no default was set for this realm setting in the [`POST /register`](zulip.com/api/register-queue) response. The documentation for both that endpoint and this event incorrectly stated that the only representation for no default language was `null`. This event in fact uses the empty string to indicate that no default has been set in all server versions.
	DefaultCodeBlockLanguage string `json:"default_code_block_language,omitempty"`
	// The default language for the organization.
	DefaultLanguage string `json:"default_language,omitempty"`
	// The description of the organization, used on login and registration pages.
	Description string `json:"description,omitempty"`
	// Whether the organization has enabled [weekly digest emails](zulip.com/help/digest-emails.
	DigestEmailsEnabled bool `json:"digest_emails_enabled,omitempty"`
	// The day of the week when the organization will send its weekly digest email to inactive users.
	DigestWeekday                int               `json:"digest_weekday,omitempty"`
	DirectMessageInitiatorGroup  GroupSettingValue `json:"direct_message_initiator_group,omitempty"`
	DirectMessagePermissionGroup GroupSettingValue `json:"direct_message_permission_group,omitempty"`
	// Whether the organization disallows disposable email addresses.
	DisallowDisposableEmailAddresses bool `json:"disallow_disposable_email_addresses,omitempty"`
	// Whether users are allowed to change their own email address in this organization. This is typically disabled for organizations that synchronize accounts from LDAP or a similar corporate database.
	EmailChangesDisabled bool `json:"email_changes_disabled,omitempty"`
	// Whether read receipts is enabled in the organization or not.  If disabled, read receipt data will be unavailable to clients, regardless of individual users' personal read receipt settings. See also the `send_read_receipts` setting within `realm_user_settings_defaults`.  **Changes**: New in Zulip 6.0 (feature level 137).
	EnableReadReceipts bool `json:"enable_read_receipts,omitempty"`
	// Whether [new users joining](zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions this organization are required to have an email address in one of the `realm_domains` configured for the organization.
	EmailsRestrictedToDomains bool `json:"emails_restricted_to_domains,omitempty"`
	// Whether clients should show a warning when a user is composing a DM to a guest user in this organization.  **Changes**: New in Zulip 10.0 (feature level 348).
	EnableGuestUserDmWarning bool `json:"enable_guest_user_dm_warning,omitempty"`
	// Whether clients should display \"(guest)\" after the names of guest users to prominently highlight their status.  **Changes**: New in Zulip 8.0 (feature level 216).
	EnableGuestUserIndicator bool `json:"enable_guest_user_indicator,omitempty"`
	// Whether web-public channels are enabled in this organization.  Can only be enabled if the `WEB_PUBLIC_STREAMS_ENABLED` [server setting][server-settings] is enabled on the Zulip server. See also the `can_create_web_public_channel_group` realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 109).
	EnableSpectatorAccess bool `json:"enable_spectator_access,omitempty"`
	// Maximum rating of the GIFs that will be retrieved from GIPHY.  **Changes**: New in Zulip 4.0 (feature level 55).
	GiphyRating int32 `json:"giphy_rating,omitempty"`
	// String indicating whether the organization's [profile icon](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's icon.  - \"G\" means generated by Gravatar (the default). - \"U\" means uploaded by an organization administrator.
	IconSource string `json:"icon_source,omitempty"`
	// The URL of the organization's [profile icon](zulip.com/help/create-your-organization-profile.
	IconUrl string `json:"icon_url,omitempty"`
	// Whether this organization has been configured to enable [previews of linked images](zulip.com/help/image-video-and-website-previews.
	InlineImagePreview bool `json:"inline_image_preview,omitempty"`
	// Whether this organization has been configured to enable [previews of linked websites](zulip.com/help/image-video-and-website-previews.
	InlineUrlEmbedPreview bool `json:"inline_url_embed_preview,omitempty"`
	// Whether an invitation is required to join this organization.
	InviteRequired bool `json:"invite_required,omitempty"`
	// The URL of the custom Jitsi Meet server configured in this organization's settings.  `null`, the default, means that the organization is using the should use the server-level configuration, `server_jitsi_server_url`.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this was only available as a server-level configuration, and required a server restart to change.
	JitsiServerUrl *string `json:"jitsi_server_url,omitempty"`
	// String indicating whether the organization's [profile wide logo](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - \"D\" means the logo is the default Zulip logo. - \"U\" means uploaded by an organization administrator.
	LogoSource string `json:"logo_source,omitempty"`
	// The URL of the organization's wide logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.
	LogoUrl string `json:"logo_url,omitempty"`
	// The organization's default policy for sending channel messages to the [empty \"general chat\" topic](zulip.com/help/general-chat-topic.  - `\"allow_empty_topic\"`: Channel messages can be sent to the empty topic. - `\"disable_empty_topic\"`: Channel messages cannot be sent to the empty topic.  **Changes**: New in Zulip 11.0 (feature level 392). Previously, this was controlled by the boolean realm `mandatory_topics` setting, which is now deprecated.
	TopicsPolicy TopicsPolicy `json:"topics_policy,omitempty"`
	// Whether [topics are required](zulip.com/help/require-topics) for messages in this organization.  **Changes**: Deprecated in Zulip 11.0 (feature level 392). This is now controlled by the realm `topics_policy` setting.
	// Deprecated
	MandatoryTopics bool `json:"mandatory_topics,omitempty"`
	// The new maximum file size that can be uploaded to this Zulip organization.  **Changes**: New in Zulip 10.0 (feature level 306). Previously, this field of the core state did not support being updated via the events system, as it was typically hardcoded for a given Zulip installation.
	MaxFileUploadSizeMib int64 `json:"max_file_upload_size_mib,omitempty"`
	// Whether notification emails in this organization are allowed to contain Zulip the message content, or simply indicate that a new message was sent.
	MessageContentAllowedInEmailNotifications bool `json:"message_content_allowed_in_email_notifications,omitempty"`
	// Messages sent more than this many seconds ago cannot be deleted with this organization's [message deletion policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be 0. A `null` value means no limit: messages can be deleted regardless of how long ago they were sent.  **Changes**: No limit was represented using the special value `0` before Zulip 5.0 (feature level 100).
	MessageContentDeleteLimitSeconds *int64 `json:"message_content_delete_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be edited with this organization's [message edit policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be `0`. A `null` value means no limit, so messages can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: Before Zulip 6.0 (feature level 138), no limit was represented using the special value `0`.
	MessageContentEditLimitSeconds *int64 `json:"message_content_edit_limit_seconds,omitempty"`
	// Which type of message edit history is configured to allow users to access [message edit history](zulip.com/help/view-a-messages-edit-history.  - \"all\" = All edit history is visible. - \"moves\" = Only moves are visible. - \"none\" = No edit history is visible.  **Changes**: New in Zulip 10.0 (feature level 358), replacing the previous `allow_edit_history` boolean setting; `true` corresponds to `all`, and `false` to `none`.
	MessageEditHistoryVisibilityPolicy string `json:"message_edit_history_visibility_policy,omitempty"`
	// The Id of the private channel to which messages flagged by users for moderation are sent. Moderators can use this channel to review and act on reported content.  Will be `-1` if moderation requests are disabled.  Clients should check whether moderation requests are disabled to determine whether to present a \"report message\" feature in their UI within a given organization.  **Changes**: New in Zulip 10.0 (feature level 331). Previously, no \"report message\" features existed in Zulip.
	ModerationRequestChannelId int64 `json:"moderation_request_channel_id,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved within a channel to another topic by users who have permission to do so based on this organization's [topic edit policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so message topics can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, this time limit was always 72 hours for users who were not administrators or moderators.
	MoveMessagesWithinStreamLimitSeconds *int64 `json:"move_messages_within_stream_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved between channels by users who have permission to do so based on this organization's [message move policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so messages can be moved regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, there was no time limit for moving messages between channels for users with permission to do so.
	MoveMessagesBetweenStreamsLimitSeconds *int64 `json:"move_messages_between_streams_limit_seconds,omitempty"`
	// The name of the organization, used in login pages etc.
	Name string `json:"name,omitempty"`
	// Indicates whether users are [allowed to change](zulip.com/help/restrict-name-and-email-changes) their name via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.
	NameChangesDisabled bool `json:"name_changes_disabled,omitempty"`
	// String indicating whether the organization's dark theme [profile wide logo](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - \"D\" means the logo is the default Zulip logo. - \"U\" means uploaded by an organization administrator.
	NightLogoSource string `json:"night_logo_source,omitempty"`
	// The URL of the organization's dark theme wide-format logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.
	NightLogoUrl string `json:"night_logo_url,omitempty"`
	// The Id of the channel to which automated messages announcing the [creation of new channels][new-channel-announce] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-channel-announce]: /help/configure-automated-notices#new-channel-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed `notifications_stream_id` to `new_stream_announcements_stream_id`.
	NewStreamAnnouncementsStreamId int64 `json:"new_stream_announcements_stream_id,omitempty"`
	// The [organization type](zulip.com/help/organization-type) for the realm.  - 0 = Unspecified - 10 = Business - 20 = Open-source project - 30 = Education (non-profit) - 35 = Education (for-profit) - 40 = Research - 50 = Event or conference - 60 = Non-profit (registered) - 70 = Government - 80 = Political group - 90 = Community - 100 = Personal - 1000 = Other  **Changes**: New in Zulip 6.0 (feature level 128).
	OrgType OrgType `json:"org_type,omitempty"`
	// The plan type of the organization.  - 1 = Self-hosted organization (SELF_HOSTED) - 2 = Zulip Cloud free plan (LIMITED) - 3 = Zulip Cloud Standard plan (STANDARD) - 4 = Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)
	PlanType PlanType `json:"plan_type,omitempty"`
	// Whether online presence of other users is shown in this organization.
	PresenceDisabled bool `json:"presence_disabled,omitempty"`
	// Whether push notifications are enabled for this organization. Typically `true` for Zulip Cloud and self-hosted realms that have a valid registration for the [Mobile push notifications service](https://zulip.readthedocs.io/en/latest/production/mobile-push-notifications.html), and `false` for self-hosted servers that do not.  **Changes**: New in Zulip 8.0 (feature level 231). Previously, this value was never updated via events.
	PushNotificationsEnabled bool `json:"push_notifications_enabled,omitempty"`
	// If the server expects the realm's push notifications access to end at a definite time in the future, the time at which this is expected to happen. Mobile clients should use this field to display warnings to users when the indicated timestamp is near.  **Changes**: New in Zulip 8.0 (feature level 231).
	PushNotificationsEnabledEndTimestamp *time.Time `json:"push_notifications_enabled_end_timestamp,omitempty"`
	// Whether this realm is configured to disallow sending mobile push notifications with message content through the legacy mobile push notifications APIs. The new API uses end-to-end encryption to protect message content and metadata from being accessible to the push bouncer service, APNs, and FCM. Clients that support the new E2EE API will use it automatically regardless of this setting.  If `true`, mobile push notifications sent to clients that lack support for E2EE push notifications will always have \"New message\" as their content. Note that these legacy mobile notifications will still contain metadata, which may include the message's Id, the sender's name, email address, and avatar.  In a future release, once the official mobile apps have implemented fully validated their E2EE protocol support, this setting will become strict, and disable the legacy protocol entirely.  **Changes**: New in Zulip 11.0 (feature level 409). Previously, this behavior was available only via the `PUSH_NOTIFICATION_REDACT_CONTENT` global server setting.
	RequireE2eePushNotifications bool `json:"require_e2ee_push_notifications,omitempty"`
	// Indicates whether the organization is configured to require users to have unique full names. If true, the server will reject attempts to create a new user, or change the name of an existing user, where doing so would lead to two users whose names are identical modulo case and unicode normalization.  **Changes**: New in Zulip 9.0 (feature level 246). Previously, the Zulip server could not be configured to enforce unique names.
	RequireUniqueNames bool `json:"require_unique_names,omitempty"`
	// Whether or not this organization is configured to send the standard Zulip [welcome emails](zulip.com/help/disable-welcome-emails) to new users joining the organization.
	SendWelcomeEmails bool `json:"send_welcome_emails,omitempty"`
	// The Id of the channel to which automated messages announcing that [new users have joined the organization][new-user-announce] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-user-announce]: /help/configure-automated-notices#new-user-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed `signup_notifications_stream_id` to `signup_announcements_stream_id`.
	SignupAnnouncementsStreamId int64 `json:"signup_announcements_stream_id,omitempty"`
	// The new upload quota for the Zulip organization.  If `null`, there is no limit.  **Changes**: New in Zulip 10.0 (feature level 306). Previously, this was present changed via an `upload_quota` field in `extra_data` property of [realm/update](#realm-update) event format for `plan_type` events.
	UploadQuotaMib *int64 `json:"upload_quota_mib,omitempty"`
	// The configured [video call provider](zulip.com/help/configure-call-provider) for the organization.  - 0 = None - 1 = Jitsi Meet - 3 = Zoom (User OAuth integration) - 4 = BigBlueButton - 5 = Zoom (Server to Server OAuth integration)  Note that only one of the [Zoom integrations][zoom-video-calls] can be configured on a Zulip server.  **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.  [zoom-video-calls]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom
	VideoChatProvider VideoChatProvider `json:"video_chat_provider,omitempty"`
	// Members whose accounts have been created at least this many days ago will be treated as [full members][calc-full-member] for the purpose of settings that restrict access to new members.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	WaitingPeriodThreshold int64 `json:"waiting_period_threshold,omitempty"`
	// Whether the organization has given permission to be advertised in the Zulip [communities directory](zulip.com/help/communities-directory.  **Changes**: New in Zulip 6.0 (feature level 129).
	WantAdvertiseInCommunitiesDirectory bool `json:"want_advertise_in_communities_directory,omitempty"`
	// This organization's configured custom message for Welcome Bot to send to new user accounts, in Zulip Markdown format.  Maximum length is 8000 characters.  **Changes**: New in Zulip 11.0 (feature level 416).
	WelcomeMessageCustomText string `json:"welcome_message_custom_text,omitempty"`
	// The Id of the channel to which automated messages announcing new features or other end-user updates about the Zulip software are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  **Changes**: New in Zulip 9.0 (feature level 242).
	ZulipUpdateAnnouncementsStreamId int64 `json:"zulip_update_announcements_stream_id,omitempty"`
}

// UpdateFlagsNarrowClause - struct for UpdateFlagsNarrowClause
type UpdateFlagsNarrowClause struct {
	UpdateFlagsNarrowFilter *UpdateFlagsNarrowFilter
	ArrayOfString           *[]string
}

// UpdateFlagsNarrowFilter struct for UpdateFlagsNarrowFilter
type UpdateFlagsNarrowFilter struct {
	Operator string                   `json:"operator"`
	Operand  UpdateFlagsNarrowOperand `json:"operand"`
	Negated  *bool                    `json:"negated,omitempty"`
}

// UpdateFlagsNarrowOperand - struct for UpdateFlagsNarrowOperand
type UpdateFlagsNarrowOperand struct {
	ArrayOfInt32 *[]int32
	Int32        *int32
	String       *string
}

// UserSettingsUpdateEvent1MutedTopicsInnerInner - struct for UserSettingsUpdateEvent1MutedTopicsInnerInner
type UserSettingsUpdateEvent1MutedTopicsInnerInner struct {
	Int32  *int32
	String *string
}

// RealmFilterTuple - struct for RealmFilterTuple
type RealmFilterTuple struct {
	Int32  *int32
	String *string
}

// DefaultChannelGroup Dictionary containing details of a default channel group.
type DefaultChannelGroup struct {
	// Name of the default channel group.
	Name *string `json:"name,omitempty"`
	// Description of the default channel group.
	Description *string `json:"description,omitempty"`
	// The Id of the default channel group.
	Id *int32 `json:"id,omitempty"`
	// An array of Ids of all the channels in the default stream group.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single stream in the default stream group.
	Streams []int32 `json:"streams,omitempty"`
}

// UnreadMsgs Present if `message` and `update_message_flags` are both present in `event_types`.  A set of data structures describing the conversations containing the 50000 most recent unread messages the user has received. This will usually contain every unread message the user has received, but clients should support users with even more unread messages (and not hardcode the number 50000).
type UnreadMsgs struct {
	// The total number of unread messages to display. This includes one-on-one and group direct messages, as well as channel messages that are not [muted](zulip.com/help/mute-a-topic.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.
	Count *int32 `json:"count,omitempty"`
	// An array of objects where each object contains details of unread one-on-one direct messages with a specific user.  Note that it is possible for a message that the current user sent to the specified user to be marked as unread and thus appear here.
	Pms []UnreadMsgsPms `json:"pms,omitempty"`
	// An array of dictionaries where each dictionary contains details of all unread messages of a single subscribed channel. This includes muted channels and muted topics, even though those messages are excluded from `count`.  **Changes**: Prior to Zulip 5.0 (feature level 90), these objects included a `sender_ids` property, which listed the set of Ids of users who had sent the unread messages.
	Streams []UnreadMsgsStreams `json:"streams,omitempty"`
	// An array of objects where each object contains details of unread group direct messages with a specific group of users.
	Huddles []UnreadMsgsHuddles `json:"huddles,omitempty"`
	// Array containing the Ids of all unread messages in which the user was mentioned directly, and unread [non-muted](zulip.com/help/mute-a-topic) messages in which the user was mentioned through a wildcard.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.
	Mentions []int32 `json:"mentions,omitempty"`
	// Whether this data set was truncated because the user has too many unread messages. When truncation occurs, only the most recent `MAX_UNREAD_MESSAGES` (currently 50000) messages will be considered when forming this response. When `true`, we recommend that clients display a warning, as they are likely to produce erroneous results until reloaded with the user having fewer than `MAX_UNREAD_MESSAGES` unread messages.  **Changes**: New in Zulip 4.0 (feature level 44).
	OldUnreadsMissing *bool `json:"old_unreads_missing,omitempty"`
}

// UnreadMsgsPms struct for UnreadMsgsPms
type UnreadMsgsPms struct {
	// The user Id of the other participant in this one-on-one direct message conversation. Will be the current user's Id for messages that they sent in a one-on-one direct message conversation with themself.  **Changes**: New in Zulip 5.0 (feature level 119), replacing the less clearly named `sender_id` field.
	OtherUserId *int32 `json:"other_user_id,omitempty"`
	// Old name for the `other_user_id` field. Clients should access this field in Zulip server versions that do not yet support `other_user_id`.  **Changes**: Deprecated in Zulip 5.0 (feature level 119). We expect to provide a next version of the full `unread_msgs` API before removing this legacy name.
	// Deprecated
	SenderId *int32 `json:"sender_id,omitempty"`
	// The message Ids of the recent unread direct messages sent by either user in this one-on-one direct message conversation, sorted in ascending order.
	UnreadMessageIds []int32 `json:"unread_message_ids,omitempty"`
}

// UnreadMsgsStreams struct for UnreadMsgsStreams
type UnreadMsgsStreams struct {
	// The topic under which the messages were sent.  Note that the empty string topic may have been rewritten by the server to the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue) response depending on the value of the `empty_topic_name` [client capability][client-capabilities].  **Changes**: The `empty_topic_name` client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	Topic *string `json:"topic,omitempty"`
	// The Id of the channel to which the messages were sent.
	StreamId *int64 `json:"stream_id,omitempty"`
	// The message Ids of the recent unread messages sent in this channel, sorted in ascending order.
	UnreadMessageIds []int32 `json:"unread_message_ids,omitempty"`
}

// UnreadMsgsHuddles struct for UnreadMsgsHuddles
type UnreadMsgsHuddles struct {
	// A string containing the Ids of all users in the group direct message conversation, including the current user, separated by commas and sorted numerically; for example: `\"1,2,3\"`.
	UserIdsString *string `json:"user_ids_string,omitempty"`
	// The message Ids of the recent unread messages which have been sent in this group direct message conversation, sorted in ascending order.
	UnreadMessageIds []int32 `json:"unread_message_ids,omitempty"`
}

// UserTopics Object describing the user's configuration for a given topic.
type UserTopics struct {
	// The Id of the channel to which the topic belongs.
	StreamId *int64 `json:"stream_id,omitempty"`
	// The name of the topic.  Note that the empty string topic may have been rewritten by the server to the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue) response depending on the value of the `empty_topic_name` [client capability][client-capabilities].  **Changes**: The `empty_topic_name` client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	TopicName *string `json:"topic_name,omitempty"`
	// An integer UNIX timestamp representing when the user-topic relationship was changed.
	LastUpdated *int32 `json:"last_updated,omitempty"`
	// An integer indicating the user's visibility configuration for the topic.  - 1 = Muted. Used to record [muted topics](zulip.com/help/mute-a-topic. - 2 = Unmuted. Used to record [unmuted topics](zulip.com/help/mute-a-topic. - 3 = Followed. Used to record [followed topics](zulip.com/help/follow-a-topic.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.
	VisibilityPolicy *int32 `json:"visibility_policy,omitempty"`
}

// RealmBilling Present if `realm_billing` is present in `fetch_event_types`.  A dictionary containing billing information of the organization.  **Changes**: New in Zulip 10.0 (feature level 363).
type RealmBilling struct {
	// Whether there is a pending sponsorship request for the organization. Note that this field will always be `false` if the user is not in `can_manage_billing_group`.  **Changes**: New in Zulip 10.0 (feature level 363).
	HasPendingSponsorshipRequest bool `json:"has_pending_sponsorship_request,omitempty"`
}

// RealmIncomingWebhookBot Object containing details of the bot.
type RealmIncomingWebhookBot struct {
	// A machine-readable unique name identifying the integration, all-lower-case without spaces.
	Name string `json:"name,omitempty"`
	// A human-readable display name identifying the integration that this bot implements, intended to be used in menus for selecting which integration to create.  **Changes**: New in Zulip 8.0 (feature level 207).
	DisplayName string `json:"display_name,omitempty"`
	// For incoming webhook integrations that support the Zulip server filtering incoming events, the list of event types supported by it.  A null value will be present if this incoming webhook integration doesn't support such filtering.  **Changes**: New in Zulip 8.0 (feature level 207).
	AllEventTypes []string `json:"all_event_types,omitempty"`
	// An array of configuration options that can be set when creating a bot user for this incoming webhook integration.  This is an unstable API. Please discuss in chat.zulip.org before using it.  **Changes**: As of Zulip 11.0 (feature level 403), this object is reserved for integration-specific configuration options that can be set when creating a bot user. Previously, this object also included optional webhook URL parameters, which are now specified in the `url_options` object.  Before Zulip 10.0 (feature level 318), this field was named `config`, and was reserved for configuration data key-value pairs.
	ConfigOptions []WebhookOption `json:"config_options,omitempty"`
	// An array of optional URL parameter options for the incoming webhook integration. In the web app, these are used when [generating a URL for an integration](zulip.com/help/generate-integration-url.  This is an unstable API expected to be used only by the Zulip web app. Please discuss in chat.zulip.org before using it.  **Changes**: New in Zulip 11.0 (feature level 403). Previously, these optional URL parameter options were included in the `config_options` object.
	UrlOptions []WebhookOption `json:"url_options,omitempty"`
}

// WebhookConfigOptionInner struct for WebhookConfigOptionInner
type WebhookOption struct {
	// A key for the configuration option.
	Key string `json:"key,omitempty"`
	// A human-readable label of the configuration option.
	Label string `json:"label,omitempty"`
	// The name of the validator function for the configuration option.
	Validator string `json:"validator,omitempty"`
}

// SendInvites400Response - struct for SendInvites400Response
type SendInvites400Response struct {
	CodedError            *CodedError
	InvitationFailedError *InvitationFailedError
}

// SendMessage400Response - struct for SendMessage400Response
type SendMessage400Response struct {
	CodedError                  *CodedError
	NonExistingChannelNameError *NonExistingChannelNameError
}

// CreateScheduledMessage400Response - struct for CreateScheduledMessage400Response
type CreateScheduledMessage400Response struct {
	CodedError                *CodedError
	NonExistingChannelIdError *NonExistingChannelIdError
}

// RestErrorHandling400Response - struct for RestErrorHandling400Response
type RestErrorHandling400Response struct {
	IncompatibleParametersError *IncompatibleParametersError
	CodedError                  *CodedError
	MissingArgumentError        *MissingArgumentError
}

// RealmDefaultExternalAccounts `{site_name}`: Dictionary containing the details of the default external account provider with the name of the website as the key.
type RealmDefaultExternalAccounts struct {
	// The name of the external account provider
	Name *string `json:"name,omitempty"`
	// The text describing the external account.
	Text *string `json:"text,omitempty"`
	// The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields for this account.
	Hint *string `json:"hint,omitempty"`
	// The regex pattern of the URL of a profile page on the external site.
	UrlPattern *string `json:"url_pattern,omitempty"`
}

// GiphyRatingOptionsValue `{rating_name}`: Dictionary containing the details of the rating with the name of the rating as the key.
type GiphyRatingOptionsValue struct {
	// The description of the rating option.
	Name *string `json:"name,omitempty"`
	// The Id of the rating option.
	Id *int32 `json:"id,omitempty"`
}

// UserUpdateEventEnvalop - Object containing the changed details of the user. It has multiple forms depending on the value changed.  **Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was replaced by the `can_manage_billing_group` realm setting.
type UserUpdateEventEnvalop struct {
	UserUpdateEventFullName      *UserUpdateEventFullName
	UserUpdateEventAvatar        *UserUpdateEventAvatar
	UserUpdateEventTimezone      *UserUpdateEventTimezone
	UserUpdateEventBotOwner      *UserUpdateEventBotOwner
	UserUpdateEventRole          *UserUpdateEventRole
	UserUpdateEventDeliveryEmail *UserUpdateEventDeliveryEmail
	UserUpdateEventCustomField   *UserUpdateEventCustomField
	UserUpdateEventEmail         *UserUpdateEventEmail
	UserUpdateEventActivation    *UserUpdateEventActivation
}

// UserUpdateEventFullName When a user changes their full name.
type UserUpdateEventFullName struct {
	// The Id of modified user.
	UserId *int32 `json:"user_id,omitempty"`
	// The new full name for the user.
	FullName *string `json:"full_name,omitempty"`
}

// UserUpdateEventAvatar When a user changes their avatar.
type UserUpdateEventAvatar struct {
	// The Id of the user who is affected by this change.
	UserId *int32 `json:"user_id,omitempty"`
	// The URL of the new avatar for the user.
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// The new avatar data source type for the user.  Value values are `G` (gravatar) and `U` (uploaded by user).
	AvatarSource *string `json:"avatar_source,omitempty"`
	// The new medium-size avatar URL for user.
	AvatarUrlMedium *string `json:"avatar_url_medium,omitempty"`
	// The version number for the user's avatar. This is useful for cache-busting.
	AvatarVersion *int32 `json:"avatar_version,omitempty"`
}

// UserUpdateEventTimezone When a user changes their [profile time zone](zulip.com/help/change-your-timezone.
type UserUpdateEventTimezone struct {
	// The Id of modified user.
	UserId *int32 `json:"user_id,omitempty"`
	// The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the `user_id`.
	// Deprecated
	Email *string `json:"email,omitempty"`
	// The IANA identifier of the new profile time zone for the user.
	Timezone *string `json:"timezone,omitempty"`
}

// UserUpdateEventBotOwner When the owner of a bot changes.
type UserUpdateEventBotOwner struct {
	// The Id of the user/bot whose owner has changed.
	UserId *int32 `json:"user_id,omitempty"`
	// The user Id of the new bot owner.
	BotOwnerId *int32 `json:"bot_owner_id,omitempty"`
}

// UserUpdateEventRole When the [role](zulip.com/help/user-roles) of a user changes.
type UserUpdateEventRole struct {
	// The Id of the user affected by this change.
	UserId *int32 `json:"user_id,omitempty"`
	// The new [role](zulip.com/api/roles-and-permissions) of the user.
	Role *int32 `json:"role,omitempty"`
}

// UserUpdateEventDeliveryEmail When the value of a user's delivery email as visible to you changes, either due to the email address changing or your access to the user's email changing via an update to their `email_address_visibility` setting.  **Changes**: Prior to Zulip 7.0 (feature level 163), this event was sent only to the affected user, and this event would only be triggered by changing the affected user's delivery email.
type UserUpdateEventDeliveryEmail struct {
	// The Id of the user affected by this change.
	UserId *int32 `json:"user_id,omitempty"`
	// The new delivery email of the user.  This value can be `null` if the affected user changed their `email_address_visibility` setting such that you cannot access their real email.  **Changes**: Before Zulip 7.0 (feature level 163), `null` was not a possible value for this event as it was only sent to the affected user when their email address was changed.
	DeliveryEmail *string `json:"delivery_email,omitempty"`
}

// UserUpdateEventCustomField When the user updates one of their custom profile fields.
type UserUpdateEventCustomField struct {
	// The Id of the user affected by this change.
	UserId             *int32                             `json:"user_id,omitempty"`
	CustomProfileField *UserUpdateEventCustomFieldDetails `json:"custom_profile_field,omitempty"`
}

// UserUpdateEventEmail When the Zulip API email address of a user changes, either due to the user's email address changing, or due to changes in the user's [email address visibility][help-email-visibility].  [help-email-visibility]: /help/configure-email-visibility
type UserUpdateEventEmail struct {
	// The Id of the user affected by this change.
	UserId *int32 `json:"user_id,omitempty"`
	// The new value of `email` for the user. The client should update any data structures associated with this user to use this new value as the user's Zulip API email address.
	NewEmail *string `json:"new_email,omitempty"`
}

// UserUpdateEventActivation When a user is deactivated or reactivated. Only users who can access the modified user under the organization's `can_access_all_users_group` policy will receive this event.  Clients receiving a deactivation event should remove the user from all user groups in their data structures, because deactivated users cannot be members of groups.  **Changes**: Prior to Zulip 10.0 (feature level 303), reactivation events were sent to users who could not access the reactivated user due to a `can_access_all_users_group` policy. Also, previously, Clients were not required to update group membership records during user deactivation.  New in Zulip 8.0 (feature level 222). Previously the server sent a `realm_user` event with `op` field set to `remove` when deactivating a user and a `realm_user` event with `op` field set to `add` when reactivating a user.
type UserUpdateEventActivation struct {
	// The Id of the user affected by this change.
	UserId *int32 `json:"user_id,omitempty"`
	// A boolean describing whether the user account has been deactivated.
	IsActive *bool `json:"is_active,omitempty"`
}

// UserUpdateEventCustomFieldDetails Object containing details about the custom profile data change.
type UserUpdateEventCustomFieldDetails struct {
	// The Id of the custom profile field which user updated.
	Id *int32 `json:"id,omitempty"`
	// User's personal value for this custom profile field, or `null` if unset.
	Value *string `json:"value,omitempty"`
	// The `value` rendered in HTML. Will only be present for custom profile field types that support Markdown rendering.  This user-generated HTML content should be rendered using the same CSS and client-side security protections as are used for message content.  See [Markdown message formatting](zulip.com/api/message-formatting) for details on Zulip's HTML format.
	RenderedValue *string `json:"rendered_value,omitempty"`
}
