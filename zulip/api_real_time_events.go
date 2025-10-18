package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type RealTimeEventsAPI interface {

	// DeleteQueue Delete an event queue
	//
	// Delete a previously registered queue.
	//
	DeleteQueue(ctx context.Context) DeleteQueueRequest

	// DeleteQueueExecute executes the request
	DeleteQueueExecute(r DeleteQueueRequest) (*Response, *http.Response, error)

	// GetEvents Get events from an event queue
	//
	// This endpoint allows you to receive new events from
	// [a registered event queue].
	//
	// Long-lived clients should use the
	// `event_queue_longpoll_timeout_seconds` property returned by
	// `POST /register` as the client-side HTTP request timeout for
	// calls to this endpoint. It is guaranteed to be higher than
	// heartbeat frequency and should be respected by clients to
	// avoid breaking when heartbeat frequency increases.
	//
	// [a registered event queue]: https://zulip.com/api/register-queue
	GetEvents(ctx context.Context) GetEventsRequest

	// GetEventsExecute executes the request
	GetEventsExecute(r GetEventsRequest) (*GetEventsResponse, *http.Response, error)

	// RegisterQueue Register an event queue
	//
	// This powerful endpoint can be used to register a Zulip "event queue"
	// (subscribed to certain types of "events", or updates to the messages
	// and other Zulip data the current user has access to), as well as to
	// fetch the current state of that data.
	//
	// (`register` also powers the `call_on_each_event` Python API, and is
	// intended primarily for complex applications for which the more convenient
	// `call_on_each_event` API is insufficient).
	//
	// This endpoint returns a `queue_id` and a `last_event_id`; these can be
	// used in subsequent calls to the
	// ["events" endpoint] to request events from
	// the Zulip server using long-polling.
	//
	// The server will queue events for up to 10 minutes of inactivity.
	// After 10 minutes, your event queue will be garbage-collected. The
	// server will send `heartbeat` events every minute, which makes it easy
	// to implement a robust client that does not miss events unless the
	// client loses network connectivity with the Zulip server for 10 minutes
	// or longer.
	//
	// Once the server garbage-collects your event queue, the server will
	// [return an error](https://zulip.com/api/get-events#bad_event_queue_id-errors
	// with a code of `BAD_EVENT_QUEUE_Id` if you try to fetch events from
	// the event queue. Your software will need to handle that error
	// condition by re-initializing itself (e.g. this is what triggers your
	// browser reloading the Zulip web app when your laptop comes back online
	// after being offline for more than 10 minutes).
	//
	// When prototyping with this API, we recommend first calling `register`
	// with no `event_types` parameter to see all the available data from all
	// supported event types. Before using your client in production, you
	// should set appropriate `event_types` and `fetch_event_types` filters
	// so that your client only requests the data it needs. A few minutes
	// doing this often saves 90% of the total bandwidth and other resources
	// consumed by a client using this API.
	//
	// See the [events system developer documentation]
	// if you need deeper details about how the Zulip event queue system
	// works, avoids clients needing to worry about large classes of
	// potentially messy races, etc.
	//
	// *Changes**: Removed `dense_mode` setting in Zulip 10.0 (feature level 364)
	// as we now have `web_font_size_px` and `web_line_height_percent`
	// settings for more control.
	//
	// Before Zulip 7.0 (feature level 183), the
	// `realm_community_topic_editing_limit_seconds` property
	// was returned by the response. It was removed because it
	// had not been in use since the realm setting
	// `move_messages_within_stream_limit_seconds` was introduced
	// in feature level 162.
	//
	// In Zulip 7.0 (feature level 163), the realm setting
	// `email_address_visibility` was removed. It was replaced by a [user setting](https://zulip.com/api/update-settings#parameter-email_address_visibility with
	// a [realm user default], with the encoding of different
	// values preserved. Clients can support all versions by supporting the
	// current API and treating every user as having the realm's
	// `email_address_visibility` value.
	//
	// [realm user default]: https://zulip.com/api/update-realm-user-settings-defaults#parameter-email_address_visibility
	// [events system developer documentation]: https://zulip.readthedocs.io/en/latest/subsystems/events-system.html
	//
	// ["events" endpoint]: https://zulip.com/api/get-events
	RegisterQueue(ctx context.Context) RegisterQueueRequest

	// RegisterQueueExecute executes the request
	RegisterQueueExecute(r RegisterQueueRequest) (*RegisterQueueResponse, *http.Response, error)
}

type DeleteQueueRequest struct {
	ctx        context.Context
	ApiService RealTimeEventsAPI
	queueId    *string
}

