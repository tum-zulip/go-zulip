package zulip

import "time"

type Realm struct {
	// The UNIX timestamp (UTC) for when the organization was created.
	//
	// **Changes**: New in Zulip 8.0 (feature level 203).
	DateCreated time.Time `json:"realm_date_created,omitempty"`
	// The maximum allowed length for a channel name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required `stream` in `fetch_event_types`, was called `stream_name_max_length`, and always had a value of 60.
	MaxChannelNameLength int64 `json:"max_stream_name_length,omitempty"`
	// The maximum allowed length for a channel description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required `stream` in `fetch_event_types`, was called `stream_description_max_length`, and always had a value of 1024.
	MaxStreamDescriptionLength int64 `json:"max_stream_description_length,omitempty"`
	// The maximum allowed length for a channel folder name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 60 as a fallback value on previous feature levels.
	MaxChannelFolderNameLength int64 `json:"max_channel_folder_name_length,omitempty"`
	// The maximum allowed length for a channel folder description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 1024 as a fallback value on previous feature levels.
	MaxChannelFolderDescriptionLength int64 `json:"max_channel_folder_description_length,omitempty"`
	// The maximum allowed length for a topic, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 60.
	MaxTopicLength int64 `json:"max_topic_length,omitempty"`
	// The maximum allowed length for a message, in Unicode code points. Clients should use this property rather than hardcoding field sizes.
	//
	// **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 10000.
	MaxMessageLength int64 `json:"max_message_length,omitempty"`
	// The minimum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If `null`, then a deactivated organization's data can be deleted immediately.
	//
	// **Changes**: New in Zulip 10.0 (feature level 332)
	ServerMinDeactivatedRealmDeletionDays int64 `json:"server_min_deactivated_realm_deletion_days,omitempty"`
	// The maximum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If `null`, then a deactivated organization's data can be retained indefinitely.
	//
	// **Changes**: New in Zulip 10.0 (feature level 332).
	ServerMaxDeactivatedRealmDeletionDays int64 `json:"server_max_deactivated_realm_deletion_days,omitempty"`
	// For clients implementing the [presence] system, the time interval the client should use for sending presence requests to the server (and thus receive presence updates from the server).  It is important for presence implementations to use both this and `server_presence_offline_threshold_seconds` correctly, so that a Zulip server can change these values to manage the trade-off between load and freshness of presence data.
	//
	// **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 60 for older Zulip servers, since that's the value that was hardcoded in the Zulip mobile apps prior to this parameter being introduced.
	//
	// [presence]: https://zulip.com/api/get-presence
	ServerPresencePingIntervalSeconds int64 `json:"server_presence_ping_interval_seconds,omitempty"`
	// How old a presence timestamp for a given user can be before the user should be displayed as offline by clients displaying Zulip presence data. See the related `server_presence_ping_interval_seconds` for details.
	//
	// **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 140 for older Zulip servers, since that's the value that was hardcoded in the Zulip client apps prior to this parameter being introduced.
	ServerPresenceOfflineThresholdSeconds int64 `json:"server_presence_offline_threshold_seconds,omitempty"`
	// For clients implementing [typing notifications] protocol, the time interval in milliseconds that the client should wait for additional [typing start] events from the server before removing an active typing indicator.
	//
	// **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 15000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	//
	// [typing notifications]: https://zulip.com/api/set-typing-status
	// [typing start]: https://zulip.com/api/get-events#typing-start
	ServerTypingStartedExpiryPeriodMilliseconds int64 `json:"server_typing_started_expiry_period_milliseconds,omitempty"`
	// For clients implementing [typing notifications] protocol, the time interval in milliseconds that the client should wait when a user stops interacting with the compose UI before sending a stop notification to the server.
	//
	// **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 5000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	//
	// [typing notifications]: https://zulip.com/api/set-typing-status
	ServerTypingStoppedWaitPeriodMilliseconds int64 `json:"server_typing_stopped_wait_period_milliseconds,omitempty"`
	// For clients implementing [typing notifications] protocol, the time interval in milliseconds that the client should use to send regular start notifications to the server to indicate that the user is still actively interacting with the compose UI.
	//
	// **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 10000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	//
	// [typing notifications]: https://zulip.com/api/set-typing-status
	ServerTypingStartedWaitPeriodMilliseconds int64 `json:"server_typing_started_wait_period_milliseconds,omitempty"`

	// A deprecated representation of a superset of the users who have permission to create public channels in the organization, available for backwards-compatibility. Clients should use `can_create_public_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:
	//   - 1 = Members only
	//   - 2 = Admins only
	//   - 3 = [Full members] only
	//   - 4 = Admins and moderators only
	//
	// **Changes**: Deprecated in Zulip 9.0 (feature level 264) and replaced by `realm_can_create_public_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the `realm_create_stream_policy` setting.
	//
	// [permission-level]: https://zulip.com/api/roles-and-permissions#permission-levels
	// [Full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	CreatePublicStreamPolicy int32 `json:"realm_create_public_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to create private channels in the organization, available for backwards-compatibility. Clients should use `can_create_private_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:
	//   - 1 = Members only
	//   - 2 = Admins only
	//   - 3 = [Full members] only
	//   - 4 = Admins and moderators only
	//
	// **Changes**: Deprecated in Zulip 9.0 (feature level 266) and replaced by `realm_can_create_private_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.
	//
	// **Changes**: Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the `realm_create_stream_policy` setting.
	//
	// [permission-level]: https://zulip.com/api/roles-and-permissions#permission-levels
	// [Full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	CreatePrivateStreamPolicy int32 `json:"realm_create_private_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to create web-public channels in the organization, available for backwards-compatibility. Clients should use `can_create_web_public_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:
	//   - 2 = Admins only
	//   - 4 = Admins and moderators only
	//   - 6 = Nobody
	//   - 7 = Owners only
	//
	// **Changes**: Deprecated in Zulip 10.0 (feature level 280) and replaced by `realm_can_create_web_public_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.
	//
	// **Changes**: Added in Zulip 5.0 (feature level 103).
	//
	// [permission-level]: https://zulip.com/api/roles-and-permissions#permission-levels
	// Deprecated
	CreateWebPublicStreamPolicy int32 `json:"realm_create_web_public_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to use wildcard mentions in large channels, available for backwards-compatibility. Clients should use `can_mention_many_users_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:
	//   - 1 = Any user can use wildcard mentions in large channels.
	//   - 2 = Only members can use wildcard mentions in large channels.
	//   - 3 = Only [full members] can use wildcard mentions in large channels.
	//   - 5 = Only organization administrators can use wildcard mentions in large channels.
	//   - 6 = Nobody can use wildcard mentions in large channels.
	//   - 7 = Only organization administrators and moderators can use wildcard mentions in large channels.
	//
	// All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.
	//
	// **Changes**: Deprecated in Zulip 10.0 (feature level 352) and replaced by `realm_can_mention_many_users_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Channel administrators option removed in Zulip 6.0 (feature level 133).  Moderators option added in Zulip 4.0 (feature level 62).  New in Zulip 4.0 (feature level 33).
	//
	// [permission-level]: https://zulip.com/api/roles-and-permissions#permission-levels
	// [full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	WildcardMentionPolicy int32 `json:"realm_wildcard_mention_policy,omitempty"`

	// Whether this organization is configured to allow users to access [message edit history].  The value of `realm_allow_edit_history` is set as `false` if the `realm_message_edit_history_visibility_policy` is configured as "None" and `true` if it is configured as "Moves only" or "All".
	//
	// **Changes**: Deprecated in Zulip 10.0 (feature level 358) and will be removed in the future, as it is an inaccurate version `realm_message_edit_history_visibility_policy`, which replaces this field.
	// Deprecated
	//
	// [message edit history]: https://zulip.com/help/view-a-messages-edit-history
	AllowEditHistory bool `json:"realm_allow_edit_history,omitempty"`

	RealmConfiguration

	// The default [message retention policy] for this organization. It can have one special value:  - `-1` denoting that the messages will be retained forever for this realm, by default.
	//
	// **Changes**: Prior to Zulip 3.0 (feature level 22), no limit was encoded as `null` instead of `-1`. Clients can correctly handle all server versions by treating both `-1` and `null` as indicating unlimited message retention.
	//
	// [message retention policy]: https://zulip.com/help/message-retention-policy
	MessageRetentionDays int64 `json:"realm_message_retention_days,omitempty"`

	// Indicates whether users are [allowed to change] their avatar via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.
	//
	// [allowed to change]: https://zulip.com/help/restrict-name-and-email-changes
	AvatarChangesDisabled bool `json:"realm_avatar_changes_disabled,omitempty"`

	// The maximum file size allowed for the organization's icon. Useful for UI allowing editing the organization's icon.
	//
	// **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called `max_icon_file_size`.
	MaxIconFileSizeMib int64 `json:"max_icon_file_size_mib,omitempty"`
	// The maximum file size allowed for the uploaded organization logos.
	//
	// **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called `max_logo_file_size`.
	MaxLogoFileSizeMib int64 `json:"max_logo_file_size_mib,omitempty"`
	// The fake email domain that will be used for new bots created this organization. Useful for UI for creating bots.
	BotDomain string `json:"realm_bot_domain,omitempty"`
	// The URL for the organization. Alias of `realm_url`.
	//
	// **Changes**: Deprecated in Zulip 9.0 (feature level 257). The term "URI" is deprecated in [web standards].
	// Deprecated
	//
	// [web standards]: https://url.spec.whatwg.org/#goals
	Uri string `json:"realm_uri,omitempty"`
	// The URL for the organization.
	//
	// **Changes**: New in Zulip 9.0 (feature level 257), replacing the deprecated `realm_uri`.
	Url string `json:"realm_url,omitempty"`
	// Dictionary where each entry describes a supported [video call provider] that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the realm setting `video_chat_provider`.
	//
	// [video call provider]: https://zulip.com/help/configure-call-provider
	AvailableVideoChatProviders map[string]VideoChatProviderInfo `json:"realm_available_video_chat_providers,omitempty"`
	// Whether this Zulip server is configured to allow organizations to enable [digest emails].  Relevant for administrative settings UI that can change the digest email settings.
	//
	// [digest emails]: https://zulip.com/help/digest-emails
	SettingsSendDigestEmails bool `json:"settings_send_digest_emails,omitempty"`
	// Whether the organization is a Zephyr mirror realm.
	IsZephyrMirrorRealm bool `json:"realm_is_zephyr_mirror_realm,omitempty"`
	// Whether the organization has enabled Zulip's default email and password authentication feature. Determines whether Zulip stores a password for the user and clients should offer any UI for changing the user's Zulip password.
	EmailAuthEnabled bool `json:"realm_email_auth_enabled,omitempty"`
	// Whether the organization allows any sort of password-based authentication (whether via EmailAuthBackend or LDAP passwords).  Determines whether a client might ever need to display a password prompt (clients will primarily look at this attribute in [server_settings] before presenting a login page).
	//
	// [server_settings]: https://zulip.com/api/get-server-settings
	PasswordAuthEnabled bool `json:"realm_password_auth_enabled,omitempty"`
	// Whether the organization is using a limited (Zulip Cloud Free) plan.
	ZulipPlanIsNotLimited bool `json:"zulip_plan_is_not_limited,omitempty"`
	// Text to use when displaying UI for wide organization logos, a feature that is currently not available on the Zulip Cloud Free plan.  Useful only for clients supporting administrative UI for uploading a new wide organization logo to brand the organization.
	UpgradeTextForWideOrganizationLogo string `json:"upgrade_text_for_wide_organization_logo,omitempty"`
	// Dictionary where each entry describes a default external account type that can be configured with Zulip's [custom profile fields feature].
	//
	// **Changes**: New in Zulip 2.1.0.
	//
	// [custom profile fields feature]: https://zulip.com/help/custom-profile-fields
	DefaultExternalAccounts map[string]RealmDefaultExternalAccounts `json:"realm_default_external_accounts,omitempty"`
	// The base URL to be used to create Jitsi video calls. Equals `realm_jitsi_server_url || server_jitsi_server_url`.
	//
	// **Changes**: Deprecated in Zulip 8.0 (feature level 212) and will eventually be removed. Previously, the Jitsi server to use was not configurable on a per-realm basis, and this field contained the server's configured Jitsi server. (Which is now provided as `server_jitsi_server_url`). Clients supporting older versions should fall back to this field when creating calls: using `realm_jitsi_server_url || server_jitsi_server_url` with newer servers and using `jitsi_server_url` with servers below feature level 212.
	// Deprecated
	JitsiServerUrlDeprecated string `json:"jitsi_server_url,omitempty"`
	// Whether this Zulip server is a development environment. Used to control certain features or UI (such as error popups) that should only apply when connected to a Zulip development environment.
	DevelopmentEnvironment bool `json:"development_environment,omitempty"`
	// A timestamp indicating when the process hosting this event queue was started. Clients will likely only find this value useful for inclusion in detailed error reports.
	ServerGeneration time.Time `json:"server_generation,omitempty"`
	// This Zulip server's configured minimum required length for passwords. Necessary for password change UI to show whether the password will be accepted.
	PasswordMinLength int64 `json:"password_min_length,omitempty"`
	// This Zulip server's configured maximum length for passwords. Necessary for password change UI to show whether the password will be accepted.
	//
	// **Changes**: New in Zulip 10.0 (feature level 338).
	PasswordMaxLength int64 `json:"password_max_length,omitempty"`
	// This Zulip server's configured minimum `zxcvbn` minimum guesses. Necessary for password change UI to show whether the password will be accepted.
	PasswordMinGuesses int64 `json:"password_min_guesses,omitempty"`
	// Dictionary where each entry describes a valid rating that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the allowed rating of GIFs.
	GiphyRatingOptions map[string]GiphyRatingOptionsValue `json:"giphy_rating_options,omitempty"`
	// The maximum avatar size that can be uploaded to this Zulip server.
	MaxAvatarFileSizeMib int64 `json:"max_avatar_file_size_mib,omitempty"`
	// Whether the server is configured with support for inline image previews. Clients containing administrative UI for changing `realm_inline_image_preview` should consult this field before offering that feature.
	ServerInlineImagePreview bool `json:"server_inline_image_preview,omitempty"`
	// Whether the server is configured with support for inline URL previews. Clients containing administrative UI for changing `realm_inline_url_embed_preview` should consult this field before offering that feature.
	ServerInlineUrlEmbedPreview bool `json:"server_inline_url_embed_preview,omitempty"`
	// A list describing the image formats that uploaded images will be thumbnailed into. Any image with a source starting with `/user_uploads/thumbnail/` can have its last path component replaced with any of the names contained in this list, to obtain the desired thumbnail size.
	//
	// **Changes**: New in Zulip 9.0 (feature level 273).
	ServerThumbnailFormats []ServerThumbnailFormat `json:"server_thumbnail_formats,omitempty"`
	// Whether the server allows avatar changes. Similar to `realm_avatar_changes_disabled` but based on the `AVATAR_CHANGES_DISABLED` Zulip server level setting.
	ServerAvatarChangesDisabled bool `json:"server_avatar_changes_disabled,omitempty"`
	// Whether the server allows name changes. Similar to `realm_name_changes_disabled` but based on the `NAME_CHANGES_DISABLED` Zulip server level setting.
	ServerNameChangesDisabled bool `json:"server_name_changes_disabled,omitempty"`
	// Whether the server is running an old version based on the Zulip [server release lifecycle], such that the web app will display to the current user a prominent warning.
	//
	// **Changes**: New in Zulip 5.0 (feature level 74).
	//
	// [server release lifecycle]: https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#upgrade-nag
	ServerNeedsUpgrade bool `json:"server_needs_upgrade,omitempty"`
	// The value of the `WEB_PUBLIC_STREAMS_ENABLED` Zulip server level setting. A server that has disabled this setting intends to not offer [web public channels] to realms it hosts. (Zulip Cloud defaults to `true`; self-hosted servers default to `false`).  Clients should use this to determine whether to offer UI for the realm-level setting for enabling web-public channels (`realm_enable_spectator_access`).
	//
	// **Changes**: New in Zulip 5.0 (feature level 110).
	//
	// [web public channels]: https://zulip.com/help/public-access-option
	ServerWebPublicStreamsEnabled bool `json:"server_web_public_streams_enabled,omitempty"`
	// The URL to a JSON file that describes which emoji names map to which emoji codes, for all Unicode emoji this Zulip server accepts.  The data at the given URL is a JSON object with one property, `code_to_names`. The value of that property is a JSON object where each key is an [emoji code].
	//
	// [emoji code]: https://zulip.com/api/add-reaction#parameter-emoji_code for an available Unicode emoji, and each value is the corresponding [emoji names] for this emoji, with the canonical name for the emoji always appearing first.  The HTTP response at that URL will have appropriate HTTP caching headers, such any HTTP implementation should get a cached version if emoji haven't changed since the last request.
	//
	// **Changes**: New in Zulip 6.0 (feature level 140
	//
	// [emoji names]: https://zulip.com/api/add-reaction#parameter-emoji_name
	ServerEmojiDataUrl string `json:"server_emoji_data_url,omitempty"`
	// The URL of the Jitsi server that the Zulip server is configured to use by default; the organization-level setting `realm_jitsi_server_url` takes precedence over this setting when both are set.
	//
	// **Changes**: New in Zulip 8.0 (feature level 212). Previously, this value was available as the now-deprecated `jitsi_server_url`.
	ServerJitsiServerUrl *string `json:"server_jitsi_server_url,omitempty"`
	// Present if `realm` is present in `fetch_event_types`  Whether topic summarization is enabled in the server or not depending upon whether `TOPIC_SUMMARIZATION_MODEL` is set or not.
	//
	// **Changes**: New in Zulip 10.0 (feature level 350).
	ServerCanSummarizeTopics bool `json:"server_can_summarize_topics,omitempty"`
	// Recommended client-side HTTP request timeout for [`GET /events`] calls. This is guaranteed to be somewhat greater than the heartbeat frequency. It is important that clients respect this parameter, so that increases in the heartbeat frequency do not break clients.
	//
	// **Changes**: New in Zulip 5.0 (feature level 74). Previously, this was hardcoded to 90 seconds, and clients should use that as a fallback value when interacting with servers where this field is not present.
	//
	// [`GET /events`]: https://zulip.com/api/get-events
	EventQueueLongpollTimeoutSeconds int64        `json:"event_queue_longpoll_timeout_seconds,omitempty"`
	Billing                          RealmBilling `json:"realm_billing,omitempty"`
	// Clients declaring the `empty_topic_name` client capability should use the value of `realm_empty_topic_display_name` to determine how to display the empty string topic.  Clients not declaring the `empty_topic_name` client capability receive `realm_empty_topic_display_name` value as the topic name replacing empty string.
	//
	// **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic name.
	EmptyTopicDisplayName string `json:"realm_empty_topic_display_name,omitempty"`

	// RealmUserSettingsDefaults Present if `realm_user_settings_defaults` is present in `fetch_event_types`.  A dictionary containing the default values of settings for new users.
	//
	//
	// **Changes**: New in Zulip 5.0 (feature level 95).
	// Prior to Zulip 5.0 (feature level 84), this field was present in response if `realm_user` was present in `fetch_event_types`, not `update_display_settings`.
	// **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.
	// Deprecated
	//
	// [capabilities]: https://zulip.com/api/register-queue#parameter-client_capabilities
	// [sending on pressing Enter]: https://zulip.com/help/configure-send-message-keys
	UserSettingsDefaults *UserSettings `json:"realm_user_settings_defaults,omitempty"`
}

