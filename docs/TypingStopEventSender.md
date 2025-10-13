# TypingStopEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** | Type of message being composed. Must be &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;direct\&quot;&#x60;.  **Changes**: In Zulip 8.0 (feature level 215), replaced the value &#x60;\&quot;private\&quot;&#x60; with &#x60;\&quot;direct\&quot;&#x60;.  New in Zulip 4.0 (feature level 58). Previously all typing notifications were implicitly direct messages.  | [optional] 
**Sender** | Pointer to [**EventEnvelopeOneOf38Sender**](EventEnvelopeOneOf38Sender.md) |  | [optional] 
**Recipients** | Pointer to [**[]EventEnvelopeOneOf38RecipientsInner**](EventEnvelopeOneOf38RecipientsInner.md) | Only present if &#x60;message_type&#x60; is &#x60;\&quot;direct\&quot;&#x60;.  Array of dictionaries describing the set of users who would be recipients of the message that was previously being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The unique ID of the channel to which message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 
**Topic** | Pointer to **string** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  Topic within the channel where the message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 

## Methods

### NewEventEnvelopeOneOf38

`func NewEventEnvelopeOneOf38() *TypingStopEvent`

NewEventEnvelopeOneOf38 instantiates a new TypingStopEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf38WithDefaults

`func NewEventEnvelopeOneOf38WithDefaults() *TypingStopEvent`

NewEventEnvelopeOneOf38WithDefaults instantiates a new TypingStopEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypingStopEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypingStopEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypingStopEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TypingStopEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TypingStopEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypingStopEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypingStopEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypingStopEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *TypingStopEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TypingStopEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TypingStopEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TypingStopEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetMessageType

`func (o *TypingStopEvent) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *TypingStopEvent) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *TypingStopEvent) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *TypingStopEvent) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSender

`func (o *TypingStopEvent) GetSender() EventEnvelopeOneOf38Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TypingStopEvent) GetSenderOk() (*EventEnvelopeOneOf38Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TypingStopEvent) SetSender(v EventEnvelopeOneOf38Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TypingStopEvent) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipients

`func (o *TypingStopEvent) GetRecipients() []EventEnvelopeOneOf38RecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *TypingStopEvent) GetRecipientsOk() (*[]EventEnvelopeOneOf38RecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *TypingStopEvent) SetRecipients(v []EventEnvelopeOneOf38RecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *TypingStopEvent) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetStreamId

`func (o *TypingStopEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *TypingStopEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *TypingStopEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *TypingStopEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopic

`func (o *TypingStopEvent) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TypingStopEvent) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TypingStopEvent) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TypingStopEvent) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