// The Id of an event queue that was previously registered via `POST /api/v1/register` (see [Register a queue].
//
// [Register a queue]: https://zulip.com/api/register-queue
func (r DeleteQueueRequest) QueueId(queueId string) DeleteQueueRequest {
	r.queueId = &queueId
	return r
}

func (r DeleteQueueRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteQueueExecute(r)
}

// DeleteQueue Delete an event queue
//
// Delete a previously registered queue.
func (c *simpleClient) DeleteQueue(ctx context.Context) DeleteQueueRequest {
	return DeleteQueueRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) DeleteQueueExecute(r DeleteQueueRequest) (*Response, *http.Response, error) {
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

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.queueId == nil {
		return localVarReturnValue, nil, reportError("queueId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "queue_id", r.queueId, "", "")
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

type GetEventsRequest struct {
	ctx         context.Context
	ApiService  RealTimeEventsAPI
	queueId     *string
	lastEventId *int64
	dontBlock   *bool
}

// The Id of an event queue that was previously registered via `POST /api/v1/register` (see [Register a queue].
//
// [Register a queue]: https://zulip.com/api/register-queue
func (r GetEventsRequest) QueueId(queueId string) GetEventsRequest {
	r.queueId = &queueId
	return r
}

// The highest event Id in this queue that you&#39;ve received and wish to acknowledge. See the [code for `call_on_each_event`] in the [zulip Python module] for an example implementation of correctly processing each event exactly once.
//
// [code for `call_on_each_event`]: https://github.com/zulip/python-zulip-api/blob/main/zulip/zulip/__init__.py
// [zulip Python module]: https://github.com/zulip/python-zulip-api
func (r GetEventsRequest) LastEventId(lastEventId int64) GetEventsRequest {
	r.lastEventId = &lastEventId
	return r
}

// Set to `true` if the client is requesting a nonblocking reply. If not specified, the request will block until either a new event is available or a few minutes have passed, in which case the server will send the client a heartbeat event.
func (r GetEventsRequest) DontBlock(dontBlock bool) GetEventsRequest {
	r.dontBlock = &dontBlock
	return r
}

func (r GetEventsRequest) Execute() (*GetEventsResponse, *http.Response, error) {
	return r.ApiService.GetEventsExecute(r)
}

// GetEvents Get events from an event queue
//
// This endpoint allows you to receive new events from
// [a registered event queue].
//
// Long-lived clients should use the
// `event_queue_longpoll_timeout_seconds` property returned by
// `POST /register` as the client-side HTTP request timeout for
// calls to this endpoint. It is guaranteed to be higher than
// heartbeat frequency and should be respected by clients to
// avoid breaking when heartbeat frequency increases.
//
// [a registered event queue]: https://zulip.com/api/register-queue
func (c *simpleClient) GetEvents(ctx context.Context) GetEventsRequest {
	return GetEventsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetEventsExecute(r GetEventsRequest) (*GetEventsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetEventsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.queueId == nil {
		return localVarReturnValue, nil, reportError("queueId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "queue_id", r.queueId, "form", "")
	if r.lastEventId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_event_id", r.lastEventId, "form", "")
	}
	if r.dontBlock != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dont_block", r.dontBlock, "form", "")
	} else {
		var defaultValue bool = false
		r.dontBlock = &defaultValue
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

type RegisterQueueRequest struct {
	ctx                      context.Context
	ApiService               RealTimeEventsAPI
	applyMarkdown            *bool
	clientGravatar           *bool
	includeSubscribers       *string
	slimPresence             *bool
	presenceHistoryLimitDays *int32
	eventTypes               *[]string
	allPublicChannels        *bool
	clientCapabilities       *map[string]interface{}
	fetchEventTypes          *[]string
	narrow                   *Narrow
}

// Set to `true` if you would like the content to be rendered in HTML format (otherwise the API will return the raw text that the user entered)
func (r RegisterQueueRequest) ApplyMarkdown(applyMarkdown bool) RegisterQueueRequest {
	r.applyMarkdown = &applyMarkdown
	return r
}

// Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  The default value is `true` for authenticated requests and `false` for [unauthenticated requests]. Passing `true` in an unauthenticated request is an error.  **Changes**: Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were `false` in unauthenticated requests.
//
// [unauthenticated requests]: https://zulip.com/help/public-access-option
func (r RegisterQueueRequest) ClientGravatar(clientGravatar bool) RegisterQueueRequest {
	r.clientGravatar = &clientGravatar
	return r
}

// Whether each returned channel object should include a `subscribers` field containing a list of the user Ids of its subscribers.  Client apps supporting organizations with many thousands of users should not pass `true`, because the full subscriber matrix may be several megabytes of data. The `partial` value, combined with the `subscriber_count` and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes `partial` for this parameter, the server may, for some channels, return a subset of the channel&#39;s subscribers in the `partial_subscribers` field instead of the `subscribers` field, which always contains the complete set of subscribers.  The server guarantees that it will always return a `subscribers` field for channels with fewer than 250 total subscribers. When returning a `partial_subscribers` field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  Passing `true` in an [unauthenticated request] is an error.  **Changes**: The `partial` value is new in Zulip 11.0 (feature level 412).  Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were `false` in unauthenticated requests.  New in Zulip 2.1.0.
//
// [unauthenticated request]: https://zulip.com/help/public-access-option
func (r RegisterQueueRequest) IncludeSubscribers(includeSubscribers string) RegisterQueueRequest {
	r.includeSubscribers = &includeSubscribers
	return r
}

// If `true`, the `presences` object returned in the response will be keyed by user Id and the entry for each user&#39;s presence data will be in the modern format.  **Changes**: New in Zulip 3.0 (no feature level; API unstable).
func (r RegisterQueueRequest) SlimPresence(slimPresence bool) RegisterQueueRequest {
	r.slimPresence = &slimPresence
	return r
}

// Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  **Changes**: New in Zulip 10.0 (feature level 288).
func (r RegisterQueueRequest) PresenceHistoryLimitDays(presenceHistoryLimitDays int32) RegisterQueueRequest {
	r.presenceHistoryLimitDays = &presenceHistoryLimitDays
	return r
}

// A JSON-encoded array indicating which types of events you&#39;re interested in. Values that you might find useful include:  - **message** (messages) - **subscription** (changes in your subscriptions) - **realm_user** (changes to users in the organization and   their properties, such as their name).  If you do not specify this parameter, you will receive all events, and have to filter out the events not relevant to your client in your client code. For most applications, one is only interested in messages, so one specifies: `\\"event_types\\": [\\"message\\"]`  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.
func (r RegisterQueueRequest) EventTypes(eventTypes []string) RegisterQueueRequest {
	r.eventTypes = &eventTypes
	return r
}

// Whether you would like to request message events from all public channels. Useful for workflow bots that you&#39;d like to see all new messages sent to public channels. (You can also subscribe the user to private channels).
func (r RegisterQueueRequest) AllPublicChannels(allPublicChannels bool) RegisterQueueRequest {
	r.allPublicChannels = &allPublicChannels
	return r
}

// Dictionary containing details on features the client supports that are relevant to the format of responses sent by the server.  - `notification_settings_null`: Boolean for whether the   client can handle the current API with `null` values for   channel-level notification settings (which means the channel   is not customized and should inherit the user&#39;s global   notification settings for channel messages).
//
//	**Changes**: New in Zulip 2.1.0. In earlier Zulip releases,   channel-level notification settings were simple booleans.  - `bulk_message_deletion`: Boolean for whether the client&#39;s   handler for the `delete_message` event type has been   updated to process the new bulk format (with a   `message_ids`, rather than a singleton `message_id`).   Otherwise, the server will send `delete_message` events   in a loop.
//	**Changes**: New in Zulip 3.0 (feature level 13). This   capability is for backwards-compatibility; it will be   required in a future server release.  - `user_avatar_url_field_optional`: Boolean for whether the   client required avatar URLs for all users, or supports   using `GET /avatar/{user_id}` to access user avatars. If the   client has this capability, the server may skip sending a   `avatar_url` field in the `realm_user` at its sole discretion   to optimize network performance. This is an important optimization   in organizations with 10,000s of users.
//	**Changes**: New in Zulip 3.0 (feature level 18).  - `stream_typing_notifications`: Boolean for whether the client   supports channel typing notifications.
//	**Changes**: New in Zulip 4.0 (feature level 58). This capability is   for backwards-compatibility; it will be required in a   future server release.  - `user_settings_object`: Boolean for whether the client supports the modern   `user_settings` event type. If false, the server will additionally send the   legacy `update_global_notifications` and `update_display_settings` event   types.
//	**Changes**: New in Zulip 5.0 (feature level 89). This capability is for   backwards-compatibility; it will be removed in a future server release.   Because the feature level 89 API changes were merged together, clients can   safely make a request with this client capability and also request all three   event types (`user_settings`, `update_display_settings`,   `update_global_notifications`), and get exactly one copy of settings data on   any server version. Clients can then use the `zulip_feature_level` in the   `/register` response or the presence/absence of a `user_settings` key to   determine where to look for the data.  - `linkifier_url_template`: Boolean for whether the client accepts   [linkifiers] that use [RFC 6570] compliant   URL templates for linkifying matches. If false or unset, then the   `realm_linkifiers` array in the `/register` response will be empty   if present, and no `realm_linkifiers` [events] will be sent to the client.
//	**Changes**: New in Zulip 7.0 (feature level 176). This capability   is for backwards-compatibility.  - `user_list_incomplete`: Boolean for whether the client supports not having an   incomplete user database. If true, then the `realm_users` array in the `register`   response will not include data for inaccessible users and clients of guest users will   not receive `realm_user op:add` events for newly created users that are not accessible   to the current user.
//	**Changes**: New in Zulip 8.0 (feature level 232). This   capability is for backwards-compatibility.  - `include_deactivated_groups`: Boolean for whether the client can handle   deactivated user groups by themselves. If false, then the `realm_user_groups`   array in the `/register` response will only include active groups, clients   will receive a `remove` event instead of `update` event when a group is   deactivated and no `update` event will be sent to the client if a deactivated   user group is renamed.
//	**Changes**: New in Zulip 10.0 (feature level 294). This   capability is for backwards-compatibility.  - `archived_channels`: Boolean for whether the client supports processing   [archived channels] in the `stream` and   `subscription` event types. If `false`, the server will not include data   related to archived channels in the `register` response or in events.
//	**Changes**: New in Zulip 10.0 (feature level 315). This allows clients to   access archived channels, without breaking backwards-compatibility for   existing clients.  - `empty_topic_name`: Boolean for whether the client supports processing   the empty string as a topic name. Clients not declaring this capability   will be sent the value of `realm_empty_topic_display_name` found in the   [POST /register] response instead of the empty string   wherever topic names appear in the register response or events involving   topic names.   &lt;br/&gt;   **Changes**: New in Zulip 10.0 (feature level 334). Previously,   the empty string was not a valid topic name.  - `simplified_presence_events`: Boolean for whether the client supports   receiving the [`presence` event type] with   user presence data in the modern format. If true, the server will   send these events with the `presences` field that has the user presence   data in the modern format. Otherwise, these event will contain fields   with legacy format user presence data.
//	**Changes**: New in Zulip 11.0 (feature level 419).
//
// [events]: https://zulip.com/api/get-events#realm_linkifiers
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
// [archived channels]: https://zulip.com/help/archive-a-channel
// [POST /register]: https://zulip.com/api/register-queue
// [`presence` event type]: https://zulip.com/api/get-events#presence
func (r RegisterQueueRequest) ClientCapabilities(clientCapabilities map[string]interface{}) RegisterQueueRequest {
	r.clientCapabilities = &clientCapabilities
	return r
}

// Same as the `event_types` parameter except that the values in `fetch_event_types` are used to fetch initial data. If `fetch_event_types` is not provided, `event_types` is used and if `event_types` is not provided, this parameter defaults to `null`.  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.
func (r RegisterQueueRequest) FetchEventTypes(fetchEventTypes []string) RegisterQueueRequest {
	r.fetchEventTypes = &fetchEventTypes
	return r
}

// A JSON-encoded array of arrays of length 2 indicating the [narrow filter(s)] for which you&#39;d like to receive events for.  For example, to receive events for direct messages (including group direct messages) received by the user, one can use `\\"narrow\\": [[\\"is\\", \\"dm\\"]]`.  Unlike the API for [fetching messages], this narrow parameter is simply a filter on messages that the user receives through their channel subscriptions (or because they are a recipient of a direct message).  This means that a client that requests a `narrow` filter of `[[\\"channel\\", \\"Denmark\\"]]` will receive events for new messages sent to that channel while the user is subscribed to that channel. The client will not receive any message events at all if the user is not subscribed to `\\"Denmark\\"`.  Newly created bot users are not usually subscribed to any channels, so bots using this API need to be [subscribed] to any channels whose messages you&#39;d like them to process using this endpoint.  See the `all_public_streams` parameter for how to process all public channel messages in an organization.  **Changes**: See [changes section] of search/narrow filter documentation.
//
// [narrow filter(s)]: https://zulip.com/api/construct-narrow
// [fetching messages]: https://zulip.com/api/get-messages
// [subscribed]: https://zulip.com/api/subscribe
// [changes section]: https://zulip.com/api/construct-narrow#changes
func (r RegisterQueueRequest) Narrow(narrow *Narrow) RegisterQueueRequest {
	r.narrow = narrow
	return r
}

func (r RegisterQueueRequest) Execute() (*RegisterQueueResponse, *http.Response, error) {
	return r.ApiService.RegisterQueueExecute(r)
}

// RegisterQueue Register an event queue
//
// This powerful endpoint can be used to register a Zulip "event queue"
// (subscribed to certain types of "events", or updates to the messages
// and other Zulip data the current user has access to), as well as to
// fetch the current state of that data.
//
// (`register` also powers the `call_on_each_event` Python API, and is
// intended primarily for complex applications for which the more convenient
// `call_on_each_event` API is insufficient).
//
// This endpoint returns a `queue_id` and a `last_event_id`; these can be
// used in subsequent calls to the
// ["events" endpoint] to request events from
// the Zulip server using long-polling.
//
// The server will queue events for up to 10 minutes of inactivity.
// After 10 minutes, your event queue will be garbage-collected. The
// server will send `heartbeat` events every minute, which makes it easy
// to implement a robust client that does not miss events unless the
// client loses network connectivity with the Zulip server for 10 minutes
// or longer.
//
// Once the server garbage-collects your event queue, the server will
// [return an error](https://zulip.com/api/get-events#bad_event_queue_id-errors
// with a code of `BAD_EVENT_QUEUE_Id` if you try to fetch events from
// the event queue. Your software will need to handle that error
// condition by re-initializing itself (e.g. this is what triggers your
// browser reloading the Zulip web app when your laptop comes back online
// after being offline for more than 10 minutes).
//
// When prototyping with this API, we recommend first calling `register`
// with no `event_types` parameter to see all the available data from all
// supported event types. Before using your client in production, you
// should set appropriate `event_types` and `fetch_event_types` filters
// so that your client only requests the data it needs. A few minutes
// doing this often saves 90% of the total bandwidth and other resources
// consumed by a client using this API.
//
// See the [events system developer documentation]
// if you need deeper details about how the Zulip event queue system
// works, avoids clients needing to worry about large classes of
// potentially messy races, etc.
//
// *Changes**: Removed `dense_mode` setting in Zulip 10.0 (feature level 364)
// as we now have `web_font_size_px` and `web_line_height_percent`
// settings for more control.
//
// Before Zulip 7.0 (feature level 183), the
// `realm_community_topic_editing_limit_seconds` property
// was returned by the response. It was removed because it
// had not been in use since the realm setting
// `move_messages_within_stream_limit_seconds` was introduced
// in feature level 162.
//
// In Zulip 7.0 (feature level 163), the realm setting
// `email_address_visibility` was removed. It was replaced by a [user setting](https://zulip.com/api/update-settings#parameter-email_address_visibility with
// a [realm user default], with the encoding of different
// values preserved. Clients can support all versions by supporting the
// current API and treating every user as having the realm's
// `email_address_visibility` value.
//
// [realm user default]: https://zulip.com/api/update-realm-user-settings-defaults#parameter-email_address_visibility
// [events system developer documentation]: https://zulip.readthedocs.io/en/latest/subsystems/events-system.html
// ["events" endpoint]: https://zulip.com/api/get-events
func (c *simpleClient) RegisterQueue(ctx context.Context) RegisterQueueRequest {
	return RegisterQueueRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) RegisterQueueExecute(r RegisterQueueRequest) (*RegisterQueueResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RegisterQueueResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/register"

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
	if r.applyMarkdown != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "apply_markdown", r.applyMarkdown, "form", "")
	}
	if r.clientGravatar != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_gravatar", r.clientGravatar, "form", "")
	}
	if r.includeSubscribers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "include_subscribers", r.includeSubscribers, "", "")
	}
	if r.slimPresence != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "slim_presence", r.slimPresence, "form", "")
	}
	if r.presenceHistoryLimitDays != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "presence_history_limit_days", r.presenceHistoryLimitDays, "", "")
	}
	if r.eventTypes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "event_types", r.eventTypes, "form", "multi")
	}
	if r.allPublicChannels != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "all_public_streams", r.allPublicChannels, "form", "")
	}
	if r.clientCapabilities != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_capabilities", r.clientCapabilities, "form", "")
	}
	if r.fetchEventTypes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fetch_event_types", r.fetchEventTypes, "form", "multi")
	}
	if r.narrow != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "narrow", r.narrow, "form", "multi")
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