// RealmBilling Present if `realm_billing` is present in `fetch_event_types`.  A dictionary containing billing information of the organization.
//
//	**Changes**: New in Zulip 10.0 (feature level 363).
type RealmBilling struct {
	// Whether there is a pending sponsorship request for the organization. Note that this field will always be `false` if the user is not in `can_manage_billing_group`.
	//
	// **Changes**: New in Zulip 10.0 (feature level 363).
	HasPendingSponsorshipRequest bool `json:"has_pending_sponsorship_request,omitempty"`
}

// GiphyRatingOptionsValue `{rating_name}`: Dictionary containing the details of the rating with the name of the rating as the key.
type GiphyRatingOptionsValue struct {
	// The Id of the rating option.
	Id int64 `json:"id,omitempty"`
	// The description of the rating option.
	Name string `json:"name,omitempty"`
}

// RealmDefaultExternalAccounts `{site_name}`: Dictionary containing the details of the default external account provider with the name of the website as the key.
type RealmDefaultExternalAccounts struct {
	// The name of the external account provider
	Name string `json:"name,omitempty"`
	// The text describing the external account.
	Text string `json:"text,omitempty"`
	// The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields for this account.
	Hint string `json:"hint,omitempty"`
	// The regex pattern of the URL of a profile page on the external site.
	UrlPattern string `json:"url_pattern,omitempty"`
}

