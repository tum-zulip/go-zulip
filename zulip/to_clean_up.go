package zulip

import "time"

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
	Channels []int32 `json:"streams,omitempty"`
}

// UnreadMsgs Present if `message` and `update_message_flags` are both present in `event_types`.  A set of data structures describing the conversations containing the 50000 most recent unread messages the user has received. This will usually contain every unread message the user has received, but clients should support users with even more unread messages (and not hardcode the number 50000).
type UnreadMsgs struct {
	// The total number of unread messages to display. This includes one-on-one and group direct messages, as well as channel messages that are not [muted](zulip.com/help/mute-a-topic.  **Changes**: Before Zulip 8.0 (feature level 213), the unmute and follow topic features were not handled correctly in calculating this field.
	Count *int32 `json:"count,omitempty"`
	// An array of objects where each object contains details of unread one-on-one direct messages with a specific user.  Note that it is possible for a message that the current user sent to the specified user to be marked as unread and thus appear here.
	Pms []UnreadMsgsPms `json:"pms,omitempty"`
	// An array of dictionaries where each dictionary contains details of all unread messages of a single subscribed channel. This includes muted channels and muted topics, even though those messages are excluded from `count`.  **Changes**: Prior to Zulip 5.0 (feature level 90), these objects included a `sender_ids` property, which listed the set of Ids of users who had sent the unread messages.
	Channels []UnreadMsgsChannels `json:"streams,omitempty"`
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

// UnreadMsgsChannels struct for UnreadMsgsChannels
type UnreadMsgsChannels struct {
	// The topic under which the messages were sent.  Note that the empty string topic may have been rewritten by the server to the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue) response depending on the value of the `empty_topic_name` [client capability][client-capabilities].  **Changes**: The `empty_topic_name` client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	Topic *string `json:"topic,omitempty"`
	// The Id of the channel to which the messages were sent.
	ChannelId *int64 `json:"stream_id,omitempty"`
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
	ChannelId *int64 `json:"stream_id,omitempty"`
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
