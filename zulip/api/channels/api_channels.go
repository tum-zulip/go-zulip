package channels

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
	. "github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"
)

type APIChannels interface {

	// AddDefaultChannel Add a default channel
	//
	// Add a channel to the set of [default channels]
	// for new users joining the organization.
	//
	// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
	//
	AddDefaultChannel(ctx context.Context) AddDefaultChannelRequest

	// AddDefaultChannelExecute executes the request
	AddDefaultChannelExecute(r AddDefaultChannelRequest) (*zulip.Response, *http.Response, error)

	// ArchiveChannel Archive a channel
	//
	// [Archive the channel] with the Id `channelId`.
	//
	// [Archive the channel]: https://zulip.com/help/archive-a-channel
	ArchiveChannel(ctx context.Context, channelId int64) ArchiveChannelRequest

	// ArchiveChannelExecute executes the request
	ArchiveChannelExecute(r ArchiveChannelRequest) (*zulip.Response, *http.Response, error)

	// CreateBigBlueButtonVideoCall Create BigBlueButton video call
	//
	// Create a video call URL for a BigBlueButton video call.
	// Requires [BigBlueButton 2.4+]
	// to be configured on the Zulip server.
	//
	// The acting user will be given the moderator role on the call.
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 337), every
	// user was given the moderator role on BigBlueButton calls, via
	// encoding a moderator password in the generated URLs.
	//
	// [BigBlueButton 2.4+]: https://zulip.com/integrations/doc/big-blue-button
	CreateBigBlueButtonVideoCall(ctx context.Context) CreateBigBlueButtonVideoCallRequest

	// CreateBigBlueButtonVideoCallExecute executes the request
	CreateBigBlueButtonVideoCallExecute(r CreateBigBlueButtonVideoCallRequest) (*CreateBigBlueButtonVideoCallResponse, *http.Response, error)

	// CreateChannel Create a channel
	//
	// Create a new [channel] and optionally
	// subscribe users to the newly created channel. The initial [channel settings]
	// will be determined by the optional parameters, like `invite_only`, detailed below.
	//
	// *Changes**: New in Zulip 11.0 (feature level 417). Previously, this was only possible via
	// the [`POST /api/subscribe`] endpoint, which handled both creation and subscription.
	//
	// [channel]: https://zulip.com/help/create-channels
	// [channel settings]: https://zulip.com/api/update-stream
	// [`POST /api/subscribe`]: https://zulip.com/api/subscribe
	CreateChannel(ctx context.Context) CreateChannelRequest

	// CreateChannelExecute executes the request
	CreateChannelExecute(r CreateChannelRequest) (*CreateChannelResponse, *http.Response, error)

	// CreateChannelFolder Create a channel folder
	//
	// Create a new channel folder, that will be used to organize
	// channels in left sidebar.
	//
	// Only organization administrators can create a new channel
	// folder.
	//
	// *Changes**: New in Zulip 11.0 (feature level 389).
	//
	CreateChannelFolder(ctx context.Context) CreateChannelFolderRequest

	// CreateChannelFolderExecute executes the request
	CreateChannelFolderExecute(r CreateChannelFolderRequest) (*CreateChannelFolderResponse, *http.Response, error)

	// DeleteTopic Delete a topic
	//
	// Delete all messages in a topic.
	//
	// Topics are a field on messages (not an independent data structure), so
	// deleting all the messages in the topic deletes the topic from Zulip.
	//
	// Because this endpoint deletes messages in batches, it is possible for
	// the request to time out after only deleting some messages in the topic.
	// When this happens, the `complete` boolean field in the success response
	// will be `false`. Clients should repeat the request when handling such a
	// response. If all messages in the topic were deleted, then the success
	// response will return `"complete": true`.
	//
	// *Changes**: Before Zulip 9.0 (feature level 256), the server never sent
	// [`stream` op: `update`] events with an updated `first_message_id` for a channel when the oldest message that
	// had been sent to it changed.
	//
	// Before Zulip 8.0 (feature level 211), if the server's
	// processing was interrupted by a timeout, but some messages in the topic
	// were deleted, then it would return `"result": "partially_completed"`,
	// along with a `code` field for an error string, in the success response
	// to indicate that there was a timeout and that the client should repeat
	// the request.
	//
	// As of Zulip 6.0 (feature level 154), instead of returning an error
	// response when a request times out after successfully deleting some of
	// the messages in the topic, a success response is returned with
	// `"result": "partially_completed"` to indicate that some messages were
	// deleted.
	//
	// Before Zulip 6.0 (feature level 147), this request did a single atomic
	// operation, which could time out for very large topics. As of this
	// feature level, messages are deleted in batches, starting with the newest
	// messages, so that progress is made even if the request times out and
	// returns an error.
	//
	// [`stream` op: `update`]: https://zulip.com/api/get-events#stream-update
	DeleteTopic(ctx context.Context, channelId int64) DeleteTopicRequest

	// DeleteTopicExecute executes the request
	DeleteTopicExecute(r DeleteTopicRequest) (*MarkAllAsReadResponse, *http.Response, error)

	// GetChannelFolders Get channel folders
	//
	// Fetches all of the channel folders in the organization.
	// The folders are sorted by the `order` field.
	//
	// *Changes**: Before Zulip 11.0 (feature level 414),
	// these were sorted by Id. (The `order` field didn't exist).
	//
	// New in Zulip 11.0 (feature level 389).
	//
	GetChannelFolders(ctx context.Context) GetChannelFoldersRequest

	// GetChannelFoldersExecute executes the request
	GetChannelFoldersExecute(r GetChannelFoldersRequest) (*GetChannelFoldersResponse, *http.Response, error)

	// GetChannelById Get a channel by Id
	//
	// Fetch details for the channel with the Id `channelId`.
	//
	// *Changes**: New in Zulip 6.0 (feature level 132).
	//
	GetChannelById(ctx context.Context, channelId int64) GetChannelByIdRequest

	// GetChannelByIdExecute executes the request
	GetChannelByIdExecute(r GetChannelByIdRequest) (*GetChannelResponse, *http.Response, error)

	// GetChannelEmailAddress Get channel's email address
	//
	// Get email address of a channel.
	//
	// *Changes**: New in Zulip 8.0 (feature level 226).
	//
	GetChannelEmailAddress(ctx context.Context, channelId int64) GetChannelEmailAddressRequest

	// GetChannelEmailAddressExecute executes the request
	GetChannelEmailAddressExecute(r GetChannelEmailAddressRequest) (*GetChannelEmailAddressResponse, *http.Response, error)

	// GetChannelId Get channel Id
	//
	// Get the unique Id of a given channel.
	//
	GetChannelId(ctx context.Context) GetChannelIdRequest

	// GetChannelIdExecute executes the request
	GetChannelIdExecute(r GetChannelIdRequest) (*GetChannelIdResponse, *http.Response, error)

	// GetChannelTopics Get topics in a channel
	//
	// Get all topics the user has access to in a specific channel.
	//
	// Note that for [private channels with protected history],
	// the user will only have access to topics of messages sent after they
	// [subscribed to] the channel. Similarly, a user's
	// [bot] will only have access to messages
	// sent after the bot was subscribed to the channel, instead of when the
	// user subscribed.
	//
	// [subscribed to]: https://zulip.com/api/subscribe
	// [bot]: https://zulip.com/help/bots-overview#bot-type
	// [private channels with protected history]: https://zulip.com/help/channel-permissions#private-channels
	GetChannelTopics(ctx context.Context, channelId int64) GetChannelTopicsRequest

	// GetChannelTopicsExecute executes the request
	GetChannelTopicsExecute(r GetChannelTopicsRequest) (*GetChannelTopicsResponse, *http.Response, error)

	// GetChannels Get all channels
	//
	// Get all channels that the user [has access to].
	//
	// [has access to]: https://zulip.com/help/channel-permissions
	GetChannels(ctx context.Context) GetChannelsRequest

	// GetChannelsExecute executes the request
	GetChannelsExecute(r GetChannelsRequest) (*GetChannelsResponse, *http.Response, error)

	// GetSubscribers Get channel subscribers
	//
	// Get all users subscribed to a channel.
	//
	GetSubscribers(ctx context.Context, channelId int64) GetSubscribersRequest

	// GetSubscribersExecute executes the request
	GetSubscribersExecute(r GetSubscribersRequest) (*GetSubscribersResponse, *http.Response, error)

	// GetSubscriptionStatus Get subscription status
	//
	// Check whether a user is subscribed to a channel.
	//
	// *Changes**: New in Zulip 3.0 (feature level 12).
	//
	GetSubscriptionStatus(ctx context.Context, userId int64, channelId int64) GetSubscriptionStatusRequest

	// GetSubscriptionStatusExecute executes the request
	GetSubscriptionStatusExecute(r GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, *http.Response, error)

	// GetSubscriptions Get subscribed channels
	//
	// Get all channels that the user is subscribed to.
	//
	GetSubscriptions(ctx context.Context) GetSubscriptionsRequest

	// GetSubscriptionsExecute executes the request
	GetSubscriptionsExecute(r GetSubscriptionsRequest) (*GetSubscriptionsResponse, *http.Response, error)

	// MuteTopic Topic muting
	//
	// [Mute or unmute a topic] within a channel that
	// the current user is subscribed to.
	//
	// *Changes**: Deprecated in Zulip 7.0 (feature level 170). Clients connecting
	// to newer servers should use the [POST /user_topics]
	// endpoint, as this endpoint may be removed in a future release.
	//
	// Before Zulip 7.0 (feature level 169), this endpoint
	// returned an error if asked to mute a topic that was already muted
	// or asked to unmute a topic that had not previously been muted.
	//
	//
	// Deprecated
	//
	// [Mute or unmute a topic]: https://zulip.com/help/mute-a-topic
	// [POST /user_topics]: https://zulip.com/api/update-user-topic
	MuteTopic(ctx context.Context) MuteTopicRequest

	// MuteTopicExecute executes the request
	// Deprecated
	MuteTopicExecute(r MuteTopicRequest) (*zulip.Response, *http.Response, error)

	// PatchChannelFolders Reorder channel folders
	//
	// Given an array of channel folder Ids, this method will set the `order`
	// property of all of the channel folders in the organization according to
	// the order of the channel folder Ids specified in the request.
	//
	// *Changes**: New in Zulip 11.0 (feature level 414).
	//
	PatchChannelFolders(ctx context.Context) PatchChannelFoldersRequest

	// PatchChannelFoldersExecute executes the request
	PatchChannelFoldersExecute(r PatchChannelFoldersRequest) (*zulip.Response, *http.Response, error)

	// RemoveDefaultChannel Remove a default channel
	//
	// Remove a channel from the set of [default channels]
	// for new users joining the organization.
	//
	// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
	//
	RemoveDefaultChannel(ctx context.Context) RemoveDefaultChannelRequest

	// RemoveDefaultChannelExecute executes the request
	RemoveDefaultChannelExecute(r RemoveDefaultChannelRequest) (*zulip.Response, *http.Response, error)

	// Subscribe Subscribe to a channel
	//
	// Subscribe one or more users to one or more channels.
	//
	// If any of the specified channels do not exist, they are automatically
	// created. The initial [channel settings] will be determined
	// by the optional parameters, like `invite_only`, detailed below.
	//
	// Note that the ability to subscribe oneself and/or other users
	// to a specified channel depends on the [channel's permissions settings].
	//
	// *Changes**: Before Zulip 10.0 (feature level 362),
	// subscriptions in archived channels could not be modified.
	//
	// Before Zulip 10.0 (feature level 357), the
	// `can_subscribe_group` permission, which allows members of the
	// group to subscribe themselves to the channel, did not exist.
	//
	// Before Zulip 10.0 (feature level 349), a user cannot subscribe
	// other users to a private channel without being subscribed
	// to that channel themselves. Now, If a user is part of
	// `can_add_subscribers_group`, they can subscribe themselves or other
	// users to a private channel without being subscribed to that channel.
	//
	// Removed `stream_post_policy` and `is_announcement_only`
	// parameters in Zulip 10.0 (feature level 333), as permission to post
	// in the channel is now controlled by `can_send_message_group`.
	//
	// Before Zulip 8.0 (feature level 208), if a user specified by the
	// [`principals`] parameter was a deactivated user,
	// or did not exist, then an HTTP status code of 403 was returned with
	// `code: "UNAUTHORIZED_PRINCIPAL"` in the error response. As of this
	// feature level, an HTTP status code of 400 is returned with
	// `code: "BAD_REQUEST"` in the error response for these cases.
	//
	// [channel's permissions settings]: https://zulip.com/help/channel-permissions
	// [`principals`]: https://zulip.com/api/subscribe#parameter-principals
	//
	// [channel settings]: https://zulip.com/api/update-stream
	Subscribe(ctx context.Context) SubscribeRequest

	// SubscribeExecute executes the request
	SubscribeExecute(r SubscribeRequest) (*SubscribeResponse, *http.Response, error)

	// Unsubscribe Unsubscribe from a channel
	//
	// Unsubscribe yourself or other users from one or more channels.
	//
	// In addition to managing the current user's subscriptions, this
	// endpoint can be used to remove other users from channels. This
	// is possible in 3 situations:
	//
	//   - Organization administrators can remove any user from any
	// channel.
	//   - Users can remove a bot that they own from any channel that
	// the user [can access].
	//   - Users can unsubscribe any user from a channel if they [have access] to the channel and are a
	// member of the [user group] specified
	// by the [`can_remove_subscribers_group`]
	// for the channel.
	//
	// *Changes**: Before Zulip 10.0 (feature level 362),
	// subscriptions in archived channels could not be modified.
	//
	// Before Zulip 8.0 (feature level 208), if a user specified by
	// the [`principals`] parameter was a
	// deactivated user, or did not exist, then an HTTP status code
	// of 403 was returned with `code: "UNAUTHORIZED_PRINCIPAL"` in
	// the error response. As of this feature level, an HTTP status
	// code of 400 is returned with `code: "BAD_REQUEST"` in the
	// error response for these cases.
	//
	// Before Zulip 8.0 (feature level 197),
	// the `can_remove_subscribers_group` setting
	// was named `can_remove_subscribers_group_id`.
	//
	// Before Zulip 7.0 (feature level 161), the
	// `can_remove_subscribers_group_id` for all channels was always
	// the system group for organization administrators.
	//
	// Before Zulip 6.0 (feature level 145), users had no special
	// privileges for managing bots that they own.
	//
	// [`principals`]: https://zulip.com/api/unsubscribe#parameter-principals
	// [`can_remove_subscribers_group`]: https://zulip.com/api/subscribe#parameter-can_remove_subscribers_group
	//
	// [can access]: https://zulip.com/help/channel-permissions
	// [user group]: https://zulip.com/api/get-user-groups
	// [have access]: https://zulip.com/help/channel-permissions
	Unsubscribe(ctx context.Context) UnsubscribeRequest

	// UnsubscribeExecute executes the request
	UnsubscribeExecute(r UnsubscribeRequest) (*UnsubscribeResponse, *http.Response, error)

	// UpdateChannelFolder Update a channel folder
	//
	// Update the name or description of a channel folder.
	//
	// This endpoint is also used to archive and unarchive
	// a channel folder.
	//
	// Only organization administrators can update a
	// channel folder.
	//
	// *Changes**: New in Zulip 11.0 (feature level 389).
	//
	UpdateChannelFolder(ctx context.Context, channelFolderId int64) UpdateChannelFolderRequest

	// UpdateChannelFolderExecute executes the request
	UpdateChannelFolderExecute(r UpdateChannelFolderRequest) (*zulip.Response, *http.Response, error)

	// UpdateChannel Update a channel
	//
	// Configure the channel with the Id `channelId`. This endpoint supports
	// an organization administrator editing any property of a channel,
	// including:
	//
	//   - Channel [name] and [description]
	//   - Channel [permissions], including
	// [privacy] and [who can send].
	//
	// Note that an organization administrator's ability to change a
	// [private channel's permissions]
	// depends on them being subscribed to the channel.
	//
	// *Changes**: Before Zulip 10.0 (feature level 362), channel privacy could not be
	// edited for archived channels.
	//
	// Removed `stream_post_policy` and `is_announcement_only`
	// parameters in Zulip 10.0 (feature level 333), as permission to post
	// in the channel is now controlled by `can_send_message_group`.
	//
	// [name]: https://zulip.com/help/rename-a-channel
	// [description]: https://zulip.com/help/change-the-channel-description
	// [permissions]: https://zulip.com/help/channel-permissions
	// [privacy]: https://zulip.com/help/change-the-privacy-of-a-channel
	// [private channel's permissions]: https://zulip.com/help/channel-permissions#private-channels
	// [who can send]: https://zulip.com/help/channel-posting-policy
	UpdateChannel(ctx context.Context, channelId int64) UpdateChannelRequest

	// UpdateChannelExecute executes the request
	UpdateChannelExecute(r UpdateChannelRequest) (*zulip.Response, *http.Response, error)

	// UpdateSubscriptionSettings Update subscription settings
	//
	// This endpoint is used to update the user's personal settings for the
	// channels they are subscribed to, including muting, color, pinning, and
	// per-channel notification settings.
	//
	// *Changes**: Prior to Zulip 5.0 (feature level 111), response
	// object included the `subscription_data` in the
	// request. The endpoint now returns the more ergonomic
	// [`ignored_parameters_unsupported`] array instead.
	//
	// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
	//
	UpdateSubscriptionSettings(ctx context.Context) UpdateSubscriptionSettingsRequest

	// UpdateSubscriptionSettingsExecute executes the request
	UpdateSubscriptionSettingsExecute(r UpdateSubscriptionSettingsRequest) (*zulip.Response, *http.Response, error)

	// UpdateSubscriptions Update subscriptions
	//
	// Update which channels you are subscribed to.
	//
	// *Changes**: Before Zulip 10.0 (feature level 362),
	// subscriptions in archived channels could not be modified.
	//
	UpdateSubscriptions(ctx context.Context) UpdateSubscriptionsRequest

	// UpdateSubscriptionsExecute executes the request
	UpdateSubscriptionsExecute(r UpdateSubscriptionsRequest) (*UpdateSubscriptionsResponse, *http.Response, error)

	// UpdateUserTopic Update personal preferences for a topic
	//
	// This endpoint is used to update the personal preferences for a topic,
	// such as the topic's visibility policy, which is used to implement
	// [mute a topic] and related features.
	//
	// This endpoint can be used to update the visibility policy for the single
	// channel and topic pair indicated by the parameters for a user.
	//
	// *Changes**: New in Zulip 7.0 (feature level 170). Previously,
	// toggling whether a topic was muted or unmuted was managed by the
	// [PATCH /users/me/subscriptions/muted_topics] endpoint.
	//
	// [mute a topic]: https://zulip.com/help/mute-a-topic
	// [PATCH /users/me/subscriptions/muted_topics]: https://zulip.com/api/mute-topic
	UpdateUserTopic(ctx context.Context) UpdateUserTopicRequest

	// UpdateUserTopicExecute executes the request
	UpdateUserTopicExecute(r UpdateUserTopicRequest) (*zulip.Response, *http.Response, error)
}

