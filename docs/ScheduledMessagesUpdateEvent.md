# ScheduledMessagesUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ScheduledMessage** | Pointer to [**ScheduledMessage**](ScheduledMessage.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf81

`func NewEventEnvelopeOneOf81() *ScheduledMessagesUpdateEvent`

NewEventEnvelopeOneOf81 instantiates a new ScheduledMessagesUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf81WithDefaults

`func NewEventEnvelopeOneOf81WithDefaults() *ScheduledMessagesUpdateEvent`

NewEventEnvelopeOneOf81WithDefaults instantiates a new ScheduledMessagesUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledMessagesUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledMessagesUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledMessagesUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledMessagesUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ScheduledMessagesUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledMessagesUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledMessagesUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduledMessagesUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ScheduledMessagesUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ScheduledMessagesUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ScheduledMessagesUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ScheduledMessagesUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetScheduledMessage

`func (o *ScheduledMessagesUpdateEvent) GetScheduledMessage() ScheduledMessage`

GetScheduledMessage returns the ScheduledMessage field if non-nil, zero value otherwise.

### GetScheduledMessageOk

`func (o *ScheduledMessagesUpdateEvent) GetScheduledMessageOk() (*ScheduledMessage, bool)`

GetScheduledMessageOk returns a tuple with the ScheduledMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessage

`func (o *ScheduledMessagesUpdateEvent) SetScheduledMessage(v ScheduledMessage)`

SetScheduledMessage sets ScheduledMessage field to given value.

### HasScheduledMessage

`func (o *ScheduledMessagesUpdateEvent) HasScheduledMessage() bool`

HasScheduledMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