type RealmConfiguration struct {
	RealmIdentity
	RealmPresentation
	RealmLocalization
	RealmPermissions
	SpecialChannels

	// Dictionary of authentication method keys mapped to dictionaries that describe the properties of the named authentication method for the organization - its enabled status and availability for use by the organization.  Clients should use this to implement server-settings UI to change which methods are enabled for the organization. For authentication UI itself, clients should use the pre-authentication metadata returned by [`GET /server_settings`].
	//
	// **Changes**: In Zulip 9.0 (feature level 243), the values in this dictionary were changed. Previously, the values were a simple boolean indicating whether the backend is enabled or not.
	//
	// [`GET /server_settings`]: https://zulip.com/api/get-server-settings
	AuthenticationMethods map[string]RealmAuthenticationMethod `json:"authentication_methods,omitempty"`

	// Whether the organization has enabled [weekly digest emails].
	//
	// [weekly digest emails]: https://zulip.com/help/digest-emails
	DigestEmailsEnabled bool `json:"digest_emails_enabled,omitempty"`
	// The day of the week when the organization will send its weekly digest email to inactive users.
	DigestWeekday int `json:"digest_weekday,omitempty"`
	// Whether read receipts is enabled in the organization or not.  If disabled, read receipt data will be unavailable to clients, regardless of individual users' personal read receipt settings. See also the `send_read_receipts` setting within `realm_user_settings_defaults`.
	//
	// **Changes**: New in Zulip 6.0 (feature level 137).
	EnableReadReceipts bool `json:"enable_read_receipts,omitempty"`
	// Whether [new users joining] this organization are required to have an email address in one of the `realm_domains` configured for the organization.
	//
	// [new users joining]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
	EmailsRestrictedToDomains bool `json:"emails_restricted_to_domains,omitempty"`
	// Whether clients should show a warning when a user is composing a DM to a guest user in this organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 348).
	EnableGuestUserDmWarning bool `json:"enable_guest_user_dm_warning,omitempty"`
	// Whether clients should display "(guest)" after the names of guest users to prominently highlight their status.
	//
	// **Changes**: New in Zulip 8.0 (feature level 216).
	EnableGuestUserIndicator bool `json:"enable_guest_user_indicator,omitempty"`
	// Whether web-public channels are enabled in this organization.  Can only be enabled if the `WEB_PUBLIC_STREAMS_ENABLED` [server setting] is enabled on the Zulip server. See also the `can_create_web_public_channel_group` realm setting.
	//
	// **Changes**: New in Zulip 5.0 (feature level 109).
	//
	// [server setting]: https://zulip.readthedocs.io/en/stable/production/settings.html
	EnableSpectatorAccess bool `json:"enable_spectator_access,omitempty"`
	// Whether an invitation is required to join this organization.
	InviteRequired bool `json:"invite_required,omitempty"`
	// The URL of the custom Jitsi Meet server configured in this organization's settings.  `null`, the default, means that the organization is using the should use the server-level configuration, `server_jitsi_server_url`.
	//
	// **Changes**: New in Zulip 8.0 (feature level 212). Previously, this was only available as a server-level configuration, and required a server restart to change.
	JitsiServerUrl *string `json:"jitsi_server_url,omitempty"`

	// The organization's default policy for sending channel messages to the [empty "general chat" topic].  - `"allow_empty_topic"`: Channel messages can be sent to the empty topic. - `"disable_empty_topic"`: Channel messages cannot be sent to the empty topic.
	//
	// **Changes**: New in Zulip 11.0 (feature level 392). Previously, this was controlled by the boolean realm `mandatory_topics` setting, which is now deprecated.
	//
	// [empty "general chat" topic]: https://zulip.com/help/general-chat-topic
	TopicsPolicy TopicsPolicy `json:"topics_policy,omitempty"`
	// Whether [topics are required] for messages in this organization.
	//
	// **Changes**: Deprecated in Zulip 11.0 (feature level 392). This is now controlled by the realm `topics_policy` setting.
	// Deprecated
	//
	// [topics are required]: https://zulip.com/help/require-topics
	MandatoryTopics bool `json:"mandatory_topics,omitempty"`
	// The new maximum file size that can be uploaded to this Zulip organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 306). Previously, this field of the core state did not support being updated via the events system, as it was typically hardcoded for a given Zulip installation.
	MaxFileUploadSizeMib int64 `json:"max_file_upload_size_mib,omitempty"`
	// Whether notification emails in this organization are allowed to contain Zulip the message content, or simply indicate that a new message was sent.
	MessageContentAllowedInEmailNotifications bool `json:"message_content_allowed_in_email_notifications,omitempty"`
	// Messages sent more than this many seconds ago cannot be deleted with this organization's [message deletion policy].  Will not be 0. A `null` value means no limit: messages can be deleted regardless of how long ago they were sent.
	//
	// **Changes**: No limit was represented using the special value `0` before Zulip 5.0 (feature level 100).
	//
	// [message deletion policy]: https://zulip.com/help/restrict-message-editing-and-deletion
	MessageContentDeleteLimitSeconds *int64 `json:"message_content_delete_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be edited with this organization's [message edit policy].  Will not be `0`. A `null` value means no limit, so messages can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`] for details and history of how message editing permissions work.
	//
	// **Changes**: Before Zulip 6.0 (feature level 138), no limit was represented using the special value `0`.
	//
	// [message edit policy]: https://zulip.com/help/restrict-message-editing-and-deletion
	// [`PATCH /messages/{message_id}`]: https://zulip.com/api/update-message
	MessageContentEditLimitSeconds *int64 `json:"message_content_edit_limit_seconds,omitempty"`
	// Which type of message edit history is configured to allow users to access [message edit history].
	//  - "all" = All edit history is visible.
	//  - "moves" = Only moves are visible.
	//  - "none" = No edit history is visible.
	//
	// **Changes**: New in Zulip 10.0 (feature level 358), replacing the previous `allow_edit_history` boolean setting; `true` corresponds to `all`, and `false` to `none`.
	//
	// [message edit history]: https://zulip.com/help/view-a-messages-edit-history
	MessageEditHistoryVisibilityPolicy string `json:"message_edit_history_visibility_policy,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved within a channel to another topic by users who have permission to do so based on this organization's [topic edit policy]. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so message topics can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`] for details and history of how message editing permissions work.
	//
	// **Changes**: New in Zulip 7.0 (feature level 162). Previously, this time limit was always 72 hours for users who were not administrators or moderators.
	//
	// [topic edit policy]: https://zulip.com/help/restrict-moving-messages
	// [`PATCH /messages/{message_id}`]: https://zulip.com/api/update-message
	MoveMessagesWithinStreamLimitSeconds *int64 `json:"move_messages_within_stream_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved between channels by users who have permission to do so based on this organization's [message move policy]. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so messages can be moved regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`] for details and history of how message editing permissions work.
	//
	// **Changes**: New in Zulip 7.0 (feature level 162). Previously, there was no time limit for moving messages between channels for users with permission to do so.
	//
	// [message move policy]: https://zulip.com/help/restrict-moving-messages
	// [`PATCH /messages/{message_id}`]: https://zulip.com/api/update-message
	MoveMessagesBetweenStreamsLimitSeconds *int64 `json:"move_messages_between_streams_limit_seconds,omitempty"`

	// Whether online presence of other users is shown in this organization.
	PresenceDisabled bool `json:"presence_disabled,omitempty"`
	// Whether push notifications are enabled for this organization. Typically `true` for Zulip Cloud and self-hosted realms that have a valid registration for the [Mobile push notifications service], and `false` for self-hosted servers that do not.
	//
	// **Changes**: New in Zulip 8.0 (feature level 231). Previously, this value was never updated via events.
	//
	// [Mobile push notifications service]: https://zulip.readthedocs.io/en/latest/production/mobile-push-notifications.html
	PushNotificationsEnabled bool `json:"push_notifications_enabled,omitempty"`
	// If the server expects the realm's push notifications access to end at a definite time in the future, the time at which this is expected to happen. Mobile clients should use this field to display warnings to users when the indicated timestamp is near.
	//
	// **Changes**: New in Zulip 8.0 (feature level 231).
	PushNotificationsEnabledEndTimestamp *time.Time `json:"push_notifications_enabled_end_timestamp,omitempty"`
	// Whether this realm is configured to disallow sending mobile push notifications with message content through the legacy mobile push notifications APIs. The new API uses end-to-end encryption to protect message content and metadata from being accessible to the push bouncer service, APNs, and FCM. Clients that support the new E2EE API will use it automatically regardless of this setting.  If `true`, mobile push notifications sent to clients that lack support for E2EE push notifications will always have "New message" as their content. Note that these legacy mobile notifications will still contain metadata, which may include the message's Id, the sender's name, email address, and avatar.  In a future release, once the official mobile apps have implemented fully validated their E2EE protocol support, this setting will become strict, and disable the legacy protocol entirely.
	//
	// **Changes**: New in Zulip 11.0 (feature level 409). Previously, this behavior was available only via the `PUSH_NOTIFICATION_REDACT_CONTENT` global server setting.
	RequireE2eePushNotifications bool `json:"require_e2ee_push_notifications,omitempty"`
	// Indicates whether the organization is configured to require users to have unique full names. If true, the server will reject attempts to create a new user, or change the name of an existing user, where doing so would lead to two users whose names are identical modulo case and unicode normalization.
	//
	// **Changes**: New in Zulip 9.0 (feature level 246). Previously, the Zulip server could not be configured to enforce unique names.
	RequireUniqueNames bool `json:"require_unique_names,omitempty"`
	// Whether or not this organization is configured to send the standard Zulip [welcome emails] to new users joining the organization.
	//
	// [welcome emails]: https://zulip.com/help/disable-welcome-emails
	SendWelcomeEmails bool `json:"send_welcome_emails,omitempty"`
	// The new upload quota for the Zulip organization.  If `null`, there is no limit.
	//
	// **Changes**: New in Zulip 10.0 (feature level 306). Previously, this was present changed via an `upload_quota` field in `extra_data` property of [realm/update] event format for `plan_type` events.
	//
	// [realm/update]: https://zulip.com/api/get-events#realm-update
	UploadQuotaMib *int64 `json:"upload_quota_mib,omitempty"`
	// The configured [video call provider] for the organization.
	//   - 0 = None
	//   - 1 = Jitsi Meet
	//   - 3 = Zoom (User OAuth integration)
	//   - 4 = BigBlueButton
	//   - 5 = Zoom (Server to Server OAuth integration)
	//
	// Note that only one of the [Zoom integrations] can be configured on a Zulip server.
	//
	// **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.
	// In Zulip 3.0 (feature level 1), added the None option to disable video call UI.
	// [Zoom integrations]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom
	// [video call provider]: https://zulip.com/help/configure-call-provider
	VideoChatProvider VideoChatProvider `json:"video_chat_provider,omitempty"`
	// Members whose accounts have been created at least this many days ago will be treated as [full members] for the purpose of settings that restrict access to new members.
	//
	// [full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	WaitingPeriodThreshold int64 `json:"waiting_period_threshold,omitempty"`
	// Whether the organization has given permission to be advertised in the Zulip [communities directory].
	//
	// **Changes**: New in Zulip 6.0 (feature level 129).
	//
	// [communities directory]: https://zulip.com/help/communities-directory
	WantAdvertiseInCommunitiesDirectory bool `json:"want_advertise_in_communities_directory,omitempty"`
}

