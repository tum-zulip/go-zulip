# ScheduledMessagesAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ScheduledMessages** | Pointer to [**[]ScheduledMessage**](ScheduledMessage.md) | An array of objects containing details of the newly created scheduled messages.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf80

`func NewGetEvents200ResponseAllOfEventsInnerOneOf80() *ScheduledMessagesAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf80 instantiates a new ScheduledMessagesAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf80WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf80WithDefaults() *ScheduledMessagesAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf80WithDefaults instantiates a new ScheduledMessagesAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledMessagesAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledMessagesAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledMessagesAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledMessagesAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ScheduledMessagesAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledMessagesAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledMessagesAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduledMessagesAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ScheduledMessagesAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ScheduledMessagesAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ScheduledMessagesAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ScheduledMessagesAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetScheduledMessages

`func (o *ScheduledMessagesAddEvent) GetScheduledMessages() []ScheduledMessage`

GetScheduledMessages returns the ScheduledMessages field if non-nil, zero value otherwise.

### GetScheduledMessagesOk

`func (o *ScheduledMessagesAddEvent) GetScheduledMessagesOk() (*[]ScheduledMessage, bool)`

GetScheduledMessagesOk returns a tuple with the ScheduledMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessages

`func (o *ScheduledMessagesAddEvent) SetScheduledMessages(v []ScheduledMessage)`

SetScheduledMessages sets ScheduledMessages field to given value.

### HasScheduledMessages

`func (o *ScheduledMessagesAddEvent) HasScheduledMessages() bool`

HasScheduledMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


