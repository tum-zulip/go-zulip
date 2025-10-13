# ScheduledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledMessageId** | **interface{}** |  | 
**Type** | **interface{}** |  | 
**To** | **interface{}** |  | 
**Topic** | Pointer to **interface{}** |  | [optional] 
**Content** | **interface{}** |  | 
**RenderedContent** | **interface{}** |  | 
**ScheduledDeliveryTimestamp** | **interface{}** |  | 
**Failed** | **interface{}** |  | 

## Methods

### NewScheduledMessage

`func NewScheduledMessage(scheduledMessageId interface{}, type_ interface{}, to interface{}, content interface{}, renderedContent interface{}, scheduledDeliveryTimestamp interface{}, failed interface{}, ) *ScheduledMessage`

NewScheduledMessage instantiates a new ScheduledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledMessageWithDefaults

`func NewScheduledMessageWithDefaults() *ScheduledMessage`

NewScheduledMessageWithDefaults instantiates a new ScheduledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledMessageId

`func (o *ScheduledMessage) GetScheduledMessageId() interface{}`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *ScheduledMessage) GetScheduledMessageIdOk() (*interface{}, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *ScheduledMessage) SetScheduledMessageId(v interface{})`

SetScheduledMessageId sets ScheduledMessageId field to given value.


### SetScheduledMessageIdNil

`func (o *ScheduledMessage) SetScheduledMessageIdNil(b bool)`

 SetScheduledMessageIdNil sets the value for ScheduledMessageId to be an explicit nil

### UnsetScheduledMessageId
`func (o *ScheduledMessage) UnsetScheduledMessageId()`

UnsetScheduledMessageId ensures that no value is present for ScheduledMessageId, not even an explicit nil
### GetType

`func (o *ScheduledMessage) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledMessage) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledMessage) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ScheduledMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ScheduledMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTo

`func (o *ScheduledMessage) GetTo() interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ScheduledMessage) GetToOk() (*interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ScheduledMessage) SetTo(v interface{})`

SetTo sets To field to given value.


### SetToNil

`func (o *ScheduledMessage) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ScheduledMessage) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTopic

`func (o *ScheduledMessage) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ScheduledMessage) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ScheduledMessage) SetTopic(v interface{})`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ScheduledMessage) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *ScheduledMessage) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *ScheduledMessage) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetContent

`func (o *ScheduledMessage) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ScheduledMessage) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ScheduledMessage) SetContent(v interface{})`

SetContent sets Content field to given value.


### SetContentNil

`func (o *ScheduledMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ScheduledMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRenderedContent

`func (o *ScheduledMessage) GetRenderedContent() interface{}`

GetRenderedContent returns the RenderedContent field if non-nil, zero value otherwise.

### GetRenderedContentOk

`func (o *ScheduledMessage) GetRenderedContentOk() (*interface{}, bool)`

GetRenderedContentOk returns a tuple with the RenderedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedContent

`func (o *ScheduledMessage) SetRenderedContent(v interface{})`

SetRenderedContent sets RenderedContent field to given value.


### SetRenderedContentNil

`func (o *ScheduledMessage) SetRenderedContentNil(b bool)`

 SetRenderedContentNil sets the value for RenderedContent to be an explicit nil

### UnsetRenderedContent
`func (o *ScheduledMessage) UnsetRenderedContent()`

UnsetRenderedContent ensures that no value is present for RenderedContent, not even an explicit nil
### GetScheduledDeliveryTimestamp

`func (o *ScheduledMessage) GetScheduledDeliveryTimestamp() interface{}`

GetScheduledDeliveryTimestamp returns the ScheduledDeliveryTimestamp field if non-nil, zero value otherwise.

### GetScheduledDeliveryTimestampOk

`func (o *ScheduledMessage) GetScheduledDeliveryTimestampOk() (*interface{}, bool)`

GetScheduledDeliveryTimestampOk returns a tuple with the ScheduledDeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryTimestamp

`func (o *ScheduledMessage) SetScheduledDeliveryTimestamp(v interface{})`

SetScheduledDeliveryTimestamp sets ScheduledDeliveryTimestamp field to given value.


### SetScheduledDeliveryTimestampNil

`func (o *ScheduledMessage) SetScheduledDeliveryTimestampNil(b bool)`

 SetScheduledDeliveryTimestampNil sets the value for ScheduledDeliveryTimestamp to be an explicit nil

### UnsetScheduledDeliveryTimestamp
`func (o *ScheduledMessage) UnsetScheduledDeliveryTimestamp()`

UnsetScheduledDeliveryTimestamp ensures that no value is present for ScheduledDeliveryTimestamp, not even an explicit nil
### GetFailed

`func (o *ScheduledMessage) GetFailed() interface{}`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ScheduledMessage) GetFailedOk() (*interface{}, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ScheduledMessage) SetFailed(v interface{})`

SetFailed sets Failed field to given value.


### SetFailedNil

`func (o *ScheduledMessage) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ScheduledMessage) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


