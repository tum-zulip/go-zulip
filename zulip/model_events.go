package zulip

import (
	"encoding/json"
	"time"
)

type Event interface {
	GetType() EventType
	GetId() int64
	HasOp() bool
	GetOp() EventOp
	GetOpOk() (*EventOp, bool)
}

type EventCommon struct {
	// The Id of the event. Events appear in increasing order but may not be consecutive.
	Id   int64     `json:"id,omitempty"`
	Type EventType `json:"type,omitempty"`
}

func (e EventCommon) GetType() EventType {
	return e.Type
}

func (e EventCommon) GetId() int64 {
	return e.Id
}

func (e EventCommon) HasOp() bool {
	return false
}

func (e EventCommon) GetOp() EventOp {
	return ""
}

func (e EventCommon) GetOpOk() (*EventOp, bool) {
	return nil, false
}

type EventCommonWithOp struct {
	EventCommon
	Op EventOp `json:"op,omitempty"`
}

func (e EventCommonWithOp) HasOp() bool {
	return true
}

func (e EventCommonWithOp) GetOp() EventOp {
	return e.Op
}

func (e EventCommonWithOp) GetOpOk() (*EventOp, bool) {
	return &e.Op, true
}

// HeartbeatEvent Event sent periodically to indicate that the event queue is still active.
type HeartbeatEvent struct{ EventCommon }

// InvitesChangedEvent A simple event sent when the set of invitations changes. This event is sent to organization administrators and the creator of the changed invitation; this tells clients they need to refetch data from `GET /invites` if they are displaying UI containing active invitations.
//
//	**Changes**: Before Zulip 8.0 (feature level 209), this event was only sent to organization administrators.
type InvitesChangedEvent struct{ EventCommon }

type AlertWordsEvent struct {
	EventCommon

	// An array of strings, where each string is an alert word (or phrase) configured by the user.
	AlertWords []string `json:"alert_words,omitempty"`
}

// MessageEvent Event type for messages.
//
//	**Changes**: In Zulip 3.1 (feature level 26), the `sender_short_name` field was removed from message objects.
type MessageEvent struct {
	EventCommon

	Message *MessagesEventData `json:"message,omitempty"`
	// The user's [message flags] for the message.  Clients should inspect the flags field rather than assuming that new messages are unread; [muted users], messages sent by the current user, and more subtle scenarios can result in a new message that the server has already marked as read for the user.
	//
	// **Changes**: In Zulip 8.0 (feature level 224), the `wildcard_mentioned` flag was deprecated in favor of the `stream_wildcard_mentioned` and `topic_wildcard_mentioned` flags. The `wildcard_mentioned` flag exists for backwards compatibility with older clients and equals `stream_wildcard_mentioned || topic_wildcard_mentioned`. Clients supporting older server versions should treat this field as a previous name for the `stream_wildcard_mentioned` flag as topic wildcard mentions were not available prior to this feature level.  [message flags]: https://zulip.com/api/update-message-flags#available-flags
	//
	// [muted users]: https://zulip.com/api/mute-user
	Flags []string `json:"flags,omitempty"`
}

// MessagesEvent struct for MessagesEvent
type MessagesEventData struct {
	AvatarUrl interface{} `json:"avatar_url,omitempty"`
	// A Zulip "client" string, describing what Zulip client sent the message.
	Client *string `json:"client,omitempty"`
	// The content/body of the message. When `apply_markdown` is set, it will be in HTML format.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	Content *string `json:"content,omitempty"`
	// The HTTP `content_type` for the message content. This will be `text/html` or `text/x-markdown`, depending on whether `apply_markdown` was set.  See the help center article on [message formatting] for details on Zulip-flavored Markdown.
	//
	// [message formatting]: https://zulip.com/help/format-your-message-using-markdown
	ContentType      *string           `json:"content_type,omitempty"`
	DisplayRecipient *DisplayRecipient `json:"display_recipient,omitempty"`
	// An array of objects, with each object documenting the changes in a previous edit made to the message, ordered chronologically from most recent to least recent edit.  Not present if the message has never been edited or moved, or if [viewing message edit history] is not allowed in the organization.  Every object will contain `user_id` and `timestamp`.  The other fields are optional, and will be present or not depending on whether the channel, topic, and/or message content were modified in the edit event. For example, if only the topic was edited, only `prev_topic` and `topic` will be present in addition to `user_id` and `timestamp`.
	//
	// **Changes**: In Zulip 10.0 (feature level 284), removed the `prev_rendered_content_version` field as it is an internal server implementation detail not used by any client.
	//
	// [viewing message edit history]: https://zulip.com/help/restrict-message-edit-history-access
	EditHistory []EditHistory `json:"edit_history,omitempty"`
	// The unique message Id. Messages should always be displayed sorted by Id.
	Id int64 `json:"id,omitempty"`
	// Whether the message is a [/me status message]
	//
	// [/me status message]: https://zulip.com/help/format-your-message-using-markdown#status-messages
	IsMeMessage bool `json:"is_me_message,omitempty"`
	// The UNIX timestamp for when the message's content was last edited, in UTC seconds.  Not present if the message's content has never been edited.  Clients should use this field, rather than parsing the `edit_history` array, to display an indicator that the message has been edited.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 365), this was the time when the message was last edited or moved.
	LastEditTimestamp *time.Time `json:"last_edit_timestamp,omitempty"`
	// The UNIX timestamp for when the message was last moved to a different channel or topic, in UTC seconds.  Not present if the message has never been moved, or if the only topic moves for the message are [resolving or unresolving] the message's topic.  Clients should use this field, rather than parsing the `edit_history` array, to display an indicator that the message has been moved.
	//
	// **Changes**: New in Zulip 10.0 (feature level 365). Previously, parsing the `edit_history` array was required in order to correctly display moved message indicators.
	//
	// [resolving or unresolving]: https://zulip.com/help/resolve-a-topic
	LastMovedTimestamp *time.Time `json:"last_moved_timestamp,omitempty"`
	// Data on any reactions to the message.
	Reactions []EmojiReaction `json:"reactions,omitempty"`
	// A unique Id for the set of users receiving the message (either a channel or group of users). Useful primarily for hashing.
	//
	// **Changes**: Before Zulip 10.0 (feature level 327), `recipient_id` was the same across all incoming 1:1 direct messages. Now, each incoming message uniquely shares a `recipient_id` with outgoing messages in the same conversation.
	RecipientId int64 `json:"recipient_id,omitempty"`
	// The Zulip API email address of the message's sender.
	SenderEmail *string `json:"sender_email,omitempty"`
	// The full name of the message's sender.
	SenderFullName *string `json:"sender_full_name,omitempty"`
	// The user Id of the message's sender.
	SenderId int64 `json:"sender_id,omitempty"`
	// A string identifier for the realm the sender is in. Unique only within the context of a given Zulip server.  E.g. on `example.zulip.com`, this will be `example`.
	SenderRealmStr string `json:"sender_realm_str,omitempty"`
	// Only present for channel messages; the Id of the channel.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// The `topic` of the message. Currently always `""` for direct messages, though this could change if Zulip adds support for topics in direct message conversations.  The field name is a legacy holdover from when topics were called "subjects" and will eventually change.  For clients that don't support the `empty_topic_name` [client capability], the empty string value is replaced with the value of `realm_empty_topic_display_name` found in the [POST /register] response, for channel messages.
	//
	// **Changes**: Before Zulip 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [POST /register]: https://zulip.com/api/register-queue
	Subject string `json:"subject,omitempty"`
	// Data used for certain experimental Zulip integrations.
	Submessages []Submessage `json:"submessages,omitempty"`
	// The UNIX timestamp for when the message was sent, in UTC seconds.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Data on any links to be included in the `topic` line (these are generated by [custom linkification filters] that match content in the message's topic.)  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called `subject_links`; clients are recommended to rename `subject_links` to `topic_links` if present for compatibility with older Zulip servers.
	//
	// [custom linkification filters]: https://zulip.com/help/add-a-custom-linkifier
	TopicLinks []TopicLink `json:"topic_links,omitempty"`
	// The type of the message: `"stream"` or `"private"`.
	Type *string `json:"type,omitempty"`
}

type messagesEventDataJSON struct {
	AvatarUrl          interface{}       `json:"avatar_url,omitempty"`
	Client             *string           `json:"client,omitempty"`
	Content            *string           `json:"content,omitempty"`
	ContentType        *string           `json:"content_type,omitempty"`
	DisplayRecipient   *DisplayRecipient `json:"display_recipient,omitempty"`
	EditHistory        []EditHistory     `json:"edit_history,omitempty"`
	Id                 int64             `json:"id,omitempty"`
	IsMeMessage        bool              `json:"is_me_message,omitempty"`
	LastEditTimestamp  *int64            `json:"last_edit_timestamp,omitempty"`
	LastMovedTimestamp *int64            `json:"last_moved_timestamp,omitempty"`
	Reactions          []EmojiReaction   `json:"reactions,omitempty"`
	RecipientId        int64             `json:"recipient_id,omitempty"`
	SenderEmail        *string           `json:"sender_email,omitempty"`
	SenderFullName     *string           `json:"sender_full_name,omitempty"`
	SenderId           int64             `json:"sender_id,omitempty"`
	SenderRealmStr     string            `json:"sender_realm_str,omitempty"`
	ChannelId          *int64            `json:"stream_id,omitempty"`
	Subject            string            `json:"subject,omitempty"`
	Submessages        []Submessage      `json:"submessages,omitempty"`
	Timestamp          int64             `json:"timestamp,omitempty"`
	TopicLinks         []TopicLink       `json:"topic_links,omitempty"`
	Type               *string           `json:"type,omitempty"`
}

func (m *MessagesEventData) MarshalJSON() ([]byte, error) {
	obj := messagesEventDataJSON{
		AvatarUrl:        m.AvatarUrl,
		Client:           m.Client,
		Content:          m.Content,
		ContentType:      m.ContentType,
		DisplayRecipient: m.DisplayRecipient,
		EditHistory:      m.EditHistory,
		Id:               m.Id,
		IsMeMessage:      m.IsMeMessage,
		Reactions:        m.Reactions,
		RecipientId:      m.RecipientId,
		SenderEmail:      m.SenderEmail,
		SenderFullName:   m.SenderFullName,
		SenderId:         m.SenderId,
		SenderRealmStr:   m.SenderRealmStr,
		ChannelId:        m.ChannelId,
		Subject:          m.Subject,
		Submessages:      m.Submessages,
		TopicLinks:       m.TopicLinks,
		Type:             m.Type,
		Timestamp:        m.Timestamp.Unix(),
	}
	if m.LastEditTimestamp != nil {
		t := m.LastEditTimestamp.Unix()
		obj.LastEditTimestamp = &t
	}
	if m.LastMovedTimestamp != nil {
		t := m.LastMovedTimestamp.Unix()
		obj.LastMovedTimestamp = &t
	}
	return json.Marshal(obj)
}

