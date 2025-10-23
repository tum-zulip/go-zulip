package zulip

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type MessagesAPI interface {

	// AddReaction Add an emoji reaction
	//
	// Add an [emoji reaction] to a message.
	//
	// [emoji reaction]: https://zulip.com/help/emoji-reactions
	AddReaction(ctx context.Context, messageId int64) AddReactionRequest

	// AddReactionExecute executes the request
	AddReactionExecute(r AddReactionRequest) (*Response, *http.Response, error)

	// CheckMessagesMatchNarrow Check if messages match a narrow
	//
	// Check whether a set of messages match a [narrow].
	//
	// For many common narrows (e.g. a topic), clients can write an efficient
	// client-side check to determine whether a newly arrived message belongs
	// in the view.
	//
	// This endpoint is designed to allow clients to handle more complex narrows
	// for which the client does not (or in the case of full-text search, cannot)
	// implement this check.
	//
	// The format of the `match_subject` and `match_content` objects is designed
	// to match those returned by the [`GET /messages`]
	// endpoint, so that a client can splice these fields into a `message` object
	// received from [`GET /events`] and end up with an
	// extended message object identical to how a [`GET /messages`]
	// request for the current narrow would have returned the message.
	//
	// [narrow]: https://zulip.com/api/construct-narrow
	// [`GET /messages`]: https://zulip.com/api/get-messages#response
	// [`GET /events`]: https://zulip.com/api/get-events#message
	CheckMessagesMatchNarrow(ctx context.Context) CheckMessagesMatchNarrowRequest

	// CheckMessagesMatchNarrowExecute executes the request
	CheckMessagesMatchNarrowExecute(r CheckMessagesMatchNarrowRequest) (*CheckMessagesMatchNarrowResponse, *http.Response, error)

	// DeleteMessage Delete a message
	//
	// Permanently delete a message.
	//
	// This API corresponds to the
	// [delete a message completely] feature documented in
	// the Zulip Help Center.
	//
	// [delete a message completely]: https://zulip.com/help/delete-a-message#delete-a-message-completely
	//
	DeleteMessage(ctx context.Context, messageId int64) DeleteMessageRequest

	// DeleteMessageExecute executes the request
	DeleteMessageExecute(r DeleteMessageRequest) (*Response, *http.Response, error)

	// GetFileTemporaryUrl Get public temporary URL
	//
	// Get a temporary URL for access to the file that doesn't require authentication.
	//
	// *Changes**: New in Zulip 3.0 (feature level 1).
	//
	GetFileTemporaryUrl(ctx context.Context, realmId int64, filename string) GetFileTemporaryUrlRequest

	// GetFileTemporaryUrlExecute executes the request
	GetFileTemporaryUrlExecute(r GetFileTemporaryUrlRequest) (*GetFileTemporaryUrlResponse, *http.Response, error)

	// GetMessage Fetch a single message
	//
	// Given a message Id, return the message object.
	//
	// Additionally, a `raw_content` field is included. This field is
	// useful for clients that primarily work with HTML-rendered
	// messages but might need to occasionally fetch the message's
	// raw [Zulip-flavored Markdown] (e.g. for [view source] or
	// prefilling a message edit textarea).
	//
	// *Changes**: Before Zulip 5.0 (feature level 120), this
	// endpoint only returned the `raw_content` field.
	//
	// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
	// [view source]: https://zulip.com/help/view-the-markdown-source-of-a-message
	GetMessage(ctx context.Context, messageId int64) GetMessageRequest

	// GetMessageExecute executes the request
	GetMessageExecute(r GetMessageRequest) (*GetMessageResponse, *http.Response, error)

	// GetMessageHistory Get a message's edit history
	//
	// Fetch the message edit history of a previously edited message.
	//
	// Note that edit history may be disabled in some organizations; see the
	// [Zulip Help Center documentation on editing messages].
	//
	// [Zulip Help Center documentation on editing messages]: https://zulip.com/help/view-a-messages-edit-history
	//
	GetMessageHistory(ctx context.Context, messageId int64) GetMessageHistoryRequest

	// GetMessageHistoryExecute executes the request
	GetMessageHistoryExecute(r GetMessageHistoryRequest) (*GetMessageHistoryResponse, *http.Response, error)

	// GetMessages Get messages
	//
	// This endpoint is the primary way to fetch a messages. It is used by all official
	// Zulip clients (e.g. the web, desktop, mobile, and terminal clients) as well as
	// many bots, API clients, backup scripts, etc.
	//
	// Most queries will specify a [narrow filter],
	// to fetch the messages matching any supported [search query]. If not specified, it will return messages
	// corresponding to the user's [combined feed]. There are two
	// ways to specify which messages matching the narrow filter to fetch:
	//
	//   - A range of messages, described by an `anchor` message Id (or a string-format
	// specification of how the server should computer an anchor to use) and a maximum
	// number of messages in each direction from that anchor.
	//
	//   - A rarely used variant (`message_ids`) where the client specifies the message Ids
	// to fetch.
	//
	// The server returns the matching messages, sorted by message Id, as well as some
	// metadata that makes it easy for a client to determine whether there are more
	// messages matching the query that were not returned due to the `num_before` and
	// `num_after` limits.
	//
	// Note that a user's message history does not contain messages sent to
	// channels before they [subscribe], and newly created
	// bot users are not usually subscribed to any channels.
	//
	// We recommend requesting at most 1000 messages in a batch, to avoid generating very
	// large HTTP responses. A maximum of 5000 messages can be obtained per request;
	// attempting to exceed this will result in an error.
	//
	// *Changes**: The `message_ids` option is new in Zulip 10.0 (feature level 300).
	//
	// [narrow filter]: https://zulip.com/api/get-messages#parameter-narrow
	// [combined feed]: https://zulip.com/help/combined-feed
	// [subscribe]: https://zulip.com/api/subscribe
	// [search query]: https://zulip.com/help/search-for-messages
	GetMessages(ctx context.Context) GetMessagesRequest

	// GetMessagesExecute executes the request
	GetMessagesExecute(r GetMessagesRequest) (*GetMessagesResponse, *http.Response, error)

	// GetReadReceipts Get a message's read receipts
	//
	// Returns a list containing the Ids for all users who have
	// marked the message as read (and whose privacy settings allow
	// sharing that information).
	//
	// The list of users Ids will include any bots who have marked
	// the message as read via the API (providing a way for bots to
	// indicate whether they have processed a message successfully in
	// a way that can be easily inspected in a Zulip client). Bots
	// for which this behavior is not desired may disable the
	// `send_read_receipts` setting via the API.
	//
	// It will never contain the message's sender.
	//
	// *Changes**: New in Zulip 6.0 (feature level 137).
	//
	GetReadReceipts(ctx context.Context, messageId int64) GetReadReceiptsRequest

	// GetReadReceiptsExecute executes the request
	GetReadReceiptsExecute(r GetReadReceiptsRequest) (*GetReadReceiptsResponse, *http.Response, error)

	// MarkAllAsRead Mark all messages as read
	//
	// Marks all of the current user's unread messages as read.
	//
	// Because this endpoint marks messages as read in batches, it is possible
	// for the request to time out after only marking some messages as read.
	// When this happens, the `complete` boolean field in the success response
	// will be `false`. Clients should repeat the request when handling such a
	// response. If all messages were marked as read, then the success response
	// will return `"complete": true`.
	//
	// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
	// as this endpoint will be removed in a future release.
	//
	// Before Zulip 8.0 (feature level 211), if the server's
	// processing was interrupted by a timeout, but some messages were marked
	// as read, then it would return `"result": "partially_completed"`, along
	// with a `code` field for an error string, in the success response to
	// indicate that there was a timeout and that the client should repeat the
	// request.
	//
	// Before Zulip 6.0 (feature level 153), this request did a single atomic
	// operation, which could time out with 10,000s of unread messages to mark
	// as read. As of this feature level, messages are marked as read in
	// batches, starting with the newest messages, so that progress is made
	// even if the request times out. And, instead of returning an error when
	// the request times out and some messages have been marked as read, a
	// success response with `"result": "partially_completed"` is returned.
	//
	//
	// Deprecated
	//
	// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
	MarkAllAsRead(ctx context.Context) MarkAllAsReadRequest

	// MarkAllAsReadExecute executes the request
	// Deprecated
	MarkAllAsReadExecute(r MarkAllAsReadRequest) (*MarkAllAsReadResponse, *http.Response, error)

	// MarkChannelAsRead Mark messages in a channel as read
	//
	// Mark all the unread messages in a channel as read.
	//
	// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
	// as this endpoint will be removed in a future release.
	//
	//
	// Deprecated
	//
	// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
	MarkChannelAsRead(ctx context.Context) MarkChannelAsReadRequest

	// MarkChannelAsReadExecute executes the request
	// Deprecated
	MarkChannelAsReadExecute(r MarkChannelAsReadRequest) (*Response, *http.Response, error)

	// MarkTopicAsRead Mark messages in a topic as read
	//
	// Mark all the unread messages in a topic as read.
	//
	// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
	// as this endpoint will be removed in a future release.
	//
	//
	// Deprecated
	//
	// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
	MarkTopicAsRead(ctx context.Context) MarkTopicAsReadRequest

	// MarkTopicAsReadExecute executes the request
	// Deprecated
	MarkTopicAsReadExecute(r MarkTopicAsReadRequest) (*Response, *http.Response, error)

	// RemoveReaction Remove an emoji reaction
	//
	// Remove an [emoji reaction] from a message.
	//
	// [emoji reaction]: https://zulip.com/help/emoji-reactions
	RemoveReaction(ctx context.Context, messageId int64) RemoveReactionRequest

	// RemoveReactionExecute executes the request
	RemoveReactionExecute(r RemoveReactionRequest) (*Response, *http.Response, error)

	// RenderMessage Render a message
	//
	// Render a message to HTML.
	//
	RenderMessage(ctx context.Context) RenderMessageRequest

	// RenderMessageExecute executes the request
	RenderMessageExecute(r RenderMessageRequest) (*RenderMessageResponse, *http.Response, error)

	// ReportMessage Report a message
	//
	// Sends a notification to the organization's moderation request channel,
	// if it is configured, that reports the targeted message for review and
	// moderation.
	//
	// Clients should check the `moderation_request_channel` realm setting to
	// decide whether to show the option to report messages in the UI.
	//
	// If the `report_type` parameter value is `"other"`, the `description`
	// parameter is required. Clients should also enforce and communicate this
	// behavior in the UI.
	//
	// *Changes**: New in Zulip 11.0 (feature level 382). This API builds on
	// the `moderation_request_channel` realm setting, which was added in
	// feature level 331.
	//
	ReportMessage(ctx context.Context, messageId int64) ReportMessageRequest

	// ReportMessageExecute executes the request
	ReportMessageExecute(r ReportMessageRequest) (*Response, *http.Response, error)

	// SendMessage Send a message
	//
	// Send a [channel message] or a
	// [direct message].
	//
	// [channel message]: https://zulip.com/help/introduction-to-topics
	// [direct message]: https://zulip.com/help/direct-messages
	SendMessage(ctx context.Context) SendMessageRequest

	// SendMessageExecute executes the request
	SendMessageExecute(r SendMessageRequest) (*SendMessageResponse, *http.Response, error)

	// UpdateMessage Edit a message
	//
	// Update the content, topic, or channel of the message with the specified
	// Id.
	//
	// You can [resolve topics] by editing the topic to
	// `✔ {original_topic}` with the `propagate_mode` parameter set to
	// `"change_all"`.
	//
	// See [configuring message editing] for detailed
	// documentation on when users are allowed to edit message content, and
	// [restricting moving messages] for detailed
	// documentation on when users are allowed to change a message's topic
	// and/or channel.
	//
	// The relevant realm settings in the API that are related to the above
	// linked documentation on when users are allowed to update messages are:
	//
	//   - `allow_message_editing`
	//   - `can_resolve_topics_group`
	//   - `can_move_messages_between_channels_group`
	//   - `can_move_messages_between_topics_group`
	//   - `message_content_edit_limit_seconds`
	//   - `move_messages_within_stream_limit_seconds`
	//   - `move_messages_between_streams_limit_seconds`
	//
	// More details about these realm settings can be found in the
	// [`POST /register`] response or in the documentation
	// of the [`realm op: update_dict`] event in [`GET /events`].
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 367), the permission for
	// resolving a topic was managed by `can_move_messages_between_topics_group`.
	// As of this feature level, users belonging to the `can_resolve_topics_group`
	// will have the permission to [resolve topics] in the organization.
	//
	// In Zulip 10.0 (feature level 316), `edit_topic_policy`
	// was removed and replaced by `can_move_messages_between_topics_group`
	// realm setting.
	//
	// *Changes**: In Zulip 10.0 (feature level 310), `move_messages_between_streams_policy`
	// was removed and replaced by `can_move_messages_between_channels_group`
	// realm setting.
	//
	// Prior to Zulip 7.0 (feature level 172), anyone could add a
	// topic to channel messages without a topic, regardless of the organization's
	// [topic editing permissions]. As of this
	// feature level, messages without topics have the same restrictions for
	// topic edits as messages with topics.
	//
	// Before Zulip 7.0 (feature level 172), by using the `change_all` value for
	// the `propagate_mode` parameter, users could move messages after the
	// organization's configured time limits for changing a message's topic or
	// channel had passed. As of this feature level, the server will [return an error] with `"code":
	// "MOVE_MESSAGES_TIME_LIMIT_EXCEEDED"` if users, other than organization
	// administrators or moderators, try to move messages after these time
	// limits have passed.
	//
	// Before Zulip 7.0 (feature level 162), users who were not administrators or
	// moderators could only edit topics if the target message was sent within the
	// last 3 days. As of this feature level, that time limit is now controlled by
	// the realm setting `move_messages_within_stream_limit_seconds`. Also at this
	// feature level, a similar time limit for moving messages between channels was
	// added, controlled by the realm setting
	// `move_messages_between_streams_limit_seconds`. Previously, all users who
	// had permission to move messages between channels did not have any time limit
	// restrictions when doing so.
	//
	// Before Zulip 7.0 (feature level 159), editing channels and topics of messages
	// was forbidden if the realm setting for `allow_message_editing` was `false`,
	// regardless of an organization's configuration for the realm settings
	// `edit_topic_policy` or `move_messages_between_streams_policy`.
	//
	// Before Zulip 7.0 (feature level 159), message senders were allowed to edit
	// the topic of their messages indefinitely.
	//
	// In Zulip 5.0 (feature level 75), the `edit_topic_policy` realm setting
	// was added, replacing the `allow_community_topic_editing` boolean.
	//
	// In Zulip 4.0 (feature level 56), the `move_messages_between_streams_policy`
	// realm setting was added.
	//
	// [configuring message editing]: https://zulip.com/help/restrict-message-editing-and-deletion
	// [restricting moving messages]: https://zulip.com/help/restrict-moving-messages
	// [`realm op: update_dict`] https://zulip.com/api/get-events#realm-update_dict
	// [resolve topics]: https://zulip.com/help/resolve-a-topic
	// [`POST /register`]: https://zulip.com/api/register-queue
	// [`GET /events`]: https://zulip.com/api/get-events
	// [topic editing permissions]: https://zulip.com/help/restrict-moving-messages
	// [return an error]: https://zulip.com/api/update-message#response
	UpdateMessage(ctx context.Context, messageId int64) UpdateMessageRequest

	// UpdateMessageExecute executes the request
	UpdateMessageExecute(r UpdateMessageRequest) (*UpdateMessageResponse, *http.Response, error)

	// UpdateMessageFlags Update personal message flags
	//
	// Add or remove personal message flags like `read` and `starred`
	// on a collection of message Ids.
	//
	// See also the endpoint for [updating flags on a range of messages within a narrow].
	//
	// [updating flags on a range of messages within a narrow]: https://zulip.com/api/update-message-flags-for-narrow
	UpdateMessageFlags(ctx context.Context) UpdateMessageFlagsRequest

	// UpdateMessageFlagsExecute executes the request
	UpdateMessageFlagsExecute(r UpdateMessageFlagsRequest) (*UpdateMessageFlagsResponse, *http.Response, error)

	// UpdateMessageFlagsForNarrow Update personal message flags for narrow
	//
	// Add or remove personal message flags like `read` and `starred`
	// on a range of messages within a narrow.
	//
	// See also [the endpoint for updating flags on specific message Ids].
	//
	// *Changes**: New in Zulip 6.0 (feature level 155).
	//
	// [the endpoint for updating flags on specific message Ids]: https://zulip.com/api/update-message-flags
	UpdateMessageFlagsForNarrow(ctx context.Context) UpdateMessageFlagsForNarrowRequest

	// UpdateMessageFlagsForNarrowExecute executes the request
	UpdateMessageFlagsForNarrowExecute(r UpdateMessageFlagsForNarrowRequest) (*UpdateMessageFlagsForNarrowResponse, *http.Response, error)

	// UploadFile Upload a file
	//
	// [Upload] a single file and get the corresponding URL.
	//
	// Initially, only you will be able to access the link. To share the
	// uploaded file, you'll need to [send a message]
	// containing the resulting link. Users who can already access the link
	// can reshare it with other users by sending additional Zulip messages
	// containing the link.
	//
	// The maximum allowed file size is available in the `max_file_upload_size_mib`
	// field in the [`POST /register`] response. Note that
	// large files (25MB+) may fail to upload using this API endpoint due to
	// network-layer timeouts, depending on the quality of your connection to the
	// Zulip server.
	//
	// For uploading larger files, `/api/v1/tus` is an endpoint implementing the
	// [`tus` resumable upload protocol],
	// which supports uploading arbitrarily large files limited only by the server's
	// `max_file_upload_size_mib` (Configured via `MAX_FILE_UPLOAD_SIZE` in
	// `/etc/zulip/settings.py`). Clients which send authenticated credentials
	// (either via browser-based cookies, or API key via `Authorization` header) may
	// use this endpoint to upload files.
	//
	// *Changes**: The `api/v1/tus` endpoint supporting resumable uploads was
	// introduced in Zulip 10.0 (feature level 296). Previously,
	// `max_file_upload_size_mib` was typically 25MB.
	//
	// [uploaded-files]: https://zulip.com/help/manage-your-uploaded-files
	// [send a message]: https://zulip.com/api/send-message
	//
	// [Upload]: https://zulip.com/help/share-and-upload-files
	// [`POST /register`]: https://zulip.com/api/register-queue
	// [`tus` resumable upload protocol]: https://tus.io/protocols/resumable-upload
	UploadFile(ctx context.Context) UploadFileRequest

	// UploadFileExecute executes the request
	UploadFileExecute(r UploadFileRequest) (*UploadFileResponse, *http.Response, error)
}