type RealmIdentity struct {
	// The name of the organization, used in login pages etc.
	Name string `json:"name,omitempty"`
	// The [organization type] for the realm.
	//   - OrgTypeUnspecified
	//   - OrgTypeBusiness
	//   - OrgTypeOpenSource
	//   - OrgTypeEducationNonProfit
	//   - OrgTypeEducationForProfit
	//   - OrgTypeResearch
	//   - OrgTypeEventOrConference
	//   - OrgTypeNonProfitRegistered
	//   - OrgTypeGovernment
	//   - OrgTypePoliticalGroup
	//   - OrgTypeCommunity
	//   - OrgTypePersonal
	//   - OrgTypeOther
	//
	// **Changes**: New in Zulip 6.0 (feature level 128).
	//
	// [organization type]: https://zulip.com/help/organization-type
	OrgType OrgType `json:"org_type,omitempty"`
	// The plan type of the organization.
	//   - PlanTypeSelfHosted
	//   - PlanTypeLimited
	//   - PlanTypeStandard
	//   - PlanTypeStandardFree
	PlanType PlanType `json:"plan_type,omitempty"`
}

type RealmPresentation struct {
	// The description of the organization, used on login and registration pages.
	Description string `json:"description,omitempty"`
	// This organization's configured custom message for Welcome Bot to send to new user accounts, in Zulip Markdown format.  Maximum length is 8000 characters.
	//
	// **Changes**: New in Zulip 11.0 (feature level 416).
	WelcomeMessageCustomText string `json:"welcome_message_custom_text,omitempty"`
	// String indicating whether the organization's [profile wide logo] was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - "D" means the logo is the default Zulip logo. - "U" means uploaded by an organization administrator.
	//
	// [profile wide logo]: https://zulip.com/help/create-your-organization-profile
	LogoSource string `json:"logo_source,omitempty"`
	// The URL of the organization's wide logo configured in the [organization profile].
	//
	// [organization profile]: https://zulip.com/help/create-your-organization-profile
	LogoUrl string `json:"logo_url,omitempty"`
	// String indicating whether the organization's dark theme [profile wide logo] was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - "D" means the logo is the default Zulip logo. - "U" means uploaded by an organization administrator.
	//
	// [profile wide logo]: https://zulip.com/help/create-your-organization-profile
	NightLogoSource string `json:"night_logo_source,omitempty"`
	// The URL of the organization's dark theme wide-format logo configured in the [organization profile].
	//
	// [organization profile]: https://zulip.com/help/create-your-organization-profile
	NightLogoUrl string `json:"night_logo_url,omitempty"`
	// Maximum rating of the GIFs that will be retrieved from GIPHY.
	//
	// **Changes**: New in Zulip 4.0 (feature level 55).
	GiphyRating int32 `json:"giphy_rating,omitempty"`
	// String indicating whether the organization's [profile icon] was uploaded by a user or is the default. Useful for UI allowing editing the organization's icon.  - "G" means generated by Gravatar (the default). - "U" means uploaded by an organization administrator.
	//
	// [profile icon]: https://zulip.com/help/create-your-organization-profile
	IconSource string `json:"icon_source,omitempty"`
	// The URL of the organization's [profile icon].
	//
	// [profile icon]: https://zulip.com/help/create-your-organization-profile
	IconUrl string `json:"icon_url,omitempty"`
	// Whether this organization has been configured to enable [previews of linked images].
	//
	// [previews of linked images]: https://zulip.com/help/image-video-and-website-previews
	InlineImagePreview bool `json:"inline_image_preview,omitempty"`
	// Whether this organization has been configured to enable [previews of linked websites].
	//
	// [previews of linked websites]: https://zulip.com/help/image-video-and-website-previews
	InlineUrlEmbedPreview bool `json:"inline_url_embed_preview,omitempty"`
}