func (m *MessagesEventData) UnmarshalJSON(data []byte) error {
	var obj messagesEventDataJSON
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	m.AvatarUrl = obj.AvatarUrl
	m.Client = obj.Client
	m.Content = obj.Content
	m.ContentType = obj.ContentType
	m.DisplayRecipient = obj.DisplayRecipient
	m.EditHistory = obj.EditHistory
	m.Id = obj.Id
	m.IsMeMessage = obj.IsMeMessage
	m.Reactions = obj.Reactions
	m.RecipientId = obj.RecipientId
	m.SenderEmail = obj.SenderEmail
	m.SenderFullName = obj.SenderFullName
	m.SenderId = obj.SenderId
	m.SenderRealmStr = obj.SenderRealmStr
	m.ChannelId = obj.ChannelId
	m.Subject = obj.Subject
	m.Submessages = obj.Submessages
	m.TopicLinks = obj.TopicLinks
	m.Type = obj.Type
	m.Timestamp = time.Unix(obj.Timestamp, 0)

	if obj.LastEditTimestamp != nil {
		t := time.Unix(*obj.LastEditTimestamp, 0)
		m.LastEditTimestamp = &t
	}

	if obj.LastMovedTimestamp != nil {
		t := time.Unix(*obj.LastMovedTimestamp, 0)
		m.LastMovedTimestamp = &t
	}

	return nil
}

// UpdateDisplaySettingsEvent Event sent to clients that have requested the `update_display_settings` event type and did not include `user_settings_object` in their `client_capabilities` when registering the event queue.
//
//	**Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.
type UpdateDisplaySettingsEvent struct {
	EventCommon

	// The Zulip API email of the user.
	User string `json:"user,omitempty"`
	// Name of the changed display setting.
	SettingName string      `json:"setting_name,omitempty"`
	Setting     UpdateValue `json:"setting,omitempty"`
	// Present only if the setting to be changed is `default_language`. Contains the name of the new default language in English.
	LanguageName *string `json:"language_name,omitempty"`
}

// HasZoomTokenEvent Event sent to a user's clients when the user completes the OAuth flow for the [Zoom integration]. Clients need to know whether initiating Zoom OAuth is required before creating a Zoom call.
//
// [Zoom integration]: https://zulip.com/help/configure-call-provider
type HasZoomTokenEvent struct {
	EventCommon
	// A boolean specifying whether the user has zoom token or not.
	Value bool `json:"value,omitempty"`
}

// RealmUserAddEvent Event sent to all users in a Zulip organization when a new user joins or when a guest user gains access to a user. Processing this event is important to being able to display basic details on other users given only their Id.  If the current user is a guest whose access to a newly created user is limited by a `can_access_all_users_group` policy, and the event queue was registered with the `user_list_incomplete` client capability, then the event queue will not receive an event for such a new user. If a newly created user is inaccessible to the current user via such a policy, but the client lacks `user_list_incomplete` client capability, then this event will be delivered to the queue, with an "Unknown user" object with the usual format but placeholder data whose only variable content is the user Id.
//
//	**Changes**: Before Zulip 8.0 (feature level 232), the `user_list_incomplete` client capability did not exist, and so all clients whose access to a new user was prevented by `can_access_all_users_group` policy would receive a fake "Unknown user" event for such a user.  Starting with Zulip 8.0 (feature level 228), this event is also sent when a guest user gains access to a user.
type RealmUserAddEvent struct {
	EventCommonWithOp

	Person User `json:"person,omitempty"`
}

// RealmUserRemoveEvent Event sent to guest users when they lose access to a user.
//
//	**Changes**: As of Zulip 8.0 (feature level 228), this event is no longer deprecated.  In Zulip 8.0 (feature level 222), this event was deprecated and no longer sent to clients. Prior to this feature level, it was sent to all users in a Zulip organization when a user was deactivated.
type RealmUserRemoveEvent struct {
	EventCommonWithOp

	Person UserInfo `json:"person,omitempty"`
}

// UserInfo Object containing details of the deactivated user.
type UserInfo struct {
	// The Id of the deactivated user.
	UserId int64 `json:"user_id,omitempty"`
	// The full name of the user.  **Deprecated**: We expect to remove this field in the future.
	// Deprecated
	FullName *string `json:"full_name,omitempty"`
}

// PresenceEvent Event sent to all users in an organization when a user comes back online after being offline for a while.  In addition to handling these events, a client that wants to maintain presence data must poll the [main presence endpoint]. Most updates to presence data, refreshing the timestamps of users who are already online, do not appear in the event queue. This design is an optimization by allowing those updates to be batched up, because there is no urgency in the information that an already-online user is still online.  These events are provided because when a user transitions from offline to online, that is information the client may want to show promptly in the UI to avoid showing a confusing state (for example, if the newly-online user sends a message or otherwise demonstrates they're online).  If the client supports the `simplified_presence_events` [client capability], the `simplified_presence_events` client capability did not exist. Therefore, all events were in the legacy format, and did not include the `presences` field.  Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
//
//	**Changes**: Prior to Zulip 11.0 (feature level 419
//
// [main presence endpoint]: https://zulip.com/api/get-presence
// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities, these events will include the `presences` field, which provides the modified user's presence data in the modern format. Clients are strongly encouraged to implement this client capability, as legacy format support will be removed in a future release.  If the `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is set to `true`, then the event is only sent to users who can access the user who came back online.
type PresenceEvent struct {
	EventCommon
	// Only present for clients that support the `simplified_presence_events` [client capability] for the modified user(s). Clients should support updating multiple users in a single event.
	//
	// **Changes**: New in Zulip 11.0 (feature level 419).
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities.  A dictionary mapping user Ids to the presence data (modern format
	Presences map[string]ModernPresenceFormat `json:"presences,omitempty"`

	// todo: custom unmarshal
	// Not present for clients that support the `simplified_presence_events` [client capability](https://zulip.com/api/register-queue#parameter-client_capabilities.
	Deprecated *PresenceEventDeprecated
}

type PresenceEventDeprecated struct {
	// The Id of the modified user.
	UserId int64 `json:"user_id,omitempty"`
	// The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the `user_id`.
	// Deprecated
	Email string `json:"email,omitempty"`
	// The timestamp of when the Zulip server received the user's presence as a UNIX timestamp.
	ServerTimestamp float32 `json:"server_timestamp,omitempty"`
	// Object containing the presence data (legacy format) of of the modified user.
	// PresenceFormatDeprecated `{client_name}`: Object containing the details of the user's presence.
	//
	// **Changes**: Starting with Zulip 7.0 (feature level 178), this will always be `"website"` as the server no longer stores which client submitted presence updates.  Previously, the object key was the client's platform name, for example `website` or `ZulipDesktop`.
	Presence map[string]interface{} `json:"presence,omitempty"`
}

// ChannelCreateEvent Event sent when a new channel is created to users who can see the new channel exists (for private channels, only subscribers and organization administrators will receive this event).  This event is also sent when a user gains access to a channel they previously [could not access], such as when their [role] changes, a private channel is made public, or a guest user is subscribed to a public (or private) channel.  This event is also sent when a channel is unarchived but only to clients that did not declare the `archived_channels` [client capability].  Note that organization administrators who are not subscribed will not be able to see content on the channel; just that it exists.
//
//	**Changes**: Prior to Zulip 11.0 (feature level 378), this event was sent to all the users who could see the channel when it was unarchived.  Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users a web-public channel was created.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user gained access to a channel due to their role changing.  Prior to Zulip 8.0 (feature level 192), this event was not sent when guest users gained access to a public channel by being subscribed.  Prior to Zulip 6.0 (feature level 134), this event was not sent when a private channel was made public.
//
// [could not access]: https://zulip.com/help/channel-permissions
// [role]: https://zulip.com/help/user-roles
type ChannelCreateEvent struct {
	EventCommonWithOp
	// Array of objects, each containing details about the newly added channel(s).
	Channels []Channel `json:"streams,omitempty"`
}

// ChannelDeleteEvent Event sent when a user loses access to a channel they previously [could access] because they are unsubscribed from a private channel or their [role] has changed.  This event is also sent when a channel is archived but only to clients that did not declare the `archived_channels` [client capability].
//
//	**Changes**: Prior to Zulip 11.0 (feature level 378), this event was sent to all the users who could see the channel when it was archived.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user lost access to a channel due to their role changing.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
//
// [could access]: https://zulip.com/help/channel-permissions
// [role]: https://zulip.com/help/user-roles
type ChannelDeleteEvent struct {
	EventCommonWithOp

	// Array of objects, each containing Id of the channel that was deleted.
	//
	// **Changes**: **Deprecated** in Zulip 10.0 (feature level 343) and will be removed in a future release. Previously, these objects additionally contained all the standard fields for a channel object.
	// Deprecated
	Channels []interface{} `json:"streams,omitempty"`

	// Array containing the Ids of the channels that were deleted.
	//
	// **Changes**: New in Zulip 10.0 (feature level 343). Previously, these Ids were available only via the legacy `streams` array.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
}

// ChannelUpdateEvent Event sent to all users who can see that a channel exists when a property of that channel changes. See [GET /streams] response for details on the various properties of a channel.  This event is also sent when archiving or unarchiving a channel to all the users who can see that channel exists but only to the clients that declared the `archived_channels` [client capability].
//
//	**Changes**: Prior to Zulip 11.0 (feature level 378), this event was never sent when archiving or unarchiving a channel.  Before Zulip 9.0 (feature level 256), this event was never sent when the `first_message_id` property of a channel was updated because the oldest message that had been sent to it changed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
//
// [GET /streams]: https://zulip.com/api/get-streams#response
type ChannelUpdateEvent struct {
	EventCommonWithOp
	// The Id of the channel whose details have changed.
	ChannelId int64 `json:"stream_id,omitempty"`
	// The name of the channel whose details have changed.
	Name string `json:"name,omitempty"`
	// The property of the channel which has changed. See [GET /streams] response for details on the various properties of a channel.  Clients should handle an "unknown" property received here without crashing, since that can happen when connecting to a server running a newer version of Zulip with new features.
	//
	// [GET /streams]: https://zulip.com/api/get-streams#response
	Property string `json:"property,omitempty"`
	// ChannelEventUpdateValue - The new value of the changed property.
	//
	// **Changes**: Starting with Zulip 11.0 (feature level 389), this value can be `null` when a channel is removed from the folder.  Starting with Zulip 10.0 (feature level 320), this field can be an object for `can_remove_subscribers_group` property, which is a [group-setting value], when the setting is set to a combination of users and groups.  [group-setting value]: https://zulip.com/api/group-setting-values
	Value *ChannelEventUpdateValue `json:"value,omitempty"`
	// Note: Only present if the changed property was `description`.  The short description of the channel rendered as HTML, intended to be used when displaying the channel description in a UI.  One should use the standard Zulip rendered_markdown CSS when displaying this content so that emoji, LaTeX, and other syntax work correctly. And any client-side security logic for user-generated message content should be applied when displaying this HTML as though it were the body of a Zulip message.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	RenderedDescription *string `json:"rendered_description,omitempty"`
	// Note: Only present if the changed property was `invite_only`.  Whether the history of the channel is public to its subscribers.  Currently always true for public channels (i.e. `"invite_only": false` implies `"history_public_to_subscribers": true`), but clients should not make that assumption, as we may change that behavior in the future.
	HistoryPublicToSubscribers *bool `json:"history_public_to_subscribers,omitempty"`
	// Note: Only present if the changed property was `invite_only`.  Whether the channel's history is now readable by web-public spectators.
	//
	// **Changes**: New in Zulip 5.0 (feature level 71).
	IsWebPublic *bool `json:"is_web_public,omitempty"`
}

type ChannelEventUpdateValue struct {
	GroupSettingValue *GroupSettingValue
	Bool              *bool
	Int64             *int64
	String            *string
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChannelEventUpdateValue) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChannelEventUpdateValue) MarshalJSON() ([]byte, error) {
	return marshalUnionType(src)
}

