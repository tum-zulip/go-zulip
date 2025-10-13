# TypingEditMessageStopEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**SenderId** | Pointer to **int32** | The ID of the user who sent the message.  | [optional] 
**MessageId** | Pointer to **int32** | Indicates the message id of the message that is being edited.  | [optional] 
**Recipient** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf39Recipient**](GetEvents200ResponseAllOfEventsInnerOneOf39Recipient.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf40

`func NewGetEvents200ResponseAllOfEventsInnerOneOf40() *TypingEditMessageStopEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf40 instantiates a new TypingEditMessageStopEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf40WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf40WithDefaults() *TypingEditMessageStopEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf40WithDefaults instantiates a new TypingEditMessageStopEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypingEditMessageStopEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypingEditMessageStopEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypingEditMessageStopEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TypingEditMessageStopEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TypingEditMessageStopEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypingEditMessageStopEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypingEditMessageStopEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypingEditMessageStopEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *TypingEditMessageStopEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TypingEditMessageStopEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TypingEditMessageStopEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TypingEditMessageStopEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSenderId

`func (o *TypingEditMessageStopEvent) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *TypingEditMessageStopEvent) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *TypingEditMessageStopEvent) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *TypingEditMessageStopEvent) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetMessageId

`func (o *TypingEditMessageStopEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *TypingEditMessageStopEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *TypingEditMessageStopEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *TypingEditMessageStopEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetRecipient

`func (o *TypingEditMessageStopEvent) GetRecipient() GetEvents200ResponseAllOfEventsInnerOneOf39Recipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TypingEditMessageStopEvent) GetRecipientOk() (*GetEvents200ResponseAllOfEventsInnerOneOf39Recipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TypingEditMessageStopEvent) SetRecipient(v GetEvents200ResponseAllOfEventsInnerOneOf39Recipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TypingEditMessageStopEvent) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


