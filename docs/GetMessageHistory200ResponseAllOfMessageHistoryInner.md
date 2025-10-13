# GetMessageHistory200ResponseAllOfMessageHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The topic of the message immediately after this edit event.  | [optional] 
**PrevTopic** | Pointer to **string** | Only present if message&#39;s topic was edited.  The topic of the message immediately prior to this edit event.  | [optional] 
**Stream** | Pointer to **int32** | Only present if message&#39;s channel was edited.  The ID of the channel containing the message immediately after this edit event.  **Changes**: New in Zulip 5.0 (feature level 118).  | [optional] 
**PrevStream** | Pointer to **int32** | Only present if message&#39;s channel was edited.  The ID of the channel containing the message immediately prior to this edit event.  **Changes**: New in Zulip 3.0 (feature level 1).  | [optional] 
**Content** | Pointer to **string** | The raw [Zulip-flavored Markdown](/help/format-your-message-using-markdown) content of the message immediately after this edit event.  | [optional] 
**RenderedContent** | Pointer to **string** | The rendered HTML representation of &#x60;content&#x60;.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**PrevContent** | Pointer to **string** | Only present if message&#39;s content was edited.  The raw [Zulip-flavored Markdown](/help/format-your-message-using-markdown) content of the message immediately prior to this edit event.  | [optional] 
**PrevRenderedContent** | Pointer to **string** | Only present if message&#39;s content was edited.  The rendered HTML representation of &#x60;prev_content&#x60;.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 
**UserId** | Pointer to **NullableInt32** | The ID of the user that made the edit.  Will be &#x60;null&#x60; only for edit history events predating March 2017.  Clients can display edit history events where this is &#x60;null&#x60; as modified by either the sender (for content edits) or an unknown user (for topic edits).  | [optional] 
**ContentHtmlDiff** | Pointer to **string** | Only present if message&#39;s content was edited.  An HTML diff between this version of the message and the previous one.  | [optional] 
**Timestamp** | Pointer to **int32** | The UNIX timestamp for this edit.  | [optional] 

## Methods

### NewGetMessageHistory200ResponseAllOfMessageHistoryInner

`func NewGetMessageHistory200ResponseAllOfMessageHistoryInner() *GetMessageHistory200ResponseAllOfMessageHistoryInner`

NewGetMessageHistory200ResponseAllOfMessageHistoryInner instantiates a new GetMessageHistory200ResponseAllOfMessageHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageHistory200ResponseAllOfMessageHistoryInnerWithDefaults

`func NewGetMessageHistory200ResponseAllOfMessageHistoryInnerWithDefaults() *GetMessageHistory200ResponseAllOfMessageHistoryInner`

NewGetMessageHistory200ResponseAllOfMessageHistoryInnerWithDefaults instantiates a new GetMessageHistory200ResponseAllOfMessageHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPrevTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevTopic() string`

GetPrevTopic returns the PrevTopic field if non-nil, zero value otherwise.

### GetPrevTopicOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevTopicOk() (*string, bool)`

GetPrevTopicOk returns a tuple with the PrevTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetPrevTopic(v string)`

SetPrevTopic sets PrevTopic field to given value.

### HasPrevTopic

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasPrevTopic() bool`

HasPrevTopic returns a boolean if a field has been set.

### GetStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetStream() int32`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetStreamOk() (*int32, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetStream(v int32)`

SetStream sets Stream field to given value.

### HasStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetPrevStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevStream() int32`

GetPrevStream returns the PrevStream field if non-nil, zero value otherwise.

### GetPrevStreamOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevStreamOk() (*int32, bool)`

GetPrevStreamOk returns a tuple with the PrevStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetPrevStream(v int32)`

SetPrevStream sets PrevStream field to given value.

### HasPrevStream

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasPrevStream() bool`

HasPrevStream returns a boolean if a field has been set.

### GetContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.

### HasRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasRenderedContent() bool`

HasRenderedContent returns a boolean if a field has been set.

### GetPrevContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevContent() string`

GetPrevContent returns the PrevContent field if non-nil, zero value otherwise.

### GetPrevContentOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevContentOk() (*string, bool)`

GetPrevContentOk returns a tuple with the PrevContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetPrevContent(v string)`

SetPrevContent sets PrevContent field to given value.

### HasPrevContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasPrevContent() bool`

HasPrevContent returns a boolean if a field has been set.

### GetPrevRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevRenderedContent() string`

GetPrevRenderedContent returns the PrevRenderedContent field if non-nil, zero value otherwise.

### GetPrevRenderedContentOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetPrevRenderedContentOk() (*string, bool)`

GetPrevRenderedContentOk returns a tuple with the PrevRenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetPrevRenderedContent(v string)`

SetPrevRenderedContent sets PrevRenderedContent field to given value.

### HasPrevRenderedContent

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasPrevRenderedContent() bool`

HasPrevRenderedContent returns a boolean if a field has been set.

### GetUserId

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetContentHtmlDiff

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetContentHtmlDiff() string`

GetContentHtmlDiff returns the ContentHtmlDiff field if non-nil, zero value otherwise.

### GetContentHtmlDiffOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetContentHtmlDiffOk() (*string, bool)`

GetContentHtmlDiffOk returns a tuple with the ContentHtmlDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHtmlDiff

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetContentHtmlDiff(v string)`

SetContentHtmlDiff sets ContentHtmlDiff field to given value.

### HasContentHtmlDiff

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasContentHtmlDiff() bool`

HasContentHtmlDiff returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetMessageHistory200ResponseAllOfMessageHistoryInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


