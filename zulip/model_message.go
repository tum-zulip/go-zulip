package zulip

import (
	"encoding/json"
	"time"
)

// Message An object containing details of the message.  **Changes**: New in Zulip 5.0 (feature level 120).
type Message struct {
	// The unique message Id. Messages should always be displayed sorted by Id.
	Id int64 `json:"id,omitempty"`

	AvatarUrl *string `json:"avatar_url,omitempty"`
	// A Zulip \"client\" string, describing what Zulip client sent the message.
	Client string `json:"client,omitempty"`
	// The content/body of the message. When `apply_markdown` is set, it will be in HTML format.  See [Markdown message formatting](zulip.com/api/message-formatting) for details on Zulip's HTML format.
	Content string `json:"content,omitempty"`
	// The HTTP `content_type` for the message content. This will be `text/html` or `text/x-markdown`, depending on whether `apply_markdown` was set.  See the help center article on [message formatting](zulip.com/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.
	ContentType      string           `json:"content_type,omitempty"`
	DisplayRecipient DisplayRecipient `json:"display_recipient,omitempty"`
	// An array of objects, with each object documenting the changes in a previous edit made to the message, ordered chronologically from most recent to least recent edit.  Not present if the message has never been edited or moved, or if [viewing message edit history][edit-history-access] is not allowed in the organization.  Every object will contain `user_id` and `timestamp`.  The other fields are optional, and will be present or not depending on whether the channel, topic, and/or message content were modified in the edit event. For example, if only the topic was edited, only `prev_topic` and `topic` will be present in addition to `user_id` and `timestamp`.  [edit-history-access]: /help/restrict-message-edit-history-access  **Changes**: In Zulip 10.0 (feature level 284), removed the `prev_rendered_content_version` field as it is an internal server implementation detail not used by any client.
	EditHistory []EditHistory `json:"edit_history,omitempty"`
	// Whether the message is a [/me status message][status-messages]  [status-messages]: /help/format-your-message-using-markdown#status-messages
	IsMeMessage bool `json:"is_me_message,omitempty"`
	// The UNIX timestamp for when the message's content was last edited, in UTC seconds.  Not present if the message's content has never been edited.  Clients should use this field, rather than parsing the `edit_history` array, to display an indicator that the message has been edited.  **Changes**: Prior to Zulip 10.0 (feature level 365), this was the time when the message was last edited or moved.
	LastEditTimestamp time.Time `json:"last_edit_timestamp,omitempty"`
	// The UNIX timestamp for when the message was last moved to a different channel or topic, in UTC seconds.  Not present if the message has never been moved, or if the only topic moves for the message are [resolving or unresolving](zulip.com/help/resolve-a-topic) the message's topic.  Clients should use this field, rather than parsing the `edit_history` array, to display an indicator that the message has been moved.  **Changes**: New in Zulip 10.0 (feature level 365). Previously, parsing the `edit_history` array was required in order to correctly display moved message indicators.
	LastMovedTimestamp time.Time `json:"last_moved_timestamp,omitempty"`
	// Data on any reactions to the message.
	Reactions []EmojiReaction `json:"reactions,omitempty"`
	// A unique Id for the set of users receiving the message (either a channel or group of users). Useful primarily for hashing.  **Changes**: Before Zulip 10.0 (feature level 327), `recipient_id` was the same across all incoming 1:1 direct messages. Now, each incoming message uniquely shares a `recipient_id` with outgoing messages in the same conversation.
	RecipientId int64 `json:"recipient_id,omitempty"`
	// The Zulip API email address of the message's sender.
	SenderEmail string `json:"sender_email,omitempty"`
	// The full name of the message's sender.
	SenderFullName string `json:"sender_full_name,omitempty"`
	// The user Id of the message's sender.
	SenderId int64 `json:"sender_id,omitempty"`
	// A string identifier for the realm the sender is in. Unique only within the context of a given Zulip server.  E.g. on `example.zulip.com`, this will be `example`.
	SenderRealmStr string `json:"sender_realm_str,omitempty"`
	// Only present for channel messages; the Id of the channel.
	StreamId *int64 `json:"stream_id,omitempty"`
	// The `topic` of the message. Currently always `\"\"` for direct messages, though this could change if Zulip adds support for topics in direct message conversations.  The field name is a legacy holdover from when topics were called \"subjects\" and will eventually change.  For clients that don't support the `empty_topic_name` [client capability][client-capabilities], the empty string value is replaced with the value of `realm_empty_topic_display_name` found in the [POST /register](zulip.com/api/register-queue) response, for channel messages.  **Changes**: Before Zulip 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities
	Subject string `json:"subject,omitempty"`
	// Data used for certain experimental Zulip integrations.
	Submessages []Submessage `json:"submessages,omitempty"`
	// The UNIX timestamp for when the message was sent, in UTC seconds.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Data on any links to be included in the `topic` line (these are generated by [custom linkification filters](zulip.com/help/add-a-custom-linkifier) that match content in the message's topic.)  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called `subject_links`; clients are recommended to rename `subject_links` to `topic_links` if present for compatibility with older Zulip servers.
	TopicLinks []TopicLink `json:"topic_links,omitempty"`
	// The type of the message: `\"stream\"` or `\"private\"`.
	Type RecipientType `json:"type,omitempty"`
	// The user's [message flags][message-flags] for the message.  **Changes**: In Zulip 8.0 (feature level 224), the `wildcard_mentioned` flag was deprecated in favor of the `stream_wildcard_mentioned` and `topic_wildcard_mentioned` flags. The `wildcard_mentioned` flag exists for backwards compatibility with older clients and equals `stream_wildcard_mentioned || topic_wildcard_mentioned`. Clients supporting older server versions should treat this field as a previous name for the `stream_wildcard_mentioned` flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags
	Flags []string `json:"flags,omitempty"`
}

