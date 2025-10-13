# ChannelFolderReorderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **[]int32** | A list of channel folder IDs representing the new order.  | [optional] 

## Methods

### NewEventEnvelopeOneOf85

`func NewEventEnvelopeOneOf85() *ChannelFolderReorderEvent`

NewEventEnvelopeOneOf85 instantiates a new ChannelFolderReorderEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf85WithDefaults

`func NewEventEnvelopeOneOf85WithDefaults() *ChannelFolderReorderEvent`

NewEventEnvelopeOneOf85WithDefaults instantiates a new ChannelFolderReorderEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelFolderReorderEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFolderReorderEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFolderReorderEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelFolderReorderEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ChannelFolderReorderEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelFolderReorderEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelFolderReorderEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelFolderReorderEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ChannelFolderReorderEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ChannelFolderReorderEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ChannelFolderReorderEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ChannelFolderReorderEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOrder

`func (o *ChannelFolderReorderEvent) GetOrder() []int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ChannelFolderReorderEvent) GetOrderOk() (*[]int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ChannelFolderReorderEvent) SetOrder(v []int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ChannelFolderReorderEvent) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


