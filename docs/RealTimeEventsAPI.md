# \RealTimeEventsAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteQueue**](RealTimeEventsAPI.md#DeleteQueue) | **Delete** /events | Delete an event queue
[**GetEvents**](RealTimeEventsAPI.md#GetEvents) | **Get** /events | Get events from an event queue
[**RealTimePost**](RealTimeEventsAPI.md#RealTimePost) | **Post** /real-time | 
[**RegisterQueue**](RealTimeEventsAPI.md#RegisterQueue) | **Post** /register | Register an event queue
[**RestErrorHandling**](RealTimeEventsAPI.md#RestErrorHandling) | **Post** /rest-error-handling | Error handling



## DeleteQueue

> JsonSuccess DeleteQueue(ctx).QueueId(queueId).Execute()

Delete an event queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	queueId := "queueId_example" // string | The ID of an event queue that was previously registered via `POST /api/v1/register` (see [Register a queue](zulip.com/api/register-queue). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteQueue(context.Background()).QueueId(queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealTimeEventsAPI.DeleteQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteQueue`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `RealTimeEventsAPI.DeleteQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queueId** | **string** | The ID of an event queue that was previously registered via &#x60;POST /api/v1/register&#x60; (see [Register a queue](zulip.com/api/register-queue).  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> GetEvents200Response GetEvents(ctx).QueueId(queueId).LastEventId(lastEventId).DontBlock(dontBlock).Execute()

Get events from an event queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	queueId := "fb67bf8a-c031-47cc-84cf-ed80accacda8" // string | The ID of an event queue that was previously registered via `POST /api/v1/register` (see [Register a queue](zulip.com/api/register-queue). 
	lastEventId := int32(-1) // int32 | The highest event ID in this queue that you've received and wish to acknowledge. See the [code for `call_on_each_event`](https://github.com/zulip/python-zulip-api/blob/main/zulip/zulip/__init__.py) in the [zulip Python module](https://github.com/zulip/python-zulip-api) for an example implementation of correctly processing each event exactly once.  (optional)
	dontBlock := true // bool | Set to `true` if the client is requesting a nonblocking reply. If not specified, the request will block until either a new event is available or a few minutes have passed, in which case the server will send the client a heartbeat event.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetEvents(context.Background()).QueueId(queueId).LastEventId(lastEventId).DontBlock(dontBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealTimeEventsAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: GetEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `RealTimeEventsAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queueId** | **string** | The ID of an event queue that was previously registered via &#x60;POST /api/v1/register&#x60; (see [Register a queue](zulip.com/api/register-queue).  | 
 **lastEventId** | **int32** | The highest event ID in this queue that you&#39;ve received and wish to acknowledge. See the [code for &#x60;call_on_each_event&#x60;](https://github.com/zulip/python-zulip-api/blob/main/zulip/zulip/__init__.py) in the [zulip Python module](https://github.com/zulip/python-zulip-api) for an example implementation of correctly processing each event exactly once.  | 
 **dontBlock** | **bool** | Set to &#x60;true&#x60; if the client is requesting a nonblocking reply. If not specified, the request will block until either a new event is available or a few minutes have passed, in which case the server will send the client a heartbeat event.  | [default to false]

### Return type

[**GetEvents200Response**](GetEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RealTimePost

> RealTimePost(ctx).EventTypes(eventTypes).Narrow(narrow).AllPublicStreams(allPublicStreams).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	eventTypes := []string{"Inner_example"} // []string | A JSON-encoded array indicating which types of events you're interested in. Values that you might find useful include:  - **message** (messages) - **subscription** (changes in your subscriptions) - **realm_user** (changes to users in the organization and   their properties, such as their name).  If you do not specify this parameter, you will receive all events, and have to filter out the events not relevant to your client in your client code. For most applications, one is only interested in messages, so one specifies: `\\\"event_types\\\": [\\\"message\\\"]`  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  (optional)
	narrow := [][]string{[]string{"Inner_example"}} // [][]string | A JSON-encoded array of arrays of length 2 indicating the [narrow filter(s)](zulip.com/api/construct-narrow for which you'd like to receive events for.  For example, to receive events for direct messages (including group direct messages) received by the user, one can use `\\\"narrow\\\": [[\\\"is\\\", \\\"dm\\\"]]`.  Unlike the API for [fetching messages](zulip.com/api/get-messages, this narrow parameter is simply a filter on messages that the user receives through their channel subscriptions (or because they are a recipient of a direct message).  This means that a client that requests a `narrow` filter of `[[\\\"channel\\\", \\\"Denmark\\\"]]` will receive events for new messages sent to that channel while the user is subscribed to that channel. The client will not receive any message events at all if the user is not subscribed to `\\\"Denmark\\\"`.  Newly created bot users are not usually subscribed to any channels, so bots using this API need to be [subscribed](zulip.com/api/subscribe to any channels whose messages you'd like them to process using this endpoint.  See the `all_public_streams` parameter for how to process all public channel messages in an organization.  **Changes**: See [changes section](zulip.com/api/construct-narrow#changes of search/narrow filter documentation.  (optional)
	allPublicStreams := true // bool | Whether you would like to request message events from all public channels. Useful for workflow bots that you'd like to see all new messages sent to public channels. (You can also subscribe the user to private channels).  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealTimePost(context.Background()).EventTypes(eventTypes).Narrow(narrow).AllPublicStreams(allPublicStreams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealTimeEventsAPI.RealTimePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRealTimePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventTypes** | **[]string** | A JSON-encoded array indicating which types of events you&#39;re interested in. Values that you might find useful include:  - **message** (messages) - **subscription** (changes in your subscriptions) - **realm_user** (changes to users in the organization and   their properties, such as their name).  If you do not specify this parameter, you will receive all events, and have to filter out the events not relevant to your client in your client code. For most applications, one is only interested in messages, so one specifies: &#x60;\\\&quot;event_types\\\&quot;: [\\\&quot;message\\\&quot;]&#x60;  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  | 
 **narrow** | **[][]string** | A JSON-encoded array of arrays of length 2 indicating the [narrow filter(s)](zulip.com/api/construct-narrow for which you&#39;d like to receive events for.  For example, to receive events for direct messages (including group direct messages) received by the user, one can use &#x60;\\\&quot;narrow\\\&quot;: [[\\\&quot;is\\\&quot;, \\\&quot;dm\\\&quot;]]&#x60;.  Unlike the API for [fetching messages](zulip.com/api/get-messages, this narrow parameter is simply a filter on messages that the user receives through their channel subscriptions (or because they are a recipient of a direct message).  This means that a client that requests a &#x60;narrow&#x60; filter of &#x60;[[\\\&quot;channel\\\&quot;, \\\&quot;Denmark\\\&quot;]]&#x60; will receive events for new messages sent to that channel while the user is subscribed to that channel. The client will not receive any message events at all if the user is not subscribed to &#x60;\\\&quot;Denmark\\\&quot;&#x60;.  Newly created bot users are not usually subscribed to any channels, so bots using this API need to be [subscribed](zulip.com/api/subscribe to any channels whose messages you&#39;d like them to process using this endpoint.  See the &#x60;all_public_streams&#x60; parameter for how to process all public channel messages in an organization.  **Changes**: See [changes section](zulip.com/api/construct-narrow#changes of search/narrow filter documentation.  | 
 **allPublicStreams** | **bool** | Whether you would like to request message events from all public channels. Useful for workflow bots that you&#39;d like to see all new messages sent to public channels. (You can also subscribe the user to private channels).  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterQueue

> RegisterQueue200Response RegisterQueue(ctx).ApplyMarkdown(applyMarkdown).ClientGravatar(clientGravatar).IncludeSubscribers(includeSubscribers).SlimPresence(slimPresence).PresenceHistoryLimitDays(presenceHistoryLimitDays).EventTypes(eventTypes).AllPublicStreams(allPublicStreams).ClientCapabilities(clientCapabilities).FetchEventTypes(fetchEventTypes).Narrow(narrow).Execute()

Register an event queue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	applyMarkdown := true // bool | Set to `true` if you would like the content to be rendered in HTML format (otherwise the API will return the raw text that the user entered)  (optional) (default to false)
	clientGravatar := true // bool | Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  The default value is `true` for authenticated requests and `false` for [unauthenticated requests](zulip.com/help/public-access-option. Passing `true` in an unauthenticated request is an error.  **Changes**: Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were `false` in unauthenticated requests.  (optional)
	includeSubscribers := "includeSubscribers_example" // string | Whether each returned channel object should include a `subscribers` field containing a list of the user IDs of its subscribers.  Client apps supporting organizations with many thousands of users should not pass `true`, because the full subscriber matrix may be several megabytes of data. The `partial` value, combined with the `subscriber_count` and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes `partial` for this parameter, the server may, for some channels, return a subset of the channel's subscribers in the `partial_subscribers` field instead of the `subscribers` field, which always contains the complete set of subscribers.  The server guarantees that it will always return a `subscribers` field for channels with fewer than 250 total subscribers. When returning a `partial_subscribers` field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  Passing `true` in an [unauthenticated request](zulip.com/help/public-access-option is an error.  **Changes**: The `partial` value is new in Zulip 11.0 (feature level 412).  Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were `false` in unauthenticated requests.  New in Zulip 2.1.0.  (optional) (default to "false")
	slimPresence := true // bool | If `true`, the `presences` object returned in the response will be keyed by user ID and the entry for each user's presence data will be in the modern format.  **Changes**: New in Zulip 3.0 (no feature level; API unstable).  (optional) (default to false)
	presenceHistoryLimitDays := int32(56) // int32 | Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  **Changes**: New in Zulip 10.0 (feature level 288).  (optional)
	eventTypes := []string{"Inner_example"} // []string | A JSON-encoded array indicating which types of events you're interested in. Values that you might find useful include:  - **message** (messages) - **subscription** (changes in your subscriptions) - **realm_user** (changes to users in the organization and   their properties, such as their name).  If you do not specify this parameter, you will receive all events, and have to filter out the events not relevant to your client in your client code. For most applications, one is only interested in messages, so one specifies: `\\\"event_types\\\": [\\\"message\\\"]`  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  (optional)
	allPublicStreams := true // bool | Whether you would like to request message events from all public channels. Useful for workflow bots that you'd like to see all new messages sent to public channels. (You can also subscribe the user to private channels).  (optional) (default to false)
	clientCapabilities := map[string]interface{}{ ... } // map[string]interface{} | Dictionary containing details on features the client supports that are relevant to the format of responses sent by the server.  - `notification_settings_null`: Boolean for whether the   client can handle the current API with `null` values for   channel-level notification settings (which means the channel   is not customized and should inherit the user's global   notification settings for channel messages).   <br />   **Changes**: New in Zulip 2.1.0. In earlier Zulip releases,   channel-level notification settings were simple booleans.  - `bulk_message_deletion`: Boolean for whether the client's   handler for the `delete_message` event type has been   updated to process the new bulk format (with a   `message_ids`, rather than a singleton `message_id`).   Otherwise, the server will send `delete_message` events   in a loop.   <br />   **Changes**: New in Zulip 3.0 (feature level 13). This   capability is for backwards-compatibility; it will be   required in a future server release.  - `user_avatar_url_field_optional`: Boolean for whether the   client required avatar URLs for all users, or supports   using `GET /avatar/{user_id}` to access user avatars. If the   client has this capability, the server may skip sending a   `avatar_url` field in the `realm_user` at its sole discretion   to optimize network performance. This is an important optimization   in organizations with 10,000s of users.   <br />   **Changes**: New in Zulip 3.0 (feature level 18).  - `stream_typing_notifications`: Boolean for whether the client   supports channel typing notifications.   <br />   **Changes**: New in Zulip 4.0 (feature level 58). This capability is   for backwards-compatibility; it will be required in a   future server release.  - `user_settings_object`: Boolean for whether the client supports the modern   `user_settings` event type. If false, the server will additionally send the   legacy `update_global_notifications` and `update_display_settings` event   types.   <br />   **Changes**: New in Zulip 5.0 (feature level 89). This capability is for   backwards-compatibility; it will be removed in a future server release.   Because the feature level 89 API changes were merged together, clients can   safely make a request with this client capability and also request all three   event types (`user_settings`, `update_display_settings`,   `update_global_notifications`), and get exactly one copy of settings data on   any server version. Clients can then use the `zulip_feature_level` in the   `/register` response or the presence/absence of a `user_settings` key to   determine where to look for the data.  - `linkifier_url_template`: Boolean for whether the client accepts   [linkifiers][help-linkifiers] that use [RFC 6570][rfc6570] compliant   URL templates for linkifying matches. If false or unset, then the   `realm_linkifiers` array in the `/register` response will be empty   if present, and no `realm_linkifiers` [events][events-linkifiers]   will be sent to the client.   <br />   **Changes**: New in Zulip 7.0 (feature level 176). This capability   is for backwards-compatibility.  - `user_list_incomplete`: Boolean for whether the client supports not having an   incomplete user database. If true, then the `realm_users` array in the `register`   response will not include data for inaccessible users and clients of guest users will   not receive `realm_user op:add` events for newly created users that are not accessible   to the current user.   <br />   **Changes**: New in Zulip 8.0 (feature level 232). This   capability is for backwards-compatibility.  - `include_deactivated_groups`: Boolean for whether the client can handle   deactivated user groups by themselves. If false, then the `realm_user_groups`   array in the `/register` response will only include active groups, clients   will receive a `remove` event instead of `update` event when a group is   deactivated and no `update` event will be sent to the client if a deactivated   user group is renamed.   <br />   **Changes**: New in Zulip 10.0 (feature level 294). This   capability is for backwards-compatibility.  - `archived_channels`: Boolean for whether the client supports processing   [archived channels](zulip.com/help/archive-a-channel in the `stream` and   `subscription` event types. If `false`, the server will not include data   related to archived channels in the `register` response or in events.   <br />   **Changes**: New in Zulip 10.0 (feature level 315). This allows clients to   access archived channels, without breaking backwards-compatibility for   existing clients.  - `empty_topic_name`: Boolean for whether the client supports processing   the empty string as a topic name. Clients not declaring this capability   will be sent the value of `realm_empty_topic_display_name` found in the   [POST /register](zulip.com/api/register-queue response instead of the empty string   wherever topic names appear in the register response or events involving   topic names.   <br/>   **Changes**: New in Zulip 10.0 (feature level 334). Previously,   the empty string was not a valid topic name.  - `simplified_presence_events`: Boolean for whether the client supports   receiving the [`presence` event type](zulip.com/api/get-events#presence with   user presence data in the modern format. If true, the server will   send these events with the `presences` field that has the user presence   data in the modern format. Otherwise, these event will contain fields   with legacy format user presence data.   <br />   **Changes**: New in Zulip 11.0 (feature level 419).  [help-linkifiers]: /help/add-a-custom-linkifier [rfc6570]: https://www.rfc-editor.org/rfc/rfc6570.html [events-linkifiers]: /api/get-events#realm_linkifiers  (optional)
	fetchEventTypes := []string{"Inner_example"} // []string | Same as the `event_types` parameter except that the values in `fetch_event_types` are used to fetch initial data. If `fetch_event_types` is not provided, `event_types` is used and if `event_types` is not provided, this parameter defaults to `null`.  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  (optional)
	narrow := [][]string{[]string{"Inner_example"}} // [][]string | A JSON-encoded array of arrays of length 2 indicating the [narrow filter(s)](zulip.com/api/construct-narrow for which you'd like to receive events for.  For example, to receive events for direct messages (including group direct messages) received by the user, one can use `\\\"narrow\\\": [[\\\"is\\\", \\\"dm\\\"]]`.  Unlike the API for [fetching messages](zulip.com/api/get-messages, this narrow parameter is simply a filter on messages that the user receives through their channel subscriptions (or because they are a recipient of a direct message).  This means that a client that requests a `narrow` filter of `[[\\\"channel\\\", \\\"Denmark\\\"]]` will receive events for new messages sent to that channel while the user is subscribed to that channel. The client will not receive any message events at all if the user is not subscribed to `\\\"Denmark\\\"`.  Newly created bot users are not usually subscribed to any channels, so bots using this API need to be [subscribed](zulip.com/api/subscribe to any channels whose messages you'd like them to process using this endpoint.  See the `all_public_streams` parameter for how to process all public channel messages in an organization.  **Changes**: See [changes section](zulip.com/api/construct-narrow#changes of search/narrow filter documentation.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegisterQueue(context.Background()).ApplyMarkdown(applyMarkdown).ClientGravatar(clientGravatar).IncludeSubscribers(includeSubscribers).SlimPresence(slimPresence).PresenceHistoryLimitDays(presenceHistoryLimitDays).EventTypes(eventTypes).AllPublicStreams(allPublicStreams).ClientCapabilities(clientCapabilities).FetchEventTypes(fetchEventTypes).Narrow(narrow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealTimeEventsAPI.RegisterQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterQueue`: RegisterQueue200Response
	fmt.Fprintf(os.Stdout, "Response from `RealTimeEventsAPI.RegisterQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyMarkdown** | **bool** | Set to &#x60;true&#x60; if you would like the content to be rendered in HTML format (otherwise the API will return the raw text that the user entered)  | [default to false]
 **clientGravatar** | **bool** | Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  The default value is &#x60;true&#x60; for authenticated requests and &#x60;false&#x60; for [unauthenticated requests](zulip.com/help/public-access-option. Passing &#x60;true&#x60; in an unauthenticated request is an error.  **Changes**: Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were &#x60;false&#x60; in unauthenticated requests.  | 
 **includeSubscribers** | **string** | Whether each returned channel object should include a &#x60;subscribers&#x60; field containing a list of the user IDs of its subscribers.  Client apps supporting organizations with many thousands of users should not pass &#x60;true&#x60;, because the full subscriber matrix may be several megabytes of data. The &#x60;partial&#x60; value, combined with the &#x60;subscriber_count&#x60; and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes &#x60;partial&#x60; for this parameter, the server may, for some channels, return a subset of the channel&#39;s subscribers in the &#x60;partial_subscribers&#x60; field instead of the &#x60;subscribers&#x60; field, which always contains the complete set of subscribers.  The server guarantees that it will always return a &#x60;subscribers&#x60; field for channels with fewer than 250 total subscribers. When returning a &#x60;partial_subscribers&#x60; field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  Passing &#x60;true&#x60; in an [unauthenticated request](zulip.com/help/public-access-option is an error.  **Changes**: The &#x60;partial&#x60; value is new in Zulip 11.0 (feature level 412).  Before Zulip 6.0 (feature level 149), this parameter was silently ignored and processed as though it were &#x60;false&#x60; in unauthenticated requests.  New in Zulip 2.1.0.  | [default to &quot;false&quot;]
 **slimPresence** | **bool** | If &#x60;true&#x60;, the &#x60;presences&#x60; object returned in the response will be keyed by user ID and the entry for each user&#39;s presence data will be in the modern format.  **Changes**: New in Zulip 3.0 (no feature level; API unstable).  | [default to false]
 **presenceHistoryLimitDays** | **int32** | Limits how far back in time to fetch user presence data. If not specified, defaults to 14 days. A value of N means that the oldest presence data fetched will be from at most N days ago.  **Changes**: New in Zulip 10.0 (feature level 288).  | 
 **eventTypes** | **[]string** | A JSON-encoded array indicating which types of events you&#39;re interested in. Values that you might find useful include:  - **message** (messages) - **subscription** (changes in your subscriptions) - **realm_user** (changes to users in the organization and   their properties, such as their name).  If you do not specify this parameter, you will receive all events, and have to filter out the events not relevant to your client in your client code. For most applications, one is only interested in messages, so one specifies: &#x60;\\\&quot;event_types\\\&quot;: [\\\&quot;message\\\&quot;]&#x60;  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  | 
 **allPublicStreams** | **bool** | Whether you would like to request message events from all public channels. Useful for workflow bots that you&#39;d like to see all new messages sent to public channels. (You can also subscribe the user to private channels).  | [default to false]
 **clientCapabilities** | [**map[string]interface{}**](map[string]interface{}.md) | Dictionary containing details on features the client supports that are relevant to the format of responses sent by the server.  - &#x60;notification_settings_null&#x60;: Boolean for whether the   client can handle the current API with &#x60;null&#x60; values for   channel-level notification settings (which means the channel   is not customized and should inherit the user&#39;s global   notification settings for channel messages).   &lt;br /&gt;   **Changes**: New in Zulip 2.1.0. In earlier Zulip releases,   channel-level notification settings were simple booleans.  - &#x60;bulk_message_deletion&#x60;: Boolean for whether the client&#39;s   handler for the &#x60;delete_message&#x60; event type has been   updated to process the new bulk format (with a   &#x60;message_ids&#x60;, rather than a singleton &#x60;message_id&#x60;).   Otherwise, the server will send &#x60;delete_message&#x60; events   in a loop.   &lt;br /&gt;   **Changes**: New in Zulip 3.0 (feature level 13). This   capability is for backwards-compatibility; it will be   required in a future server release.  - &#x60;user_avatar_url_field_optional&#x60;: Boolean for whether the   client required avatar URLs for all users, or supports   using &#x60;GET /avatar/{user_id}&#x60; to access user avatars. If the   client has this capability, the server may skip sending a   &#x60;avatar_url&#x60; field in the &#x60;realm_user&#x60; at its sole discretion   to optimize network performance. This is an important optimization   in organizations with 10,000s of users.   &lt;br /&gt;   **Changes**: New in Zulip 3.0 (feature level 18).  - &#x60;stream_typing_notifications&#x60;: Boolean for whether the client   supports channel typing notifications.   &lt;br /&gt;   **Changes**: New in Zulip 4.0 (feature level 58). This capability is   for backwards-compatibility; it will be required in a   future server release.  - &#x60;user_settings_object&#x60;: Boolean for whether the client supports the modern   &#x60;user_settings&#x60; event type. If false, the server will additionally send the   legacy &#x60;update_global_notifications&#x60; and &#x60;update_display_settings&#x60; event   types.   &lt;br /&gt;   **Changes**: New in Zulip 5.0 (feature level 89). This capability is for   backwards-compatibility; it will be removed in a future server release.   Because the feature level 89 API changes were merged together, clients can   safely make a request with this client capability and also request all three   event types (&#x60;user_settings&#x60;, &#x60;update_display_settings&#x60;,   &#x60;update_global_notifications&#x60;), and get exactly one copy of settings data on   any server version. Clients can then use the &#x60;zulip_feature_level&#x60; in the   &#x60;/register&#x60; response or the presence/absence of a &#x60;user_settings&#x60; key to   determine where to look for the data.  - &#x60;linkifier_url_template&#x60;: Boolean for whether the client accepts   [linkifiers][help-linkifiers] that use [RFC 6570][rfc6570] compliant   URL templates for linkifying matches. If false or unset, then the   &#x60;realm_linkifiers&#x60; array in the &#x60;/register&#x60; response will be empty   if present, and no &#x60;realm_linkifiers&#x60; [events][events-linkifiers]   will be sent to the client.   &lt;br /&gt;   **Changes**: New in Zulip 7.0 (feature level 176). This capability   is for backwards-compatibility.  - &#x60;user_list_incomplete&#x60;: Boolean for whether the client supports not having an   incomplete user database. If true, then the &#x60;realm_users&#x60; array in the &#x60;register&#x60;   response will not include data for inaccessible users and clients of guest users will   not receive &#x60;realm_user op:add&#x60; events for newly created users that are not accessible   to the current user.   &lt;br /&gt;   **Changes**: New in Zulip 8.0 (feature level 232). This   capability is for backwards-compatibility.  - &#x60;include_deactivated_groups&#x60;: Boolean for whether the client can handle   deactivated user groups by themselves. If false, then the &#x60;realm_user_groups&#x60;   array in the &#x60;/register&#x60; response will only include active groups, clients   will receive a &#x60;remove&#x60; event instead of &#x60;update&#x60; event when a group is   deactivated and no &#x60;update&#x60; event will be sent to the client if a deactivated   user group is renamed.   &lt;br /&gt;   **Changes**: New in Zulip 10.0 (feature level 294). This   capability is for backwards-compatibility.  - &#x60;archived_channels&#x60;: Boolean for whether the client supports processing   [archived channels](zulip.com/help/archive-a-channel in the &#x60;stream&#x60; and   &#x60;subscription&#x60; event types. If &#x60;false&#x60;, the server will not include data   related to archived channels in the &#x60;register&#x60; response or in events.   &lt;br /&gt;   **Changes**: New in Zulip 10.0 (feature level 315). This allows clients to   access archived channels, without breaking backwards-compatibility for   existing clients.  - &#x60;empty_topic_name&#x60;: Boolean for whether the client supports processing   the empty string as a topic name. Clients not declaring this capability   will be sent the value of &#x60;realm_empty_topic_display_name&#x60; found in the   [POST /register](zulip.com/api/register-queue response instead of the empty string   wherever topic names appear in the register response or events involving   topic names.   &lt;br/&gt;   **Changes**: New in Zulip 10.0 (feature level 334). Previously,   the empty string was not a valid topic name.  - &#x60;simplified_presence_events&#x60;: Boolean for whether the client supports   receiving the [&#x60;presence&#x60; event type](zulip.com/api/get-events#presence with   user presence data in the modern format. If true, the server will   send these events with the &#x60;presences&#x60; field that has the user presence   data in the modern format. Otherwise, these event will contain fields   with legacy format user presence data.   &lt;br /&gt;   **Changes**: New in Zulip 11.0 (feature level 419).  [help-linkifiers]: /help/add-a-custom-linkifier [rfc6570]: https://www.rfc-editor.org/rfc/rfc6570.html [events-linkifiers]: /api/get-events#realm_linkifiers  | 
 **fetchEventTypes** | **[]string** | Same as the &#x60;event_types&#x60; parameter except that the values in &#x60;fetch_event_types&#x60; are used to fetch initial data. If &#x60;fetch_event_types&#x60; is not provided, &#x60;event_types&#x60; is used and if &#x60;event_types&#x60; is not provided, this parameter defaults to &#x60;null&#x60;.  Event types not supported by the server are ignored, in order to simplify the implementation of client apps that support multiple server versions.  | 
 **narrow** | **[][]string** | A JSON-encoded array of arrays of length 2 indicating the [narrow filter(s)](zulip.com/api/construct-narrow for which you&#39;d like to receive events for.  For example, to receive events for direct messages (including group direct messages) received by the user, one can use &#x60;\\\&quot;narrow\\\&quot;: [[\\\&quot;is\\\&quot;, \\\&quot;dm\\\&quot;]]&#x60;.  Unlike the API for [fetching messages](zulip.com/api/get-messages, this narrow parameter is simply a filter on messages that the user receives through their channel subscriptions (or because they are a recipient of a direct message).  This means that a client that requests a &#x60;narrow&#x60; filter of &#x60;[[\\\&quot;channel\\\&quot;, \\\&quot;Denmark\\\&quot;]]&#x60; will receive events for new messages sent to that channel while the user is subscribed to that channel. The client will not receive any message events at all if the user is not subscribed to &#x60;\\\&quot;Denmark\\\&quot;&#x60;.  Newly created bot users are not usually subscribed to any channels, so bots using this API need to be [subscribed](zulip.com/api/subscribe to any channels whose messages you&#39;d like them to process using this endpoint.  See the &#x60;all_public_streams&#x60; parameter for how to process all public channel messages in an organization.  **Changes**: See [changes section](zulip.com/api/construct-narrow#changes of search/narrow filter documentation.  | 

### Return type

[**RegisterQueue200Response**](RegisterQueue200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestErrorHandling

> RestErrorHandling(ctx).Execute()

Error handling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RestErrorHandling(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealTimeEventsAPI.RestErrorHandling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRestErrorHandlingRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

