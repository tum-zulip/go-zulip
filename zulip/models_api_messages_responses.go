package zulip

// CheckMessagesMatchNarrowResponse struct for CheckMessagesMatchNarrowResponse
type CheckMessagesMatchNarrowResponse struct {
	Response

	// A dictionary with a key for each queried message that matches the narrow, with message Ids as keys and search rendering data as values.
	Messages map[string]NarrowMatch `json:"messages,omitempty"`
}

// NarrowMatch `message_id`: The Id of the message that matches the narrow. No record will be returned for queried messages that do not match the narrow.
type NarrowMatch struct {
	// HTML content of a queried message that matches the narrow. If the narrow is a search narrow, `<span class=\"highlight\">` elements will be included, wrapping the matches for the search keywords.
	MatchContent *string `json:"match_content,omitempty"`
	// HTML-escaped topic of a queried message that matches the narrow. If the narrow is a search narrow, `<span class=\"highlight\">` elements will be included wrapping the matches for the search keywords.
	MatchSubject *string `json:"match_subject,omitempty"`
}

// GetFileTemporaryUrlResponse struct for GetFileTemporaryUrlResponse
type GetFileTemporaryUrlResponse struct {
	Response

	// A temporary URL that can be used to access the uploaded file without Zulip's normal API authentication.
	Url string `json:"url,omitempty"`
}

// GetMessageResponse struct for GetMessageResponse
type GetMessageResponse struct {
	Response

	// The raw Markdown content of the message.  See the help center article on [message formatting](https://zulip.com/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.  **Deprecated** and to be removed once no longer required for legacy clients. Modern clients should prefer passing `\"apply_markdown\": false` to request raw message content.
	// Deprecated
	RawContent string  `json:"raw_content,omitempty"`
	Message    Message `json:"message,omitempty"`
}

// GetMessageHistoryResponse struct for GetMessageHistoryResponse
type GetMessageHistoryResponse struct {
	Response

	// A chronologically sorted, oldest to newest, array of `snapshot` objects, each one with the values of the message after the edit.
	MessageHistory []Snapshot `json:"message_history,omitempty"`
}

// GetMessagesResponse struct for GetMessagesResponse
type GetMessagesResponse struct {
	Response

	// The same `anchor` specified in the request (or the computed one, if `use_first_unread_anchor` is `true`).  Only present if `message_ids` is not provided.
	Anchor int32 `json:"anchor,omitempty"`
	// Whether the server promises that the `messages` list includes the very newest messages matching the narrow (used by clients that paginate their requests to decide whether there may be more messages to fetch).
	FoundNewest bool `json:"found_newest,omitempty"`
	// Whether the server promises that the `messages` list includes the very oldest messages matching the narrow (used by clients that paginate their requests to decide whether there may be more messages to fetch).
	FoundOldest bool `json:"found_oldest,omitempty"`
	// Whether the anchor message is included in the response. If the message with the Id specified in the request does not exist, did not match the narrow, or was excluded via `\"include_anchor\": false`, this will be false.
	FoundAnchor bool `json:"found_anchor,omitempty"`
	// Whether the message history was limited due to plan restrictions. This flag is set to `true` only when the oldest messages(`found_oldest`) matching the narrow is fetched.
	HistoryLimited bool `json:"history_limited,omitempty"`
	// An array of `message` objects.  **Changes**: In Zulip 3.1 (feature level 26), the `sender_short_name` field was removed from message objects.
	Messages []MessageWithOptMatch `json:"messages"`
}

// GetReadReceiptsResponse struct for GetReadReceiptsResponse
type GetReadReceiptsResponse struct {
	Response

	// An array of Ids of users who have marked the target message as read and whose read status is available to the current user.  The Ids of users who have disabled sending read receipts (`\"send_read_receipts\": false`) will never appear in the response, nor will the message's sender. Additionally, the Ids of any users who have been muted by the current user or who have muted the current user will not be included in the response.  The current user's Id will appear if they have marked the target message as read.  **Changes**: Prior to Zulip 6.0 (feature level 143), the Ids of users who have been muted by or have muted the current user were included in the response.
	UserIds []int64 `json:"user_ids,omitempty"`
}

// RenderMessageResponse struct for RenderMessageResponse
type RenderMessageResponse struct {
	Response

	// The rendered HTML.
	Rendered string `json:"rendered,omitempty"`
}

