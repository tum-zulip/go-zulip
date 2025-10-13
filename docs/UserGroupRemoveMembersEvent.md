# UserGroupRemoveMembersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the user group whose details have changed.  | [optional] 
**UserIds** | Pointer to **[]int32** | Array containing the IDs of the users who have been removed from the user group.  | [optional] 

## Methods

### NewEventEnvelopeOneOf46

`func NewEventEnvelopeOneOf46() *UserGroupRemoveMembersEvent`

NewEventEnvelopeOneOf46 instantiates a new UserGroupRemoveMembersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf46WithDefaults

`func NewEventEnvelopeOneOf46WithDefaults() *UserGroupRemoveMembersEvent`

NewEventEnvelopeOneOf46WithDefaults instantiates a new UserGroupRemoveMembersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupRemoveMembersEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupRemoveMembersEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupRemoveMembersEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupRemoveMembersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupRemoveMembersEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupRemoveMembersEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupRemoveMembersEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupRemoveMembersEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserGroupRemoveMembersEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserGroupRemoveMembersEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserGroupRemoveMembersEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserGroupRemoveMembersEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetGroupId

`func (o *UserGroupRemoveMembersEvent) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserGroupRemoveMembersEvent) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserGroupRemoveMembersEvent) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserGroupRemoveMembersEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserIds

`func (o *UserGroupRemoveMembersEvent) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UserGroupRemoveMembersEvent) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UserGroupRemoveMembersEvent) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UserGroupRemoveMembersEvent) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


