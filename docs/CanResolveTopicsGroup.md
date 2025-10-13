# CanResolveTopicsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectMembers** | Pointer to **[]int32** | The list of IDs of individual users in the collection of users with this permission.  **Changes**: Prior to Zulip 10.0 (feature level 303), this list would include deactivated users who had the permission before being deactivated.  | [optional] 
**DirectSubgroups** | Pointer to **[]int32** | The list of IDs of the groups in the collection of users with this permission.  | [optional] 

## Methods

### NewCanResolveTopicsGroup

`func NewCanResolveTopicsGroup() *CanResolveTopicsGroup`

NewCanResolveTopicsGroup instantiates a new CanResolveTopicsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanResolveTopicsGroupWithDefaults

`func NewCanResolveTopicsGroupWithDefaults() *CanResolveTopicsGroup`

NewCanResolveTopicsGroupWithDefaults instantiates a new CanResolveTopicsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectMembers

`func (o *CanResolveTopicsGroup) GetDirectMembers() []int32`

GetDirectMembers returns the DirectMembers field if non-nil, zero value otherwise.

### GetDirectMembersOk

`func (o *CanResolveTopicsGroup) GetDirectMembersOk() (*[]int32, bool)`

GetDirectMembersOk returns a tuple with the DirectMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMembers

`func (o *CanResolveTopicsGroup) SetDirectMembers(v []int32)`

SetDirectMembers sets DirectMembers field to given value.

### HasDirectMembers

`func (o *CanResolveTopicsGroup) HasDirectMembers() bool`

HasDirectMembers returns a boolean if a field has been set.

### GetDirectSubgroups

`func (o *CanResolveTopicsGroup) GetDirectSubgroups() []int32`

GetDirectSubgroups returns the DirectSubgroups field if non-nil, zero value otherwise.

### GetDirectSubgroupsOk

`func (o *CanResolveTopicsGroup) GetDirectSubgroupsOk() (*[]int32, bool)`

GetDirectSubgroupsOk returns a tuple with the DirectSubgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroups

`func (o *CanResolveTopicsGroup) SetDirectSubgroups(v []int32)`

SetDirectSubgroups sets DirectSubgroups field to given value.

### HasDirectSubgroups

`func (o *CanResolveTopicsGroup) HasDirectSubgroups() bool`

HasDirectSubgroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