// ReactionEvent Event sent when a reaction is added to or removed from a message. Sent to all users who were recipients of the message.
type ReactionEvent struct {
	EventCommonWithOp
	// The Id of the message to which a reaction was added or removed.
	MessageId int64 `json:"message_id,omitempty"`

	// Name of the emoji.
	EmojiName string `json:"emoji_name,omitempty"`
	// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
	EmojiCode string `json:"emoji_code,omitempty"`
	// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the Id of   the uploaded [custom emoji].  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   "zulip").
	//
	// [custom emoji]: https://zulip.com/help/custom-emoji
	ReactionType ReactionType `json:"reaction_type,omitempty"`
	// The Id of the user who added the reaction.
	//
	// **Changes**: New in Zulip 3.0 (feature level 2). The `user` object is deprecated and will be removed in the future.
	UserId int64 `json:"user_id,omitempty"`

	// Deprecated
	User interface{} `json:"user,omitempty"`
}

// UpdateGlobalNotificationsEvent Event sent to a user's clients when that user's [notification settings] have changed with an additional rule that it is only sent to clients that did not include `user_settings_object` in their `client_capabilities` when registering the event queue.
//
//	**Changes**: Deprecated in Zulip 5.0 (feature level 89). Clients connecting to newer servers should declare the `user_settings_object` client capability and process the `user_settings` event type instead.
//
// [notification settings]: https://zulip.com/api/update-settings
type UpdateGlobalNotificationsEvent struct {
	EventCommon
	// The Zulip API email of the user.
	User string `json:"user,omitempty"`
	// Name of the changed notification setting.
	NotificationName string      `json:"notification_name,omitempty"`
	Setting          UpdateValue `json:"setting,omitempty"`
}

// UpdateValue - New value of the changed setting.
type UpdateValue struct {
	Bool   *bool
	Int64  *int64
	String *string
}

// boolAsUpdateValue is a convenience function that returns bool wrapped in UpdateValue
func BoolAsUpdateValue(v bool) UpdateValue {
	return UpdateValue{
		Bool: &v,
	}
}

// Int64AsUpdateValue is a convenience function that returns int64 wrapped in UpdateValue
func Int64AsUpdateValue(v int64) UpdateValue {
	return UpdateValue{
		Int64: &v,
	}
}

// StringAsUpdateValue is a convenience function that returns string wrapped in UpdateValue
func StringAsUpdateValue(v string) UpdateValue {
	return UpdateValue{
		String: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateValue) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateValue) MarshalJSON() ([]byte, error) {
	return marshalUnionType(src)
}

// AttachmentAddEvent Event sent to a user's clients when the user uploads a new file in a Zulip message. Useful to implement live update in UI showing all files the current user has uploaded.
type AttachmentAddEvent struct {
	EventCommonWithOp
	Attachment Attachment `json:"attachment,omitempty"`
	// The total size of all files uploaded by in the organization, in bytes.
	UploadSpaceUsed int64 `json:"upload_space_used,omitempty"`
}

// AttachmentUpdateEvent Event sent to a user's clients when details of a file that user uploaded are changed. Most updates will be changes in the list of messages that reference the uploaded file.
type AttachmentUpdateEvent struct{ AttachmentAddEvent }

// AttachmentRemoveEvent Event sent to a user's clients when the user deletes a file they had uploaded. Useful primarily for UI showing all the files the current user has uploaded.
type AttachmentRemoveEvent struct {
	EventCommonWithOp
	Attachment AttachmentId `json:"attachment,omitempty"`
	// The total size of all files uploaded by in the organization, in bytes.
	UploadSpaceUsed int64 `json:"upload_space_used,omitempty"`
}

type AttachmentId struct {
	// The Id of the deleted attachment.
	Id int64 `json:"id,omitempty"`
}

// PushDeviceEvent Event sent to a user's clients when the metadata in the `push_devices` dictionary for the user changes.  Helps clients to live-update the `push_devices` dictionary returned in [`POST /register`] response.
//
//	**Changes**: New in Zulip 11.0 (feature level 406).
//
// [`POST /register`]: https://zulip.com/api/register-queue
type PushDeviceEvent struct {
	EventCommon
	// The push account Id for this client registration.  See [`POST /mobile_push/register`] for details on push account Ids.
	//
	// [`POST /mobile_push/register`]: https://zulip.com/api/register-push-device
	PushAccountId string `json:"push_account_id,omitempty"`
	// The updated registration status. Will be `"active"`, `"failed"`, or `"pending"`.
	Status string `json:"status,omitempty"`
	// If the status is `"failed"`, a [Zulip API error code] indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - `"INVALId_BOUNCER_PUBLIC_KEY"` - Inform the user to update app. - `"REQUEST_EXPIRED` - Retry with a fresh payload.   If the status is "failed", an error code explaining the failure.
	//
	// [Zulip API error code]: https://zulip.com/api/rest-error-handling
	ErrorCode *string `json:"error_code,omitempty"`
}

// SubmessageEvent Event sent when a submessage is added to a message.  Submessages are an **experimental** API used for widgets such as the `/poll` widget in Zulip.
type SubmessageEvent struct {
	EventCommon
	// The type of the message.
	MsgType string `json:"msg_type,omitempty"`
	// The new content of the submessage.
	Content string `json:"content,omitempty"`
	// The Id of the message to which the submessage has been added.
	MessageId int64 `json:"message_id,omitempty"`
	// The Id of the user who sent the message.
	SenderId int64 `json:"sender_id,omitempty"`
	// The Id of the submessage.
	SubmessageId int64 `json:"submessage_id,omitempty"`
}

// CustomProfileFieldsEvent Event sent to all users in a Zulip organization when new custom profile field types are configured for that Zulip organization.
type CustomProfileFieldsEvent struct {
	EventCommon
	// An array of dictionaries where each dictionary contains details of a single new custom profile field for the Zulip organization.
	Fields []CustomProfileField `json:"fields,omitempty"`
}

// CustomProfileField Dictionary containing the details of a custom profile field configured for this organization.
type CustomProfileField struct {
	// The Id of the custom profile field. This will be referenced in the custom profile fields section of user objects.
	Id int64 `json:"id"`
	// An integer indicating the type of the custom profile field, which determines how it is configured and displayed to users.  See the [Custom profile fields] article for details on what each type means.  - **1**: Short text - **2**: Long text - **3**: List of options - **4**: Date picker - **5**: Link - **6**: Person picker - **7**: External account - **8**: Pronouns  **Changes**: Field type `8` added in Zulip 6.0 (feature level 151).
	//
	// [Custom profile fields]: https://zulip.com/help/custom-profile-fields#profile-field-types
	Type CustomFieldType `json:"type"`
	// Custom profile fields are displayed in both settings UI and UI showing users' profiles in increasing `order`.
	Order int32 `json:"order"`
	// The name of the custom profile field.
	Name string `json:"name"`
	// The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.
	Hint string `json:"hint"`
	// Field types 3 (List of options) and 7 (External account) support storing additional configuration for the field type in the `field_data` attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type 7 is not yet stabilized.
	FieldData *string `json:"field_data,omitempty"`
	// Whether the custom profile field, display or not on the user card.  Currently it's value not allowed to be `true` of `Long text` and `Person picker` [profile field types].  This field is only included when its value is `true`.
	//
	// **Changes**: New in Zulip 6.0 (feature level 146).
	//
	// [profile field types]: https://zulip.com/help/custom-profile-fields#profile-field-types
	DisplayInProfileSummary *bool `json:"display_in_profile_summary,omitempty"`
	// Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.
	//
	// **Changes**: New in Zulip 9.0 (feature level 244).
	Required bool `json:"required"`
	// Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.
	//
	// **Changes**: New in Zulip 10.0 (feature level 296).
	EditableByUser bool `json:"editable_by_user"`
}

// DefaultChannelGroupsEvent Event sent to all users in a Zulip organization when an organization administrator changes the organization's configured default channel groups.  Default channel groups are an **experimental** feature that is not yet stabilized.
type DefaultChannelGroupsEvent struct {
	EventCommon
	// An array of dictionaries where each dictionary contains details about a single default channel group.
	DefaultChannelGroups []DefaultChannelGroup `json:"default_stream_groups,omitempty"`
}

// DefaultChannelsEvent Event sent to all users in a Zulip organization when the default channels in the organization are changed by an organization administrator.
type DefaultChannelsEvent struct {
	EventCommon
	// An array of Ids of all the [default channels] in the organization.
	//
	// **Changes**: Before Zulip 10.0 (feature level 330), we sent array of dictionaries where each dictionary contained details about a single default stream for the Zulip organization.
	//
	// [default channels]: https://zulip.com/help/set-default-streams-for-new-users
	DefaultChannels []int64 `json:"default_streams,omitempty"`
}

