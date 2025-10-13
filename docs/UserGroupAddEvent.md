# UserGroupAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf43

`func NewEventEnvelopeOneOf43() *UserGroupAddEvent`

NewEventEnvelopeOneOf43 instantiates a new UserGroupAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf43WithDefaults

`func NewEventEnvelopeOneOf43WithDefaults() *UserGroupAddEvent`

NewEventEnvelopeOneOf43WithDefaults instantiates a new UserGroupAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserGroupAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserGroupAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserGroupAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserGroupAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetGroup

`func (o *UserGroupAddEvent) GetGroup() UserGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserGroupAddEvent) GetGroupOk() (*UserGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserGroupAddEvent) SetGroup(v UserGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserGroupAddEvent) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