type channelsService struct {
	client clients.Client
}

func NewChannelsService(client clients.Client) *channelsService {
	return &channelsService{client: client}
}

var _ APIChannels = (*channelsService)(nil)

type AddDefaultChannelRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  *int64
}

// The Id of the target channel.
func (r AddDefaultChannelRequest) ChannelId(channelId int64) AddDefaultChannelRequest {
	r.channelId = &channelId
	return r
}

func (r AddDefaultChannelRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.AddDefaultChannelExecute(r)
}

// AddDefaultChannel Add a default channel
//
// Add a channel to the set of [default channels]
// for new users joining the organization.
//
// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
func (s *channelsService) AddDefaultChannel(ctx context.Context) AddDefaultChannelRequest {
	return AddDefaultChannelRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) AddDefaultChannelExecute(r AddDefaultChannelRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/default_streams"
	)
	if r.channelId == nil {
		return nil, nil, fmt.Errorf("channelId is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "stream_id", r.channelId)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type ArchiveChannelRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  int64
}

func (r ArchiveChannelRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.ArchiveChannelExecute(r)
}

// ArchiveChannel Archive a channel
//
// [Archive the channel] with the Id `channelId`.
//
// [Archive the channel]: https://zulip.com/help/archive-a-channel
func (s *channelsService) ArchiveChannel(ctx context.Context, channelId int64) ArchiveChannelRequest {
	return ArchiveChannelRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) ArchiveChannelExecute(r ArchiveChannelRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/streams/{stream_id}"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type CreateBigBlueButtonVideoCallRequest struct {
	ctx         context.Context
	apiService  APIChannels
	meetingName *string
	voiceOnly   *bool
}

// Meeting name for the BigBlueButton video call.
func (r CreateBigBlueButtonVideoCallRequest) MeetingName(meetingName string) CreateBigBlueButtonVideoCallRequest {
	r.meetingName = &meetingName
	return r
}

// Configures whether the call is voice-only; if true, disables cameras for all users. Only the call creator/moderator can edit this configuration.
//
// **Changes**: New in Zulip 10.0 (feature level 337).
func (r CreateBigBlueButtonVideoCallRequest) VoiceOnly(voiceOnly bool) CreateBigBlueButtonVideoCallRequest {
	r.voiceOnly = &voiceOnly
	return r
}

func (r CreateBigBlueButtonVideoCallRequest) Execute() (*CreateBigBlueButtonVideoCallResponse, *http.Response, error) {
	return r.apiService.CreateBigBlueButtonVideoCallExecute(r)
}

// CreateBigBlueButtonVideoCall Create BigBlueButton video call
//
// Create a video call URL for a BigBlueButton video call.
// Requires [BigBlueButton 2.4+]
// to be configured on the Zulip server.
//
// The acting user will be given the moderator role on the call.
//
// *Changes**: Prior to Zulip 10.0 (feature level 337), every
// user was given the moderator role on BigBlueButton calls, via
// encoding a moderator password in the generated URLs.
//
// [BigBlueButton 2.4+]: /integrations/doc/big-blue-button
func (s *channelsService) CreateBigBlueButtonVideoCall(ctx context.Context) CreateBigBlueButtonVideoCallRequest {
	return CreateBigBlueButtonVideoCallRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) CreateBigBlueButtonVideoCallExecute(r CreateBigBlueButtonVideoCallRequest) (*CreateBigBlueButtonVideoCallResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateBigBlueButtonVideoCallResponse{}
		endpoint = "/calls/bigbluebutton/create"
	)
	if r.meetingName == nil {
		return nil, nil, fmt.Errorf("meetingName is required and must be specified")
	}

	AddParam(query, "meeting_name", r.meetingName)
	AddOptionalParam(query, "voice_only", r.voiceOnly)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type CreateChannelRequest struct {
	ctx                               context.Context
	apiService                        APIChannels
	name                              *string
	subscribers                       *[]int64
	description                       *string
	announce                          *bool
	inviteOnly                        *bool
	isWebPublic                       *bool
	isDefaultChannel                  *bool
	folderId                          *int64
	sendNewSubscriptionMessages       *bool
	topicsPolicy                      *zulip.TopicsPolicy
	historyPublicToSubscribers        *bool
	messageRetentionDays              *zulip.MessageRetentionDaysValue
	canAddSubscribersGroup            *zulip.GroupSettingValue
	canDeleteAnyMessageGroup          *zulip.GroupSettingValue
	canDeleteOwnMessageGroup          *zulip.GroupSettingValue
	canRemoveSubscribersGroup         *zulip.GroupSettingValue
	canAdministerChannelGroup         *zulip.GroupSettingValue
	canMoveMessagesOutOfChannelGroup  *zulip.GroupSettingValue
	canMoveMessagesWithinChannelGroup *zulip.GroupSettingValue
	canSendMessageGroup               *zulip.GroupSettingValue
	canSubscribeGroup                 *zulip.GroupSettingValue
	canResolveTopicsGroup             *zulip.GroupSettingValue
}

// The name of the new channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`] endpoint to determine the maximum channel name length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateChannelRequest) Name(name string) CreateChannelRequest {
	r.name = &name
	return r
}