type messageJSON struct {
	Id                 int64            `json:"id,omitempty"`
	AvatarUrl          *string          `json:"avatar_url,omitempty"`
	Client             string           `json:"client,omitempty"`
	Content            string           `json:"content,omitempty"`
	ContentType        string           `json:"content_type,omitempty"`
	DisplayRecipient   DisplayRecipient `json:"display_recipient,omitempty"`
	EditHistory        []EditHistory    `json:"edit_history,omitempty"`
	IsMeMessage        bool             `json:"is_me_message,omitempty"`
	LastEditTimestamp  int64            `json:"last_edit_timestamp,omitempty"`
	LastMovedTimestamp int64            `json:"last_moved_timestamp,omitempty"`
	Reactions          []EmojiReaction  `json:"reactions,omitempty"`
	RecipientId        int64            `json:"recipient_id,omitempty"`
	SenderEmail        string           `json:"sender_email,omitempty"`
	SenderFullName     string           `json:"sender_full_name,omitempty"`
	SenderId           int64            `json:"sender_id,omitempty"`
	SenderRealmStr     string           `json:"sender_realm_str,omitempty"`
	StreamId           *int64           `json:"stream_id,omitempty"`
	Subject            string           `json:"subject,omitempty"`
	Submessages        []Submessage     `json:"submessages,omitempty"`
	Timestamp          int64            `json:"timestamp,omitempty"`
	TopicLinks         []TopicLink      `json:"topic_links,omitempty"`
	Type               RecipientType    `json:"type,omitempty"`
	Flags              []string         `json:"flags,omitempty"`
}

func (o Message) MarshalJSON() ([]byte, error) {
	aux := messageJSON{
		Id:               o.Id,
		AvatarUrl:        o.AvatarUrl,
		Client:           o.Client,
		Content:          o.Content,
		ContentType:      o.ContentType,
		DisplayRecipient: o.DisplayRecipient,
		EditHistory:      o.EditHistory,
		IsMeMessage:      o.IsMeMessage,
		Reactions:        o.Reactions,
		RecipientId:      o.RecipientId,
		SenderEmail:      o.SenderEmail,
		SenderFullName:   o.SenderFullName,
		SenderId:         o.SenderId,
		SenderRealmStr:   o.SenderRealmStr,
		StreamId:         o.StreamId,
		Subject:          o.Subject,
		Submessages:      o.Submessages,
		TopicLinks:       o.TopicLinks,
		Type:             o.Type,
		Flags:            o.Flags,
	}

	if !o.LastEditTimestamp.IsZero() {
		aux.LastEditTimestamp = o.LastEditTimestamp.Unix()
	}

	if !o.LastMovedTimestamp.IsZero() {
		aux.LastMovedTimestamp = o.LastMovedTimestamp.Unix()
	}

	if !o.Timestamp.IsZero() {
		aux.Timestamp = o.Timestamp.Unix()
	}

	return json.Marshal(aux)
}

