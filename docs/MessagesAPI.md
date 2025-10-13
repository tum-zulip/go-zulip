# \MessagesAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReaction**](MessagesAPI.md#AddReaction) | **Post** /messages/{message_id}/reactions | Add an emoji reaction
[**CheckMessagesMatchNarrow**](MessagesAPI.md#CheckMessagesMatchNarrow) | **Get** /messages/matches_narrow | Check if messages match a narrow
[**DeleteMessage**](MessagesAPI.md#DeleteMessage) | **Delete** /messages/{message_id} | Delete a message
[**GetFileTemporaryUrl**](MessagesAPI.md#GetFileTemporaryUrl) | **Get** /user_uploads/{realm_id_str}/{filename} | Get public temporary URL
[**GetMessage**](MessagesAPI.md#GetMessage) | **Get** /messages/{message_id} | Fetch a single message
[**GetMessageHistory**](MessagesAPI.md#GetMessageHistory) | **Get** /messages/{message_id}/history | Get a message&#39;s edit history
[**GetMessages**](MessagesAPI.md#GetMessages) | **Get** /messages | Get messages
[**GetReadReceipts**](MessagesAPI.md#GetReadReceipts) | **Get** /messages/{message_id}/read_receipts | Get a message&#39;s read receipts
[**MarkAllAsRead**](MessagesAPI.md#MarkAllAsRead) | **Post** /mark_all_as_read | Mark all messages as read
[**MarkStreamAsRead**](MessagesAPI.md#MarkStreamAsRead) | **Post** /mark_stream_as_read | Mark messages in a channel as read
[**MarkTopicAsRead**](MessagesAPI.md#MarkTopicAsRead) | **Post** /mark_topic_as_read | Mark messages in a topic as read
[**RemoveReaction**](MessagesAPI.md#RemoveReaction) | **Delete** /messages/{message_id}/reactions | Remove an emoji reaction
[**RenderMessage**](MessagesAPI.md#RenderMessage) | **Post** /messages/render | Render a message
[**ReportMessage**](MessagesAPI.md#ReportMessage) | **Post** /messages/{message_id}/report | Report a message
[**SendMessage**](MessagesAPI.md#SendMessage) | **Post** /messages | Send a message
[**UpdateMessage**](MessagesAPI.md#UpdateMessage) | **Patch** /messages/{message_id} | Edit a message
[**UpdateMessageFlags**](MessagesAPI.md#UpdateMessageFlags) | **Post** /messages/flags | Update personal message flags
[**UpdateMessageFlagsForNarrow**](MessagesAPI.md#UpdateMessageFlagsForNarrow) | **Post** /messages/flags/narrow | Update personal message flags for narrow
[**UploadFile**](MessagesAPI.md#UploadFile) | **Post** /user_uploads | Upload a file



## AddReaction

> JsonSuccess AddReaction(ctx, messageId).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()

Add an emoji reaction



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
	messageId := int32(43) // int32 | The target message's ID. 
	emojiName := "emojiName_example" // string | The target emoji's human-readable name.  To find an emoji's name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji's text name. 
	emojiCode := "emojiCode_example" // string | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  For most API clients, you won't need this, but it's important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the `emoji_code` in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a \\\"different\\\" emoji.  (optional)
	reactionType := "reactionType_example" // string | A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity. Supported values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \\\"zulip\\\").  **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji](/help/custom-emoji); previously, this endpoint assumed `unicode_emoji` if this parameter was not specified.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddReaction(context.Background(), messageId).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.AddReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddReaction`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.AddReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emojiName** | **string** | The target emoji&#39;s human-readable name.  To find an emoji&#39;s name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji&#39;s text name.  | 
 **emojiCode** | **string** | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  For most API clients, you won&#39;t need this, but it&#39;s important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the &#x60;emoji_code&#x60; in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a \\\&quot;different\\\&quot; emoji.  | 
 **reactionType** | **string** | A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity. Supported values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji](/help/custom-emoji); previously, this endpoint assumed &#x60;unicode_emoji&#x60; if this parameter was not specified.  | 

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


## CheckMessagesMatchNarrow

> CheckMessagesMatchNarrow200Response CheckMessagesMatchNarrow(ctx).MsgIds(msgIds).Narrow(narrow).Execute()

Check if messages match a narrow



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
	msgIds := []int32{int32(123)} // []int32 | List of IDs for the messages to check.
	narrow := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A structure defining the narrow to check against. See how to [construct a narrow](/api/construct-narrow).  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckMessagesMatchNarrow(context.Background()).MsgIds(msgIds).Narrow(narrow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CheckMessagesMatchNarrow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMessagesMatchNarrow`: CheckMessagesMatchNarrow200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CheckMessagesMatchNarrow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMessagesMatchNarrowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgIds** | **[]int32** | List of IDs for the messages to check. | 
 **narrow** | **[]map[string]interface{}** | A structure defining the narrow to check against. See how to [construct a narrow](/api/construct-narrow).  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation.  | 

### Return type

[**CheckMessagesMatchNarrow200Response**](CheckMessagesMatchNarrow200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessage

> JsonSuccess DeleteMessage(ctx, messageId).Execute()

Delete a message



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
	messageId := int32(43) // int32 | The target message's ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteMessage(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.DeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessage`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.DeleteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


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


## GetFileTemporaryUrl

> GetFileTemporaryUrl200Response GetFileTemporaryUrl(ctx, realmIdStr, filename).Execute()

Get public temporary URL



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
	realmIdStr := int32(1) // int32 | The realm ID. 
	filename := "4e/m2A3MSqFnWRLUf9SaPzQ0Up_/zulip.txt" // string | Path to the URL. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetFileTemporaryUrl(context.Background(), realmIdStr, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetFileTemporaryUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileTemporaryUrl`: GetFileTemporaryUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetFileTemporaryUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmIdStr** | **int32** | The realm ID.  | 
**filename** | **string** | Path to the URL.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileTemporaryUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFileTemporaryUrl200Response**](GetFileTemporaryUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessage

> GetMessage200Response GetMessage(ctx, messageId).ApplyMarkdown(applyMarkdown).AllowEmptyTopicName(allowEmptyTopicName).Execute()

Fetch a single message



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
	messageId := int32(43) // int32 | The target message's ID. 
	applyMarkdown := false // bool | If `true`, message content is returned in the rendered HTML format. If `false`, message content is returned in the raw [Zulip-flavored Markdown format](/help/format-your-message-using-markdown) text that user entered.  **Changes**: New in Zulip 5.0 (feature level 120).  (optional) (default to true)
	allowEmptyTopicName := true // bool | Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If `false`, the server will use the value of `realm_empty_topic_display_name` found in the [`POST /register`](/api/register-queue) response instead of empty string to represent the empty string topic in its response.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMessage(context.Background(), messageId).ApplyMarkdown(applyMarkdown).AllowEmptyTopicName(allowEmptyTopicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessage`: GetMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyMarkdown** | **bool** | If &#x60;true&#x60;, message content is returned in the rendered HTML format. If &#x60;false&#x60;, message content is returned in the raw [Zulip-flavored Markdown format](/help/format-your-message-using-markdown) text that user entered.  **Changes**: New in Zulip 5.0 (feature level 120).  | [default to true]
 **allowEmptyTopicName** | **bool** | Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If &#x60;false&#x60;, the server will use the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response instead of empty string to represent the empty string topic in its response.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  | [default to false]

### Return type

[**GetMessage200Response**](GetMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageHistory

> GetMessageHistory200Response GetMessageHistory(ctx, messageId).AllowEmptyTopicName(allowEmptyTopicName).Execute()

Get a message's edit history



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
	messageId := int32(43) // int32 | The target message's ID. 
	allowEmptyTopicName := true // bool | Whether the topic names i.e. `topic` and `prev_topic` fields in the `message_history` objects returned can be empty string.  If `false`, the value of `realm_empty_topic_display_name` found in the [`POST /register`](/api/register-queue) response is returned replacing the empty string as the topic name.  **Changes**: New in Zulip 10.0 (feature level 334).  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMessageHistory(context.Background(), messageId).AllowEmptyTopicName(allowEmptyTopicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageHistory`: GetMessageHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowEmptyTopicName** | **bool** | Whether the topic names i.e. &#x60;topic&#x60; and &#x60;prev_topic&#x60; fields in the &#x60;message_history&#x60; objects returned can be empty string.  If &#x60;false&#x60;, the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response is returned replacing the empty string as the topic name.  **Changes**: New in Zulip 10.0 (feature level 334).  | [default to false]

### Return type

[**GetMessageHistory200Response**](GetMessageHistory200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessages

> GetMessages200Response GetMessages(ctx).Anchor(anchor).IncludeAnchor(includeAnchor).NumBefore(numBefore).NumAfter(numAfter).Narrow(narrow).ClientGravatar(clientGravatar).ApplyMarkdown(applyMarkdown).UseFirstUnreadAnchor(useFirstUnreadAnchor).MessageIds(messageIds).AllowEmptyTopicName(allowEmptyTopicName).Execute()

Get messages



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
	anchor := "43" // string | Integer message ID to anchor fetching of new messages. Supports special string values for when the client wants the server to compute the anchor to use:  - `newest`: The most recent message. - `oldest`: The oldest message. - `first_unread`: The oldest unread message matching the   query, if any; otherwise, the most recent message.  **Changes**: String values are new in Zulip 3.0 (feature level 1). The `first_unread` functionality was supported in Zulip 2.1.x and older by not sending `anchor` and using `use_first_unread_anchor`.  In Zulip 2.1.x and older, `oldest` can be emulated with `\"anchor\": 0`, and `newest` with `\"anchor\": 10000000000000000` (that specific large value works around a bug in Zulip 2.1.x and older in the `found_newest` return value).  (optional)
	includeAnchor := false // bool | Whether a message with the specified ID matching the narrow should be included.  **Changes**: New in Zulip 6.0 (feature level 155).  (optional) (default to true)
	numBefore := int32(4) // int32 | The number of messages with IDs less than the anchor to retrieve. Required if `message_ids` is not provided.  (optional)
	numAfter := int32(8) // int32 | The number of messages with IDs greater than the anchor to retrieve. Required if `message_ids` is not provided.  (optional)
	narrow := []OneOfobjectarray{"TODO"} // []OneOfobjectarray | The narrow where you want to fetch the messages from. See how to [construct a narrow](/api/construct-narrow).  Note that many narrows, including all that lack a `channel`, `channels`, `stream`, or `streams` operator, search the user's personal message history. See [searching shared history](/help/search-for-messages#search-shared-history) for details.  For example, if you would like to fetch messages from all public channels instead of only the user's message history, then a specific narrow for messages sent to all public channels can be used: `{\"operator\": \"channels\", \"operand\": \"public\"}`.  Newly created bot users are not usually subscribed to any channels, so bots using this API should either be subscribed to appropriate channels or use a shared history search narrow with this endpoint.  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation.  (optional) (default to [])
	clientGravatar := false // bool | Whether the client supports computing gravatars URLs. If enabled, `avatar_url` will be included in the response only if there is a Zulip avatar, and will be `null` for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The `client_gravatar` field is set to `true` if clients can compute their own gravatars.  **Changes**: The default value of this parameter was `false` prior to Zulip 5.0 (feature level 92).  (optional) (default to true)
	applyMarkdown := false // bool | If `true`, message content is returned in the rendered HTML format. If `false`, message content is returned in the raw Markdown-format text that user entered.  See [Markdown message formatting](/api/message-formatting) for details on Zulip's HTML format.  (optional) (default to true)
	useFirstUnreadAnchor := true // bool | Legacy way to specify `\"anchor\": \"first_unread\"` in Zulip 2.1.x and older.  Whether to use the (computed by the server) first unread message matching the narrow as the `anchor`. Mutually exclusive with `anchor`.  **Changes**: Deprecated in Zulip 3.0 (feature level 1) and replaced by `\"anchor\": \"first_unread\"`.  (optional) (default to false)
	messageIds := []int32{int32(123)} // []int32 | A list of message IDs to fetch. The server will return messages corresponding to the subset of the requested message IDs that exist and the current user has access to, potentially filtered by the narrow (if that parameter is provided).  It is an error to pass this parameter as well as any of the parameters involved in specifying a range of messages: `anchor`, `include_anchor`, `use_first_unread_anchor`, `num_before`, and `num_after`.  **Changes**: New in Zulip 10.0 (feature level 300). Previously, there was no way to request a specific set of messages IDs.  (optional)
	allowEmptyTopicName := true // bool | Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If `false`, the server will use the value of `realm_empty_topic_display_name` found in the [`POST /register`](/api/register-queue) response instead of empty string to represent the empty string topic in its response.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetMessages(context.Background()).Anchor(anchor).IncludeAnchor(includeAnchor).NumBefore(numBefore).NumAfter(numAfter).Narrow(narrow).ClientGravatar(clientGravatar).ApplyMarkdown(applyMarkdown).UseFirstUnreadAnchor(useFirstUnreadAnchor).MessageIds(messageIds).AllowEmptyTopicName(allowEmptyTopicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessages`: GetMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anchor** | **string** | Integer message ID to anchor fetching of new messages. Supports special string values for when the client wants the server to compute the anchor to use:  - &#x60;newest&#x60;: The most recent message. - &#x60;oldest&#x60;: The oldest message. - &#x60;first_unread&#x60;: The oldest unread message matching the   query, if any; otherwise, the most recent message.  **Changes**: String values are new in Zulip 3.0 (feature level 1). The &#x60;first_unread&#x60; functionality was supported in Zulip 2.1.x and older by not sending &#x60;anchor&#x60; and using &#x60;use_first_unread_anchor&#x60;.  In Zulip 2.1.x and older, &#x60;oldest&#x60; can be emulated with &#x60;\&quot;anchor\&quot;: 0&#x60;, and &#x60;newest&#x60; with &#x60;\&quot;anchor\&quot;: 10000000000000000&#x60; (that specific large value works around a bug in Zulip 2.1.x and older in the &#x60;found_newest&#x60; return value).  | 
 **includeAnchor** | **bool** | Whether a message with the specified ID matching the narrow should be included.  **Changes**: New in Zulip 6.0 (feature level 155).  | [default to true]
 **numBefore** | **int32** | The number of messages with IDs less than the anchor to retrieve. Required if &#x60;message_ids&#x60; is not provided.  | 
 **numAfter** | **int32** | The number of messages with IDs greater than the anchor to retrieve. Required if &#x60;message_ids&#x60; is not provided.  | 
 **narrow** | [**[]OneOfobjectarray**](oneOf&lt;object,array&gt;.md) | The narrow where you want to fetch the messages from. See how to [construct a narrow](/api/construct-narrow).  Note that many narrows, including all that lack a &#x60;channel&#x60;, &#x60;channels&#x60;, &#x60;stream&#x60;, or &#x60;streams&#x60; operator, search the user&#39;s personal message history. See [searching shared history](/help/search-for-messages#search-shared-history) for details.  For example, if you would like to fetch messages from all public channels instead of only the user&#39;s message history, then a specific narrow for messages sent to all public channels can be used: &#x60;{\&quot;operator\&quot;: \&quot;channels\&quot;, \&quot;operand\&quot;: \&quot;public\&quot;}&#x60;.  Newly created bot users are not usually subscribed to any channels, so bots using this API should either be subscribed to appropriate channels or use a shared history search narrow with this endpoint.  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation.  | [default to []]
 **clientGravatar** | **bool** | Whether the client supports computing gravatars URLs. If enabled, &#x60;avatar_url&#x60; will be included in the response only if there is a Zulip avatar, and will be &#x60;null&#x60; for users who are using gravatar as their avatar. This option significantly reduces the compressed size of user data, since gravatar URLs are long, random strings and thus do not compress well. The &#x60;client_gravatar&#x60; field is set to &#x60;true&#x60; if clients can compute their own gravatars.  **Changes**: The default value of this parameter was &#x60;false&#x60; prior to Zulip 5.0 (feature level 92).  | [default to true]
 **applyMarkdown** | **bool** | If &#x60;true&#x60;, message content is returned in the rendered HTML format. If &#x60;false&#x60;, message content is returned in the raw Markdown-format text that user entered.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [default to true]
 **useFirstUnreadAnchor** | **bool** | Legacy way to specify &#x60;\&quot;anchor\&quot;: \&quot;first_unread\&quot;&#x60; in Zulip 2.1.x and older.  Whether to use the (computed by the server) first unread message matching the narrow as the &#x60;anchor&#x60;. Mutually exclusive with &#x60;anchor&#x60;.  **Changes**: Deprecated in Zulip 3.0 (feature level 1) and replaced by &#x60;\&quot;anchor\&quot;: \&quot;first_unread\&quot;&#x60;.  | [default to false]
 **messageIds** | **[]int32** | A list of message IDs to fetch. The server will return messages corresponding to the subset of the requested message IDs that exist and the current user has access to, potentially filtered by the narrow (if that parameter is provided).  It is an error to pass this parameter as well as any of the parameters involved in specifying a range of messages: &#x60;anchor&#x60;, &#x60;include_anchor&#x60;, &#x60;use_first_unread_anchor&#x60;, &#x60;num_before&#x60;, and &#x60;num_after&#x60;.  **Changes**: New in Zulip 10.0 (feature level 300). Previously, there was no way to request a specific set of messages IDs.  | 
 **allowEmptyTopicName** | **bool** | Whether the client supports processing the empty string as a topic in the topic name fields in the returned data, including in returned edit_history data.  If &#x60;false&#x60;, the server will use the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response instead of empty string to represent the empty string topic in its response.  **Changes**: New in Zulip 10.0 (feature level 334). Previously, the empty string was not a valid topic.  | [default to false]

### Return type

[**GetMessages200Response**](GetMessages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReadReceipts

> GetReadReceipts200Response GetReadReceipts(ctx, messageId).Execute()

Get a message's read receipts



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
	messageId := int32(43) // int32 | The target message's ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetReadReceipts(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetReadReceipts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReadReceipts`: GetReadReceipts200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetReadReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReadReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReadReceipts200Response**](GetReadReceipts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAllAsRead

> MarkAllAsRead200Response MarkAllAsRead(ctx).Execute()

Mark all messages as read



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
	resp, r, err := apiClient.MarkAllAsRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MarkAllAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkAllAsRead`: MarkAllAsRead200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MarkAllAsRead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAllAsReadRequest struct via the builder pattern


### Return type

[**MarkAllAsRead200Response**](MarkAllAsRead200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkStreamAsRead

> JsonSuccess MarkStreamAsRead(ctx).StreamId(streamId).Execute()

Mark messages in a channel as read



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkStreamAsRead(context.Background()).StreamId(streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MarkStreamAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkStreamAsRead`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MarkStreamAsRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkStreamAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **int32** | The ID of the channel to access.  | 

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


## MarkTopicAsRead

> JsonSuccess MarkTopicAsRead(ctx).StreamId(streamId).TopicName(topicName).Execute()

Mark messages in a topic as read



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
	topicName := "topicName_example" // string | The name of the topic whose messages should be marked as read.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarkTopicAsRead(context.Background()).StreamId(streamId).TopicName(topicName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MarkTopicAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTopicAsRead`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MarkTopicAsRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkTopicAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **int32** | The ID of the channel to access.  | 
 **topicName** | **string** | The name of the topic whose messages should be marked as read.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  | 

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


## RemoveReaction

> JsonSuccess RemoveReaction(ctx, messageId).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()

Remove an emoji reaction



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
	messageId := int32(43) // int32 | The target message's ID. 
	emojiName := "emojiName_example" // string | The target emoji's human-readable name.  To find an emoji's name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji's text name.  (optional)
	emojiCode := "emojiCode_example" // string | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.  For most API clients, you won't need this, but it's important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the `emoji_code` in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a \\\"different\\\" emoji.  (optional)
	reactionType := "reactionType_example" // string | A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity. Supported values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \\\"zulip\\\").  **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji](/help/custom-emoji); previously, this endpoint assumed `unicode_emoji` if this parameter was not specified.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveReaction(context.Background(), messageId).EmojiName(emojiName).EmojiCode(emojiCode).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.RemoveReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveReaction`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.RemoveReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emojiName** | **string** | The target emoji&#39;s human-readable name.  To find an emoji&#39;s name, hover over a message to reveal three icons on the right, then click the smiley face icon. Images of available reaction emojis appear. Hover over the emoji you want, and note that emoji&#39;s text name.  | 
 **emojiCode** | **string** | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  For most API clients, you won&#39;t need this, but it&#39;s important for Zulip apps to handle rare corner cases when adding/removing votes on an emoji reaction added previously by another user.  If the existing reaction was added when the Zulip server was using a previous version of the emoji data mapping between Unicode codepoints and human-readable names, sending the &#x60;emoji_code&#x60; in the data for the original reaction allows the Zulip server to correctly interpret your upvote as an upvote rather than a reaction with a \\\&quot;different\\\&quot; emoji.  | 
 **reactionType** | **string** | A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  If an API client is adding/removing a vote on an existing reaction, it should pass this parameter using the value the server provided for the existing reaction for specificity. Supported values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \\\&quot;zulip\\\&quot;).  **Changes**: In Zulip 3.0 (feature level 2), this parameter became optional for [custom emoji](/help/custom-emoji); previously, this endpoint assumed &#x60;unicode_emoji&#x60; if this parameter was not specified.  | 

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


## RenderMessage

> RenderMessage200Response RenderMessage(ctx).Content(content).Execute()

Render a message



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
	content := "content_example" // string | The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum message size. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RenderMessage(context.Background()).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.RenderMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderMessage`: RenderMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.RenderMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenderMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | The content of the message.  Clients should use the &#x60;max_message_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum message size.  | 

### Return type

[**RenderMessage200Response**](RenderMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportMessage

> JsonSuccess ReportMessage(ctx, messageId).ReportType(reportType).Description(description).Execute()

Report a message



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
	messageId := int32(43) // int32 | The target message's ID. 
	reportType := "reportType_example" // string | The reason that best describes why the current user is reporting the target message for moderation. 
	description := "description_example" // string | A short description with additional context about why the current user is reporting the target message for moderation.  Clients should limit this string to a maximum length of 1000 characters.  If the `report_type` parameter is `\\\"other\\\"`, this parameter is required, and its value cannot be an empty string.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportMessage(context.Background(), messageId).ReportType(reportType).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ReportMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportMessage`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ReportMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportType** | **string** | The reason that best describes why the current user is reporting the target message for moderation.  | 
 **description** | **string** | A short description with additional context about why the current user is reporting the target message for moderation.  Clients should limit this string to a maximum length of 1000 characters.  If the &#x60;report_type&#x60; parameter is &#x60;\\\&quot;other\\\&quot;&#x60;, this parameter is required, and its value cannot be an empty string.  | 

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


## SendMessage

> SendMessage200Response SendMessage(ctx).Type_(type_).To(to).Content(content).Topic(topic).QueueId(queueId).LocalId(localId).ReadBySender(readBySender).Execute()

Send a message



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
	type_ := "type__example" // string | The type of message to be sent.  `\\\"direct\\\"` for a direct message and `\\\"stream\\\"` or `\\\"channel\\\"` for a channel message.  **Changes**: In Zulip 9.0 (feature level 248), `\\\"channel\\\"` was added as an additional value for this parameter to request a channel message.  In Zulip 7.0 (feature level 174), `\\\"direct\\\"` was added as the preferred way to request a direct message, deprecating the original `\\\"private\\\"`. While `\\\"private\\\"` is still supported for requesting direct messages, clients are encouraged to use to the modern convention with servers that support it, because support for `\\\"private\\\"` will eventually be removed. 
	to := openapiclient.send_message_request_to{ArrayOfInt32: new([]int32)} // SendMessageRequestTo | 
	content := "content_example" // string | The content of the message.  Clients should use the `max_message_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum message size. 
	topic := "topic_example" // string | The topic of the message. Only required for channel messages (`\\\"type\\\": \\\"stream\\\"` or `\\\"type\\\": \\\"channel\\\"`), ignored otherwise.  Clients should use the `max_topic_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum topic length.  Note: When `\\\"(no topic)\\\"` or the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can't be `\\\"(no topic)\\\"`, an empty string, or the value of `realm_empty_topic_display_name`.  **Changes**: Before Zulip 10.0 (feature level 370), `\\\"(no topic)\\\"` was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as `subject`, which is currently a deprecated alias.  (optional)
	queueId := "queueId_example" // string | For clients supporting [local echo](https://zulip.readthedocs.io/en/latest/subsystems/sending-messages.html#local-echo), the [event queue](/api/register-queue) ID for the client. If passed, `local_id` is required. If the message is successfully sent, the server will include `local_id` in the `message` event that the client with this `queue_id` will receive notifying it of the new message via [`GET /events`](/api/get-events). This lets the client know unambiguously that it should replace the locally echoed message, rather than adding this new message (which would be correct if the user had sent the new message from another device).  (optional)
	localId := "localId_example" // string | For clients supporting local echo, a unique string-format identifier chosen freely by the client; the server will pass it back to the client without inspecting it, as described in the `queue_id` description.  (optional)
	readBySender := true // bool | Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name.  **Changes**: New in Zulip 8.0 (feature level 236).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendMessage(context.Background()).Type_(type_).To(to).Content(content).Topic(topic).QueueId(queueId).LocalId(localId).ReadBySender(readBySender).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: SendMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of message to be sent.  &#x60;\\\&quot;direct\\\&quot;&#x60; for a direct message and &#x60;\\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;channel\\\&quot;&#x60; for a channel message.  **Changes**: In Zulip 9.0 (feature level 248), &#x60;\\\&quot;channel\\\&quot;&#x60; was added as an additional value for this parameter to request a channel message.  In Zulip 7.0 (feature level 174), &#x60;\\\&quot;direct\\\&quot;&#x60; was added as the preferred way to request a direct message, deprecating the original &#x60;\\\&quot;private\\\&quot;&#x60;. While &#x60;\\\&quot;private\\\&quot;&#x60; is still supported for requesting direct messages, clients are encouraged to use to the modern convention with servers that support it, because support for &#x60;\\\&quot;private\\\&quot;&#x60; will eventually be removed.  | 
 **to** | [**SendMessageRequestTo**](SendMessageRequestTo.md) |  | 
 **content** | **string** | The content of the message.  Clients should use the &#x60;max_message_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum message size.  | 
 **topic** | **string** | The topic of the message. Only required for channel messages (&#x60;\\\&quot;type\\\&quot;: \\\&quot;stream\\\&quot;&#x60; or &#x60;\\\&quot;type\\\&quot;: \\\&quot;channel\\\&quot;&#x60;), ignored otherwise.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum topic length.  Note: When &#x60;\\\&quot;(no topic)\\\&quot;&#x60; or the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can&#39;t be &#x60;\\\&quot;(no topic)\\\&quot;&#x60;, an empty string, or the value of &#x60;realm_empty_topic_display_name&#x60;.  **Changes**: Before Zulip 10.0 (feature level 370), &#x60;\\\&quot;(no topic)\\\&quot;&#x60; was not interpreted as an empty string.  Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as &#x60;subject&#x60;, which is currently a deprecated alias.  | 
 **queueId** | **string** | For clients supporting [local echo](https://zulip.readthedocs.io/en/latest/subsystems/sending-messages.html#local-echo), the [event queue](/api/register-queue) ID for the client. If passed, &#x60;local_id&#x60; is required. If the message is successfully sent, the server will include &#x60;local_id&#x60; in the &#x60;message&#x60; event that the client with this &#x60;queue_id&#x60; will receive notifying it of the new message via [&#x60;GET /events&#x60;](/api/get-events). This lets the client know unambiguously that it should replace the locally echoed message, rather than adding this new message (which would be correct if the user had sent the new message from another device).  | 
 **localId** | **string** | For clients supporting local echo, a unique string-format identifier chosen freely by the client; the server will pass it back to the client without inspecting it, as described in the &#x60;queue_id&#x60; description.  | 
 **readBySender** | **bool** | Whether the message should be initially marked read by its sender. If unspecified, the server uses a heuristic based on the client name.  **Changes**: New in Zulip 8.0 (feature level 236).  | 

### Return type

[**SendMessage200Response**](SendMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> UpdateMessage200Response UpdateMessage(ctx, messageId).Topic(topic).PropagateMode(propagateMode).SendNotificationToOldThread(sendNotificationToOldThread).SendNotificationToNewThread(sendNotificationToNewThread).Content(content).PrevContentSha256(prevContentSha256).StreamId(streamId).Execute()

Edit a message



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
	messageId := int32(43) // int32 | The target message's ID. 
	topic := "topic_example" // string | The topic to move the message(s) to, to request changing the topic.  Clients should use the `max_topic_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum topic length  Should only be sent when changing the topic, and will throw an error if the target message is not a channel message.  Note: When the value of `realm_empty_topic_display_name` found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can't be `\\\"(no topic)\\\"`, an empty string, or the value of `realm_empty_topic_display_name`.  You can [resolve topics](/help/resolve-a-topic) by editing the topic to `✔ {original_topic}` with the `propagate_mode` parameter set to `\\\"change_all\\\"`. The empty string topic cannot be marked as resolved.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as `subject`, which is currently a deprecated alias.  (optional)
	propagateMode := "propagateMode_example" // string | Which message(s) should be edited:  - `\\\"change_later\\\"`: The target message and all following messages. - `\\\"change_one\\\"`: Only the target message. - `\\\"change_all\\\"`: All messages in this topic.  Only the default value of `\\\"change_one\\\"` is valid when editing only the content of a message.  This parameter determines both which messages get moved and also whether clients that are currently narrowed to the topic containing the message should navigate or adjust their compose box recipient to point to the post-edit channel/topic.  (optional) (default to "change_one")
	sendNotificationToOldThread := true // bool | Whether to send an automated message to the old topic to notify users where the messages were moved to.  **Changes**: Before Zulip 6.0 (feature level 152), this parameter had a default of `true` and was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).  (optional) (default to false)
	sendNotificationToNewThread := true // bool | Whether to send an automated message to the new topic to notify users where the messages came from.  If the move is just [resolving/unresolving a topic](/help/resolve-a-topic), this parameter will not trigger an additional notification.  **Changes**: Before Zulip 6.0 (feature level 152), this parameter was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).  (optional) (default to true)
	content := "content_example" // string | The updated content of the target message.  Clients should use the `max_message_length` returned by the [`POST /register`](/api/register-queue) endpoint to determine the maximum message size.  Note that a message's content and channel cannot be changed at the same time, so sending both `content` and `stream_id` parameters will throw an error.  (optional)
	prevContentSha256 := "prevContentSha256_example" // string | An optional SHA-256 hash of the previous raw content of the message that the client has at the time of the request.  If provided, the server will return an error if it does not match the SHA-256 hash of the message's content stored in the database.  Clients can use this feature to prevent races where multiple clients save conflicting edits to a message.  **Changes**: New in Zulip 11.0 (feature level 379).  (optional)
	streamId := int32(56) // int32 | The channel ID to move the message(s) to, to request moving messages to another channel.  Should only be sent when changing the channel, and will throw an error if the target message is not a channel message.  Note that a message's content and channel cannot be changed at the same time, so sending both `content` and `stream_id` parameters will throw an error.  **Changes**: New in Zulip 3.0 (feature level 1).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateMessage(context.Background(), messageId).Topic(topic).PropagateMode(propagateMode).SendNotificationToOldThread(sendNotificationToOldThread).SendNotificationToNewThread(sendNotificationToNewThread).Content(content).PrevContentSha256(prevContentSha256).StreamId(streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessage`: UpdateMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The target message&#39;s ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topic** | **string** | The topic to move the message(s) to, to request changing the topic.  Clients should use the &#x60;max_topic_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum topic length  Should only be sent when changing the topic, and will throw an error if the target message is not a channel message.  Note: When the value of &#x60;realm_empty_topic_display_name&#x60; found in the [POST /register](/api/register-queue) response is used for this parameter, it is interpreted as an empty string.  When [topics are required](/help/require-topics), this parameter can&#39;t be &#x60;\\\&quot;(no topic)\\\&quot;&#x60;, an empty string, or the value of &#x60;realm_empty_topic_display_name&#x60;.  You can [resolve topics](/help/resolve-a-topic) by editing the topic to &#x60;✔ {original_topic}&#x60; with the &#x60;propagate_mode&#x60; parameter set to &#x60;\\\&quot;change_all\\\&quot;&#x60;. The empty string topic cannot be marked as resolved.  **Changes**: Before Zulip 10.0 (feature level 334), empty string was not a valid topic name for channel messages.  New in Zulip 2.0.0. Previous Zulip releases encoded this as &#x60;subject&#x60;, which is currently a deprecated alias.  | 
 **propagateMode** | **string** | Which message(s) should be edited:  - &#x60;\\\&quot;change_later\\\&quot;&#x60;: The target message and all following messages. - &#x60;\\\&quot;change_one\\\&quot;&#x60;: Only the target message. - &#x60;\\\&quot;change_all\\\&quot;&#x60;: All messages in this topic.  Only the default value of &#x60;\\\&quot;change_one\\\&quot;&#x60; is valid when editing only the content of a message.  This parameter determines both which messages get moved and also whether clients that are currently narrowed to the topic containing the message should navigate or adjust their compose box recipient to point to the post-edit channel/topic.  | [default to &quot;change_one&quot;]
 **sendNotificationToOldThread** | **bool** | Whether to send an automated message to the old topic to notify users where the messages were moved to.  **Changes**: Before Zulip 6.0 (feature level 152), this parameter had a default of &#x60;true&#x60; and was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).  | [default to false]
 **sendNotificationToNewThread** | **bool** | Whether to send an automated message to the new topic to notify users where the messages came from.  If the move is just [resolving/unresolving a topic](/help/resolve-a-topic), this parameter will not trigger an additional notification.  **Changes**: Before Zulip 6.0 (feature level 152), this parameter was ignored unless the channel was changed.  New in Zulip 3.0 (feature level 9).  | [default to true]
 **content** | **string** | The updated content of the target message.  Clients should use the &#x60;max_message_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum message size.  Note that a message&#39;s content and channel cannot be changed at the same time, so sending both &#x60;content&#x60; and &#x60;stream_id&#x60; parameters will throw an error.  | 
 **prevContentSha256** | **string** | An optional SHA-256 hash of the previous raw content of the message that the client has at the time of the request.  If provided, the server will return an error if it does not match the SHA-256 hash of the message&#39;s content stored in the database.  Clients can use this feature to prevent races where multiple clients save conflicting edits to a message.  **Changes**: New in Zulip 11.0 (feature level 379).  | 
 **streamId** | **int32** | The channel ID to move the message(s) to, to request moving messages to another channel.  Should only be sent when changing the channel, and will throw an error if the target message is not a channel message.  Note that a message&#39;s content and channel cannot be changed at the same time, so sending both &#x60;content&#x60; and &#x60;stream_id&#x60; parameters will throw an error.  **Changes**: New in Zulip 3.0 (feature level 1).  | 

### Return type

[**UpdateMessage200Response**](UpdateMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageFlags

> UpdateMessageFlags200Response UpdateMessageFlags(ctx).Messages(messages).Op(op).Flag(flag).Execute()

Update personal message flags



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
	messages := []int32{int32(123)} // []int32 | An array containing the IDs of the target messages. 
	op := "op_example" // string | Whether to `add` the flag or `remove` it. 
	flag := "flag_example" // string | The flag that should be added/removed. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateMessageFlags(context.Background()).Messages(messages).Op(op).Flag(flag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessageFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessageFlags`: UpdateMessageFlags200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateMessageFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messages** | **[]int32** | An array containing the IDs of the target messages.  | 
 **op** | **string** | Whether to &#x60;add&#x60; the flag or &#x60;remove&#x60; it.  | 
 **flag** | **string** | The flag that should be added/removed.  | 

### Return type

[**UpdateMessageFlags200Response**](UpdateMessageFlags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageFlagsForNarrow

> UpdateMessageFlagsForNarrow200Response UpdateMessageFlagsForNarrow(ctx).Anchor(anchor).NumBefore(numBefore).NumAfter(numAfter).Narrow(narrow).Op(op).Flag(flag).IncludeAnchor(includeAnchor).Execute()

Update personal message flags for narrow



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
	anchor := TODO // string | Integer message ID to anchor updating of flags. Supports special string values for when the client wants the server to compute the anchor to use:  - `newest`: The most recent message. - `oldest`: The oldest message. - `first_unread`: The oldest unread message matching the   query, if any; otherwise, the most recent message. 
	numBefore := int32(56) // int32 | Limit the number of messages preceding the anchor in the update range. The server may decrease this to bound transaction sizes. 
	numAfter := int32(56) // int32 | Limit the number of messages following the anchor in the update range. The server may decrease this to bound transaction sizes. 
	narrow := []openapiclient.UpdateFlagsNarrowClause{openapiclient.update_message_flags_for_narrow_request_narrow_inner{UpdateFlagsNarrowFilter: openapiclient.NewUpdateFlagsNarrowFilter("Operator_example", openapiclient.update_message_flags_for_narrow_request_narrow_inner_oneOf_operand{ArrayOfInt32: new([]int32)})}} // []UpdateFlagsNarrowClause | The narrow you want update flags within. See how to [construct a narrow](/api/construct-narrow).  Note that, when adding the `read` flag to messages, clients should consider including a narrow with the `is:unread` filter as an optimization. Including that filter takes advantage of the fact that the server has a database index for unread messages.  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation. 
	op := "op_example" // string | Whether to `add` the flag or `remove` it. 
	flag := "flag_example" // string | The flag that should be added/removed. See [available flags](/api/update-message-flags#available-flags). 
	includeAnchor := true // bool | Whether a message with the specified ID matching the narrow should be included in the update range.  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateMessageFlagsForNarrow(context.Background()).Anchor(anchor).NumBefore(numBefore).NumAfter(numAfter).Narrow(narrow).Op(op).Flag(flag).IncludeAnchor(includeAnchor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessageFlagsForNarrow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessageFlagsForNarrow`: UpdateMessageFlagsForNarrow200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateMessageFlagsForNarrow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageFlagsForNarrowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anchor** | [**string**](string.md) | Integer message ID to anchor updating of flags. Supports special string values for when the client wants the server to compute the anchor to use:  - &#x60;newest&#x60;: The most recent message. - &#x60;oldest&#x60;: The oldest message. - &#x60;first_unread&#x60;: The oldest unread message matching the   query, if any; otherwise, the most recent message.  | 
 **numBefore** | **int32** | Limit the number of messages preceding the anchor in the update range. The server may decrease this to bound transaction sizes.  | 
 **numAfter** | **int32** | Limit the number of messages following the anchor in the update range. The server may decrease this to bound transaction sizes.  | 
 **narrow** | [**[]UpdateFlagsNarrowClause**](UpdateFlagsNarrowClause.md) | The narrow you want update flags within. See how to [construct a narrow](/api/construct-narrow).  Note that, when adding the &#x60;read&#x60; flag to messages, clients should consider including a narrow with the &#x60;is:unread&#x60; filter as an optimization. Including that filter takes advantage of the fact that the server has a database index for unread messages.  **Changes**: See [changes section](/api/construct-narrow#changes) of search/narrow filter documentation.  | 
 **op** | **string** | Whether to &#x60;add&#x60; the flag or &#x60;remove&#x60; it.  | 
 **flag** | **string** | The flag that should be added/removed. See [available flags](/api/update-message-flags#available-flags).  | 
 **includeAnchor** | **bool** | Whether a message with the specified ID matching the narrow should be included in the update range.  | [default to true]

### Return type

[**UpdateMessageFlagsForNarrow200Response**](UpdateMessageFlagsForNarrow200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> UploadFile200Response UploadFile(ctx).Filename(filename).Execute()

Upload a file



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
	filename := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadFile(context.Background()).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: UploadFile200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | ***os.File** |  | 

### Return type

[**UploadFile200Response**](UploadFile200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

