# UpdateMessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | 
**Type** | **string** |  | 
**UserId** | **NullableInt32** | The ID of the user who sent the message.  Is &#x60;null&#x60; when event is for a rendering update of the original message, such as for an [inline URL preview][inline-url-previews].  **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all &#x60;update_message&#x60; events. Previously, this field was omitted for [inline URL preview][inline-url-previews] updates.  | 
**RenderingOnly** | **bool** | Whether the event only updates the rendered content of the message.  This field should be used by clients to determine if the event only provides a rendering update to the message content, such as for an [inline URL preview][inline-url-previews]. When &#x60;true&#x60;, the event does not reflect a user-generated edit and does not modify the message history.  **Changes**: New in Zulip 5.0 (feature level 114). Clients can correctly identify these rendering update event with earlier Zulip versions by checking whether the &#x60;user_id&#x60; field was omitted.  | 
**MessageId** | **int32** | The ID of the message which was edited or updated.  This field should be used to apply content edits to the client&#39;s cached message history, or to apply rendered content updates.  If the channel or topic was changed, the set of moved messages is encoded in the separate &#x60;message_ids&#x60; field, which is guaranteed to include &#x60;message_id&#x60;.  | 
**MessageIds** | **[]int32** | A sorted list of IDs of messages to which any channel or topic changes encoded in this event should be applied.  This list always includes &#x60;message_id&#x60;, even when there are no channel or topic changes to apply.  These messages are guaranteed to have all been previously sent to channel &#x60;stream_id&#x60; with topic &#x60;orig_subject&#x60;, and have been moved to &#x60;new_stream_id&#x60; with topic &#x60;subject&#x60; (if those fields are present in the event).  Clients processing these events should update all cached message history associated with the moved messages (including adjusting &#x60;unread_msgs&#x60; data structures, where the client may not have the message itself in its history) to reflect the new channel and topic.  Content changes should be applied only to the single message indicated by &#x60;message_id&#x60;.  **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.  | 
**Flags** | **[]string** | The user&#39;s personal [message flags][message-flags] for the message with ID &#x60;message_id&#x60; following the edit.  A client application should compare these to the original flags to identify cases where a mention or alert word was added by the edit.  **Changes**: In Zulip 8.0 (feature level 224), the &#x60;wildcard_mentioned&#x60; flag was deprecated in favor of the &#x60;stream_wildcard_mentioned&#x60; and &#x60;topic_wildcard_mentioned&#x60; flags. The &#x60;wildcard_mentioned&#x60; flag exists for backwards compatibility with older clients and equals &#x60;stream_wildcard_mentioned || topic_wildcard_mentioned&#x60;. Clients supporting older server versions should treat this field as a previous name for the &#x60;stream_wildcard_mentioned&#x60; flag as topic wildcard mentions were not available prior to this feature level.  [message-flags]: /api/update-message-flags#available-flags  | 
**EditTimestamp** | **int32** | The time when this message edit operation was processed by the server.  **Changes**: As of Zulip 5.0 (feature level 114), this field is present for all &#x60;update_message&#x60; events. Previously, this field was omitted for [inline URL preview][inline-url-previews] updates.  | 
**StreamName** | Pointer to **string** | Only present if the message was edited and originally sent to a channel.  The name of the channel that the message was sent to. Clients are recommended to use the &#x60;stream_id&#x60; field instead.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if the message was edited and originally sent to a channel.  The pre-edit channel for all of the messages with IDs in &#x60;message_ids&#x60;.  **Changes**: As of Zulip 5.0 (feature level 112), this field is present for all edits to a channel message. Previously, it was not present when only the content of the channel message was edited.  | [optional] 
**NewStreamId** | Pointer to **int32** | Only present if message(s) were moved to a different channel.  The post-edit channel for all of the messages with IDs in &#x60;message_ids&#x60;.  | [optional] 
**PropagateMode** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The choice the editing user made about which messages should be affected by a channel/topic edit:  - &#x60;\&quot;change_one\&quot;&#x60;: Just change the one indicated in &#x60;message_id&#x60;. - &#x60;\&quot;change_later\&quot;&#x60;: Change messages in the same topic that had   been sent after this one. - &#x60;\&quot;change_all\&quot;&#x60;: Change all messages in that topic.  This parameter should be used to decide whether to change navigation and compose box state in response to the edit. For example, if the user was previously in topic narrow, and the topic was edited with &#x60;\&quot;change_later\&quot;&#x60; or &#x60;\&quot;change_all\&quot;&#x60;, the Zulip web app will automatically navigate to the new topic narrow. Similarly, a message being composed to the old topic should have its recipient changed to the new topic.  This navigation makes it much more convenient to move content between topics without disruption or messages continuing to be sent to the pre-edit topic by accident.  | [optional] 
**OrigSubject** | Pointer to **string** | Only present if this event moved messages to a different topic and/or channel.  The pre-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual pre-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**Subject** | Pointer to **string** | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  The post-edit topic for all of the messages with IDs in &#x60;message_ids&#x60;.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual post-edit topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**TopicLinks** | Pointer to [**[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner**](GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner.md) | Only present if this event moved messages to a different topic; this field will not be present when moving messages to the same topic name in a different channel.  Data on any links to be included in the &#x60;topic&#x60; line (these are generated by [custom linkification filter](/help/add-a-custom-linkifier) that match content in the message&#39;s topic.), corresponding to the post-edit topic.  **Changes**: This field contained a list of urls before Zulip 4.0 (feature level 46).  New in Zulip 3.0 (feature level 1). Previously, this field was called &#x60;subject_links&#x60;; clients are recommended to rename &#x60;subject_links&#x60; to &#x60;topic_links&#x60; if present for compatibility with older Zulip servers.  | [optional] 
**OrigContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, in the original [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  | [optional] 
**OrigRenderedContent** | Pointer to **string** | Only present if this event changed the message content.  The original content of the message with ID &#x60;message_id&#x60; immediately prior to this edit, rendered as HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**Content** | Pointer to **string** | Only present if this event changed the message content or updated the message content for an [inline URL preview][inline-url-previews].  The new content of the message with ID &#x60;message_id&#x60;, in the original [Zulip-flavored Markdown](/help/format-your-message-using-markdown) format.  | [optional] 
**RenderedContent** | Pointer to **string** | Only present if this event changed the message content or updated the message content for an [inline URL preview][inline-url-previews].  The new content of the message with ID &#x60;message_id&#x60;, rendered in HTML.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**IsMeMessage** | Pointer to **bool** | Only present if this event changed the message content.  Whether the message with ID &#x60;message_id&#x60; is now a [/me status message][status-messages].  [status-messages]: /help/format-your-message-using-markdown#status-messages  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf36

`func NewGetEvents200ResponseAllOfEventsInnerOneOf36(id int32, type_ string, userId NullableInt32, renderingOnly bool, messageId int32, messageIds []int32, flags []string, editTimestamp int32, ) *UpdateMessageEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf36 instantiates a new UpdateMessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf36WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf36WithDefaults() *UpdateMessageEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf36WithDefaults instantiates a new UpdateMessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMessageEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMessageEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMessageEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateMessageEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMessageEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMessageEvent) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *UpdateMessageEvent) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateMessageEvent) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateMessageEvent) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *UpdateMessageEvent) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UpdateMessageEvent) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetRenderingOnly

`func (o *UpdateMessageEvent) GetRenderingOnly() bool`

GetRenderingOnly returns the RenderingOnly field if non-nil, zero value otherwise.

### GetRenderingOnlyOk

`func (o *UpdateMessageEvent) GetRenderingOnlyOk() (*bool, bool)`

GetRenderingOnlyOk returns a tuple with the RenderingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderingOnly

`func (o *UpdateMessageEvent) SetRenderingOnly(v bool)`

SetRenderingOnly sets RenderingOnly field to given value.


### GetMessageId

`func (o *UpdateMessageEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *UpdateMessageEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *UpdateMessageEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetMessageIds

`func (o *UpdateMessageEvent) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *UpdateMessageEvent) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *UpdateMessageEvent) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.


### GetFlags

`func (o *UpdateMessageEvent) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateMessageEvent) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateMessageEvent) SetFlags(v []string)`

SetFlags sets Flags field to given value.


### GetEditTimestamp

`func (o *UpdateMessageEvent) GetEditTimestamp() int32`

GetEditTimestamp returns the EditTimestamp field if non-nil, zero value otherwise.

### GetEditTimestampOk

`func (o *UpdateMessageEvent) GetEditTimestampOk() (*int32, bool)`

GetEditTimestampOk returns a tuple with the EditTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditTimestamp

`func (o *UpdateMessageEvent) SetEditTimestamp(v int32)`

SetEditTimestamp sets EditTimestamp field to given value.


### GetStreamName

`func (o *UpdateMessageEvent) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *UpdateMessageEvent) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *UpdateMessageEvent) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *UpdateMessageEvent) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetStreamId

`func (o *UpdateMessageEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *UpdateMessageEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *UpdateMessageEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *UpdateMessageEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetNewStreamId

`func (o *UpdateMessageEvent) GetNewStreamId() int32`

GetNewStreamId returns the NewStreamId field if non-nil, zero value otherwise.

### GetNewStreamIdOk

`func (o *UpdateMessageEvent) GetNewStreamIdOk() (*int32, bool)`

GetNewStreamIdOk returns a tuple with the NewStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStreamId

`func (o *UpdateMessageEvent) SetNewStreamId(v int32)`

SetNewStreamId sets NewStreamId field to given value.

### HasNewStreamId

`func (o *UpdateMessageEvent) HasNewStreamId() bool`

HasNewStreamId returns a boolean if a field has been set.

### GetPropagateMode

`func (o *UpdateMessageEvent) GetPropagateMode() string`

GetPropagateMode returns the PropagateMode field if non-nil, zero value otherwise.

### GetPropagateModeOk

`func (o *UpdateMessageEvent) GetPropagateModeOk() (*string, bool)`

GetPropagateModeOk returns a tuple with the PropagateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateMode

`func (o *UpdateMessageEvent) SetPropagateMode(v string)`

SetPropagateMode sets PropagateMode field to given value.

### HasPropagateMode

`func (o *UpdateMessageEvent) HasPropagateMode() bool`

HasPropagateMode returns a boolean if a field has been set.

### GetOrigSubject

`func (o *UpdateMessageEvent) GetOrigSubject() string`

GetOrigSubject returns the OrigSubject field if non-nil, zero value otherwise.

### GetOrigSubjectOk

`func (o *UpdateMessageEvent) GetOrigSubjectOk() (*string, bool)`

GetOrigSubjectOk returns a tuple with the OrigSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSubject

`func (o *UpdateMessageEvent) SetOrigSubject(v string)`

SetOrigSubject sets OrigSubject field to given value.

### HasOrigSubject

`func (o *UpdateMessageEvent) HasOrigSubject() bool`

HasOrigSubject returns a boolean if a field has been set.

### GetSubject

`func (o *UpdateMessageEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateMessageEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateMessageEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateMessageEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTopicLinks

`func (o *UpdateMessageEvent) GetTopicLinks() []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner`

GetTopicLinks returns the TopicLinks field if non-nil, zero value otherwise.

### GetTopicLinksOk

`func (o *UpdateMessageEvent) GetTopicLinksOk() (*[]GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner, bool)`

GetTopicLinksOk returns a tuple with the TopicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicLinks

`func (o *UpdateMessageEvent) SetTopicLinks(v []GetEvents200ResponseAllOfEventsInnerOneOf36TopicLinksInner)`

SetTopicLinks sets TopicLinks field to given value.

### HasTopicLinks

`func (o *UpdateMessageEvent) HasTopicLinks() bool`

HasTopicLinks returns a boolean if a field has been set.

### GetOrigContent

`func (o *UpdateMessageEvent) GetOrigContent() string`

GetOrigContent returns the OrigContent field if non-nil, zero value otherwise.

### GetOrigContentOk

`func (o *UpdateMessageEvent) GetOrigContentOk() (*string, bool)`

GetOrigContentOk returns a tuple with the OrigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigContent

`func (o *UpdateMessageEvent) SetOrigContent(v string)`

SetOrigContent sets OrigContent field to given value.

### HasOrigContent

`func (o *UpdateMessageEvent) HasOrigContent() bool`

HasOrigContent returns a boolean if a field has been set.

### GetOrigRenderedContent

`func (o *UpdateMessageEvent) GetOrigRenderedContent() string`

GetOrigRenderedContent returns the OrigRenderedContent field if non-nil, zero value otherwise.

### GetOrigRenderedContentOk

`func (o *UpdateMessageEvent) GetOrigRenderedContentOk() (*string, bool)`

GetOrigRenderedContentOk returns a tuple with the OrigRenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigRenderedContent

`func (o *UpdateMessageEvent) SetOrigRenderedContent(v string)`

SetOrigRenderedContent sets OrigRenderedContent field to given value.

### HasOrigRenderedContent

`func (o *UpdateMessageEvent) HasOrigRenderedContent() bool`

HasOrigRenderedContent returns a boolean if a field has been set.

### GetContent

`func (o *UpdateMessageEvent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateMessageEvent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateMessageEvent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateMessageEvent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRenderedContent

`func (o *UpdateMessageEvent) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *UpdateMessageEvent) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *UpdateMessageEvent) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.

### HasRenderedContent

`func (o *UpdateMessageEvent) HasRenderedContent() bool`

HasRenderedContent returns a boolean if a field has been set.

### GetIsMeMessage

`func (o *UpdateMessageEvent) GetIsMeMessage() bool`

GetIsMeMessage returns the IsMeMessage field if non-nil, zero value otherwise.

### GetIsMeMessageOk

`func (o *UpdateMessageEvent) GetIsMeMessageOk() (*bool, bool)`

GetIsMeMessageOk returns a tuple with the IsMeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeMessage

`func (o *UpdateMessageEvent) SetIsMeMessage(v bool)`

SetIsMeMessage sets IsMeMessage field to given value.

### HasIsMeMessage

`func (o *UpdateMessageEvent) HasIsMeMessage() bool`

HasIsMeMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


