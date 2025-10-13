# GetUserGroups200ResponseAllOfUserGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The human-readable description of the user group.  | [optional] 
**Id** | Pointer to **int32** | The user group&#39;s integer ID.  | [optional] 
**DateCreated** | Pointer to **NullableInt32** | The UNIX timestamp for when the user group was created, in UTC seconds.  A &#x60;null&#x60; value means the user group has no recorded date, which is often because the group predates the metadata being tracked starting in Zulip 8.0, or because it was created via a data import tool or [management command][management-commands].  **Changes**: New in Zulip 10.0 (feature level 292).  [management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html  | [optional] 
**CreatorId** | Pointer to **NullableInt32** | The ID of the user who created this user group.  A &#x60;null&#x60; value means the user group has no recorded creator, which is often because the group predates the metadata being tracked starting in Zulip 8.0, or because it was created via a data import tool or [management command][management-commands].  **Changes**: New in Zulip 10.0 (feature level 292).  [management-commands]: https://zulip.readthedocs.io/en/latest/production/management-commands.html  | [optional] 
**Members** | Pointer to **[]int32** | The integer user IDs of the user group&#39;s members, which are guaranteed to be non-deactivated users in the organization.  **Changes**: Prior to Zulip 10.0 (feature level 303), this list also included deactivated users who were members of the user group before being deactivated.  | [optional] 
**DirectSubgroupIds** | Pointer to **[]int32** | The integer user group IDs of the direct subgroups.  **Changes**: New in Zulip 6.0 (feature level 131). Introduced in feature level 127 as &#x60;subgroups&#x60;, but clients can ignore older events as this feature level predates subgroups being fully implemented.  | [optional] 
**Name** | Pointer to **string** | User group name.  | [optional] 
**IsSystemGroup** | Pointer to **bool** | Whether the user group is a system group which cannot be modified by users.  **Changes**: New in Zulip 5.0 (feature level 93).  | [optional] 
**CanAddMembersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to add members to this user group.  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanJoinGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to join this user group.  **Changes**: New in Zulip 10.0 (feature level 301).  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanLeaveGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to leave this user group.  **Changes**: New in Zulip 10.0 (feature level 308).  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanManageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [manage this user group][manage-user-groups].  **Changes**: New in Zulip 10.0 (feature level 283).  Will be one of the following:  [setting-values]: /api/group-setting-values [manage-user-groups]: /help/manage-user-groups  | [optional] 
**CanMentionGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [mention this user group][mentions].  **Changes**: Before Zulip 9.0 (feature level 258), this setting was always the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this setting was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  Will be one of the following:  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [mentions]: /help/mention-a-user-or-group  | [optional] 
**CanRemoveMembersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to remove members from this user group.  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**Deactivated** | Pointer to **bool** | Whether the user group is deactivated. Deactivated groups cannot be used as a subgroup of another group or used for any other purpose.  **Changes**: New in Zulip 10.0 (feature level 290).  | [optional] 

## Methods

### NewGetUserGroups200ResponseAllOfUserGroupsInner

`func NewGetUserGroups200ResponseAllOfUserGroupsInner() *GetUserGroups200ResponseAllOfUserGroupsInner`

NewGetUserGroups200ResponseAllOfUserGroupsInner instantiates a new GetUserGroups200ResponseAllOfUserGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGroups200ResponseAllOfUserGroupsInnerWithDefaults

`func NewGetUserGroups200ResponseAllOfUserGroupsInnerWithDefaults() *GetUserGroups200ResponseAllOfUserGroupsInner`

