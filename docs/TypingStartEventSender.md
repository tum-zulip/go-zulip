# TypingStartEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** | Type of message being composed. Must be &#x60;\&quot;stream\&quot;&#x60; or &#x60;\&quot;direct\&quot;&#x60;.  **Changes**: In Zulip 8.0 (feature level 215), replaced the value &#x60;\&quot;private\&quot;&#x60; with &#x60;\&quot;direct\&quot;&#x60;.  New in Zulip 4.0 (feature level 58). Previously, all typing notifications were implicitly direct messages.  | [optional] 
**Sender** | Pointer to [**EventEnvelopeOneOf37Sender**](EventEnvelopeOneOf37Sender.md) |  | [optional] 
**Recipients** | Pointer to [**[]EventEnvelopeOneOf37RecipientsInner**](EventEnvelopeOneOf37RecipientsInner.md) | Only present if &#x60;message_type&#x60; is &#x60;\&quot;direct\&quot;&#x60;.  Array of dictionaries describing the set of users who would be recipients of the message being typed. Each dictionary contains details about one of the recipients. The sending user is guaranteed to appear among the recipients.  | [optional] 
**StreamId** | Pointer to **int32** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  The unique ID of the channel to which message is being typed.  **Changes**: New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  | [optional] 
**Topic** | Pointer to **string** | Only present if &#x60;message_type&#x60; is &#x60;\&quot;stream\&quot;&#x60;.  Topic within the channel where the message is being typed.  For clients that don&#39;t support the &#x60;empty_topic_name&#x60; [client capability][client-capabilities], if the actual topic name is empty string, this field&#39;s value will instead be the value of &#x60;realm_empty_topic_display_name&#x60; found in the [&#x60;POST /register&#x60;](zulip.com/api/register-queue response.  **Changes**: Before 10.0 (feature level 334), &#x60;empty_topic_name&#x60; client capability didn&#39;t exist and empty string as the topic name for channel messages wasn&#39;t allowed.  New in Zulip 4.0 (feature level 58). Previously, typing notifications were only for direct messages.  [client-capabilities]: /api/register-queue#parameter-client_capabilities  | [optional] 

## Methods

### NewEventEnvelopeOneOf37

`func NewEventEnvelopeOneOf37() *TypingStartEvent`

NewEventEnvelopeOneOf37 instantiates a new TypingStartEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf37WithDefaults

`func NewEventEnvelopeOneOf37WithDefaults() *TypingStartEvent`

NewEventEnvelopeOneOf37WithDefaults instantiates a new TypingStartEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypingStartEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypingStartEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypingStartEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TypingStartEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TypingStartEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypingStartEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypingStartEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypingStartEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *TypingStartEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TypingStartEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TypingStartEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TypingStartEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetMessageType

`func (o *TypingStartEvent) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *TypingStartEvent) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *TypingStartEvent) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *TypingStartEvent) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSender

`func (o *TypingStartEvent) GetSender() EventEnvelopeOneOf37Sender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TypingStartEvent) GetSenderOk() (*EventEnvelopeOneOf37Sender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TypingStartEvent) SetSender(v EventEnvelopeOneOf37Sender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TypingStartEvent) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipients

`func (o *TypingStartEvent) GetRecipients() []EventEnvelopeOneOf37RecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *TypingStartEvent) GetRecipientsOk() (*[]EventEnvelopeOneOf37RecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *TypingStartEvent) SetRecipients(v []EventEnvelopeOneOf37RecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *TypingStartEvent) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetStreamId

`func (o *TypingStartEvent) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *TypingStartEvent) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *TypingStartEvent) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *TypingStartEvent) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTopic

`func (o *TypingStartEvent) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TypingStartEvent) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TypingStartEvent) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TypingStartEvent) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


