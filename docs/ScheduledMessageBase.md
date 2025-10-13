# ScheduledMessageBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledMessageId** | **int32** | The unique ID of the scheduled message, which can be used to modify or delete the scheduled message.  This is different from the unique ID that the message will have after it is sent.  | 
**Type** | **string** | The type of the scheduled message. Either &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;private\&quot;&#x60;.  | 
**To** | [**Recipients**](Recipients.md) |  | 
**Topic** | Pointer to **string** | Only present if &#x60;type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The topic for the channel message.  | [optional] 
**Content** | **string** | The content/body of the scheduled message, in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  See [Markdown message formatting](zulip.com/api/message-formatting for details on Zulip&#39;s HTML format.  | 
**RenderedContent** | **string** | The content/body of the scheduled message rendered in HTML.  | 
**ScheduledDeliveryTimestamp** | **int32** | The UNIX timestamp for when the message will be sent by the server, in UTC seconds.  | 
**Failed** | **bool** | Whether the server has tried to send the scheduled message and it failed to successfully send.  Clients that support unscheduling and editing scheduled messages should display scheduled messages with &#x60;\&quot;failed\&quot;: true&#x60; with an indicator that the server failed to send the message at the scheduled time, so that the user is aware of the failure and can get the content of the scheduled message.  **Changes**: New in Zulip 7.0 (feature level 181).  | 

## Methods

### NewScheduledMessageBase

`func NewScheduledMessageBase(scheduledMessageId int32, type_ string, to Recipients, content string, renderedContent string, scheduledDeliveryTimestamp int32, failed bool, ) *ScheduledMessageBase`

NewScheduledMessageBase instantiates a new ScheduledMessageBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledMessageBaseWithDefaults

`func NewScheduledMessageBaseWithDefaults() *ScheduledMessageBase`

NewScheduledMessageBaseWithDefaults instantiates a new ScheduledMessageBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledMessageId

`func (o *ScheduledMessageBase) GetScheduledMessageId() int32`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *ScheduledMessageBase) GetScheduledMessageIdOk() (*int32, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *ScheduledMessageBase) SetScheduledMessageId(v int32)`

SetScheduledMessageId sets ScheduledMessageId field to given value.


### GetType

`func (o *ScheduledMessageBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledMessageBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledMessageBase) SetType(v string)`

SetType sets Type field to given value.


### GetTo

`func (o *ScheduledMessageBase) GetTo() Recipients`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ScheduledMessageBase) GetToOk() (*Recipients, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ScheduledMessageBase) SetTo(v Recipients)`

SetTo sets To field to given value.


### GetTopic

`func (o *ScheduledMessageBase) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ScheduledMessageBase) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ScheduledMessageBase) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ScheduledMessageBase) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetContent

`func (o *ScheduledMessageBase) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ScheduledMessageBase) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ScheduledMessageBase) SetContent(v string)`

SetContent sets Content field to given value.


### GetRenderedContent

`func (o *ScheduledMessageBase) GetRenderedContent() string`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *ScheduledMessageBase) GetRenderedContentOk() (*string, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *ScheduledMessageBase) SetRenderedContent(v string)`

SetRenderedContent sets RenderedContent field to given value.


### GetScheduledDeliveryTimestamp

`func (o *ScheduledMessageBase) GetScheduledDeliveryTimestamp() int32`

GetScheduledDeliveryTimestamp returns the ScheduledDeliveryTimestamp field if non-nil, zero value otherwise.

### GetScheduledDeliveryTimestampOk

`func (o *ScheduledMessageBase) GetScheduledDeliveryTimestampOk() (*int32, bool)`

GetScheduledDeliveryTimestampOk returns a tuple with the ScheduledDeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryTimestamp

`func (o *ScheduledMessageBase) SetScheduledDeliveryTimestamp(v int32)`

SetScheduledDeliveryTimestamp sets ScheduledDeliveryTimestamp field to given value.


### GetFailed

`func (o *ScheduledMessageBase) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ScheduledMessageBase) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ScheduledMessageBase) SetFailed(v bool)`

SetFailed sets Failed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


