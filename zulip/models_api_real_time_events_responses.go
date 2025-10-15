package zulip

import "time"

// GetEventsResponse struct for GetEventsResponse
type GetEventsResponse struct {
	Response

	// An array of `event` objects (possibly zero-length if `dont_block` is set) with Ids newer than `last_event_id`. Event Ids are guaranteed to be increasing, but they are not guaranteed to be consecutive.
	Events []EventEnvelope `json:"events,omitempty"`
	// The Id of the registered queue.
	QueueId string `json:"queue_id,omitempty"`
}

// RegisterQueueResponse struct for RegisterQueueResponse
type RegisterQueueResponse struct {
	Response

	// The Id of the queue that has been allocated for your client.  Will be `null` only for unauthenticated access in realms that have enabled the [public access option](zulip.com/help/public-access-option.
	QueueId *string `json:"queue_id,omitempty"`
	// The initial value of `last_event_id` to pass to `GET /api/v1/events`.
	LastEventId int64 `json:"last_event_id,omitempty"`
	// The server's current [Zulip feature level](zulip.com/api/changelog.  **Changes**: As of Zulip 3.0 (feature level 3), this is always present in the endpoint's response. Previously, it was only present if `event_types` included `zulip_version`.  New in Zulip 3.0 (feature level 1).
	ZulipFeatureLevel int64 `json:"zulip_feature_level,omitempty"`
	// The server's version number. This is often a release version number, like `2.1.7`. But for a server running a [version from Git][git-release], it will be a Git reference to the commit, like `5.0-dev-1650-gc3fd37755f`.  **Changes**: As of Zulip 3.0 (feature level 3), this is always present in the endpoint's response. Previously, it was only present if `event_types` included `zulip_version`.  [git-release]: https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#git-versions
	ZulipVersion string `json:"zulip_version,omitempty"`
	// The `git merge-base` between `zulip_version` and official branches in the public [Zulip server and web app repository](https://github.com/zulip/zulip), in the same format as `zulip_version`. This will equal `zulip_version` if the server is not running a fork of the Zulip server.  This will be `\"\"` if the server does not know its `merge-base`.  **Changes**: New in Zulip 5.0 (feature level 88).
	ZulipMergeBase string `json:"zulip_merge_base,omitempty"`
	// Present if `alert_words` is present in `fetch_event_types`.  An array of strings, each an [alert word](zulip.com/help/dm-mention-alert-notifications#alert-words that the current user has configured.
	AlertWords []string `json:"alert_words,omitempty"`
	// Present if `custom_profile_fields` is present in `fetch_event_types`.  An array of dictionaries where each dictionary contains the details of a single custom profile field that is available to users in this Zulip organization. This must be combined with the custom profile field values on individual user objects to display users' profiles.
	CustomProfileFields []CustomProfileField `json:"custom_profile_fields,omitempty"`
	// Present if `custom_profile_fields` is present in `fetch_event_types`.  An array of objects; each object describes a type of custom profile field that could be configured on this Zulip server. Each custom profile type has an Id and the `type` property of a custom profile field is equal to one of these Ids.  This attribute is only useful for clients containing UI for changing the set of configured custom profile fields in a Zulip organization.
	CustomProfileFieldTypes map[string]CustomProfileFieldType `json:"custom_profile_field_types,omitempty"`

	// Present if `realm` is present in `fetch_event_types`, and the realm is a demo organization.  The UNIX timestamp (UTC) when the demo organization will be automatically deleted. Clients should use this to display a prominent warning to the user that the organization will be deleted at the indicated time.  **Changes**: New in Zulip 5.0 (feature level 94).
	DemoOrganizationScheduledDeletionDate *time.Time `json:"demo_organization_scheduled_deletion_date,omitempty"`
	// An array containing draft objects for the user. These drafts are being stored on the backend for the purpose of syncing across devices. This array will be empty if `enable_drafts_synchronization` is set to `false`.
	Drafts []Draft `json:"drafts,omitempty"`
	// Present if `onboarding_steps` is present in `fetch_event_types`.  An array of dictionaries, where each dictionary contains details about a single onboarding step that should be shown to the user.  We expect that only official Zulip clients will interact with this data.  **Changes**: Before Zulip 8.0 (feature level 233), this array was named `hotspots`. Prior to this feature level, one-time notice onboarding steps were not supported, and the `type` field in these objects did not exist as all onboarding steps were implicitly hotspots.
	OnboardingSteps []OnboardingStep `json:"onboarding_steps,omitempty"`
	// Present if `onboarding_steps` is present in `fetch_event_types`.  URL of the navigation tour video to display to new users during onboarding. If `null`, the onboarding video experience is disabled.  **Changes**: New in Zulip 10.0 (feature level 369).
	NavigationTourVideoUrl *string `json:"navigation_tour_video_url,omitempty"`
	// Present if `message` is present in `fetch_event_types`.  The highest message Id among all messages the user has received as of the moment of this request.  **Deprecated**: This field may be removed in future versions as it no longer has a clear purpose. Clients wishing to fetch the latest messages should pass `\"anchor\": \"latest\"` to `GET /messages`.
	// Deprecated
	MaxMessageId *int64 `json:"max_message_id,omitempty"`
	// The maximum allowed length for a reminder note.  **Changes**: New in Zulip 11.0 (feature level 415).
	MaxReminderNoteLength *int64 `json:"max_reminder_note_length,omitempty"`

	// Present if `scheduled_messages` is present in `fetch_event_types`.  An array of all undelivered scheduled messages by the user.  **Changes**: New in Zulip 7.0 (feature level 179).
	ScheduledMessages []ScheduledMessage `json:"scheduled_messages,omitempty"`
	// Present if `reminders` is present in `fetch_event_types`.  An array of all undelivered reminders scheduled by the user.  **Changes**: New in Zulip 11.0 (feature level 399).
	Reminders []ScheduledMessage `json:"reminders,omitempty"`
	// Present if `muted_topics` is present in `fetch_event_types`.  Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.  **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, `muted_topics` will only be present in the response if the `user_topic` object, which generalizes and replaces this field, is not explicitly requested via `fetch_event_types`.  Before Zulip 3.0 (feature level 1), the `muted_topics` array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.
	// Deprecated
	MutedTopics [][]UserSettingsUpdateEvent1MutedTopicsInnerInner `json:"muted_topics,omitempty"`
	// Present if `muted_users` is present in `fetch_event_types`.  A list of dictionaries where each dictionary describes a [muted user](zulip.com/api/mute-user.  **Changes**: New in Zulip 4.0 (feature level 48).
	MutedUsers []MutedUser `json:"muted_users,omitempty"`
	// Present if `presence` is present in `fetch_event_types`.  A dictionary where each entry describes the presence details of a user in the Zulip organization.  The format of the entry (modern or legacy) depends on the value of [`slim_presence`](#parameter-slim_presence).  Users who have been offline for multiple weeks may not appear in this object.
	Presences map[string]PresenceUpdateValue `json:"presences,omitempty"`
	// Present if `presence` is present in `fetch_event_types`.  Provides the `last_update_id` value of the latest presence data fetched by the server and included in the response in `presences`. This can be used as the value of the `presence_last_update_id` parameter when polling for presence data at the [/users/me/presence](zulip.com/api/update-presence) endpoint to tell the server to only fetch the relevant newer data in order to skip redundant already-known presence information.  **Changes**: New in Zulip 9.0 (feature level 263).
	PresenceLastUpdateId *int64 `json:"presence_last_update_id,omitempty"`
	// Present if `presence` is present in `fetch_event_types`.  The time when the server fetched the `presences` data included in the response. Matches the similar field in presence responses.  **Changes**: New in Zulip 5.0 (feature level 70).
	ServerTimestamp *float32 `json:"server_timestamp,omitempty"`
	// Present if `realm_domains` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes a domain within which users can join the organization without and invitation.
	RealmDomains []RealmDomain `json:"realm_domains,omitempty"`
	// Present if `realm_emoji` is present in `fetch_event_types`.  A dictionary of objects where each object describes a custom emoji that has been uploaded in this Zulip organization.
	RealmEmoji map[string]RealmEmoji `json:"realm_emoji,omitempty"`
	// Present if `realm_linkifiers` is present in `fetch_event_types`.  An ordered array of objects where each object describes a single [linkifier](zulip.com/help/add-a-custom-linkifier.  The order of the array reflects the order that each linkifier should be processed when linkifying messages and topics. By default, new linkifiers are ordered last. This order can be modified with [`PATCH /realm/linkifiers`](zulip.com/api/reorder-linkifiers.  Clients will receive an empty array unless the event queue is registered with the client capability `{\"linkifier_url_template\": true}`. See [`client_capabilities`](zulip.com/api/register-queue#parameter-client_capabilities parameter for how this can be specified.  **Changes**: Before Zulip 7.0 (feature level 176), the `linkifier_url_template` client capability was not required. The requirement was added because linkifiers were updated to contain a URL template instead of a URL format string, which was a not backwards-compatible change.  New in Zulip 4.0 (feature level 54). Clients can access this data for servers on earlier feature levels via the legacy `realm_filters` property.
	RealmLinkifiers []RealmLinkifiers `json:"realm_linkifiers,omitempty"`
	// Legacy property for [linkifiers](zulip.com/help/add-a-custom-linkifier. Present if `realm_filters` is present in `fetch_event_types`.  When present, this is always an empty array.  **Changes**: Prior to Zulip 7.0 (feature level 176), this was an array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example `\"#(?P<id>[123])\"`. The second element was a URL format string that the pattern should be linkified with. A URL format string for the above example would be `\"https://realm.com/my_realm_filter/%(id)s\"`. And the third element was the Id of the realm filter.  **Deprecated** in Zulip 4.0 (feature level 54), replaced by the `realm_linkifiers` key.
	// Deprecated
	RealmFilters [][]RealmFilterTuple `json:"realm_filters,omitempty"`
	// Present if `realm_playgrounds` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes a [code playground](zulip.com/help/code-blocks#code-playgrounds configured for this Zulip organization.  **Changes**: New in Zulip 4.0 (feature level 49).
	RealmPlaygrounds []RealmPlayground `json:"realm_playgrounds,omitempty"`
	// Present if `realm_user_groups` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes a [user group](zulip.com/help/user-groups) in the Zulip organization.  Deactivated groups will only be included if `include_deactivated_groups` client capability is set to `true`.  **Changes**: Prior to Zulip 10.0 (feature level 294), deactivated groups were included for all the clients.
	RealmUserGroups []UserGroup `json:"realm_user_groups,omitempty"`
	// Present if `realm_bot` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes a bot that the current user can administer. If the current user is an organization administrator, this will include all bots in the organization. Otherwise, it will only include bots owned by the user (either because the user created the bot or an administrator transferred the bot's ownership to the user).
	RealmBots []Bot `json:"realm_bots,omitempty"`
	// Present if `realm_embedded_bots` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes an type of embedded bot that is available to be configured on this Zulip server.  Clients only need these data if they contain UI for creating or administering bots.
	RealmEmbeddedBots []RealmEmbeddedBots `json:"realm_embedded_bots,omitempty"`
	// Present if `realm_incoming_webhook_bots` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes a type of incoming webhook integration that is available to be configured on this Zulip server.  Clients only need these data if they contain UI for creating or administering bots.
	RealmIncomingWebhookBots []RealmIncomingWebhookBot `json:"realm_incoming_webhook_bots,omitempty"`
	// Present if `recent_private_conversations` is present in `fetch_event_types`.  An array of dictionaries containing data on all direct message and group direct message conversations that the user has received (or sent) messages in, organized by conversation. This data set is designed to support UI elements such as the \"Direct messages\" widget in the web application showing recent direct message conversations that the user has participated in.  \"Recent\" is defined as the server's discretion; the original implementation interpreted that as \"the 1000 most recent direct messages the user received\".
	RecentPrivateConversations []RecentPrivateConversation `json:"recent_private_conversations,omitempty"`
	// Present if `navigation_views` is present in `fetch_event_types`. An array of dictionaries containing data on all of the current user's navigation views.  **Changes**: New in Zulip 11.0 (feature level 390).
	NavigationViews []NavigationView `json:"navigation_views,omitempty"`
	// Present if `saved_snippets` is present in `fetch_event_types`.  An array of dictionaries containing data on all of the current user's saved snippets.  **Changes**: New in Zulip 10.0 (feature level 297).
	SavedSnippets []SavedSnippet `json:"saved_snippets,omitempty"`
	// Present if `subscription` is present in `fetch_event_types`.  A array of dictionaries where each dictionary describes the properties of a channel the user is subscribed to (as well as that user's personal per-channel settings).  **Changes**: Removed `email_address` field from the dictionary in Zulip 8.0 (feature level 226).  Removed `role` field from the dictionary in Zulip 6.0 (feature level 133).
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	// Present if `subscription` is present in `fetch_event_types`.  A array of dictionaries where each dictionary describes one of the channels the user has unsubscribed from but was previously subscribed to along with the subscription details.  Unlike `never_subscribed`, the user might have messages in their personal message history that were sent to these channels.  **Changes**: Prior to Zulip 10.0 (feature level 349), if a user was in `can_administer_channel_group` of a channel that they had unsubscribed from, but not an organization administrator, the channel in question would not be part of this array.  Removed `email_address` field from the dictionary in Zulip 8.0 (feature level 226).  Removed `role` field from the dictionary in Zulip 6.0 (feature level 133).
	Unsubscribed []Subscription `json:"unsubscribed,omitempty"`
	// Present if `subscription` is present in `fetch_event_types`.  A array of dictionaries where each dictionary describes one of the channels that is visible to the user and the user has never been subscribed to.  Important for clients containing UI where one can browse channels to subscribe to.  **Changes**: Before Zulip 10.0 (feature level 362), archived channels did not appear in this list, even if the `archived_channels` [client capability][client-capabilities] was declared by the client.  Prior to Zulip 10.0 (feature level 349), if a user was in `can_administer_channel_group` of a channel that they never subscribed to, but not an organization administrator, the channel in question would not be part of this array.
	NeverSubscribed []ChannelWithSubscribers `json:"never_subscribed,omitempty"`
	// Present if `channel_folders` is present in `fetch_event_types`.  An array of dictionaries where each dictionary describes one of the channel folders in the organization.  Only channel folders with one or more public web channels are visible to spectators.  **Changes**: New in Zulip 11.0 (feature level 389).
	ChannelFolders []ChannelFolder `json:"channel_folders,omitempty"`
	UnreadMsgs     *UnreadMsgs     `json:"unread_msgs,omitempty"`
	// Present if `starred_messages` is present in `fetch_event_types`.  Array containing the Ids of all messages which have been [starred](zulip.com/help/star-a-message) by the user.
	StarredMessages []int64 `json:"starred_messages,omitempty"`
	// Present if `stream` is present in `fetch_event_types`.  Array of dictionaries where each dictionary contains details about a single channel in the organization that is visible to the user.  For organization administrators, this will include all private channels in the organization.  **Changes**: Before Zulip 11.0 (feature level 378), archived channels did not appear in this list, even if the `archived_channels` [client capability][client-capabilities] was declared by the client.  As of Zulip 8.0 (feature level 205), this will include all web-public channels in the organization as well.
	Streams []Channel `json:"streams,omitempty"`
	// Present if `default_streams` is present in `fetch_event_types`.  An array of Ids of all the [default channels](zulip.com/help/set-default-streams-for-new-users) in the organization.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.
	RealmDefaultStreams []int64 `json:"realm_default_streams,omitempty"`
	// Present if `default_stream_groups` is present in `fetch_event_types`.  An array of dictionaries where each dictionary contains details about a single default channel group configured for this Zulip organization.  Default channel groups are an experimental feature.
	RealmDefaultStreamGroups []DefaultChannelGroup `json:"realm_default_stream_groups,omitempty"`
	// Present if `stop_words` is present in `fetch_event_types`.  An array containing the stop words used by the Zulip server's full-text search implementation. Useful for showing helpful error messages when a search returns limited results because a stop word in the query was ignored.
	StopWords []string `json:"stop_words,omitempty"`
	// Present if `user_status` is present in `fetch_event_types`.  A dictionary which contains the [status](zulip.com/help/status-and-availability) of all users in the Zulip organization who have set a status.  **Changes**: The emoji parameters are new in Zulip 5.0 (feature level 86). Previously, Zulip did not support emoji associated with statuses.
	UserStatus   map[string]UserStatus `json:"user_status,omitempty"`
	UserSettings *UserSettings         `json:"user_settings,omitempty"`
	// Present if `user_topic` is present in `fetch_event_types`.  **Changes**: New in Zulip 6.0 (feature level 134), deprecating and replacing the previous `muted_topics` structure.
	UserTopics []UserTopics `json:"user_topics,omitempty"`
	// Present if `video_calls` is present in `fetch_event_types`.  A boolean which signifies whether the user has a Zoom token and has thus completed OAuth flow for the [Zoom integration](zulip.com/help/configure-call-provider. Clients need to know whether initiating Zoom OAuth is required before creating a Zoom call.
	HasZoomToken *bool `json:"has_zoom_token,omitempty"`
	// Present if `giphy` is present in `fetch_event_types`.  GIPHY's client-side SDKs needs this API key to use the GIPHY API. GIPHY API keys are not secret (their main purpose appears to be allowing GIPHY to block a problematic app). Please don't use our API key for an app unrelated to Zulip.  Developers of clients should also read the [GIPHY API TOS](https://support.giphy.com/hc/en-us/articles/360028134111-GIPHY-API-Terms-of-Service-) before using this API key.  **Changes**: Added in Zulip 4.0 (feature level 47).
	GiphyApiKey *string `json:"giphy_api_key,omitempty"`
	// Present if `push_device` is present in `fetch_event_types`.  Dictionary where each entry describes the user's push device's registration status and error code (if registration failed).  **Changes**: New in Zulip 11.0 (feature level 406).
	PushDevices map[string]PushDevicesValue `json:"push_devices,omitempty"`
	// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
	ReceivesTypingNotifications *bool `json:"receives_typing_notifications,omitempty"`

	ServerSupportedPermissionSettings ServerSupportedPermissionSettings `json:"server_supported_permission_settings,omitempty"`
	// Maximum number of new subscribers for which the server will respect the `send_new_subscription_messages` parameter when [adding subscribers to a channel](zulip.com/api/subscribe#parameter-send_new_subscription_messages.  **Changes**: New in Zulip 11.0 (feature level 397).
	MaxBulkNewSubscriptionMessages float32 `json:"max_bulk_new_subscription_messages,omitempty"`

	// Present if `update_display_settings` is present in `fetch_event_types` and only for clients that did not include `user_settings_object` in their [`client_capabilities`][capabilities] when registering the event queue.
	UpdateDisplaySettings *UpdateDisplaySettings `json:"update_display_settings,omitempty"`

	// Present if `realm_user` is present in `fetch_event_types`.
	RealmUser *RealmUser

	// Present if `update_global_notifications` is present in `fetch_event_types` and only for clients that did not include `user_settings_object` in their [`client_capabilities`][capabilities] when registering the event queue.
	GlobalNotifications *GlobalNotifications

	// Present if `realm` is present in `fetch_event_types`.
	Realm *Realm
}