// A list of user Ids of the users to be subscribed to the new channel.
func (r CreateChannelRequest) Subscribers(subscribers ...int64) CreateChannelRequest {
	r.subscribers = &subscribers
	return r
}

// The [description] to use for the new channel being created, in text/markdown format.  Clients should use the `max_stream_description_length` returned by the [`POST /register`] endpoint to determine the maximum channel description length.
//
// [description]: https://zulip.com/help/change-the-channel-description
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateChannelRequest) Description(description string) CreateChannelRequest {
	r.description = &description
	return r
}

// This determines whether [notification bot] will send an announcement about the new channel's creation.
//
// [notification bot]: https://zulip.com/help/configure-automated-notices
func (r CreateChannelRequest) Announce(announce bool) CreateChannelRequest {
	r.announce = &announce
	return r
}

// This parameter and the ones that follow are used to request an initial configuration of the new channel.  This parameter determines whether the newly created channel will be a private channel.
func (r CreateChannelRequest) InviteOnly(inviteOnly bool) CreateChannelRequest {
	r.inviteOnly = &inviteOnly
	return r
}

// This parameter determines whether the newly created channel will be a web-public channel.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED`
//
// [server setting] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current user to have permission under the organization's `can_create_web_public_channel_group` realm setting.
//
// [server setting]: https://zulip.readthedocs.io/en/stable/production/settings.html
func (r CreateChannelRequest) IsWebPublic(isWebPublic bool) CreateChannelRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// This parameter determines whether the newly created channel will be added as a [default channel] for new users joining the organization.
//
// [default channel]: https://zulip.com/help/set-default-channels-for-new-users
func (r CreateChannelRequest) IsDefaultChannel(isDefaultChannel bool) CreateChannelRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

// The Id of the folder to which the channel belongs.  Is `null` if channel does not belong to any folder.
//
// **Changes**: New in Zulip 11.0 (feature level 389).
func (r CreateChannelRequest) FolderId(folderId int64) CreateChannelRequest {
	r.folderId = &folderId
	return r
}

// Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than `max_bulk_new_subscription_messages` (available in the [`POST /register`] response) users were subscribed in this request.
//
// **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateChannelRequest) SendNewSubscriptionMessages(sendNewSubscriptionMessages bool) CreateChannelRequest {
	r.sendNewSubscriptionMessages = &sendNewSubscriptionMessages
	return r
}

func (r CreateChannelRequest) TopicsPolicy(topicsPolicy zulip.TopicsPolicy) CreateChannelRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

// Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels].
//
// [private channels]: https://zulip.com/help/channel-permissions#private-channels
func (r CreateChannelRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) CreateChannelRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

func (r CreateChannelRequest) MessageRetentionDays(messageRetentionDays zulip.MessageRetentionDaysValue) CreateChannelRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