type RealmLocalization struct {
	// The default pygments language code to be used for code blocks in this organization. If an empty string, no default has been set.
	//
	// **Changes**: Prior to Zulip 8.0 (feature level 195), a server bug meant that both `null` and an empty string could represent that no default was set for this realm setting in the [`POST /register`] response. The documentation for both that endpoint and this event incorrectly stated that the only representation for no default language was `null`. This event in fact uses the empty string to indicate that no default has been set in all server versions.
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	DefaultCodeBlockLanguage string `json:"default_code_block_language,omitempty"`
	// The default language for the organization.
	DefaultLanguage string `json:"default_language,omitempty"`
}

type SpecialChannels struct {
	// The Id of the private channel to which messages flagged by users for moderation are sent. Moderators can use this channel to review and act on reported content.  Will be `-1` if moderation requests are disabled.  Clients should check whether moderation requests are disabled to determine whether to present a "report message" feature in their UI within a given organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 331). Previously, no "report message" feature existed in Zulip.
	ModerationRequestChannelId int64 `json:"realm_moderation_request_channel_id,omitempty"`
	// The Id of the channel to which automated messages announcing the [creation of new channels] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.
	//
	// **Changes**: In Zulip 9.0 (feature level 241), renamed 'realm_notifications_stream_id' to `realm_new_stream_announcements_stream_id`.
	//
	// [creation of new channels]: https://zulip.com/help/configure-automated-notices#new-channel-announcements
	NewStreamAnnouncementsChannelId int64 `json:"realm_new_stream_announcements_stream_id,omitempty"`
	// The Id of the channel to which automated messages announcing that [new users have joined the organization] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.
	//
	// **Changes**: In Zulip 9.0 (feature level 241), renamed 'realm_signup_notifications_stream_id' to `realm_signup_announcements_stream_id`.
	//
	// [new users have joined the organization]: https://zulip.com/help/configure-automated-notices#new-user-announcements
	SignupAnnouncementsChannelId int64 `json:"realm_signup_announcements_stream_id,omitempty"`
	// The Id of the channel to which automated messages announcing new features or other end-user updates about the Zulip software are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.
	//
	// **Changes**: New in Zulip 9.0 (feature level 242).
	ZulipUpdateAnnouncementsChannelId int64 `json:"realm_zulip_update_announcements_stream_id,omitempty"`
}

