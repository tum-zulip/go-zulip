# UserGroupAddMembersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int32** | The ID of the user group with new members.  | [optional] 
**UserIds** | Pointer to **[]int32** | Array containing the IDs of the users who have been added to the user group.  | [optional] 

## Methods

### NewEventEnvelopeOneOf45

`func NewEventEnvelopeOneOf45() *UserGroupAddMembersEvent`

NewEventEnvelopeOneOf45 instantiates a new UserGroupAddMembersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf45WithDefaults

`func NewEventEnvelopeOneOf45WithDefaults() *UserGroupAddMembersEvent`

NewEventEnvelopeOneOf45WithDefaults instantiates a new UserGroupAddMembersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroupAddMembersEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroupAddMembersEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroupAddMembersEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroupAddMembersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserGroupAddMembersEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroupAddMembersEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroupAddMembersEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroupAddMembersEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserGroupAddMembersEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserGroupAddMembersEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserGroupAddMembersEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserGroupAddMembersEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetGroupId

`func (o *UserGroupAddMembersEvent) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UserGroupAddMembersEvent) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UserGroupAddMembersEvent) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UserGroupAddMembersEvent) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUserIds

`func (o *UserGroupAddMembersEvent) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UserGroupAddMembersEvent) GetUserIdsOk() (*[]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UserGroupAddMembersEvent) SetUserIds(v []int32)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UserGroupAddMembersEvent) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


