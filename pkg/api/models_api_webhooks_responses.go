package api

import "time"

// ZulipOutgoingWebhooksResponse This is an example of the JSON payload that the Zulip server will `POST` to your server:
type ZulipOutgoingWebhooksResponse struct {
	// Email of the bot user.
	BotEmail string `json:"bot_email,omitempty"`
	// The full name of the bot user.
	BotFullName string `json:"bot_full_name,omitempty"`
	// The message content, in raw [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown) format (not rendered to HTML).
	Data string `json:"data,omitempty"`
	// What aspect of the message triggered the outgoing webhook notification. Possible values include `direct_message` and `mention`.  **Changes**: In Zulip 8.0 (feature level 201), renamed the trigger `private_message` to `direct_message`.
	Trigger WebhookTrigger `json:"trigger,omitempty"`
	// A string of alphanumeric characters that can be used to authenticate the webhook request (each bot user uses a fixed token). You can get the token used by a given outgoing webhook bot in the `zuliprc` file downloaded when creating the bot.
	Token   string                  `json:"token,omitempty"`
	Message OutgoingWebhooksMessage `json:"message,omitempty"`
}

type OutgoingWebhooksMessage struct {
	AvatarUrl interface{} `json:"avatar_url,omitempty"`
	// A Zulip \"client\" string, describing what Zulip client sent the message.
	Client string `json:"client,omitempty"`
	// The content/body of the message. When `apply_markdown` is set, it will be in HTML format.  See [Markdown message formatting](zulip.com/api/message-formatting) for details on Zulip's HTML format.
	Content string `json:"content,omitempty"`
	// The HTTP `content_type` for the message content. This will be `text/html` or `text/x-markdown`, depending on whether `apply_markdown` was set.  See the help center article on [message formatting](zulip.com/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.
	ContentType      string           `json:"content_type,omitempty"`
	DisplayRecipient DisplayRecipient `json:"display_recipient,omitempty"`
	// An array of objects, with each object documenting the changes in a previous edit made to the message, ordered chronologically from most recent to least recent edit.  Not present if the message has never been edited or moved, or if [viewing message edit history][edit-history-access] is not allowed in the organization.  Every object will contain `user_id` and `timestamp`.  The other fields are optional, and will be present or not depending on whether the channel, topic, and/or message content were modified in the edit event. For example, if only the topic was edited, only `prev_topic` and `topic` will be present in addition to `user_id` and `timestamp`.  [edit-history-access]: /help/restrict-message-edit-history-access  **Changes**: In Zulip 10.0 (feature level 284), removed the `prev_rendered_content_version` field as it is an internal server implementation detail not used by any client.
	EditHistory []EditHistory `json:"edit_history,omitempty"`
	// The unique message Id. Messages should always be displayed sorted by Id.
	Id int64 `json:"id,omitempty"`
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
	StreamId int64 `json:"stream_id,omitempty"`
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
	// The content/body of the message rendered in HTML.  See [Markdown message formatting](zulip.com/api/message-formatting) for details on Zulip's HTML format.
	RenderedContent string `json:"rendered_content,omitempty"`
}
