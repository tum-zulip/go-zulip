# SubmessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MsgType** | Pointer to **string** | The type of the message.  | [optional] 
**Content** | Pointer to **string** | The new content of the submessage.  | [optional] 
**MessageId** | Pointer to **int32** | The ID of the message to which the submessage has been added.  | [optional] 
**SenderId** | Pointer to **int32** | The ID of the user who sent the message.  | [optional] 
**SubmessageId** | Pointer to **int32** | The ID of the submessage.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf25

`func NewGetEvents200ResponseAllOfEventsInnerOneOf25() *SubmessageEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf25 instantiates a new SubmessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf25WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf25WithDefaults() *SubmessageEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf25WithDefaults instantiates a new SubmessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmessageEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmessageEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmessageEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubmessageEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubmessageEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmessageEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmessageEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubmessageEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMsgType

`func (o *SubmessageEvent) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *SubmessageEvent) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *SubmessageEvent) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *SubmessageEvent) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetContent

`func (o *SubmessageEvent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SubmessageEvent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SubmessageEvent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SubmessageEvent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMessageId

`func (o *SubmessageEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SubmessageEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SubmessageEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SubmessageEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSenderId

`func (o *SubmessageEvent) GetSenderId() int32`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *SubmessageEvent) GetSenderIdOk() (*int32, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *SubmessageEvent) SetSenderId(v int32)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *SubmessageEvent) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSubmessageId

`func (o *SubmessageEvent) GetSubmessageId() int32`

GetSubmessageId returns the SubmessageId field if non-nil, zero value otherwise.

### GetSubmessageIdOk

`func (o *SubmessageEvent) GetSubmessageIdOk() (*int32, bool)`

GetSubmessageIdOk returns a tuple with the SubmessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmessageId

`func (o *SubmessageEvent) SetSubmessageId(v int32)`

SetSubmessageId sets SubmessageId field to given value.

### HasSubmessageId

`func (o *SubmessageEvent) HasSubmessageId() bool`

HasSubmessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