NewGetUserGroups200ResponseAllOfUserGroupsInnerWithDefaults instantiates a new GetUserGroups200ResponseAllOfUserGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDateCreated() int32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDateCreatedOk() (*int32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetDateCreated(v int32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCreatorId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetMembers

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetMembersOk() (*[]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetMembers(v []int32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetDirectSubgroupIds

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDirectSubgroupIds() []int32`

GetDirectSubgroupIds returns the DirectSubgroupIds field if non-nil, zero value otherwise.

### GetDirectSubgroupIdsOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDirectSubgroupIdsOk() (*[]int32, bool)`

GetDirectSubgroupIdsOk returns a tuple with the DirectSubgroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubgroupIds

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetDirectSubgroupIds(v []int32)`

SetDirectSubgroupIds sets DirectSubgroupIds field to given value.

### HasDirectSubgroupIds

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasDirectSubgroupIds() bool`

HasDirectSubgroupIds returns a boolean if a field has been set.

### GetName

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsSystemGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetIsSystemGroup() bool`

GetIsSystemGroup returns the IsSystemGroup field if non-nil, zero value otherwise.

### GetIsSystemGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetIsSystemGroupOk() (*bool, bool)`

GetIsSystemGroupOk returns a tuple with the IsSystemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetIsSystemGroup(v bool)`

SetIsSystemGroup sets IsSystemGroup field to given value.

### HasIsSystemGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasIsSystemGroup() bool`

HasIsSystemGroup returns a boolean if a field has been set.

### GetCanAddMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanAddMembersGroup() GroupSettingValue`

GetCanAddMembersGroup returns the CanAddMembersGroup field if non-nil, zero value otherwise.

### GetCanAddMembersGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanAddMembersGroupOk() (*GroupSettingValue, bool)`

GetCanAddMembersGroupOk returns a tuple with the CanAddMembersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanAddMembersGroup(v GroupSettingValue)`

SetCanAddMembersGroup sets CanAddMembersGroup field to given value.

### HasCanAddMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanAddMembersGroup() bool`

HasCanAddMembersGroup returns a boolean if a field has been set.

### GetCanJoinGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanJoinGroup() GroupSettingValue`

GetCanJoinGroup returns the CanJoinGroup field if non-nil, zero value otherwise.

### GetCanJoinGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanJoinGroupOk() (*GroupSettingValue, bool)`

GetCanJoinGroupOk returns a tuple with the CanJoinGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanJoinGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanJoinGroup(v GroupSettingValue)`

SetCanJoinGroup sets CanJoinGroup field to given value.

### HasCanJoinGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanJoinGroup() bool`

HasCanJoinGroup returns a boolean if a field has been set.

### GetCanLeaveGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanLeaveGroup() GroupSettingValue`

GetCanLeaveGroup returns the CanLeaveGroup field if non-nil, zero value otherwise.

### GetCanLeaveGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanLeaveGroupOk() (*GroupSettingValue, bool)`

GetCanLeaveGroupOk returns a tuple with the CanLeaveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLeaveGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanLeaveGroup(v GroupSettingValue)`

SetCanLeaveGroup sets CanLeaveGroup field to given value.

### HasCanLeaveGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanLeaveGroup() bool`

HasCanLeaveGroup returns a boolean if a field has been set.

### GetCanManageGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanManageGroup() GroupSettingValue`

GetCanManageGroup returns the CanManageGroup field if non-nil, zero value otherwise.

### GetCanManageGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanManageGroupOk() (*GroupSettingValue, bool)`

GetCanManageGroupOk returns a tuple with the CanManageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanManageGroup(v GroupSettingValue)`

SetCanManageGroup sets CanManageGroup field to given value.

### HasCanManageGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanManageGroup() bool`

HasCanManageGroup returns a boolean if a field has been set.

### GetCanMentionGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanMentionGroup() GroupSettingValue`

GetCanMentionGroup returns the CanMentionGroup field if non-nil, zero value otherwise.

### GetCanMentionGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanMentionGroupOk() (*GroupSettingValue, bool)`

GetCanMentionGroupOk returns a tuple with the CanMentionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMentionGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanMentionGroup(v GroupSettingValue)`

SetCanMentionGroup sets CanMentionGroup field to given value.

### HasCanMentionGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanMentionGroup() bool`

HasCanMentionGroup returns a boolean if a field has been set.

### GetCanRemoveMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanRemoveMembersGroup() GroupSettingValue`

GetCanRemoveMembersGroup returns the CanRemoveMembersGroup field if non-nil, zero value otherwise.

### GetCanRemoveMembersGroupOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetCanRemoveMembersGroupOk() (*GroupSettingValue, bool)`

GetCanRemoveMembersGroupOk returns a tuple with the CanRemoveMembersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetCanRemoveMembersGroup(v GroupSettingValue)`

SetCanRemoveMembersGroup sets CanRemoveMembersGroup field to given value.

### HasCanRemoveMembersGroup

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasCanRemoveMembersGroup() bool`

HasCanRemoveMembersGroup returns a boolean if a field has been set.

### GetDeactivated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *GetUserGroups200ResponseAllOfUserGroupsInner) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


