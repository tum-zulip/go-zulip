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
	MutedTopics [][]interface{} `json:"muted_topics,omitempty"`
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
	RealmFilters [][]interface{} `json:"realm_filters,omitempty"`
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
	Channels []Channel `json:"streams,omitempty"`
	// Present if `default_streams` is present in `fetch_event_types`.  An array of Ids of all the [default channels](zulip.com/help/set-default-streams-for-new-users) in the organization.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.
	RealmDefaultChannels []int64 `json:"realm_default_streams,omitempty"`
	// Present if `default_stream_groups` is present in `fetch_event_types`.  An array of dictionaries where each dictionary contains details about a single default channel group configured for this Zulip organization.  Default channel groups are an experimental feature.
	RealmDefaultChannelGroups []DefaultChannelGroup `json:"realm_default_stream_groups,omitempty"`
	// Present if `stop_words` is present in `fetch_event_types`.  An array containing the stop words used by the Zulip server's full-text search implementation. Useful for showing helpful error messages when a search returns limited results because a stop word in the query was ignored.
	StopWords []string `json:"stop_words,omitempty"`
	// Present if `user_status` is present in `fetch_event_types`.  A dictionary which contains the [status](zulip.com/help/status-and-availability) of all users in the Zulip organization who have set a status.  **Changes**: The emoji parameters are new in Zulip 5.0 (feature level 86). Previously, Zulip did not support emoji associated with statuses.
	UserStatus   map[string]UserStatus `json:"user_status,omitempty"`
	UserSettings *UserSettings         `json:"user_settings,omitempty"`
	// Present if `user_topic` is present in `fetch_event_types`.  **Changes**: New in Zulip 6.0 (feature level 134), deprecating and replacing the previous `muted_topics` structure.
	UserTopics []UserTopic `json:"user_topics,omitempty"`
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
	Channel map[string]GroupPermissionSetting `json:"stream,omitempty"`
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
	DemoteInactiveChannels int64 `json:"demote_inactive_streams,omitempty"`
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

	Avatar

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
	CrossRealmBots []User `json:"cross_realm_bots,omitempty"`
}