func (r CreateChannelRequest) CanAddSubscribersGroup(canAddSubscribersGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

func (r CreateChannelRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

func (r CreateChannelRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

func (r CreateChannelRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

func (r CreateChannelRequest) CanAdministerChannelGroup(canAdministerChannelGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

func (r CreateChannelRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

func (r CreateChannelRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

func (r CreateChannelRequest) CanSendMessageGroup(canSendMessageGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

func (r CreateChannelRequest) CanSubscribeGroup(canSubscribeGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

func (r CreateChannelRequest) CanResolveTopicsGroup(canResolveTopicsGroup zulip.GroupSettingValue) CreateChannelRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

func (r CreateChannelRequest) Execute() (*CreateChannelResponse, *http.Response, error) {
	return r.apiService.CreateChannelExecute(r)
}

// CreateChannel Create a channel
//
// Create a new [channel] and optionally
// subscribe users to the newly created channel. The initial [channel settings]
// will be determined by the optional parameters, like `invite_only`, detailed below.
//
// *Changes**: New in Zulip 11.0 (feature level 417). Previously, this was only possible via
// the [`POST /api/subscribe`] endpoint, which handled both creation and subscription.
//
// [channel]: https://zulip.com/help/create-channels
// [channel settings]: https://zulip.com/api/update-stream
// [`POST /api/subscribe`]: https://zulip.com/api/subscribe
func (s *channelsService) CreateChannel(ctx context.Context) CreateChannelRequest {
	return CreateChannelRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) CreateChannelExecute(r CreateChannelRequest) (*CreateChannelResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateChannelResponse{}
		endpoint = "/channels/create"
	)
	if r.name == nil {
		return nil, nil, fmt.Errorf("name is required and must be specified")
	}
	if r.subscribers == nil {
		return nil, nil, fmt.Errorf("subscribers is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "name", r.name)
	AddOptionalParam(form, "description", r.description)
	if err := AddOptionalJSONParam(form, "subscribers", *r.subscribers); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "announce", r.announce)
	AddOptionalParam(form, "invite_only", r.inviteOnly)
	AddOptionalParam(form, "is_web_public", r.isWebPublic)
	AddOptionalParam(form, "is_default_stream", r.isDefaultChannel)
	AddOptionalParam(form, "folder_id", r.folderId)
	AddOptionalParam(form, "send_new_subscription_messages", r.sendNewSubscriptionMessages)
	AddOptionalParam(form, "topics_policy", r.topicsPolicy)
	AddOptionalParam(form, "history_public_to_subscribers", r.historyPublicToSubscribers)
	if err := AddOptionalJSONParam(form, "message_retention_days", r.messageRetentionDays); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_add_subscribers_group", r.canAddSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_any_message_group", r.canDeleteAnyMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_own_message_group", r.canDeleteOwnMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_remove_subscribers_group", r.canRemoveSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_administer_channel_group", r.canAdministerChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_out_of_channel_group", r.canMoveMessagesOutOfChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_within_channel_group", r.canMoveMessagesWithinChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_send_message_group", r.canSendMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_subscribe_group", r.canSubscribeGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_resolve_topics_group", r.canResolveTopicsGroup); err != nil {
		return nil, nil, err
	}
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type CreateChannelFolderRequest struct {
	ctx         context.Context
	apiService  APIChannels
	name        *string
	description *string
}

// The name of the channel folder.  Clients should use the `max_channel_folder_name_length` returned by the [`POST /register`] endpoint to determine the maximum channel folder name length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateChannelFolderRequest) Name(name string) CreateChannelFolderRequest {
	r.name = &name
	return r
}

// The description of the channel folder.  Clients should use the `max_channel_folder_description_length` returned by the [`POST /register`] endpoint to determine the maximum channel folder description length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r CreateChannelFolderRequest) Description(description string) CreateChannelFolderRequest {
	r.description = &description
	return r
}

func (r CreateChannelFolderRequest) Execute() (*CreateChannelFolderResponse, *http.Response, error) {
	return r.apiService.CreateChannelFolderExecute(r)
}

// CreateChannelFolder Create a channel folder
//
// Create a new channel folder, that will be used to organize
// channels in left sidebar.
//
// Only organization administrators can create a new channel
// folder.
//
// *Changes**: New in Zulip 11.0 (feature level 389).
func (s *channelsService) CreateChannelFolder(ctx context.Context) CreateChannelFolderRequest {
	return CreateChannelFolderRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) CreateChannelFolderExecute(r CreateChannelFolderRequest) (*CreateChannelFolderResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &CreateChannelFolderResponse{}
		endpoint = "/channel_folders/create"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "name", r.name)
	AddOptionalParam(form, "description", r.description)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type DeleteTopicRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  int64
	topicName  *string
}

// The name of the topic to delete.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.
//
// **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [POST /register]: https://zulip.com/api/register-queue
func (r DeleteTopicRequest) TopicName(topicName string) DeleteTopicRequest {
	r.topicName = &topicName
	return r
}

func (r DeleteTopicRequest) Execute() (*MarkAllAsReadResponse, *http.Response, error) {
	return r.apiService.DeleteTopicExecute(r)
}

// DeleteTopic Delete a topic
//
// Delete all messages in a topic.
//
// Topics are a field on messages (not an independent data structure), so
// deleting all the messages in the topic deletes the topic from Zulip.
//
// Because this endpoint deletes messages in batches, it is possible for
// the request to time out after only deleting some messages in the topic.
// When this happens, the `complete` boolean field in the success response
// will be `false`. Clients should repeat the request when handling such a
// response. If all messages in the topic were deleted, then the success
// response will return `"complete": true`.
//
// *Changes**: Before Zulip 9.0 (feature level 256), the server never sent
// [`stream` op: `update`] events with an
// updated `first_message_id` for a channel when the oldest message that
// had been sent to it changed.
//
// Before Zulip 8.0 (feature level 211), if the server's
// processing was interrupted by a timeout, but some messages in the topic
// were deleted, then it would return `"result": "partially_completed"`,
// along with a `code` field for an error string, in the success response
// to indicate that there was a timeout and that the client should repeat
// the request.
//
// As of Zulip 6.0 (feature level 154), instead of returning an error
// response when a request times out after successfully deleting some of
// the messages in the topic, a success response is returned with
// `"result": "partially_completed"` to indicate that some messages were
// deleted.
//
// Before Zulip 6.0 (feature level 147), this request did a single atomic
// operation, which could time out for very large topics. As of this
// feature level, messages are deleted in batches, starting with the newest
// messages, so that progress is made even if the request times out and
// returns an error.
//
// [`stream` op: `update`]: https://zulip.com/api/get-events#stream-update
func (s *channelsService) DeleteTopic(ctx context.Context, channelId int64) DeleteTopicRequest {
	return DeleteTopicRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) DeleteTopicExecute(r DeleteTopicRequest) (*MarkAllAsReadResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &MarkAllAsReadResponse{}
		endpoint = "/streams/{stream_id}/delete_topic"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	if r.topicName == nil {
		return nil, nil, fmt.Errorf("topicName is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "topic_name", r.topicName)
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelFoldersRequest struct {
	ctx             context.Context
	apiService      APIChannels
	includeArchived *bool
}

// Whether to include archived channel folders in the response.
func (r GetChannelFoldersRequest) IncludeArchived(includeArchived bool) GetChannelFoldersRequest {
	r.includeArchived = &includeArchived
	return r
}

func (r GetChannelFoldersRequest) Execute() (*GetChannelFoldersResponse, *http.Response, error) {
	return r.apiService.GetChannelFoldersExecute(r)
}

// GetChannelFolders Get channel folders
//
// Fetches all of the channel folders in the organization.
// The folders are sorted by the `order` field.
//
// *Changes**: Before Zulip 11.0 (feature level 414),
// these were sorted by Id. (The `order` field didn't exist).
//
// New in Zulip 11.0 (feature level 389).
func (s *channelsService) GetChannelFolders(ctx context.Context) GetChannelFoldersRequest {
	return GetChannelFoldersRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelFoldersExecute(r GetChannelFoldersRequest) (*GetChannelFoldersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelFoldersResponse{}
		endpoint = "/channel_folders"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "include_archived", r.includeArchived)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelByIdRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  int64
}

func (r GetChannelByIdRequest) Execute() (*GetChannelResponse, *http.Response, error) {
	return r.apiService.GetChannelByIdExecute(r)
}

// GetChannelById Get a channel by Id
//
// Fetch details for the channel with the Id `channelId`.
//
// *Changes**: New in Zulip 6.0 (feature level 132).
func (s *channelsService) GetChannelById(ctx context.Context, channelId int64) GetChannelByIdRequest {
	return GetChannelByIdRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelByIdExecute(r GetChannelByIdRequest) (*GetChannelResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelResponse{}
		endpoint = "/streams/{stream_id}"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelEmailAddressRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  int64
	senderId   *int64
}

// The Id of a user or bot which should appear as the sender when messages are sent to the channel using the returned channel email address.  `sender_id` can be:  - Id of the current user. - Id of the Email gateway bot. (Default value) - Id of a bot owned by the current user.
//
// **Changes**: New in Zulip 10.0 (feature level 335).  Previously, the sender was always Email gateway bot.
func (r GetChannelEmailAddressRequest) SenderId(senderId int64) GetChannelEmailAddressRequest {
	r.senderId = &senderId
	return r
}

func (r GetChannelEmailAddressRequest) Execute() (*GetChannelEmailAddressResponse, *http.Response, error) {
	return r.apiService.GetChannelEmailAddressExecute(r)
}

// GetChannelEmailAddress Get channel's email address
//
// Get email address of a channel.
//
// *Changes**: New in Zulip 8.0 (feature level 226).
func (s *channelsService) GetChannelEmailAddress(ctx context.Context, channelId int64) GetChannelEmailAddressRequest {
	return GetChannelEmailAddressRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelEmailAddressExecute(r GetChannelEmailAddressRequest) (*GetChannelEmailAddressResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelEmailAddressResponse{}
		endpoint = "/streams/{stream_id}/email_address"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	AddOptionalParam(query, "sender_id", r.senderId)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelIdRequest struct {
	ctx        context.Context
	apiService APIChannels
	channel    *string
}

// The name of the channel to access.
func (r GetChannelIdRequest) Channel(channel string) GetChannelIdRequest {
	r.channel = &channel
	return r
}

func (r GetChannelIdRequest) Execute() (*GetChannelIdResponse, *http.Response, error) {
	return r.apiService.GetChannelIdExecute(r)
}

// GetChannelId Get channel Id
//
// Get the unique Id of a given channel.
func (s *channelsService) GetChannelId(ctx context.Context) GetChannelIdRequest {
	return GetChannelIdRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelIdExecute(r GetChannelIdRequest) (*GetChannelIdResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelIdResponse{}
		endpoint = "/get_stream_id"
	)
	if r.channel == nil {
		return nil, nil, fmt.Errorf("channel is required and must be specified")
	}

	AddParam(query, "stream", r.channel)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelTopicsRequest struct {
	ctx                 context.Context
	apiService          APIChannels
	channelId           int64
	allowEmptyTopicName *bool
}

// Whether the client supports processing the empty string as a topic name in the returned data.  If `false`, the value of `realm_empty_topic_display_name` found in the [`POST /register`] response is returned replacing the empty string as the topic name.
//
// **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r GetChannelTopicsRequest) AllowEmptyTopicName(allowEmptyTopicName bool) GetChannelTopicsRequest {
	r.allowEmptyTopicName = &allowEmptyTopicName
	return r
}

func (r GetChannelTopicsRequest) Execute() (*GetChannelTopicsResponse, *http.Response, error) {
	return r.apiService.GetChannelTopicsExecute(r)
}

// GetChannelTopics Get topics in a channel
//
// Get all topics the user has access to in a specific channel.
//
// Note that for [private channels with protected history],
// the user will only have access to topics of messages sent after they
// [subscribed to] the channel. Similarly, a user's
// [bot] will only have access to messages
// sent after the bot was subscribed to the channel, instead of when the
// user subscribed.
//
// [subscribed to]: https://zulip.com/api/subscribe
// [bot]: https://zulip.com/help/bots-overview#bot-type
// [private channels with protected history]: https://zulip.com/help/channel-permissions#private-channels
func (s *channelsService) GetChannelTopics(ctx context.Context, channelId int64) GetChannelTopicsRequest {
	return GetChannelTopicsRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelTopicsExecute(r GetChannelTopicsRequest) (*GetChannelTopicsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelTopicsResponse{}
		endpoint = "/users/me/{stream_id}/topics"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	AddOptionalParam(query, "allow_empty_topic_name", r.allowEmptyTopicName)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetChannelsRequest struct {
	ctx                     context.Context
	apiService              APIChannels
	includePublic           *bool
	includeWebPublic        *bool
	includeSubscribed       *bool
	excludeArchived         *bool
	includeAllActive        *bool
	includeAll              *bool
	includeDefault          *bool
	includeOwnerSubscribed  *bool
	includeCanAccessContent *bool
}

// Include all public channels.
func (r GetChannelsRequest) IncludePublic(includePublic bool) GetChannelsRequest {
	r.includePublic = &includePublic
	return r
}

// Include all web-public channels.
func (r GetChannelsRequest) IncludeWebPublic(includeWebPublic bool) GetChannelsRequest {
	r.includeWebPublic = &includeWebPublic
	return r
}

// Include all channels that the user is subscribed to.
func (r GetChannelsRequest) IncludeSubscribed(includeSubscribed bool) GetChannelsRequest {
	r.includeSubscribed = &includeSubscribed
	return r
}

// Whether to exclude archived streams from the results.
//
// **Changes**: New in Zulip 10.0 (feature level 315).
func (r GetChannelsRequest) ExcludeArchived(excludeArchived bool) GetChannelsRequest {
	r.excludeArchived = &excludeArchived
	return r
}

// Deprecated parameter to include all channels. The user must have administrative privileges to use this parameter.
//
// **Changes**: Deprecated in Zulip 10.0 (feature level 356). Clients interacting with newer servers should use the equivalent `include_all` parameter, which does not incorrectly hint that this parameter, and not `exclude_archived`, controls whether archived channels appear in the response.
//
// Deprecated
func (r GetChannelsRequest) IncludeAllActive(includeAllActive bool) GetChannelsRequest {
	r.includeAllActive = &includeAllActive
	return r
}

// Include all channels that the user has metadata access to.  For organization administrators, this will be all channels in the organization, since organization administrators implicitly have metadata access to all channels.
//
// **Changes**: New in Zulip 10.0 (feature level 356). On older versions, use `include_all_active`, which this replaces.
func (r GetChannelsRequest) IncludeAll(includeAll bool) GetChannelsRequest {
	r.includeAll = &includeAll
	return r
}

// Include all default channels for the user's realm.
func (r GetChannelsRequest) IncludeDefault(includeDefault bool) GetChannelsRequest {
	r.includeDefault = &includeDefault
	return r
}

// If the user is a bot, include all channels that the bot's owner is subscribed to.
func (r GetChannelsRequest) IncludeOwnerSubscribed(includeOwnerSubscribed bool) GetChannelsRequest {
	r.includeOwnerSubscribed = &includeOwnerSubscribed
	return r
}

// Include all the channels that the user has content access to.
//
// **Changes**: New in Zulip 10.0 (feature level 356).
func (r GetChannelsRequest) IncludeCanAccessContent(includeCanAccessContent bool) GetChannelsRequest {
	r.includeCanAccessContent = &includeCanAccessContent
	return r
}

func (r GetChannelsRequest) Execute() (*GetChannelsResponse, *http.Response, error) {
	return r.apiService.GetChannelsExecute(r)
}

// GetChannels Get all channels
//
// Get all channels that the user [has access to].
//
// [has access to]: https://zulip.com/help/channel-permissions
func (s *channelsService) GetChannels(ctx context.Context) GetChannelsRequest {
	return GetChannelsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) GetChannelsExecute(r GetChannelsRequest) (*GetChannelsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetChannelsResponse{}
		endpoint = "/streams"
	)
	AddOptionalParam(query, "include_public", r.includePublic)
	AddOptionalParam(query, "include_web_public", r.includeWebPublic)
	AddOptionalParam(query, "include_subscribed", r.includeSubscribed)
	AddOptionalParam(query, "exclude_archived", r.excludeArchived)
	AddOptionalParam(query, "include_all_active", r.includeAllActive)
	AddOptionalParam(query, "include_all", r.includeAll)
	AddOptionalParam(query, "include_default", r.includeDefault)
	AddOptionalParam(query, "include_owner_subscribed", r.includeOwnerSubscribed)
	AddOptionalParam(query, "include_can_access_content", r.includeCanAccessContent)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetSubscribersRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  int64
}

func (r GetSubscribersRequest) Execute() (*GetSubscribersResponse, *http.Response, error) {
	return r.apiService.GetSubscribersExecute(r)
}

// GetSubscribers Get channel subscribers
//
// Get all users subscribed to a channel.
func (s *channelsService) GetSubscribers(ctx context.Context, channelId int64) GetSubscribersRequest {
	return GetSubscribersRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) GetSubscribersExecute(r GetSubscribersRequest) (*GetSubscribersResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetSubscribersResponse{}
		endpoint = "/streams/{stream_id}/members"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetSubscriptionStatusRequest struct {
	ctx        context.Context
	apiService APIChannels
	userId     int64
	channelId  int64
}

func (r GetSubscriptionStatusRequest) Execute() (*GetSubscriptionStatusResponse, *http.Response, error) {
	return r.apiService.GetSubscriptionStatusExecute(r)
}

// GetSubscriptionStatus Get subscription status
//
// Check whether a user is subscribed to a channel.
//
// *Changes**: New in Zulip 3.0 (feature level 12).
func (s *channelsService) GetSubscriptionStatus(ctx context.Context, userId int64, channelId int64) GetSubscriptionStatusRequest {
	return GetSubscriptionStatusRequest{
		apiService: s,
		ctx:        ctx,
		userId:     userId,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) GetSubscriptionStatusExecute(r GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetSubscriptionStatusResponse{}
		endpoint = "/users/{user_id}/subscriptions/{stream_id}"
	)

	path := strings.Replace(endpoint, "{user_id}", IdToString(r.userId), -1)
	path = strings.Replace(path, "{stream_id}", IdToString(r.channelId), -1)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type GetSubscriptionsRequest struct {
	ctx                context.Context
	apiService         APIChannels
	includeSubscribers *string
}

// Whether each returned channel object should include a `subscribers` field containing a list of the user Ids of its subscribers.  Client apps supporting organizations with many thousands of users should not pass `true`, because the full subscriber matrix may be several megabytes of data. The `partial` value, combined with the `subscriber_count` and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes `partial` for this parameter, the server may, for some channels, return a subset of the channel's subscribers in the `partial_subscribers` field instead of the `subscribers` field, which always contains the complete set of subscribers.  The server guarantees that it will always return a `subscribers` field for channels with fewer than 250 total subscribers. When returning a `partial_subscribers` field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.
//
// **Changes**: The `partial` value is new in Zulip 11.0 (feature level 412).  New in Zulip 2.1.0.
func (r GetSubscriptionsRequest) IncludeSubscribers(includeSubscribers string) GetSubscriptionsRequest {
	r.includeSubscribers = &includeSubscribers
	return r
}

func (r GetSubscriptionsRequest) Execute() (*GetSubscriptionsResponse, *http.Response, error) {
	return r.apiService.GetSubscriptionsExecute(r)
}

// GetSubscriptions Get subscribed channels
//
// Get all channels that the user is subscribed to.
func (s *channelsService) GetSubscriptions(ctx context.Context) GetSubscriptionsRequest {
	return GetSubscriptionsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) GetSubscriptionsExecute(r GetSubscriptionsRequest) (*GetSubscriptionsResponse, *http.Response, error) {
	var (
		method   = http.MethodGet
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &GetSubscriptionsResponse{}
		endpoint = "/users/me/subscriptions"
	)
	AddOptionalParam(query, "include_subscribers", r.includeSubscribers)

	headers["Accept"] = "application/json"
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type MuteTopicRequest struct {
	ctx        context.Context
	apiService APIChannels
	topic      *string
	op         *string
	channelId  *int64
	channel    *string
}

// The topic to (un)mute. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r MuteTopicRequest) Topic(topic string) MuteTopicRequest {
	r.topic = &topic
	return r
}

// Whether to mute (`add`) or unmute (`remove`) the provided topic.
func (r MuteTopicRequest) Op(op string) MuteTopicRequest {
	r.op = &op
	return r
}

// The Id of the channel to access.  Clients must provide either `stream` or `stream_id` as a parameter to this endpoint, but not both.
//
// **Changes**: New in Zulip 2.0.0.
func (r MuteTopicRequest) ChannelId(channelId int64) MuteTopicRequest {
	r.channelId = &channelId
	return r
}

// The name of the channel to access.  Clients must provide either `stream` or `stream_id` as a parameter to this endpoint, but not both. Clients should use `stream_id` instead of the `stream` parameter when possible.
func (r MuteTopicRequest) Channel(channel string) MuteTopicRequest {
	r.channel = &channel
	return r
}

func (r MuteTopicRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.MuteTopicExecute(r)
}

// MuteTopic Topic muting
//
// [Mute or unmute a topic] within a channel that
// the current user is subscribed to.
//
// *Changes**: Deprecated in Zulip 7.0 (feature level 170). Clients connecting
// to newer servers should use the [POST /user_topics]
// endpoint, as this endpoint may be removed in a future release.
//
// Before Zulip 7.0 (feature level 169), this endpoint
// returned an error if asked to mute a topic that was already muted
// or asked to unmute a topic that had not previously been muted.
//
// # Deprecated
//
// [Mute or unmute a topic]: https://zulip.com/help/mute-a-topic
// [POST /user_topics]: https://zulip.com/api/update-user-topic
func (s *channelsService) MuteTopic(ctx context.Context) MuteTopicRequest {
	return MuteTopicRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
//
// Deprecated
func (s *channelsService) MuteTopicExecute(r MuteTopicRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/subscriptions/muted_topics"
	)
	if r.topic == nil {
		return nil, nil, fmt.Errorf("topic is required and must be specified")
	}
	if r.op == nil {
		return nil, nil, fmt.Errorf("op is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "stream_id", r.channelId)
	AddOptionalParam(form, "stream", r.channel)
	AddParam(form, "topic", r.topic)
	AddParam(form, "op", r.op)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type PatchChannelFoldersRequest struct {
	ctx        context.Context
	apiService APIChannels
	order      *[]int64
}

// A list of channel folder Ids representing the new order.  This list must include the Ids of all the organization's channel folders, including archived folders.
func (r PatchChannelFoldersRequest) Order(order []int64) PatchChannelFoldersRequest {
	r.order = &order
	return r
}

func (r PatchChannelFoldersRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.PatchChannelFoldersExecute(r)
}

// PatchChannelFolders Reorder channel folders
//
// Given an array of channel folder Ids, this method will set the `order`
// property of all of the channel folders in the organization according to
// the order of the channel folder Ids specified in the request.
//
// *Changes**: New in Zulip 11.0 (feature level 414).
func (s *channelsService) PatchChannelFolders(ctx context.Context) PatchChannelFoldersRequest {
	return PatchChannelFoldersRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) PatchChannelFoldersExecute(r PatchChannelFoldersRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/channel_folders"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "order", r.order)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type RemoveDefaultChannelRequest struct {
	ctx        context.Context
	apiService APIChannels
	channelId  *int64
}

// The Id of the target channel.
func (r RemoveDefaultChannelRequest) ChannelId(channelId int64) RemoveDefaultChannelRequest {
	r.channelId = &channelId
	return r
}

func (r RemoveDefaultChannelRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.RemoveDefaultChannelExecute(r)
}

// RemoveDefaultChannel Remove a default channel
//
// Remove a channel from the set of [default channels]
// for new users joining the organization.
//
// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
func (s *channelsService) RemoveDefaultChannel(ctx context.Context) RemoveDefaultChannelRequest {
	return RemoveDefaultChannelRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) RemoveDefaultChannelExecute(r RemoveDefaultChannelRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/default_streams"
	)
	if r.channelId == nil {
		return nil, nil, fmt.Errorf("channelId is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "stream_id", r.channelId)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type SubscribeRequest struct {
	ctx                               context.Context
	apiService                        APIChannels
	subscriptions                     *[]SubscriptionRequest
	principals                        *zulip.Principals
	authorizationErrorsFatal          *bool
	announce                          *bool
	inviteOnly                        *bool
	isWebPublic                       *bool
	isDefaultChannel                  *bool
	folderId                          *int64
	sendNewSubscriptionMessages       *bool
	topicsPolicy                      *zulip.TopicsPolicy
	historyPublicToSubscribers        *bool
	messageRetentionDays              *zulip.MessageRetentionDaysValue
	canAddSubscribersGroup            *zulip.GroupSettingValue
	canDeleteAnyMessageGroup          *zulip.GroupSettingValue
	canDeleteOwnMessageGroup          *zulip.GroupSettingValue
	canRemoveSubscribersGroup         *zulip.GroupSettingValue
	canAdministerChannelGroup         *zulip.GroupSettingValue
	canMoveMessagesOutOfChannelGroup  *zulip.GroupSettingValue
	canMoveMessagesWithinChannelGroup *zulip.GroupSettingValue
	canSendMessageGroup               *zulip.GroupSettingValue
	canSubscribeGroup                 *zulip.GroupSettingValue
	canResolveTopicsGroup             *zulip.GroupSettingValue
}

// SubscriptionRequest struct for SubscriptionRequest
type SubscriptionRequest struct {
	// The name of the channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`] endpoint to determine the maximum channel name length.
	//
	// [`POST /register`]: https://zulip.com/api/register-queue
	Name string `json:"name"`
	// The [description] to use for a new channel being created, in text/markdown format.  See the help center article on [message formatting] for details on Zulip-flavored Markdown.  Clients should use the `max_stream_description_length` returned by the [`POST /register`] endpoint to determine the maximum channel description length.
	//
	// [description]: https://zulip.com/help/change-the-channel-description
	// [message formatting]: https://zulip.com/help/format-your-message-using-markdown
	// [`POST /register`]: https://zulip.com/api/register-queue
	Description *string `json:"description,omitempty"`
}

// A list of dictionaries containing the key `name` and value specifying the name of the channel to subscribe. If the channel does not exist a new channel is created. The description of the channel created can be specified by setting the dictionary key `description` with an appropriate value.
func (r SubscribeRequest) Subscriptions(subscriptions []SubscriptionRequest) SubscribeRequest {
	r.subscriptions = &subscriptions
	return r
}

func (r SubscribeRequest) Principals(principals zulip.Principals) SubscribeRequest {
	r.principals = &principals
	return r
}

// A boolean specifying whether authorization errors (such as when the requesting user is not authorized to access a private channel) should be considered fatal or not. When `true`, an authorization error is reported as such. When set to `false`, the response will be a 200 and any channels where the request encountered an authorization error will be listed in the `unauthorized` key.
func (r SubscribeRequest) AuthorizationErrorsFatal(authorizationErrorsFatal bool) SubscribeRequest {
	r.authorizationErrorsFatal = &authorizationErrorsFatal
	return r
}

// If one of the channels specified did not exist previously and is thus created by this call, this determines whether [notification bot] will send an announcement about the new channel's creation.
//
// [notification bot]: https://zulip.com/help/configure-automated-notices
func (r SubscribeRequest) Announce(announce bool) SubscribeRequest {
	r.announce = &announce
	return r
}

// As described above, this endpoint will create a new channel if passed a channel name that doesn't already exist. This parameters and the ones that follow are used to request an initial configuration of a created channel; they are ignored for channels that already exist.  This parameter determines whether any newly created channels will be private channels.
func (r SubscribeRequest) InviteOnly(inviteOnly bool) SubscribeRequest {
	r.inviteOnly = &inviteOnly
	return r
}

// This parameter determines whether any newly created channels will be web-public channels.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED`
//
// [server setting] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current use to have permission under the organization's `can_create_web_public_channel_group` realm setting.
//
// **Changes**: New in Zulip 5.0 (feature level 98).
//
// [server setting]: https://zulip.readthedocs.io/en/stable/production/settings.html
func (r SubscribeRequest) IsWebPublic(isWebPublic bool) SubscribeRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// This parameter determines whether any newly created channels will be added as [default channels] for new users joining the organization.
//
// **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint].
//
// [default channels]: https://zulip.com/help/set-default-channels-for-new-users
// [dedicated API endpoint]: https://zulip.com/api/add-default-stream
func (r SubscribeRequest) IsDefaultChannel(isDefaultChannel bool) SubscribeRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

// Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels].
//
// [private channels]: https://zulip.com/help/channel-permissions#private-channels
func (r SubscribeRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) SubscribeRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

func (r SubscribeRequest) MessageRetentionDays(messageRetentionDays zulip.MessageRetentionDaysValue) SubscribeRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

func (r SubscribeRequest) TopicsPolicy(topicsPolicy zulip.TopicsPolicy) SubscribeRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

func (r SubscribeRequest) CanAddSubscribersGroup(canAddSubscribersGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

func (r SubscribeRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

func (r SubscribeRequest) CanAdministerChannelGroup(canAdministerChannelGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

func (r SubscribeRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

func (r SubscribeRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

func (r SubscribeRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

func (r SubscribeRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

func (r SubscribeRequest) CanSendMessageGroup(canSendMessageGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

func (r SubscribeRequest) CanSubscribeGroup(canSubscribeGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

func (r SubscribeRequest) CanResolveTopicsGroup(canResolveTopicsGroup zulip.GroupSettingValue) SubscribeRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

// This parameter determines the folder to which the newly created channel will be added.  If the value is `None`, the channel will not be added to any folder.
//
// **Changes**: New in Zulip 11.0 (feature level 389).
func (r SubscribeRequest) FolderId(folderId int64) SubscribeRequest {
	r.folderId = &folderId
	return r
}

// Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than `max_bulk_new_subscription_messages` (available in the [`POST /register`] response) users were subscribed in this request.
//
// **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r SubscribeRequest) SendNewSubscriptionMessages(sendNewSubscriptionMessages bool) SubscribeRequest {
	r.sendNewSubscriptionMessages = &sendNewSubscriptionMessages
	return r
}

func (r SubscribeRequest) Execute() (*SubscribeResponse, *http.Response, error) {
	return r.apiService.SubscribeExecute(r)
}

// Subscribe Subscribe to a channel
//
// Subscribe one or more users to one or more channels.
//
// If any of the specified channels do not exist, they are automatically
// created. The initial [channel settings] will be determined
// by the optional parameters, like `invite_only`, detailed below.
//
// Note that the ability to subscribe oneself and/or other users
// to a specified channel depends on the [channel's permissions settings].
//
// *Changes**: Before Zulip 10.0 (feature level 362),
// subscriptions in archived channels could not be modified.
//
// Before Zulip 10.0 (feature level 357), the
// `can_subscribe_group` permission, which allows members of the
// group to subscribe themselves to the channel, did not exist.
//
// Before Zulip 10.0 (feature level 349), a user cannot subscribe
// other users to a private channel without being subscribed
// to that channel themselves. Now, If a user is part of
// `can_add_subscribers_group`, they can subscribe themselves or other
// users to a private channel without being subscribed to that channel.
//
// Removed `stream_post_policy` and `is_announcement_only`
// parameters in Zulip 10.0 (feature level 333), as permission to post
// in the channel is now controlled by `can_send_message_group`.
//
// Before Zulip 8.0 (feature level 208), if a user specified by the
// [`principals`] parameter was a deactivated user,
// or did not exist, then an HTTP status code of 403 was returned with
// `code: "UNAUTHORIZED_PRINCIPAL"` in the error response. As of this
// feature level, an HTTP status code of 400 is returned with
// `code: "BAD_REQUEST"` in the error response for these cases.
//
// [channel's permissions settings]: https://zulip.com/help/channel-permissions
// [`principals`]: https://zulip.com/api/subscribe#parameter-principals
// [channel settings]: https://zulip.com/api/update-stream
func (s *channelsService) Subscribe(ctx context.Context) SubscribeRequest {
	return SubscribeRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) SubscribeExecute(r SubscribeRequest) (*SubscribeResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &SubscribeResponse{}
		endpoint = "/users/me/subscriptions"
	)
	if r.subscriptions == nil {
		return nil, nil, fmt.Errorf("subscriptions is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "subscriptions", r.subscriptions)
	if err := AddOptionalJSONParam(form, "principals", r.principals); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "authorization_errors_fatal", r.authorizationErrorsFatal)
	AddOptionalParam(form, "announce", r.announce)
	AddOptionalParam(form, "invite_only", r.inviteOnly)
	AddOptionalParam(form, "is_web_public", r.isWebPublic)
	AddOptionalParam(form, "is_default_stream", r.isDefaultChannel)
	AddOptionalParam(form, "history_public_to_subscribers", r.historyPublicToSubscribers)
	if err := AddOptionalJSONParam(form, "message_retention_days", r.messageRetentionDays); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "topics_policy", r.topicsPolicy)
	if err := AddOptionalJSONParam(form, "can_add_subscribers_group", r.canAddSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_remove_subscribers_group", r.canRemoveSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_administer_channel_group", r.canAdministerChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_any_message_group", r.canDeleteAnyMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_own_message_group", r.canDeleteOwnMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_out_of_channel_group", r.canMoveMessagesOutOfChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_within_channel_group", r.canMoveMessagesWithinChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_send_message_group", r.canSendMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_subscribe_group", r.canSubscribeGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_resolve_topics_group", r.canResolveTopicsGroup); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "folder_id", r.folderId)
	AddOptionalParam(form, "send_new_subscription_messages", r.sendNewSubscriptionMessages)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UnsubscribeRequest struct {
	ctx           context.Context
	apiService    APIChannels
	subscriptions *[]string
	principals    *zulip.Principals
}

// A list of channel names to unsubscribe from. This parameter is called `streams` in our Python API.
func (r UnsubscribeRequest) Subscriptions(subscriptions []string) UnsubscribeRequest {
	r.subscriptions = &subscriptions
	return r
}

func (r UnsubscribeRequest) Principals(principals zulip.Principals) UnsubscribeRequest {
	r.principals = &principals
	return r
}

func (r UnsubscribeRequest) Execute() (*UnsubscribeResponse, *http.Response, error) {
	return r.apiService.UnsubscribeExecute(r)
}

// Unsubscribe Unsubscribe from a channel
//
// Unsubscribe yourself or other users from one or more channels.
//
// In addition to managing the current user's subscriptions, this
// endpoint can be used to remove other users from channels. This
// is possible in 3 situations:
//
//   - Organization administrators can remove any user from any channel.
//   - Users can remove a bot that they own from any channel that the user [can access].
//   - Users can unsubscribe any user from a channel if they [have access] to the channel and are a member of the [user group] specified by the [`can_remove_subscribers_group`] for the channel.
//
// *Changes**: Before Zulip 10.0 (feature level 362),
// subscriptions in archived channels could not be modified.
//
// Before Zulip 8.0 (feature level 208), if a user specified by
// the [`principals`] parameter was a
// deactivated user, or did not exist, then an HTTP status code
// of 403 was returned with `code: "UNAUTHORIZED_PRINCIPAL"` in
// the error response. As of this feature level, an HTTP status
// code of 400 is returned with `code: "BAD_REQUEST"` in the
// error response for these cases.
//
// Before Zulip 8.0 (feature level 197),
// the `can_remove_subscribers_group` setting
// was named `can_remove_subscribers_group_id`.
//
// Before Zulip 7.0 (feature level 161), the
// `can_remove_subscribers_group_id` for all channels was always
// the system group for organization administrators.
//
// Before Zulip 6.0 (feature level 145), users had no special
// privileges for managing bots that they own.
//
// [`can_remove_subscribers_group`]: https://zulip.com/api/subscribe#parameter-can_remove_subscribers_group
// [`principals`]: https://zulip.com/api/unsubscribe#parameter-principals
// [can access]: https://zulip.com/help/channel-permissions
// [user group]: https://zulip.com/api/get-user-groups
// [have access]: https://zulip.com/help/channel-permissions
func (s *channelsService) Unsubscribe(ctx context.Context) UnsubscribeRequest {
	return UnsubscribeRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) UnsubscribeExecute(r UnsubscribeRequest) (*UnsubscribeResponse, *http.Response, error) {
	var (
		method   = http.MethodDelete
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UnsubscribeResponse{}
		endpoint = "/users/me/subscriptions"
	)
	if r.subscriptions == nil {
		return nil, nil, fmt.Errorf("subscriptions is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "subscriptions", r.subscriptions)
	if err := AddOptionalJSONParam(form, "principals", r.principals); err != nil {
		return nil, nil, err
	}
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UpdateChannelFolderRequest struct {
	ctx             context.Context
	apiService      APIChannels
	channelFolderId int64
	name            *string
	description     *string
	isArchived      *bool
}

// The new name of the channel folder.  Clients should use the `max_channel_folder_name_length` returned by the [`POST /register`] endpoint to determine the maximum channel folder name length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateChannelFolderRequest) Name(name string) UpdateChannelFolderRequest {
	r.name = &name
	return r
}

// The new description of the channel folder.  Clients should use the `max_channel_folder_description_length` returned by the [`POST /register`] endpoint to determine the maximum channel folder description length.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateChannelFolderRequest) Description(description string) UpdateChannelFolderRequest {
	r.description = &description
	return r
}

// Whether to archive or unarchive the channel folder.
func (r UpdateChannelFolderRequest) IsArchived(isArchived bool) UpdateChannelFolderRequest {
	r.isArchived = &isArchived
	return r
}

func (r UpdateChannelFolderRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateChannelFolderExecute(r)
}

// UpdateChannelFolder Update a channel folder
//
// Update the name or description of a channel folder.
//
// This endpoint is also used to archive and unarchive
// a channel folder.
//
// Only organization administrators can update a
// channel folder.
//
// *Changes**: New in Zulip 11.0 (feature level 389).
func (s *channelsService) UpdateChannelFolder(ctx context.Context, channelFolderId int64) UpdateChannelFolderRequest {
	return UpdateChannelFolderRequest{
		apiService:      s,
		ctx:             ctx,
		channelFolderId: channelFolderId,
	}
}

// Execute executes the request
func (s *channelsService) UpdateChannelFolderExecute(r UpdateChannelFolderRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/channel_folders/{channel_folder_id}"
	)

	path := strings.Replace(endpoint, "{channel_folder_id}", IdToString(r.channelFolderId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "name", r.name)
	AddOptionalParam(form, "description", r.description)
	AddOptionalParam(form, "is_archived", r.isArchived)
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UpdateChannelRequest struct {
	ctx                               context.Context
	apiService                        APIChannels
	channelId                         int64
	description                       *string
	newName                           *string
	isPrivate                         *bool
	isWebPublic                       *bool
	historyPublicToSubscribers        *bool
	isDefaultChannel                  *bool
	messageRetentionDays              *zulip.MessageRetentionDaysValue
	isArchived                        *bool
	folderId                          *int64
	topicsPolicy                      *zulip.TopicsPolicy
	canAddSubscribersGroup            *zulip.GroupSettingValueUpdate
	canRemoveSubscribersGroup         *zulip.GroupSettingValueUpdate
	canAdministerChannelGroup         *zulip.GroupSettingValueUpdate
	canDeleteAnyMessageGroup          *zulip.GroupSettingValueUpdate
	canDeleteOwnMessageGroup          *zulip.GroupSettingValueUpdate
	canMoveMessagesOutOfChannelGroup  *zulip.GroupSettingValueUpdate
	canMoveMessagesWithinChannelGroup *zulip.GroupSettingValueUpdate
	canSendMessageGroup               *zulip.GroupSettingValueUpdate
	canSubscribeGroup                 *zulip.GroupSettingValueUpdate
	canResolveTopicsGroup             *zulip.GroupSettingValueUpdate
}

// The new [description] for the channel, in [Zulip-flavored Markdown] format.  Clients should use the `max_stream_description_length` returned by the [`POST /register`] endpoint to determine the maximum channel description length.
//
// **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).
//
// [description]: https://zulip.com/help/change-the-channel-description
// [Zulip-flavored Markdown]: https://zulip.com/help/format-your-message-using-markdown
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateChannelRequest) Description(description string) UpdateChannelRequest {
	r.description = &description
	return r
}

// The new name for the channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`] endpoint to determine the maximum channel name length.
//
// **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateChannelRequest) NewName(newName string) UpdateChannelRequest {
	r.newName = &newName
	return r
}

// Change whether the channel is a private channel.
func (r UpdateChannelRequest) IsPrivate(isPrivate bool) UpdateChannelRequest {
	r.isPrivate = &isPrivate
	return r
}

// Change whether the channel is a web-public channel.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED` [server setting] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current use to have permission under the organization's `can_create_web_public_channel_group` realm setting.
//
// **Changes**: New in Zulip 5.0 (feature level 98).
//
// [server setting]: https://zulip.readthedocs.io/en/stable/production/settings.html
func (r UpdateChannelRequest) IsWebPublic(isWebPublic bool) UpdateChannelRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels].  It's an error for this parameter to be false for a public or web-public channel and when is_private is false.
//
// **Changes**: Before Zulip 6.0 (feature level 136), `history_public_to_subscribers` was silently ignored unless the request also contained either `is_private` or `is_web_public`.
//
// [private channels]: https://zulip.com/help/channel-permissions#private-channels
func (r UpdateChannelRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) UpdateChannelRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

// Add or remove the channel as a [default channel] for new users joining the organization.
//
// **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint].
//
// [default channel]: https://zulip.com/help/set-default-channels-for-new-users
// [dedicated API endpoint]: https://zulip.com/api/add-default-stream
func (r UpdateChannelRequest) IsDefaultChannel(isDefaultChannel bool) UpdateChannelRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

func (r UpdateChannelRequest) MessageRetentionDays(messageRetentionDays zulip.MessageRetentionDaysValue) UpdateChannelRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

// A boolean indicating whether the channel is [archived] or unarchived. Currently only allows unarchiving previously archived channels.
//
// **Changes**: New in Zulip 11.0 (feature level 388).
//
// [archived]: https://zulip.com/help/archive-a-channel
func (r UpdateChannelRequest) IsArchived(isArchived bool) UpdateChannelRequest {
	r.isArchived = &isArchived
	return r
}

// Id of the new folder to which the channel should belong.  It can be `None` if the user wants to just remove the channel from its existing folder.
//
// **Changes**: New in Zulip 11.0 (feature level 389).
func (r UpdateChannelRequest) FolderId(folderId int64) UpdateChannelRequest {
	r.folderId = &folderId
	return r
}

func (r UpdateChannelRequest) TopicsPolicy(topicsPolicy zulip.TopicsPolicy) UpdateChannelRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

// The set of users who have permission to add subscribers to this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 342). Previously, there was no channel-level setting for this permission.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Users who can administer the channel or have similar realm-level permissions can add subscribers to a public channel regardless of the value of this setting.  Users in this group need not be subscribed to a private channel to add subscribers to it.  Note that a user must [have content access] to a channel and permission to administer the channel in order to modify this setting.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanAddSubscribersGroup(canAddSubscribersGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

// The set of users who have permission to unsubscribe others from this channel expressed as an [update to a group-setting value].
//
// **Changes**: Prior to Zulip 10.0 (feature level 349), channel administrators could not unsubscribe other users if they were not an organization administrator or part of `can_remove_subscribers_group`. Realm administrators were not allowed to unsubscribe other users from a private channel if they were not subscribed to that channel.  Prior to Zulip 10.0 (feature level 320), this value was always the integer Id of a system group.  Before Zulip 8.0 (feature level 197), the `can_remove_subscribers_group` setting was named `can_remove_subscribers_group_id`.  New in Zulip 7.0 (feature level 161).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Organization administrators can unsubscribe others from a channel as though they were in this group without being explicitly listed here.  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.
func (r UpdateChannelRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

// The set of users who have permission to administer this channel expressed as an [update to a group-setting value].
//
// **Changes**: Prior to Zulip 10.0 (feature level 349) a user needed to [have content access] to a channel in order to modify it. The exception to this rule was that organization administrators can edit channel names and descriptions without having full access to the channel.  New in Zulip 10.0 (feature level 325). Prior to this change, the permission to administer channels was limited to realm administrators.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Organization administrators can administer every channel as though they were in this group without being explicitly listed here.  Note that a user must [have content access] to a channel in order to add other subscribers to the channel.
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanAdministerChannelGroup(canAdministerChannelGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

// The set of users who have permission to delete any message in the channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in `can_delete_any_message_group` were able delete any message in the organization.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Note that a user must [have content access] to a channel in order to delete any message in the channel.  Users present in the organization-level `can_delete_any_message_group` setting can always delete any message in the channel if they [have content access] to that channel.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

// The set of users who have permission to delete the messages that they have sent in the channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in the organization-level `can_delete_any_message_group` and `can_delete_own_message_group` settings were able delete their own messages in the organization.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Note that a user must [have content access] to a channel in order to delete their own message in the channel.  Users with permission to delete any message in the channel and users present in the organization-level `can_delete_own_message_group` setting can always delete their own messages in the channel if they [have content access] to that channel.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

// The set of users who have permission to move messages out of this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in `can_move_messages_between_channels_group` were able move messages between channels.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Note that a user must [have content access] to a channel in order to move messages out of the channel.  Channel administrators and users present in the organization-level `can_move_messages_between_channels_group` setting can always move messages out of the channel if they [have content access] to the channel.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

// The set of users who have permission to move messages within this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in `can_move_messages_between_topics_group` were able move messages between topics of a channel.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Note that a user must [have content access] to a channel in order to move messages within the channel.  Channel administrators and users present in the organization-level `can_move_messages_between_topics_group` setting can always move messages within the channel if they [have content access] to the channel.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

// The set of users who have permission to post in this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 333). Previously `stream_post_policy` field used to control the permission to post in the channel.
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.
func (r UpdateChannelRequest) CanSendMessageGroup(canSendMessageGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

// The set of users who have permission to subscribe themselves to this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 10.0 (feature level 357).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Everyone, excluding guests, can subscribe to any public channel irrespective of this setting.  Users in this group can subscribe to a private channel as well.  Note that a user must [have content access] to a channel and permission to administer the channel in order to modify this setting.
//
// [have content access]: https://zulip.com/help/channel-permissions
func (r UpdateChannelRequest) CanSubscribeGroup(canSubscribeGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

// The set of users who have permission to to resolve topics in this channel expressed as an [update to a group-setting value].
//
// **Changes**: New in Zulip 11.0 (feature level 402).
//
// [update to a group-setting value]: https://zulip.com/api/group-setting-values#updating-group-setting-values  Users who have similar realm-level permissions can resolve topics in a channel regardless of the value of this setting.
func (r UpdateChannelRequest) CanResolveTopicsGroup(canResolveTopicsGroup zulip.GroupSettingValueUpdate) UpdateChannelRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

func (r UpdateChannelRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateChannelExecute(r)
}

// UpdateChannel Update a channel
//
// Configure the channel with the Id `channelId`. This endpoint supports
// an organization administrator editing any property of a channel,
// including:
//
//   - Channel [name] and [description]
//   - Channel [permissions], including
//
// [privacy] and [who can send].
//
// Note that an organization administrator's ability to change a
// [private channel's permissions]
// depends on them being subscribed to the channel.
//
// *Changes**: Before Zulip 10.0 (feature level 362), channel privacy could not be
// edited for archived channels.
//
// Removed `stream_post_policy` and `is_announcement_only`
// parameters in Zulip 10.0 (feature level 333), as permission to post
// in the channel is now controlled by `can_send_message_group`.
//
// [name]: https://zulip.com/help/rename-a-channel
// [description]: https://zulip.com/help/change-the-channel-description
// [permissions]: https://zulip.com/help/channel-permissions
// [privacy]: https://zulip.com/help/change-the-privacy-of-a-channel
// [private channel's permissions]: https://zulip.com/help/channel-permissions#private-channels
// [who can send]: https://zulip.com/help/channel-posting-policy
func (s *channelsService) UpdateChannel(ctx context.Context, channelId int64) UpdateChannelRequest {
	return UpdateChannelRequest{
		apiService: s,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
func (s *channelsService) UpdateChannelExecute(r UpdateChannelRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/streams/{stream_id}"
	)

	path := strings.Replace(endpoint, "{stream_id}", IdToString(r.channelId), -1)

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "description", r.description)
	AddOptionalParam(form, "new_name", r.newName)
	AddOptionalParam(form, "is_private", r.isPrivate)
	AddOptionalParam(form, "is_web_public", r.isWebPublic)
	AddOptionalParam(form, "history_public_to_subscribers", r.historyPublicToSubscribers)
	AddOptionalParam(form, "is_default_stream", r.isDefaultChannel)
	if err := AddOptionalJSONParam(form, "message_retention_days", r.messageRetentionDays); err != nil {
		return nil, nil, err
	}
	AddOptionalParam(form, "is_archived", r.isArchived)
	AddOptionalParam(form, "folder_id", r.folderId)
	AddOptionalParam(form, "topics_policy", r.topicsPolicy)
	if err := AddOptionalJSONParam(form, "can_add_subscribers_group", r.canAddSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_remove_subscribers_group", r.canRemoveSubscribersGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_administer_channel_group", r.canAdministerChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_any_message_group", r.canDeleteAnyMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_delete_own_message_group", r.canDeleteOwnMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_out_of_channel_group", r.canMoveMessagesOutOfChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_move_messages_within_channel_group", r.canMoveMessagesWithinChannelGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_send_message_group", r.canSendMessageGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_subscribe_group", r.canSubscribeGroup); err != nil {
		return nil, nil, err
	}
	if err := AddOptionalJSONParam(form, "can_resolve_topics_group", r.canResolveTopicsGroup); err != nil {
		return nil, nil, err
	}
	req, err := PrepareRequest(r.ctx, s.client, path, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UpdateSubscriptionSettingsRequest struct {
	ctx              context.Context
	apiService       APIChannels
	subscriptionData *[]zulip.SubscriptionData
}

// A list of objects that describe the changes that should be applied in each subscription. Each object represents a subscription, and must have a `stream_id` key that identifies the channel, as well as the `property` being modified and its new `value`.
func (r UpdateSubscriptionSettingsRequest) SubscriptionData(subscriptionData []zulip.SubscriptionData) UpdateSubscriptionSettingsRequest {
	r.subscriptionData = &subscriptionData
	return r
}

func (r UpdateSubscriptionSettingsRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateSubscriptionSettingsExecute(r)
}

// UpdateSubscriptionSettings Update subscription settings
//
// This endpoint is used to update the user's personal settings for the
// channels they are subscribed to, including muting, color, pinning, and
// per-channel notification settings.
//
// *Changes**: Prior to Zulip 5.0 (feature level 111), response
// object included the `subscription_data` in the
// request. The endpoint now returns the more ergonomic
// [`ignored_parameters_unsupported`] array instead.
//
// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
func (s *channelsService) UpdateSubscriptionSettings(ctx context.Context) UpdateSubscriptionSettingsRequest {
	return UpdateSubscriptionSettingsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) UpdateSubscriptionSettingsExecute(r UpdateSubscriptionSettingsRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/users/me/subscriptions/properties"
	)
	if r.subscriptionData == nil {
		return nil, nil, fmt.Errorf("subscriptionData is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "subscription_data", r.subscriptionData)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UpdateSubscriptionsRequest struct {
	ctx        context.Context
	apiService APIChannels
	delete     *[]string
	add        *[]SubscriptionRequestWithColor
}

// SubscriptionRequestWithColor struct for SubscriptionRequestWithColor
type SubscriptionRequestWithColor struct {
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
}

// A list of channel names to unsubscribe from.
func (r UpdateSubscriptionsRequest) Delete(delete []string) UpdateSubscriptionsRequest {
	r.delete = &delete
	return r
}

// A list of objects describing which channels to subscribe to, optionally including per-user subscription parameters (e.g. color) and if the channel is to be created, its description.
func (r UpdateSubscriptionsRequest) Add(add []SubscriptionRequestWithColor) UpdateSubscriptionsRequest {
	r.add = &add
	return r
}

func (r UpdateSubscriptionsRequest) Execute() (*UpdateSubscriptionsResponse, *http.Response, error) {
	return r.apiService.UpdateSubscriptionsExecute(r)
}

// UpdateSubscriptions Update subscriptions
//
// Update which channels you are subscribed to.
//
// *Changes**: Before Zulip 10.0 (feature level 362),
// subscriptions in archived channels could not be modified.
func (s *channelsService) UpdateSubscriptions(ctx context.Context) UpdateSubscriptionsRequest {
	return UpdateSubscriptionsRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) UpdateSubscriptionsExecute(r UpdateSubscriptionsRequest) (*UpdateSubscriptionsResponse, *http.Response, error) {
	var (
		method   = http.MethodPatch
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &UpdateSubscriptionsResponse{}
		endpoint = "/users/me/subscriptions"
	)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddOptionalParam(form, "delete", r.delete)
	AddOptionalParam(form, "add", r.add)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type UpdateUserTopicRequest struct {
	ctx              context.Context
	apiService       APIChannels
	channelId        *int64
	topic            *string
	visibilityPolicy *zulip.VisibilityPolicy
}

// The Id of the channel to access.
func (r UpdateUserTopicRequest) ChannelId(channelId int64) UpdateUserTopicRequest {
	r.channelId = &channelId
	return r
}

// The topic for which the personal preferences needs to be updated. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the `max_topic_length` returned by the [`POST /register`] endpoint to determine the maximum topic length.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register] response is used for this parameter, it is interpreted as an empty string.
//
// **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
//
// [`POST /register`]: https://zulip.com/api/register-queue
func (r UpdateUserTopicRequest) Topic(topic string) UpdateUserTopicRequest {
	r.topic = &topic
	return r
}

// Controls which visibility policy to set.
//   - VisibilityPolicyNone
//   - VisibilityPolicyMuted
//   - VisibilityPolicyUnmuted
//   - VisibilityPolicyFollowed
//
// In an unmuted channel, a topic visibility policy of unmuted will have the same effect as the "None" visibility policy.
//
// **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.
func (r UpdateUserTopicRequest) VisibilityPolicy(visibilityPolicy zulip.VisibilityPolicy) UpdateUserTopicRequest {
	r.visibilityPolicy = &visibilityPolicy
	return r
}

func (r UpdateUserTopicRequest) Execute() (*zulip.Response, *http.Response, error) {
	return r.apiService.UpdateUserTopicExecute(r)
}

// UpdateUserTopic Update personal preferences for a topic
//
// This endpoint is used to update the personal preferences for a topic,
// such as the topic's visibility policy, which is used to implement
// [mute a topic] and related features.
//
// This endpoint can be used to update the visibility policy for the single
// channel and topic pair indicated by the parameters for a user.
//
// *Changes**: New in Zulip 7.0 (feature level 170). Previously,
// toggling whether a topic was muted or unmuted was managed by the
// [PATCH /users/me/subscriptions/muted_topics] endpoint.
//
// [mute a topic]: https://zulip.com/help/mute-a-topic
// [PATCH /users/me/subscriptions/muted_topics]: https://zulip.com/api/mute-topic
func (s *channelsService) UpdateUserTopic(ctx context.Context) UpdateUserTopicRequest {
	return UpdateUserTopicRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request
func (s *channelsService) UpdateUserTopicExecute(r UpdateUserTopicRequest) (*zulip.Response, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &zulip.Response{}
		endpoint = "/user_topics"
	)
	if r.channelId == nil {
		return nil, nil, fmt.Errorf("channelId is required and must be specified")
	}
	if r.topic == nil {
		return nil, nil, fmt.Errorf("topic is required and must be specified")
	}
	if r.visibilityPolicy == nil {
		return nil, nil, fmt.Errorf("visibilityPolicy is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	AddParam(form, "stream_id", r.channelId)
	AddParam(form, "topic", r.topic)
	AddParam(form, "visibility_policy", r.visibilityPolicy)
	req, err := PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}
