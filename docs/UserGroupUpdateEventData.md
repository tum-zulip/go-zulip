# UserGroupUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the user group whose details have changed.  | [optional] 
**Data** | Pointer to [**EventEnvelopeOneOf44Data**](EventEnvelopeOneOf44Data.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf44

`func NewEventEnvelopeOneOf44() *UserGroupUpdateEvent`

NewEventEnvelopeOneOf44 instantiates a new UserGroupUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf44WithDefaults

`func NewEventEnvelopeOneOf44WithDefaults() *UserGroupUpdateEvent`

NewEventEnvelopeOneOf44WithDefaults instantiates a new UserGroupUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserGroupUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserGroupUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserGroupUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserGroupUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetGroupId

`func (o *UserGroupUpdateEvent) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserGroupUpdateEvent) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserGroupUpdateEvent) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserGroupUpdateEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *UserGroupUpdateEvent) GetData() EventEnvelopeOneOf44Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserGroupUpdateEvent) GetDataOk() (*EventEnvelopeOneOf44Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserGroupUpdateEvent) SetData(v EventEnvelopeOneOf44Data)`

SetData sets Data field to given value.

### HasData

`func (o *UserGroupUpdateEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