type Avatar struct {
	// The URL of the new avatar for the user.
	AvatarUrl string `json:"avatar_url,omitempty"`
	// The new avatar data source type for the user.  Value values are `G` (gravatar) and `U` (uploaded by user).
	AvatarSource AvatarSource `json:"avatar_source,omitempty"`
	// The new medium-size avatar URL for user.
	AvatarUrlMedium string `json:"avatar_url_medium,omitempty"`
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
	EnableChannelDesktopNotifications bool `json:"enable_stream_desktop_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableChannelEmailNotifications bool `json:"enable_stream_email_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableChannelPushNotifications bool `json:"enable_stream_push_notifications,omitempty"`
	// The current value of this global notification setting for the user. See [PATCH /settings](zulip.com/api/update-settings) for details on the meaning of this setting.  **Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and access the `user_settings` object instead.  [capabilities]: /api/register-queue#parameter-client_capabilities
	// Deprecated
	EnableChannelAudibleNotifications bool `json:"enable_stream_audible_notifications,omitempty"`
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
	DemoteInactiveChannels *int32 `json:"demote_inactive_streams,omitempty"`
	// The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).
	UserListStyle *int32 `json:"user_list_style,omitempty"`
	// Controls how animated images should be played in the message feed in the web/desktop application.  - \"always\" - Always play the animated images in the message feed. - \"on_hover\" - Play the animated images on hover over them in the message feed. - \"never\" - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275).
	WebAnimateImagePreviews *string `json:"web_animate_image_previews,omitempty"`
	// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).
	WebChannelUnreadsCountDisplayPolicy *int32 `json:"web_stream_unreads_count_display_policy,omitempty"`
	// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).
	HideAiFeatures *bool `json:"hide_ai_features,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).
	WebLeftSidebarShowChannelFolders *bool `json:"web_left_sidebar_show_channel_folders,omitempty"`
	// Determines whether the web/desktop application's left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).
	WebLeftSidebarUnreadsCountSummary *bool `json:"web_left_sidebar_unreads_count_summary,omitempty"`
	// Enable visual desktop notifications for channel messages.
	EnableChannelDesktopNotifications *bool `json:"enable_stream_desktop_notifications,omitempty"`
	// Enable email notifications for channel messages.
	EnableChannelEmailNotifications *bool `json:"enable_stream_email_notifications,omitempty"`
	// Enable mobile notifications for channel messages.
	EnableChannelPushNotifications *bool `json:"enable_stream_push_notifications,omitempty"`
	// Enable audible desktop notifications for channel messages.
	EnableChannelAudibleNotifications *bool `json:"enable_stream_audible_notifications,omitempty"`
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
	AutomaticallyUnmuteTopicsInMutedChannelsPolicy *int32 `json:"automatically_unmute_topics_in_muted_streams_policy,omitempty"`
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
	SendChannelTypingNotifications *bool `json:"send_stream_typing_notifications,omitempty"`
	// Whether other users are allowed to see whether you've read messages.  **Changes**: New in Zulip 5.0 (feature level 105).
	SendReadReceipts *bool `json:"send_read_receipts,omitempty"`
	// Whether organization administrators are allowed to export your private data.  **Changes**: New in Zulip 10.0 (feature level 293).
	AllowPrivateDataExport *bool `json:"allow_private_data_export,omitempty"`
	// The [policy][permission-level] for [which other users][help-email-visibility] in this organization can see the user's real email address.  - 1 = Everyone - 2 = Members only - 3 = Administrators only - 4 = Nobody - 5 = Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility
	EmailAddressVisibility *int32 `json:"email_address_visibility,omitempty"`
	// Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
	WebNavigateToSentMessage *bool `json:"web_navigate_to_sent_message,omitempty"`
}

// UserTopics Object describing the user's configuration for a given topic.
type UserTopic struct {
	// The Id of the channel to which the topic belongs.
	ChannelId int64 `json:"stream_id,omitempty"`
	// The name of the topic.  For clients that don't support the `empty_topic_name` [client capability][client-capabilities], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	TopicName string `json:"topic_name,omitempty"`
	// An integer UNIX timestamp representing when the user-topic relationship was last changed.
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// An integer indicating the user's visibility preferences for the topic, such as whether the topic is muted.  - 0 = None. Used to indicate that the user no   longer has a special visibility policy for this topic. - 1 = Muted. Used to record [muted topics](zulip.com/help/mute-a-topic. - 2 = Unmuted. Used to record unmuted topics. - 3 = Followed. Used to record [followed topics](zulip.com/help/follow-a-topic.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  In Zulip 7.0 (feature level 170), added unmuted as a visibility policy option.
	VisibilityPolicy VisibilityPolicy `json:"visibility_policy,omitempty"`
}

// DefaultChannelGroup Dictionary containing details of a default channel group.
type DefaultChannelGroup struct {
	// Name of the default channel group.
	Name string `json:"name,omitempty"`
	// Description of the default channel group.
	Description string `json:"description,omitempty"`
	// The Id of the default channel group.
	Id int64 `json:"id,omitempty"`
	// An array of Ids of all the channels in the default stream group.  **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single stream in the default stream group.
	Channels []int64 `json:"streams,omitempty"`
}

// UnreadMsgs Present if `message` and `update_message_flags` are both present in `event_types`.  A set of data structures describing the conversations containing the 50000 most recent unread messages the user has received. This will usually contain every unread message the user has received, but clients should support users with even more unread messages (and not hardcode the number 50000).
type UnreadMsgs struct {
	// The total number of unread messages to display. This includes one-on-one and group direct messages, as well as channel messages that are not [muted](zulip.com/help/mute-a-topic.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.
	Count int64 `json:"count,omitempty"`
	// An array of objects where each object contains details of unread one-on-one direct messages with a specific user.  Note that it is possible for a message that the current user sent to the specified user to be marked as unread and thus appear here.
	Pms []UnreadMsgsPms `json:"pms,omitempty"`
	// An array of dictionaries where each dictionary contains details of all unread messages of a single subscribed channel. This includes muted channels and muted topics, even though those messages are excluded from `count`.  **Changes**: Prior to Zulip 5.0 (feature level 90), these objects included a `sender_ids` property, which listed the set of Ids of users who had sent the unread messages.
	Channels []UnreadMsgsChannels `json:"streams,omitempty"`
	// An array of objects where each object contains details of unread group direct messages with a specific group of users.
	Huddles []UnreadMsgsHuddles `json:"huddles,omitempty"`
	// Array containing the Ids of all unread messages in which the user was mentioned directly, and unread [non-muted](zulip.com/help/mute-a-topic) messages in which the user was mentioned through a wildcard.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.
	Mentions []int64 `json:"mentions,omitempty"`
	// Whether this data set was truncated because the user has too many unread messages. When truncation occurs, only the most recent `MAX_UNREAD_MESSAGES` (currently 50000) messages will be considered when forming this response. When `true`, we recommend that clients display a warning, as they are likely to produce erroneous results until reloaded with the user having fewer than `MAX_UNREAD_MESSAGES` unread messages.  **Changes**: New in Zulip 4.0 (feature level 44).
	OldUnreadsMissing *bool `json:"old_unreads_missing,omitempty"`
}

// UnreadMsgsPms struct for UnreadMsgsPms
type UnreadMsgsPms struct {
	// The user Id of the other participant in this one-on-one direct message conversation. Will be the current user's Id for messages that they sent in a one-on-one direct message conversation with themself.  **Changes**: New in Zulip 5.0 (feature level 119), replacing the less clearly named `sender_id` field.
	OtherUserId int64 `json:"other_user_id,omitempty"`
	// Old name for the `other_user_id` field. Clients should access this field in Zulip server versions that do not yet support `other_user_id`.  **Changes**: Deprecated in Zulip 5.0 (feature level 119). We expect to provide a next version of the full `unread_msgs` API before removing this legacy name.
	// Deprecated
	SenderId int64 `json:"sender_id,omitempty"`
	// The message Ids of the recent unread direct messages sent by either user in this one-on-one direct message conversation, sorted in ascending order.
	UnreadMessageIds []int64 `json:"unread_message_ids,omitempty"`
}

// UnreadMsgsChannels struct for UnreadMsgsChannels
type UnreadMsgsChannels struct {
	// The topic under which the messages were sent.  Note that the empty string topic may have been rewritten by the server to the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue) response depending on the value of the `empty_topic_name` [client capability][client-capabilities].  **Changes**: The `empty_topic_name` client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	Topic string `json:"topic,omitempty"`
	// The Id of the channel to which the messages were sent.
	ChannelId int64 `json:"stream_id,omitempty"`
	// The message Ids of the recent unread messages sent in this channel, sorted in ascending order.
	UnreadMessageIds []int64 `json:"unread_message_ids,omitempty"`
}

// UnreadMsgsHuddles struct for UnreadMsgsHuddles
type UnreadMsgsHuddles struct {
	// A string containing the Ids of all users in the group direct message conversation, including the current user, separated by commas and sorted numerically; for example: `\"1,2,3\"`.
	UserIdsString string `json:"user_ids_string,omitempty"`
	// The message Ids of the recent unread messages which have been sent in this group direct message conversation, sorted in ascending order.
	UnreadMessageIds []int64 `json:"unread_message_ids,omitempty"`
}
