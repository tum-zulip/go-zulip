# TypingEditMessageStartEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**SenderId** | Pointer to **int32** | The ID of the user who is typing the edit of the message.  Clients should be careful to display this user as the person who is typing, not that of the sender of the message, in case a collaborative editing feature be might be added in the future.  | [optional] 
**MessageId** | Pointer to **int32** | Indicates the message id of the message that is being edited.  | [optional] 
**Recipient** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf39Recipient**](GetEvents200ResponseAllOfEventsInnerOneOf39Recipient.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf39

`func NewGetEvents200ResponseAllOfEventsInnerOneOf39() *TypingEditMessageStartEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf39 instantiates a new TypingEditMessageStartEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf39WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf39WithDefaults() *TypingEditMessageStartEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf39WithDefaults instantiates a new TypingEditMessageStartEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypingEditMessageStartEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypingEditMessageStartEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypingEditMessageStartEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TypingEditMessageStartEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TypingEditMessageStartEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypingEditMessageStartEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypingEditMessageStartEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypingEditMessageStartEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *TypingEditMessageStartEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TypingEditMessageStartEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TypingEditMessageStartEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TypingEditMessageStartEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSenderId

`func (o *TypingEditMessageStartEvent) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *TypingEditMessageStartEvent) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *TypingEditMessageStartEvent) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *TypingEditMessageStartEvent) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetMessageId

`func (o *TypingEditMessageStartEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *TypingEditMessageStartEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *TypingEditMessageStartEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *TypingEditMessageStartEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetRecipient

`func (o *TypingEditMessageStartEvent) GetRecipient() GetEvents200ResponseAllOfEventsInnerOneOf39Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TypingEditMessageStartEvent) GetRecipientOk() (*GetEvents200ResponseAllOfEventsInnerOneOf39Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TypingEditMessageStartEvent) SetRecipient(v GetEvents200ResponseAllOfEventsInnerOneOf39Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TypingEditMessageStartEvent) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