func (o *Message) UnmarshalJSON(data []byte) error {
	var aux messageJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	o.Id = aux.Id
	o.AvatarUrl = aux.AvatarUrl
	o.Client = aux.Client
	o.Content = aux.Content
	o.ContentType = aux.ContentType
	o.DisplayRecipient = aux.DisplayRecipient
	o.EditHistory = aux.EditHistory
	o.IsMeMessage = aux.IsMeMessage
	if aux.LastEditTimestamp != 0 {
		o.LastEditTimestamp = time.Unix(aux.LastEditTimestamp, 0).UTC()
	} else {
		o.LastEditTimestamp = time.Time{}
	}
	if aux.LastMovedTimestamp != 0 {
		o.LastMovedTimestamp = time.Unix(aux.LastMovedTimestamp, 0).UTC()
	} else {
		o.LastMovedTimestamp = time.Time{}
	}
	o.Reactions = aux.Reactions
	o.RecipientId = aux.RecipientId
	o.SenderEmail = aux.SenderEmail
	o.SenderFullName = aux.SenderFullName
	o.SenderId = aux.SenderId
	o.SenderRealmStr = aux.SenderRealmStr
	o.StreamId = aux.StreamId
	o.Subject = aux.Subject
	o.Submessages = aux.Submessages
	if aux.Timestamp != 0 {
		o.Timestamp = time.Unix(aux.Timestamp, 0).UTC()
	} else {
		o.Timestamp = time.Time{}
	}
	o.TopicLinks = aux.TopicLinks
	o.Type = aux.Type
	o.Flags = aux.Flags

	return nil
}

// DisplayRecipient - Data on the recipient of the message; either the name of a channel or a dictionary containing basic data on the users who received the message.
type DisplayRecipient struct {
	UserRecipents *[]UserRecipent
	Name          *string
}

// UserRecipent struct for UserRecipent
type UserRecipent struct {
	// Id of the user.
	Id *int64 `json:"id,omitempty"`
	// Zulip API email of the user.
	Email *string `json:"email,omitempty"`
	// Full name of the user.
	FullName *string `json:"full_name,omitempty"`
	// Whether the user is a mirror dummy.
	IsMirrorDummy *bool `json:"is_mirror_dummy,omitempty"`
}

// Submessage struct for Submessage
type Submessage struct {
	// The type of the message.
	MsgType string `json:"msg_type,omitempty"`
	// The new content of the submessage.
	Content string `json:"content,omitempty"`
	// The Id of the message to which the submessage has been added.
	MessageId int64 `json:"message_id,omitempty"`
	// The Id of the user who sent the message.
	SenderId int64 `json:"sender_id,omitempty"`
	// The Id of the submessage.
	Id int64 `json:"id,omitempty"`
}

// TopicLink struct for TopicLink
type TopicLink struct {
	// The original link text present in the topic.
	Text string `json:"text,omitempty"`
	// The expanded target url which the link points to.
	Url string `json:"url,omitempty"`
}

type MessageWithOptMatch struct {
	Message

	// Only present if keyword search was included among the narrow parameters.  HTML content of a queried message that matches the narrow, with `<span class=\"highlight\">` elements wrapping the matches for the search keywords.
	MatchContent *string `json:"match_content,omitempty"`
	// Only present if keyword search was included among the narrow parameters.  HTML-escaped topic of a queried message that matches the narrow, with `<span class=\"highlight\">` elements wrapping the matches for the search keywords.
	MatchSubject *string `json:"match_subject,omitempty"`
}

// []DisplayRecipientFromUserRecipentArray is a convenience function that returns []UserRecipent wrapped in DisplayRecipient
func DisplayRecipientFromUserRecipentArray(v []UserRecipent) DisplayRecipient {
	return DisplayRecipient{
		UserRecipents: &v,
	}
}

// ChannelNameAsDisplayRecipient is a convenience function that returns string wrapped in DisplayRecipient
func ChannelNameAsDisplayRecipient(v *string) DisplayRecipient {
	return DisplayRecipient{
		Name: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DisplayRecipient) UnmarshalJSON(data []byte) error {
	return UnionUnmarshalJSON(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DisplayRecipient) MarshalJSON() ([]byte, error) {
	return UnionMarshalJSON(src)
}