// UserSettingsUpdateEvent Event sent to a user's clients when that user's settings have changed.
//
//	**Changes**: New in Zulip 5.0 (feature level 89), replacing the previous `update_display_settings` and `update_global_notifications` event types, which are still present for backwards compatibility reasons.
type UserSettingsUpdateEvent struct {
	EventCommonWithOp
	// Name of the changed setting.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
	// Present only if the setting to be changed is `default_language`. Contains the name of the new default language in English.
	LanguageName *string `json:"language_name,omitempty"`
}

// UserStatusEvent Event sent to all users who can access the modified user when the status of a user changes.
//
//	**Changes**: Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
type UserStatusEvent struct {
	EventCommon
	// Whether the user has marked themself "away" with this status.
	//
	// **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, `away` is a legacy way to access the user's `presence_enabled` setting, with `away = !presence_enabled`. To be removed in a future release.
	// Deprecated
	Away *bool `json:"away,omitempty"`

	// The text content of the status message.  This will be `""` for users who set a status without selecting or writing a message.
	StatusText string `json:"status_text,omitempty"`
	// The [emoji name].
	//
	// [emoji name]: https://zulip.com/api/update-status#parameter-emoji_name for the emoji the user selected for their new status.  This will be `""` for users who set a status without selecting an emoji.
	//
	// **Changes**: New in Zulip 5.0 (feature level 86
	EmojiName string `json:"emoji_name,omitempty"`
	// The [emoji code].
	//
	// [emoji code]: https://zulip.com/api/update-status#parameter-emoji_code for the emoji the user selected for their new status.  This will be `""` for users who set a status without selecting an emoji.
	//
	// **Changes**: New in Zulip 5.0 (feature level 86
	EmojiCode string `json:"emoji_code,omitempty"`
	// The [emoji type].
	//
	// [emoji type]: https://zulip.com/api/update-status#parameter-reaction_type for the emoji the user selected for their new status.  This will be `""` for users who set a status without selecting an emoji.
	//
	// **Changes**: New in Zulip 5.0 (feature level 86
	ReactionType ReactionType `json:"reaction_type,omitempty"`
	// The Id of the user whose status changed.
	UserId int64 `json:"user_id,omitempty"`
}

// DeleteMessageEvent Event sent when a message has been deleted.  Sent to all users who currently are subscribed to the messages' recipient. May also be sent to additional users who had access to it, including, in particular, an administrator user deleting messages in a stream that they are not subscribed to.  This means that clients can assume that they will always receive an event of this type for deletions that the client itself initiated.  This event is also sent when the user loses access to a message, such as when it is [moved to a channel] that the user does not [have permission to access].
//
//	**Changes**: Before Zulip 9.0 (feature level 274), this event was only sent to subscribers of the message's recipient.  Before Zulip 5.0 (feature level 77), events for direct messages contained additional `sender_id` and `recipient_id` fields.
//
// [moved to a channel]: https://zulip.com/help/move-content-to-another-channel
// [have permission to access]: https://zulip.com/help/channel-permissions
type DeleteMessageEvent struct {
	EventCommon
	// Only present for clients that support the `bulk_message_deletion` [client capability].  A sorted list containing the Ids of the newly deleted messages.
	//
	// **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	MessageIds []int64 `json:"message_ids,omitempty"`
	// Only present for clients that do not support the `bulk_message_deletion` [client capability].  The Id of the newly deleted message.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	MessageId *int64 `json:"message_id,omitempty"`

	// The type of message. Either `"stream"` or `"private"`.
	MessageType RecipientType `json:"message_type,omitempty"`
	// Only present if `message_type` is `"stream"`.  The Id of the channel to which the message was sent.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// Only present if `message_type` is `"stream"`.  The topic to which the message was sent.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name was empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
}

// UserTopicEvent Event sent to a user's clients when the user mutes/unmutes a topic, or otherwise modifies their personal per-topic configuration.
//
//	**Changes**: New in Zulip 6.0 (feature level 134). Previously, clients were notified about changes in muted topic configuration via the `muted_topics` event type.
type UserTopicEvent struct {
	EventCommon
	UserTopic
}

// MutedTopicsEvent Event sent to a user's clients when that user's set of configured muted topics have changed.
type MutedTopicsEvent struct {
	EventCommon
	// Array of tuples, where each tuple describes a muted topic. The first element of the tuple is the channel name in which the topic has to be muted, the second element is the topic name to be muted and the third element is an integer UNIX timestamp representing when the topic was muted.
	//
	// **Changes**: Deprecated in Zulip 6.0 (feature level 134). Starting with this version, clients that explicitly requested the replacement `user_topic` event type when registering their event queue will not receive this legacy event type.  Before Zulip 3.0 (feature level 1), the `muted_topics` array objects were 2-item tuples and did not include the timestamp information for when the topic was muted.
	// Deprecated
	MutedTopics [][]interface{} `json:"muted_topics,omitempty"`
}

// MutedUsersEvent Event sent to a user's clients when that user's set of configured [muted users] have changed.
//
//	**Changes**: New in Zulip 4.0 (feature level 48).
//
// [muted users]: https://zulip.com/api/mute-user
type MutedUsersEvent struct {
	EventCommon
	// A list of dictionaries where each dictionary describes a muted user.
	MutedUsers []MutedUser `json:"muted_users,omitempty"`
}

// OnboardingStepsEvent Event sent when the set of onboarding steps to show for the current user has changed (e.g. because the user dismissed one).  Clients that feature a similar tutorial experience to the Zulip web app may want to handle these events.
//
//	**Changes**: Before Zulip 8.0 (feature level 233), this event was named `hotspots`. Prior to this feature level, one-time notice onboarding steps were not supported.
type OnboardingStepsEvent struct {
	EventCommon
	// An array of dictionaries where each dictionary contains details about a single onboarding step.
	//
	// **Changes**: Before Zulip 8.0 (feature level 233), this array was named `hotspots`. Prior to this feature level, one-time notice onboarding steps were not supported, and the `type` field in these objects did not exist as all onboarding steps were implicitly hotspots.
	OnboardingSteps []OnboardingStep `json:"onboarding_steps,omitempty"`
}

// OnboardingStep Dictionary containing details of a single onboarding step.
type OnboardingStep struct {
	// The type of the onboarding step. Valid value is `"one_time_notice"`.
	//
	// **Changes**: Removed type `"hotspot"` in Zulip 9.0 (feature level 259).  New in Zulip 8.0 (feature level 233).
	Type string `json:"type,omitempty"`
	// The name of the onboarding step.
	Name string `json:"name,omitempty"`
}

// UpdateMessageEvent Event sent when a message's content, topic and/or channel has been edited or when a message's content has a rendering update, such as for an [inline URL preview]. Sent to all users who had received the original message.  [inline URL preview]: https://zulip.readthedocs.io/en/latest/subsystems/sending-messages.html#inline-url-previews  **Changes**: In Zulip 10.0 (feature level 284), removed the `prev_rendered_content_version` field as it is an internal server implementation detail not used by any client.
type UpdateMessageEvent struct {
	EventCommon
	// The Id of the user who sent the message.  Is `null` when event is for a rendering update of the original message, such as for an [inline URL preview].
	//
	// **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all `update_message` events. Previously, this field was omitted for [inline URL preview] updates.
	UserId *int64 `json:"user_id"`
	// Whether the event only updates the rendered content of the message.  This field should be used by clients to determine if the event only provides a rendering update to the message content, such as for an [inline URL preview]. When `true`, the event does not reflect a user-generated edit and does not modify the message history.
	//
	// **Changes**: New in Zulip 5.0 (feature level 114). Clients can correctly identify these rendering update event with earlier Zulip versions by checking whether the `user_id` field was omitted.
	RenderingOnly bool `json:"rendering_only"`
	// The Id of the message which was edited or updated.  This field should be used to apply content edits to the client's cached message history, or to apply rendered content updates.  If the channel or topic was changed, the set of moved messages is encoded in the separate `message_ids` field, which is guaranteed to include `message_id`.
	MessageId int64 `json:"message_id"`
	// A sorted list of Ids of messages to which any channel or topic changes encoded in this event should be applied.  This list always includes `message_id`, even when there are no channel or topic changes to apply.  These messages are guaranteed to have all been previously sent to channel `stream_id` with topic `orig_subject`, and have been moved to `new_stream_id` with topic `subject` (if those fields are present in the event).  Clients processing these events should update all cached message history associated with the moved messages (including adjusting `unread_msgs` data structures, where the client may not have the message itself in its history) to reflect the new channel and topic.  Content changes should be applied only to the single message indicated by `message_id`.
	//
	// **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.
	MessageIds []int64 `json:"message_ids"`
	// The user's personal [message flags] for the message with Id `message_id` following the edit.  A client application should compare these to the original flags to identify cases where a mention or alert word was added by the edit.
	//
	// **Changes**: In Zulip 8.0 (feature level 224), the `wildcard_mentioned` flag was deprecated in favor of the `stream_wildcard_mentioned` and `topic_wildcard_mentioned` flags. The `wildcard_mentioned` flag exists for backwards compatibility with older clients and equals `stream_wildcard_mentioned || topic_wildcard_mentioned`. Clients supporting older server versions should treat this field as a previous name for the `stream_wildcard_mentioned` flag as topic wildcard mentions were not available prior to this feature level.  [message flags]: https://zulip.com/api/update-message-flags#available-flags
	Flags []string `json:"flags"`
	// The time when this message edit operation was processed by the server.
	//
	// **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all `update_message` events. Previously, this field was omitted for [inline URL preview] updates.
	EditTimestamp time.Time `json:"edit_timestamp"`
	// Only present if the message was edited and originally sent to a channel.  The name of the channel that the message was sent to. Clients are recommended to use the `stream_id` field instead.
	ChannelName *string `json:"stream_name,omitempty"`
	// Only present if the message was edited and originally sent to a channel.  The pre-edit channel for all of the messages with Ids in `message_ids`.
	//
	// **Changes**: As of Zulip 5.0 (feature level 112), this field is present for all edits to a channel message. Previously, it was not present when only the content of the channel message was edited.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// Only present if message(s) were moved to a different channel.  The post-edit channel for all of the messages with Ids in `message_ids`.
	NewChannelId *int64 `json:"new_stream_id,omitempty"`
	// Only present if this event moved messages to a different topic and/or channel.  The choice the editing user made about which messages should be affected by a channel/topic edit:  - `"change_one"`: Just change the one indicated in `message_id`. - `"change_later"`: Change messages in the same topic that had   been sent after this one. - `"change_all"`: Change all messages in that topic.  This parameter should be used to decide whether to change navigation and compose box state in response to the edit. For example, if the user was previously in topic narrow, and the topic was edited with `"change_later"` or `"change_all"`, the Zulip web app will automatically navigate to the new topic narrow. Similarly, a message being composed to the old topic should have its recipient changed to the new topic.  This navigation makes it much more convenient to move content between topics without disruption or messages continuing to be sent to the pre-edit topic by accident.
	PropagateMode *string `json:"propagate_mode,omitempty"`
	// Only present if this event moved messages to a different topic and/or channel.  The pre-edit topic for all of the messages with Ids in `message_ids`.  For clients that don't support the `empty_topic_name` [client capability], if the actual pre-edit topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	OrigSubject *string `json:"orig_subject,omitempty"`
	// Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  The post-edit topic for all of the messages with Ids in `message_ids`.  For clients that don't support the `empty_topic_name` [client capability], if the actual post-edit topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Subject *string `json:"subject,omitempty"`
	// Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  Data on any links to be included in the `topic` line (these are generated by [custom linkification filter] that match content in the message's topic.), corresponding to the post-edit topic.
	//
	// **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called `subject_links`; clients are recommended to rename `subject_links` to `topic_links` if present for compatibility with older Zulip servers.
	//
	// [custom linkification filter]: https://zulip.com/help/add-a-custom-linkifier
	TopicLinks []TopicLink `json:"topic_links,omitempty"`
	// Only present if this event changed the message content.  The original content of the message with Id `message_id` immediately prior to this edit, in the original [Zulip-flavored Markdown] format.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	OrigContent *string `json:"orig_content,omitempty"`
	// Only present if this event changed the message content.  The original content of the message with Id `message_id` immediately prior to this edit, rendered as HTML.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	OrigRenderedContent *string `json:"orig_rendered_content,omitempty"`
	// Only present if this event changed the message content or updated the message content for an [inline URL preview].  The new content of the message with Id `message_id`, in the original [Zulip-flavored Markdown] format.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	Content *string `json:"content,omitempty"`
	// Only present if this event changed the message content or updated the message content for an [inline URL preview].  The new content of the message with Id `message_id`, rendered in HTML.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	RenderedContent *string `json:"rendered_content,omitempty"`
	// Only present if this event changed the message content.  Whether the message with Id `message_id` is now a [/me status message].
	//
	// [/me status message]: https://zulip.com/help/format-your-message-using-markdown#status-messages
	IsMeMessage *bool `json:"is_me_message,omitempty"`
}

