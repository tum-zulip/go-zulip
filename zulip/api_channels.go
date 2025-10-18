package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ChannelsAPI interface {

	/*
			AddDefaultChannel Add a default channel

			Add a channel to the set of [default channels][default-channels]
		for new users joining the organization.

		[default-channels]: /help/set-default-channels-for-new-users


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AddDefaultChannelRequest
	*/
	AddDefaultChannel(ctx context.Context) AddDefaultChannelRequest

	// AddDefaultChannelExecute executes the request
	//  @return Response
	AddDefaultChannelExecute(r AddDefaultChannelRequest) (*Response, *http.Response, error)

	/*
		ArchiveChannel Archive a channel

		[Archive the channel](zulip.com/help/archive-a-channel) with the Id `channelId`.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param channelId The Id of the channel to access.
		@return ArchiveChannelRequest
	*/
	ArchiveChannel(ctx context.Context, channelId int64) ArchiveChannelRequest

	// ArchiveChannelExecute executes the request
	//  @return Response
	ArchiveChannelExecute(r ArchiveChannelRequest) (*Response, *http.Response, error)

	/*
			CreateBigBlueButtonVideoCall Create BigBlueButton video call

			Create a video call URL for a BigBlueButton video call.
		Requires [BigBlueButton 2.4+](/integrations/doc/big-blue-button)
		to be configured on the Zulip server.

		The acting user will be given the moderator role on the call.

		**Changes**: Prior to Zulip 10.0 (feature level 337), every
		user was given the moderator role on BigBlueButton calls, via
		encoding a moderator password in the generated URLs.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return CreateBigBlueButtonVideoCallRequest
	*/
	CreateBigBlueButtonVideoCall(ctx context.Context) CreateBigBlueButtonVideoCallRequest

	// CreateBigBlueButtonVideoCallExecute executes the request
	//  @return CreateBigBlueButtonVideoCallResponse
	CreateBigBlueButtonVideoCallExecute(r CreateBigBlueButtonVideoCallRequest) (*CreateBigBlueButtonVideoCallResponse, *http.Response, error)

	/*
			CreateChannel Create a channel

			Create a new [channel](zulip.com/help/create-channels) and optionally
		subscribe users to the newly created channel. The initial [channel settings](zulip.com/api/update-stream
		will be determined by the optional parameters, like `invite_only`, detailed below.

		**Changes**: New in Zulip 11.0 (feature level 417). Previously, this was only possible via
		the [`POST /api/subscribe`](zulip.com/api/subscribe) endpoint, which handled both creation and subscription.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return CreateChannelRequest
	*/
	CreateChannel(ctx context.Context) CreateChannelRequest

	// CreateChannelExecute executes the request
	//  @return CreateChannelResponse
	CreateChannelExecute(r CreateChannelRequest) (*CreateChannelResponse, *http.Response, error)

	/*
			CreateChannelFolder Create a channel folder

			Create a new channel folder, that will be used to organize
		channels in left sidebar.

		Only organization administrators can create a new channel
		folder.

		**Changes**: New in Zulip 11.0 (feature level 389).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return CreateChannelFolderRequest
	*/
	CreateChannelFolder(ctx context.Context) CreateChannelFolderRequest

	// CreateChannelFolderExecute executes the request
	//  @return CreateChannelFolderResponse
	CreateChannelFolderExecute(r CreateChannelFolderRequest) (*CreateChannelFolderResponse, *http.Response, error)

	/*
			DeleteTopic Delete a topic

			Delete all messages in a topic.

		Topics are a field on messages (not an independent data structure), so
		deleting all the messages in the topic deletes the topic from Zulip.

		Because this endpoint deletes messages in batches, it is possible for
		the request to time out after only deleting some messages in the topic.
		When this happens, the `complete` boolean field in the success response
		will be `false`. Clients should repeat the request when handling such a
		response. If all messages in the topic were deleted, then the success
		response will return `"complete": true`.

		**Changes**: Before Zulip 9.0 (feature level 256), the server never sent
		[`stream` op: `update`](zulip.com/api/get-events#stream-update events with an
		updated `first_message_id` for a channel when the oldest message that
		had been sent to it changed.

		Before Zulip 8.0 (feature level 211), if the server's
		processing was interrupted by a timeout, but some messages in the topic
		were deleted, then it would return `"result": "partially_completed"`,
		along with a `code` field for an error string, in the success response
		to indicate that there was a timeout and that the client should repeat
		the request.

		As of Zulip 6.0 (feature level 154), instead of returning an error
		response when a request times out after successfully deleting some of
		the messages in the topic, a success response is returned with
		`"result": "partially_completed"` to indicate that some messages were
		deleted.

		Before Zulip 6.0 (feature level 147), this request did a single atomic
		operation, which could time out for very large topics. As of this
		feature level, messages are deleted in batches, starting with the newest
		messages, so that progress is made even if the request times out and
		returns an error.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelId The Id of the channel to access.
			@return DeleteTopicRequest
	*/
	DeleteTopic(ctx context.Context, channelId int64) DeleteTopicRequest

	// DeleteTopicExecute executes the request
	//  @return MarkAllAsReadResponse
	DeleteTopicExecute(r DeleteTopicRequest) (*MarkAllAsReadResponse, *http.Response, error)

	/*
			GetChannelFolders Get channel folders

			Fetches all of the channel folders in the organization.
		The folders are sorted by the `order` field.

		**Changes**: Before Zulip 11.0 (feature level 414),
		these were sorted by Id. (The `order` field didn't exist).

		New in Zulip 11.0 (feature level 389).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return GetChannelFoldersRequest
	*/
	GetChannelFolders(ctx context.Context) GetChannelFoldersRequest

	// GetChannelFoldersExecute executes the request
	//  @return GetChannelFoldersResponse
	GetChannelFoldersExecute(r GetChannelFoldersRequest) (*GetChannelFoldersResponse, *http.Response, error)

	/*
			GetChannelById Get a channel by Id

			Fetch details for the channel with the Id `channelId`.

		**Changes**: New in Zulip 6.0 (feature level 132).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelId The Id of the channel to access.
			@return GetChannelByIdRequest
	*/
	GetChannelById(ctx context.Context, channelId int64) GetChannelByIdRequest

	// GetChannelByIdExecute executes the request
	//  @return GetChannelResponse
	GetChannelByIdExecute(r GetChannelByIdRequest) (*GetChannelResponse, *http.Response, error)

	/*
			GetChannelEmailAddress Get channel's email address

			Get email address of a channel.

		**Changes**: New in Zulip 8.0 (feature level 226).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelId The Id of the channel to access.
			@return GetChannelEmailAddressRequest
	*/
	GetChannelEmailAddress(ctx context.Context, channelId int64) GetChannelEmailAddressRequest

	// GetChannelEmailAddressExecute executes the request
	//  @return GetChannelEmailAddressResponse
	GetChannelEmailAddressExecute(r GetChannelEmailAddressRequest) (*GetChannelEmailAddressResponse, *http.Response, error)

	/*
		GetChannelId Get channel Id

		Get the unique Id of a given channel.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetChannelIdRequest
	*/
	GetChannelId(ctx context.Context) GetChannelIdRequest

	// GetChannelIdExecute executes the request
	//  @return GetChannelIdResponse
	GetChannelIdExecute(r GetChannelIdRequest) (*GetChannelIdResponse, *http.Response, error)

	/*
			GetChannelTopics Get topics in a channel

			Get all topics the user has access to in a specific channel.

		Note that for [private channels with
		protected history](zulip.com/help/channel-permissions#private-channels,
		the user will only have access to topics of messages sent after they
		[subscribed to](zulip.com/api/subscribe) the channel. Similarly, a user's
		[bot](zulip.com/help/bots-overview#bot-type will only have access to messages
		sent after the bot was subscribed to the channel, instead of when the
		user subscribed.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelId The Id of the channel to access.
			@return GetChannelTopicsRequest
	*/
	GetChannelTopics(ctx context.Context, channelId int64) GetChannelTopicsRequest

	// GetChannelTopicsExecute executes the request
	//  @return GetChannelTopicsResponse
	GetChannelTopicsExecute(r GetChannelTopicsRequest) (*GetChannelTopicsResponse, *http.Response, error)

	/*
		GetChannels Get all channels

		Get all channels that the user [has access to](zulip.com/help/channel-permissions.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetChannelsRequest
	*/
	GetChannels(ctx context.Context) GetChannelsRequest

	// GetChannelsExecute executes the request
	//  @return GetChannelsResponse
	GetChannelsExecute(r GetChannelsRequest) (*GetChannelsResponse, *http.Response, error)

	/*
		GetSubscribers Get channel subscribers

		Get all users subscribed to a channel.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param channelId The Id of the channel to access.
		@return GetSubscribersRequest
	*/
	GetSubscribers(ctx context.Context, channelId int64) GetSubscribersRequest

	// GetSubscribersExecute executes the request
	//  @return GetSubscribersResponse
	GetSubscribersExecute(r GetSubscribersRequest) (*GetSubscribersResponse, *http.Response, error)

	/*
			GetSubscriptionStatus Get subscription status

			Check whether a user is subscribed to a channel.

		**Changes**: New in Zulip 3.0 (feature level 12).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The target user's Id.
			@param channelId The Id of the channel to access.
			@return GetSubscriptionStatusRequest
	*/
	GetSubscriptionStatus(ctx context.Context, userId int64, channelId int64) GetSubscriptionStatusRequest

	// GetSubscriptionStatusExecute executes the request
	//  @return GetSubscriptionStatusResponse
	GetSubscriptionStatusExecute(r GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, *http.Response, error)

	/*
		GetSubscriptions Get subscribed channels

		Get all channels that the user is subscribed to.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetSubscriptionsRequest
	*/
	GetSubscriptions(ctx context.Context) GetSubscriptionsRequest

	// GetSubscriptionsExecute executes the request
	//  @return GetSubscriptionsResponse
	GetSubscriptionsExecute(r GetSubscriptionsRequest) (*GetSubscriptionsResponse, *http.Response, error)

	/*
			MuteTopic Topic muting

			[Mute or unmute a topic](zulip.com/help/mute-a-topic) within a channel that
		the current user is subscribed to.

		**Changes**: Deprecated in Zulip 7.0 (feature level 170). Clients connecting
		to newer servers should use the [POST /user_topics](zulip.com/api/update-user-topic
		endpoint, as this endpoint may be removed in a future release.

		Before Zulip 7.0 (feature level 169), this endpoint
		returned an error if asked to mute a topic that was already muted
		or asked to unmute a topic that had not previously been muted.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return MuteTopicRequest

			Deprecated
	*/
	MuteTopic(ctx context.Context) MuteTopicRequest

	// MuteTopicExecute executes the request
	//  @return Response
	// Deprecated
	MuteTopicExecute(r MuteTopicRequest) (*Response, *http.Response, error)

	/*
			PatchChannelFolders Reorder channel folders

			Given an array of channel folder Ids, this method will set the `order`
		property of all of the channel folders in the organization according to
		the order of the channel folder Ids specified in the request.

		**Changes**: New in Zulip 11.0 (feature level 414).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return PatchChannelFoldersRequest
	*/
	PatchChannelFolders(ctx context.Context) PatchChannelFoldersRequest

	// PatchChannelFoldersExecute executes the request
	//  @return Response
	PatchChannelFoldersExecute(r PatchChannelFoldersRequest) (*Response, *http.Response, error)

	/*
			RemoveDefaultChannel Remove a default channel

			Remove a channel from the set of [default channels][default-channels]
		for new users joining the organization.

		[default-channels]: /help/set-default-channels-for-new-users


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return RemoveDefaultChannelRequest
	*/
	RemoveDefaultChannel(ctx context.Context) RemoveDefaultChannelRequest

	// RemoveDefaultChannelExecute executes the request
	//  @return Response
	RemoveDefaultChannelExecute(r RemoveDefaultChannelRequest) (*Response, *http.Response, error)

	/*
			Subscribe Subscribe to a channel

			Subscribe one or more users to one or more channels.

		If any of the specified channels do not exist, they are automatically
		created. The initial [channel settings](zulip.com/api/update-stream) will be determined
		by the optional parameters, like `invite_only`, detailed below.

		Note that the ability to subscribe oneself and/or other users
		to a specified channel depends on the [channel's permissions
		settings](zulip.com/help/channel-permissions.

		**Changes**: Before Zulip 10.0 (feature level 362),
		subscriptions in archived channels could not be modified.

		Before Zulip 10.0 (feature level 357), the
		`can_subscribe_group` permission, which allows members of the
		group to subscribe themselves to the channel, did not exist.

		Before Zulip 10.0 (feature level 349), a user cannot subscribe
		other users to a private channel without being subscribed
		to that channel themselves. Now, If a user is part of
		`can_add_subscribers_group`, they can subscribe themselves or other
		users to a private channel without being subscribed to that channel.

		Removed `stream_post_policy` and `is_announcement_only`
		parameters in Zulip 10.0 (feature level 333), as permission to post
		in the channel is now controlled by `can_send_message_group`.

		Before Zulip 8.0 (feature level 208), if a user specified by the
		[`principals`][principals-param] parameter was a deactivated user,
		or did not exist, then an HTTP status code of 403 was returned with
		`code: "UNAUTHORIZED_PRINCIPAL"` in the error response. As of this
		feature level, an HTTP status code of 400 is returned with
		`code: "BAD_REQUEST"` in the error response for these cases.

		[principals-param]: /api/subscribe#parameter-principals


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return SubscribeRequest
	*/
	Subscribe(ctx context.Context) SubscribeRequest

	// SubscribeExecute executes the request
	//  @return SubscribeResponse
	SubscribeExecute(r SubscribeRequest) (*SubscribeResponse, *http.Response, error)

	/*
			Unsubscribe Unsubscribe from a channel

			Unsubscribe yourself or other users from one or more channels.

		In addition to managing the current user's subscriptions, this
		endpoint can be used to remove other users from channels. This
		is possible in 3 situations:

		- Organization administrators can remove any user from any
		  channel.
		- Users can remove a bot that they own from any channel that
		  the user [can access](zulip.com/help/channel-permissions.
		- Users can unsubscribe any user from a channel if they [have
		  access](zulip.com/help/channel-permissions) to the channel and are a
		  member of the [user group](zulip.com/api/get-user-groups) specified
		  by the [`can_remove_subscribers_group`][can-remove-parameter]
		  for the channel.

		**Changes**: Before Zulip 10.0 (feature level 362),
		subscriptions in archived channels could not be modified.

		Before Zulip 8.0 (feature level 208), if a user specified by
		the [`principals`][principals-param] parameter was a
		deactivated user, or did not exist, then an HTTP status code
		of 403 was returned with `code: "UNAUTHORIZED_PRINCIPAL"` in
		the error response. As of this feature level, an HTTP status
		code of 400 is returned with `code: "BAD_REQUEST"` in the
		error response for these cases.

		Before Zulip 8.0 (feature level 197),
		the `can_remove_subscribers_group` setting
		was named `can_remove_subscribers_group_id`.

		Before Zulip 7.0 (feature level 161), the
		`can_remove_subscribers_group_id` for all channels was always
		the system group for organization administrators.

		Before Zulip 6.0 (feature level 145), users had no special
		privileges for managing bots that they own.

		[principals-param]: /api/unsubscribe#parameter-principals
		[can-remove-parameter]: /api/subscribe#parameter-can_remove_subscribers_group


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UnsubscribeRequest
	*/
	Unsubscribe(ctx context.Context) UnsubscribeRequest

	// UnsubscribeExecute executes the request
	//  @return UnsubscribeResponse
	UnsubscribeExecute(r UnsubscribeRequest) (*UnsubscribeResponse, *http.Response, error)

	/*
			UpdateChannelFolder Update a channel folder

			Update the name or description of a channel folder.

		This endpoint is also used to archive and unarchive
		a channel folder.

		Only organization administrators can update a
		channel folder.

		**Changes**: New in Zulip 11.0 (feature level 389).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelFolderId The Id of the target channel folder.
			@return UpdateChannelFolderRequest
	*/
	UpdateChannelFolder(ctx context.Context, channelFolderId int64) UpdateChannelFolderRequest

	// UpdateChannelFolderExecute executes the request
	//  @return Response
	UpdateChannelFolderExecute(r UpdateChannelFolderRequest) (*Response, *http.Response, error)

	/*
			UpdateChannel Update a channel

			Configure the channel with the Id `channelId`. This endpoint supports
		an organization administrator editing any property of a channel,
		including:

		- Channel [name](zulip.com/help/rename-a-channel) and [description](zulip.com/help/change-the-channel-description
		- Channel [permissions](zulip.com/help/channel-permissions, including
		  [privacy](zulip.com/help/change-the-privacy-of-a-channel) and [who can
		  send](zulip.com/help/channel-posting-policy.

		Note that an organization administrator's ability to change a
		[private channel's permissions](zulip.com/help/channel-permissions#private-channels
		depends on them being subscribed to the channel.

		**Changes**: Before Zulip 10.0 (feature level 362), channel privacy could not be
		edited for archived channels.

		Removed `stream_post_policy` and `is_announcement_only`
		parameters in Zulip 10.0 (feature level 333), as permission to post
		in the channel is now controlled by `can_send_message_group`.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param channelId The Id of the channel to access.
			@return UpdateChannelRequest
	*/
	UpdateChannel(ctx context.Context, channelId int64) UpdateChannelRequest

	// UpdateChannelExecute executes the request
	//  @return Response
	UpdateChannelExecute(r UpdateChannelRequest) (*Response, *http.Response, error)

	/*
			UpdateSubscriptionSettings Update subscription settings

			This endpoint is used to update the user's personal settings for the
		channels they are subscribed to, including muting, color, pinning, and
		per-channel notification settings.

		**Changes**: Prior to Zulip 5.0 (feature level 111), response
		object included the `subscription_data` in the
		request. The endpoint now returns the more ergonomic
		[`ignored_parameters_unsupported`][ignored-parameters] array instead.

		[ignored-parameters]: /api/rest-error-handling#ignored-parameters


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdateSubscriptionSettingsRequest
	*/
	UpdateSubscriptionSettings(ctx context.Context) UpdateSubscriptionSettingsRequest

	// UpdateSubscriptionSettingsExecute executes the request
	//  @return Response
	UpdateSubscriptionSettingsExecute(r UpdateSubscriptionSettingsRequest) (*Response, *http.Response, error)

	/*
			UpdateSubscriptions Update subscriptions

			Update which channels you are subscribed to.

		**Changes**: Before Zulip 10.0 (feature level 362),
		subscriptions in archived channels could not be modified.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdateSubscriptionsRequest
	*/
	UpdateSubscriptions(ctx context.Context) UpdateSubscriptionsRequest

	// UpdateSubscriptionsExecute executes the request
	//  @return UpdateSubscriptionsResponse
	UpdateSubscriptionsExecute(r UpdateSubscriptionsRequest) (*UpdateSubscriptionsResponse, *http.Response, error)

	/*
			UpdateUserTopic Update personal preferences for a topic

			This endpoint is used to update the personal preferences for a topic,
		such as the topic's visibility policy, which is used to implement
		[mute a topic](zulip.com/help/mute-a-topic) and related features.

		This endpoint can be used to update the visibility policy for the single
		channel and topic pair indicated by the parameters for a user.

		**Changes**: New in Zulip 7.0 (feature level 170). Previously,
		toggling whether a topic was muted or unmuted was managed by the
		[PATCH /users/me/subscriptions/muted_topics](zulip.com/api/mute-topic) endpoint.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return UpdateUserTopicRequest
	*/
	UpdateUserTopic(ctx context.Context) UpdateUserTopicRequest

	// UpdateUserTopicExecute executes the request
	//  @return Response
	UpdateUserTopicExecute(r UpdateUserTopicRequest) (*Response, *http.Response, error)
}

type AddDefaultChannelRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  *int64
}

// The Id of the target channel.
func (r AddDefaultChannelRequest) ChannelId(channelId int64) AddDefaultChannelRequest {
	r.channelId = &channelId
	return r
}

func (r AddDefaultChannelRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.AddDefaultChannelExecute(r)
}

/*
AddDefaultChannel Add a default channel

Add a channel to the set of [default channels][default-channels]
for new users joining the organization.

[default-channels]: /help/set-default-channels-for-new-users

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddDefaultChannelRequest
*/
func (c *simpleClient) AddDefaultChannel(ctx context.Context) AddDefaultChannelRequest {
	return AddDefaultChannelRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) AddDefaultChannelExecute(r AddDefaultChannelRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/default_streams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.channelId == nil {
		return localVarReturnValue, nil, reportError("channelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "stream_id", r.channelId, "form", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ArchiveChannelRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  int64
}

func (r ArchiveChannelRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.ArchiveChannelExecute(r)
}

/*
ArchiveChannel Archive a channel

[Archive the channel](zulip.com/help/archive-a-channel) with the Id `channelId`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return ArchiveChannelRequest
*/
func (c *simpleClient) ArchiveChannel(ctx context.Context, channelId int64) ArchiveChannelRequest {
	return ArchiveChannelRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) ArchiveChannelExecute(r ArchiveChannelRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateBigBlueButtonVideoCallRequest struct {
	ctx         context.Context
	ApiService  ChannelsAPI
	meetingName *string
	voiceOnly   *bool
}

// Meeting name for the BigBlueButton video call.
func (r CreateBigBlueButtonVideoCallRequest) MeetingName(meetingName string) CreateBigBlueButtonVideoCallRequest {
	r.meetingName = &meetingName
	return r
}

// Configures whether the call is voice-only; if true, disables cameras for all users. Only the call creator/moderator can edit this configuration.  **Changes**: New in Zulip 10.0 (feature level 337).
func (r CreateBigBlueButtonVideoCallRequest) VoiceOnly(voiceOnly bool) CreateBigBlueButtonVideoCallRequest {
	r.voiceOnly = &voiceOnly
	return r
}

func (r CreateBigBlueButtonVideoCallRequest) Execute() (*CreateBigBlueButtonVideoCallResponse, *http.Response, error) {
	return r.ApiService.CreateBigBlueButtonVideoCallExecute(r)
}

/*
CreateBigBlueButtonVideoCall Create BigBlueButton video call

Create a video call URL for a BigBlueButton video call.
Requires [BigBlueButton 2.4+](/integrations/doc/big-blue-button)
to be configured on the Zulip server.

The acting user will be given the moderator role on the call.

**Changes**: Prior to Zulip 10.0 (feature level 337), every
user was given the moderator role on BigBlueButton calls, via
encoding a moderator password in the generated URLs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateBigBlueButtonVideoCallRequest
*/
func (c *simpleClient) CreateBigBlueButtonVideoCall(ctx context.Context) CreateBigBlueButtonVideoCallRequest {
	return CreateBigBlueButtonVideoCallRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateBigBlueButtonVideoCallResponse
func (c *simpleClient) CreateBigBlueButtonVideoCallExecute(r CreateBigBlueButtonVideoCallRequest) (*CreateBigBlueButtonVideoCallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBigBlueButtonVideoCallResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calls/bigbluebutton/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.meetingName == nil {
		return localVarReturnValue, nil, reportError("meetingName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "meeting_name", r.meetingName, "form", "")
	if r.voiceOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "voice_only", r.voiceOnly, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateChannelRequest struct {
	ctx                               context.Context
	ApiService                        ChannelsAPI
	name                              *string
	subscribers                       *[]int64
	description                       *string
	announce                          *bool
	inviteOnly                        *bool
	isWebPublic                       *bool
	isDefaultChannel                  *bool
	folderId                          *int64
	sendNewSubscriptionMessages       *bool
	topicsPolicy                      *TopicsPolicy
	historyPublicToSubscribers        *bool
	messageRetentionDays              *MessageRetentionDays
	canAddSubscribersGroup            *GroupSettingValue
	canDeleteAnyMessageGroup          *GroupSettingValue
	canDeleteOwnMessageGroup          *GroupSettingValue
	canRemoveSubscribersGroup         *GroupSettingValue
	canAdministerChannelGroup         *GroupSettingValue
	canMoveMessagesOutOfChannelGroup  *GroupSettingValue
	canMoveMessagesWithinChannelGroup *GroupSettingValue
	canSendMessageGroup               *GroupSettingValue
	canSubscribeGroup                 *GroupSettingValue
	canResolveTopicsGroup             *GroupSettingValue
}

// The name of the new channel.  Clients should use the &#x60;max_stream_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel name length.
func (r CreateChannelRequest) Name(name string) CreateChannelRequest {
	r.name = &name
	return r
}

// A list of user Ids of the users to be subscribed to the new channel.
func (r CreateChannelRequest) Subscribers(subscribers []int64) CreateChannelRequest {
	r.subscribers = &subscribers
	return r
}

// The [description](zulip.com/help/change-the-channel-description) to use for the new channel being created, in text/markdown format.  Clients should use the &#x60;max_stream_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel description length.
func (r CreateChannelRequest) Description(description string) CreateChannelRequest {
	r.description = &description
	return r
}

// This determines whether [notification bot](zulip.com/help/configure-automated-notices) will send an announcement about the new channel&#39;s creation.
func (r CreateChannelRequest) Announce(announce bool) CreateChannelRequest {
	r.announce = &announce
	return r
}

// This parameter and the ones that follow are used to request an initial configuration of the new channel.  This parameter determines whether the newly created channel will be a private channel.
func (r CreateChannelRequest) InviteOnly(inviteOnly bool) CreateChannelRequest {
	r.inviteOnly = &inviteOnly
	return r
}

// This parameter determines whether the newly created channel will be a web-public channel.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current user to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html
func (r CreateChannelRequest) IsWebPublic(isWebPublic bool) CreateChannelRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// This parameter determines whether the newly created channel will be added as a [default channel][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users
func (r CreateChannelRequest) IsDefaultChannel(isDefaultChannel bool) CreateChannelRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

// The Id of the folder to which the channel belongs.  Is &#x60;null&#x60; if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).
func (r CreateChannelRequest) FolderId(folderId int64) CreateChannelRequest {
	r.folderId = &folderId
	return r
}

// Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than &#x60;max_bulk_new_subscription_messages&#x60; (available in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.
func (r CreateChannelRequest) SendNewSubscriptionMessages(sendNewSubscriptionMessages bool) CreateChannelRequest {
	r.sendNewSubscriptionMessages = &sendNewSubscriptionMessages
	return r
}

func (r CreateChannelRequest) TopicsPolicy(topicsPolicy TopicsPolicy) CreateChannelRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

// Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.
func (r CreateChannelRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) CreateChannelRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

func (r CreateChannelRequest) MessageRetentionDays(messageRetentionDays MessageRetentionDays) CreateChannelRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

func (r CreateChannelRequest) CanAddSubscribersGroup(canAddSubscribersGroup GroupSettingValue) CreateChannelRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

func (r CreateChannelRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup GroupSettingValue) CreateChannelRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

func (r CreateChannelRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup GroupSettingValue) CreateChannelRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

func (r CreateChannelRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup GroupSettingValue) CreateChannelRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

func (r CreateChannelRequest) CanAdministerChannelGroup(canAdministerChannelGroup GroupSettingValue) CreateChannelRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

func (r CreateChannelRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup GroupSettingValue) CreateChannelRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

func (r CreateChannelRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup GroupSettingValue) CreateChannelRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

func (r CreateChannelRequest) CanSendMessageGroup(canSendMessageGroup GroupSettingValue) CreateChannelRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

func (r CreateChannelRequest) CanSubscribeGroup(canSubscribeGroup GroupSettingValue) CreateChannelRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

func (r CreateChannelRequest) CanResolveTopicsGroup(canResolveTopicsGroup GroupSettingValue) CreateChannelRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

func (r CreateChannelRequest) Execute() (*CreateChannelResponse, *http.Response, error) {
	return r.ApiService.CreateChannelExecute(r)
}

/*
CreateChannel Create a channel

Create a new [channel](zulip.com/help/create-channels) and optionally
subscribe users to the newly created channel. The initial [channel settings](zulip.com/api/update-stream
will be determined by the optional parameters, like `invite_only`, detailed below.

**Changes**: New in Zulip 11.0 (feature level 417). Previously, this was only possible via
the [`POST /api/subscribe`](zulip.com/api/subscribe) endpoint, which handled both creation and subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateChannelRequest
*/
func (c *simpleClient) CreateChannel(ctx context.Context) CreateChannelRequest {
	return CreateChannelRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateChannelResponse
func (c *simpleClient) CreateChannelExecute(r CreateChannelRequest) (*CreateChannelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateChannelResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.subscribers == nil {
		return localVarReturnValue, nil, reportError("subscribers is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.subscribers != nil {
		paramJsonSubscribers, err := parameterToJson(*r.subscribers)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("subscribers", paramJsonSubscribers)
	}
	if r.announce != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "announce", r.announce, "form", "")
	}
	if r.inviteOnly != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_only", r.inviteOnly, "form", "")
	}
	if r.isWebPublic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_web_public", r.isWebPublic, "form", "")
	}
	if r.isDefaultChannel != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_default_stream", r.isDefaultChannel, "form", "")
	}
	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "folder_id", r.folderId, "form", "")
	}
	if r.sendNewSubscriptionMessages != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_new_subscription_messages", r.sendNewSubscriptionMessages, "", "")
	}
	if r.topicsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topics_policy", r.topicsPolicy, "form", "")
	}
	if r.historyPublicToSubscribers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "history_public_to_subscribers", r.historyPublicToSubscribers, "form", "")
	}
	if r.messageRetentionDays != nil {
		paramJson, err := parameterToJson(*r.messageRetentionDays)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("message_retention_days", paramJson)
	}
	if r.canAddSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canAddSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_add_subscribers_group", paramJson)
	}
	if r.canDeleteAnyMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteAnyMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_any_message_group", paramJson)
	}
	if r.canDeleteOwnMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteOwnMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_own_message_group", paramJson)
	}
	if r.canRemoveSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canRemoveSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_remove_subscribers_group", paramJson)
	}
	if r.canAdministerChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canAdministerChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_administer_channel_group", paramJson)
	}
	if r.canMoveMessagesOutOfChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesOutOfChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_out_of_channel_group", paramJson)
	}
	if r.canMoveMessagesWithinChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesWithinChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_within_channel_group", paramJson)
	}
	if r.canSendMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canSendMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_send_message_group", paramJson)
	}
	if r.canSubscribeGroup != nil {
		paramJson, err := parameterToJson(*r.canSubscribeGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_subscribe_group", paramJson)
	}
	if r.canResolveTopicsGroup != nil {
		paramJson, err := parameterToJson(*r.canResolveTopicsGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_resolve_topics_group", paramJson)
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateChannelFolderRequest struct {
	ctx         context.Context
	ApiService  ChannelsAPI
	name        *string
	description *string
}

// The name of the channel folder.  Clients should use the &#x60;max_channel_folder_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel folder name length.
func (r CreateChannelFolderRequest) Name(name string) CreateChannelFolderRequest {
	r.name = &name
	return r
}

// The description of the channel folder.  Clients should use the &#x60;max_channel_folder_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel folder description length.
func (r CreateChannelFolderRequest) Description(description string) CreateChannelFolderRequest {
	r.description = &description
	return r
}

func (r CreateChannelFolderRequest) Execute() (*CreateChannelFolderResponse, *http.Response, error) {
	return r.ApiService.CreateChannelFolderExecute(r)
}

/*
CreateChannelFolder Create a channel folder

Create a new channel folder, that will be used to organize
channels in left sidebar.

Only organization administrators can create a new channel
folder.

**Changes**: New in Zulip 11.0 (feature level 389).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateChannelFolderRequest
*/
func (c *simpleClient) CreateChannelFolder(ctx context.Context) CreateChannelFolderRequest {
	return CreateChannelFolderRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateChannelFolderResponse
func (c *simpleClient) CreateChannelFolderExecute(r CreateChannelFolderRequest) (*CreateChannelFolderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateChannelFolderResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channel_folders/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteTopicRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  int64
	topicName  *string
}

// The name of the topic to delete.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](zulip.com/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
func (r DeleteTopicRequest) TopicName(topicName string) DeleteTopicRequest {
	r.topicName = &topicName
	return r
}

func (r DeleteTopicRequest) Execute() (*MarkAllAsReadResponse, *http.Response, error) {
	return r.ApiService.DeleteTopicExecute(r)
}

/*
DeleteTopic Delete a topic

Delete all messages in a topic.

Topics are a field on messages (not an independent data structure), so
deleting all the messages in the topic deletes the topic from Zulip.

Because this endpoint deletes messages in batches, it is possible for
the request to time out after only deleting some messages in the topic.
When this happens, the `complete` boolean field in the success response
will be `false`. Clients should repeat the request when handling such a
response. If all messages in the topic were deleted, then the success
response will return `"complete": true`.

**Changes**: Before Zulip 9.0 (feature level 256), the server never sent
[`stream` op: `update`](zulip.com/api/get-events#stream-update events with an
updated `first_message_id` for a channel when the oldest message that
had been sent to it changed.

Before Zulip 8.0 (feature level 211), if the server's
processing was interrupted by a timeout, but some messages in the topic
were deleted, then it would return `"result": "partially_completed"`,
along with a `code` field for an error string, in the success response
to indicate that there was a timeout and that the client should repeat
the request.

As of Zulip 6.0 (feature level 154), instead of returning an error
response when a request times out after successfully deleting some of
the messages in the topic, a success response is returned with
`"result": "partially_completed"` to indicate that some messages were
deleted.

Before Zulip 6.0 (feature level 147), this request did a single atomic
operation, which could time out for very large topics. As of this
feature level, messages are deleted in batches, starting with the newest
messages, so that progress is made even if the request times out and
returns an error.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return DeleteTopicRequest
*/
func (c *simpleClient) DeleteTopic(ctx context.Context, channelId int64) DeleteTopicRequest {
	return DeleteTopicRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return MarkAllAsReadResponse
func (c *simpleClient) DeleteTopicExecute(r DeleteTopicRequest) (*MarkAllAsReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MarkAllAsReadResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}/delete_topic"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.topicName == nil {
		return localVarReturnValue, nil, reportError("topicName is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "topic_name", r.topicName, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelFoldersRequest struct {
	ctx             context.Context
	ApiService      ChannelsAPI
	includeArchived *bool
}

// Whether to include archived channel folders in the response.
func (r GetChannelFoldersRequest) IncludeArchived(includeArchived bool) GetChannelFoldersRequest {
	r.includeArchived = &includeArchived
	return r
}

func (r GetChannelFoldersRequest) Execute() (*GetChannelFoldersResponse, *http.Response, error) {
	return r.ApiService.GetChannelFoldersExecute(r)
}

/*
GetChannelFolders Get channel folders

Fetches all of the channel folders in the organization.
The folders are sorted by the `order` field.

**Changes**: Before Zulip 11.0 (feature level 414),
these were sorted by Id. (The `order` field didn't exist).

New in Zulip 11.0 (feature level 389).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetChannelFoldersRequest
*/
func (c *simpleClient) GetChannelFolders(ctx context.Context) GetChannelFoldersRequest {
	return GetChannelFoldersRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetChannelFoldersResponse
func (c *simpleClient) GetChannelFoldersExecute(r GetChannelFoldersRequest) (*GetChannelFoldersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelFoldersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channel_folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.includeArchived != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "include_archived", r.includeArchived, "form", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelByIdRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  int64
}

func (r GetChannelByIdRequest) Execute() (*GetChannelResponse, *http.Response, error) {
	return r.ApiService.GetChannelByIdExecute(r)
}

/*
GetChannelById Get a channel by Id

Fetch details for the channel with the Id `channelId`.

**Changes**: New in Zulip 6.0 (feature level 132).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return GetChannelByIdRequest
*/
func (c *simpleClient) GetChannelById(ctx context.Context, channelId int64) GetChannelByIdRequest {
	return GetChannelByIdRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return GetChannelResponse
func (c *simpleClient) GetChannelByIdExecute(r GetChannelByIdRequest) (*GetChannelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelEmailAddressRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  int64
	senderId   *int64
}

// The Id of a user or bot which should appear as the sender when messages are sent to the channel using the returned channel email address.  &#x60;sender_id&#x60; can be:  - Id of the current user. - Id of the Email gateway bot. (Default value) - Id of a bot owned by the current user.  **Changes**: New in Zulip 10.0 (feature level 335).  Previously, the sender was always Email gateway bot.
func (r GetChannelEmailAddressRequest) SenderId(senderId int64) GetChannelEmailAddressRequest {
	r.senderId = &senderId
	return r
}

func (r GetChannelEmailAddressRequest) Execute() (*GetChannelEmailAddressResponse, *http.Response, error) {
	return r.ApiService.GetChannelEmailAddressExecute(r)
}

/*
GetChannelEmailAddress Get channel's email address

Get email address of a channel.

**Changes**: New in Zulip 8.0 (feature level 226).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return GetChannelEmailAddressRequest
*/
func (c *simpleClient) GetChannelEmailAddress(ctx context.Context, channelId int64) GetChannelEmailAddressRequest {
	return GetChannelEmailAddressRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return GetChannelEmailAddressResponse
func (c *simpleClient) GetChannelEmailAddressExecute(r GetChannelEmailAddressRequest) (*GetChannelEmailAddressResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelEmailAddressResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}/email_address"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.senderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sender_id", r.senderId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelIdRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channel    *string
}

// The name of the channel to access.
func (r GetChannelIdRequest) Channel(channel string) GetChannelIdRequest {
	r.channel = &channel
	return r
}

func (r GetChannelIdRequest) Execute() (*GetChannelIdResponse, *http.Response, error) {
	return r.ApiService.GetChannelIdExecute(r)
}

/*
GetChannelId Get channel Id

Get the unique Id of a given channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetChannelIdRequest
*/
func (c *simpleClient) GetChannelId(ctx context.Context) GetChannelIdRequest {
	return GetChannelIdRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetChannelIdResponse
func (c *simpleClient) GetChannelIdExecute(r GetChannelIdRequest) (*GetChannelIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelIdResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/get_stream_id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.channel == nil {
		return localVarReturnValue, nil, reportError("channel is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "stream", r.channel, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelTopicsRequest struct {
	ctx                 context.Context
	ApiService          ChannelsAPI
	channelId           int64
	allowEmptyTopicName *bool
}

// Whether the client supports processing the empty string as a topic name in the returned data.  If &#x60;false&#x60;, the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) response is returned replacing the empty string as the topic name.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.
func (r GetChannelTopicsRequest) AllowEmptyTopicName(allowEmptyTopicName bool) GetChannelTopicsRequest {
	r.allowEmptyTopicName = &allowEmptyTopicName
	return r
}

func (r GetChannelTopicsRequest) Execute() (*GetChannelTopicsResponse, *http.Response, error) {
	return r.ApiService.GetChannelTopicsExecute(r)
}

/*
GetChannelTopics Get topics in a channel

Get all topics the user has access to in a specific channel.

Note that for [private channels with
protected history](zulip.com/help/channel-permissions#private-channels,
the user will only have access to topics of messages sent after they
[subscribed to](zulip.com/api/subscribe) the channel. Similarly, a user's
[bot](zulip.com/help/bots-overview#bot-type will only have access to messages
sent after the bot was subscribed to the channel, instead of when the
user subscribed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return GetChannelTopicsRequest
*/
func (c *simpleClient) GetChannelTopics(ctx context.Context, channelId int64) GetChannelTopicsRequest {
	return GetChannelTopicsRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return GetChannelTopicsResponse
func (c *simpleClient) GetChannelTopicsExecute(r GetChannelTopicsRequest) (*GetChannelTopicsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelTopicsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/{stream_id}/topics"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.allowEmptyTopicName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "allow_empty_topic_name", r.allowEmptyTopicName, "form", "")
	} else {
		var defaultValue bool = false
		r.allowEmptyTopicName = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetChannelsRequest struct {
	ctx                     context.Context
	ApiService              ChannelsAPI
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

// Whether to exclude archived streams from the results.  **Changes**: New in Zulip 10.0 (feature level 315).
func (r GetChannelsRequest) ExcludeArchived(excludeArchived bool) GetChannelsRequest {
	r.excludeArchived = &excludeArchived
	return r
}

// Deprecated parameter to include all channels. The user must have administrative privileges to use this parameter.  **Changes**: Deprecated in Zulip 10.0 (feature level 356). Clients interacting with newer servers should use the equivalent &#x60;include_all&#x60; parameter, which does not incorrectly hint that this parameter, and not &#x60;exclude_archived&#x60;, controls whether archived channels appear in the response.
// Deprecated
func (r GetChannelsRequest) IncludeAllActive(includeAllActive bool) GetChannelsRequest {
	r.includeAllActive = &includeAllActive
	return r
}

// Include all channels that the user has metadata access to.  For organization administrators, this will be all channels in the organization, since organization administrators implicitly have metadata access to all channels.  **Changes**: New in Zulip 10.0 (feature level 356). On older versions, use &#x60;include_all_active&#x60;, which this replaces.
func (r GetChannelsRequest) IncludeAll(includeAll bool) GetChannelsRequest {
	r.includeAll = &includeAll
	return r
}

// Include all default channels for the user&#39;s realm.
func (r GetChannelsRequest) IncludeDefault(includeDefault bool) GetChannelsRequest {
	r.includeDefault = &includeDefault
	return r
}

// If the user is a bot, include all channels that the bot&#39;s owner is subscribed to.
func (r GetChannelsRequest) IncludeOwnerSubscribed(includeOwnerSubscribed bool) GetChannelsRequest {
	r.includeOwnerSubscribed = &includeOwnerSubscribed
	return r
}

// Include all the channels that the user has content access to.  **Changes**: New in Zulip 10.0 (feature level 356).
func (r GetChannelsRequest) IncludeCanAccessContent(includeCanAccessContent bool) GetChannelsRequest {
	r.includeCanAccessContent = &includeCanAccessContent
	return r
}

func (r GetChannelsRequest) Execute() (*GetChannelsResponse, *http.Response, error) {
	return r.ApiService.GetChannelsExecute(r)
}

/*
GetChannels Get all channels

Get all channels that the user [has access to](zulip.com/help/channel-permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetChannelsRequest
*/
func (c *simpleClient) GetChannels(ctx context.Context) GetChannelsRequest {
	return GetChannelsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetChannelsResponse
func (c *simpleClient) GetChannelsExecute(r GetChannelsRequest) (*GetChannelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChannelsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includePublic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_public", r.includePublic, "form", "")
	} else {
		var defaultValue bool = true
		r.includePublic = &defaultValue
	}
	if r.includeWebPublic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_web_public", r.includeWebPublic, "form", "")
	} else {
		var defaultValue bool = false
		r.includeWebPublic = &defaultValue
	}
	if r.includeSubscribed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_subscribed", r.includeSubscribed, "form", "")
	} else {
		var defaultValue bool = true
		r.includeSubscribed = &defaultValue
	}
	if r.excludeArchived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_archived", r.excludeArchived, "form", "")
	} else {
		var defaultValue bool = true
		r.excludeArchived = &defaultValue
	}
	if r.includeAllActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_all_active", r.includeAllActive, "form", "")
	} else {
		var defaultValue bool = false
		r.includeAllActive = &defaultValue
	}
	if r.includeAll != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_all", r.includeAll, "form", "")
	} else {
		var defaultValue bool = false
		r.includeAll = &defaultValue
	}
	if r.includeDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_default", r.includeDefault, "form", "")
	} else {
		var defaultValue bool = false
		r.includeDefault = &defaultValue
	}
	if r.includeOwnerSubscribed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_owner_subscribed", r.includeOwnerSubscribed, "form", "")
	} else {
		var defaultValue bool = false
		r.includeOwnerSubscribed = &defaultValue
	}
	if r.includeCanAccessContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_can_access_content", r.includeCanAccessContent, "form", "")
	} else {
		var defaultValue bool = false
		r.includeCanAccessContent = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetSubscribersRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  int64
}

func (r GetSubscribersRequest) Execute() (*GetSubscribersResponse, *http.Response, error) {
	return r.ApiService.GetSubscribersExecute(r)
}

/*
GetSubscribers Get channel subscribers

Get all users subscribed to a channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return GetSubscribersRequest
*/
func (c *simpleClient) GetSubscribers(ctx context.Context, channelId int64) GetSubscribersRequest {
	return GetSubscribersRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return GetSubscribersResponse
func (c *simpleClient) GetSubscribersExecute(r GetSubscribersRequest) (*GetSubscribersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSubscribersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetSubscriptionStatusRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	userId     int64
	channelId  int64
}

func (r GetSubscriptionStatusRequest) Execute() (*GetSubscriptionStatusResponse, *http.Response, error) {
	return r.ApiService.GetSubscriptionStatusExecute(r)
}

/*
GetSubscriptionStatus Get subscription status

Check whether a user is subscribed to a channel.

**Changes**: New in Zulip 3.0 (feature level 12).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The target user's Id.
	@param channelId The Id of the channel to access.
	@return GetSubscriptionStatusRequest
*/
func (c *simpleClient) GetSubscriptionStatus(ctx context.Context, userId int64, channelId int64) GetSubscriptionStatusRequest {
	return GetSubscriptionStatusRequest{
		ApiService: c,
		ctx:        ctx,
		userId:     userId,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return GetSubscriptionStatusResponse
func (c *simpleClient) GetSubscriptionStatusExecute(r GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSubscriptionStatusResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/subscriptions/{stream_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetSubscriptionsRequest struct {
	ctx                context.Context
	ApiService         ChannelsAPI
	includeSubscribers *string
}

// Whether each returned channel object should include a &#x60;subscribers&#x60; field containing a list of the user Ids of its subscribers.  Client apps supporting organizations with many thousands of users should not pass &#x60;true&#x60;, because the full subscriber matrix may be several megabytes of data. The &#x60;partial&#x60; value, combined with the &#x60;subscriber_count&#x60; and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes &#x60;partial&#x60; for this parameter, the server may, for some channels, return a subset of the channel&#39;s subscribers in the &#x60;partial_subscribers&#x60; field instead of the &#x60;subscribers&#x60; field, which always contains the complete set of subscribers.  The server guarantees that it will always return a &#x60;subscribers&#x60; field for channels with fewer than 250 total subscribers. When returning a &#x60;partial_subscribers&#x60; field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  **Changes**: The &#x60;partial&#x60; value is new in Zulip 11.0 (feature level 412).  New in Zulip 2.1.0.
func (r GetSubscriptionsRequest) IncludeSubscribers(includeSubscribers string) GetSubscriptionsRequest {
	r.includeSubscribers = &includeSubscribers
	return r
}

func (r GetSubscriptionsRequest) Execute() (*GetSubscriptionsResponse, *http.Response, error) {
	return r.ApiService.GetSubscriptionsExecute(r)
}

/*
GetSubscriptions Get subscribed channels

Get all channels that the user is subscribed to.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetSubscriptionsRequest
*/
func (c *simpleClient) GetSubscriptions(ctx context.Context) GetSubscriptionsRequest {
	return GetSubscriptionsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSubscriptionsResponse
func (c *simpleClient) GetSubscriptionsExecute(r GetSubscriptionsRequest) (*GetSubscriptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSubscriptionsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeSubscribers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_subscribers", r.includeSubscribers, "form", "")
	} else {
		var defaultValue string = "false"
		r.includeSubscribers = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MuteTopicRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	topic      *string
	op         *string
	channelId  *int64
	channel    *string
}

// The topic to (un)mute. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum topic length.
func (r MuteTopicRequest) Topic(topic string) MuteTopicRequest {
	r.topic = &topic
	return r
}

// Whether to mute (&#x60;add&#x60;) or unmute (&#x60;remove&#x60;) the provided topic.
func (r MuteTopicRequest) Op(op string) MuteTopicRequest {
	r.op = &op
	return r
}

// The Id of the channel to access.  Clients must provide either &#x60;stream&#x60; or &#x60;stream_id&#x60; as a parameter to this endpoint, but not both.  **Changes**: New in Zulip 2.0.0.
func (r MuteTopicRequest) ChannelId(channelId int64) MuteTopicRequest {
	r.channelId = &channelId
	return r
}

// The name of the channel to access.  Clients must provide either &#x60;stream&#x60; or &#x60;stream_id&#x60; as a parameter to this endpoint, but not both. Clients should use &#x60;stream_id&#x60; instead of the &#x60;stream&#x60; parameter when possible.
func (r MuteTopicRequest) Channel(channel string) MuteTopicRequest {
	r.channel = &channel
	return r
}

func (r MuteTopicRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.MuteTopicExecute(r)
}

/*
MuteTopic Topic muting

[Mute or unmute a topic](zulip.com/help/mute-a-topic) within a channel that
the current user is subscribed to.

**Changes**: Deprecated in Zulip 7.0 (feature level 170). Clients connecting
to newer servers should use the [POST /user_topics](zulip.com/api/update-user-topic
endpoint, as this endpoint may be removed in a future release.

Before Zulip 7.0 (feature level 169), this endpoint
returned an error if asked to mute a topic that was already muted
or asked to unmute a topic that had not previously been muted.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MuteTopicRequest

Deprecated
*/
func (c *simpleClient) MuteTopic(ctx context.Context) MuteTopicRequest {
	return MuteTopicRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
//
// Deprecated
func (c *simpleClient) MuteTopicExecute(r MuteTopicRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions/muted_topics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.topic == nil {
		return localVarReturnValue, nil, reportError("topic is required and must be specified")
	}
	if r.op == nil {
		return localVarReturnValue, nil, reportError("op is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.channelId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "stream_id", r.channelId, "form", "")
	}
	if r.channel != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "stream", r.channel, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "op", r.op, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PatchChannelFoldersRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	order      *[]int64
}

// A list of channel folder Ids representing the new order.  This list must include the Ids of all the organization&#39;s channel folders, including archived folders.
func (r PatchChannelFoldersRequest) Order(order []int64) PatchChannelFoldersRequest {
	r.order = &order
	return r
}

func (r PatchChannelFoldersRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.PatchChannelFoldersExecute(r)
}

/*
PatchChannelFolders Reorder channel folders

Given an array of channel folder Ids, this method will set the `order`
property of all of the channel folders in the organization according to
the order of the channel folder Ids specified in the request.

**Changes**: New in Zulip 11.0 (feature level 414).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PatchChannelFoldersRequest
*/
func (c *simpleClient) PatchChannelFolders(ctx context.Context) PatchChannelFoldersRequest {
	return PatchChannelFoldersRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) PatchChannelFoldersExecute(r PatchChannelFoldersRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channel_folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "order", r.order, "form", "multi")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveDefaultChannelRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
	channelId  *int64
}

// The Id of the target channel.
func (r RemoveDefaultChannelRequest) ChannelId(channelId int64) RemoveDefaultChannelRequest {
	r.channelId = &channelId
	return r
}

func (r RemoveDefaultChannelRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveDefaultChannelExecute(r)
}

/*
RemoveDefaultChannel Remove a default channel

Remove a channel from the set of [default channels][default-channels]
for new users joining the organization.

[default-channels]: /help/set-default-channels-for-new-users

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RemoveDefaultChannelRequest
*/
func (c *simpleClient) RemoveDefaultChannel(ctx context.Context) RemoveDefaultChannelRequest {
	return RemoveDefaultChannelRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) RemoveDefaultChannelExecute(r RemoveDefaultChannelRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/default_streams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.channelId == nil {
		return localVarReturnValue, nil, reportError("channelId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "stream_id", r.channelId, "form", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SubscribeRequest struct {
	ctx                               context.Context
	ApiService                        ChannelsAPI
	subscriptions                     *[]SubscriptionRequest
	principals                        *Principals
	authorizationErrorsFatal          *bool
	announce                          *bool
	inviteOnly                        *bool
	isWebPublic                       *bool
	isDefaultChannel                  *bool
	historyPublicToSubscribers        *bool
	messageRetentionDays              *MessageRetentionDays
	topicsPolicy                      *TopicsPolicy
	canAddSubscribersGroup            *GroupSettingValue
	canRemoveSubscribersGroup         *GroupSettingValue
	canAdministerChannelGroup         *GroupSettingValue
	canDeleteAnyMessageGroup          *GroupSettingValue
	canDeleteOwnMessageGroup          *GroupSettingValue
	canMoveMessagesOutOfChannelGroup  *GroupSettingValue
	canMoveMessagesWithinChannelGroup *GroupSettingValue
	canSendMessageGroup               *GroupSettingValue
	canSubscribeGroup                 *GroupSettingValue
	canResolveTopicsGroup             *GroupSettingValue
	folderId                          *int64
	sendNewSubscriptionMessages       *bool
}

// SubscriptionRequest struct for SubscriptionRequest
type SubscriptionRequest struct {
	// The name of the channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`](zulip.com/api/register-queue) endpoint to determine the maximum channel name length.
	Name string `json:"name"`
	// The [description](zulip.com/help/change-the-channel-description) to use for a new channel being created, in text/markdown format.  See the help center article on [message formatting](zulip.com/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.  Clients should use the `max_stream_description_length` returned by the [`POST /register`](zulip.com/api/register-queue) endpoint to determine the maximum channel description length.
	Description *string `json:"description,omitempty"`
}

// A list of dictionaries containing the key &#x60;name&#x60; and value specifying the name of the channel to subscribe. If the channel does not exist a new channel is created. The description of the channel created can be specified by setting the dictionary key &#x60;description&#x60; with an appropriate value.
func (r SubscribeRequest) Subscriptions(subscriptions []SubscriptionRequest) SubscribeRequest {
	r.subscriptions = &subscriptions
	return r
}

func (r SubscribeRequest) Principals(principals Principals) SubscribeRequest {
	r.principals = &principals
	return r
}

// A boolean specifying whether authorization errors (such as when the requesting user is not authorized to access a private channel) should be considered fatal or not. When &#x60;true&#x60;, an authorization error is reported as such. When set to &#x60;false&#x60;, the response will be a 200 and any channels where the request encountered an authorization error will be listed in the &#x60;unauthorized&#x60; key.
func (r SubscribeRequest) AuthorizationErrorsFatal(authorizationErrorsFatal bool) SubscribeRequest {
	r.authorizationErrorsFatal = &authorizationErrorsFatal
	return r
}

// If one of the channels specified did not exist previously and is thus created by this call, this determines whether [notification bot](zulip.com/help/configure-automated-notices) will send an announcement about the new channel&#39;s creation.
func (r SubscribeRequest) Announce(announce bool) SubscribeRequest {
	r.announce = &announce
	return r
}

// As described above, this endpoint will create a new channel if passed a channel name that doesn&#39;t already exist. This parameters and the ones that follow are used to request an initial configuration of a created channel; they are ignored for channels that already exist.  This parameter determines whether any newly created channels will be private channels.
func (r SubscribeRequest) InviteOnly(inviteOnly bool) SubscribeRequest {
	r.inviteOnly = &inviteOnly
	return r
}

// This parameter determines whether any newly created channels will be web-public channels.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current use to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).
func (r SubscribeRequest) IsWebPublic(isWebPublic bool) SubscribeRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// This parameter determines whether any newly created channels will be added as [default channels][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.
func (r SubscribeRequest) IsDefaultChannel(isDefaultChannel bool) SubscribeRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

// Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.
func (r SubscribeRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) SubscribeRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

func (r SubscribeRequest) MessageRetentionDays(messageRetentionDays MessageRetentionDays) SubscribeRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

func (r SubscribeRequest) TopicsPolicy(topicsPolicy TopicsPolicy) SubscribeRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

func (r SubscribeRequest) CanAddSubscribersGroup(canAddSubscribersGroup GroupSettingValue) SubscribeRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

func (r SubscribeRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup GroupSettingValue) SubscribeRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

func (r SubscribeRequest) CanAdministerChannelGroup(canAdministerChannelGroup GroupSettingValue) SubscribeRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

func (r SubscribeRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup GroupSettingValue) SubscribeRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

func (r SubscribeRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup GroupSettingValue) SubscribeRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

func (r SubscribeRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup GroupSettingValue) SubscribeRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

func (r SubscribeRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup GroupSettingValue) SubscribeRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

func (r SubscribeRequest) CanSendMessageGroup(canSendMessageGroup GroupSettingValue) SubscribeRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

func (r SubscribeRequest) CanSubscribeGroup(canSubscribeGroup GroupSettingValue) SubscribeRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

func (r SubscribeRequest) CanResolveTopicsGroup(canResolveTopicsGroup GroupSettingValue) SubscribeRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

// This parameter determines the folder to which the newly created channel will be added.  If the value is &#x60;None&#x60;, the channel will not be added to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).
func (r SubscribeRequest) FolderId(folderId int64) SubscribeRequest {
	r.folderId = &folderId
	return r
}

// Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than &#x60;max_bulk_new_subscription_messages&#x60; (available in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.
func (r SubscribeRequest) SendNewSubscriptionMessages(sendNewSubscriptionMessages bool) SubscribeRequest {
	r.sendNewSubscriptionMessages = &sendNewSubscriptionMessages
	return r
}

func (r SubscribeRequest) Execute() (*SubscribeResponse, *http.Response, error) {
	return r.ApiService.SubscribeExecute(r)
}

/*
Subscribe Subscribe to a channel

Subscribe one or more users to one or more channels.

If any of the specified channels do not exist, they are automatically
created. The initial [channel settings](zulip.com/api/update-stream) will be determined
by the optional parameters, like `invite_only`, detailed below.

Note that the ability to subscribe oneself and/or other users
to a specified channel depends on the [channel's permissions
settings](zulip.com/help/channel-permissions.

**Changes**: Before Zulip 10.0 (feature level 362),
subscriptions in archived channels could not be modified.

Before Zulip 10.0 (feature level 357), the
`can_subscribe_group` permission, which allows members of the
group to subscribe themselves to the channel, did not exist.

Before Zulip 10.0 (feature level 349), a user cannot subscribe
other users to a private channel without being subscribed
to that channel themselves. Now, If a user is part of
`can_add_subscribers_group`, they can subscribe themselves or other
users to a private channel without being subscribed to that channel.

Removed `stream_post_policy` and `is_announcement_only`
parameters in Zulip 10.0 (feature level 333), as permission to post
in the channel is now controlled by `can_send_message_group`.

Before Zulip 8.0 (feature level 208), if a user specified by the
[`principals`][principals-param] parameter was a deactivated user,
or did not exist, then an HTTP status code of 403 was returned with
`code: "UNAUTHORIZED_PRINCIPAL"` in the error response. As of this
feature level, an HTTP status code of 400 is returned with
`code: "BAD_REQUEST"` in the error response for these cases.

[principals-param]: /api/subscribe#parameter-principals

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SubscribeRequest
*/
func (c *simpleClient) Subscribe(ctx context.Context) SubscribeRequest {
	return SubscribeRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeResponse
func (c *simpleClient) SubscribeExecute(r SubscribeRequest) (*SubscribeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubscribeResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptions == nil {
		return localVarReturnValue, nil, reportError("subscriptions is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "subscriptions", r.subscriptions, "form", "multi")
	if r.principals != nil {
		paramJson, err := parameterToJson(*r.principals)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("principals", paramJson)
	}
	if r.authorizationErrorsFatal != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "authorization_errors_fatal", r.authorizationErrorsFatal, "form", "")
	}
	if r.announce != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "announce", r.announce, "form", "")
	}
	if r.inviteOnly != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "invite_only", r.inviteOnly, "form", "")
	}
	if r.isWebPublic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_web_public", r.isWebPublic, "form", "")
	}
	if r.isDefaultChannel != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_default_stream", r.isDefaultChannel, "form", "")
	}
	if r.historyPublicToSubscribers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "history_public_to_subscribers", r.historyPublicToSubscribers, "form", "")
	}
	if r.messageRetentionDays != nil {
		paramJson, err := parameterToJson(*r.messageRetentionDays)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("message_retention_days", paramJson)
	}
	if r.topicsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topics_policy", r.topicsPolicy, "form", "")
	}
	if r.canAddSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canAddSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_add_subscribers_group", paramJson)
	}
	if r.canRemoveSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canRemoveSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_remove_subscribers_group", paramJson)
	}
	if r.canAdministerChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canAdministerChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_administer_channel_group", paramJson)
	}
	if r.canDeleteAnyMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteAnyMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_any_message_group", paramJson)
	}
	if r.canDeleteOwnMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteOwnMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_own_message_group", paramJson)
	}
	if r.canMoveMessagesOutOfChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesOutOfChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_out_of_channel_group", paramJson)
	}
	if r.canMoveMessagesWithinChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesWithinChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_within_channel_group", paramJson)
	}
	if r.canSendMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canSendMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_send_message_group", paramJson)
	}
	if r.canSubscribeGroup != nil {
		paramJson, err := parameterToJson(*r.canSubscribeGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_subscribe_group", paramJson)
	}
	if r.canResolveTopicsGroup != nil {
		paramJson, err := parameterToJson(*r.canResolveTopicsGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_resolve_topics_group", paramJson)
	}
	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "folder_id", r.folderId, "form", "")
	}
	if r.sendNewSubscriptionMessages != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_new_subscription_messages", r.sendNewSubscriptionMessages, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UnsubscribeRequest struct {
	ctx           context.Context
	ApiService    ChannelsAPI
	subscriptions *[]string
	principals    *Principals
}

