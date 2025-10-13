# DeleteMessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MessageIds** | Pointer to **[]int32** | Only present for clients that support the &#x60;bulk_message_deletion&#x60; [client capability][client-capabilities].  A sorted list containing the IDs of the newly deleted messages.  **Changes**: Before Zulip 11.0 (feature level 393), this list was not guaranteed to be sorted.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**MessageId** | Pointer to **int32** | Only present for clients that do not support the &#x60;bulk_message_deletion&#x60; [client capability][client-capabilities].  The ID of the newly deleted message.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 
**MessageType** | Pointer to **string** | The type of message. Either &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;private\&quot;&#x60;.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The ID of the channel to which the message was sent.  | [optional] 
**Topic** | Pointer to **string** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The topic to which the message was sent.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name was empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](/api/register-queue) response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 

## Methods

### NewEventEnvelopeOneOf30

`func NewEventEnvelopeOneOf30() *DeleteMessageEvent`

NewEventEnvelopeOneOf30 instantiates a new DeleteMessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf30WithDefaults

`func NewEventEnvelopeOneOf30WithDefaults() *DeleteMessageEvent`

NewEventEnvelopeOneOf30WithDefaults instantiates a new DeleteMessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteMessageEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteMessageEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteMessageEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteMessageEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DeleteMessageEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeleteMessageEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeleteMessageEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeleteMessageEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessageIds

`func (o *DeleteMessageEvent) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *DeleteMessageEvent) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *DeleteMessageEvent) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.

### HasMessageIds

`func (o *DeleteMessageEvent) HasMessageIds() bool`

HasMessageIds returns a boolean if a field has been set.

### GetMessageId

`func (o *DeleteMessageEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *DeleteMessageEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *DeleteMessageEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *DeleteMessageEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *DeleteMessageEvent) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *DeleteMessageEvent) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *DeleteMessageEvent) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *DeleteMessageEvent) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetStreamId

`func (o *DeleteMessageEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *DeleteMessageEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *DeleteMessageEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *DeleteMessageEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopic

`func (o *DeleteMessageEvent) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *DeleteMessageEvent) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *DeleteMessageEvent) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *DeleteMessageEvent) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