// SendMessageResponse struct for SendMessageResponse
type SendMessageResponse struct {
	Response

	// The unique Id assigned to the sent message.
	Id int64 `json:"id"`
	// If the message's sender had configured their [visibility policy settings](https://zulip.com/help/mute-a-topic) to potentially automatically follow or unmute topics when sending messages, and one of these policies did in fact change the user's visibility policy for the topic where this message was sent, the new value for that user's visibility policy for the recipient topic.  Only present if the sender's visibility was in fact changed.  The value can be either [unmuted or followed](https://zulip.com/api/update-user-topic#parameter-visibility_policy).  Clients will also be notified about the change in policy via a `user_topic` event as usual. This field is intended to be used by clients to explicitly inform the user when a topic's visibility policy was changed automatically due to sending a message.  For example, the Zulip web application uses this field to decide whether to display a warning or notice suggesting to unmute the topic after sending a message to a muted channel. Such a notice would be confusing in the event that the act of sending the message had already resulted in the user automatically unmuting or following the topic in question.  **Changes**: New in Zulip 8.0 (feature level 218).
	AutomaticNewVisibilityPolicy *VisibilityPolicy `json:"automatic_new_visibility_policy,omitempty"`
}

// UpdateMessageResponse struct for UpdateMessageResponse
type UpdateMessageResponse struct {
	Response

	// Details on all files uploaded by the acting user whose only references were removed when editing this message. Clients should ask the acting user if they wish to delete the uploaded files returned in this response, which might otherwise remain visible only in message edit history.  Note that [access to message edit history][edit-history-access] is configurable; this detail may be important in presenting the question clearly to users.  New in Zulip 10.0 (feature level 285).  [edit-history-access]: https://zulip.com/help/restrict-message-edit-history-access
	DetachedUploads []Attachment `json:"detached_uploads,omitempty"`
}

// UpdateMessageFlagsResponse struct for UpdateMessageFlagsResponse
type UpdateMessageFlagsResponse struct {
	Response

	// An array with the Ids of the modified messages.
	Messages []int64 `json:"messages,omitempty"`
	// Only present if the flag is `read` and the operation is `remove`.  Zulip has an invariant that all unread messages must be in channels the user is subscribed to. This field will contain a list of the channels whose messages were skipped to mark as unread because the user is not subscribed to them.  **Changes**: New in Zulip 10.0 (feature level 355).
	IgnoredBecauseNotSubscribedChannels []int64 `json:"ignored_because_not_subscribed_channels,omitempty"`
}

// UpdateMessageFlagsForNarrowResponse struct for UpdateMessageFlagsForNarrowResponse
type UpdateMessageFlagsForNarrowResponse struct {
	Response

	// The number of messages that were within the update range (at most `num_before + 1 + num_after`).
	ProcessedCount int64 `json:"processed_count"`
	// The number of messages where the flag's value was changed (at most `processed_count`).
	UpdatedCount int64 `json:"updated_count"`
	// The Id of the oldest message within the update range, or `null` if the range was empty.
	FirstProcessedId *int64 `json:"first_processed_id"`
	// The Id of the newest message within the update range, or `null` if the range was empty.
	LastProcessedId *int64 `json:"last_processed_id"`
	// Whether the update range reached backward far enough to include very oldest message matching the narrow (used by clients doing a bulk update to decide whether to issue another request anchored at `first_processed_id`).
	FoundOldest bool `json:"found_oldest"`
	// Whether the update range reached forward far enough to include very oldest message matching the narrow (used by clients doing a bulk update to decide whether to issue another request anchored at `last_processed_id`).
	FoundNewest bool `json:"found_newest"`
	// Only present if the flag is `read` and the operation is `remove`.  Zulip has an invariant that all unread messages must be in channels the user is subscribed to. This field will contain a list of the channels whose messages were skipped to mark as unread because the user is not subscribed to them.  **Changes**: New in Zulip 10.0 (feature level 355).
	IgnoredBecauseNotSubscribedChannels []int64 `json:"ignored_because_not_subscribed_channels,omitempty"`
}

// UploadFileResponse struct for UploadFileResponse
type UploadFileResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
	// An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling](https://zulip.com/api/rest-error-handling#ignored-parameters) documentation for details on this and its change history.
	IgnoredParametersUnsupported []string `json:"ignored_parameters_unsupported,omitempty"`
	// The URL of the uploaded file. Alias of `url`.  **Changes**: Deprecated in Zulip 9.0 (feature level 272). The term \"URI\" is deprecated in [web standards](https://url.spec.whatwg.org/#goals).
	// Deprecated
	Uri string `json:"uri,omitempty"`
	// The URL of the uploaded file.  **Changes**: New in Zulip 9.0 (feature level 272). Previously, this property was only available under the legacy `uri` name.
	Url string `json:"url,omitempty"`
	// The filename that Zulip stored the upload as. This usually differs from the basename of the URL when HTML escaping is required to generate a valid URL.  Clients generating a Markdown link to a newly uploaded file should do so by combining the `url` and `filename` fields in the response as follows: `[{filename}]({url})`, with care taken to clean `filename` of `[` and `]` characters that might break Markdown rendering.  **Changes**: New in Zulip 10.0 (feature level 285).
	Filename string `json:"filename,omitempty"`
}
