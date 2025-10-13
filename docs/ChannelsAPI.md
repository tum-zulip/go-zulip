# \ChannelsAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDefaultStream**](ChannelsAPI.md#AddDefaultStream) | **Post** /default_streams | Add a default channel
[**ArchiveStream**](ChannelsAPI.md#ArchiveStream) | **Delete** /streams/{stream_id} | Archive a channel
[**CreateBigBlueButtonVideoCall**](ChannelsAPI.md#CreateBigBlueButtonVideoCall) | **Get** /calls/bigbluebutton/create | Create BigBlueButton video call
[**CreateChannel**](ChannelsAPI.md#CreateChannel) | **Post** /channels/create | Create a channel
[**CreateChannelFolder**](ChannelsAPI.md#CreateChannelFolder) | **Post** /channel_folders/create | Create a channel folder
[**DeleteTopic**](ChannelsAPI.md#DeleteTopic) | **Post** /streams/{stream_id}/delete_topic | Delete a topic
[**GetChannelFolders**](ChannelsAPI.md#GetChannelFolders) | **Get** /channel_folders | Get channel folders
[**GetStreamById**](ChannelsAPI.md#GetStreamById) | **Get** /streams/{stream_id} | Get a channel by ID
[**GetStreamEmailAddress**](ChannelsAPI.md#GetStreamEmailAddress) | **Get** /streams/{stream_id}/email_address | Get channel&#39;s email address
[**GetStreamId**](ChannelsAPI.md#GetStreamId) | **Get** /get_stream_id | Get channel ID
[**GetStreamTopics**](ChannelsAPI.md#GetStreamTopics) | **Get** /users/me/{stream_id}/topics | Get topics in a channel
[**GetStreams**](ChannelsAPI.md#GetStreams) | **Get** /streams | Get all channels
[**GetSubscribers**](ChannelsAPI.md#GetSubscribers) | **Get** /streams/{stream_id}/members | Get channel subscribers
[**GetSubscriptionStatus**](ChannelsAPI.md#GetSubscriptionStatus) | **Get** /users/{user_id}/subscriptions/{stream_id} | Get subscription status
[**GetSubscriptions**](ChannelsAPI.md#GetSubscriptions) | **Get** /users/me/subscriptions | Get subscribed channels
[**MuteTopic**](ChannelsAPI.md#MuteTopic) | **Patch** /users/me/subscriptions/muted_topics | Topic muting
[**PatchChannelFolders**](ChannelsAPI.md#PatchChannelFolders) | **Patch** /channel_folders | Reorder channel folders
[**RemoveDefaultStream**](ChannelsAPI.md#RemoveDefaultStream) | **Delete** /default_streams | Remove a default channel
[**Subscribe**](ChannelsAPI.md#Subscribe) | **Post** /users/me/subscriptions | Subscribe to a channel
[**Unsubscribe**](ChannelsAPI.md#Unsubscribe) | **Delete** /users/me/subscriptions | Unsubscribe from a channel
[**UpdateChannelFolder**](ChannelsAPI.md#UpdateChannelFolder) | **Patch** /channel_folders/{channel_folder_id} | Update a channel folder
[**UpdateStream**](ChannelsAPI.md#UpdateStream) | **Patch** /streams/{stream_id} | Update a channel
[**UpdateSubscriptionSettings**](ChannelsAPI.md#UpdateSubscriptionSettings) | **Post** /users/me/subscriptions/properties | Update subscription settings
[**UpdateSubscriptions**](ChannelsAPI.md#UpdateSubscriptions) | **Patch** /users/me/subscriptions | Update subscriptions
[**UpdateUserTopic**](ChannelsAPI.md#UpdateUserTopic) | **Post** /user_topics | Update personal preferences for a topic



## AddDefaultStream

> JsonSuccess AddDefaultStream(ctx).StreamId(streamId).Execute()

Add a default channel



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
	streamId := int32(56) // int32 | The ID of the target channel. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddDefaultStream(context.Background()).StreamId(streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.AddDefaultStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDefaultStream`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.AddDefaultStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDefaultStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **int32** | The ID of the target channel.  | 

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


## ArchiveStream

> JsonSuccess ArchiveStream(ctx, streamId).Execute()

Archive a channel



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
	streamId := int32(1) // int32 | The ID of the channel to access. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArchiveStream(context.Background(), streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.ArchiveStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveStream`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.ArchiveStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveStreamRequest struct via the builder pattern


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


## CreateBigBlueButtonVideoCall

> CreateBigBlueButtonVideoCall200Response CreateBigBlueButtonVideoCall(ctx).MeetingName(meetingName).VoiceOnly(voiceOnly).Execute()

Create BigBlueButton video call



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
	meetingName := "test_channel meeting" // string | Meeting name for the BigBlueButton video call. 
	voiceOnly := true // bool | Configures whether the call is voice-only; if true, disables cameras for all users. Only the call creator/moderator can edit this configuration.  **Changes**: New in Zulip 10.0 (feature level 337).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateBigBlueButtonVideoCall(context.Background()).MeetingName(meetingName).VoiceOnly(voiceOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CreateBigBlueButtonVideoCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBigBlueButtonVideoCall`: CreateBigBlueButtonVideoCall200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CreateBigBlueButtonVideoCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigBlueButtonVideoCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meetingName** | **string** | Meeting name for the BigBlueButton video call.  | 
 **voiceOnly** | **bool** | Configures whether the call is voice-only; if true, disables cameras for all users. Only the call creator/moderator can edit this configuration.  **Changes**: New in Zulip 10.0 (feature level 337).  | 

### Return type

[**CreateBigBlueButtonVideoCall200Response**](CreateBigBlueButtonVideoCall200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannel

> CreateChannel200Response CreateChannel(ctx).Name(name).Subscribers(subscribers).Description(description).Announce(announce).InviteOnly(inviteOnly).IsWebPublic(isWebPublic).IsDefaultStream(isDefaultStream).FolderId(folderId).SendNewSubscriptionMessages(sendNewSubscriptionMessages).TopicsPolicy(topicsPolicy).HistoryPublicToSubscribers(historyPublicToSubscribers).MessageRetentionDays(messageRetentionDays).CanAddSubscribersGroup(canAddSubscribersGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).Execute()

Create a channel



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
	name := "name_example" // string | The name of the new channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel name length. 
	subscribers := []int32{int32(123)} // []int32 | A list of user IDs of the users to be subscribed to the new channel. 
	description := "description_example" // string | The [description](zulip.com/help/change-the-channel-description to use for the new channel being created, in text/markdown format.  Clients should use the `max_stream_description_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel description length.  (optional)
	announce := true // bool | This determines whether [notification bot](zulip.com/help/configure-automated-notices will send an announcement about the new channel's creation.  (optional) (default to false)
	inviteOnly := true // bool | This parameter and the ones that follow are used to request an initial configuration of the new channel.  This parameter determines whether the newly created channel will be a private channel.  (optional) (default to false)
	isWebPublic := true // bool | This parameter determines whether the newly created channel will be a web-public channel.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED` [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current user to have permission under the organization's `can_create_web_public_channel_group` realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  (optional) (default to false)
	isDefaultStream := true // bool | This parameter determines whether the newly created channel will be added as a [default channel][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users  (optional) (default to false)
	folderId := int32(56) // int32 | The ID of the folder to which the channel belongs.  Is `null` if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  (optional)
	sendNewSubscriptionMessages := true // bool | Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than `max_bulk_new_subscription_messages` (available in the [`POST /register`](zulip.com/api/register-queue response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.  (optional) (default to true)
	topicsPolicy := openapiclient.TopicsPolicy("inherit") // TopicsPolicy |  (optional)
	historyPublicToSubscribers := true // bool | Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  (optional)
	messageRetentionDays := openapiclient.MessageRetentionDays{Int32: new(int32)} // MessageRetentionDays |  (optional)
	canAddSubscribersGroup := *openapiclient.NewChannelCanAddSubscribersGroup() // ChannelCanAddSubscribersGroup |  (optional)
	canDeleteAnyMessageGroup := *openapiclient.NewCanDeleteAnyMessageGroup() // CanDeleteAnyMessageGroup |  (optional)
	canDeleteOwnMessageGroup := *openapiclient.NewCanDeleteOwnMessageGroup() // CanDeleteOwnMessageGroup |  (optional)
	canRemoveSubscribersGroup := *openapiclient.NewCanRemoveSubscribersGroup() // CanRemoveSubscribersGroup |  (optional)
	canAdministerChannelGroup := *openapiclient.NewCanAdministerChannelGroup() // CanAdministerChannelGroup |  (optional)
	canMoveMessagesOutOfChannelGroup := *openapiclient.NewCanMoveMessagesOutOfChannelGroup() // CanMoveMessagesOutOfChannelGroup |  (optional)
	canMoveMessagesWithinChannelGroup := *openapiclient.NewCanMoveMessagesWithinChannelGroup() // CanMoveMessagesWithinChannelGroup |  (optional)
	canSendMessageGroup := *openapiclient.NewCanSendMessageGroup() // CanSendMessageGroup |  (optional)
	canSubscribeGroup := *openapiclient.NewCanSubscribeGroup() // CanSubscribeGroup |  (optional)
	canResolveTopicsGroup := *openapiclient.NewCanResolveTopicsGroup() // CanResolveTopicsGroup |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateChannel(context.Background()).Name(name).Subscribers(subscribers).Description(description).Announce(announce).InviteOnly(inviteOnly).IsWebPublic(isWebPublic).IsDefaultStream(isDefaultStream).FolderId(folderId).SendNewSubscriptionMessages(sendNewSubscriptionMessages).TopicsPolicy(topicsPolicy).HistoryPublicToSubscribers(historyPublicToSubscribers).MessageRetentionDays(messageRetentionDays).CanAddSubscribersGroup(canAddSubscribersGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannel`: CreateChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the new channel.  Clients should use the &#x60;max_stream_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel name length.  | 
 **subscribers** | **[]int32** | A list of user IDs of the users to be subscribed to the new channel.  | 
 **description** | **string** | The [description](zulip.com/help/change-the-channel-description to use for the new channel being created, in text/markdown format.  Clients should use the &#x60;max_stream_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel description length.  | 
 **announce** | **bool** | This determines whether [notification bot](zulip.com/help/configure-automated-notices will send an announcement about the new channel&#39;s creation.  | [default to false]
 **inviteOnly** | **bool** | This parameter and the ones that follow are used to request an initial configuration of the new channel.  This parameter determines whether the newly created channel will be a private channel.  | [default to false]
 **isWebPublic** | **bool** | This parameter determines whether the newly created channel will be a web-public channel.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current user to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  | [default to false]
 **isDefaultStream** | **bool** | This parameter determines whether the newly created channel will be added as a [default channel][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users  | [default to false]
 **folderId** | **int32** | The ID of the folder to which the channel belongs.  Is &#x60;null&#x60; if channel does not belong to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  | 
 **sendNewSubscriptionMessages** | **bool** | Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than &#x60;max_bulk_new_subscription_messages&#x60; (available in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.  | [default to true]
 **topicsPolicy** | [**TopicsPolicy**](TopicsPolicy.md) |  | 
 **historyPublicToSubscribers** | **bool** | Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  | 
 **messageRetentionDays** | [**MessageRetentionDays**](MessageRetentionDays.md) |  | 
 **canAddSubscribersGroup** | [**ChannelCanAddSubscribersGroup**](ChannelCanAddSubscribersGroup.md) |  | 
 **canDeleteAnyMessageGroup** | [**CanDeleteAnyMessageGroup**](CanDeleteAnyMessageGroup.md) |  | 
 **canDeleteOwnMessageGroup** | [**CanDeleteOwnMessageGroup**](CanDeleteOwnMessageGroup.md) |  | 
 **canRemoveSubscribersGroup** | [**CanRemoveSubscribersGroup**](CanRemoveSubscribersGroup.md) |  | 
 **canAdministerChannelGroup** | [**CanAdministerChannelGroup**](CanAdministerChannelGroup.md) |  | 
 **canMoveMessagesOutOfChannelGroup** | [**CanMoveMessagesOutOfChannelGroup**](CanMoveMessagesOutOfChannelGroup.md) |  | 
 **canMoveMessagesWithinChannelGroup** | [**CanMoveMessagesWithinChannelGroup**](CanMoveMessagesWithinChannelGroup.md) |  | 
 **canSendMessageGroup** | [**CanSendMessageGroup**](CanSendMessageGroup.md) |  | 
 **canSubscribeGroup** | [**CanSubscribeGroup**](CanSubscribeGroup.md) |  | 
 **canResolveTopicsGroup** | [**CanResolveTopicsGroup**](CanResolveTopicsGroup.md) |  | 

### Return type

[**CreateChannel200Response**](CreateChannel200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelFolder

> CreateChannelFolder200Response CreateChannelFolder(ctx).Name(name).Description(description).Execute()

Create a channel folder



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
	name := "name_example" // string | The name of the channel folder.  Clients should use the `max_channel_folder_name_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel folder name length.  (optional)
	description := "description_example" // string | The description of the channel folder.  Clients should use the `max_channel_folder_description_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel folder description length.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateChannelFolder(context.Background()).Name(name).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CreateChannelFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannelFolder`: CreateChannelFolder200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CreateChannelFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the channel folder.  Clients should use the &#x60;max_channel_folder_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel folder name length.  | 
 **description** | **string** | The description of the channel folder.  Clients should use the &#x60;max_channel_folder_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel folder description length.  | 

### Return type

[**CreateChannelFolder200Response**](CreateChannelFolder200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopic

> MarkAllAsRead200Response DeleteTopic(ctx, streamId).TopicName(topicName).Execute()

Delete a topic



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
	streamId := int32(1) // int32 | The ID of the channel to access. 
	topicName := "topicName_example" // string | The name of the topic to delete.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register](zulip.com/api/register-queue response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteTopic(context.Background(), streamId).TopicName(topicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.DeleteTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTopic`: MarkAllAsRead200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.DeleteTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topicName** | **string** | The name of the topic to delete.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](zulip.com/api/register-queue response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  | 

### Return type

[**MarkAllAsRead200Response**](MarkAllAsRead200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelFolders

> GetChannelFolders200Response GetChannelFolders(ctx).IncludeArchived(includeArchived).Execute()

Get channel folders



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
	includeArchived := true // bool | Whether to include archived channel folders in the response.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetChannelFolders(context.Background()).IncludeArchived(includeArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelFolders`: GetChannelFolders200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArchived** | **bool** | Whether to include archived channel folders in the response.  | [default to false]

### Return type

[**GetChannelFolders200Response**](GetChannelFolders200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamById

> GetStreamById200Response GetStreamById(ctx, streamId).Execute()

Get a channel by ID



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
	streamId := int32(1) // int32 | The ID of the channel to access. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStreamById(context.Background(), streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetStreamById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamById`: GetStreamById200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetStreamById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStreamById200Response**](GetStreamById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamEmailAddress

> GetStreamEmailAddress200Response GetStreamEmailAddress(ctx, streamId).SenderId(senderId).Execute()

Get channel's email address



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
	streamId := int32(1) // int32 | The ID of the channel to access. 
	senderId := int32(1) // int32 | The ID of a user or bot which should appear as the sender when messages are sent to the channel using the returned channel email address.  `sender_id` can be:  - ID of the current user. - ID of the Email gateway bot. (Default value) - ID of a bot owned by the current user.  **Changes**: New in Zulip 10.0 (feature level 335).  Previously, the sender was always Email gateway bot.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStreamEmailAddress(context.Background(), streamId).SenderId(senderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetStreamEmailAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamEmailAddress`: GetStreamEmailAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetStreamEmailAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **senderId** | **int32** | The ID of a user or bot which should appear as the sender when messages are sent to the channel using the returned channel email address.  &#x60;sender_id&#x60; can be:  - ID of the current user. - ID of the Email gateway bot. (Default value) - ID of a bot owned by the current user.  **Changes**: New in Zulip 10.0 (feature level 335).  Previously, the sender was always Email gateway bot.  | 

### Return type

[**GetStreamEmailAddress200Response**](GetStreamEmailAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamId

> GetStreamId200Response GetStreamId(ctx).Stream(stream).Execute()

Get channel ID



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
	stream := "Denmark" // string | The name of the channel to access. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStreamId(context.Background()).Stream(stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetStreamId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamId`: GetStreamId200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetStreamId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stream** | **string** | The name of the channel to access.  | 

### Return type

[**GetStreamId200Response**](GetStreamId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamTopics

> GetStreamTopics200Response GetStreamTopics(ctx, streamId).AllowEmptyTopicName(allowEmptyTopicName).Execute()

Get topics in a channel



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
	streamId := int32(1) // int32 | The ID of the channel to access. 
	allowEmptyTopicName := true // bool | Whether the client supports processing the empty string as a topic name in the returned data.  If `false`, the value of `realm_empty_topic_display_name` found in the [`POST /register`](zulip.com/api/register-queue response is returned replacing the empty string as the topic name.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStreamTopics(context.Background(), streamId).AllowEmptyTopicName(allowEmptyTopicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetStreamTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamTopics`: GetStreamTopics200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetStreamTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowEmptyTopicName** | **bool** | Whether the client supports processing the empty string as a topic name in the returned data.  If &#x60;false&#x60;, the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response is returned replacing the empty string as the topic name.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  | [default to false]

### Return type

[**GetStreamTopics200Response**](GetStreamTopics200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreams

> GetStreams200Response GetStreams(ctx).IncludePublic(includePublic).IncludeWebPublic(includeWebPublic).IncludeSubscribed(includeSubscribed).ExcludeArchived(excludeArchived).IncludeAllActive(includeAllActive).IncludeAll(includeAll).IncludeDefault(includeDefault).IncludeOwnerSubscribed(includeOwnerSubscribed).IncludeCanAccessContent(includeCanAccessContent).Execute()

Get all channels



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
	includePublic := false // bool | Include all public channels.  (optional) (default to true)
	includeWebPublic := true // bool | Include all web-public channels.  (optional) (default to false)
	includeSubscribed := false // bool | Include all channels that the user is subscribed to.  (optional) (default to true)
	excludeArchived := true // bool | Whether to exclude archived streams from the results.  **Changes**: New in Zulip 10.0 (feature level 315).  (optional) (default to true)
	includeAllActive := true // bool | Deprecated parameter to include all channels. The user must have administrative privileges to use this parameter.  **Changes**: Deprecated in Zulip 10.0 (feature level 356). Clients interacting with newer servers should use the equivalent `include_all` parameter, which does not incorrectly hint that this parameter, and not `exclude_archived`, controls whether archived channels appear in the response.  (optional) (default to false)
	includeAll := true // bool | Include all channels that the user has metadata access to.  For organization administrators, this will be all channels in the organization, since organization administrators implicitly have metadata access to all channels.  **Changes**: New in Zulip 10.0 (feature level 356). On older versions, use `include_all_active`, which this replaces.  (optional) (default to false)
	includeDefault := true // bool | Include all default channels for the user's realm.  (optional) (default to false)
	includeOwnerSubscribed := true // bool | If the user is a bot, include all channels that the bot's owner is subscribed to.  (optional) (default to false)
	includeCanAccessContent := true // bool | Include all the channels that the user has content access to.  **Changes**: New in Zulip 10.0 (feature level 356).  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetStreams(context.Background()).IncludePublic(includePublic).IncludeWebPublic(includeWebPublic).IncludeSubscribed(includeSubscribed).ExcludeArchived(excludeArchived).IncludeAllActive(includeAllActive).IncludeAll(includeAll).IncludeDefault(includeDefault).IncludeOwnerSubscribed(includeOwnerSubscribed).IncludeCanAccessContent(includeCanAccessContent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreams`: GetStreams200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includePublic** | **bool** | Include all public channels.  | [default to true]
 **includeWebPublic** | **bool** | Include all web-public channels.  | [default to false]
 **includeSubscribed** | **bool** | Include all channels that the user is subscribed to.  | [default to true]
 **excludeArchived** | **bool** | Whether to exclude archived streams from the results.  **Changes**: New in Zulip 10.0 (feature level 315).  | [default to true]
 **includeAllActive** | **bool** | Deprecated parameter to include all channels. The user must have administrative privileges to use this parameter.  **Changes**: Deprecated in Zulip 10.0 (feature level 356). Clients interacting with newer servers should use the equivalent &#x60;include_all&#x60; parameter, which does not incorrectly hint that this parameter, and not &#x60;exclude_archived&#x60;, controls whether archived channels appear in the response.  | [default to false]
 **includeAll** | **bool** | Include all channels that the user has metadata access to.  For organization administrators, this will be all channels in the organization, since organization administrators implicitly have metadata access to all channels.  **Changes**: New in Zulip 10.0 (feature level 356). On older versions, use &#x60;include_all_active&#x60;, which this replaces.  | [default to false]
 **includeDefault** | **bool** | Include all default channels for the user&#39;s realm.  | [default to false]
 **includeOwnerSubscribed** | **bool** | If the user is a bot, include all channels that the bot&#39;s owner is subscribed to.  | [default to false]
 **includeCanAccessContent** | **bool** | Include all the channels that the user has content access to.  **Changes**: New in Zulip 10.0 (feature level 356).  | [default to false]

### Return type

[**GetStreams200Response**](GetStreams200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribers

> GetSubscribers200Response GetSubscribers(ctx, streamId).Execute()

Get channel subscribers



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
	streamId := int32(1) // int32 | The ID of the channel to access. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSubscribers(context.Background(), streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscribers`: GetSubscribers200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSubscribers200Response**](GetSubscribers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionStatus

> GetSubscriptionStatus200Response GetSubscriptionStatus(ctx, userId, streamId).Execute()

Get subscription status



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
	userId := int32(12) // int32 | The target user's ID. 
	streamId := int32(1) // int32 | The ID of the channel to access. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSubscriptionStatus(context.Background(), userId, streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetSubscriptionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionStatus`: GetSubscriptionStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetSubscriptionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The target user&#39;s ID.  | 
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSubscriptionStatus200Response**](GetSubscriptionStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> GetSubscriptions200Response GetSubscriptions(ctx).IncludeSubscribers(includeSubscribers).Execute()

Get subscribed channels



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
	includeSubscribers := "true" // string | Whether each returned channel object should include a `subscribers` field containing a list of the user IDs of its subscribers.  Client apps supporting organizations with many thousands of users should not pass `true`, because the full subscriber matrix may be several megabytes of data. The `partial` value, combined with the `subscriber_count` and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes `partial` for this parameter, the server may, for some channels, return a subset of the channel's subscribers in the `partial_subscribers` field instead of the `subscribers` field, which always contains the complete set of subscribers.  The server guarantees that it will always return a `subscribers` field for channels with fewer than 250 total subscribers. When returning a `partial_subscribers` field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  **Changes**: The `partial` value is new in Zulip 11.0 (feature level 412).  New in Zulip 2.1.0.  (optional) (default to "false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSubscriptions(context.Background()).IncludeSubscribers(includeSubscribers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: GetSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSubscribers** | **string** | Whether each returned channel object should include a &#x60;subscribers&#x60; field containing a list of the user IDs of its subscribers.  Client apps supporting organizations with many thousands of users should not pass &#x60;true&#x60;, because the full subscriber matrix may be several megabytes of data. The &#x60;partial&#x60; value, combined with the &#x60;subscriber_count&#x60; and fetching subscribers for individual channels as needed, is recommended to support client app features where channel subscriber data is useful.  If a client passes &#x60;partial&#x60; for this parameter, the server may, for some channels, return a subset of the channel&#39;s subscribers in the &#x60;partial_subscribers&#x60; field instead of the &#x60;subscribers&#x60; field, which always contains the complete set of subscribers.  The server guarantees that it will always return a &#x60;subscribers&#x60; field for channels with fewer than 250 total subscribers. When returning a &#x60;partial_subscribers&#x60; field, the server guarantees that all bot users and users active within the last 14 days will be included. For other cases, the server may use its discretion to determine which channels and users to include, balancing between payload size and usefulness of the data provided to the client.  **Changes**: The &#x60;partial&#x60; value is new in Zulip 11.0 (feature level 412).  New in Zulip 2.1.0.  | [default to &quot;false&quot;]

### Return type

[**GetSubscriptions200Response**](GetSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteTopic

> JsonSuccess MuteTopic(ctx).Topic(topic).Op(op).StreamId(streamId).Stream(stream).Execute()

Topic muting



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
	topic := "topic_example" // string | The topic to (un)mute. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the `max_topic_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum topic length. 
	op := "op_example" // string | Whether to mute (`add`) or unmute (`remove`) the provided topic. 
	streamId := int32(56) // int32 | The ID of the channel to access.  Clients must provide either `stream` or `stream_id` as a parameter to this endpoint, but not both.  **Changes**: New in Zulip 2.0.0.  (optional)
	stream := "stream_example" // string | The name of the channel to access.  Clients must provide either `stream` or `stream_id` as a parameter to this endpoint, but not both. Clients should use `stream_id` instead of the `stream` parameter when possible.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteTopic(context.Background()).Topic(topic).Op(op).StreamId(streamId).Stream(stream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.MuteTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteTopic`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.MuteTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMuteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string** | The topic to (un)mute. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum topic length.  | 
 **op** | **string** | Whether to mute (&#x60;add&#x60;) or unmute (&#x60;remove&#x60;) the provided topic.  | 
 **streamId** | **int32** | The ID of the channel to access.  Clients must provide either &#x60;stream&#x60; or &#x60;stream_id&#x60; as a parameter to this endpoint, but not both.  **Changes**: New in Zulip 2.0.0.  | 
 **stream** | **string** | The name of the channel to access.  Clients must provide either &#x60;stream&#x60; or &#x60;stream_id&#x60; as a parameter to this endpoint, but not both. Clients should use &#x60;stream_id&#x60; instead of the &#x60;stream&#x60; parameter when possible.  | 

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


## PatchChannelFolders

> JsonSuccess PatchChannelFolders(ctx).Order(order).Execute()

Reorder channel folders



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
	order := []int32{int32(123)} // []int32 | A list of channel folder IDs representing the new order.  This list must include the IDs of all the organization's channel folders, including archived folders.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchChannelFolders(context.Background()).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PatchChannelFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchChannelFolders`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PatchChannelFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchChannelFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **[]int32** | A list of channel folder IDs representing the new order.  This list must include the IDs of all the organization&#39;s channel folders, including archived folders.  | 

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


## RemoveDefaultStream

> JsonSuccess RemoveDefaultStream(ctx).StreamId(streamId).Execute()

Remove a default channel



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
	streamId := int32(56) // int32 | The ID of the target channel. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveDefaultStream(context.Background()).StreamId(streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.RemoveDefaultStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDefaultStream`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.RemoveDefaultStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDefaultStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **int32** | The ID of the target channel.  | 

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


## Subscribe

> Subscribe200Response Subscribe(ctx).Subscriptions(subscriptions).Principals(principals).AuthorizationErrorsFatal(authorizationErrorsFatal).Announce(announce).InviteOnly(inviteOnly).IsWebPublic(isWebPublic).IsDefaultStream(isDefaultStream).HistoryPublicToSubscribers(historyPublicToSubscribers).MessageRetentionDays(messageRetentionDays).TopicsPolicy(topicsPolicy).CanAddSubscribersGroup(canAddSubscribersGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).FolderId(folderId).SendNewSubscriptionMessages(sendNewSubscriptionMessages).Execute()

Subscribe to a channel



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
	subscriptions := []openapiclient.SubscribeRequestSubscriptionsInner{*openapiclient.NewSubscribeRequestSubscriptionsInner("Name_example")} // []SubscribeRequestSubscriptionsInner | A list of dictionaries containing the key `name` and value specifying the name of the channel to subscribe. If the channel does not exist a new channel is created. The description of the channel created can be specified by setting the dictionary key `description` with an appropriate value. 
	principals := openapiclient.Principals{ArrayOfInt32: new([]int32)} // Principals |  (optional)
	authorizationErrorsFatal := true // bool | A boolean specifying whether authorization errors (such as when the requesting user is not authorized to access a private channel) should be considered fatal or not. When `true`, an authorization error is reported as such. When set to `false`, the response will be a 200 and any channels where the request encountered an authorization error will be listed in the `unauthorized` key.  (optional) (default to true)
	announce := true // bool | If one of the channels specified did not exist previously and is thus created by this call, this determines whether [notification bot](zulip.com/help/configure-automated-notices will send an announcement about the new channel's creation.  (optional) (default to false)
	inviteOnly := true // bool | As described above, this endpoint will create a new channel if passed a channel name that doesn't already exist. This parameters and the ones that follow are used to request an initial configuration of a created channel; they are ignored for channels that already exist.  This parameter determines whether any newly created channels will be private channels.  (optional) (default to false)
	isWebPublic := true // bool | This parameter determines whether any newly created channels will be web-public channels.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED` [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current use to have permission under the organization's `can_create_web_public_channel_group` realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).  (optional) (default to false)
	isDefaultStream := true // bool | This parameter determines whether any newly created channels will be added as [default channels][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.  (optional) (default to false)
	historyPublicToSubscribers := true // bool | Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  (optional)
	messageRetentionDays := openapiclient.MessageRetentionDays{Int32: new(int32)} // MessageRetentionDays |  (optional)
	topicsPolicy := openapiclient.TopicsPolicy("inherit") // TopicsPolicy |  (optional)
	canAddSubscribersGroup := *openapiclient.NewChannelCanAddSubscribersGroup() // ChannelCanAddSubscribersGroup |  (optional)
	canRemoveSubscribersGroup := *openapiclient.NewCanRemoveSubscribersGroup() // CanRemoveSubscribersGroup |  (optional)
	canAdministerChannelGroup := *openapiclient.NewCanAdministerChannelGroup() // CanAdministerChannelGroup |  (optional)
	canDeleteAnyMessageGroup := *openapiclient.NewCanDeleteAnyMessageGroup() // CanDeleteAnyMessageGroup |  (optional)
	canDeleteOwnMessageGroup := *openapiclient.NewCanDeleteOwnMessageGroup() // CanDeleteOwnMessageGroup |  (optional)
	canMoveMessagesOutOfChannelGroup := *openapiclient.NewCanMoveMessagesOutOfChannelGroup() // CanMoveMessagesOutOfChannelGroup |  (optional)
	canMoveMessagesWithinChannelGroup := *openapiclient.NewCanMoveMessagesWithinChannelGroup() // CanMoveMessagesWithinChannelGroup |  (optional)
	canSendMessageGroup := *openapiclient.NewCanSendMessageGroup() // CanSendMessageGroup |  (optional)
	canSubscribeGroup := *openapiclient.NewCanSubscribeGroup() // CanSubscribeGroup |  (optional)
	canResolveTopicsGroup := *openapiclient.NewCanResolveTopicsGroup() // CanResolveTopicsGroup |  (optional)
	folderId := int32(56) // int32 | This parameter determines the folder to which the newly created channel will be added.  If the value is `None`, the channel will not be added to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  (optional)
	sendNewSubscriptionMessages := true // bool | Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than `max_bulk_new_subscription_messages` (available in the [`POST /register`](zulip.com/api/register-queue response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscribe(context.Background()).Subscriptions(subscriptions).Principals(principals).AuthorizationErrorsFatal(authorizationErrorsFatal).Announce(announce).InviteOnly(inviteOnly).IsWebPublic(isWebPublic).IsDefaultStream(isDefaultStream).HistoryPublicToSubscribers(historyPublicToSubscribers).MessageRetentionDays(messageRetentionDays).TopicsPolicy(topicsPolicy).CanAddSubscribersGroup(canAddSubscribersGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).FolderId(folderId).SendNewSubscriptionMessages(sendNewSubscriptionMessages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.Subscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Subscribe`: Subscribe200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.Subscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptions** | [**[]SubscribeRequestSubscriptionsInner**](SubscribeRequestSubscriptionsInner.md) | A list of dictionaries containing the key &#x60;name&#x60; and value specifying the name of the channel to subscribe. If the channel does not exist a new channel is created. The description of the channel created can be specified by setting the dictionary key &#x60;description&#x60; with an appropriate value.  | 
 **principals** | [**Principals**](Principals.md) |  | 
 **authorizationErrorsFatal** | **bool** | A boolean specifying whether authorization errors (such as when the requesting user is not authorized to access a private channel) should be considered fatal or not. When &#x60;true&#x60;, an authorization error is reported as such. When set to &#x60;false&#x60;, the response will be a 200 and any channels where the request encountered an authorization error will be listed in the &#x60;unauthorized&#x60; key.  | [default to true]
 **announce** | **bool** | If one of the channels specified did not exist previously and is thus created by this call, this determines whether [notification bot](zulip.com/help/configure-automated-notices will send an announcement about the new channel&#39;s creation.  | [default to false]
 **inviteOnly** | **bool** | As described above, this endpoint will create a new channel if passed a channel name that doesn&#39;t already exist. This parameters and the ones that follow are used to request an initial configuration of a created channel; they are ignored for channels that already exist.  This parameter determines whether any newly created channels will be private channels.  | [default to false]
 **isWebPublic** | **bool** | This parameter determines whether any newly created channels will be web-public channels.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current use to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).  | [default to false]
 **isDefaultStream** | **bool** | This parameter determines whether any newly created channels will be added as [default channels][default-channels] for new users joining the organization.  [default-channels]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.  | [default to false]
 **historyPublicToSubscribers** | **bool** | Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  | 
 **messageRetentionDays** | [**MessageRetentionDays**](MessageRetentionDays.md) |  | 
 **topicsPolicy** | [**TopicsPolicy**](TopicsPolicy.md) |  | 
 **canAddSubscribersGroup** | [**ChannelCanAddSubscribersGroup**](ChannelCanAddSubscribersGroup.md) |  | 
 **canRemoveSubscribersGroup** | [**CanRemoveSubscribersGroup**](CanRemoveSubscribersGroup.md) |  | 
 **canAdministerChannelGroup** | [**CanAdministerChannelGroup**](CanAdministerChannelGroup.md) |  | 
 **canDeleteAnyMessageGroup** | [**CanDeleteAnyMessageGroup**](CanDeleteAnyMessageGroup.md) |  | 
 **canDeleteOwnMessageGroup** | [**CanDeleteOwnMessageGroup**](CanDeleteOwnMessageGroup.md) |  | 
 **canMoveMessagesOutOfChannelGroup** | [**CanMoveMessagesOutOfChannelGroup**](CanMoveMessagesOutOfChannelGroup.md) |  | 
 **canMoveMessagesWithinChannelGroup** | [**CanMoveMessagesWithinChannelGroup**](CanMoveMessagesWithinChannelGroup.md) |  | 
 **canSendMessageGroup** | [**CanSendMessageGroup**](CanSendMessageGroup.md) |  | 
 **canSubscribeGroup** | [**CanSubscribeGroup**](CanSubscribeGroup.md) |  | 
 **canResolveTopicsGroup** | [**CanResolveTopicsGroup**](CanResolveTopicsGroup.md) |  | 
 **folderId** | **int32** | This parameter determines the folder to which the newly created channel will be added.  If the value is &#x60;None&#x60;, the channel will not be added to any folder.  **Changes**: New in Zulip 11.0 (feature level 389).  | 
 **sendNewSubscriptionMessages** | **bool** | Whether any other users newly subscribed via this request should be sent a Notification Bot DM notifying them about their new subscription.  The server will never send Notification Bot DMs if more than &#x60;max_bulk_new_subscription_messages&#x60; (available in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response) users were subscribed in this request.  **Changes**: Before Zulip 11.0 (feature level 397), new subscribers were always sent a Notification Bot DM, which was unduly expensive when bulk-subscribing thousands of users to a channel.  | [default to true]

### Return type

[**Subscribe200Response**](Subscribe200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unsubscribe

> Unsubscribe200Response Unsubscribe(ctx).Subscriptions(subscriptions).Principals(principals).Execute()

Unsubscribe from a channel



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
	subscriptions := []string{"Inner_example"} // []string | A list of channel names to unsubscribe from. This parameter is called `streams` in our Python API. 
	principals := openapiclient.Principals{ArrayOfInt32: new([]int32)} // Principals |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Unsubscribe(context.Background()).Subscriptions(subscriptions).Principals(principals).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.Unsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unsubscribe`: Unsubscribe200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.Unsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptions** | **[]string** | A list of channel names to unsubscribe from. This parameter is called &#x60;streams&#x60; in our Python API.  | 
 **principals** | [**Principals**](Principals.md) |  | 

### Return type

[**Unsubscribe200Response**](Unsubscribe200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelFolder

> JsonSuccess UpdateChannelFolder(ctx, channelFolderId).Name(name).Description(description).IsArchived(isArchived).Execute()

Update a channel folder



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
	channelFolderId := int32(1) // int32 | The ID of the target channel folder. 
	name := "name_example" // string | The new name of the channel folder.  Clients should use the `max_channel_folder_name_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel folder name length.  (optional)
	description := "description_example" // string | The new description of the channel folder.  Clients should use the `max_channel_folder_description_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel folder description length.  (optional)
	isArchived := true // bool | Whether to archive or unarchive the channel folder.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateChannelFolder(context.Background(), channelFolderId).Name(name).Description(description).IsArchived(isArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateChannelFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChannelFolder`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateChannelFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelFolderId** | **int32** | The ID of the target channel folder.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The new name of the channel folder.  Clients should use the &#x60;max_channel_folder_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel folder name length.  | 
 **description** | **string** | The new description of the channel folder.  Clients should use the &#x60;max_channel_folder_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel folder description length.  | 
 **isArchived** | **bool** | Whether to archive or unarchive the channel folder.  | 

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


## UpdateStream

> JsonSuccess UpdateStream(ctx, streamId).Description(description).NewName(newName).IsPrivate(isPrivate).IsWebPublic(isWebPublic).HistoryPublicToSubscribers(historyPublicToSubscribers).IsDefaultStream(isDefaultStream).MessageRetentionDays(messageRetentionDays).IsArchived(isArchived).FolderId(folderId).TopicsPolicy(topicsPolicy).CanAddSubscribersGroup(canAddSubscribersGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).Execute()

Update a channel



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
	streamId := int32(1) // int32 | The ID of the channel to access. 
	description := "description_example" // string | The new [description](zulip.com/help/change-the-channel-description for the channel, in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should use the `max_stream_description_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel description length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).  (optional)
	newName := "newName_example" // string | The new name for the channel.  Clients should use the `max_stream_name_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum channel name length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).  (optional)
	isPrivate := true // bool | Change whether the channel is a private channel.  (optional)
	isWebPublic := true // bool | Change whether the channel is a web-public channel.  Note that creating web-public channels requires the `WEB_PUBLIC_STREAMS_ENABLED` [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the `enable_spectator_access` realm setting, and the current use to have permission under the organization's `can_create_web_public_channel_group` realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).  (optional)
	historyPublicToSubscribers := true // bool | Whether the channel's message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  It's an error for this parameter to be false for a public or web-public channel and when is_private is false.  **Changes**: Before Zulip 6.0 (feature level 136), `history_public_to_subscribers` was silently ignored unless the request also contained either `is_private` or `is_web_public`.  (optional)
	isDefaultStream := true // bool | Add or remove the channel as a [default channel][default-channel] for new users joining the organization.  [default-channel]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.  (optional)
	messageRetentionDays := openapiclient.MessageRetentionDays{Int32: new(int32)} // MessageRetentionDays |  (optional)
	isArchived := true // bool | A boolean indicating whether the channel is [archived](zulip.com/help/archive-a-channel or unarchived. Currently only allows unarchiving previously archived channels.  **Changes**: New in Zulip 11.0 (feature level 388).  (optional)
	folderId := int32(56) // int32 | ID of the new folder to which the channel should belong.  It can be `None` if the user wants to just remove the channel from its existing folder.  **Changes**: New in Zulip 11.0 (feature level 389).  (optional)
	topicsPolicy := openapiclient.TopicsPolicy("inherit") // TopicsPolicy |  (optional)
	canAddSubscribersGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to add subscribers to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who can administer the channel or have similar realm-level permissions can add subscribers to a public channel regardless of the value of this setting.  Users in this group need not be subscribed to a private channel to add subscribers to it.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 342). Previously, there was no channel-level setting for this permission.  (optional)
	canRemoveSubscribersGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to unsubscribe others from this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can unsubscribe others from a channel as though they were in this group without being explicitly listed here.  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: Prior to Zulip 10.0 (feature level 349), channel administrators could not unsubscribe other users if they were not an organization administrator or part of `can_remove_subscribers_group`. Realm administrators were not allowed to unsubscribe other users from a private channel if they were not subscribed to that channel.  Prior to Zulip 10.0 (feature level 320), this value was always the integer ID of a system group.  Before Zulip 8.0 (feature level 197), the `can_remove_subscribers_group` setting was named `can_remove_subscribers_group_id`.  New in Zulip 7.0 (feature level 161).  (optional)
	canAdministerChannelGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to administer this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can administer every channel as though they were in this group without being explicitly listed here.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to add other subscribers to the channel.  **Changes**: Prior to Zulip 10.0 (feature level 349) a user needed to [have content access](zulip.com/help/channel-permissions to a channel in order to modify it. The exception to this rule was that organization administrators can edit channel names and descriptions without having full access to the channel.  New in Zulip 10.0 (feature level 325). Prior to this change, the permission to administer channels was limited to realm administrators.  (optional)
	canDeleteAnyMessageGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to delete any message in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to delete any message in the channel.  Users present in the organization-level `can_delete_any_message_group` setting can always delete any message in the channel if they [have content access](zulip.com/help/channel-permissions to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in `can_delete_any_message_group` were able delete any message in the organization.  (optional)
	canDeleteOwnMessageGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to delete the messages that they have sent in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to delete their own message in the channel.  Users with permission to delete any message in the channel and users present in the organization-level `can_delete_own_message_group` setting can always delete their own messages in the channel if they [have content access](zulip.com/help/channel-permissions to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in the organization-level `can_delete_any_message_group` and `can_delete_own_message_group` settings were able delete their own messages in the organization.  (optional)
	canMoveMessagesOutOfChannelGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to move messages out of this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to move messages out of the channel.  Channel administrators and users present in the organization-level `can_move_messages_between_channels_group` setting can always move messages out of the channel if they [have content access](zulip.com/help/channel-permissions to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in `can_move_messages_between_channels_group` were able move messages between channels.  (optional)
	canMoveMessagesWithinChannelGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to move messages within this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to move messages within the channel.  Channel administrators and users present in the organization-level `can_move_messages_between_topics_group` setting can always move messages within the channel if they [have content access](zulip.com/help/channel-permissions to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in `can_move_messages_between_topics_group` were able move messages between topics of a channel.  (optional)
	canSendMessageGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to post in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 333). Previously `stream_post_policy` field used to control the permission to post in the channel.  (optional)
	canSubscribeGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to subscribe themselves to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Everyone, excluding guests, can subscribe to any public channel irrespective of this setting.  Users in this group can subscribe to a private channel as well.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 357).  (optional)
	canResolveTopicsGroup := *openapiclient.NewGroupSettingValueUpdate(openapiclient.GroupSettingValue{GroupSettingValueOneOf: openapiclient.NewGroupSettingValueOneOf()}) // GroupSettingValueUpdate | The set of users who have permission to to resolve topics in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who have similar realm-level permissions can resolve topics in a channel regardless of the value of this setting.  **Changes**: New in Zulip 11.0 (feature level 402).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateStream(context.Background(), streamId).Description(description).NewName(newName).IsPrivate(isPrivate).IsWebPublic(isWebPublic).HistoryPublicToSubscribers(historyPublicToSubscribers).IsDefaultStream(isDefaultStream).MessageRetentionDays(messageRetentionDays).IsArchived(isArchived).FolderId(folderId).TopicsPolicy(topicsPolicy).CanAddSubscribersGroup(canAddSubscribersGroup).CanRemoveSubscribersGroup(canRemoveSubscribersGroup).CanAdministerChannelGroup(canAdministerChannelGroup).CanDeleteAnyMessageGroup(canDeleteAnyMessageGroup).CanDeleteOwnMessageGroup(canDeleteOwnMessageGroup).CanMoveMessagesOutOfChannelGroup(canMoveMessagesOutOfChannelGroup).CanMoveMessagesWithinChannelGroup(canMoveMessagesWithinChannelGroup).CanSendMessageGroup(canSendMessageGroup).CanSubscribeGroup(canSubscribeGroup).CanResolveTopicsGroup(canResolveTopicsGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStream`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **int32** | The ID of the channel to access.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** | The new [description](zulip.com/help/change-the-channel-description for the channel, in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should use the &#x60;max_stream_description_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel description length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).  | 
 **newName** | **string** | The new name for the channel.  Clients should use the &#x60;max_stream_name_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum channel name length.  **Changes**: Removed unnecessary JSON-encoding of this parameter in Zulip 4.0 (feature level 64).  | 
 **isPrivate** | **bool** | Change whether the channel is a private channel.  | 
 **isWebPublic** | **bool** | Change whether the channel is a web-public channel.  Note that creating web-public channels requires the &#x60;WEB_PUBLIC_STREAMS_ENABLED&#x60; [server setting][server-settings] to be enabled on the Zulip server in question, the organization to have enabled the &#x60;enable_spectator_access&#x60; realm setting, and the current use to have permission under the organization&#39;s &#x60;can_create_web_public_channel_group&#x60; realm setting.  [server-settings]: https://zulip.readthedocs.io/en/stable/production/settings.html  **Changes**: New in Zulip 5.0 (feature level 98).  | 
 **historyPublicToSubscribers** | **bool** | Whether the channel&#39;s message history should be available to newly subscribed members, or users can only access messages they actually received while subscribed to the channel.  Corresponds to the shared history option for [private channels](zulip.com/help/channel-permissions#private-channels.  It&#39;s an error for this parameter to be false for a public or web-public channel and when is_private is false.  **Changes**: Before Zulip 6.0 (feature level 136), &#x60;history_public_to_subscribers&#x60; was silently ignored unless the request also contained either &#x60;is_private&#x60; or &#x60;is_web_public&#x60;.  | 
 **isDefaultStream** | **bool** | Add or remove the channel as a [default channel][default-channel] for new users joining the organization.  [default-channel]: /help/set-default-channels-for-new-users  **Changes**: New in Zulip 8.0 (feature level 200). Previously, default channel status could only be changed using the [dedicated API endpoint](zulip.com/api/add-default-stream.  | 
 **messageRetentionDays** | [**MessageRetentionDays**](MessageRetentionDays.md) |  | 
 **isArchived** | **bool** | A boolean indicating whether the channel is [archived](zulip.com/help/archive-a-channel or unarchived. Currently only allows unarchiving previously archived channels.  **Changes**: New in Zulip 11.0 (feature level 388).  | 
 **folderId** | **int32** | ID of the new folder to which the channel should belong.  It can be &#x60;None&#x60; if the user wants to just remove the channel from its existing folder.  **Changes**: New in Zulip 11.0 (feature level 389).  | 
 **topicsPolicy** | [**TopicsPolicy**](TopicsPolicy.md) |  | 
 **canAddSubscribersGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to add subscribers to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who can administer the channel or have similar realm-level permissions can add subscribers to a public channel regardless of the value of this setting.  Users in this group need not be subscribed to a private channel to add subscribers to it.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 342). Previously, there was no channel-level setting for this permission.  | 
 **canRemoveSubscribersGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to unsubscribe others from this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can unsubscribe others from a channel as though they were in this group without being explicitly listed here.  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: Prior to Zulip 10.0 (feature level 349), channel administrators could not unsubscribe other users if they were not an organization administrator or part of &#x60;can_remove_subscribers_group&#x60;. Realm administrators were not allowed to unsubscribe other users from a private channel if they were not subscribed to that channel.  Prior to Zulip 10.0 (feature level 320), this value was always the integer ID of a system group.  Before Zulip 8.0 (feature level 197), the &#x60;can_remove_subscribers_group&#x60; setting was named &#x60;can_remove_subscribers_group_id&#x60;.  New in Zulip 7.0 (feature level 161).  | 
 **canAdministerChannelGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to administer this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Organization administrators can administer every channel as though they were in this group without being explicitly listed here.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to add other subscribers to the channel.  **Changes**: Prior to Zulip 10.0 (feature level 349) a user needed to [have content access](zulip.com/help/channel-permissions to a channel in order to modify it. The exception to this rule was that organization administrators can edit channel names and descriptions without having full access to the channel.  New in Zulip 10.0 (feature level 325). Prior to this change, the permission to administer channels was limited to realm administrators.  | 
 **canDeleteAnyMessageGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to delete any message in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to delete any message in the channel.  Users present in the organization-level &#x60;can_delete_any_message_group&#x60; setting can always delete any message in the channel if they [have content access](zulip.com/help/channel-permissions to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in &#x60;can_delete_any_message_group&#x60; were able delete any message in the organization.  | 
 **canDeleteOwnMessageGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to delete the messages that they have sent in the channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to delete their own message in the channel.  Users with permission to delete any message in the channel and users present in the organization-level &#x60;can_delete_own_message_group&#x60; setting can always delete their own messages in the channel if they [have content access](zulip.com/help/channel-permissions to that channel.  **Changes**: New in Zulip 11.0 (feature level 407). Prior to this change, only the users in the organization-level &#x60;can_delete_any_message_group&#x60; and &#x60;can_delete_own_message_group&#x60; settings were able delete their own messages in the organization.  | 
 **canMoveMessagesOutOfChannelGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to move messages out of this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to move messages out of the channel.  Channel administrators and users present in the organization-level &#x60;can_move_messages_between_channels_group&#x60; setting can always move messages out of the channel if they [have content access](zulip.com/help/channel-permissions to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in &#x60;can_move_messages_between_channels_group&#x60; were able move messages between channels.  | 
 **canMoveMessagesWithinChannelGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to move messages within this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel in order to move messages within the channel.  Channel administrators and users present in the organization-level &#x60;can_move_messages_between_topics_group&#x60; setting can always move messages within the channel if they [have content access](zulip.com/help/channel-permissions to the channel.  **Changes**: New in Zulip 11.0 (feature level 396). Prior to this change, only the users in &#x60;can_move_messages_between_topics_group&#x60; were able move messages between topics of a channel.  | 
 **canSendMessageGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to post in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Note that a user must have metadata access to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 333). Previously &#x60;stream_post_policy&#x60; field used to control the permission to post in the channel.  | 
 **canSubscribeGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to subscribe themselves to this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Everyone, excluding guests, can subscribe to any public channel irrespective of this setting.  Users in this group can subscribe to a private channel as well.  Note that a user must [have content access](zulip.com/help/channel-permissions to a channel and permission to administer the channel in order to modify this setting.  **Changes**: New in Zulip 10.0 (feature level 357).  | 
 **canResolveTopicsGroup** | [**GroupSettingValueUpdate**](GroupSettingValueUpdate.md) | The set of users who have permission to to resolve topics in this channel expressed as an [update to a group-setting value][update-group-setting].  [update-group-setting]: /api/group-setting-values#updating-group-setting-values  Users who have similar realm-level permissions can resolve topics in a channel regardless of the value of this setting.  **Changes**: New in Zulip 11.0 (feature level 402).  | 

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


## UpdateSubscriptionSettings

> IgnoredParametersSuccess UpdateSubscriptionSettings(ctx).SubscriptionData(subscriptionData).Execute()

Update subscription settings



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
	subscriptionData := []openapiclient.UpdateSubscriptionSetting{*openapiclient.NewUpdateSubscriptionSetting(int32(123), "Property_example", openapiclient.update_subscription_setting_value{Bool: new(bool)})} // []UpdateSubscriptionSetting | A list of objects that describe the changes that should be applied in each subscription. Each object represents a subscription, and must have a `stream_id` key that identifies the channel, as well as the `property` being modified and its new `value`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateSubscriptionSettings(context.Background()).SubscriptionData(subscriptionData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateSubscriptionSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscriptionSettings`: IgnoredParametersSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateSubscriptionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionData** | [**[]UpdateSubscriptionSetting**](UpdateSubscriptionSetting.md) | A list of objects that describe the changes that should be applied in each subscription. Each object represents a subscription, and must have a &#x60;stream_id&#x60; key that identifies the channel, as well as the &#x60;property&#x60; being modified and its new &#x60;value&#x60;.  | 

### Return type

[**IgnoredParametersSuccess**](IgnoredParametersSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscriptions

> UpdateSubscriptions200Response UpdateSubscriptions(ctx).Delete(delete).Add(add).Execute()

Update subscriptions



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
	delete := []string{"Inner_example"} // []string | A list of channel names to unsubscribe from.  (optional)
	add := []openapiclient.UpdateSubscriptionsRequestAddInner{*openapiclient.NewUpdateSubscriptionsRequestAddInner()} // []UpdateSubscriptionsRequestAddInner | A list of objects describing which channels to subscribe to, optionally including per-user subscription parameters (e.g. color) and if the channel is to be created, its description.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateSubscriptions(context.Background()).Delete(delete).Add(add).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscriptions`: UpdateSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delete** | **[]string** | A list of channel names to unsubscribe from.  | 
 **add** | [**[]UpdateSubscriptionsRequestAddInner**](UpdateSubscriptionsRequestAddInner.md) | A list of objects describing which channels to subscribe to, optionally including per-user subscription parameters (e.g. color) and if the channel is to be created, its description.  | 

### Return type

[**UpdateSubscriptions200Response**](UpdateSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserTopic

> JsonSuccess UpdateUserTopic(ctx).StreamId(streamId).Topic(topic).VisibilityPolicy(visibilityPolicy).Execute()

Update personal preferences for a topic



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
	streamId := int32(56) // int32 | The ID of the channel to access. 
	topic := "topic_example" // string | The topic for which the personal preferences needs to be updated. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the `max_topic_length` returned by the [`POST /register`](zulip.com/api/register-queue endpoint to determine the maximum topic length.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register](zulip.com/api/register-queue response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages. 
	visibilityPolicy := int32(56) // int32 | Controls which visibility policy to set.  - 0 = None. Removes the visibility policy previously set for the topic. - 1 = Muted. [Mutes the topic](zulip.com/help/mute-a-topic in a channel. - 2 = Unmuted. [Unmutes the topic](zulip.com/help/mute-a-topic in a muted channel. - 3 = Followed. [Follows the topic](zulip.com/help/follow-a-topic.  In an unmuted channel, a topic visibility policy of unmuted will have the same effect as the \\\"None\\\" visibility policy.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateUserTopic(context.Background()).StreamId(streamId).Topic(topic).VisibilityPolicy(visibilityPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.UpdateUserTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserTopic`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.UpdateUserTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **int32** | The ID of the channel to access.  | 
 **topic** | **string** | The topic for which the personal preferences needs to be updated. Note that the request will succeed regardless of whether any messages have been sent to the specified topic.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](zulip.com/api/register-queue endpoint to determine the maximum topic length.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](zulip.com/api/register-queue response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  | 
 **visibilityPolicy** | **int32** | Controls which visibility policy to set.  - 0 &#x3D; None. Removes the visibility policy previously set for the topic. - 1 &#x3D; Muted. [Mutes the topic](zulip.com/help/mute-a-topic in a channel. - 2 &#x3D; Unmuted. [Unmutes the topic](zulip.com/help/mute-a-topic in a muted channel. - 3 &#x3D; Followed. [Follows the topic](zulip.com/help/follow-a-topic.  In an unmuted channel, a topic visibility policy of unmuted will have the same effect as the \\\&quot;None\\\&quot; visibility policy.  **Changes**: In Zulip 7.0 (feature level 219), added followed as a visibility policy option.  | 

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

