# \RemindersAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageReminder**](RemindersAPI.md#CreateMessageReminder) | **Post** /reminders | Create a message reminder
[**DeleteReminder**](RemindersAPI.md#DeleteReminder) | **Delete** /reminders/{reminder_id} | Delete a reminder
[**GetReminders**](RemindersAPI.md#GetReminders) | **Get** /reminders | Get reminders



## CreateMessageReminder

> CreateMessageReminder200Response CreateMessageReminder(ctx).MessageId(messageId).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Note(note).Execute()

Create a message reminder



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
	messageId := int32(56) // int32 | The ID of the previously sent message to reference in the reminder message.  (optional)
	scheduledDeliveryTimestamp := int32(56) // int32 | The UNIX timestamp for when the reminder will be sent, in UTC seconds.  (optional)
	note := "note_example" // string | A note associated with the reminder shown in the Notification Bot message.  **Changes**: New in Zulip 11.0 (feature level 415).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateMessageReminder(context.Background()).MessageId(messageId).ScheduledDeliveryTimestamp(scheduledDeliveryTimestamp).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.CreateMessageReminder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessageReminder`: CreateMessageReminder200Response
	fmt.Fprintf(os.Stdout, "Response from `RemindersAPI.CreateMessageReminder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageReminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **int32** | The ID of the previously sent message to reference in the reminder message.  | 
 **scheduledDeliveryTimestamp** | **int32** | The UNIX timestamp for when the reminder will be sent, in UTC seconds.  | 
 **note** | **string** | A note associated with the reminder shown in the Notification Bot message.  **Changes**: New in Zulip 11.0 (feature level 415).  | 

### Return type

[**CreateMessageReminder200Response**](CreateMessageReminder200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReminder

> JsonSuccess DeleteReminder(ctx, reminderId).Execute()

Delete a reminder



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
	reminderId := int32(1) // int32 | The ID of the reminder to delete.  This is different from the unique ID that the message would have after being sent. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteReminder(context.Background(), reminderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.DeleteReminder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReminder`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `RemindersAPI.DeleteReminder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reminderId** | **int32** | The ID of the reminder to delete.  This is different from the unique ID that the message would have after being sent.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReminderRequest struct via the builder pattern


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


## GetReminders

> GetReminders200Response GetReminders(ctx).Execute()

Get reminders



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
	resp, r, err := apiClient.GetReminders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemindersAPI.GetReminders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReminders`: GetReminders200Response
	fmt.Fprintf(os.Stdout, "Response from `RemindersAPI.GetReminders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemindersRequest struct via the builder pattern


### Return type

[**GetReminders200Response**](GetReminders200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

