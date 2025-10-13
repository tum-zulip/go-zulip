# MessagesBaseEditHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevContent** | Pointer to **string** | Only present if message&#39;s content was edited.  The content of the message immediately prior to this edit event.  | [optional] 
**PrevRenderedContent** | Pointer to **string** | Only present if message&#39;s content was edited.  The rendered HTML representation of &#x60;prev_content&#x60;.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**PrevStream** | Pointer to **int32** | Only present if message&#39;s channel was edited.  The channel ID of the message immediately prior to this edit event.  **Changes**: New in Zulip 3.0 (feature level 1).  | [optional] 
**PrevTopic** | Pointer to **string** | Only present if message&#39;s topic was edited.  The topic of the message immediately prior to this edit event.  **Changes**: New in Zulip 5.0 (feature level 118). Previously, this field was called &#x60;prev_subject&#x60;; clients are recommended to rename &#x60;prev_subject&#x60; to &#x60;prev_topic&#x60; if present for compatibility with older Zulip servers.  | [optional] 
**Stream** | Pointer to **int32** | Only present if message&#39;s channel was edited.  The ID of the channel containing the message immediately after this edit event.  **Changes**: New in Zulip 5.0 (feature level 118).  | [optional] 
**Timestamp** | **int32** | The UNIX timestamp for the edit.  | 
**Topic** | Pointer to **string** | Only present if message&#39;s topic was edited.  The topic of the message immediately after this edit event.  **Changes**: New in Zulip 5.0 (feature level 118).  | [optional] 
**UserId** | **NullableInt32** | The ID of the user that made the edit.  Will be &#x60;null&#x60; only for edit history events predating March 2017.  Clients can display edit history events where this is &#x60;null&#x60; as modified by either the sender (for content edits) or an unknown user (for topic edits).  | 

## Methods

### NewMessagesBaseEditHistoryInner

`func NewMessagesBaseEditHistoryInner(timestamp int32, userId NullableInt32, ) *MessagesBaseEditHistoryInner`

NewMessagesBaseEditHistoryInner instantiates a new MessagesBaseEditHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesBaseEditHistoryInnerWithDefaults

`func NewMessagesBaseEditHistoryInnerWithDefaults() *MessagesBaseEditHistoryInner`

NewMessagesBaseEditHistoryInnerWithDefaults instantiates a new MessagesBaseEditHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevContent

`func (o *MessagesBaseEditHistoryInner) GetPrevContent() string`

GetPrevContent returns the PrevContent field if non-nil, zero value otherwise.

### GetPrevContentOk

`func (o *MessagesBaseEditHistoryInner) GetPrevContentOk() (*string, bool)`

GetPrevContentOk returns a tuple with the PrevContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevContent

`func (o *MessagesBaseEditHistoryInner) SetPrevContent(v string)`

SetPrevContent sets PrevContent field to given value.

### HasPrevContent

`func (o *MessagesBaseEditHistoryInner) HasPrevContent() bool`

HasPrevContent returns a boolean if a field has been set.

### GetPrevRenderedContent

`func (o *MessagesBaseEditHistoryInner) GetPrevRenderedContent() string`

GetPrevRenderedContent returns the PrevRenderedContent field if non-nil, zero value otherwise.

### GetPrevRenderedContentOk

`func (o *MessagesBaseEditHistoryInner) GetPrevRenderedContentOk() (*string, bool)`

GetPrevRenderedContentOk returns a tuple with the PrevRenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRenderedContent

`func (o *MessagesBaseEditHistoryInner) SetPrevRenderedContent(v string)`

SetPrevRenderedContent sets PrevRenderedContent field to given value.

### HasPrevRenderedContent

`func (o *MessagesBaseEditHistoryInner) HasPrevRenderedContent() bool`

HasPrevRenderedContent returns a boolean if a field has been set.

### GetPrevStream

`func (o *MessagesBaseEditHistoryInner) GetPrevStream() int32`

GetPrevStream returns the PrevStream field if non-nil, zero value otherwise.

### GetPrevStreamOk

`func (o *MessagesBaseEditHistoryInner) GetPrevStreamOk() (*int32, bool)`

GetPrevStreamOk returns a tuple with the PrevStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevStream

`func (o *MessagesBaseEditHistoryInner) SetPrevStream(v int32)`

SetPrevStream sets PrevStream field to given value.

### HasPrevStream

`func (o *MessagesBaseEditHistoryInner) HasPrevStream() bool`

HasPrevStream returns a boolean if a field has been set.

### GetPrevTopic

`func (o *MessagesBaseEditHistoryInner) GetPrevTopic() string`

GetPrevTopic returns the PrevTopic field if non-nil, zero value otherwise.

### GetPrevTopicOk

`func (o *MessagesBaseEditHistoryInner) GetPrevTopicOk() (*string, bool)`

GetPrevTopicOk returns a tuple with the PrevTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTopic

`func (o *MessagesBaseEditHistoryInner) SetPrevTopic(v string)`

SetPrevTopic sets PrevTopic field to given value.

### HasPrevTopic

`func (o *MessagesBaseEditHistoryInner) HasPrevTopic() bool`

HasPrevTopic returns a boolean if a field has been set.

### GetStream

`func (o *MessagesBaseEditHistoryInner) GetStream() int32`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *MessagesBaseEditHistoryInner) GetStreamOk() (*int32, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *MessagesBaseEditHistoryInner) SetStream(v int32)`

SetStream sets Stream field to given value.

### HasStream

`func (o *MessagesBaseEditHistoryInner) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessagesBaseEditHistoryInner) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessagesBaseEditHistoryInner) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessagesBaseEditHistoryInner) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTopic

`func (o *MessagesBaseEditHistoryInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MessagesBaseEditHistoryInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MessagesBaseEditHistoryInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MessagesBaseEditHistoryInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUserId

`func (o *MessagesBaseEditHistoryInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessagesBaseEditHistoryInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessagesBaseEditHistoryInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *MessagesBaseEditHistoryInner) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MessagesBaseEditHistoryInner) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


