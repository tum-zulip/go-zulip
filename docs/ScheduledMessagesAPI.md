# \ScheduledMessagesAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduledMessage**](ScheduledMessagesAPI.md#CreateScheduledMessage) | **Post** /scheduled_messages | Create a scheduled message
[**DeleteScheduledMessage**](ScheduledMessagesAPI.md#DeleteScheduledMessage) | **Delete** /scheduled_messages/{scheduled_message_id} | Delete a scheduled message
[**GetScheduledMessages**](ScheduledMessagesAPI.md#GetScheduledMessages) | **Get** /scheduled_messages | Get scheduled messages
[**UpdateScheduledMessage**](ScheduledMessagesAPI.md#UpdateScheduledMessage) | **Patch** /scheduled_messages/{scheduled_message_id} | Edit a scheduled message



## CreateScheduledMessage

> CreateScheduledMessage200Response CreateScheduledMessage(ctx).Type_(type_).To(to).Content(content).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Topic(topic).ReadBySender(readBySender).Execute()

Create a scheduled message



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
	type_ := "type__example" // string | The type of scheduled message to be sent. `\\\"direct\\\"` for a direct message and `\\\"stream\\\"` or `\\\"channel\\\"` for a channel message.  Note that, while `\\\"private\\\"` is supported for scheduling direct messages, clients are encouraged to use to the modern convention of `\\\"direct\\\"` to indicate this message type, because support for `\\\"private\\\"` may eventually be removed.  **Changes**: In Zulip 9.0 (feature level 248), `\\\"channel\\\"` was added as an additional value for this parameter to indicate the type of a channel message. 
	to := openapiclient.create_scheduled_message_request_to{ArrayOfInt32: new([]int32)} // CreateScheduledMessageRequestTo | 
	content := "content_example" // string | The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum message size. 
	scheduledDeliveryTimestamp := int32(56) // int32 | The UNIX timestamp for when the message will be sent, in UTC seconds. 
	topic := "topic_example" // string | The topic of the message. Only required for channel messages (`\\\"type\\\": \\\"stream\\\"` or `\\\"type\\\": \\\"channel\\\"`), ignored otherwise.  Clients should use the `max_topic_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum topic length.  Note: When `\\\"(no topic)\\\"` or the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can't be `\\\"(no topic)\\\"`, an empty string, or the value of `realm_empty_topic_display_name`.  **Changes**: Before Zulip 10.0 (feature level 370), `\\\"(no topic)\\\"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  (optional)
	readBySender := true // bool | Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name and the recipient.  **Changes**: New in Zulip 8.0 (feature level 236).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateScheduledMessage(context.Background()).Type_(type_).To(to).Content(content).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Topic(topic).ReadBySender(readBySender).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledMessagesAPI.CreateScheduledMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScheduledMessage`: CreateScheduledMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `ScheduledMessagesAPI.CreateScheduledMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduledMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of scheduled message to be sent. &#x60;\\\&quot;direct\\\&quot;&#x60; for a direct message and &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; for a channel message.  Note that, while &#x60;\\\&quot;private\\\&quot;&#x60; is supported for scheduling direct messages, clients are encouraged to use to the modern convention of &#x60;\\\&quot;direct\\\&quot;&#x60; to indicate this message type, because support for &#x60;\\\&quot;private\\\&quot;&#x60; may eventually be removed.  **Changes**: In Zulip 9.0 (feature level 248), &#x60;\\\&quot;channel\\\&quot;&#x60; was added as an additional value for this parameter to indicate the type of a channel message.  | 
 **to** | [**CreateScheduledMessageRequestTo**](CreateScheduledMessageRequestTo.md) |  | 
 **content** | **string** | The content of the message.  Clients should use the &#x60;max_message_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum message size.  | 
 **scheduledDeliveryTimestamp** | **int32** | The UNIX timestamp for when the message will be sent, in UTC seconds.  | 
 **topic** | **string** | The topic of the message. Only required for channel messages (&#x60;\\\&quot;type\\\&quot;: \\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;type\\\&quot;: \\\&quot;channel\\\&quot;&#x60;), ignored otherwise.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum topic length.  Note: When &#x60;\\\&quot;(no topic)\\\&quot;&#x60; or the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can&#39;t be &#x60;\\\&quot;(no topic)\\\&quot;&#x60;, an empty string, or the value of &#x60;realm_empty_topic_display_name&#x60;.  **Changes**: Before Zulip 10.0 (feature level 370), &#x60;\\\&quot;(no topic)\\\&quot;&#x60; was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  | 
 **readBySender** | **bool** | Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name and the recipient.  **Changes**: New in Zulip 8.0 (feature level 236).  | 

### Return type

[**CreateScheduledMessage200Response**](CreateScheduledMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledMessage

> JsonSuccess DeleteScheduledMessage(ctx, scheduledMessageId).Execute()

Delete a scheduled message



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
	scheduledMessageId := int32(1) // int32 | The ID of the scheduled message to delete.  This is different from the unique ID that the message would have after being sent. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteScheduledMessage(context.Background(), scheduledMessageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledMessagesAPI.DeleteScheduledMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteScheduledMessage`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ScheduledMessagesAPI.DeleteScheduledMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledMessageId** | **int32** | The ID of the scheduled message to delete.  This is different from the unique ID that the message would have after being sent.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledMessages

> GetScheduledMessages200Response GetScheduledMessages(ctx).Execute()

Get scheduled messages



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
	resp, r, err := apiClient.GetScheduledMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledMessagesAPI.GetScheduledMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledMessages`: GetScheduledMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `ScheduledMessagesAPI.GetScheduledMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledMessagesRequest struct via the builder pattern


### Return type

[**GetScheduledMessages200Response**](GetScheduledMessages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScheduledMessage

> UpdateScheduledMessage200Response UpdateScheduledMessage(ctx, scheduledMessageId).Type_(type_).To(to).Content(content).Topic(topic).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Execute()

Edit a scheduled message



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
	scheduledMessageId := int32(2) // int32 | The ID of the scheduled message to update.  This is different from the unique ID that the message would have after being sent. 
	type_ := "type__example" // string | The type of scheduled message to be sent. `\\\"direct\\\"` for a direct message and `\\\"stream\\\"` or `\\\"channel\\\"` for a channel message.  When updating the type of the scheduled message, the `to` parameter is required. And, if updating the type of the scheduled message to `\\\"stream\\\"`/`\\\"channel\\\"`, then the `topic` parameter is also required.  Note that, while `\\\"private\\\"` is supported for scheduling direct messages, clients are encouraged to use to the modern convention of `\\\"direct\\\"` to indicate this message type, because support for `\\\"private\\\"` may eventually be removed.  **Changes**: In Zulip 9.0 (feature level 248), `\\\"channel\\\"` was added as an additional value for this parameter to indicate the type of a channel message.  (optional)
	to := openapiclient.update_scheduled_message_request_to{ArrayOfInt32: new([]int32)} // UpdateScheduledMessageRequestTo |  (optional)
	content := "content_example" // string | The updated content of the scheduled message.  Clients should use the `max_message_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum message size.  (optional)
	topic := "topic_example" // string | The updated topic of the scheduled message.  Required when updating the `type` of the scheduled message to `\\\"stream\\\"` or `\\\"channel\\\"`. Ignored when the existing or updated `type` of the scheduled message is `\\\"direct\\\"` (or `\\\"private\\\"`).  Clients should use the `max_topic_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum topic length.  Note: When `\\\"(no topic)\\\"` or the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can't be `\\\"(no topic)\\\"`, an empty string, or the value of `realm_empty_topic_display_name`.  **Changes**: Before Zulip 10.0 (feature level 370), `\\\"(no topic)\\\"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  (optional)
	scheduledDeliveryTimestamp := int32(56) // int32 | The UNIX timestamp for when the message will be sent, in UTC seconds.  Required when updating a scheduled message that the server has already tried and failed to send. This state is indicated with `\\\"failed\\\": true` in `scheduled_messages` objects; see response description at [`GET /scheduled_messages`](/api/get-scheduled-messages#response).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateScheduledMessage(context.Background(), scheduledMessageId).Type_(type_).To(to).Content(content).Topic(topic).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledMessagesAPI.UpdateScheduledMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScheduledMessage`: UpdateScheduledMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `ScheduledMessagesAPI.UpdateScheduledMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledMessageId** | **int32** | The ID of the scheduled message to update.  This is different from the unique ID that the message would have after being sent.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduledMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | The type of scheduled message to be sent. &#x60;\\\&quot;direct\\\&quot;&#x60; for a direct message and &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; for a channel message.  When updating the type of the scheduled message, the &#x60;to&#x60; parameter is required. And, if updating the type of the scheduled message to &#x60;\\\&quot;stream\\\&quot;&#x60;/&#x60;\\\&quot;channel\\\&quot;&#x60;, then the &#x60;topic&#x60; parameter is also required.  Note that, while &#x60;\\\&quot;private\\\&quot;&#x60; is supported for scheduling direct messages, clients are encouraged to use to the modern convention of &#x60;\\\&quot;direct\\\&quot;&#x60; to indicate this message type, because support for &#x60;\\\&quot;private\\\&quot;&#x60; may eventually be removed.  **Changes**: In Zulip 9.0 (feature level 248), &#x60;\\\&quot;channel\\\&quot;&#x60; was added as an additional value for this parameter to indicate the type of a channel message.  | 
 **to** | [**UpdateScheduledMessageRequestTo**](UpdateScheduledMessageRequestTo.md) |  | 
 **content** | **string** | The updated content of the scheduled message.  Clients should use the &#x60;max_message_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum message size.  | 
 **topic** | **string** | The updated topic of the scheduled message.  Required when updating the &#x60;type&#x60; of the scheduled message to &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60;. Ignored when the existing or updated &#x60;type&#x60; of the scheduled message is &#x60;\\\&quot;direct\\\&quot;&#x60; (or &#x60;\\\&quot;private\\\&quot;&#x60;).  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum topic length.  Note: When &#x60;\\\&quot;(no topic)\\\&quot;&#x60; or the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can&#39;t be &#x60;\\\&quot;(no topic)\\\&quot;&#x60;, an empty string, or the value of &#x60;realm_empty_topic_display_name&#x60;.  **Changes**: Before Zulip 10.0 (feature level 370), &#x60;\\\&quot;(no topic)\\\&quot;&#x60; was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  | 
 **scheduledDeliveryTimestamp** | **int32** | The UNIX timestamp for when the message will be sent, in UTC seconds.  Required when updating a scheduled message that the server has already tried and failed to send. This state is indicated with &#x60;\\\&quot;failed\\\&quot;: true&#x60; in &#x60;scheduled_messages&#x60; objects; see response description at [&#x60;GET /scheduled_messages&#x60;](/api/get-scheduled-messages#response).  | 

### Return type

[**UpdateScheduledMessage200Response**](UpdateScheduledMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