type RealmPermissions struct {
	// Indicates whether users are [allowed to change] their name via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.
	//
	// [allowed to change]: https://zulip.com/help/restrict-name-and-email-changes
	NameChangesDisabled bool `json:"name_changes_disabled,omitempty"`
	// Whether the organization disallows disposable email addresses.
	DisallowDisposableEmailAddresses bool `json:"realm_disallow_disposable_email_addresses,omitempty"`
	// Whether users are allowed to change their own email address in this organization. This is typically disabled for organizations that synchronize accounts from LDAP or a similar corporate database.
	EmailChangesDisabled bool `json:"email_changes_disabled,omitempty"`
	// Whether this organization's [message edit policy] allows editing the content of messages.  See [`PATCH /messages/{message_id}`] for details and history of how message editing permissions work.
	// [message edit policy]: https://zulip.com/help/restrict-message-editing-and-deletion
	// [`PATCH /messages/{message_id}`]: https://zulip.com/api/update-message
	AllowMessageEditing *bool `json:"realm_allow_message_editing,omitempty"`
	// A [group-setting value] defining the set of users who have permission to add custom emoji in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 307). Previously, this permission was controlled by the enum `add_custom_emoji_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators.  Before Zulip 5.0 (feature level 85), the `realm_add_emoji_by_admins_only` boolean setting controlled this permission; `true` corresponded to `Admins`, and `false` to `Everyone`.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanAddCustomEmojiGroup GroupSettingValue `json:"realm_can_add_custom_emoji_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to add subscribers to channels in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 341). Previously, this permission was controlled by the enum `invite_to_stream_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanAddSubscribersGroup GroupSettingValue `json:"realm_can_add_subscribers_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to delete any message in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 281). Previously, this permission was limited to administrators only and was uneditable.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanDeleteAnyMessageGroup GroupSettingValue `json:"realm_can_delete_any_message_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to delete messages that they have sent in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 291). Previously, this permission was controlled by the enum `delete_own_message_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 5=Everyone.  Before Zulip 5.0 (feature level 101), the `allow_message_deleting` boolean setting controlled this permission; `true` corresponded to `Everyone`, and `false` to `Admins`.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanDeleteOwnMessageGroup GroupSettingValue `json:"realm_can_delete_own_message_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to change per-channel `can_delete_any_message_group` and `can_delete_own_message_group` permission settings. Note that the user must be a member of both this group and the `can_administer_channel_group` of the channel whose message delete settings they want to change.  Organization administrators can always change these settings of every channel.
	//
	// **Changes**: New in Zulip 11.0 (feature level 407).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanSetDeleteMessagePolicyGroup GroupSettingValue `json:"realm_can_set_delete_message_policy_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to change per-channel `topics_policy` setting. Note that the user must be a member of both this group and the `can_administer_channel_group` of the channel whose `topics_policy` they want to change.  Organization administrators can always change the `topics_policy` setting of every channel.
	//
	// **Changes**: New in Zulip 11.0 (feature level 392).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanSetTopicsPolicyGroup GroupSettingValue `json:"realm_can_set_topics_policy_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to send email invitations for inviting other users to the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 321). Previously, this permission was controlled by the enum `invite_to_realm_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 6=Nobody.  Before Zulip 4.0 (feature level 50), the `invite_by_admins_only` boolean setting controlled this permission; `true` corresponded to `Admins`, and `false` to `Members`.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanInviteUsersGroup GroupSettingValue `json:"realm_can_invite_users_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.
	//
	// **Changes**: New in Zulip 10.0 (feature level 352). Previously, this permission was controlled by the enum `wildcard_mention_policy`.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanMentionManyUsersGroup GroupSettingValue `json:"realm_can_mention_many_users_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to move messages from one channel to another in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 310). Previously, this permission was controlled by the enum `move_messages_between_streams_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 6=Nobody.  In Zulip 7.0 (feature level 159), `Nobody` was added as an option to `move_messages_between_streams_policy` enum.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanMoveMessagesBetweenChannelsGroup GroupSettingValue `json:"realm_can_move_messages_between_channels_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to move messages from one topic to another within a channel in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 316). Previously, this permission was controlled by the enum `edit_topic_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 5=Everyone, 6=Nobody.  In Zulip 7.0 (feature level 159), `Nobody` was added as an option to `edit_topic_policy` enum.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanMoveMessagesBetweenTopicsGroup GroupSettingValue `json:"realm_can_move_messages_between_topics_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create user groups in this organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 299). Previously `realm_user_group_edit_policy` field used to control the permission to create user groups.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreateGroups GroupSettingValue `json:"realm_can_create_groups,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create all types of bot users in the organization. See also `can_create_write_only_bots_group`.
	//
	// **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum `bot_creation_policy`. Values were 1=Members, 2=Generic bots limited to administrators, 3=Administrators.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreateBotsGroup GroupSettingValue `json:"realm_can_create_bots_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create bot users that can only send messages in the organization, i.e. incoming webhooks, in addition to the users who are present in `can_create_bots_group`.
	//
	// **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum `bot_creation_policy`. Values were 1=Members, 2=Generic bots limited to administrators, 3=Administrators.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreateWriteOnlyBotsGroup GroupSettingValue `json:"realm_can_create_write_only_bots_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to administer all existing groups in this organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 305), only users who were a member of the group or had the moderator role or above could exercise the permission on a given group.  New in Zulip 10.0 (feature level 299). Previously the `user_group_edit_policy` field controlled the permission to manage user groups. Valid values were as follows:
	//   - 1 = All members can create and edit user groups
	//   - 2 = Only organization administrators can create and edit   user groups
	//   - 3 = Only [full members] can create and   edit user groups.
	//   - 4 = Only organization administrators and moderators can   create and edit user groups.
	//
	// [full members]: https://zulip.com/api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanManageAllGroups GroupSettingValue `json:"realm_can_manage_all_groups,omitempty"`
	// A [group-setting value] defining the set of users who have permission to manage plans and billing in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 363). Previously, only owners and users with `is_billing_admin` property set to `true` were allowed to manage plans and billing.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanManageBillingGroup GroupSettingValue `json:"realm_can_manage_billing_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create public channels in this organization.
	//
	// **Changes**: New in Zulip 9.0 (feature level 264). Previously `realm_create_public_stream_policy` field used to control the permission to create public channels.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreatePublicChannelGroup GroupSettingValue `json:"realm_can_create_public_channel_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create private channels in this organization.
	//
	// **Changes**: New in Zulip 9.0 (feature level 266). Previously `realm_create_private_stream_policy` field used to control the permission to create private channels.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreatePrivateChannelGroup GroupSettingValue `json:"realm_can_create_private_channel_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to create web-public channels in this organization.  Has no effect and should not be displayed in settings UI unless the Zulip server has the `WEB_PUBLIC_STREAMS_ENABLED` server-level setting enabled and the organization has enabled the `enable_spectator_access` realm setting.
	//
	// **Changes**: New in Zulip 10.0 (feature level 280). Previously `realm_create_web_public_stream_policy` field used to control the permission to create web-public channels.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanCreateWebPublicChannelGroup GroupSettingValue `json:"realm_can_create_web_public_channel_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to [resolve topics] in the organization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 367). Previously, permission to resolve topics was controlled by the more general `can_move_messages_between_topics_group permission for moving messages`.
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	// [resolve topics]: https://zulip.com/help/resolve-a-topic
	CanResolveTopicsGroup GroupSettingValue `json:"realm_can_resolve_topics_group,omitempty"`
	// A [group-setting value] defining the set of users who are allowed to create [reusable invitation links] to the organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 209).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	// [reusable invitation links]: https://zulip.com/help/invite-new-users#create-a-reusable-invitation-link
	CreateMultiuseInviteGroup GroupSettingValue `json:"realm_create_multiuse_invite_group,omitempty"`
	// A [group-setting value] defining the set of users who are allowed to access all users in the organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 225).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanAccessAllUsersGroup GroupSettingValue `json:"realm_can_access_all_users_group,omitempty"`
	// A [group-setting value] defining the set of users who are allowed to use AI summarization.
	//
	// **Changes**: New in Zulip 10.0 (feature level 350).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	CanSummarizeTopicsGroup GroupSettingValue `json:"realm_can_summarize_topics_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to start a new direct message conversation involving other non-bot users. Users who are outside this group and attempt to send the first direct message to a given collection of recipient users will receive an error, unless all other recipients are bots or the sender.
	//
	// **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the `private_message_policy` realm setting, which supported values of 1 (enabled) and 2 (disabled).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	DirectMessageInitiatorGroup GroupSettingValue `json:"realm_direct_message_initiator_group,omitempty"`
	// A [group-setting value] defining the set of users who have permission to fully use direct messages. Users outside this group can only send direct messages to conversations where all the recipients are in this group, are bots, or are the sender, ensuring that every direct message conversation will be visible to at least one user in this group.
	//
	// **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the `private_message_policy` realm setting, which supported values of 1 (enabled) and 2 (disabled).
	//
	// [group-setting value]: https://zulip.com/api/group-setting-values
	DirectMessagePermissionGroup GroupSettingValue `json:"realm_direct_message_permission_group,omitempty"`
}
