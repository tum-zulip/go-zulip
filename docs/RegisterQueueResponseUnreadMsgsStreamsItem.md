# RegisterQueueResponseUnreadMsgsStreamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The topic under which the messages were sent.  Note that the empty string topic may have been rewritten by the server to the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response depending on the value of the &#x60;empty_topic_name&#x60; [client capability][client-capabilities].  **Changes**: The &#x60;empty_topic_name&#x60; client capability is new in Zulip 10.0 (feature level 334).  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**StreamId** | Pointer to **int32** | The ID of the channel to which the messages were sent.  | [optional] 
**UnreadMessageIds** | Pointer to **[]int32** | The message IDs of the recent unread messages sent in this channel, sorted in ascending order.  | [optional] 

## Methods

### NewRegisterQueueResponseUnreadMsgsStreamsInner

`func NewRegisterQueueResponseUnreadMsgsStreamsInner() *RegisterQueueResponseUnreadMsgsStreamsInner`

NewRegisterQueueResponseUnreadMsgsStreamsInner instantiates a new RegisterQueueResponseUnreadMsgsStreamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseUnreadMsgsStreamsInnerWithDefaults

`func NewRegisterQueueResponseUnreadMsgsStreamsInnerWithDefaults() *RegisterQueueResponseUnreadMsgsStreamsInner`

NewRegisterQueueResponseUnreadMsgsStreamsInnerWithDefaults instantiates a new RegisterQueueResponseUnreadMsgsStreamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetStreamId

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetUnreadMessageIds() []int32`

GetUnreadMessageIds returns the UnreadMessageIds field if non-nil, zero value otherwise.

### GetUnreadMessageIdsOk

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) GetUnreadMessageIdsOk() (*[]int32, bool)`

GetUnreadMessageIdsOk returns a tuple with the UnreadMessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) SetUnreadMessageIds(v []int32)`

SetUnreadMessageIds sets UnreadMessageIds field to given value.

### HasUnreadMessageIds

`func (o *RegisterQueueResponseUnreadMsgsStreamsInner) HasUnreadMessageIds() bool`

HasUnreadMessageIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


