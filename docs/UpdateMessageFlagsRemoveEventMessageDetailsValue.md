# UpdateMessageFlagsRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** | Old name for the &#x60;op&#x60; field in this event type.  **Deprecated** in Zulip 4.0 (feature level 32), and replaced by the &#x60;op&#x60; field.  | [optional] 
**Flag** | Pointer to **string** | The [flag][message-flags] to be removed.  | [optional] 
**Messages** | Pointer to **[]int32** | Array containing the IDs of the messages from which the flag was removed.  | [optional] 
**All** | Pointer to **bool** | Will be &#x60;false&#x60; for all specified flags.  **Deprecated** and will be removed in a future release.  | [optional] 
**MessageDetails** | Pointer to [**map[string]EventEnvelopeOneOf42MessageDetailsValue**](EventEnvelopeOneOf42MessageDetailsValue.md) | Only present if the specified &#x60;flag&#x60; is &#x60;\&quot;read\&quot;&#x60;.  A set of data structures describing the messages that are being marked as unread with additional details to allow clients to update the &#x60;unread_msgs&#x60; data structure for these messages (which may not be otherwise known to the client).  **Changes**: New in Zulip 5.0 (feature level 121). Previously, marking already read messages as unread was not supported by the Zulip API.  | [optional] 

## Methods

### NewEventEnvelopeOneOf42

`func NewEventEnvelopeOneOf42() *UpdateMessageFlagsRemoveEvent`

NewEventEnvelopeOneOf42 instantiates a new UpdateMessageFlagsRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf42WithDefaults

`func NewEventEnvelopeOneOf42WithDefaults() *UpdateMessageFlagsRemoveEvent`

NewEventEnvelopeOneOf42WithDefaults instantiates a new UpdateMessageFlagsRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMessageFlagsRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMessageFlagsRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMessageFlagsRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMessageFlagsRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateMessageFlagsRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMessageFlagsRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMessageFlagsRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateMessageFlagsRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UpdateMessageFlagsRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UpdateMessageFlagsRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UpdateMessageFlagsRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UpdateMessageFlagsRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOperation

`func (o *UpdateMessageFlagsRemoveEvent) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *UpdateMessageFlagsRemoveEvent) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *UpdateMessageFlagsRemoveEvent) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *UpdateMessageFlagsRemoveEvent) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetFlag

`func (o *UpdateMessageFlagsRemoveEvent) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *UpdateMessageFlagsRemoveEvent) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *UpdateMessageFlagsRemoveEvent) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *UpdateMessageFlagsRemoveEvent) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetMessages

`func (o *UpdateMessageFlagsRemoveEvent) GetMessages() []int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *UpdateMessageFlagsRemoveEvent) GetMessagesOk() (*[]int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *UpdateMessageFlagsRemoveEvent) SetMessages(v []int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *UpdateMessageFlagsRemoveEvent) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAll

`func (o *UpdateMessageFlagsRemoveEvent) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateMessageFlagsRemoveEvent) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateMessageFlagsRemoveEvent) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateMessageFlagsRemoveEvent) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetMessageDetails

`func (o *UpdateMessageFlagsRemoveEvent) GetMessageDetails() map[string]EventEnvelopeOneOf42MessageDetailsValue`

GetMessageDetails returns the MessageDetails field if non-nil, zero value otherwise.

### GetMessageDetailsOk

`func (o *UpdateMessageFlagsRemoveEvent) GetMessageDetailsOk() (*map[string]EventEnvelopeOneOf42MessageDetailsValue, bool)`

GetMessageDetailsOk returns a tuple with the MessageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetails

`func (o *UpdateMessageFlagsRemoveEvent) SetMessageDetails(v map[string]EventEnvelopeOneOf42MessageDetailsValue)`

SetMessageDetails sets MessageDetails field to given value.

### HasMessageDetails

`func (o *UpdateMessageFlagsRemoveEvent) HasMessageDetails() bool`

HasMessageDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


