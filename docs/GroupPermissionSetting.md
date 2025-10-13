# GroupPermissionSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireSystemGroup** | Pointer to **bool** | Whether the setting can only be set to a system user group.  | [optional] 
**AllowInternetGroup** | Pointer to **bool** | Whether the setting can be set to &#x60;role:internet&#x60; system group.  | [optional] 
**AllowNobodyGroup** | Pointer to **bool** | Whether the setting can be set to &#x60;role:nobody&#x60; system group.  | [optional] 
**AllowEveryoneGroup** | Pointer to **bool** | Whether the setting can be set to &#x60;role:everyone&#x60; system group.  If false, guest users cannot exercise this permission even if they are part of the [group-setting value](zulip.com/api/group-setting-values for this setting.  | [optional] 
**DefaultGroupName** | Pointer to **string** | Name of the default group for the setting.  | [optional] 
**DefaultForSystemGroups** | Pointer to **NullableString** | Name of the default group for the setting for system groups.  This is non-null only for group-level settings.  | [optional] 
**AllowedSystemGroups** | Pointer to **[]string** | An array of names of system groups to which the setting can be set to.  If the list is empty, the setting can be set to system groups based on the other boolean fields.  **Changes**: New in Zulip 8.0 (feature level 225).  | [optional] 

## Methods

### NewGroupPermissionSetting

`func NewGroupPermissionSetting() *GroupPermissionSetting`

NewGroupPermissionSetting instantiates a new GroupPermissionSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionSettingWithDefaults

`func NewGroupPermissionSettingWithDefaults() *GroupPermissionSetting`

NewGroupPermissionSettingWithDefaults instantiates a new GroupPermissionSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireSystemGroup

`func (o *GroupPermissionSetting) GetRequireSystemGroup() bool`

GetRequireSystemGroup returns the RequireSystemGroup field if non-nil, zero value otherwise.

### GetRequireSystemGroupOk

`func (o *GroupPermissionSetting) GetRequireSystemGroupOk() (*bool, bool)`

GetRequireSystemGroupOk returns a tuple with the RequireSystemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSystemGroup

`func (o *GroupPermissionSetting) SetRequireSystemGroup(v bool)`

SetRequireSystemGroup sets RequireSystemGroup field to given value.

### HasRequireSystemGroup

`func (o *GroupPermissionSetting) HasRequireSystemGroup() bool`

HasRequireSystemGroup returns a boolean if a field has been set.

### GetAllowInternetGroup

`func (o *GroupPermissionSetting) GetAllowInternetGroup() bool`

GetAllowInternetGroup returns the AllowInternetGroup field if non-nil, zero value otherwise.

### GetAllowInternetGroupOk

`func (o *GroupPermissionSetting) GetAllowInternetGroupOk() (*bool, bool)`

GetAllowInternetGroupOk returns a tuple with the AllowInternetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInternetGroup

`func (o *GroupPermissionSetting) SetAllowInternetGroup(v bool)`

SetAllowInternetGroup sets AllowInternetGroup field to given value.

### HasAllowInternetGroup

`func (o *GroupPermissionSetting) HasAllowInternetGroup() bool`

HasAllowInternetGroup returns a boolean if a field has been set.

### GetAllowNobodyGroup

`func (o *GroupPermissionSetting) GetAllowNobodyGroup() bool`

GetAllowNobodyGroup returns the AllowNobodyGroup field if non-nil, zero value otherwise.

### GetAllowNobodyGroupOk

`func (o *GroupPermissionSetting) GetAllowNobodyGroupOk() (*bool, bool)`

GetAllowNobodyGroupOk returns a tuple with the AllowNobodyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNobodyGroup

`func (o *GroupPermissionSetting) SetAllowNobodyGroup(v bool)`

SetAllowNobodyGroup sets AllowNobodyGroup field to given value.

### HasAllowNobodyGroup

`func (o *GroupPermissionSetting) HasAllowNobodyGroup() bool`

HasAllowNobodyGroup returns a boolean if a field has been set.

### GetAllowEveryoneGroup

`func (o *GroupPermissionSetting) GetAllowEveryoneGroup() bool`

GetAllowEveryoneGroup returns the AllowEveryoneGroup field if non-nil, zero value otherwise.

### GetAllowEveryoneGroupOk

`func (o *GroupPermissionSetting) GetAllowEveryoneGroupOk() (*bool, bool)`

GetAllowEveryoneGroupOk returns a tuple with the AllowEveryoneGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEveryoneGroup

`func (o *GroupPermissionSetting) SetAllowEveryoneGroup(v bool)`

SetAllowEveryoneGroup sets AllowEveryoneGroup field to given value.

### HasAllowEveryoneGroup

`func (o *GroupPermissionSetting) HasAllowEveryoneGroup() bool`

HasAllowEveryoneGroup returns a boolean if a field has been set.

### GetDefaultGroupName

`func (o *GroupPermissionSetting) GetDefaultGroupName() string`

GetDefaultGroupName returns the DefaultGroupName field if non-nil, zero value otherwise.

### GetDefaultGroupNameOk

`func (o *GroupPermissionSetting) GetDefaultGroupNameOk() (*string, bool)`

GetDefaultGroupNameOk returns a tuple with the DefaultGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupName

`func (o *GroupPermissionSetting) SetDefaultGroupName(v string)`

SetDefaultGroupName sets DefaultGroupName field to given value.

### HasDefaultGroupName

`func (o *GroupPermissionSetting) HasDefaultGroupName() bool`

HasDefaultGroupName returns a boolean if a field has been set.

### GetDefaultForSystemGroups

`func (o *GroupPermissionSetting) GetDefaultForSystemGroups() string`

GetDefaultForSystemGroups returns the DefaultForSystemGroups field if non-nil, zero value otherwise.

### GetDefaultForSystemGroupsOk

`func (o *GroupPermissionSetting) GetDefaultForSystemGroupsOk() (*string, bool)`

GetDefaultForSystemGroupsOk returns a tuple with the DefaultForSystemGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForSystemGroups

`func (o *GroupPermissionSetting) SetDefaultForSystemGroups(v string)`

SetDefaultForSystemGroups sets DefaultForSystemGroups field to given value.

### HasDefaultForSystemGroups

`func (o *GroupPermissionSetting) HasDefaultForSystemGroups() bool`

HasDefaultForSystemGroups returns a boolean if a field has been set.

### SetDefaultForSystemGroupsNil

`func (o *GroupPermissionSetting) SetDefaultForSystemGroupsNil(b bool)`

 SetDefaultForSystemGroupsNil sets the value for DefaultForSystemGroups to be an explicit nil

### UnsetDefaultForSystemGroups
`func (o *GroupPermissionSetting) UnsetDefaultForSystemGroups()`

UnsetDefaultForSystemGroups ensures that no value is present for DefaultForSystemGroups, not even an explicit nil
### GetAllowedSystemGroups

`func (o *GroupPermissionSetting) GetAllowedSystemGroups() []string`

GetAllowedSystemGroups returns the AllowedSystemGroups field if non-nil, zero value otherwise.

### GetAllowedSystemGroupsOk

`func (o *GroupPermissionSetting) GetAllowedSystemGroupsOk() (*[]string, bool)`

GetAllowedSystemGroupsOk returns a tuple with the AllowedSystemGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSystemGroups

`func (o *GroupPermissionSetting) SetAllowedSystemGroups(v []string)`

SetAllowedSystemGroups sets AllowedSystemGroups field to given value.

### HasAllowedSystemGroups

`func (o *GroupPermissionSetting) HasAllowedSystemGroups() bool`

HasAllowedSystemGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