type AddReactionRequest struct {
	ctx          context.Context
	apiService   MessagesAPI
	messageId    int64
	emojiName    *string
	emojiCode    *string
	reactionType *string
}

// The target emoji's human-readable name.  To find an emoji's name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji's text name.
func (r AddReactionRequest) EmojiName(emojiName string) AddReactionRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  For most API clients, you won't need this, but it's important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the `emoji_code` in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a "different" emoji.
func (r AddReactionRequest) EmojiCode(emojiCode string) AddReactionRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity.
//
// Supported values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
//
// **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji]; previously, this endpoint assumed `unicode_emoji` if this parameter was not specified.
//
// [custom emoji]: https://zulip.com/help/custom-emoji
func (r AddReactionRequest) ReactionType(reactionType string) AddReactionRequest {
	r.reactionType = &reactionType
	return r
}

func (r AddReactionRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.AddReactionExecute(r)
}

// AddReaction Add an emoji reaction
//
// Add an [emoji reaction] to a message.
//
// [emoji reaction]: https://zulip.com/help/emoji-reactions
func (c *simpleClient) AddReaction(ctx context.Context, messageId int64) AddReactionRequest {
	return AddReactionRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) AddReactionExecute(r AddReactionRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/messages/{message_id}/reactions"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	if r.emojiName == nil {
		return nil, nil, reportError("emojiName is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "emoji_name", r.emojiName)
	addOptionalParam(form, "emoji_code", r.emojiCode)
	addOptionalParam(form, "reaction_type", r.reactionType)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type CheckMessagesMatchNarrowRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	msgIds     *[]int64
	narrow     *Narrow
}

// List of Ids for the messages to check.
func (r CheckMessagesMatchNarrowRequest) MsgIds(msgIds []int64) CheckMessagesMatchNarrowRequest {
	r.msgIds = &msgIds
	return r
}

// A structure defining the narrow to check against. See how to [construct a narrow].
//
// **Changes**: See [changes section] of search/narrow filter documentation.
//
// [construct a narrow]: https://zulip.com/api/construct-narrow
// [changes section]: https://zulip.com/api/construct-narrow#changes
func (r CheckMessagesMatchNarrowRequest) Narrow(narrow *Narrow) CheckMessagesMatchNarrowRequest {
	r.narrow = narrow
	return r
}

func (r CheckMessagesMatchNarrowRequest) Execute() (*CheckMessagesMatchNarrowResponse, *http.Response, error) {
	return r.apiService.CheckMessagesMatchNarrowExecute(r)
}

// CheckMessagesMatchNarrow Check if messages match a narrow
//
// Check whether a set of messages match a [narrow].
//
// For many common narrows (e.g. a topic), clients can write an efficient
// client-side check to determine whether a newly arrived message belongs
// in the view.
//
// This endpoint is designed to allow clients to handle more complex narrows
// for which the client does not (or in the case of full-text search, cannot)
// implement this check.
//
// The format of the `match_subject` and `match_content` objects is designed
// to match those returned by the [`GET /messages`]
// endpoint, so that a client can splice these fields into a `message` object
// received from [`GET /events`] and end up with an
// extended message object identical to how a [`GET /messages`]
// request for the current narrow would have returned the message.
//
// [narrow]: https://zulip.com/api/construct-narrow
// [`GET /messages`]: https://zulip.com/api/get-messages#response
// [`GET /events`]: https://zulip.com/api/get-events#message
func (c *simpleClient) CheckMessagesMatchNarrow(ctx context.Context) CheckMessagesMatchNarrowRequest {
	return CheckMessagesMatchNarrowRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CheckMessagesMatchNarrowExecute(r CheckMessagesMatchNarrowRequest) (*CheckMessagesMatchNarrowResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CheckMessagesMatchNarrowResponse{}
		endpoint = "/messages/matches_narrow"
	)
	if r.msgIds == nil {
		return nil, nil, reportError("msgIds is required and must be specified")
	}
	if r.narrow == nil {
		return nil, nil, reportError("narrow is required and must be specified")
	}

	addCSVParam(query, "msg_ids", r.msgIds)
	addCSVParam(query, "narrow", r.narrow)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type DeleteMessageRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	messageId  int64
}

func (r DeleteMessageRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.DeleteMessageExecute(r)
}

// DeleteMessage Delete a message
//
// Permanently delete a message.
//
// This API corresponds to the
// [delete a message completely] feature documented in
// the Zulip Help Center.
//
// [delete a message completely]: https://zulip.com/help/delete-a-message#delete-a-message-completely
func (c *simpleClient) DeleteMessage(ctx context.Context, messageId int64) DeleteMessageRequest {
	return DeleteMessageRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteMessageExecute(r DeleteMessageRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/messages/{message_id}"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetFileTemporaryUrlRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	realmId    int64
	filename   string
}

func (r GetFileTemporaryUrlRequest) Execute() (*GetFileTemporaryUrlResponse, *http.Response, error) {
	return r.apiService.GetFileTemporaryUrlExecute(r)
}

// GetFileTemporaryUrl Get public temporary URL
//
// Get a temporary URL for access to the file that doesn't require authentication.
//
// *Changes**: New in Zulip 3.0 (feature level 1).
func (c *simpleClient) GetFileTemporaryUrl(ctx context.Context, realmId int64, filename string) GetFileTemporaryUrlRequest {
	return GetFileTemporaryUrlRequest{
		apiService: c,
		ctx:        ctx,
		realmId:    realmId,
		filename:   filename,
	}
}

// Execute executes the request
func (c *simpleClient) GetFileTemporaryUrlExecute(r GetFileTemporaryUrlRequest) (*GetFileTemporaryUrlResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetFileTemporaryUrlResponse{}
		endpoint = "/user_uploads/{realm_id_str}/{filename}"
	)

	endpoint = strings.Replace(endpoint, "{realm_id_str}", idToString(r.realmId), -1)
	endpoint = strings.Replace(endpoint, "{filename}", url.PathEscape(r.filename), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetMessageRequest struct {
	ctx                 context.Context
	apiService          MessagesAPI
	messageId           int64
	applyMarkdown       *bool
	allowEmptyTopicName *bool
}

// If `true`, message content is returned in the rendered HTML format. If `false`, message content is returned in the raw [Zulip-flavored Markdown format] text that user entered.
//
// **Changes**: New in Zulip 5.0 (feature level 120).
//
// [Zulip-flavored Markdown format]: https://zulip.com/help/format-your-message-using-markdown
func (r GetMessageRequest) ApplyMarkdown(applyMarkdown bool) GetMessageRequest {
	r.applyMarkdown = &applyMarkdown
	return r
}

// Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If `false`, the server will use the value of `realm_empty_topic_display_name` found in the [`POST /register`] response instead of empty string to represent the empty string topic in its response.
//
// **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r GetMessageRequest) AllowEmptyTopicName(allowEmptyTopicName bool) GetMessageRequest {
	r.allowEmptyTopicName = &allowEmptyTopicName
	return r
}

func (r GetMessageRequest) Execute() (*GetMessageResponse, *http.Response, error) {
	return r.apiService.GetMessageExecute(r)
}

// GetMessage Fetch a single message
//
// Given a message Id, return the message object.
//
// Additionally, a `raw_content` field is included. This field is
// useful for clients that primarily work with HTML-rendered
// messages but might need to occasionally fetch the message's
// raw [Zulip-flavored Markdown] (e.g. for [view source] or
// prefilling a message edit textarea).
//
// *Changes**: Before Zulip 5.0 (feature level 120), this
// endpoint only returned the `raw_content` field.
//
// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
// [view source]: https://zulip.com/help/view-the-markdown-source-of-a-message
func (c *simpleClient) GetMessage(ctx context.Context, messageId int64) GetMessageRequest {
	return GetMessageRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) GetMessageExecute(r GetMessageRequest) (*GetMessageResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetMessageResponse{}
		endpoint = "/messages/{message_id}"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	addOptionalParam(query, "apply_markdown", r.applyMarkdown)
	addOptionalParam(query, "allow_empty_topic_name", r.allowEmptyTopicName)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetMessageHistoryRequest struct {
	ctx                 context.Context
	apiService          MessagesAPI
	messageId           int64
	allowEmptyTopicName *bool
}

// Whether the topic names i.e. `topic` and `prev_topic` fields in the `message_history` objects returned can be empty string.  If `false`, the value of `realm_empty_topic_display_name` found in the [`POST /register`] response is returned replacing the empty string as the topic name.
//
// **Changes**: New in Zulip 10.0 (feature level 334).
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r GetMessageHistoryRequest) AllowEmptyTopicName(allowEmptyTopicName bool) GetMessageHistoryRequest {
	r.allowEmptyTopicName = &allowEmptyTopicName
	return r
}

func (r GetMessageHistoryRequest) Execute() (*GetMessageHistoryResponse, *http.Response, error) {
	return r.apiService.GetMessageHistoryExecute(r)
}

// GetMessageHistory Get a message's edit history
//
// Fetch the message edit history of a previously edited message.
//
// Note that edit history may be disabled in some organizations; see the
// [Zulip Help Center documentation on editing messages].
//
// [Zulip Help Center documentation on editing messages]: https://zulip.com/help/view-a-messages-edit-history
func (c *simpleClient) GetMessageHistory(ctx context.Context, messageId int64) GetMessageHistoryRequest {
	return GetMessageHistoryRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) GetMessageHistoryExecute(r GetMessageHistoryRequest) (*GetMessageHistoryResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetMessageHistoryResponse{}
		endpoint = "/messages/{message_id}/history"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	addOptionalParam(query, "allow_empty_topic_name", r.allowEmptyTopicName)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetMessagesRequest struct {
	ctx                  context.Context
	apiService           MessagesAPI
	anchor               *string
	includeAnchor        *bool
	numBefore            *int32
	numAfter             *int32
	narrow               *Narrow
	clientGravatar       *bool
	applyMarkdown        *bool
	useFirstUnreadAnchor *bool
	messageIds           *[]int64
	allowEmptyTopicName  *bool
}

// Integer message Id to anchor fetching of new messages. Supports special string values for when the client wants the server to compute the anchor to use:  - `newest`: The most recent message. - `oldest`: The oldest message. - `first_unread`: The oldest unread message matching the   query, if any; otherwise, the most recent message.
//
// **Changes**: String values are new in Zulip 3.0 (feature level 1). The `first_unread` functionality was supported in Zulip 2.1.x and older by not sending `anchor` and using `use_first_unread_anchor`.  In Zulip 2.1.x and older, `oldest` can be emulated with `"anchor": 0`, and `newest` with `"anchor": 10000000000000000` (that specific large value works around a bug in Zulip 2.1.x and older in the `found_newest` return value).
func (r GetMessagesRequest) Anchor(anchor string) GetMessagesRequest {
	r.anchor = &anchor
	return r
}

// Whether a message with the specified Id matching the narrow should be included.
//
// **Changes**: New in Zulip 6.0 (feature level 155).
func (r GetMessagesRequest) IncludeAnchor(includeAnchor bool) GetMessagesRequest {
	r.includeAnchor = &includeAnchor
	return r
}

// The number of messages with Ids less than the anchor to retrieve. Required if `message_ids` is not provided.
func (r GetMessagesRequest) NumBefore(numBefore int32) GetMessagesRequest {
	r.numBefore = &numBefore
	return r
}

// The number of messages with Ids greater than the anchor to retrieve. Required if `message_ids` is not provided.
func (r GetMessagesRequest) NumAfter(numAfter int32) GetMessagesRequest {
	r.numAfter = &numAfter
	return r
}

// The narrow where you want to fetch the messages from. See how to [construct a narrow].  Note that many narrows, including all that lack a `channel`, `channels`, `stream`, or `streams` operator, search the user's personal message history. See [searching shared history] for details.  For example, if you would like to fetch messages from all public channels instead of only the user's message history, then a specific narrow for messages sent to all public channels can be used: `{"operator": "channels", "operand": "public"}`.  Newly created bot users are not usually subscribed to any channels, so bots using this API should either be subscribed to appropriate channels or use a shared history search narrow with this endpoint.
//
// **Changes**: See [changes section] of search/narrow filter documentation.
//
// [construct a narrow]: https://zulip.com/api/construct-narrow
// [searching shared history]: https://zulip.com/help/search-for-messages#search-shared-history
// [changes section]: https://zulip.com/api/construct-narrow#changes
func (r GetMessagesRequest) Narrow(narrow *Narrow) GetMessagesRequest {
	r.narrow = narrow
	return r
}

// Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.
//
// **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).
func (r GetMessagesRequest) ClientGravatar(clientGravatar bool) GetMessagesRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// If `true`, message content is returned in the rendered HTML format. If `false`, message content is returned in the raw Markdown-format text that user entered.  See [Markdown message formatting] for details on Zulip's HTML format.
//
// [Markdown message formatting]: https://zulip.com/api/message-formatting
func (r GetMessagesRequest) ApplyMarkdown(applyMarkdown bool) GetMessagesRequest {
	r.applyMarkdown = &applyMarkdown
	return r
}

// Legacy way to specify `"anchor": "first_unread"` in Zulip 2.1.x and older.  Whether to use the (computed by the server) first unread message matching the narrow as the `anchor`. Mutually exclusive with `anchor`.
//
// **Changes**: Deprecated in Zulip 3.0 (feature level 1) and replaced by `"anchor": "first_unread"`.
//
// Deprecated
func (r GetMessagesRequest) UseFirstUnreadAnchor(useFirstUnreadAnchor bool) GetMessagesRequest {
	r.useFirstUnreadAnchor = &useFirstUnreadAnchor
	return r
}

// A list of message Ids to fetch. The server will return messages corresponding to the subset of the requested message Ids that exist and the current user has access to, potentially filtered by the narrow (if that parameter is provided).  It is an error to pass this parameter as well as any of the parameters involved in specifying a range of messages: `anchor`, `include_anchor`, `use_first_unread_anchor`, `num_before`, and `num_after`.
//
// **Changes**: New in Zulip 10.0 (feature level 300). Previously, there was no way to request a specific set of messages Ids.
func (r GetMessagesRequest) MessageIds(messageIds []int64) GetMessagesRequest {
	r.messageIds = &messageIds
	return r
}

// Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If `false`, the server will use the value of `realm_empty_topic_display_name` found in the [`POST /register`] response instead of empty string to represent the empty string topic in its response.
//
// **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r GetMessagesRequest) AllowEmptyTopicName(allowEmptyTopicName bool) GetMessagesRequest {
	r.allowEmptyTopicName = &allowEmptyTopicName
	return r
}

func (r GetMessagesRequest) Execute() (*GetMessagesResponse, *http.Response, error) {
	return r.apiService.GetMessagesExecute(r)
}

// GetMessages Get messages
//
// This endpoint is the primary way to fetch a messages. It is used by all official
// Zulip clients (e.g. the web, desktop, mobile, and terminal clients) as well as
// many bots, API clients, backup scripts, etc.
//
// Most queries will specify a [narrow filter],
// to fetch the messages matching any supported [search query]. If not specified, it will return messages
// corresponding to the user's [combined feed]. There are two
// ways to specify which messages matching the narrow filter to fetch:
//
//   - A range of messages, described by an `anchor` message Id (or a string-format specification of how the server should computer an anchor to use) and a maximum number of messages in each direction from that anchor.
//   - A rarely used variant (`message_ids`) where the client specifies the message Ids to fetch.
//
// The server returns the matching messages, sorted by message Id, as well as some
// metadata that makes it easy for a client to determine whether there are more
// messages matching the query that were not returned due to the `num_before` and
// `num_after` limits.
//
// Note that a user's message history does not contain messages sent to
// channels before they [subscribe], and newly created
// bot users are not usually subscribed to any channels.
//
// We recommend requesting at most 1000 messages in a batch, to avoid generating very
// large HTTP responses. A maximum of 5000 messages can be obtained per request;
// attempting to exceed this will result in an error.
//
// *Changes**: The `message_ids` option is new in Zulip 10.0 (feature level 300).
//
// [narrow filter]: https://zulip.com/api/get-messages#parameter-narrow
// [combined feed]: https://zulip.com/help/combined-feed
// [subscribe]: https://zulip.com/api/subscribe
// [search query]: https://zulip.com/help/search-for-messages
func (c *simpleClient) GetMessages(ctx context.Context) GetMessagesRequest {
	return GetMessagesRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetMessagesExecute(r GetMessagesRequest) (*GetMessagesResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetMessagesResponse{}
		endpoint = "/messages"
	)
	addOptionalParam(query, "anchor", r.anchor)
	addOptionalParam(query, "include_anchor", r.includeAnchor)
	addOptionalParam(query, "num_before", r.numBefore)
	addOptionalParam(query, "num_after", r.numAfter)
	addOptionalCSVParam(query, "narrow", r.narrow)
	addOptionalParam(query, "client_gravatar", r.clientGravatar)
	addOptionalParam(query, "apply_markdown", r.applyMarkdown)
	addOptionalParam(query, "use_first_unread_anchor", r.useFirstUnreadAnchor)
	addOptionalCSVParam(query, "message_ids", r.messageIds)
	addOptionalParam(query, "allow_empty_topic_name", r.allowEmptyTopicName)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type GetReadReceiptsRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	messageId  int64
}

func (r GetReadReceiptsRequest) Execute() (*GetReadReceiptsResponse, *http.Response, error) {
	return r.apiService.GetReadReceiptsExecute(r)
}

// GetReadReceipts Get a message's read receipts
//
// Returns a list containing the Ids for all users who have
// marked the message as read (and whose privacy settings allow
// sharing that information).
//
// The list of users Ids will include any bots who have marked
// the message as read via the API (providing a way for bots to
// indicate whether they have processed a message successfully in
// a way that can be easily inspected in a Zulip client). Bots
// for which this behavior is not desired may disable the
// `send_read_receipts` setting via the API.
//
// It will never contain the message's sender.
//
// *Changes**: New in Zulip 6.0 (feature level 137).
func (c *simpleClient) GetReadReceipts(ctx context.Context, messageId int64) GetReadReceiptsRequest {
	return GetReadReceiptsRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) GetReadReceiptsExecute(r GetReadReceiptsRequest) (*GetReadReceiptsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetReadReceiptsResponse{}
		endpoint = "/messages/{message_id}/read_receipts"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type MarkAllAsReadRequest struct {
	ctx        context.Context
	apiService MessagesAPI
}

func (r MarkAllAsReadRequest) Execute() (*MarkAllAsReadResponse, *http.Response, error) {
	return r.apiService.MarkAllAsReadExecute(r)
}

// MarkAllAsRead Mark all messages as read
//
// Marks all of the current user's unread messages as read.
//
// Because this endpoint marks messages as read in batches, it is possible
// for the request to time out after only marking some messages as read.
// When this happens, the `complete` boolean field in the success response
// will be `false`. Clients should repeat the request when handling such a
// response. If all messages were marked as read, then the success response
// will return `"complete": true`.
//
// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
// as this endpoint will be removed in a future release.
//
// Before Zulip 8.0 (feature level 211), if the server's
// processing was interrupted by a timeout, but some messages were marked
// as read, then it would return `"result": "partially_completed"`, along
// with a `code` field for an error string, in the success response to
// indicate that there was a timeout and that the client should repeat the
// request.
//
// Before Zulip 6.0 (feature level 153), this request did a single atomic
// operation, which could time out with 10,000s of unread messages to mark
// as read. As of this feature level, messages are marked as read in
// batches, starting with the newest messages, so that progress is made
// even if the request times out. And, instead of returning an error when
// the request times out and some messages have been marked as read, a
// success response with `"result": "partially_completed"` is returned.
//
// # Deprecated
//
// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
func (c *simpleClient) MarkAllAsRead(ctx context.Context) MarkAllAsReadRequest {
	return MarkAllAsReadRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
// Deprecated
func (c *simpleClient) MarkAllAsReadExecute(r MarkAllAsReadRequest) (*MarkAllAsReadResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &MarkAllAsReadResponse{}
		endpoint = "/mark_all_as_read"
	)

	headers["Accept"] = "application/json"
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type MarkChannelAsReadRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	channelId  *int64
}

// The Id of the channel to access.
func (r MarkChannelAsReadRequest) ChannelId(channelId int64) MarkChannelAsReadRequest {
	r.channelId = &channelId
	return r
}

func (r MarkChannelAsReadRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.MarkChannelAsReadExecute(r)
}

// MarkChannelAsRead Mark messages in a channel as read
//
// Mark all the unread messages in a channel as read.
//
// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
// as this endpoint will be removed in a future release.
//
// # Deprecated
//
// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
func (c *simpleClient) MarkChannelAsRead(ctx context.Context) MarkChannelAsReadRequest {
	return MarkChannelAsReadRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
// Deprecated
func (c *simpleClient) MarkChannelAsReadExecute(r MarkChannelAsReadRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/mark_stream_as_read"
	)
	if r.channelId == nil {
		return nil, nil, reportError("channelId is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "stream_id", r.channelId)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type MarkTopicAsReadRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	channelId  *int64
	topicName  *string
}

// The Id of the channel to access.
func (r MarkTopicAsReadRequest) ChannelId(channelId int64) MarkTopicAsReadRequest {
	r.channelId = &channelId
	return r
}

// The name of the topic whose messages should be marked as read.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.
//
// **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [POST /register]: https://zulip.com/api/register-queue
func (r MarkTopicAsReadRequest) TopicName(topicName string) MarkTopicAsReadRequest {
	r.topicName = &topicName
	return r
}

func (r MarkTopicAsReadRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.MarkTopicAsReadExecute(r)
}

// MarkTopicAsRead Mark messages in a topic as read
//
// Mark all the unread messages in a topic as read.
//
// *Changes**: Deprecated; clients should use the [update personal message flags for narrow] endpoint instead
// as this endpoint will be removed in a future release.
//
// # Deprecated
//
// [update personal message flags for narrow]: https://zulip.com/api/update-message-flags-for-narrow
func (c *simpleClient) MarkTopicAsRead(ctx context.Context) MarkTopicAsReadRequest {
	return MarkTopicAsReadRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
// Deprecated
func (c *simpleClient) MarkTopicAsReadExecute(r MarkTopicAsReadRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/mark_topic_as_read"
	)
	if r.channelId == nil {
		return nil, nil, reportError("channelId is required and must be specified")
	}
	if r.topicName == nil {
		return nil, nil, reportError("topicName is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "stream_id", r.channelId)
	addParam(form, "topic_name", r.topicName)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RemoveReactionRequest struct {
	ctx          context.Context
	apiService   MessagesAPI
	messageId    int64
	emojiName    *string
	emojiCode    *string
	reactionType *ReactionType
}

// The target emoji's human-readable name.  To find an emoji's name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji's text name.
func (r RemoveReactionRequest) EmojiName(emojiName string) RemoveReactionRequest {
	r.emojiName = &emojiName
	return r
}

// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  For most API clients, you won't need this, but it's important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the `emoji_code` in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a "different" emoji.
func (r RemoveReactionRequest) EmojiCode(emojiCode string) RemoveReactionRequest {
	r.emojiCode = &emojiCode
	return r
}

// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`. If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity.
//
// Supported values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
//
// **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji]; previously, this endpoint assumed `unicode_emoji` if this parameter was not specified.
//
// [custom emoji]: https://zulip.com/help/custom-emoji
func (r RemoveReactionRequest) ReactionType(reactionType ReactionType) RemoveReactionRequest {
	r.reactionType = &reactionType
	return r
}

func (r RemoveReactionRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.RemoveReactionExecute(r)
}

// RemoveReaction Remove an emoji reaction
//
// Remove an [emoji reaction] from a message.
//
// [emoji reaction]: https://zulip.com/help/emoji-reactions
func (c *simpleClient) RemoveReaction(ctx context.Context, messageId int64) RemoveReactionRequest {
	return RemoveReactionRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) RemoveReactionExecute(r RemoveReactionRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/messages/{message_id}/reactions"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(form, "emoji_name", r.emojiName)
	addOptionalParam(form, "emoji_code", r.emojiCode)
	addOptionalParam(form, "reaction_type", r.reactionType)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type RenderMessageRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	content    *string
}

// The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`] endpoint to determine the maximum message size.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r RenderMessageRequest) Content(content string) RenderMessageRequest {
	r.content = &content
	return r
}

func (r RenderMessageRequest) Execute() (*RenderMessageResponse, *http.Response, error) {
	return r.apiService.RenderMessageExecute(r)
}

// RenderMessage Render a message
//
// Render a message to HTML.
func (c *simpleClient) RenderMessage(ctx context.Context) RenderMessageRequest {
	return RenderMessageRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) RenderMessageExecute(r RenderMessageRequest) (*RenderMessageResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &RenderMessageResponse{}
		endpoint = "/messages/render"
	)
	if r.content == nil {
		return nil, nil, reportError("content is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "content", r.content)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type ReportMessageRequest struct {
	ctx         context.Context
	apiService  MessagesAPI
	messageId   int64
	reportType  *string
	description *string
}

// The reason that best describes why the current user is reporting the target message for moderation.
func (r ReportMessageRequest) ReportType(reportType string) ReportMessageRequest {
	r.reportType = &reportType
	return r
}

// A short description with additional context about why the current user is reporting the target message for moderation.  Clients should limit this string to a maximum length of 1000 characters.  If the `report_type` parameter is `"other"`, this parameter is required, and its value cannot be an empty string.
func (r ReportMessageRequest) Description(description string) ReportMessageRequest {
	r.description = &description
	return r
}

func (r ReportMessageRequest) Execute() (*Response, *http.Response, error) {
	return r.apiService.ReportMessageExecute(r)
}

// ReportMessage Report a message
//
// Sends a notification to the organization's moderation request channel,
// if it is configured, that reports the targeted message for review and
// moderation.
//
// Clients should check the `moderation_request_channel` realm setting to
// decide whether to show the option to report messages in the UI.
//
// If the `report_type` parameter value is `"other"`, the `description`
// parameter is required. Clients should also enforce and communicate this
// behavior in the UI.
//
// *Changes**: New in Zulip 11.0 (feature level 382). This API builds on
// the `moderation_request_channel` realm setting, which was added in
// feature level 331.
func (c *simpleClient) ReportMessage(ctx context.Context, messageId int64) ReportMessageRequest {
	return ReportMessageRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) ReportMessageExecute(r ReportMessageRequest) (*Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &Response{}
		endpoint = "/messages/{message_id}/report"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	if r.reportType == nil {
		return nil, nil, reportError("reportType is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "report_type", r.reportType)
	addOptionalParam(form, "description", r.description)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type SendMessageRequest struct {
	ctx           context.Context
	apiService    MessagesAPI
	recipientType *RecipientType
	to            *Recipient
	content       *string
	topic         *string
	queueId       *string
	localId       *string
	readBySender  *bool
}

// The type of message to be sent.  `RecipientTypeDirect` for a direct message and `RecipientTypeStream` or `RecipientTypeChannel` for a channel message.
//
// **Changes**: In Zulip 9.0 (feature level 248), `RecipientTypeChannel` was added as an additional value for this parameter to request a channel message.  In Zulip 7.0 (feature level 174), `RecipientTypeDirect` was added as the preferred way to request a direct message, deprecating the original `RecipientTypePrivate`. While `RecipientTypePrivate` is still supported for requesting direct messages, clients are encouraged to use to the modern convention with servers that support it, because support for `RecipientTypePrivate` will eventually be removed.
func (r SendMessageRequest) RecipientType(recipientType RecipientType) SendMessageRequest {
	r.recipientType = &recipientType
	return r
}

func (r SendMessageRequest) To(to Recipient) SendMessageRequest {
	r.to = &to
	return r
}

// The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`] endpoint to determine the maximum message size.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r SendMessageRequest) Content(content string) SendMessageRequest {
	r.content = &content
	return r
}

// The topic of the message. Only required for channel messages (`"type": "stream"` or `"type": "channel"`), ignored otherwise.  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length.  Note: When `"(no topic)"` or the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.  When [topics are required], this parameter can't be `"(no topic)"`, an empty string, or the value of `realm_empty_topic_display_name`.
//
// **Changes**: Before Zulip 10.0 (feature level 370), `"(no topic)"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as `subject`, which is currently a deprecated alias.
//
// [`POST /register`]: https://zulip.com/api/register-queue
// [topics are required]: https://zulip.com/help/require-topics
func (r SendMessageRequest) Topic(topic string) SendMessageRequest {
	r.topic = &topic
	return r
}

// For clients supporting [local echo], the [event queue] Id for the client. If passed, `local_id` is required. If the message is successfully sent, the server will include `local_id` in the `message` event that the client with this `queue_id` will receive notifying it of the new message via [`GET /events`]. This lets the client know unambiguously that it should replace the locally echoed message, rather than adding this new message (which would be correct if the user had sent the new message from another device).
//
// [local echo]: https://zulip.readthedocs.io/en/latest/subsystems/sending-messages.html#local-echo
// [event queue]: https://zulip.com/api/register-queue
// [`GET /events`]: https://zulip.com/api/get-events
func (r SendMessageRequest) QueueId(queueId string) SendMessageRequest {
	r.queueId = &queueId
	return r
}

// For clients supporting local echo, a unique string-format identifier chosen freely by the client; the server will pass it back to the client without inspecting it, as described in the `queue_id` description.
func (r SendMessageRequest) LocalId(localId string) SendMessageRequest {
	r.localId = &localId
	return r
}

// Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name.
//
// **Changes**: New in Zulip 8.0 (feature level 236).
func (r SendMessageRequest) ReadBySender(readBySender bool) SendMessageRequest {
	r.readBySender = &readBySender
	return r
}

func (r SendMessageRequest) Execute() (*SendMessageResponse, *http.Response, error) {
	return r.apiService.SendMessageExecute(r)
}

// SendMessage Send a message
//
// Send a [channel message] or a
// [direct message].
//
// [channel message]: https://zulip.com/help/introduction-to-topics
// [direct message]: https://zulip.com/help/direct-messages
func (c *simpleClient) SendMessage(ctx context.Context) SendMessageRequest {
	return SendMessageRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) SendMessageExecute(r SendMessageRequest) (*SendMessageResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &SendMessageResponse{}
		endpoint = "/messages"
	)
	if r.recipientType == nil {
		return nil, nil, reportError("recipientType is required and must be specified")
	}
	if r.to == nil {
		return nil, nil, reportError("to is required and must be specified")
	}
	if r.content == nil {
		return nil, nil, reportError("content is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "type", r.recipientType)
	addParam(form, "to", r.to)
	addParam(form, "content", r.content)
	addOptionalParam(form, "topic", r.topic)
	addOptionalParam(form, "queue_id", r.queueId)
	addOptionalParam(form, "local_id", r.localId)
	addOptionalParam(form, "read_by_sender", r.readBySender)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateMessageRequest struct {
	ctx                         context.Context
	apiService                  MessagesAPI
	messageId                   int64
	topic                       *string
	propagateMode               *string
	sendNotificationToOldThread *bool
	sendNotificationToNewThread *bool
	content                     *string
	prevContentSha256           *string
	channelId                   *int64
}

// The topic to move the message(s) to, to request changing the topic.  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length  Should only be sent when changing the topic, and will throw an error if the target message is not a channel message.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.  When [topics are required], this parameter can't be `"(no topic)"`, an empty string, or the value of `realm_empty_topic_display_name`.  You can [resolve topics] by editing the topic to `✔ {original_topic}` with the `propagate_mode` parameter set to `"change_all"`. The empty string topic cannot be marked as resolved.
//
// **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as `subject`, which is currently a deprecated alias.
//
// [`POST /register`]: https://zulip.com/api/register-queue
// [topics are required]: https://zulip.com/help/require-topics
// [resolve topics]: https://zulip.com/help/resolve-a-topic
func (r UpdateMessageRequest) Topic(topic string) UpdateMessageRequest {
	r.topic = &topic
	return r
}

// Which message(s) should be edited:  - `"change_later"`: The target message and all following messages. - `"change_one"`: Only the target message. - `"change_all"`: All messages in this topic.  Only the default value of `"change_one"` is valid when editing only the content of a message.  This parameter determines both which messages get moved and also whether clients that are currently narrowed to the topic containing the message should navigate or adjust their compose box recipient to point to the post-edit channel/topic.
func (r UpdateMessageRequest) PropagateMode(propagateMode string) UpdateMessageRequest {
	r.propagateMode = &propagateMode
	return r
}

// Whether to send an automated message to the old topic to notify users where the messages were moved to.
//
// **Changes**: Before Zulip 6.0 (feature level 152), this parameter had a default of `true` and was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).
func (r UpdateMessageRequest) SendNotificationToOldThread(sendNotificationToOldThread bool) UpdateMessageRequest {
	r.sendNotificationToOldThread = &sendNotificationToOldThread
	return r
}

// Whether to send an automated message to the new topic to notify users where the messages came from.  If the move is just [resolving/unresolving a topic], this parameter will not trigger an additional notification.
//
// **Changes**: Before Zulip 6.0 (feature level 152), this parameter was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).
//
// [resolving/unresolving a topic]: https://zulip.com/help/resolve-a-topic
func (r UpdateMessageRequest) SendNotificationToNewThread(sendNotificationToNewThread bool) UpdateMessageRequest {
	r.sendNotificationToNewThread = &sendNotificationToNewThread
	return r
}

// The updated content of the target message.  Clients should use the `max_message_length` returned by the [`POST /register`] endpoint to determine the maximum message size.  Note that a message's content and channel cannot be changed at the same time, so sending both `content` and `stream_id` parameters will throw an error.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateMessageRequest) Content(content string) UpdateMessageRequest {
	r.content = &content
	return r
}

// An optional SHA-256 hash of the previous raw content of the message that the client has at the time of the request.  If provided, the server will return an error if it does not match the SHA-256 hash of the message's content stored in the database.  Clients can use this feature to prevent races where multiple clients save conflicting edits to a message.
//
// **Changes**: New in Zulip 11.0 (feature level 379).
func (r UpdateMessageRequest) PrevContentSha256(prevContentSha256 string) UpdateMessageRequest {
	r.prevContentSha256 = &prevContentSha256
	return r
}

// The channel Id to move the message(s) to, to request moving messages to another channel.  Should only be sent when changing the channel, and will throw an error if the target message is not a channel message.  Note that a message's content and channel cannot be changed at the same time, so sending both `content` and `stream_id` parameters will throw an error.
//
// **Changes**: New in Zulip 3.0 (feature level 1).
func (r UpdateMessageRequest) ChannelId(channelId int64) UpdateMessageRequest {
	r.channelId = &channelId
	return r
}

func (r UpdateMessageRequest) Execute() (*UpdateMessageResponse, *http.Response, error) {
	return r.apiService.UpdateMessageExecute(r)
}

// UpdateMessage Edit a message
//
// Update the content, topic, or channel of the message with the specified
// Id.
//
// You can [resolve topics] by editing the topic to
// `✔ {original_topic}` with the `propagate_mode` parameter set to
// `"change_all"`.
//
// See [configuring message editing] for detailed
// documentation on when users are allowed to edit message content, and
// [restricting moving messages] for detailed
// documentation on when users are allowed to change a message's topic
// and/or channel.
//
// The relevant realm settings in the API that are related to the above
// linked documentation on when users are allowed to update messages are:
//
//   - `allow_message_editing`
//   - `can_resolve_topics_group`
//   - `can_move_messages_between_channels_group`
//   - `can_move_messages_between_topics_group`
//   - `message_content_edit_limit_seconds`
//   - `move_messages_within_stream_limit_seconds`
//   - `move_messages_between_streams_limit_seconds`
//
// More details about these realm settings can be found in the
// [`POST /register`] response or in the documentation
// of the [`realm op: update_dict`] event in [`GET /events`].
//
// *Changes**: Prior to Zulip 10.0 (feature level 367), the permission for
// resolving a topic was managed by `can_move_messages_between_topics_group`.
// As of this feature level, users belonging to the `can_resolve_topics_group`
// will have the permission to [resolve topics] in the organization.
//
// In Zulip 10.0 (feature level 316), `edit_topic_policy`
// was removed and replaced by `can_move_messages_between_topics_group`
// realm setting.
//
// *Changes**: In Zulip 10.0 (feature level 310), `move_messages_between_streams_policy`
// was removed and replaced by `can_move_messages_between_channels_group`
// realm setting.
//
// Prior to Zulip 7.0 (feature level 172), anyone could add a
// topic to channel messages without a topic, regardless of the organization's
// [topic editing permissions]. As of this
// feature level, messages without topics have the same restrictions for
// topic edits as messages with topics.
//
// Before Zulip 7.0 (feature level 172), by using the `change_all` value for
// the `propagate_mode` parameter, users could move messages after the
// organization's configured time limits for changing a message's topic or
// channel had passed. As of this feature level, the server will [return an error] with `"code":
// "MOVE_MESSAGES_TIME_LIMIT_EXCEEDED"` if users, other than organization
// administrators or moderators, try to move messages after these time
// limits have passed.
//
// Before Zulip 7.0 (feature level 162), users who were not administrators or
// moderators could only edit topics if the target message was sent within the
// last 3 days. As of this feature level, that time limit is now controlled by
// the realm setting `move_messages_within_stream_limit_seconds`. Also at this
// feature level, a similar time limit for moving messages between channels was
// added, controlled by the realm setting
// `move_messages_between_streams_limit_seconds`. Previously, all users who
// had permission to move messages between channels did not have any time limit
// restrictions when doing so.
//
// Before Zulip 7.0 (feature level 159), editing channels and topics of messages
// was forbidden if the realm setting for `allow_message_editing` was `false`,
// regardless of an organization's configuration for the realm settings
// `edit_topic_policy` or `move_messages_between_streams_policy`.
//
// Before Zulip 7.0 (feature level 159), message senders were allowed to edit
// the topic of their messages indefinitely.
//
// In Zulip 5.0 (feature level 75), the `edit_topic_policy` realm setting
// was added, replacing the `allow_community_topic_editing` boolean.
//
// In Zulip 4.0 (feature level 56), the `move_messages_between_streams_policy`
// realm setting was added.
//
// [configuring message editing]: https://zulip.com/help/restrict-message-editing-and-deletion
// [restricting moving messages]: https://zulip.com/help/restrict-moving-messages
// [resolve topics]: https://zulip.com/help/resolve-a-topic
// [`POST /register`]: https://zulip.com/api/register-queue
// [`GET /events`]: https://zulip.com/api/get-events
// [topic editing permissions]: https://zulip.com/help/restrict-moving-messages
// [return an error]: https://zulip.com/api/update-message#response
// [`realm op: update_dict`]: https://zulip.com/api/get-events#realm-update_dict
func (c *simpleClient) UpdateMessage(ctx context.Context, messageId int64) UpdateMessageRequest {
	return UpdateMessageRequest{
		apiService: c,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateMessageExecute(r UpdateMessageRequest) (*UpdateMessageResponse, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UpdateMessageResponse{}
		endpoint = "/messages/{message_id}"
	)

	endpoint = strings.Replace(endpoint, "{message_id}", idToString(r.messageId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addOptionalParam(form, "topic", r.topic)
	addOptionalParam(form, "propagate_mode", r.propagateMode)
	addOptionalParam(form, "send_notification_to_old_thread", r.sendNotificationToOldThread)
	addOptionalParam(form, "send_notification_to_new_thread", r.sendNotificationToNewThread)
	addOptionalParam(form, "content", r.content)
	addOptionalParam(form, "prev_content_sha256", r.prevContentSha256)
	addOptionalParam(form, "stream_id", r.channelId)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateMessageFlagsRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	messages   *[]int64
	op         *string
	flag       *string
}

// An array containing the Ids of the target messages.
func (r UpdateMessageFlagsRequest) Messages(messages []int64) UpdateMessageFlagsRequest {
	r.messages = &messages
	return r
}

// Whether to `add` the flag or `remove` it.
func (r UpdateMessageFlagsRequest) Op(op string) UpdateMessageFlagsRequest {
	r.op = &op
	return r
}

// The flag that should be added/removed.
func (r UpdateMessageFlagsRequest) Flag(flag string) UpdateMessageFlagsRequest {
	r.flag = &flag
	return r
}

func (r UpdateMessageFlagsRequest) Execute() (*UpdateMessageFlagsResponse, *http.Response, error) {
	return r.apiService.UpdateMessageFlagsExecute(r)
}

// UpdateMessageFlags Update personal message flags
//
// Add or remove personal message flags like `read` and `starred`
// on a collection of message Ids.
//
// See also the endpoint for [updating flags on a range of messages within a narrow].
//
// [updating flags on a range of messages within a narrow]: https://zulip.com/api/update-message-flags-for-narrow
func (c *simpleClient) UpdateMessageFlags(ctx context.Context) UpdateMessageFlagsRequest {
	return UpdateMessageFlagsRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateMessageFlagsExecute(r UpdateMessageFlagsRequest) (*UpdateMessageFlagsResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UpdateMessageFlagsResponse{}
		endpoint = "/messages/flags"
	)
	if r.messages == nil {
		return nil, nil, reportError("messages is required and must be specified")
	}
	if r.op == nil {
		return nil, nil, reportError("op is required and must be specified")
	}
	if r.flag == nil {
		return nil, nil, reportError("flag is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "messages", r.messages)
	addParam(form, "op", r.op)
	addParam(form, "flag", r.flag)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UpdateMessageFlagsForNarrowRequest struct {
	ctx           context.Context
	apiService    MessagesAPI
	anchor        *string
	numBefore     *int32
	numAfter      *int32
	narrow        *Narrow
	op            *string
	flag          *string
	includeAnchor *bool
}

// Integer message Id to anchor updating of flags. Supports special string values for when the client wants the server to compute the anchor to use:  - `newest`: The most recent message. - `oldest`: The oldest message. - `first_unread`: The oldest unread message matching the   query, if any; otherwise, the most recent message.
func (r UpdateMessageFlagsForNarrowRequest) Anchor(anchor string) UpdateMessageFlagsForNarrowRequest {
	r.anchor = &anchor
	return r
}

// Limit the number of messages preceding the anchor in the update range. The server may decrease this to bound transaction sizes.
func (r UpdateMessageFlagsForNarrowRequest) NumBefore(numBefore int32) UpdateMessageFlagsForNarrowRequest {
	r.numBefore = &numBefore
	return r
}

// Limit the number of messages following the anchor in the update range. The server may decrease this to bound transaction sizes.
func (r UpdateMessageFlagsForNarrowRequest) NumAfter(numAfter int32) UpdateMessageFlagsForNarrowRequest {
	r.numAfter = &numAfter
	return r
}

// The narrow you want update flags within. See how to [construct a narrow].  Note that, when adding the `read` flag to messages, clients should consider including a narrow with the `is:unread` filter as an optimization. Including that filter takes advantage of the fact that the server has a database index for unread messages.
//
// **Changes**: See [changes section] of search/narrow filter documentation.
//
// [construct a narrow]: https://zulip.com/api/construct-narrow
// [changes section]: https://zulip.com/api/construct-narrow#changes
func (r UpdateMessageFlagsForNarrowRequest) Narrow(narrow *Narrow) UpdateMessageFlagsForNarrowRequest {
	r.narrow = narrow
	return r
}

// Whether to `add` the flag or `remove` it.
func (r UpdateMessageFlagsForNarrowRequest) Op(op string) UpdateMessageFlagsForNarrowRequest {
	r.op = &op
	return r
}

// The flag that should be added/removed. See [available flags].
//
// [available flags]: https://zulip.com/api/update-message-flags#available-flags
func (r UpdateMessageFlagsForNarrowRequest) Flag(flag string) UpdateMessageFlagsForNarrowRequest {
	r.flag = &flag
	return r
}

// Whether a message with the specified Id matching the narrow should be included in the update range.
func (r UpdateMessageFlagsForNarrowRequest) IncludeAnchor(includeAnchor bool) UpdateMessageFlagsForNarrowRequest {
	r.includeAnchor = &includeAnchor
	return r
}

func (r UpdateMessageFlagsForNarrowRequest) Execute() (*UpdateMessageFlagsForNarrowResponse, *http.Response, error) {
	return r.apiService.UpdateMessageFlagsForNarrowExecute(r)
}

// UpdateMessageFlagsForNarrow Update personal message flags for narrow
//
// Add or remove personal message flags like `read` and `starred`
// on a range of messages within a narrow.
//
// See also [the endpoint for updating flags on specific message Ids].
//
// *Changes**: New in Zulip 6.0 (feature level 155).
//
// [the endpoint for updating flags on specific message Ids]: https://zulip.com/api/update-message-flags
func (c *simpleClient) UpdateMessageFlagsForNarrow(ctx context.Context) UpdateMessageFlagsForNarrowRequest {
	return UpdateMessageFlagsForNarrowRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateMessageFlagsForNarrowExecute(r UpdateMessageFlagsForNarrowRequest) (*UpdateMessageFlagsForNarrowResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UpdateMessageFlagsForNarrowResponse{}
		endpoint = "/messages/flags/narrow"
	)
	if r.anchor == nil {
		return nil, nil, reportError("anchor is required and must be specified")
	}
	if r.numBefore == nil {
		return nil, nil, reportError("numBefore is required and must be specified")
	}
	if *r.numBefore < 0 {
		return nil, nil, reportError("numBefore must be greater than 0")
	}
	if r.numAfter == nil {
		return nil, nil, reportError("numAfter is required and must be specified")
	}
	if *r.numAfter < 0 {
		return nil, nil, reportError("numAfter must be greater than 0")
	}
	if r.narrow == nil {
		return nil, nil, reportError("narrow is required and must be specified")
	}
	if r.op == nil {
		return nil, nil, reportError("op is required and must be specified")
	}
	if r.flag == nil {
		return nil, nil, reportError("flag is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "anchor", r.anchor)
	addOptionalParam(form, "include_anchor", r.includeAnchor)
	addParam(form, "num_before", r.numBefore)
	addParam(form, "num_after", r.numAfter)
	addParam(form, "narrow", r.narrow)
	addParam(form, "op", r.op)
	addParam(form, "flag", r.flag)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type UploadFileRequest struct {
	ctx        context.Context
	apiService MessagesAPI
	filename   *os.File
}

func (r UploadFileRequest) Filename(filename *os.File) UploadFileRequest {
	r.filename = filename
	return r
}

func (r UploadFileRequest) Execute() (*UploadFileResponse, *http.Response, error) {
	return r.apiService.UploadFileExecute(r)
}

// UploadFile Upload a file
//
// [Upload] a single file and get the corresponding URL.
//
// Initially, only you will be able to access the link. To share the
// uploaded file, you'll need to [send a message]
// containing the resulting link. Users who can already access the link
// can reshare it with other users by sending additional Zulip messages
// containing the link.
//
// The maximum allowed file size is available in the `max_file_upload_size_mib`
// field in the [`POST /register`] response. Note that
// large files (25MB+) may fail to upload using this API endpoint due to
// network-layer timeouts, depending on the quality of your connection to the
// Zulip server.
//
// For uploading larger files, `/api/v1/tus` is an endpoint implementing the
// [`tus` resumable upload protocol],
// which supports uploading arbitrarily large files limited only by the server's
// `max_file_upload_size_mib` (Configured via `MAX_FILE_UPLOAD_SIZE` in
// `/etc/zulip/settings.py`). Clients which send authenticated credentials
// (either via browser-based cookies, or API key via `Authorization` header) may
// use this endpoint to upload files.
//
// *Changes**: The `api/v1/tus` endpoint supporting resumable uploads was
// introduced in Zulip 10.0 (feature level 296). Previously,
// `max_file_upload_size_mib` was typically 25MB.
//
// [send a message]: https://zulip.com/api/send-message
// [Upload]: https://zulip.com/help/share-and-upload-files
// [`POST /register`]: https://zulip.com/api/register-queue
// [`tus` resumable upload protocol]: https://tus.io/protocols/resumable-upload
//
// [uploaded-files]: https://zulip.com/help/manage-your-uploaded-files
func (c *simpleClient) UploadFile(ctx context.Context) UploadFileRequest {
	return UploadFileRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) UploadFileExecute(r UploadFileRequest) (*UploadFileResponse, *http.Response, error) {
	var (
		method    = http.MethodPost
		headers   = make(map[string]string)
		query     = url.Values{}
		form      = url.Values{}
		formFiles []formFile
		response  = &UploadFileResponse{}
		endpoint  = "/user_uploads"
	)
	headers["Content-Type"] = "multipart/form-data"
	headers["Accept"] = "application/json"

	var filenameLocalVarFormFileName string
	var filenameLocalVarFileName string
	var filenameLocalVarFileBytes []byte

	filenameLocalVarFormFileName = "filename"
	filenameLocalVarFile := r.filename

	if filenameLocalVarFile != nil {
		fbs, _ := io.ReadAll(filenameLocalVarFile)

		filenameLocalVarFileBytes = fbs
		filenameLocalVarFileName = filenameLocalVarFile.Name()
		filenameLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: filenameLocalVarFileBytes, fileName: filenameLocalVarFileName, formFileName: filenameLocalVarFormFileName})
	}
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, formFiles)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