// RecentPrivateConversation Object describing a single recent direct conversation in the user's history.
type RecentPrivateConversation struct {
	// The highest message Id of the conversation, intended to support sorting the conversations by recency.
	MaxMessageId int64 `json:"max_message_id,omitempty"`
	// The list of users other than the current user in the direct message conversation. This will be an empty list for direct messages sent to oneself.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// ServerSupportedPermissionSettings Present if `realm` is present in `fetch_event_types`.  Metadata detailing the valid values for permission settings that use [group-setting values](zulip.com/api/group-setting-values. Clients should use these data as explained in the [main documentation](zulip.com/api/group-setting-values#permitted-values to determine what values to present as possible values for these settings in UI components.  This part of the Zulip API is unstable and may change significantly in future versions.  **Changes**: New in Zulip 8.0 (feature level 221).
type ServerSupportedPermissionSettings struct {
	// Configuration for realm level group permission settings.
	Realm map[string]GroupPermissionSetting `json:"realm,omitempty"`
	// Configuration for channel level group permission settings.
	Stream map[string]GroupPermissionSetting `json:"stream,omitempty"`
	// Configuration for group level group permission settings.
	Group map[string]GroupPermissionSetting `json:"group,omitempty"`
}

// GroupPermissionSetting Configuration for a group permission setting specifying the groups to which the setting can be set to and the default values for the setting.  **Changes**: Removed `allow_owners_group` field in Zulip 10.0 (feature level 326), as we now support anonymous user groups. Previously, the `role:owners` system group was not offered when `allow_owners_group` was false.  Removed unnecessary `id_field_name` field in Zulip 10.0 (feature level 326). Previously, this always had the value of `\"{setting_name}_id\"`; it was an internal implementation detail of the server not intended to be included in the API.
type GroupPermissionSetting struct {
	// Whether the setting can only be set to a system user group.
	RequireSystemGroup bool `json:"require_system_group,omitempty"`
	// Whether the setting can be set to `role:internet` system group.
	AllowInternetGroup bool `json:"allow_internet_group,omitempty"`
	// Whether the setting can be set to `role:nobody` system group.
	AllowNobodyGroup bool `json:"allow_nobody_group,omitempty"`
	// Whether the setting can be set to `role:everyone` system group.  If false, guest users cannot exercise this permission even if they are part of the [group-setting value](zulip.com/api/group-setting-values) for this setting.
	AllowEveryoneGroup bool `json:"allow_everyone_group,omitempty"`
	// Name of the default group for the setting.
	DefaultGroupName string `json:"default_group_name,omitempty"`
	// Name of the default group for the setting for system groups.  This is non-null only for group-level settings.
	DefaultForSystemGroups *string `json:"default_for_system_groups,omitempty"`
	// An array of names of system groups to which the setting can be set to.  If the list is empty, the setting can be set to system groups based on the other boolean fields.  **Changes**: New in Zulip 8.0 (feature level 225).
	AllowedSystemGroups []string `json:"allowed_system_groups,omitempty"`
}

// PushDevicesValue `{push_account_id}`: Dictionary containing the details of a push device with the push account Id as the key.
type PushDevicesValue struct {
	// The push account's registration status. Either `\"active\"`, `\"pending\"`, or `\"failed\"`.
	Status string `json:"status,omitempty"`
	// If the status is `\"failed\"`, a [Zulip API error code](zulip.com/api/rest-error-handling) indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - `\"INVALId_BOUNCER_PUBLIC_KEY\"` - Inform the user to update app. - `\"REQUEST_EXPIRED` - Retry with a fresh payload.
	ErrorCode *string `json:"error_code,omitempty"`
}

// RealmEmbeddedBots Object containing details of an embedded bot. Embedded bots are an experimental feature not enabled in production yet.
type RealmEmbeddedBots struct {
	// The name of the bot.
	Name string `json:"name,omitempty"`
	// A dictionary of string key/value pairs, which describe the configuration for the bot. These are usually details like API keys, and are unique to the integration/bot. Can be an empty dictionary.
	Config map[string]string `json:"config,omitempty"`
}

type UpdateDisplaySettings struct {
	// The color scheme selected by the user.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	ColorScheme ColorScheme `json:"color_scheme,omitempty"`
	// The default language chosen by the user.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	DefaultLanguage string `json:"default_language,omitempty"`
	// Whether the user has chosen to hide inactive channels.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	DemoteInactiveStreams int64 `json:"demote_inactive_streams,omitempty"`
	// The name of the emoji set that the user has chosen.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	Emojiset string `json:"emojiset,omitempty"`
	// Whether drafts synchronization is enabled for the user. If disabled, clients will receive an error when trying to use the `drafts` endpoints.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  New in Zulip 5.0 (feature level 87).  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableDraftsSynchronization bool `json:"enable_drafts_synchronization,omitempty"`
	// Whether the user has chosen for the layout width to be fluid.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	FluidLayoutWidth bool `json:"fluid_layout_width,omitempty"`
	// The [home view](zulip.com/help/configure-home-view) in Zulip, represented as the URL suffix after `#` to be rendered when Zulip loads.  Currently supported values are `all_messages` and `recent_topics`.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).  **Deprecated** in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	WebHomeView string `json:"web_home_view,omitempty"`
	// Whether has switched on high contrast mode.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	HighContrastMode bool `json:"high_contrast_mode,omitempty"`
	// Whether the user has chosen for the userlist to be displayed on the left side of the screen (for desktop app and web app) in narrow windows.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	LeftSideUserlist bool `json:"left_side_userlist,omitempty"`
	// Whether the user has chosen the number of starred messages to be displayed similar to unread counts.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	StarredMessageCounts bool `json:"starred_message_counts,omitempty"`
	// The user's [profile time zone](zulip.com/help/change-your-timezone, which is used primarily to display the user's local time to other users.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	Timezone string `json:"timezone,omitempty"`
	// Whether the user has chosen for emoticons to be translated into emoji in the Zulip compose box.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	TranslateEmoticons bool `json:"translate_emoticons,omitempty"`
	// Whether the user has chosen a twenty four hour time display (true) or a twelve hour one (false).  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	TwentyFourHourTime bool `json:"twenty_four_hour_time,omitempty"`
	// Whether the user setting for [sending on pressing Enter][set-enter-send] in the compose box is enabled.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.  Prior to Zulip 5.0 (feature level 84), this field was present in response if `realm_user` was present in `fetch_event_types`, not `update_display_settings`.  [capabilities]: /api/register-queue#parameter-client_capabilities [set-enter-send]: /help/configure-send-message-keys
	// Deprecated
	EnterSends bool `json:"enter_sends,omitempty"`
	// Array of dictionaries where each dictionary describes an emoji set supported by this version of the Zulip server.  Only relevant to clients with configuration UI for choosing an emoji set; the currently selected emoji set is available in the `emojiset` key.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EmojisetChoices []UserSettingsEmojisetChoice `json:"emojiset_choices,omitempty"`
}

// UserSettingsEmojisetChoice Object describing a emoji set.
type UserSettingsEmojisetChoice struct {
	// The key or the name of the emoji set which will be the value of `emojiset` if this emoji set is chosen.
	Key string `json:"key,omitempty"`
	// The text describing the emoji set.
	Text string `json:"text,omitempty"`
}

type RealmUser struct {
	// A array of dictionaries where each entry describes a user whose account has not been deactivated. Note that unlike the usual User dictionary, this does not contain the `is_active` key, as all the users present in this array have active accounts.  If the current user is a guest whose access to users is limited by a `can_access_all_users_group` policy, and the event queue was registered with the `user_list_incomplete` client capability, then users that the current user cannot access will not be included in this array. If the current user's access to a user is restricted but the client lacks this capability, then that inaccessible user will appear in the users array as an \"Unknown user\" object with the usual format but placeholder data whose only variable content is the user Id.  See also `cross_realm_bots` and `realm_non_active_users`.  **Changes**: Before Zulip 8.0 (feature level 232), the `user_list_incomplete` client capability did not exist, and so all clients whose access to a new user was prevented by `can_access_all_users_group` policy would receive a fake \"Unknown user\" event for such users.
	RealmUsers []User `json:"realm_users,omitempty"`
	// A array of dictionaries where each entry describes a user whose account has been deactivated. Note that unlike the usual User dictionary this does not contain the `is_active` key as all the users present in this array have deactivated accounts.
	RealmNonActiveUsers []User `json:"realm_non_active_users,omitempty"`
	// The avatar data source type for the current user.  Value values are `G` (gravatar) and `U` (uploaded by user).
	AvatarSource string `json:"avatar_source,omitempty"`
	// The avatar URL for the current user at 500x500 resolution, appropriate for use in settings UI showing the user's avatar.
	AvatarUrlMedium string `json:"avatar_url_medium,omitempty"`
	// The URL of the avatar for the current user at 100x100 resolution. See also `avatar_url_medium`.
	AvatarUrl string `json:"avatar_url,omitempty"`
	// Whether the current user is allowed to create at least one type of channel with the organization's [channel creation policy](zulip.com/help/configure-who-can-create-channels. Its value will always equal `can_create_public_streams || can_create_private_streams`.  **Changes**: Deprecated in Zulip 5.0 (feature level 102), when the new `create_private_stream_policy` and `create_public_stream_policy` properties introduced the possibility that a user could only create one type of channel.  This field will be removed in a future release.
	// Deprecated
	CanCreateStreams *bool `json:"can_create_streams,omitempty"`
	// Whether the current user is allowed to create public channels with the organization's [channel creation policy](zulip.com/help/configure-who-can-create-channels.  **Changes**: New in Zulip 5.0 (feature level 102). In older versions, the deprecated `can_create_streams` property should be used to determine whether the user can create public channels.
	CanCreatePublicStreams bool `json:"can_create_public_streams,omitempty"`
	// Whether the current user is allowed to create private channels with the organization's [channel creation policy](zulip.com/help/configure-who-can-create-channels.  **Changes**: New in Zulip 5.0 (feature level 102). In older versions, the deprecated `can_create_streams` property should be used to determine whether the user can create private channels.
	CanCreatePrivateStreams bool `json:"can_create_private_streams,omitempty"`
	// Whether the current user is allowed to create public channels with the organization's [channel creation policy](zulip.com/help/configure-who-can-create-channels.  Note that this will be false if the Zulip server does not have the `WEB_PUBLIC_STREAMS_ENABLED` setting enabled or if the organization has not enabled the `enable_spectator_access` realm setting.  **Changes**: New in Zulip 5.0 (feature level 103).
	CanCreateWebPublicStreams bool `json:"can_create_web_public_streams,omitempty"`
	// Whether the current user is allowed to subscribe other users to channels with the organization's [channels policy](zulip.com/help/configure-who-can-invite-to-channels.
	CanSubscribeOtherUsers bool `json:"can_subscribe_other_users,omitempty"`
	// Whether the current user [is allowed to invite others][who-can-send-invitations] to the organization.  **Changes**: New in Zulip 4.0 (feature level 51).  [who-can-send-invitations]: /help/restrict-account-creation#change-who-can-send-invitations
	CanInviteOthersToRealm bool `json:"can_invite_others_to_realm,omitempty"`
	// Whether the current user is at least an [organization administrator](zulip.com/api/roles-and-permissions.
	IsAdmin bool `json:"is_admin,omitempty"`
	// Whether the current user is an [organization owner](zulip.com/api/roles-and-permissions.  **Changes**: New in Zulip 3.0 (feature level 11).
	IsOwner bool `json:"is_owner,omitempty"`
	// Whether the current user is at least an [organization moderator](zulip.com/api/roles-and-permissions.  **Changes**: Prior to Zulip 11.0 (feature level 380), this was only true for users whose role was exactly the moderator role.  New in Zulip 4.0 (feature level 60).
	IsModerator bool `json:"is_moderator,omitempty"`
	// Whether the current user is a [guest user](zulip.com/api/roles-and-permissions.
	IsGuest bool `json:"is_guest,omitempty"`
	// The unique Id for the current user.
	UserId int64 `json:"user_id,omitempty"`
	// The Zulip API email address for the current user. See also `delivery_email`; these may be the same or different depending on the user's `email_address_visibility` policy.
	Email string `json:"email,omitempty"`
	// The user's email address, appropriate for UI for changing the user's email address. See also `email`.
	DeliveryEmail string `json:"delivery_email,omitempty"`
	// The full name of the current user.
	FullName string `json:"full_name,omitempty"`
	// Array of dictionaries where each dictionary contains details of a single cross realm bot. Cross-realm bots are special system bot accounts like Notification Bot.  Most clients will want to combine this with `realm_users` in many contexts.
	CrossRealmBots []UserWithIsSystemBot `json:"cross_realm_bots,omitempty"`
}

type GlobalNotifications struct {
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableDesktopNotifications bool `json:"enable_desktop_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableDigestEmails bool `json:"enable_digest_emails,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableLoginEmails bool `json:"enable_login_emails,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableMarketingEmails bool `json:"enable_marketing_emails,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EmailNotificationsBatchingPeriodSeconds int64 `json:"email_notifications_batching_period_seconds,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableOfflineEmailNotifications bool `json:"enable_offline_email_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableOfflinePushNotifications bool `json:"enable_offline_push_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableOnlinePushNotifications bool `json:"enable_online_push_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableSounds bool `json:"enable_sounds,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableStreamDesktopNotifications bool `json:"enable_stream_desktop_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableStreamEmailNotifications bool `json:"enable_stream_email_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableStreamPushNotifications bool `json:"enable_stream_push_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableStreamAudibleNotifications bool `json:"enable_stream_audible_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	WildcardMentionsNotify bool `json:"wildcard_mentions_notify,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	MessageContentInEmailNotifications bool `json:"message_content_in_email_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	NotificationSound string `json:"notification_sound,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	PmContentInDesktopNotifications bool `json:"pm_content_in_desktop_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	DesktopIconCountDisplay int64 `json:"desktop_icon_count_display,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: In Zulip 7.0 (feature level 168), replaced previous `realm_name_in_notifications` global notifications setting with `realm_name_in_email_notifications_policy`.  **Deprecated** since Zulip 5.0 (feature level 89); both `realm_name_in_notifications` and the newer `realm_name_in_email_notifications_policy` are deprecated. Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	RealmNameInEmailNotificationsPolicy int64 `json:"realm_name_in_email_notifications_policy,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	PresenceEnabled bool `json:"presence_enabled,omitempty"`
	// Array containing the names of the notification sound options supported by this Zulip server. Only relevant to support UI for configuring notification sounds.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	AvailableNotificationSounds []string `json:"available_notification_sounds,omitempty"`
}

type Realm struct {
	// The UNIX timestamp (UTC) for when the organization was created.  **Changes**: New in Zulip 8.0 (feature level 203).
	DateCreated time.Time `json:"realm_date_created,omitempty"`
	// The maximum allowed length for a channel name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required `stream` in `fetch_event_types`, was called `stream_name_max_length`, and always had a value of 60.
	MaxStreamNameLength int64 `json:"max_stream_name_length,omitempty"`
	// The maximum allowed length for a channel description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this required `stream` in `fetch_event_types`, was called `stream_description_max_length`, and always had a value of 1024.
	MaxStreamDescriptionLength int64 `json:"max_stream_description_length,omitempty"`
	// The maximum allowed length for a channel folder name, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 60 as a fallback value on previous feature levels.
	MaxChannelFolderNameLength int64 `json:"max_channel_folder_name_length,omitempty"`
	// The maximum allowed length for a channel folder description, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 11.0 (feature level 410). Clients should use 1024 as a fallback value on previous feature levels.
	MaxChannelFolderDescriptionLength int64 `json:"max_channel_folder_description_length,omitempty"`
	// The maximum allowed length for a topic, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 60.
	MaxTopicLength int64 `json:"max_topic_length,omitempty"`
	// The maximum allowed length for a message, in Unicode code points. Clients should use this property rather than hardcoding field sizes.  **Changes**: New in Zulip 4.0 (feature level 53). Previously, this property always had a value of 10000.
	MaxMessageLength int64 `json:"max_message_length,omitempty"`
	// The minimum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If `null`, then a deactivated organization's data can be deleted immediately.  **Changes**: New in Zulip 10.0 (feature level 332)
	ServerMinDeactivatedRealmDeletionDays int64 `json:"server_min_deactivated_realm_deletion_days,omitempty"`
	// The maximum permitted number of days before full data deletion (users, channels, messages, etc.) of a deactivated organization. If `null`, then a deactivated organization's data can be retained indefinitely.  **Changes**: New in Zulip 10.0 (feature level 332).
	ServerMaxDeactivatedRealmDeletionDays int64 `json:"server_max_deactivated_realm_deletion_days,omitempty"`
	// For clients implementing the [presence](zulip.com/api/get-presence) system, the time interval the client should use for sending presence requests to the server (and thus receive presence updates from the server).  It is important for presence implementations to use both this and `server_presence_offline_threshold_seconds` correctly, so that a Zulip server can change these values to manage the trade-off between load and freshness of presence data.  **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 60 for older Zulip servers, since that's the value that was hardcoded in the Zulip mobile apps prior to this parameter being introduced.
	ServerPresencePingIntervalSeconds int64 `json:"server_presence_ping_interval_seconds,omitempty"`
	// How old a presence timestamp for a given user can be before the user should be displayed as offline by clients displaying Zulip presence data. See the related `server_presence_ping_interval_seconds` for details.  **Changes**: New in Zulip 7.0 (feature level 164). Clients should use 140 for older Zulip servers, since that's the value that was hardcoded in the Zulip client apps prior to this parameter being introduced.
	ServerPresenceOfflineThresholdSeconds int64 `json:"server_presence_offline_threshold_seconds,omitempty"`
	// For clients implementing [typing notifications](zulip.com/api/set-typing-status) protocol, the time interval in milliseconds that the client should wait for additional [typing start](zulip.com/api/get-events#typing-start events from the server before removing an active typing indicator.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 15000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	ServerTypingStartedExpiryPeriodMilliseconds int64 `json:"server_typing_started_expiry_period_milliseconds,omitempty"`
	// For clients implementing [typing notifications](zulip.com/api/set-typing-status) protocol, the time interval in milliseconds that the client should wait when a user stops interacting with the compose UI before sending a stop notification to the server.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 5000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	ServerTypingStoppedWaitPeriodMilliseconds int64 `json:"server_typing_stopped_wait_period_milliseconds,omitempty"`
	// For clients implementing [typing notifications](zulip.com/api/set-typing-status) protocol, the time interval in milliseconds that the client should use to send regular start notifications to the server to indicate that the user is still actively interacting with the compose UI.  **Changes**: New in Zulip 8.0 (feature level 204). Clients should use 10000 for older Zulip servers, since that's the value that was hardcoded in the Zulip apps prior to this parameter being introduced.
	ServerTypingStartedWaitPeriodMilliseconds int64 `json:"server_typing_started_wait_period_milliseconds,omitempty"`
	// What typesof message edit history are accessible to users via [message edit history](zulip.com/help/view-a-messages-edit-history.  - \"all\" = All edit history is visible. - \"moves\" = Only moves are visible. - \"none\" = No edit history is visible.  **Changes**: New in Zulip 10.0 (feature level 358), replacing the previous `allow_edit_history` boolean setting; `true` corresponds to `all`, and `false` to `none`.
	MessageEditHistoryVisibilityPolicy string `json:"realm_message_edit_history_visibility_policy,omitempty"`
	// Whether this organization is configured to allow users to access [message edit history](zulip.com/help/view-a-messages-edit-history.  The value of `realm_allow_edit_history` is set as `false` if the `realm_message_edit_history_visibility_policy` is configured as \"None\" and `true` if it is configured as \"Moves only\" or \"All\".  **Changes**: Deprecated in Zulip 10.0 (feature level 358) and will be removed in the future, as it is an inaccurate version `realm_message_edit_history_visibility_policy`, which replaces this field.
	// Deprecated
	AllowEditHistory bool `json:"realm_allow_edit_history,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to add custom emoji in the organization.  **Changes**: New in Zulip 10.0 (feature level 307). Previously, this permission was controlled by the enum `add_custom_emoji_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators.  Before Zulip 5.0 (feature level 85), the `realm_add_emoji_by_admins_only` boolean setting controlled this permission; `true` corresponded to `Admins`, and `false` to `Everyone`.
	CanAddCustomEmojiGroup GroupSettingValue `json:"realm_can_add_custom_emoji_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to add subscribers to channels in the organization.  **Changes**: New in Zulip 10.0 (feature level 341). Previously, this permission was controlled by the enum `invite_to_stream_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators.
	CanAddSubscribersGroup GroupSettingValue `json:"realm_can_add_subscribers_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to delete any message in the organization.  **Changes**: New in Zulip 10.0 (feature level 281). Previously, this permission was limited to administrators only and was uneditable.
	CanDeleteAnyMessageGroup GroupSettingValue `json:"realm_can_delete_any_message_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to delete messages that they have sent in the organization.  **Changes**: New in Zulip 10.0 (feature level 291). Previously, this permission was controlled by the enum `delete_own_message_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 5=Everyone.  Before Zulip 5.0 (feature level 101), the `allow_message_deleting` boolean setting controlled this permission; `true` corresponded to `Everyone`, and `false` to `Admins`.
	CanDeleteOwnMessageGroup GroupSettingValue `json:"realm_can_delete_own_message_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to change per-channel `can_delete_any_message_group` and `can_delete_own_message_group` permission settings. Note that the user must be a member of both this group and the `can_administer_channel_group` of the channel whose message delete settings they want to change.  Organization administrators can always change these settings of every channel.  **Changes**: New in Zulip 11.0 (feature level 407).
	CanSetDeleteMessagePolicyGroup GroupSettingValue `json:"realm_can_set_delete_message_policy_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to change per-channel `topics_policy` setting. Note that the user must be a member of both this group and the `can_administer_channel_group` of the channel whose `topics_policy` they want to change.  Organization administrators can always change the `topics_policy` setting of every channel.  **Changes**: New in Zulip 11.0 (feature level 392).
	CanSetTopicsPolicyGroup GroupSettingValue `json:"realm_can_set_topics_policy_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to send email invitations for inviting other users to the organization.  **Changes**: New in Zulip 10.0 (feature level 321). Previously, this permission was controlled by the enum `invite_to_realm_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 6=Nobody.  Before Zulip 4.0 (feature level 50), the `invite_by_admins_only` boolean setting controlled this permission; `true` corresponded to `Admins`, and `false` to `Members`.
	CanInviteUsersGroup GroupSettingValue `json:"realm_can_invite_users_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.  **Changes**: New in Zulip 10.0 (feature level 352). Previously, this permission was controlled by the enum `wildcard_mention_policy`.
	CanMentionManyUsersGroup GroupSettingValue `json:"realm_can_mention_many_users_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to move messages from one channel to another in the organization.  **Changes**: New in Zulip 10.0 (feature level 310). Previously, this permission was controlled by the enum `move_messages_between_streams_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 6=Nobody.  In Zulip 7.0 (feature level 159), `Nobody` was added as an option to `move_messages_between_streams_policy` enum.
	CanMoveMessagesBetweenChannelsGroup GroupSettingValue `json:"realm_can_move_messages_between_channels_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to move messages from one topic to another within a channel in the organization.  **Changes**: New in Zulip 10.0 (feature level 316). Previously, this permission was controlled by the enum `edit_topic_policy`. Values were 1=Members, 2=Admins, 3=Full members, 4=Moderators, 5=Everyone, 6=Nobody.  In Zulip 7.0 (feature level 159), `Nobody` was added as an option to `edit_topic_policy` enum.
	CanMoveMessagesBetweenTopicsGroup GroupSettingValue `json:"realm_can_move_messages_between_topics_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create user groups in this organization.  **Changes**: New in Zulip 10.0 (feature level 299). Previously `realm_user_group_edit_policy` field used to control the permission to create user groups.
	CanCreateGroups GroupSettingValue `json:"realm_can_create_groups,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create all types of bot users in the organization. See also `can_create_write_only_bots_group`.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum `bot_creation_policy`. Values were 1=Members, 2=Generic bots limited to administrators, 3=Administrators.
	CanCreateBotsGroup GroupSettingValue `json:"realm_can_create_bots_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create bot users that can only send messages in the organization, i.e. incoming webhooks, in addition to the users who are present in `can_create_bots_group`.  **Changes**: New in Zulip 10.0 (feature level 344). Previously, this permission was controlled by the enum `bot_creation_policy`. Values were 1=Members, 2=Generic bots limited to administrators, 3=Administrators.
	CanCreateWriteOnlyBotsGroup GroupSettingValue `json:"realm_can_create_write_only_bots_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to administer all existing groups in this organization.  **Changes**: Prior to Zulip 10.0 (feature level 305), only users who were a member of the group or had the moderator role or above could exercise the permission on a given group.  New in Zulip 10.0 (feature level 299). Previously the `user_group_edit_policy` field controlled the permission to manage user groups. Valid values were as follows:  - 1 = All members can create and edit user groups - 2 = Only organization administrators can create and edit   user groups - 3 = Only [full members][calc-full-member] can create and   edit user groups. - 4 = Only organization administrators and moderators can   create and edit user groups.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	CanManageAllGroups GroupSettingValue `json:"realm_can_manage_all_groups,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to manage plans and billing in the organization.  **Changes**: New in Zulip 10.0 (feature level 363). Previously, only owners and users with `is_billing_admin` property set to `true` were allowed to manage plans and billing.
	CanManageBillingGroup GroupSettingValue `json:"realm_can_manage_billing_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create public channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 264). Previously `realm_create_public_stream_policy` field used to control the permission to create public channels.
	CanCreatePublicChannelGroup GroupSettingValue `json:"realm_can_create_public_channel_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create private channels in this organization.  **Changes**: New in Zulip 9.0 (feature level 266). Previously `realm_create_private_stream_policy` field used to control the permission to create private channels.
	CanCreatePrivateChannelGroup GroupSettingValue `json:"realm_can_create_private_channel_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to create web-public channels in this organization.  Has no effect and should not be displayed in settings UI unless the Zulip server has the `WEB_PUBLIC_STREAMS_ENABLED` server-level setting enabled and the organization has enabled the `enable_spectator_access` realm setting.  **Changes**: New in Zulip 10.0 (feature level 280). Previously `realm_create_web_public_stream_policy` field used to control the permission to create web-public channels.
	CanCreateWebPublicChannelGroup GroupSettingValue `json:"realm_can_create_web_public_channel_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to [resolve topics](zulip.com/help/resolve-a-topic) in the organization.  **Changes**: New in Zulip 10.0 (feature level 367). Previously, permission to resolve topics was controlled by the more general `can_move_messages_between_topics_group permission for moving messages`.
	CanResolveTopicsGroup GroupSettingValue `json:"realm_can_resolve_topics_group,omitempty"`
	// A deprecated representation of a superset of the users who have permission to create public channels in the organization, available for backwards-compatibility. Clients should use `can_create_public_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 = Members only - 2 = Admins only - 3 = [Full members][calc-full-member] only - 4 = Admins and moderators only  **Changes**: Deprecated in Zulip 9.0 (feature level 264) and replaced by `realm_can_create_public_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the `realm_create_stream_policy` setting.  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	CreatePublicStreamPolicy int32 `json:"realm_create_public_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to create private channels in the organization, available for backwards-compatibility. Clients should use `can_create_private_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 = Members only - 2 = Admins only - 3 = [Full members][calc-full-member] only - 4 = Admins and moderators only  **Changes**: Deprecated in Zulip 9.0 (feature level 266) and replaced by `realm_can_create_private_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  **Changes**: Before Zulip 5.0 (feature level 102), permission to create channels was controlled by the `realm_create_stream_policy` setting.  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	CreatePrivateStreamPolicy int32 `json:"realm_create_private_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to create web-public channels in the organization, available for backwards-compatibility. Clients should use `can_create_web_public_channel_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 2 = Admins only - 4 = Admins and moderators only - 6 = Nobody - 7 = Owners only  **Changes**: Deprecated in Zulip 10.0 (feature level 280) and replaced by `realm_can_create_web_public_channel_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  **Changes**: Added in Zulip 5.0 (feature level 103).  [permission-level]: /api/roles-and-permissions#permission-levels
	// Deprecated
	CreateWebPublicStreamPolicy int32 `json:"realm_create_web_public_stream_policy,omitempty"`
	// A deprecated representation of a superset of the users who have permission to use wildcard mentions in large channels, available for backwards-compatibility. Clients should use `can_mention_many_users_group` instead.  It is an enum with the following possible values, corresponding to roles/system groups:  - 1 = Any user can use wildcard mentions in large channels. - 2 = Only members can use wildcard mentions in large channels. - 3 = Only [full members][calc-full-member] can use wildcard mentions in large channels. - 5 = Only organization administrators can use wildcard mentions in large channels. - 6 = Nobody can use wildcard mentions in large channels. - 7 = Only organization administrators and moderators can use wildcard mentions in large channels.  All users will receive a warning/reminder when using mentions in large channels, even when permitted to do so.  **Changes**: Deprecated in Zulip 10.0 (feature level 352) and replaced by `realm_can_mention_many_users_group`, which supports finer resolution of configurations, resulting in this property being inaccurate following that transition.  Channel administrators option removed in Zulip 6.0 (feature level 133).  Moderators option added in Zulip 4.0 (feature level 62).  New in Zulip 4.0 (feature level 33).  [permission-level]: /api/roles-and-permissions#permission-levels [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	// Deprecated
	WildcardMentionPolicy int32 `json:"realm_wildcard_mention_policy,omitempty"`
	// The [organization language][org-lang] for automated messages and invitation emails.  [org-lang]: /help/configure-organization-language
	DefaultLanguage string `json:"realm_default_language,omitempty"`
	// This organization's configured custom message for Welcome Bot to send to new user accounts, in Zulip Markdown format.  **Changes**: New in Zulip 11.0 (feature level 416).
	WelcomeMessageCustomText string `json:"realm_welcome_message_custom_text,omitempty"`
	// The description of the organization, used on login and registration pages.
	Description string `json:"realm_description,omitempty"`
	// Whether the organization has enabled [weekly digest emails](zulip.com/help/digest-emails.
	DigestEmailsEnabled bool `json:"realm_digest_emails_enabled,omitempty"`
	// Whether the organization disallows disposable email addresses.
	DisallowDisposableEmailAddresses bool `json:"realm_disallow_disposable_email_addresses,omitempty"`
	// Whether users are allowed to change their own email address in this organization. This is typically disabled for organizations that synchronize accounts from LDAP or a similar corporate database.
	EmailChangesDisabled bool `json:"realm_email_changes_disabled,omitempty"`
	// Whether an invitation is required to join this organization.
	InviteRequired bool `json:"realm_invite_required,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who are allowed to create [reusable invitation links](zulip.com/help/invite-new-users#create-a-reusable-invitation-link to the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 209).
	CreateMultiuseInviteGroup GroupSettingValue `json:"realm_create_multiuse_invite_group,omitempty"`
	// Whether this organization has been configured to enable [previews of linked images](zulip.com/help/image-video-and-website-previews.
	InlineImagePreview bool `json:"realm_inline_image_preview,omitempty"`
	// Whether this organization has been configured to enable [previews of linked websites](zulip.com/help/image-video-and-website-previews.
	InlineUrlEmbedPreview bool `json:"realm_inline_url_embed_preview,omitempty"`
	// The organization's default policy for sending channel messages to the [empty \"general chat\" topic](zulip.com/help/require-topics.  - `\"allow_empty_topic\"`: Channel messages can be sent to the empty topic. - `\"disable_empty_topic\"`: Channel messages cannot be sent to the empty topic.  **Changes**: New in Zulip 11.0 (feature level 392). Previously, this was controlled by the boolean `realm_mandatory_topics` setting, which is now deprecated.
	TopicsPolicy TopicsPolicy `json:"realm_topics_policy,omitempty"`
	// Whether [topics are required](zulip.com/help/require-topics) for messages in this organization.  **Changes**: Deprecated in Zulip 11.0 (feature level 392). This is now controlled by the realm `topics_policy` setting.
	// Deprecated
	MandatoryTopics bool `json:"realm_mandatory_topics,omitempty"`
	// The default [message retention policy](zulip.com/help/message-retention-policy) for this organization. It can have one special value:  - `-1` denoting that the messages will be retained forever for this realm, by default.  **Changes**: Prior to Zulip 3.0 (feature level 22), no limit was encoded as `null` instead of `-1`. Clients can correctly handle all server versions by treating both `-1` and `null` as indicating unlimited message retention.
	MessageRetentionDays int64 `json:"realm_message_retention_days,omitempty"`
	// The name of the organization, used in login pages etc.
	Name string `json:"realm_name,omitempty"`
	// Whether this realm is configured to disallow sending mobile push notifications with message content through the legacy mobile push notifications APIs. The new API uses end-to-end encryption to protect message content and metadata from being accessible to the push bouncer service, APNs, and FCM. Clients that support the new E2EE API will use it automatically regardless of this setting.  If `true`, mobile push notifications sent to clients that lack support for E2EE push notifications will always have \"New message\" as their content. Note that these legacy mobile notifications will still contain metadata, which may include the message's Id, the sender's name, email address, and avatar.  In a future release, once the official mobile apps have implemented fully validated their E2EE protocol support, this setting will become strict, and disable the legacy protocol entirely.  **Changes**: New in Zulip 11.0 (feature level 409). Previously, this behavior was available only via the `PUSH_NOTIFICATION_REDACT_CONTENT` global server setting.
	RequireE2eePushNotifications bool `json:"realm_require_e2ee_push_notifications,omitempty"`
	// Indicates whether the organization is configured to require users to have unique full names. If true, the server will reject attempts to create a new user, or change the name of an existing user, where doing so would lead to two users whose names are identical modulo case and unicode normalization.  **Changes**: New in Zulip 9.0 (feature level 246). Previously, the Zulip server could not be configured to enforce unique names.
	RequireUniqueNames bool `json:"realm_require_unique_names,omitempty"`
	// Indicates whether users are [allowed to change](zulip.com/help/restrict-name-and-email-changes) their name via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.
	NameChangesDisabled bool `json:"realm_name_changes_disabled,omitempty"`
	// Indicates whether users are [allowed to change](zulip.com/help/restrict-name-and-email-changes) their avatar via the Zulip UI in this organization. Typically disabled in organizations syncing this type of account information from an external user database like LDAP.
	AvatarChangesDisabled bool `json:"realm_avatar_changes_disabled,omitempty"`
	// Whether [new users joining](zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions this organization are required to have an email address in one of the `realm_domains` configured for the organization.
	EmailsRestrictedToDomains bool `json:"realm_emails_restricted_to_domains,omitempty"`
	// Whether or not this organization is configured to send the standard Zulip [welcome emails](zulip.com/help/disable-welcome-emails) to new users joining the organization.
	SendWelcomeEmails bool `json:"realm_send_welcome_emails,omitempty"`
	// Whether notification emails in this organization are allowed to contain Zulip the message content, or simply indicate that a new message was sent.
	MessageContentAllowedInEmailNotifications bool `json:"realm_message_content_allowed_in_email_notifications,omitempty"`
	// Whether web-public channels and related anonymous access APIs/features are enabled in this organization.  Can only be enabled if the `WEB_PUBLIC_STREAMS_ENABLED` [server setting][server-settings] is enabled on the Zulip server. See also the `can_create_web_public_channel_group` realm setting.  **Changes**: New in Zulip 5.0 (feature level 109).  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html
	EnableSpectatorAccess bool `json:"realm_enable_spectator_access,omitempty"`
	// Whether the organization has given permission to be advertised in the Zulip [communities directory](zulip.com/help/communities-directory.  Useful only to clients supporting changing this setting for the organization.  Giving permission via this setting does not guarantee that an organization will be listed in the Zulip communities directory.  **Changes**: New in Zulip 6.0 (feature level 129).
	WantAdvertiseInCommunitiesDirectory bool `json:"realm_want_advertise_in_communities_directory,omitempty"`
	// The configured [video call provider](zulip.com/help/configure-call-provider) for the organization.  - 0 = None - 1 = Jitsi Meet - 3 = Zoom (User OAuth integration) - 4 = BigBlueButton - 5 = Zoom (Server to Server OAuth integration)  Note that only one of the [Zoom integrations][zoom-video-calls] can be configured on a Zulip server.  **Changes**: In Zulip 10.0 (feature level 353), added the Zoom Server to Server OAuth option.  In Zulip 3.0 (feature level 1), added the None option to disable video call UI.  [zoom-video-calls]: https://zulip.readthedocs.io/en/latest/production/video-calls.html#zoom
	VideoChatProvider VideoChatProvider `json:"realm_video_chat_provider,omitempty"`
	// The URL of the custom Jitsi Meet server configured in this organization's settings.  `null`, the default, means that the organization is using the should use the server-level configuration, `server_jitsi_server_url`. A correct client supporting only the modern API should use `realm_jitsi_server_url || server_jitsi_server_url` to create calls.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this was only available as a server-level configuration, which was available via the `jitsi_server_url` field.
	JitsiServerUrl *string `json:"realm_jitsi_server_url,omitempty"`
	// The configured GIPHY rating for the organization.  **Changes**: New in Zulip 4.0 (feature level 55).
	GiphyRating int64 `json:"realm_giphy_rating,omitempty"`
	// Members whose accounts have been created at least this many days ago will be treated as [full members][calc-full-member] for the purpose of settings that restrict access to new members.  [calc-full-member]: /api/roles-and-permissions#determining-if-a-user-is-a-full-member
	WaitingPeriodThreshold int64 `json:"realm_waiting_period_threshold,omitempty"`
	// The day of the week when the organization will send its weekly digest email to inactive users.
	DigestWeekday int `json:"realm_digest_weekday,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to start a new direct message conversation involving other non-bot users. Users who are outside this group and attempt to send the first direct message to a given collection of recipient users will receive an error, unless all other recipients are bots or the sender.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the `private_message_policy` realm setting, which supported values of 1 (enabled) and 2 (disabled).
	DirectMessageInitiatorGroup GroupSettingValue `json:"realm_direct_message_initiator_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who have permission to fully use direct messages. Users outside this group can only send direct messages to conversations where all the recipients are in this group, are bots, or are the sender, ensuring that every direct message conversation will be visible to at least one user in this group.  **Changes**: New in Zulip 9.0 (feature level 270).  Previously, access to send direct messages was controlled by the `private_message_policy` realm setting, which supported values of 1 (enabled) and 2 (disabled).
	DirectMessagePermissionGroup GroupSettingValue `json:"realm_direct_message_permission_group,omitempty"`
	// The default pygments language code to be used for code blocks in this organization. If an empty string, no default has been set.  **Changes**: Prior to Zulip 8.0 (feature level 195), a server bug meant that both `null` and an empty string could represent that no default was set for this realm setting. Clients supporting older server versions should treat either value (`null` or `\"\"`) as no default being set.
	DefaultCodeBlockLanguage *string `json:"realm_default_code_block_language,omitempty"`
	// Messages sent more than this many seconds ago cannot be deleted with this organization's [message deletion policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be 0. A `null` value means no limit: messages can be deleted regardless of how long ago they were sent.  **Changes**: No limit was represented using the special value `0` before Zulip 5.0 (feature level 100).
	MessageContentDeleteLimitSeconds *int64 `json:"realm_message_content_delete_limit_seconds,omitempty"`
	// Dictionary of authentication method keys mapped to dictionaries that describe the properties of the named authentication method for the organization - its enabled status and availability for use by the organization.  Clients should use this to implement server-settings UI to change which methods are enabled for the organization. For authentication UI itself, clients should use the pre-authentication metadata returned by [`GET /server_settings`](zulip.com/api/get-server-settings.  **Changes**: In Zulip 9.0 (feature level 241), the values in this dictionary were changed. Previously, the values were a simple boolean indicating whether the backend is enabled or not.
	AuthenticationMethods map[string]RealmAuthenticationMethod `json:"realm_authentication_methods,omitempty"`
	// Whether this organization's [message edit policy][config-message-editing] allows editing the content of messages.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  [config-message-editing]: /help/restrict-message-editing-and-deletion
	AllowMessageEditing *bool `json:"realm_allow_message_editing,omitempty"`
	// Messages sent more than this many seconds ago cannot be edited with this organization's [message edit policy](zulip.com/help/restrict-message-editing-and-deletion.  Will not be `0`. A `null` value means no limit, so messages can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: Before Zulip 6.0 (feature level 138), no limit was represented using the special value `0`.
	MessageContentEditLimitSeconds *int64 `json:"realm_message_content_edit_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved within a channel to another topic by users who have permission to do so based on this organization's [topic edit policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so message topics can be edited regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, this time limit was always 72 hours for users who were not administrators or moderators.
	MoveMessagesWithinStreamLimitSeconds *int64 `json:"realm_move_messages_within_stream_limit_seconds,omitempty"`
	// Messages sent more than this many seconds ago cannot be moved between channels by users who have permission to do so based on this organization's [message move policy](zulip.com/help/restrict-moving-messages. This setting does not affect moderators and administrators.  Will not be `0`. A `null` value means no limit, so messages can be moved regardless of how long ago they were sent.  See [`PATCH /messages/{message_id}`](zulip.com/api/update-message) for details and history of how message editing permissions work.  **Changes**: New in Zulip 7.0 (feature level 162). Previously, there was no time limit for moving messages between channels for users with permission to do so.
	MoveMessagesBetweenStreamsLimitSeconds *int64 `json:"realm_move_messages_between_streams_limit_seconds,omitempty"`
	// Whether read receipts is enabled in the organization or not.  If disabled, read receipt data will be unavailable to clients, regardless of individual users' personal read receipt settings. See also the `send_read_receipts` setting within `realm_user_settings_defaults`.  **Changes**: New in Zulip 6.0 (feature level 137).
	EnableReadReceipts bool `json:"realm_enable_read_receipts,omitempty"`
	// The URL of the organization's [profile icon](zulip.com/help/create-your-organization-profile.
	IconUrl string `json:"realm_icon_url,omitempty"`
	// String indicating whether the organization's [profile icon](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's icon.  - \"G\" means generated by Gravatar (the default). - \"U\" means uploaded by an organization administrator.
	IconSource string `json:"realm_icon_source,omitempty"`
	// The maximum file size allowed for the organization's icon. Useful for UI allowing editing the organization's icon.  **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called `max_icon_file_size`.
	MaxIconFileSizeMib int64 `json:"max_icon_file_size_mib,omitempty"`
	// The URL of the organization's wide logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.
	LogoUrl string `json:"realm_logo_url,omitempty"`
	// String indicating whether the organization's [profile wide logo](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - \"D\" means the logo is the default Zulip logo. - \"U\" means uploaded by an organization administrator.
	LogoSource string `json:"realm_logo_source,omitempty"`
	// The URL of the organization's dark theme wide-format logo configured in the [organization profile](zulip.com/help/create-your-organization-profile.
	NightLogoUrl string `json:"realm_night_logo_url,omitempty"`
	// String indicating whether the organization's dark theme [profile wide logo](zulip.com/help/create-your-organization-profile) was uploaded by a user or is the default. Useful for UI allowing editing the organization's wide logo.  - \"D\" means the logo is the default Zulip logo. - \"U\" means uploaded by an organization administrator.
	NightLogoSource string `json:"realm_night_logo_source,omitempty"`
	// The maximum file size allowed for the uploaded organization logos.  **Changes**: New in Zulip 5.0 (feature level 72). Previously, this was called `max_logo_file_size`.
	MaxLogoFileSizeMib int64 `json:"max_logo_file_size_mib,omitempty"`
	// The fake email domain that will be used for new bots created this organization. Useful for UI for creating bots.
	BotDomain string `json:"realm_bot_domain,omitempty"`
	// The URL for the organization. Alias of `realm_url`.  **Changes**: Deprecated in Zulip 9.0 (feature level 257). The term \"URI\" is deprecated in [web standards](https://url.spec.whatwg.org/#goals).
	// Deprecated
	Uri string `json:"realm_uri,omitempty"`
	// The URL for the organization.  **Changes**: New in Zulip 9.0 (feature level 257), replacing the deprecated `realm_uri`.
	Url string `json:"realm_url,omitempty"`
	// Dictionary where each entry describes a supported [video call provider](zulip.com/help/configure-call-provider) that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the realm setting `video_chat_provider`.
	AvailableVideoChatProviders map[string]VideoChatProviderInfo `json:"realm_available_video_chat_providers,omitempty"`
	// Whether online presence of other users is shown in this organization.
	PresenceDisabled bool `json:"realm_presence_disabled,omitempty"`
	// Whether this Zulip server is configured to allow organizations to enable [digest emails](zulip.com/help/digest-emails.  Relevant for administrative settings UI that can change the digest email settings.
	SettingsSendDigestEmails bool `json:"settings_send_digest_emails,omitempty"`
	// Whether the organization is a Zephyr mirror realm.
	IsZephyrMirrorRealm bool `json:"realm_is_zephyr_mirror_realm,omitempty"`
	// Whether the organization has enabled Zulip's default email and password authentication feature. Determines whether Zulip stores a password for the user and clients should offer any UI for changing the user's Zulip password.
	EmailAuthEnabled bool `json:"realm_email_auth_enabled,omitempty"`
	// Whether the organization allows any sort of password-based authentication (whether via EmailAuthBackend or LDAP passwords).  Determines whether a client might ever need to display a password prompt (clients will primarily look at this attribute in [server_settings](zulip.com/api/get-server-settings) before presenting a login page).
	PasswordAuthEnabled bool `json:"realm_password_auth_enabled,omitempty"`
	// Whether push notifications are enabled for this organization. Typically `true` for Zulip Cloud and self-hosted realms that have a valid registration for the [Mobile push notifications service](https://zulip.readthedocs.io/en/latest/production/mobile-push-notifications.html), and `false` for self-hosted servers that do not.  **Changes**: Before Zulip 8.0 (feature level 231), this incorrectly was `true` for servers that were partly configured to use the Mobile Push Notifications Service but not properly registered.
	PushNotificationsEnabled bool `json:"realm_push_notifications_enabled,omitempty"`
	// If the server expects the realm's push notifications access to end at a definite time in the future, the UNIX timestamp (UTC) at which this is expected to happen. Mobile clients should use this field to display warnings to users when the indicated timestamp is near.  **Changes**: New in Zulip 8.0 (feature level 231).
	PushNotificationsEnabledEndTimestamp *time.Time `json:"realm_push_notifications_enabled_end_timestamp,omitempty"`
	// The total quota for uploaded files in this organization.  Clients are not responsible for checking this quota; it is included in the API only for display purposes.  If `null`, there is no limit.  **Changes**: Before Zulip 9.0 (feature level 251), this field was incorrectly measured in bytes, not MiB.  New in Zulip 5.0 (feature level 72). Previously, this was called `realm_upload_quota`.
	UploadQuotaMib *int64 `json:"realm_upload_quota_mib,omitempty"`
	// The [organization type](zulip.com/help/organization-type) for the realm. Useful only to clients supporting changing this setting for the organization, or clients implementing onboarding content or other features that varies with organization type.  - 0 = Unspecified - 10 = Business - 20 = Open-source project - 30 = Education (non-profit) - 35 = Education (for-profit) - 40 = Research - 50 = Event or conference - 60 = Non-profit (registered) - 70 = Government - 80 = Political group - 90 = Community - 100 = Personal - 1000 = Other  **Changes**: New in Zulip 6.0 (feature level 128).
	OrgType OrgType `json:"realm_org_type,omitempty"`
	// The plan type of the organization.  - 1 = Self-hosted organization (SELF_HOSTED) - 2 = Zulip Cloud free plan (LIMITED) - 3 = Zulip Cloud Standard plan (STANDARD) - 4 = Zulip Cloud Standard plan, sponsored for free (STANDARD_FREE)
	PlanType PlanType `json:"realm_plan_type,omitempty"`
	// Whether clients should show a warning when a user is composing a DM to a guest user in this organization.  **Changes**: New in Zulip 10.0 (feature level 348).
	EnableGuestUserDmWarning bool `json:"realm_enable_guest_user_dm_warning,omitempty"`
	// Whether clients should display \"(guest)\" after the names of guest users to prominently highlight their status.  **Changes**: New in Zulip 8.0 (feature level 216).
	EnableGuestUserIndicator bool `json:"realm_enable_guest_user_indicator,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who are allowed to access all users in the organization.  **Changes**: Prior to Zulip 10.0 (feature level 314), this value used to be of type integer and did not accept anonymous user groups.  New in Zulip 8.0 (feature level 225).
	CanAccessAllUsersGroup GroupSettingValue `json:"realm_can_access_all_users_group,omitempty"`
	// A [group-setting value](zulip.com/api/group-setting-values) defining the set of users who are allowed to use AI summarization.  **Changes**: New in Zulip 10.0 (feature level 350).
	CanSummarizeTopicsGroup GroupSettingValue `json:"realm_can_summarize_topics_group,omitempty"`
	// Whether the organization is using a limited (Zulip Cloud Free) plan.
	ZulipPlanIsNotLimited bool `json:"zulip_plan_is_not_limited,omitempty"`
	// Text to use when displaying UI for wide organization logos, a feature that is currently not available on the Zulip Cloud Free plan.  Useful only for clients supporting administrative UI for uploading a new wide organization logo to brand the organization.
	UpgradeTextForWideOrganizationLogo string `json:"upgrade_text_for_wide_organization_logo,omitempty"`
	// Dictionary where each entry describes a default external account type that can be configured with Zulip's [custom profile fields feature](zulip.com/help/custom-profile-fields.  **Changes**: New in Zulip 2.1.0.
	DefaultExternalAccounts *map[string]RealmDefaultExternalAccounts `json:"realm_default_external_accounts,omitempty"`
	// The base URL to be used to create Jitsi video calls. Equals `realm_jitsi_server_url || server_jitsi_server_url`.  **Changes**: Deprecated in Zulip 8.0 (feature level 212) and will eventually be removed. Previously, the Jitsi server to use was not configurable on a per-realm basis, and this field contained the server's configured Jitsi server. (Which is now provided as `server_jitsi_server_url`). Clients supporting older versions should fall back to this field when creating calls: using `realm_jitsi_server_url || server_jitsi_server_url` with newer servers and using `jitsi_server_url` with servers below feature level 212.
	// Deprecated
	JitsiServerUrlDeprecated string `json:"jitsi_server_url,omitempty"`
	// Whether this Zulip server is a development environment. Used to control certain features or UI (such as error popups) that should only apply when connected to a Zulip development environment.
	DevelopmentEnvironment bool `json:"development_environment,omitempty"`
	// A timestamp indicating when the process hosting this event queue was started. Clients will likely only find this value useful for inclusion in detailed error reports.
	ServerGeneration time.Time `json:"server_generation,omitempty"`
	// This Zulip server's configured minimum required length for passwords. Necessary for password change UI to show whether the password will be accepted.
	PasswordMinLength int64 `json:"password_min_length,omitempty"`
	// This Zulip server's configured maximum length for passwords. Necessary for password change UI to show whether the password will be accepted.  **Changes**: New in Zulip 10.0 (feature level 338).
	PasswordMaxLength int64 `json:"password_max_length,omitempty"`
	// This Zulip server's configured minimum `zxcvbn` minimum guesses. Necessary for password change UI to show whether the password will be accepted.
	PasswordMinGuesses int64 `json:"password_min_guesses,omitempty"`
	// Dictionary where each entry describes a valid rating that is configured on this server and could be selected by an organization administrator.  Useful for administrative settings UI that allows changing the allowed rating of GIFs.
	GiphyRatingOptions map[string]GiphyRatingOptionsValue `json:"giphy_rating_options,omitempty"`
	// The maximum file size that can be uploaded to this Zulip organization.
	MaxFileUploadSizeMib int64 `json:"max_file_upload_size_mib,omitempty"`
	// The maximum avatar size that can be uploaded to this Zulip server.
	MaxAvatarFileSizeMib int64 `json:"max_avatar_file_size_mib,omitempty"`
	// Whether the server is configured with support for inline image previews. Clients containing administrative UI for changing `realm_inline_image_preview` should consult this field before offering that feature.
	ServerInlineImagePreview bool `json:"server_inline_image_preview,omitempty"`
	// Whether the server is configured with support for inline URL previews. Clients containing administrative UI for changing `realm_inline_url_embed_preview` should consult this field before offering that feature.
	ServerInlineUrlEmbedPreview bool `json:"server_inline_url_embed_preview,omitempty"`
	// A list describing the image formats that uploaded images will be thumbnailed into. Any image with a source starting with `/user_uploads/thumbnail/` can have its last path component replaced with any of the names contained in this list, to obtain the desired thumbnail size.  **Changes**: New in Zulip 9.0 (feature level 273).
	ServerThumbnailFormats []ServerThumbnailFormat `json:"server_thumbnail_formats,omitempty"`
	// Whether the server allows avatar changes. Similar to `realm_avatar_changes_disabled` but based on the `AVATAR_CHANGES_DISABLED` Zulip server level setting.
	ServerAvatarChangesDisabled bool `json:"server_avatar_changes_disabled,omitempty"`
	// Whether the server allows name changes. Similar to `realm_name_changes_disabled` but based on the `NAME_CHANGES_DISABLED` Zulip server level setting.
	ServerNameChangesDisabled bool `json:"server_name_changes_disabled,omitempty"`
	// Whether the server is running an old version based on the Zulip [server release lifecycle](https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#upgrade-nag), such that the web app will display to the current user a prominent warning.  **Changes**: New in Zulip 5.0 (feature level 74).
	ServerNeedsUpgrade bool `json:"server_needs_upgrade,omitempty"`
	// The value of the `WEB_PUBLIC_STREAMS_ENABLED` Zulip server level setting. A server that has disabled this setting intends to not offer [web public channels](zulip.com/help/public-access-option) to realms it hosts. (Zulip Cloud defaults to `true`; self-hosted servers default to `false`).  Clients should use this to determine whether to offer UI for the realm-level setting for enabling web-public channels (`realm_enable_spectator_access`).  **Changes**: New in Zulip 5.0 (feature level 110).
	ServerWebPublicStreamsEnabled bool `json:"server_web_public_streams_enabled,omitempty"`
	// The URL to a JSON file that describes which emoji names map to which emoji codes, for all Unicode emoji this Zulip server accepts.  The data at the given URL is a JSON object with one property, `code_to_names`. The value of that property is a JSON object where each key is an [emoji code](zulip.com/api/add-reaction#parameter-emoji_code for an available Unicode emoji, and each value is the corresponding [emoji names](zulip.com/api/add-reaction#parameter-emoji_name for this emoji, with the canonical name for the emoji always appearing first.  The HTTP response at that URL will have appropriate HTTP caching headers, such any HTTP implementation should get a cached version if emoji haven't changed since the last request.  **Changes**: New in Zulip 6.0 (feature level 140).
	ServerEmojiDataUrl string `json:"server_emoji_data_url,omitempty"`
	// The URL of the Jitsi server that the Zulip server is configured to use by default; the organization-level setting `realm_jitsi_server_url` takes precedence over this setting when both are set.  **Changes**: New in Zulip 8.0 (feature level 212). Previously, this value was available as the now-deprecated `jitsi_server_url`.
	ServerJitsiServerUrl *string `json:"server_jitsi_server_url,omitempty"`
	// Present if `realm` is present in `fetch_event_types`  Whether topic summarization is enabled in the server or not depending upon whether `TOPIC_SUMMARIZATION_MODEL` is set or not.  **Changes**: New in Zulip 10.0 (feature level 350).
	ServerCanSummarizeTopics bool `json:"server_can_summarize_topics,omitempty"`
	// Recommended client-side HTTP request timeout for [`GET /events`](zulip.com/api/get-events) calls. This is guaranteed to be somewhat greater than the heartbeat frequency. It is important that clients respect this parameter, so that increases in the heartbeat frequency do not break clients.  **Changes**: New in Zulip 5.0 (feature level 74). Previously, this was hardcoded to 90 seconds, and clients should use that as a fallback value when interacting with servers where this field is not present.
	EventQueueLongpollTimeoutSeconds int64        `json:"event_queue_longpoll_timeout_seconds,omitempty"`
	Billing                          RealmBilling `json:"realm_billing,omitempty"`
	// The Id of the private channel to which messages flagged by users for moderation are sent. Moderators can use this channel to review and act on reported content.  Will be `-1` if moderation requests are disabled.  Clients should check whether moderation requests are disabled to determine whether to present a \"report message\" feature in their UI within a given organization.  **Changes**: New in Zulip 10.0 (feature level 331). Previously, no \"report message\" feature existed in Zulip.
	ModerationRequestChannelId int64 `json:"realm_moderation_request_channel_id,omitempty"`
	// The Id of the channel to which automated messages announcing the [creation of new channels][new-channel-announce] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-channel-announce]: /help/configure-automated-notices#new-channel-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed 'realm_notifications_stream_id' to `realm_new_stream_announcements_stream_id`.
	NewStreamAnnouncementsStreamId int64 `json:"realm_new_stream_announcements_stream_id,omitempty"`
	// The Id of the channel to which automated messages announcing that [new users have joined the organization][new-user-announce] are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  [new-user-announce]: /help/configure-automated-notices#new-user-announcements  **Changes**: In Zulip 9.0 (feature level 241), renamed 'realm_signup_notifications_stream_id' to `realm_signup_announcements_stream_id`.
	SignupAnnouncementsStreamId int64 `json:"realm_signup_announcements_stream_id,omitempty"`
	// The Id of the channel to which automated messages announcing new features or other end-user updates about the Zulip software are sent.  Will be `-1` if such automated messages are disabled.  Since these automated messages are sent by the server, this field is primarily relevant to clients containing UI for changing it.  **Changes**: New in Zulip 9.0 (feature level 242).
	ZulipUpdateAnnouncementsStreamId int64 `json:"realm_zulip_update_announcements_stream_id,omitempty"`
	// Clients declaring the `empty_topic_name` client capability should use the value of `realm_empty_topic_display_name` to determine how to display the empty string topic.  Clients not declaring the `empty_topic_name` client capability receive `realm_empty_topic_display_name` value as the topic name replacing empty string.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic name.
	EmptyTopicDisplayName string `json:"realm_empty_topic_display_name,omitempty"`

	// RealmUserSettingsDefaults Present if `realm_user_settings_defaults` is present in `fetch_event_types`.  A dictionary containing the default values of settings for new users.  **Changes**: New in Zulip 5.0 (feature level 95).
	UserSettingsDefaults *RealmUserSettingsDefaults `json:"realm_user_settings_defaults,omitempty"`
}

// VideoChatProviderInfo `{provider_name}`: Dictionary containing the details of the video call provider with the name of the chat provider as the key.
type VideoChatProviderInfo struct {
	// The name of the video call provider.
	Name string `json:"name,omitempty"`
	// The Id of the video call provider.
	Id int64 `json:"id,omitempty"`
}

// RealmAuthenticationMethod Dictionary describing the properties of an authentication method for the organization - its enabled status and availability for use by the organization.
type RealmAuthenticationMethod struct {
	// Boolean describing whether the authentication method (i.e. its key) is enabled in this organization.
	Enabled *bool `json:"enabled,omitempty"`
	// Boolean describing whether the authentication method is available for use. If false, the organization is not eligible to enable the authentication method.
	Available *bool `json:"available,omitempty"`
	// Reason why the authentication method is unavailable. This field is optional and is only present when 'available' is false.
	UnavailableReason *string `json:"unavailable_reason,omitempty"`
}

// CustomProfileFieldType `{FIELD_TYPE}`: Dictionary which contains the details of the field type with the field type as the name of the property itself. The current supported field types are as follows:  - `SHORT_TEXT` - `LONG_TEXT` - `DATE` for date-based fields. - `SELECT` for a list of options. - `URL` for links. - `EXTERNAL_ACCOUNT` for external accounts. - `USER` for selecting a user for the field. - `PRONOUNS` for a short text field with convenient typeahead for one's preferred pronouns.  **Changes**: `PRONOUNS` type added in Zulip 6.0 (feature level 151).
type CustomProfileFieldType struct {
	// The Id of the custom profile field type.
	Id int64 `json:"id,omitempty"`
	// The name of the custom profile field type.
	Name string `json:"name,omitempty"`
}

// ServerThumbnailFormat struct for ServerThumbnailFormat
type ServerThumbnailFormat struct {
	// The file path component of the thumbnail format.
	Name string `json:"name,omitempty"`
	// The maximum width of this format.
	MaxWidth int64 `json:"max_width,omitempty"`
	// The maximum height of this format.
	MaxHeight int64 `json:"max_height,omitempty"`
	// The extension of this format.
	Format string `json:"format,omitempty"`
	// If this file format is animated. These formats are only generated for uploaded imates which themselves are animated.
	Animated bool `json:"animated,omitempty"`
}

type RealmUserSettingsDefaults struct {
	// Whether time should be [displayed in 24-hour notation](zulip.com/help/change-the-time-format.  A `null` value indicates that the client should use the default time format for the user's locale.  **Changes**: Prior to Zulip 11.0 (feature level 408), `null` was not a valid value for this setting. Note that it was not possible to actually set the time format to `null` at this feature level.  New in Zulip 5.0 (feature level 99). This value was previously available as `realm_default_twenty_four_hour_time` in the top-level response object (only when `realm` was present in `fetch_event_types`).
	TwentyFourHourTime *bool `json:"twenty_four_hour_time,omitempty"`
	// Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \"Always\" setting when marking messages as read.
	WebMarkReadOnScrollPolicy MarkReadOnScrollPolicy `json:"web_mark_read_on_scroll_policy,omitempty"`
	// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \"Top unread topic in channel\" is new in Zulip 11.0 (feature level 401).  In Zulip 11.0 (feature level 383), we added a new option \"List of topics\" to this setting.  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \"Channel feed\" behavior.
	WebChannelDefaultView ChannelDefaultView `json:"web_channel_default_view,omitempty"`
	// Whether clients should display the [number of starred messages](zulip.com/help/star-a-message#display-the-number-of-starred-messages.
	StarredMessageCounts *bool `json:"starred_message_counts,omitempty"`
	// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
	ReceivesTypingNotifications *bool `json:"receives_typing_notifications,omitempty"`
	// Whether the user should be shown an alert, offering to update their [profile time zone](zulip.com/help/change-your-timezone, when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).
	WebSuggestUpdateTimezone *bool `json:"web_suggest_update_timezone,omitempty"`
	// Whether to use the [maximum available screen width](zulip.com/help/enable-full-width-display) for the web app's center panel (message feed, recent conversations) on wide screens.
	FluidLayoutWidth *bool `json:"fluid_layout_width,omitempty"`
	// This setting is reserved for use to control variations in Zulip's design to help visually impaired users.
	HighContrastMode *bool `json:"high_contrast_mode,omitempty"`
	// User-configured primary `font-size` for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.
	WebFontSizePx *int32 `json:"web_font_size_px,omitempty"`
	// User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.
	WebLineHeightPercent *int32 `json:"web_line_height_percent,omitempty"`
	// Controls which [color theme](zulip.com/help/dark-theme) to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard `prefers-color-scheme` media query.
	ColorScheme *int32 `json:"color_scheme,omitempty"`
	// Whether to [translate emoticons to emoji](zulip.com/help/configure-emoticon-translations) in messages the user sends.
	TranslateEmoticons *bool `json:"translate_emoticons,omitempty"`
	// Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).
	DisplayEmojiReactionUsers *bool `json:"display_emoji_reaction_users,omitempty"`
	// What [default language](zulip.com/help/change-your-language) to use for the account.  This controls both the Zulip UI as well as email notifications sent to the user.  The value needs to be a standard language code that the Zulip server has translation data for; for example, `\"en\"` for English or `\"de\"` for German.
	DefaultLanguage *string `json:"default_language,omitempty"`
	// The [home view](zulip.com/help/configure-home-view) used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.  - \"recent_topics\" - Recent conversations view - \"inbox\" - Inbox view - \"all_messages\" - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
	WebHomeView *string `json:"web_home_view,omitempty"`
	// Whether the escape key navigates to the [configured home view](zulip.com/help/configure-home-view.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).
	WebEscapeNavigatesToHomeView *bool `json:"web_escape_navigates_to_home_view,omitempty"`
	// Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.
	LeftSideUserlist *bool `json:"left_side_userlist,omitempty"`
	// The user's configured [emoji set](zulip.com/help/emoji-and-emoticons#use-emoticons, used to display emoji to the user everywhere they appear in the UI.  - \"google\" - Google modern - \"google-blob\" - Google classic - \"twitter\" - Twitter - \"text\" - Plain text
	Emojiset *string `json:"emojiset,omitempty"`
	// Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels) in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never
	DemoteInactiveStreams *int32 `json:"demote_inactive_streams,omitempty"`
	// The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).
	UserListStyle *int32 `json:"user_list_style,omitempty"`
	// Controls how animated images should be played in the message feed in the web/desktop application.  - \"always\" - Always play the animated images in the message feed. - \"on_hover\" - Play the animated images on hover over them in the message feed. - \"never\" - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).
	WebAnimateImagePreviews *string `json:"web_animate_image_previews,omitempty"`
	// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).
	WebStreamUnreadsCountDisplayPolicy *int32 `json:"web_stream_unreads_count_display_policy,omitempty"`
	// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).
	HideAiFeatures *bool `json:"hide_ai_features,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).
	WebLeftSidebarShowChannelFolders *bool `json:"web_left_sidebar_show_channel_folders,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).
	WebLeftSidebarUnreadsCountSummary *bool `json:"web_left_sidebar_unreads_count_summary,omitempty"`
	// Enable visual desktop notifications for channel messages.
	EnableStreamDesktopNotifications *bool `json:"enable_stream_desktop_notifications,omitempty"`
	// Enable email notifications for channel messages.
	EnableStreamEmailNotifications *bool `json:"enable_stream_email_notifications,omitempty"`
	// Enable mobile notifications for channel messages.
	EnableStreamPushNotifications *bool `json:"enable_stream_push_notifications,omitempty"`
	// Enable audible desktop notifications for channel messages.
	EnableStreamAudibleNotifications *bool `json:"enable_stream_audible_notifications,omitempty"`
	// Notification sound name.
	NotificationSound *string `json:"notification_sound,omitempty"`
	// Enable visual desktop notifications for direct messages and @-mentions.
	EnableDesktopNotifications *bool `json:"enable_desktop_notifications,omitempty"`
	// Enable audible desktop notifications for direct messages and @-mentions.
	EnableSounds *bool `json:"enable_sounds,omitempty"`
	// Enable email notifications for direct messages and @-mentions received when the user is offline.
	EnableOfflineEmailNotifications *bool `json:"enable_offline_email_notifications,omitempty"`
	// Enable mobile notification for direct messages and @-mentions received when the user is offline.
	EnableOfflinePushNotifications *bool `json:"enable_offline_push_notifications,omitempty"`
	// Enable mobile notification for direct messages and @-mentions received when the user is online.
	EnableOnlinePushNotifications *bool `json:"enable_online_push_notifications,omitempty"`
	// Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicDesktopNotifications *bool `json:"enable_followed_topic_desktop_notifications,omitempty"`
	// Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicEmailNotifications *bool `json:"enable_followed_topic_email_notifications,omitempty"`
	// Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicPushNotifications *bool `json:"enable_followed_topic_push_notifications,omitempty"`
	// Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicAudibleNotifications *bool `json:"enable_followed_topic_audible_notifications,omitempty"`
	// Enable digest emails when the user is away.
	EnableDigestEmails *bool `json:"enable_digest_emails,omitempty"`
	// Enable marketing emails. Has no function outside Zulip Cloud.
	EnableMarketingEmails *bool `json:"enable_marketing_emails,omitempty"`
	// Enable email notifications for new logins to account.
	EnableLoginEmails *bool `json:"enable_login_emails,omitempty"`
	// Include the message's content in email notifications for new messages.
	MessageContentInEmailNotifications *bool `json:"message_content_in_email_notifications,omitempty"`
	// Include content of direct messages in desktop notifications.
	PmContentInDesktopNotifications *bool `json:"pm_content_in_desktop_notifications,omitempty"`
	// Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.
	WildcardMentionsNotify *bool `json:"wildcard_mentions_notify,omitempty"`
	// Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).
	EnableFollowedTopicWildcardMentionsNotify *bool `json:"enable_followed_topic_wildcard_mentions_notify,omitempty"`
	// Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
	DesktopIconCountDisplay *int32 `json:"desktop_icon_count_display,omitempty"`
	// Whether to [include organization name in subject of message notification emails](zulip.com/help/email-notifications#include-organization-name-in-subject-line.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
	RealmNameInEmailNotificationsPolicy *int32 `json:"realm_name_in_email_notifications_policy,omitempty"`
	// Which [topics to follow automatically](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
	AutomaticallyFollowTopicsPolicy *int32 `json:"automatically_follow_topics_policy,omitempty"`
	// Which [topics to unmute automatically in muted channels](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
	AutomaticallyUnmuteTopicsInMutedStreamsPolicy *int32 `json:"automatically_unmute_topics_in_muted_streams_policy,omitempty"`
	// Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).
	AutomaticallyFollowTopicsWhereMentioned *bool `json:"automatically_follow_topics_where_mentioned,omitempty"`
	// Controls whether the resolved-topic notices are marked as read.  - \"always\" - Always mark resolved-topic notices as read. - \"except_followed\" - Mark resolved-topic notices as read in topics not followed by the user. - \"never\" - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).
	ResolvedTopicNoticeAutoReadPolicy *string `json:"resolved_topic_notice_auto_read_policy,omitempty"`
	// Display the presence status to other users when online.
	PresenceEnabled *bool `json:"presence_enabled,omitempty"`
	// Whether the user setting for [sending on pressing Enter](zulip.com/help/configure-send-message-keys) in the compose box is enabled.
	EnterSends *bool `json:"enter_sends,omitempty"`
	// A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.
	EnableDraftsSynchronization *bool `json:"enable_drafts_synchronization,omitempty"`
	// The duration (in seconds) for which the server should wait to batch email notifications before sending them.
	EmailNotificationsBatchingPeriodSeconds *int32 `json:"email_notifications_batching_period_seconds,omitempty"`
	// Array containing the names of the notification sound options supported by this Zulip server. Only relevant to support UI for configuring notification sounds.
	AvailableNotificationSounds []string `json:"available_notification_sounds,omitempty"`
	// Array of dictionaries where each dictionary describes an emoji set supported by this version of the Zulip server.  Only relevant to clients with configuration UI for choosing an emoji set; the currently selected emoji set is available in the `emojiset` key.  See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.
	EmojisetChoices []UserSettingsEmojisetChoice `json:"emojiset_choices,omitempty"`
	// Whether [typing notifications](zulip.com/help/typing-notifications) be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).
	SendPrivateTypingNotifications *bool `json:"send_private_typing_notifications,omitempty"`
	// Whether [typing notifications](zulip.com/help/typing-notifications) be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).
	SendStreamTypingNotifications *bool `json:"send_stream_typing_notifications,omitempty"`
	// Whether other users are allowed to see whether you've read messages.  **Changes**: New in Zulip 5.0 (feature level 105).
	SendReadReceipts *bool `json:"send_read_receipts,omitempty"`
	// Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).
	AllowPrivateDataExport *bool `json:"allow_private_data_export,omitempty"`
	// The [policy][permission-level] for [which other users][help-email-visibility] in this organization can see the user's real email address.  - 1 = Everyone - 2 = Members only - 3 = Administrators only - 4 = Nobody - 5 = Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility
	EmailAddressVisibility *int32 `json:"email_address_visibility,omitempty"`
	// Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
	WebNavigateToSentMessage *bool `json:"web_navigate_to_sent_message,omitempty"`
}
