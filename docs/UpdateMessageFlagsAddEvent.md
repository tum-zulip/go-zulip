# UpdateMessageFlagsAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** | Old name for the &#x60;op&#x60; field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the &#x60;op&#x60; field.  | [optional] 
**Flag** | Pointer to **string** | The [flag][message-flags] that was added.  | [optional] 
**Messages** | Pointer to **[]int32** | Array containing the IDs of all messages to which the flag was added.  | [optional] 
**All** | Pointer to **bool** | Whether the specified flag was added to all messages. This field is only relevant for the &#x60;\&quot;read\&quot;&#x60; flag, and will be &#x60;false&#x60; for all other flags.  When &#x60;true&#x60; for the &#x60;\&quot;read\&quot;&#x60; flag, then the &#x60;messages&#x60; array will be empty.  | [optional] 

## Methods

### NewEventEnvelopeOneOf41

`func NewEventEnvelopeOneOf41() *UpdateMessageFlagsAddEvent`

NewEventEnvelopeOneOf41 instantiates a new UpdateMessageFlagsAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf41WithDefaults

`func NewEventEnvelopeOneOf41WithDefaults() *UpdateMessageFlagsAddEvent`

NewEventEnvelopeOneOf41WithDefaults instantiates a new UpdateMessageFlagsAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMessageFlagsAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMessageFlagsAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMessageFlagsAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMessageFlagsAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateMessageFlagsAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMessageFlagsAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMessageFlagsAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateMessageFlagsAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UpdateMessageFlagsAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UpdateMessageFlagsAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UpdateMessageFlagsAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UpdateMessageFlagsAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOperation

`func (o *UpdateMessageFlagsAddEvent) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *UpdateMessageFlagsAddEvent) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *UpdateMessageFlagsAddEvent) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *UpdateMessageFlagsAddEvent) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetFlag

`func (o *UpdateMessageFlagsAddEvent) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *UpdateMessageFlagsAddEvent) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *UpdateMessageFlagsAddEvent) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *UpdateMessageFlagsAddEvent) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetMessages

`func (o *UpdateMessageFlagsAddEvent) GetMessages() []int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *UpdateMessageFlagsAddEvent) GetMessagesOk() (*[]int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *UpdateMessageFlagsAddEvent) SetMessages(v []int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *UpdateMessageFlagsAddEvent) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAll

`func (o *UpdateMessageFlagsAddEvent) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateMessageFlagsAddEvent) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateMessageFlagsAddEvent) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateMessageFlagsAddEvent) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