// TypingStartEvent Event sent when a user starts typing a message.  Sent to all clients for users who would receive the message being typed, with the additional rule that typing notifications for channel messages are only sent to clients that included `stream_typing_notifications` in their [client capabilities] when registering the event queue.  See [POST /typing] endpoint for more details.
//
//	**Changes**: Typing notifications for channel messages are new in Zulip 4.0 (feature level 58).
//
// [client capabilities]: https://zulip.com/api/register-queue#parameter-client_capabilities
// [POST /typing]: https://zulip.com/api/set-typing-status
type TypingEvent struct {
	EventCommonWithOp
	// Type of message being composed. Must be `"stream"` or `"direct"`.
	//
	// **Changes**: In Zulip 8.0 (feature level 215), replaced the value `"private"` with `"direct"`.  New in Zulip 4.0 (feature level 58). Previously, all typing notifications were implicitly direct messages.
	MessageType RecipientType  `json:"message_type,omitempty"`
	Sender      UserIdentifier `json:"sender,omitempty"`
	// Only present if `message_type` is `"direct"`.  Array of dictionaries describing the set of users who would be recipients of the message being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.
	Recipients []UserIdentifier `json:"recipients,omitempty"`
	// Only present if `message_type` is `"stream"`.  The unique Id of the channel to which message is being typed.
	//
	// **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// Only present if `message_type` is `"stream"`.  Topic within the channel where the message is being typed.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
}

// UserSettingsUpdateEvent7RecipientsInner Object containing the user Id and Zulip API email of a recipient.
type UserIdentifier struct {
	// The Id of the user.
	UserId int64 `json:"user_id,omitempty"`
	// The Zulip API email address for the user.
	Email string `json:"email,omitempty"`
}

// TypingEditMessageStartEvent Event sent when a user starts editing a message. Event sent when a user starts typing in a textarea to edit the content of a message. See the [edit message typing notifications endpoint].  Clients requesting `typing_edit_message` event type that have `receives_typing_notifications` enabled will receive this event if they would have been notified if the message's content edit were to be saved (E.g., because they were a direct message recipient or are a subscribe to the channel).
//
//	**Changes**: New in Zulip 10.0 (feature level 351). Previously, typing notifications were not available when editing messages.
//
// [edit message typing notifications endpoint]: https://zulip.com/api/set-typing-status-for-message-edit
type TypingEditMessageEvent struct {
	EventCommonWithOp

	// The Id of the user who is typing the edit of the message.  Clients should be careful to display this user as the person who is typing, not that of the sender of the message, in case a collaborative editing feature be might be added in the future.
	SenderId int64 `json:"sender_id,omitempty"`
	// Indicates the message id of the message that is being edited.
	MessageId int64          `json:"message_id,omitempty"`
	Recipient *RecipientData `json:"recipient,omitempty"`
}

