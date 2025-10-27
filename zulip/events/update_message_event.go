package events

import (
	"encoding/json"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
)

// UpdateMessageEvent Event sent when a message's content, topic and/or channel has been edited or when a message's content has a rendering update, such as for an [inline URL preview]. Sent to all users who had received the original message.
//
// **Changes**: In Zulip 10.0 (feature level 284), removed the `prev_rendered_content_version` field as it is an internal server implementation detail not used by any client.
//
// [inline URL preview]: https://zulip.readthedocs.io/en/latest/subsystems/sending-messages.html#inline-url-previews
type UpdateMessageEvent struct {
	event
	// The Id of the user who sent the message.  Is `null` when event is for a rendering update of the original message, such as for an [inline URL preview].
	//
	// **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all `update_message` events. Previously, this field was omitted for [inline URL preview] updates.
	UserID *int64 `json:"user_id"`
	// Whether the event only updates the rendered content of the message.  This field should be used by clients to determine if the event only provides a rendering update to the message content, such as for an [inline URL preview]. When `true`, the event does not reflect a user-generated edit and does not modify the message history.
	//
	// **Changes**: New in Zulip 5.0 (feature level 114). Clients can correctly identify these rendering update event with earlier Zulip versions by checking whether the `user_id` field was omitted.
	RenderingOnly bool `json:"rendering_only"`
	// The Id of the message which was edited or updated.  This field should be used to apply content edits to the client's cached message history, or to apply rendered content updates.  If the channel or topic was changed, the set of moved messages is encoded in the separate `message_ids` field, which is guaranteed to include `message_id`.
	MessageID int64 `json:"message_id"`
	// A sorted list of Ids of messages to which any channel or topic changes encoded in this event should be applied.  This list always includes `message_id`, even when there are no channel or topic changes to apply.  These messages are guaranteed to have all been previously sent to channel `stream_id` with topic `orig_subject`, and have been moved to `new_stream_id` with topic `subject` (if those fields are present in the event).  Clients processing these events should update all cached message history associated with the moved messages (including adjusting `unread_msgs` data structures, where the client may not have the message itself in its history) to reflect the new channel and topic.  Content changes should be applied only to the single message indicated by `message_id`.
	//
	// **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.
	MessageIds []int64 `json:"message_ids"`
	// The user's personal [message flags] for the message with Id `message_id` following the edit.  A client application should compare these to the original flags to identify cases where a mention or alert word was added by the edit.
	//
	// **Changes**: In Zulip 8.0 (feature level 224), the `wildcard_mentioned` flag was deprecated in favor of the `stream_wildcard_mentioned` and `topic_wildcard_mentioned` flags. The `wildcard_mentioned` flag exists for backwards compatibility with older clients and equals `stream_wildcard_mentioned || topic_wildcard_mentioned`. Clients supporting older server versions should treat this field as a previous name for the `stream_wildcard_mentioned` flag as topic wildcard mentions were not available prior to this feature level.
	//
	// [message flags]: https://zulip.com/api/update-message-flags#available-flags
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
	ChannelID *int64 `json:"stream_id,omitempty"`
	// Only present if message(s) were moved to a different channel.  The post-edit channel for all of the messages with Ids in `message_ids`.
	NewChannelId *int64 `json:"new_stream_id,omitempty"`
	// Only present if this event moved messages to a different topic and/or channel.  The choice the editing user made about which messages should be affected by a channel/topic edit:  - `"change_one"`: Just change the one indicated in `message_id`. - `"change_later"`: Change messages in the same topic that had   been sent after this one. - `"change_all"`: Change all messages in that topic.  This parameter should be used to decide whether to change navigation and compose box state in response to the edit. For example, if the user was previously in topic narrow, and the topic was edited with `"change_later"` or `"change_all"`, the Zulip web app will automatically navigate to the new topic narrow. Similarly, a message being composed to the old topic should have its recipient changed to the new topic.  This navigation makes it much more convenient to move content between topics without disruption or messages continuing to be sent to the pre-edit topic by accident.
	PropagateMode *string `json:"propagate_mode,omitempty"`
	// Only present if this event moved messages to a different topic and/or channel.  The pre-edit topic for all of the messages with Ids in `message_ids`.  For clients that don't support the `empty_topic_name` [client capability], if the actual pre-edit topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	OrigSubject *string `json:"orig_subject,omitempty"`
	// Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  The post-edit topic for all of the messages with Ids in `message_ids`.  For clients that don't support the `empty_topic_name` [client capability], if the actual post-edit topic name is empty string, this field's value will instead be the value of `realm_empty_topic_display_name` found in the [`POST /register`] response.
	//
	// **Changes**: Before 10.0 (feature level 334), `empty_topic_name` client capability didn't exist and empty string as the topic name for channel messages wasn't allowed.
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Subject *string `json:"subject,omitempty"`
	// Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  Data on any links to be included in the `topic` line (these are generated by [custom linkification filter] that match content in the message's topic.), corresponding to the post-edit topic.
	//
	// **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called `subject_links`; clients are recommended to rename `subject_links` to `topic_links` if present for compatibility with older Zulip servers.
	//
	// [custom linkification filter]: https://zulip.com/help/add-a-custom-linkifier
	TopicLinks []zulip.TopicLink `json:"topic_links,omitempty"`
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

func (o *UpdateMessageEvent) MarshalJSON() ([]byte, error) {
	v := updateMessageEventJSON{
		UserID:              o.UserID,
		RenderingOnly:       o.RenderingOnly,
		MessageID:           o.MessageID,
		MessageIds:          o.MessageIds,
		Flags:               o.Flags,
		EditTimestamp:       o.EditTimestamp.UnixMilli(),
		ChannelName:         o.ChannelName,
		ChannelID:           o.ChannelID,
		NewChannelId:        o.NewChannelId,
		PropagateMode:       o.PropagateMode,
		OrigSubject:         o.OrigSubject,
		Subject:             o.Subject,
		TopicLinks:          o.TopicLinks,
		OrigContent:         o.OrigContent,
		OrigRenderedContent: o.OrigRenderedContent,
		Content:             o.Content,
		RenderedContent:     o.RenderedContent,
		IsMeMessage:         o.IsMeMessage,
	}
	return json.Marshal(v)
}

func (o *UpdateMessageEvent) UnmarshalJSON(data []byte) error {
	var v updateMessageEventJSON
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.UserID = v.UserID
	o.RenderingOnly = v.RenderingOnly
	o.MessageID = v.MessageID
	o.MessageIds = v.MessageIds
	o.Flags = v.Flags
	o.EditTimestamp = time.UnixMilli(v.EditTimestamp)
	o.ChannelName = v.ChannelName
	o.ChannelID = v.ChannelID
	o.NewChannelId = v.NewChannelId
	o.PropagateMode = v.PropagateMode
	o.OrigSubject = v.OrigSubject
	o.Subject = v.Subject
	o.TopicLinks = v.TopicLinks
	o.OrigContent = v.OrigContent
	o.OrigRenderedContent = v.OrigRenderedContent
	o.Content = v.Content
	o.RenderedContent = v.RenderedContent
	o.IsMeMessage = v.IsMeMessage
	return nil
}

type updateMessageEventJSON struct {
	UserID              *int64            `json:"user_id"`
	RenderingOnly       bool              `json:"rendering_only"`
	MessageID           int64             `json:"message_id"`
	MessageIds          []int64           `json:"message_ids"`
	Flags               []string          `json:"flags"`
	EditTimestamp       int64             `json:"edit_timestamp"`
	ChannelName         *string           `json:"stream_name,omitempty"`
	ChannelID           *int64            `json:"stream_id,omitempty"`
	NewChannelId        *int64            `json:"new_stream_id,omitempty"`
	PropagateMode       *string           `json:"propagate_mode,omitempty"`
	OrigSubject         *string           `json:"orig_subject,omitempty"`
	Subject             *string           `json:"subject,omitempty"`
	TopicLinks          []zulip.TopicLink `json:"topic_links,omitempty"`
	OrigContent         *string           `json:"orig_content,omitempty"`
	OrigRenderedContent *string           `json:"orig_rendered_content,omitempty"`
	Content             *string           `json:"content,omitempty"`
	RenderedContent     *string           `json:"rendered_content,omitempty"`
	IsMeMessage         *bool             `json:"is_me_message,omitempty"`
}
