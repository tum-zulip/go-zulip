# GroupSettingValueOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectMembers** | Pointer to **[]int32** | The list of IDs of individual users in the collection of users with this permission.  **Changes**: Prior to Zulip 10.0 (feature level 303), this list would include deactivated users who had the permission before being deactivated.  | [optional] 
**DirectSubgroups** | Pointer to **[]int32** | The list of IDs of the groups in the collection of users with this permission.  | [optional] 

## Methods

### NewGroupSettingValueOneOf

`func NewGroupSettingValueOneOf() *GroupSettingValueOneOf`

NewGroupSettingValueOneOf instantiates a new GroupSettingValueOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingValueOneOfWithDefaults

`func NewGroupSettingValueOneOfWithDefaults() *GroupSettingValueOneOf`

NewGroupSettingValueOneOfWithDefaults instantiates a new GroupSettingValueOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectMembers

`func (o *GroupSettingValueOneOf) GetDirectMembers() []int32`

GetDirectMembers returns the DirectMembers field if non-nil, zero value otherwise.

### GetDirectMembersOk

`func (o *GroupSettingValueOneOf) GetDirectMembersOk() (*[]int32, bool)`

GetDirectMembersOk returns a tuple with the DirectMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMembers

`func (o *GroupSettingValueOneOf) SetDirectMembers(v []int32)`

SetDirectMembers sets DirectMembers field to given value.

### HasDirectMembers

`func (o *GroupSettingValueOneOf) HasDirectMembers() bool`

HasDirectMembers returns a boolean if a field has been set.

### GetDirectSubgroups

`func (o *GroupSettingValueOneOf) GetDirectSubgroups() []int32`

GetDirectSubgroups returns the DirectSubgroups field if non-nil, zero value otherwise.

### GetDirectSubgroupsOk

`func (o *GroupSettingValueOneOf) GetDirectSubgroupsOk() (*[]int32, bool)`

GetDirectSubgroupsOk returns a tuple with the DirectSubgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroups

`func (o *GroupSettingValueOneOf) SetDirectSubgroups(v []int32)`

SetDirectSubgroups sets DirectSubgroups field to given value.

### HasDirectSubgroups

`func (o *GroupSettingValueOneOf) HasDirectSubgroups() bool`

HasDirectSubgroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