// Recipient Object containing details about recipient of message edit typing notification.
type RecipientData struct {
	// Type of message being composed. Must be `"channel"` or `"direct"`.
	Type RecipientType `json:"type,omitempty"`
	// Only present if `type` is `"channel"`.  The unique Id of the channel to which message is being edited.
	ChannelId *int64 `json:"channel_id,omitempty"`
	// Only present if `type` is `"channel"`.  Topic within the channel where the message is being edited.
	Topic *string `json:"topic,omitempty"`
	// Present only if `type` is `direct`.  The user Ids of every recipient of this direct message.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// RealmUserUpdateEvent Event sent generally to all users who can access the modified user for changes in the set of users or those users metadata.
//
//	**Changes**: Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
type RealmUserUpdateEvent struct {
	EventCommonWithOp

	Person UserUpdate `json:"person,omitempty"`
}

// UserUpdate - Object containing the changed details of the user. It has multiple forms depending on the value changed.
//
//	**Changes**: Removed `is_billing_admin` field in Zulip 10.0 (feature level 363), as it was replaced by the `can_manage_billing_group` realm setting.
type UserUpdate struct {
	// The Id of the user who is affected by this change.
	UserId int64 `json:"user_id,omitempty"`

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

// UserUpdateEventAvatar When a user changes their avatar.
type UserUpdateEventAvatar struct {
	Avatar
	// The version number for the user's avatar. This is useful for cache-busting.
	AvatarVersion int64 `json:"avatar_version,omitempty"`
}

// UserUpdateEventFullName When a user changes their full name.
type UserUpdateEventFullName struct {
	// The new full name for the user.
	FullName string `json:"full_name,omitempty"`
}

// UserUpdateEventTimezone When a user changes their [profile time zone].
//
// [profile time zone]: https://zulip.com/help/change-your-timezone
type UserUpdateEventTimezone struct {
	// The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the `user_id`.
	// Deprecated
	Email string `json:"email,omitempty"`
	// The IANA identifier of the new profile time zone for the user.
	Timezone string `json:"timezone,omitempty"`
}

// UserUpdateEventBotOwner When the owner of a bot changes.
type UserUpdateEventBotOwner struct {
	// The user Id of the new bot owner.
	BotOwnerId int64 `json:"bot_owner_id,omitempty"`
}

// UserUpdateEventRole When the [role] of a user changes.
//
// [role]: https://zulip.com/help/user-roles
type UserUpdateEventRole struct {
	// The new [role] of the user.
	//
	// [role]: https://zulip.com/api/roles-and-permissions
	Role Role `json:"role,omitempty"`
}

// UserUpdateEventDeliveryEmail When the value of a user's delivery email as visible to you changes, either due to the email address changing or your access to the user's email changing via an update to their `email_address_visibility` setting.
//
//	**Changes**: Prior to Zulip 7.0 (feature level 163), this event was sent only to the affected user, and this event would only be triggered by changing the affected user's delivery email.
type UserUpdateEventDeliveryEmail struct {
	// The new delivery email of the user.  This value can be `null` if the affected user changed their `email_address_visibility` setting such that you cannot access their real email.
	//
	// **Changes**: Before Zulip 7.0 (feature level 163), `null` was not a possible value for this event as it was only sent to the affected user when their email address was changed.
	DeliveryEmail *string `json:"delivery_email,omitempty"`
}

// UserUpdateEventCustomField When the user updates one of their custom profile fields.
type UserUpdateEventCustomField struct {
	CustomProfileField ProfileDataValue `json:"custom_profile_field,omitempty"`
}

// UserUpdateEventEmail When the Zulip API email address of a user changes, either due to the user's email address changing, or due to changes in the user's [email address visibility].
//
// [email address visibility]: https://zulip.com/help/configure-email-visibility
type UserUpdateEventEmail struct {
	// The new value of `email` for the user. The client should update any data structures associated with this user to use this new value as the user's Zulip API email address.
	NewEmail string `json:"new_email,omitempty"`
}

// UserUpdateEventActivation When a user is deactivated or reactivated. Only users who can access the modified user under the organization's `can_access_all_users_group` policy will receive this event.  Clients receiving a deactivation event should remove the user from all user groups in their data structures, because deactivated users cannot be members of groups.
//
//	**Changes**: Prior to Zulip 10.0 (feature level 303), reactivation events were sent to users who could not access the reactivated user due to a `can_access_all_users_group` policy. Also, previously, Clients were not required to update group membership records during user deactivation.  New in Zulip 8.0 (feature level 222). Previously the server sent a `realm_user` event with `op` field set to `remove` when deactivating a user and a `realm_user` event with `op` field set to `add` when reactivating a user.
type UserUpdateEventActivation struct {
	// A boolean describing whether the user account has been deactivated.
	IsActive bool `json:"is_active,omitempty"`
}

// UpdateMessageFlagsAddEvent Event sent to a user when [message flags] are added to messages.  This can reflect a direct user action, or can be the indirect consequence of another action. Whatever the cause, if there's a change in the set of message flags that the user has for a message, then an `update_message_flags` event will be sent with the change. Note that this applies when the user already had access to the message, and continues to have access to it. When a message newly appears or disappears, a [`message`] or [`delete_message`] event is sent instead.  Some examples of actions that trigger an `update_message_flags` event:  - The `"starred"` flag is added when the user chooses to [star a   message]. - The `"read"` flag is added when the user marks messages as read by   scrolling through them, or uses [Mark all messages as read] on a conversation. - The `"read"` flag is added when the user [mutes] a   message's sender. - The `"read"` flag is added after the user unsubscribes from a channel,   or messages are moved to a not-subscribed channel, provided the user   can still access the messages at all. Note a   [`delete_message`] event is sent in the case where the   user can no longer access the messages.  In some cases, a change in message flags that's caused by another change may happen a short while after the original change, rather than simultaneously. For example, when messages that were unread are moved to a channel where the user is not subscribed, the resulting change in message flags (and the corresponding `update_message_flags` event with flag `"read"`) may happen later than the message move itself. The delay in that example is typically at most a few hundred milliseconds and can in rare cases be minutes or longer.
//
// [message flags]: https://zulip.com/api/update-message-flags#available-flags
// [`message`]: https://zulip.com/api/get-events#message
// [`delete_message`]: https://zulip.com/api/get-events#delete_message
// [Mark all messages as read]: https://zulip.com/help/marking-messages-as-read#mark-messages-in-multiple-topics-and-channels-as-read
// [star a   message]: https://zulip.com/help/star-a-message
// [mutes]: https://zulip.com/help/mute-a-user
type UpdateMessageFlagsAddEvent struct {
	EventCommonWithOp

	// Old name for the `op` field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the `op` field.
	// Deprecated
	Operation *string `json:"operation,omitempty"`
	// The [flag] that was added/removed. [flag]: https://zulip.com/api/update-message-flags#available-flags
	Flag string `json:"flag,omitempty"`
	// Array containing the Ids of all messages to which the flag was added/removed.
	Messages []int64 `json:"messages,omitempty"`
	// Whether the specified flag was added to all messages. This field is only relevant for the `"read"` flag, and will be `false` for all other flags.  When `true` for the `"read"` flag, then the `messages` array will be empty.
	All bool `json:"all,omitempty"`
}

// UpdateMessageFlagsRemoveEvent Event sent to a user when [message flags] are removed from messages.  See the description for the [`update_message_flags` op: `add`](https://zulip.com/api/get-events#update_message_flags-add event for more details about these events.  [message flags]: https://zulip.com/api/update-message-flags#available-flags
type UpdateMessageFlagsRemoveEvent struct {
	EventCommonWithOp

	// Old name for the `op` field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the `op` field.
	// Deprecated
	Operation *string `json:"operation,omitempty"`
	// The [flag] to be removed.
	//
	// [flag]: https://zulip.com/api/update-message-flags#available-flags
	Flag string `json:"flag,omitempty"`
	// Array containing the Ids of the messages from which the flag was removed.
	Messages []int64 `json:"messages,omitempty"`
	// Will be `false` for all specified flags.  **Deprecated** and will be removed in a future release.
	// Deprecated
	All *bool `json:"all,omitempty"`

	// Only present if the specified `flag` is `"read"`.  A set of data structures describing the messages that are being marked as unread with additional details to allow clients to update the `unread_msgs` data structure for these messages (which may not be otherwise known to the client).
	//
	// **Changes**: New in Zulip 5.0 (feature level 121). Previously, marking already read messages as unread was not supported by the Zulip API.
	MessageDetails map[string]MessageDetail `json:"message_details,omitempty"`
}

// UpdateMessageFlagsRemoveEventMessageDetailsValue `{message_id}`: Object containing details about the message with the specified Id.
type MessageDetail struct {
	// The type of this message. Either `"stream"` or `"private"`.
	Type RecipientType `json:"type"`
	// A flag which indicates whether the message contains a mention of the user.  Present only if the message mentions the current user.
	Mentioned *bool `json:"mentioned,omitempty"`
	// Present only if `type` is `private`.  The user Ids of every recipient of this direct message, excluding yourself. Will be the empty list for a message you had sent to only yourself.
	UserIds []int64 `json:"user_ids,omitempty"`
	// Present only if `type` is `"stream"`.  The Id of the channel where the message was sent.
	ChannelId *int64 `json:"stream_id,omitempty"`
	// Present only if `type` is `"stream"`.  Name of the topic where the message was sent.  For clients that don't support the `empty_topic_name` [client capability], if the actual topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Topic *string `json:"topic,omitempty"`
	// **Deprecated** internal implementation detail. Clients should ignore this field as it will be removed in the future.
	// Deprecated
	UnmutedStreamMsg *bool `json:"unmuted_stream_msg,omitempty"`
}

// UserGroupAddEvent Event sent to users in an organization when a [user group] is created.
//
// [user group]: https://zulip.com/help/user-groups
type UserGroupAddEvent struct {
	EventCommonWithOp

	Group UserGroup `json:"group,omitempty"`
}

// UserGroupRemoveEvent Event sent when a user group is deactivated but only to clients with `include_deactivated_groups` client capability set to `false`.
//
//	**Changes**: Prior to Zulip 10.0 (feature level 294), this event was sent when a user group was deleted.
type UserGroupRemoveEvent struct {
	EventCommonWithOp

	// The Id of the group which has been deleted.
	GroupId int64 `json:"group_id,omitempty"`
}

// UserGroupUpdateEvent Event sent to all users in a Zulip organization when a property of a user group is changed.  For group deactivation, this event is only sent if `include_deactivated_groups` client capability is set to `true`.  This event is also sent when deactivating or reactivating a user for settings set to anonymous user groups which the user is direct member of. When deactivating the user, event is only sent to users who cannot access the deactivated user.
//
//	**Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when deactivating or reactivating a user.  Prior to Zulip 10.0 (feature level 294), this event was sent to all clients when a user group was deactivated.
type UserGroupUpdateEvent struct {
	EventCommonWithOp

	// The Id of the user group whose details have changed.
	GroupId int64               `json:"group_id,omitempty"`
	Data    UserGroupUpdateData `json:"data,omitempty"`
}

// UserGroupUpdate Dictionary containing the changed details of the user group.
type UserGroupUpdateData struct {
	// The new name of the user group. Only present if the group's name changed.
	Name *string `json:"name,omitempty"`
	// The new description of the group. Only present if the description changed.
	Description           *string            `json:"description,omitempty"`
	CanAddMembersGroup    *GroupSettingValue `json:"can_add_members_group,omitempty"`
	CanJoinGroup          *GroupSettingValue `json:"can_join_group,omitempty"`
	CanLeaveGroup         *GroupSettingValue `json:"can_leave_group,omitempty"`
	CanManageGroup        *GroupSettingValue `json:"can_manage_group,omitempty"`
	CanMentionGroup       *GroupSettingValue `json:"can_mention_group,omitempty"`
	CanRemoveMembersGroup *GroupSettingValue `json:"can_remove_members_group,omitempty"`
	// Whether the user group is deactivated. Deactivated groups cannot be used as a subgroup of another group or used for any other purpose.
	//
	// **Changes**: New in Zulip 10.0 (feature level 290).
	Deactivated *bool `json:"deactivated,omitempty"`
}

// UserGroupAddMembersEvent Event sent to all users when users have been added to a user group.  This event is also sent when reactivating a user for all the user groups the reactivated user was a member of before being deactivated.
//
//	**Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when reactivating a user.
type UserGroupMembersEvent struct {
	EventCommonWithOp

	// The Id of the user group with added/removed members.
	GroupId int64 `json:"group_id,omitempty"`
	// Array containing the Ids of the users who have been added/removed to the user group.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// UserGroupAddSubgroupsEvent Event sent to all users when subgroups have been added to a user group.
//
//	**Changes**: New in Zulip 6.0 (feature level 127).
type UserGroupSubgroupsEvent struct {
	EventCommonWithOp

	// The Id of the user group whose details have changed.
	GroupId int64 `json:"group_id,omitempty"`
	// Array containing the Ids of the subgroups that have been added/removed to the user group.
	//
	// **Changes**: New in Zulip 6.0 (feature level 131). Previously, this was called `subgroup_ids`, but clients can ignore older events as this feature level predates subgroups being fully implemented.
	DirectSubgroupIds []int64 `json:"direct_subgroup_ids,omitempty"`
}

// SubscriptionAddEvent Event sent to a user's clients when that user's channel subscriptions have changed (either the set of subscriptions or their properties).
type SubscriptionAddEvent struct {
	EventCommonWithOp

	// A list of dictionaries where each dictionary contains information about one of the subscribed channels.
	//
	// **Changes**: Removed `email_address` field from the dictionary in Zulip 8.0 (feature level 226).  Removed `role` field from the dictionary in Zulip 6.0 (feature level 133).
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// SubscriptionRemoveEvent Event sent to a user's clients when that user has been unsubscribed from one or more channels.
type SubscriptionRemoveEvent struct {
	EventCommonWithOp

	// A list of dictionaries, where each dictionary contains information about one of the newly unsubscribed channels.
	Subscriptions []SubscriptionRemoveData `json:"subscriptions,omitempty"`
}

// SubscriptionRemoveData Dictionary containing details about the unsubscribed channel.
type SubscriptionRemoveData struct {
	// The Id of the channel.
	ChannelId int64 `json:"stream_id,omitempty"`
	// The name of the channel.
	Name string `json:"name,omitempty"`
}

// RealmLinkifiersEvent Event sent to all users in a Zulip organization when the set of configured [linkifiers] for the organization has changed.  Processing this event is important for doing Markdown local echo correctly.  Clients will not receive this event unless the event queue is registered with the client capability `{"linkifier_url_template": true}`. See [`POST /register`], the `linkifier_url_template` client capability was not required. The requirement was added because linkifiers were updated to contain a URL template instead of a URL format string, which was not a backwards-compatible change.  New in Zulip 4.0 (feature level 54), replacing the deprecated `realm_filters` event type.
//
//	**Changes**: Before Zulip 7.0 (feature level 176
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
// [`POST /register`]: https://zulip.com/api/register-queue#parameter-client_capabilities for how client capabilities can be specified.
type RealmLinkifiersEvent struct {
	EventCommon
	// An ordered array of dictionaries where each dictionary contains details about a single linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [`PATCH /realm/linkifiers`].
	//
	// [`PATCH /realm/linkifiers`]: https://zulip.com/api/reorder-linkifiers
	RealmLinkifiers []RealmLinkifiers `json:"realm_linkifiers,omitempty"`
}

// RealmPlaygroundsEvent Event sent to all users in a Zulip organization when the set of configured [code playgrounds] for the organization has changed.
//
//	**Changes**: New in Zulip 4.0 (feature level 49).
//
// [code playgrounds]: https://zulip.com/help/code-blocks#code-playgrounds
type RealmPlaygroundsEvent struct {
	EventCommon
	// An array of dictionaries where each dictionary contains data about a single playground entry.
	RealmPlaygrounds []RealmPlayground `json:"realm_playgrounds,omitempty"`
}

// RealmDomainsAddEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsAddEvent struct {
	EventCommonWithOp

	RealmDomain *RealmDomain `json:"realm_domain,omitempty"`
}

// RealmFiltersEvent Legacy event type that is no longer sent to clients. Previously, sent to all users in a Zulip organization when the set of configured [linkifiers] for the organization was changed.
//
//	**Changes**: Prior to Zulip 7.0 (feature level 176), this event type was sent to clients.  **Deprecated** in Zulip 4.0 (feature level 54), and replaced by the `realm_linkifiers` event type, which has a clearer name and format.
//
// # Deprecated
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
type RealmFiltersEvent struct {
	EventCommon
	// An array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example `"#(?P<id>[123])"`. The second element was the URL format string that the pattern should be linkified with. A URL format string for the above example would be `"https://realm.com/my_realm_filter/%(id)s"`. And the third element was the Id of the realm filter.
	RealmFilters []interface{}
}

// RealmEmojiUpdateEvent Event sent to all users in a Zulip organization when a [custom emoji] has been updated, typically when a new emoji has been added or an old one has been deactivated. The event contains all custom emoji configured for the organization, not just the updated custom emoji.
//
// [custom emoji]: https://zulip.com/help/custom-emoji
type RealmEmojiUpdateEvent struct {
	EventCommonWithOp

	// An object in which each key describes a realm emoji.
	RealmEmoji map[string]RealmEmoji `json:"realm_emoji,omitempty"`
}

// RealmDomainsChangeEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsChangeEvent struct {
	EventCommonWithOp

	RealmDomain RealmDomain `json:"realm_domain,omitempty"`
}

// RealmDomainsRemoveEvent Event sent to all users in a Zulip organization when the set of [allowed domains for new users] has changed.
//
// [allowed domains for new users]: https://zulip.com/help/restrict-account-creation#configuring-email-domain-restrictions
type RealmDomainsRemoveEvent struct {
	EventCommonWithOp

	// The domain to be removed.
	Domain string `json:"domain,omitempty"`
}

// RealmExportEvent Event sent to the user who requested a [data export] when the status of the data export changes.
//
// [data export]: https://zulip.com/help/export-your-organization
type RealmExportEvent struct {
	EventCommon
	// An array of dictionaries where each dictionary contains details about a data export of the organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 304), `export_type` parameter was not present as only public data export was supported via API.
	Exports []RealmExport `json:"exports,omitempty"`
}

// RealmExportConsentEvent Event sent to administrators when the [data export consent] status for a user changes, whether due to a user changing their consent preferences or a user being created or reactivated (since user creation/activation events do not contain these data).
//
// **Changes**: New in Zulip 10.0 (feature level 312). Previously, there was not event available to administrators with these data.
//
// [data export consent]: https://zulip.com/help/export-your-organization#configure-whether-administrators-can-export-your-private-data
type RealmExportConsentEvent struct {
	EventCommon

	ExportConsent
}

// RealmBotAddEvent Event sent to users who can administer a newly created bot user. Clients will also receive a `realm_user` event that contains basic details (but not the API key).  The `realm_user` events are sufficient for clients that only need to interact with the bot; this `realm_bot` event type is relevant only for administering bots.  Only organization administrators and the user who owns the bot will receive this event.
type RealmBotAddEvent struct {
	EventCommonWithOp

	Bot Bot `json:"bot,omitempty"`
}

// RealmBotUpdateEvent Event sent to users who can administer a bot user when the bot is configured. Clients may also receive a `realm_user` event that for changes in public data about the bot (name, etc.).  The `realm_user` events are sufficient for clients that only need to interact with the bot; this `realm_bot` event type is relevant only for administering bots.  Only organization administrators and the user who owns the bot will receive this event.
type RealmBotUpdateEvent struct {
	EventCommonWithOp

	// Object containing details about the changed bot. It contains two properties: the user Id of the bot and the property to be changed. The changed property is one of the remaining properties listed below.
	Bot Bot `json:"bot,omitempty"`
}

// RealmBotDeleteEvent Event sent to all users when a bot has been deactivated. Note that this is very similar to the bot_remove event and one of them will be removed soon.
type RealmBotDeleteEvent struct {
	EventCommonWithOp

	Bot UserInfo `json:"bot,omitempty"`
}

// RealmUpdateEvent The simpler of two possible event types sent to all users in a Zulip organization when the configuration of the organization (realm) has changed.  Often individual settings are migrated from this format to the [realm/update_dict] event format when additional realm settings are added whose values are coupled to each other in some way. The specific values supported by this event type are documented in the [realm/update_dict] documentation.  A correct client implementation should convert these events into the corresponding [realm/update_dict] event and then process that.
//
//	**Changes**: Removed `extra_data` optional property in Zulip 10.0 (feature level 306). The `extra_data` used to include an `upload_quota` field when changed property was `plan_type`. The server now sends a standard `realm/update_dict` event for plan changes.
//
// [realm/update_dict]: #realm-update_dict
type RealmUpdateEvent struct {
	EventCommonWithOp

	// The name of the property that was changed.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}

// RealmDeactivatedEvent Event sent to all users in a Zulip organization when the organization (realm) is deactivated. Its main purpose is to flush active longpolling connections so clients can immediately show the organization as deactivated.  Clients cannot rely on receiving this event, because they will no longer be able to authenticate to the Zulip API due to the deactivation, and thus can miss it if they did not have an active longpolling connection at the moment of deactivation.  Correct handling of realm deactivations requires that clients parse authentication errors from GET /events; if that is done correctly, the client can ignore this event type and rely on its handling of the `GET /events` request it will do immediately after processing this batch of events.
type RealmDeactivatedEvent struct {
	EventCommonWithOp

	// The Id of the deactivated realm.
	RealmId int64 `json:"realm_id,omitempty"`
}

// RestartEvent Event sent to all the users whenever the Zulip server restarts.  Specifically, this event is sent whenever the Tornado process for the user is restarted; in particular, this will always happen when the Zulip server is upgraded.  Clients should use this event to update their tracking of the server's capabilities, and to decide if they wish to get a new event queue after a server upgrade. Clients doing so must implement a random delay strategy to spread such restarts over 5 minutes or more to avoid creating a synchronized thundering herd effect.
//
//	**Changes**: Removed the `immediate` flag, which was only used by web clients in development, in Zulip 9.0 (feature level 240).
type RestartEvent struct {
	EventCommonWithOp

	// The Zulip version number, in the format where this appears in the [server_settings] and [register] responses.
	//
	// **Changes**: New in Zulip 4.0 (feature level 59).
	//
	// [server_settings]: https://zulip.com/api/get-server-settings
	// [register]: https://zulip.com/api/register-queue
	ZulipVersion string `json:"zulip_version,omitempty"`
	// The Zulip merge base number, in the format where this appears in the [server_settings] and [register] responses.
	//
	// **Changes**: New in Zulip 5.0 (feature level 88).
	//
	// [server_settings]: https://zulip.com/api/get-server-settings
	// [register]: https://zulip.com/api/register-queue
	ZulipMergeBase string `json:"zulip_merge_base,omitempty"`
	// The [Zulip feature level] of the server after the restart.  Clients should use this to update their tracking of the server's capabilities, and may choose to refetch their state and create a new event queue when the API feature level has changed in a way that the client finds significant. Clients choosing to do so must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.
	//
	// **Changes**: New in Zulip 4.0 (feature level 59).
	//
	// [Zulip feature level]: https://zulip.com/api/changelog
	ZulipFeatureLevel int32 `json:"zulip_feature_level,omitempty"`
	// The timestamp at which the server started.
	ServerGeneration time.Time `json:"server_generation,omitempty"`
}

type restartEventJSON struct {
	EventCommonWithOp

	ZulipVersion      string  `json:"zulip_version,omitempty"`
	ZulipMergeBase    string  `json:"zulip_merge_base,omitempty"`
	ZulipFeatureLevel int32   `json:"zulip_feature_level,omitempty"`
	ServerGeneration  float64 `json:"server_generation,omitempty"`
}

// UnmarshalJSON Custom unmarshaler to handle conversion of `server_generation` from a float timestamp to a `time.Time` object.
func (e *RestartEvent) UnmarshalJSON(data []byte) error {
	var rej restartEventJSON
	if err := json.Unmarshal(data, &rej); err != nil {
		return err
	}

	e.EventCommonWithOp = rej.EventCommonWithOp
	e.ZulipVersion = rej.ZulipVersion
	e.ZulipMergeBase = rej.ZulipMergeBase
	e.ZulipFeatureLevel = rej.ZulipFeatureLevel
	e.ServerGeneration = time.Unix(int64(rej.ServerGeneration), 0)

	return nil
}

func (e *RestartEvent) MarshalJSON() ([]byte, error) {
	rej := restartEventJSON{
		EventCommonWithOp: e.EventCommonWithOp,
		ZulipVersion:      e.ZulipVersion,
		ZulipMergeBase:    e.ZulipMergeBase,
		ZulipFeatureLevel: e.ZulipFeatureLevel,
		ServerGeneration:  float64(e.ServerGeneration.Unix()),
	}
	return json.Marshal(rej)
}

// WebReloadClientEvent An event which signals the official Zulip web/desktop app to update, by reloading the page and fetching a new queue; this will generally follow a `restart` event. Clients which do not obtain their code from the server (e.g. mobile and terminal clients, which store their code locally) should ignore this event.  Clients choosing to reload the application must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.
//
//	**Changes**: New in Zulip 9.0 (feature level 240).
type WebReloadClientEvent struct {
	EventCommon
	// Whether the client should fetch a new event queue immediately, rather than using a backoff strategy to avoid thundering herds. A Zulip development server uses this parameter to reload clients immediately.
	Immediate bool `json:"immediate,omitempty"`
}

// RealmUpdateDictEvent The more general of two event types that may be used when sending an event to all users in a Zulip organization when the configuration of the organization (realm) has changed.  Unlike the simpler [realm/update] event format, this event type supports multiple properties being changed in a single event.  This event is also sent when deactivating or reactivating a user for settings set to anonymous user groups which the user is direct member of. When deactivating the user, event is only sent to users who cannot access the deactivated user.
//
//	**Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when deactivating or reactivating a user.  In Zulip 7.0 (feature level 163), the realm setting `email_address_visibility` was removed. It was replaced by a [user setting](https://zulip.com/api/update-settings#parameter-email_address_visibility with a [realm user default], with the encoding of different values preserved. Clients can support all versions by supporting the current API and treating every user as having the realm's `email_address_visibility` value.
//
// [realm user default]: https://zulip.com/api/update-realm-user-settings-defaults#parameter-email_address_visibility
// [realm/update]: #realm-update
type RealmUpdateDictEvent struct {
	EventCommonWithOp

	// Always `"default"`. Present for backwards-compatibility with older clients that predate the `update_dict` event style.  **Deprecated** and will be removed in a future release.
	// Deprecated
	Property           string             `json:"property,omitempty"`
	RealmConfiguration RealmConfiguration `json:"data,omitempty"`
}

// DraftsAddEvent Event containing details of newly created drafts.
type DraftsAddEvent struct {
	EventCommonWithOp

	// An array containing objects for the newly created drafts.
	Drafts []Draft `json:"drafts,omitempty"`
}

// DraftsUpdateEvent Event containing details for an edited draft.
type DraftsUpdateEvent struct {
	EventCommonWithOp

	Draft Draft `json:"draft,omitempty"`
}

// SubscriptionUpdateEvent Event sent to a user's clients when a property of the user's subscription to a channel has been updated. This event is used only for personal properties like `is_muted` or `pin_to_top`. See the [`stream op: update` event] for updates to global properties of a channel.
//
// [`stream op: update` event]: https://zulip.com/api/get-events#stream-update
type SubscriptionUpdateEvent struct {
	EventCommonWithOp

	ChannelId int64 `json:"stream_id,omitempty"`
	// The property of the subscription which has changed. For details on the various subscription properties that a user can change, see [POST /users/me/subscriptions/properties].  Clients should generally handle an unknown property received here without crashing, since that will naturally happen when connecting to a Zulip server running a new version that adds a new subscription property.
	//
	// **Changes**: As of Zulip 6.0 (feature level 139), updates to the `is_muted` property or the deprecated `in_home_view` property will send two `subscription` update events, one for each property, to support clients fully migrating to use the `is_muted` property. Prior to this feature level, updates to either property only sent one event with the deprecated `in_home_view` property.
	//
	// [POST /users/me/subscriptions/properties]: https://zulip.com/api/update-subscription-settings
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}

// DraftsRemoveEvent Event containing the Id of a deleted draft.
type DraftsRemoveEvent struct {
	EventCommonWithOp

	// The Id of the draft that was just deleted.
	DraftId int64 `json:"draft_id,omitempty"`
}

// RealmUserSettingsDefaultsUpdateEvent Event sent to all users in a Zulip organization when the [default settings for new users] of the organization (realm) have changed.
// See [PATCH /realm/user_settings_defaults] for details on possible properties.
// **Changes**: New in Zulip 5.0 (feature level 95).
//
// [default settings for new users]: https://zulip.com/help/configure-default-new-user-settings
// [PATCH /realm/user_settings_defaults]: https://zulip.com/api/update-realm-user-settings-defaults
type RealmUserSettingsDefaultsUpdateEvent struct {
	EventCommonWithOp

	// The name of the property that was changed.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}

// NavigationViewAddEvent Event containing details of a newly configured navigation view.
//
//	**Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewAddEvent struct {
	EventCommonWithOp

	NavigationView NavigationView `json:"navigation_view,omitempty"`
}

// NavigationViewUpdateEvent Event containing details of an update to an existing navigation view.
//
//	**Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewUpdateEvent struct {
	EventCommonWithOp

	// The unique URL hash of the navigation view being updated.
	Fragment string                   `json:"fragment,omitempty"`
	Data     NavigationViewUpdateData `json:"data,omitempty"`
}

// NavigationViewUpdateData A dictionary containing the updated properties of the navigation view.
type NavigationViewUpdateData struct {
	// The user-facing name for custom navigation views. Omit this field for built-in views.
	Name *string `json:"name,omitempty"`
	// Determines whether the view is pinned (true) or hidden in the menu (false).
	IsPinned *bool `json:"is_pinned,omitempty"`
}

// NavigationViewRemoveEvent Event containing the fragment of a deleted navigation view.
//
//	**Changes**: New in Zulip 11.0 (feature level 390).
type NavigationViewRemoveEvent struct {
	EventCommonWithOp

	// The unique URL hash of the navigation view that was just deleted.
	Fragment string `json:"fragment,omitempty"`
}

// SavedSnippetsEvent Event containing details of a newly created saved snippet.
//
//	**Changes**: New in Zulip 10.0 (feature level 297).
type SavedSnippetsEvent struct {
	EventCommonWithOp

	SavedSnippet SavedSnippet `json:"saved_snippet,omitempty"`
}

// SavedSnippetsRemoveEvent Event containing the Id of a deleted saved snippet.
//
//	**Changes**: New in Zulip 10.0 (feature level 297).
type SavedSnippetsRemoveEvent struct {
	EventCommonWithOp

	// The Id of the saved snippet that was just deleted.
	//
	// **Changes**: New in Zulip 10.0 (feature level 297).
	SavedSnippetId int64 `json:"saved_snippet_id,omitempty"`
}

// RemindersAddEvent Event sent to a user's clients when a reminder is scheduled.
//
//	**Changes**: New in Zulip 11.0 (feature level 399).
type RemindersAddEvent struct {
	EventCommonWithOp

	// An array of objects containing details of the newly created reminders.
	Reminders []Reminder `json:"reminders,omitempty"`
}

// RemindersRemoveEvent Event sent to a user's clients when a reminder is deleted.
//
//	**Changes**: New in Zulip 11.0 (feature level 399).
type RemindersRemoveEvent struct {
	EventCommonWithOp

	// The Id of the reminder that was deleted.
	ReminderId int64 `json:"reminder_id,omitempty"`
}

// SubscriptionPeerAddEvent Event sent when another user subscribes to a channel, or their subscription is newly visible to the current user.  When a user subscribes to a channel, the current user will receive this event only if they [have permission to see the channel's subscriber list]. When the current user gains permission to see a given channel's subscriber list, they will receive this event for the existing subscriptions to the channel.
//
//	**Changes**: Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users when subscribers to web-public channels and subscribed public channels changed.  Prior to Zulip 8.0 (feature level 205), this event was not sent when a user gained access to a channel due to their [role changing].  Prior to Zulip 6.0 (feature level 134), this event was not sent when a private channel was made public.  In Zulip 4.0 (feature level 35), the singular `user_id` and `stream_id` integers included in this event were replaced with plural `user_ids` and `stream_ids` integer arrays.  In Zulip 3.0 (feature level 19), the `stream_id` field was added to identify the channel the user subscribed to, replacing the `name` field.
//
// [have permission to see the channel's subscriber list]: https://zulip.com/help/channel-permissions
//
// [role changing]: https://zulip.com/help/user-roles
type SubscriptionPeerAddEvent struct {
	EventCommonWithOp

	// The Ids of channels that have new or updated subscriber data.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `stream_id` integer.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
	// The Ids of the users who are newly visible as subscribed to the specified channels.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `user_id` integer.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// ScheduledMessagesAddEvent Event sent to a user's clients when scheduled messages are created.
//
//	**Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesAddEvent struct {
	EventCommonWithOp

	// An array of objects containing details of the newly created scheduled messages.
	ScheduledMessages []ScheduledMessage `json:"scheduled_messages,omitempty"`
}

// ScheduledMessagesUpdateEvent Event sent to a user's clients when a scheduled message is edited.
//
//	**Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesUpdateEvent struct {
	EventCommonWithOp

	ScheduledMessage ScheduledMessage `json:"scheduled_message,omitempty"`
}

// ScheduledMessagesRemoveEvent Event sent to a user's clients when a scheduled message is deleted.
//
//	**Changes**: New in Zulip 7.0 (feature level 179).
type ScheduledMessagesRemoveEvent struct {
	EventCommonWithOp

	// The Id of the scheduled message that was deleted.
	ScheduledMessageId int64 `json:"scheduled_message_id,omitempty"`
}

// ChannelFolderAddEvent Event sent to users in an organization when a channel folder is created.
//
//	**Changes**: New in Zulip 11.0 (feature level 389).
type ChannelFolderAddEvent struct {
	EventCommonWithOp

	ChannelFolder ChannelFolder `json:"channel_folder,omitempty"`
}

// ChannelFolderUpdateEvent Event sent to users in an organization when a channel folder is updated.
//
//	**Changes**: New in Zulip 11.0 (feature level 389).
type ChannelFolderUpdateEvent struct {
	EventCommonWithOp

	// Id of the updated channel folder.
	ChannelFolderId int64            `json:"channel_folder_id,omitempty"`
	Data            FolderUpdateData `json:"data,omitempty"`
}

// FolderUpdateData Dictionary containing the changed details of the channel folder.
type FolderUpdateData struct {
	// The new name of the channel folder. Only present if the channel folder's name changed.
	Name *string `json:"name,omitempty"`
	// The new description of the channel folder. Only present if the description changed.  See [Markdown message formatting] for details on Zulip's HTML format.
	//
	// [Markdown message formatting]: https://zulip.com/api/message-formatting
	Description *string `json:"description,omitempty"`
	// The new rendered description of the channel folder. Only present if the description changed.
	RenderedDescription *string `json:"rendered_description,omitempty"`
	// Whether the channel folder is archived or not. Only present if the channel folder is archived or unarchived.
	IsArchived *bool `json:"is_archived,omitempty"`
}

// ChannelFolderReorderEvent Event sent to users in an organization when channel folders are reordered.
//
//	**Changes**: New in Zulip 11.0 (feature level 418).
type ChannelFolderReorderEvent struct {
	EventCommonWithOp

	// A list of channel folder Ids representing the new order.
	Order []int64 `json:"order,omitempty"`
}

// SubscriptionPeerRemoveEvent Event sent to other users when users have been unsubscribed from channels. Sent to all users if the channel is public or to only the existing subscribers if the channel is private.
//
//	**Changes**: Prior to Zulip 8.0 (feature level 220), this event was incorrectly not sent to guest users when subscribers to web-public channels and subscribed public channels changed.  In Zulip 4.0 (feature level 35), the singular `user_id` and `stream_id` integers included in this event were replaced with plural `user_ids` and `stream_ids` integer arrays.  In Zulip 3.0 (feature level 19), the `stream_id` field was added to identify the channel the user unsubscribed from, replacing the `name` field.
type SubscriptionPeerRemoveEvent struct {
	EventCommonWithOp

	// The Ids of the channels from which the users have been unsubscribed from.  When a user is deactivated, the server will send this event removing the user's subscriptions before the `realm_user` event for the user's deactivation.
	//
	// **Changes**: Before Zulip 10.0 (feature level 377), this event was not sent on user deactivation. Clients supporting older server versions and maintaining peer subscriber data need to remove all channel subscriptions for a user when processing the `realm_user` event with `op="remove"`.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `stream_id` integer.
	ChannelIds []int64 `json:"stream_ids,omitempty"`
	// The Ids of the users who have been unsubscribed.
	//
	// **Changes**: New in Zulip 4.0 (feature level 35), replacing the `user_id` integer.
	UserIds []int64 `json:"user_ids,omitempty"`
}
