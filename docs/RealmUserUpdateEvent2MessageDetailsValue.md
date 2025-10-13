# RealmUserUpdateEvent2MessageDetailsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this message. Either &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;private\&quot;&#x60;.  | 
**Mentioned** | Pointer to **bool** | A flag which indicates whether the message contains a mention of the user.  Present only if the message mentions the current user.  | [optional] 
**UserIds** | Pointer to **[]int32** | Present only if &#x60;type&#x60; is &#x60;private&#x60;.  The user IDs of every recipient of this direct message, excluding yourself. Will be the empty list for a message you had sent to only yourself.  | [optional] 
**StreamId** | Pointer to **int32** | Present only if &#x60;type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The ID of the channel where the message was sent.  | [optional] 
**Topic** | Pointer to **string** | Present only if &#x60;type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  Name of the topic where the message was sent.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**UnmutedStreamMsg** | Pointer to **bool** | **Deprecated** internal implementation detail. Clients should ignore this field as it will be removed in the future.  | [optional] 

## Methods

### NewEventEnvelopeOneOf42MessageDetailsValue

`func NewEventEnvelopeOneOf42MessageDetailsValue(type_ string, ) *EventEnvelopeOneOf42MessageDetailsValue`

NewEventEnvelopeOneOf42MessageDetailsValue instantiates a new EventEnvelopeOneOf42MessageDetailsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf42MessageDetailsValueWithDefaults

`func NewEventEnvelopeOneOf42MessageDetailsValueWithDefaults() *EventEnvelopeOneOf42MessageDetailsValue`

NewEventEnvelopeOneOf42MessageDetailsValueWithDefaults instantiates a new EventEnvelopeOneOf42MessageDetailsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetType(v string)`

SetType sets Type field to given value.


### GetMentioned

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetMentioned() bool`

GetMentioned returns the Mentioned field if non-nil, zero value otherwise.

### GetMentionedOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetMentionedOk() (*bool, bool)`

GetMentionedOk returns a tuple with the Mentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentioned

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetMentioned(v bool)`

SetMentioned sets Mentioned field to given value.

### HasMentioned

`func (o *EventEnvelopeOneOf42MessageDetailsValue) HasMentioned() bool`

HasMentioned returns a boolean if a field has been set.

### GetUserIds

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *EventEnvelopeOneOf42MessageDetailsValue) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetStreamId

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *EventEnvelopeOneOf42MessageDetailsValue) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopic

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *EventEnvelopeOneOf42MessageDetailsValue) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUnmutedStreamMsg

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetUnmutedStreamMsg() bool`

GetUnmutedStreamMsg returns the UnmutedStreamMsg field if non-nil, zero value otherwise.

### GetUnmutedStreamMsgOk

`func (o *EventEnvelopeOneOf42MessageDetailsValue) GetUnmutedStreamMsgOk() (*bool, bool)`

GetUnmutedStreamMsgOk returns a tuple with the UnmutedStreamMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmutedStreamMsg

`func (o *EventEnvelopeOneOf42MessageDetailsValue) SetUnmutedStreamMsg(v bool)`

SetUnmutedStreamMsg sets UnmutedStreamMsg field to given value.

### HasUnmutedStreamMsg

`func (o *EventEnvelopeOneOf42MessageDetailsValue) HasUnmutedStreamMsg() bool`

HasUnmutedStreamMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