// A list of channel names to unsubscribe from. This parameter is called &#x60;streams&#x60; in our Python API.
func (r UnsubscribeRequest) Subscriptions(subscriptions []string) UnsubscribeRequest {
	r.subscriptions = &subscriptions
	return r
}

func (r UnsubscribeRequest) Principals(principals Principals) UnsubscribeRequest {
	r.principals = &principals
	return r
}

func (r UnsubscribeRequest) Execute() (*UnsubscribeResponse, *http.Response, error) {
	return r.ApiService.UnsubscribeExecute(r)
}

/*
Unsubscribe Unsubscribe from a channel

Unsubscribe yourself or other users from one or more channels.

In addition to managing the current user's subscriptions, this
endpoint can be used to remove other users from channels. This
is possible in 3 situations:

  - Organization administrators can remove any user from any
    channel.
  - Users can remove a bot that they own from any channel that
    the user [can access](zulip.com/help/channel-permissions.
  - Users can unsubscribe any user from a channel if they [have
    access](zulip.com/help/channel-permissions) to the channel and are a
    member of the [user group](zulip.com/api/get-user-groups) specified
    by the [`can_remove_subscribers_group`][can-remove-parameter]
    for the channel.

**Changes**: Before Zulip 10.0 (feature level 362),
subscriptions in archived channels could not be modified.

Before Zulip 8.0 (feature level 208), if a user specified by
the [`principals`][principals-param] parameter was a
deactivated user, or did not exist, then an HTTP status code
of 403 was returned with `code: "UNAUTHORIZED_PRINCIPAL"` in
the error response. As of this feature level, an HTTP status
code of 400 is returned with `code: "BAD_REQUEST"` in the
error response for these cases.

Before Zulip 8.0 (feature level 197),
the `can_remove_subscribers_group` setting
was named `can_remove_subscribers_group_id`.

Before Zulip 7.0 (feature level 161), the
`can_remove_subscribers_group_id` for all channels was always
the system group for organization administrators.

Before Zulip 6.0 (feature level 145), users had no special
privileges for managing bots that they own.

[principals-param]: /api/unsubscribe#parameter-principals
[can-remove-parameter]: /api/subscribe#parameter-can_remove_subscribers_group

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UnsubscribeRequest
*/
func (c *simpleClient) Unsubscribe(ctx context.Context) UnsubscribeRequest {
	return UnsubscribeRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UnsubscribeResponse
func (c *simpleClient) UnsubscribeExecute(r UnsubscribeRequest) (*UnsubscribeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UnsubscribeResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptions == nil {
		return localVarReturnValue, nil, reportError("subscriptions is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "subscriptions", r.subscriptions, "form", "multi")
	if r.principals != nil {
		paramJson, err := parameterToJson(*r.principals)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("principals", paramJson)
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateChannelFolderRequest struct {
	ctx             context.Context
	ApiService      ChannelsAPI
	channelFolderId int64
	name            *string
	description     *string
	isArchived      *bool
}

// The new name of the channel folder.  Clients should use the &#x60;max_channel_folder_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel folder name length.
func (r UpdateChannelFolderRequest) Name(name string) UpdateChannelFolderRequest {
	r.name = &name
	return r
}

// The new description of the channel folder.  Clients should use the &#x60;max_channel_folder_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel folder description length.
func (r UpdateChannelFolderRequest) Description(description string) UpdateChannelFolderRequest {
	r.description = &description
	return r
}

// Whether to archive or unarchive the channel folder.
func (r UpdateChannelFolderRequest) IsArchived(isArchived bool) UpdateChannelFolderRequest {
	r.isArchived = &isArchived
	return r
}

func (r UpdateChannelFolderRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateChannelFolderExecute(r)
}

/*
UpdateChannelFolder Update a channel folder

Update the name or description of a channel folder.

This endpoint is also used to archive and unarchive
a channel folder.

Only organization administrators can update a
channel folder.

**Changes**: New in Zulip 11.0 (feature level 389).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelFolderId The Id of the target channel folder.
	@return UpdateChannelFolderRequest
*/
func (c *simpleClient) UpdateChannelFolder(ctx context.Context, channelFolderId int64) UpdateChannelFolderRequest {
	return UpdateChannelFolderRequest{
		ApiService:      c,
		ctx:             ctx,
		channelFolderId: channelFolderId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateChannelFolderExecute(r UpdateChannelFolderRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channel_folders/{channel_folder_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel_folder_id"+"}", url.PathEscape(parameterValueToString(r.channelFolderId, "channelFolderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.isArchived != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_archived", r.isArchived, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateChannelRequest struct {
	ctx                               context.Context
	ApiService                        ChannelsAPI
	channelId                         int64
	description                       *string
	newName                           *string
	isPrivate                         *bool
	isWebPublic                       *bool
	historyPublicToSubscribers        *bool
	isDefaultChannel                  *bool
	messageRetentionDays              *MessageRetentionDays
	isArchived                        *bool
	folderId                          *int64
	topicsPolicy                      *TopicsPolicy
	canAddSubscribersGroup            *GroupSettingValueUpdate
	canRemoveSubscribersGroup         *GroupSettingValueUpdate
	canAdministerChannelGroup         *GroupSettingValueUpdate
	canDeleteAnyMessageGroup          *GroupSettingValueUpdate
	canDeleteOwnMessageGroup          *GroupSettingValueUpdate
	canMoveMessagesOutOfChannelGroup  *GroupSettingValueUpdate
	canMoveMessagesWithinChannelGroup *GroupSettingValueUpdate
	canSendMessageGroup               *GroupSettingValueUpdate
	canSubscribeGroup                 *GroupSettingValueUpdate
	canResolveTopicsGroup             *GroupSettingValueUpdate
}

// GroupSettingValueUpdate struct for GroupSettingValueUpdate
type GroupSettingValueUpdate struct {
	New GroupSettingValue  `json:"new"`
	Old *GroupSettingValue `json:"old,omitempty"`
}

// The new [description](zulip.com/help/change-the-channel-description) for the channel, in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown) format.  Clients should use the &#x60;max_stream_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel description length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).
func (r UpdateChannelRequest) Description(description string) UpdateChannelRequest {
	r.description = &description
	return r
}

// The new name for the channel.  Clients should use the &#x60;max_stream_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum channel name length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).
func (r UpdateChannelRequest) NewName(newName string) UpdateChannelRequest {
	r.newName = &newName
	return r
}

// Change whether the channel is a private channel.
func (r UpdateChannelRequest) IsPrivate(isPrivate bool) UpdateChannelRequest {
	r.isPrivate = &isPrivate
	return r
}

// Change whether the channel is a web-public channel.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current use to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).
func (r UpdateChannelRequest) IsWebPublic(isWebPublic bool) UpdateChannelRequest {
	r.isWebPublic = &isWebPublic
	return r
}

// Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  It&#39;s an error for this parameter to be false for a public or web-public channel and when is_private is false.  **Changes**: Before Zulip 6.0 (feature level 136), &#x60;history_public_to_subscribers&#x60; was silently ignored unless the request also contained either &#x60;is_private&#x60; or &#x60;is_web_public&#x60;.
func (r UpdateChannelRequest) HistoryPublicToSubscribers(historyPublicToSubscribers bool) UpdateChannelRequest {
	r.historyPublicToSubscribers = &historyPublicToSubscribers
	return r
}

// Add or remove the channel as a [default channel][default-channel] for new users joining the organization.  [default-channel]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.
func (r UpdateChannelRequest) IsDefaultChannel(isDefaultChannel bool) UpdateChannelRequest {
	r.isDefaultChannel = &isDefaultChannel
	return r
}

func (r UpdateChannelRequest) MessageRetentionDays(messageRetentionDays MessageRetentionDays) UpdateChannelRequest {
	r.messageRetentionDays = &messageRetentionDays
	return r
}

// A boolean indicating whether the channel is [archived](zulip.com/help/archive-a-channel) or unarchived. Currently only allows unarchiving previously archived channels.  **Changes**: New in Zulip 11.0 (feature level 388).
func (r UpdateChannelRequest) IsArchived(isArchived bool) UpdateChannelRequest {
	r.isArchived = &isArchived
	return r
}

// Id of the new folder to which the channel should belong.  It can be &#x60;None&#x60; if the user wants to just remove the channel from its existing folder.  **Changes**: New in Zulip 11.0 (feature level 389).
func (r UpdateChannelRequest) FolderId(folderId int64) UpdateChannelRequest {
	r.folderId = &folderId
	return r
}

func (r UpdateChannelRequest) TopicsPolicy(topicsPolicy TopicsPolicy) UpdateChannelRequest {
	r.topicsPolicy = &topicsPolicy
	return r
}

// The set of users who have permission to add subscribers to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who can administer the channel or have similar realm-level permissions can add subscribers to a public channel regardless of the value of this setting.  Users in this group need not be subscribed to a private channel to add subscribers to it.  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 342). Previously, there was no channel-level setting for this permission.
func (r UpdateChannelRequest) CanAddSubscribersGroup(canAddSubscribersGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canAddSubscribersGroup = &canAddSubscribersGroup
	return r
}

// The set of users who have permission to unsubscribe others from this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can unsubscribe others from a channel as though they were in this group without being explicitly listed here.  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: Prior to Zulip 10.0 (feature level 349), channel administrators could not unsubscribe other users if they were not an organization administrator or part of &#x60;can_remove_subscribers_group&#x60;. Realm administrators were not allowed to unsubscribe other users from a private channel if they were not subscribed to that channel.  Prior to Zulip 10.0 (feature level 320), this value was always the integer Id of a system group.  Before Zulip 8.0 (feature level 197), the &#x60;can_remove_subscribers_group&#x60; setting was named &#x60;can_remove_subscribers_group_id&#x60;.  New in Zulip 7.0 (feature level 161).
func (r UpdateChannelRequest) CanRemoveSubscribersGroup(canRemoveSubscribersGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canRemoveSubscribersGroup = &canRemoveSubscribersGroup
	return r
}

// The set of users who have permission to administer this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can administer every channel as though they were in this group without being explicitly listed here.  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel in order to add other subscribers to the channel.  **Changes**: Prior to Zulip 10.0 (feature level 349) a user needed to [have content access](zulip.com/help/channel-permissions) to a channel in order to modify it. The exception to this rule was that organization administrators can edit channel names and descriptions without having full access to the channel.  New in Zulip 10.0 (feature level 325). Prior to this change, the permission to administer channels was limited to realm administrators.
func (r UpdateChannelRequest) CanAdministerChannelGroup(canAdministerChannelGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canAdministerChannelGroup = &canAdministerChannelGroup
	return r
}

// The set of users who have permission to delete any message in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel in order to delete any message in the channel.  Users present in the organization-level &#x60;can_delete_any_message_group&#x60; setting can always delete any message in the channel if they [have content access](zulip.com/help/channel-permissions) to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in &#x60;can_delete_any_message_group&#x60; were able delete any message in the organization.
func (r UpdateChannelRequest) CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canDeleteAnyMessageGroup = &canDeleteAnyMessageGroup
	return r
}

// The set of users who have permission to delete the messages that they have sent in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel in order to delete their own message in the channel.  Users with permission to delete any message in the channel and users present in the organization-level &#x60;can_delete_own_message_group&#x60; setting can always delete their own messages in the channel if they [have content access](zulip.com/help/channel-permissions) to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in the organization-level &#x60;can_delete_any_message_group&#x60; and &#x60;can_delete_own_message_group&#x60; settings were able delete their own messages in the organization.
func (r UpdateChannelRequest) CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canDeleteOwnMessageGroup = &canDeleteOwnMessageGroup
	return r
}

// The set of users who have permission to move messages out of this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel in order to move messages out of the channel.  Channel administrators and users present in the organization-level &#x60;can_move_messages_between_channels_group&#x60; setting can always move messages out of the channel if they [have content access](zulip.com/help/channel-permissions) to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in &#x60;can_move_messages_between_channels_group&#x60; were able move messages between channels.
func (r UpdateChannelRequest) CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canMoveMessagesOutOfChannelGroup = &canMoveMessagesOutOfChannelGroup
	return r
}

// The set of users who have permission to move messages within this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel in order to move messages within the channel.  Channel administrators and users present in the organization-level &#x60;can_move_messages_between_topics_group&#x60; setting can always move messages within the channel if they [have content access](zulip.com/help/channel-permissions) to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in &#x60;can_move_messages_between_topics_group&#x60; were able move messages between topics of a channel.
func (r UpdateChannelRequest) CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canMoveMessagesWithinChannelGroup = &canMoveMessagesWithinChannelGroup
	return r
}

// The set of users who have permission to post in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 333). Previously &#x60;stream_post_policy&#x60; field used to control the permission to post in the channel.
func (r UpdateChannelRequest) CanSendMessageGroup(canSendMessageGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canSendMessageGroup = &canSendMessageGroup
	return r
}

// The set of users who have permission to subscribe themselves to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Everyone, excluding guests, can subscribe to any public channel irrespective of this setting.  Users in this group can subscribe to a private channel as well.  Note that a user must [have content access](zulip.com/help/channel-permissions) to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 357).
func (r UpdateChannelRequest) CanSubscribeGroup(canSubscribeGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canSubscribeGroup = &canSubscribeGroup
	return r
}

// The set of users who have permission to to resolve topics in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who have similar realm-level permissions can resolve topics in a channel regardless of the value of this setting.  **Changes**: New in Zulip 11.0 (feature level 402).
func (r UpdateChannelRequest) CanResolveTopicsGroup(canResolveTopicsGroup GroupSettingValueUpdate) UpdateChannelRequest {
	r.canResolveTopicsGroup = &canResolveTopicsGroup
	return r
}

func (r UpdateChannelRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateChannelExecute(r)
}

/*
UpdateChannel Update a channel

Configure the channel with the Id `channelId`. This endpoint supports
an organization administrator editing any property of a channel,
including:

  - Channel [name](zulip.com/help/rename-a-channel) and [description](zulip.com/help/change-the-channel-description
  - Channel [permissions](zulip.com/help/channel-permissions, including
    [privacy](zulip.com/help/change-the-privacy-of-a-channel) and [who can
    send](zulip.com/help/channel-posting-policy.

Note that an organization administrator's ability to change a
[private channel's permissions](zulip.com/help/channel-permissions#private-channels
depends on them being subscribed to the channel.

**Changes**: Before Zulip 10.0 (feature level 362), channel privacy could not be
edited for archived channels.

Removed `stream_post_policy` and `is_announcement_only`
parameters in Zulip 10.0 (feature level 333), as permission to post
in the channel is now controlled by `can_send_message_group`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId The Id of the channel to access.
	@return UpdateChannelRequest
*/
func (c *simpleClient) UpdateChannel(ctx context.Context, channelId int64) UpdateChannelRequest {
	return UpdateChannelRequest{
		ApiService: c,
		ctx:        ctx,
		channelId:  channelId,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateChannelExecute(r UpdateChannelRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams/{stream_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"stream_id"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.newName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "new_name", r.newName, "", "")
	}
	if r.isPrivate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_private", r.isPrivate, "form", "")
	}
	if r.isWebPublic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_web_public", r.isWebPublic, "form", "")
	}
	if r.historyPublicToSubscribers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "history_public_to_subscribers", r.historyPublicToSubscribers, "form", "")
	}
	if r.isDefaultChannel != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_default_stream", r.isDefaultChannel, "form", "")
	}
	if r.messageRetentionDays != nil {
		paramJson, err := parameterToJson(*r.messageRetentionDays)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("message_retention_days", paramJson)
	}
	if r.isArchived != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_archived", r.isArchived, "", "")
	}
	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "folder_id", r.folderId, "form", "")
	}
	if r.topicsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topics_policy", r.topicsPolicy, "", "")
	}
	if r.canAddSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canAddSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_add_subscribers_group", paramJson)
	}
	if r.canRemoveSubscribersGroup != nil {
		paramJson, err := parameterToJson(*r.canRemoveSubscribersGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_remove_subscribers_group", paramJson)
	}
	if r.canAdministerChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canAdministerChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_administer_channel_group", paramJson)
	}
	if r.canDeleteAnyMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteAnyMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_any_message_group", paramJson)
	}
	if r.canDeleteOwnMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canDeleteOwnMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_delete_own_message_group", paramJson)
	}
	if r.canMoveMessagesOutOfChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesOutOfChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_out_of_channel_group", paramJson)
	}
	if r.canMoveMessagesWithinChannelGroup != nil {
		paramJson, err := parameterToJson(*r.canMoveMessagesWithinChannelGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_move_messages_within_channel_group", paramJson)
	}
	if r.canSendMessageGroup != nil {
		paramJson, err := parameterToJson(*r.canSendMessageGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_send_message_group", paramJson)
	}
	if r.canSubscribeGroup != nil {
		paramJson, err := parameterToJson(*r.canSubscribeGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_subscribe_group", paramJson)
	}
	if r.canResolveTopicsGroup != nil {
		paramJson, err := parameterToJson(*r.canResolveTopicsGroup)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("can_resolve_topics_group", paramJson)
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateSubscriptionSettingsRequest struct {
	ctx              context.Context
	ApiService       ChannelsAPI
	subscriptionData *[]SubscriptionData
}

// A list of objects that describe the changes that should be applied in each subscription. Each object represents a subscription, and must have a &#x60;stream_id&#x60; key that identifies the channel, as well as the &#x60;property&#x60; being modified and its new &#x60;value&#x60;.
func (r UpdateSubscriptionSettingsRequest) SubscriptionData(subscriptionData []SubscriptionData) UpdateSubscriptionSettingsRequest {
	r.subscriptionData = &subscriptionData
	return r
}

func (r UpdateSubscriptionSettingsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateSubscriptionSettingsExecute(r)
}

/*
UpdateSubscriptionSettings Update subscription settings

This endpoint is used to update the user's personal settings for the
channels they are subscribed to, including muting, color, pinning, and
per-channel notification settings.

**Changes**: Prior to Zulip 5.0 (feature level 111), response
object included the `subscription_data` in the
request. The endpoint now returns the more ergonomic
[`ignored_parameters_unsupported`][ignored-parameters] array instead.

[ignored-parameters]: /api/rest-error-handling#ignored-parameters

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdateSubscriptionSettingsRequest
*/
func (c *simpleClient) UpdateSubscriptionSettings(ctx context.Context) UpdateSubscriptionSettingsRequest {
	return UpdateSubscriptionSettingsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateSubscriptionSettingsExecute(r UpdateSubscriptionSettingsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions/properties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionData == nil {
		return localVarReturnValue, nil, reportError("subscriptionData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "subscription_data", r.subscriptionData, "form", "multi")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateSubscriptionsRequest struct {
	ctx        context.Context
	ApiService ChannelsAPI
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
	return r.ApiService.UpdateSubscriptionsExecute(r)
}

/*
UpdateSubscriptions Update subscriptions

Update which channels you are subscribed to.

**Changes**: Before Zulip 10.0 (feature level 362),
subscriptions in archived channels could not be modified.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdateSubscriptionsRequest
*/
func (c *simpleClient) UpdateSubscriptions(ctx context.Context) UpdateSubscriptionsRequest {
	return UpdateSubscriptionsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UpdateSubscriptionsResponse
func (c *simpleClient) UpdateSubscriptionsExecute(r UpdateSubscriptionsRequest) (*UpdateSubscriptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateSubscriptionsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.delete != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delete", r.delete, "form", "multi")
	}
	if r.add != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "add", r.add, "form", "multi")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateUserTopicRequest struct {
	ctx              context.Context
	ApiService       ChannelsAPI
	channelId        *int64
	topic            *string
	visibilityPolicy *VisibilityPolicy
}

// The Id of the channel to access.
func (r UpdateUserTopicRequest) ChannelId(channelId int64) UpdateUserTopicRequest {
	r.channelId = &channelId
	return r
}

// The topic for which the personal preferences needs to be updated. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue) endpoint to determine the maximum topic length.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](zulip.com/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.
func (r UpdateUserTopicRequest) Topic(topic string) UpdateUserTopicRequest {
	r.topic = &topic
	return r
}

// Controls which visibility policy to set.  - 0 &#x3D; None. Removes the visibility policy previously set for the topic. - 1 &#x3D; Muted. [Mutes the topic](zulip.com/help/mute-a-topic) in a channel. - 2 &#x3D; Unmuted. [Unmutes the topic](zulip.com/help/mute-a-topic) in a muted channel. - 3 &#x3D; Followed. [Follows the topic](zulip.com/help/follow-a-topic.  In an unmuted channel, a topic visibility policy of unmuted will have the same effect as the \\\&quot;None\\\&quot; visibility policy.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.
func (r UpdateUserTopicRequest) VisibilityPolicy(visibilityPolicy VisibilityPolicy) UpdateUserTopicRequest {
	r.visibilityPolicy = &visibilityPolicy
	return r
}

func (r UpdateUserTopicRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateUserTopicExecute(r)
}

/*
UpdateUserTopic Update personal preferences for a topic

This endpoint is used to update the personal preferences for a topic,
such as the topic's visibility policy, which is used to implement
[mute a topic](zulip.com/help/mute-a-topic) and related features.

This endpoint can be used to update the visibility policy for the single
channel and topic pair indicated by the parameters for a user.

**Changes**: New in Zulip 7.0 (feature level 170). Previously,
toggling whether a topic was muted or unmuted was managed by the
[PATCH /users/me/subscriptions/muted_topics](zulip.com/api/mute-topic) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UpdateUserTopicRequest
*/
func (c *simpleClient) UpdateUserTopic(ctx context.Context) UpdateUserTopicRequest {
	return UpdateUserTopicRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) UpdateUserTopicExecute(r UpdateUserTopicRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_topics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.channelId == nil {
		return localVarReturnValue, nil, reportError("channelId is required and must be specified")
	}
	if r.topic == nil {
		return localVarReturnValue, nil, reportError("topic is required and must be specified")
	}
	if r.visibilityPolicy == nil {
		return localVarReturnValue, nil, reportError("visibilityPolicy is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "stream_id", r.channelId, "form", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "visibility_policy", r.visibilityPolicy, "form", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}
