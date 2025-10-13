# GetEvents200ResponseAllOfEventsInnerOneOf44Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the user group. Only present if the group&#39;s name changed.  | [optional] 
**Description** | Pointer to **string** | The new description of the group. Only present if the description changed.  | [optional] 
**CanAddMembersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to add members to this group. Only present if this user group permission setting changed.  **Changes**: New in Zulip 10.0 (feature level 305). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanJoinGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to join this group. Only present if this user group permission setting changed.  **Changes**: New in Zulip 10.0 (feature level 301).  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanLeaveGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to leave this group. Only present if this user group permission setting changed.  **Changes**: New in Zulip 10.0 (feature level 308).  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**CanManageGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [manage this group][manage-user-groups]. Only present if this user group permission setting changed.  **Changes**: New in Zulip 10.0 (feature level 283).  Will be one of the following:  [setting-values]: /api/group-setting-values [manage-user-groups]: /help/manage-user-groups  | [optional] 
**CanMentionGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to [mention this user group][mentions]. Only present if this user group permission setting changed.  **Changes**: Before Zulip 9.0 (feature level 258), this setting was always the integer form of a [group-setting value][setting-values].  Before Zulip 8.0 (feature level 198), this setting was named &#x60;can_mention_group_id&#x60;.  New in Zulip 8.0 (feature level 191). Previously, groups could be mentioned only if they were not [system groups][system-groups].  Will be one of the following:  [setting-values]: /api/group-setting-values [system-groups]: /api/group-setting-values#system-groups [mentions]: /help/mention-a-user-or-group  | [optional] 
**CanRemoveMembersGroup** | Pointer to [**GroupSettingValue**](GroupSettingValue.md) | A [group-setting value][setting-values] defining the set of users who have permission to remove members from this group. Only present if this user group permission setting changed.  **Changes**: New in Zulip 10.0 (feature level 324). Previously, this permission was controlled by the &#x60;can_manage_group&#x60; setting.  Will be one of the following:  [setting-values]: /api/group-setting-values  | [optional] 
**Deactivated** | Pointer to **bool** | Whether the user group is deactivated. Deactivated groups cannot be used as a subgroup of another group or used for any other purpose.  **Changes**: New in Zulip 10.0 (feature level 290).  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf44Data

`func NewGetEvents200ResponseAllOfEventsInnerOneOf44Data() *GetEvents200ResponseAllOfEventsInnerOneOf44Data`

NewGetEvents200ResponseAllOfEventsInnerOneOf44Data instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf44Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf44DataWithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf44DataWithDefaults() *GetEvents200ResponseAllOfEventsInnerOneOf44Data`

NewGetEvents200ResponseAllOfEventsInnerOneOf44DataWithDefaults instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf44Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCanAddMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanAddMembersGroup() GroupSettingValue`

GetCanAddMembersGroup returns the CanAddMembersGroup field if non-nil, zero value otherwise.

### GetCanAddMembersGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanAddMembersGroupOk() (*GroupSettingValue, bool)`

GetCanAddMembersGroupOk returns a tuple with the CanAddMembersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAddMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanAddMembersGroup(v GroupSettingValue)`

SetCanAddMembersGroup sets CanAddMembersGroup field to given value.

### HasCanAddMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanAddMembersGroup() bool`

HasCanAddMembersGroup returns a boolean if a field has been set.

### GetCanJoinGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanJoinGroup() GroupSettingValue`

GetCanJoinGroup returns the CanJoinGroup field if non-nil, zero value otherwise.

### GetCanJoinGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanJoinGroupOk() (*GroupSettingValue, bool)`

GetCanJoinGroupOk returns a tuple with the CanJoinGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanJoinGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanJoinGroup(v GroupSettingValue)`

SetCanJoinGroup sets CanJoinGroup field to given value.

### HasCanJoinGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanJoinGroup() bool`

HasCanJoinGroup returns a boolean if a field has been set.

### GetCanLeaveGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanLeaveGroup() GroupSettingValue`

GetCanLeaveGroup returns the CanLeaveGroup field if non-nil, zero value otherwise.

### GetCanLeaveGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanLeaveGroupOk() (*GroupSettingValue, bool)`

GetCanLeaveGroupOk returns a tuple with the CanLeaveGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLeaveGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanLeaveGroup(v GroupSettingValue)`

SetCanLeaveGroup sets CanLeaveGroup field to given value.

### HasCanLeaveGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanLeaveGroup() bool`

HasCanLeaveGroup returns a boolean if a field has been set.

### GetCanManageGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanManageGroup() GroupSettingValue`

GetCanManageGroup returns the CanManageGroup field if non-nil, zero value otherwise.

### GetCanManageGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanManageGroupOk() (*GroupSettingValue, bool)`

GetCanManageGroupOk returns a tuple with the CanManageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanManageGroup(v GroupSettingValue)`

SetCanManageGroup sets CanManageGroup field to given value.

### HasCanManageGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanManageGroup() bool`

HasCanManageGroup returns a boolean if a field has been set.

### GetCanMentionGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanMentionGroup() GroupSettingValue`

GetCanMentionGroup returns the CanMentionGroup field if non-nil, zero value otherwise.

### GetCanMentionGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanMentionGroupOk() (*GroupSettingValue, bool)`

GetCanMentionGroupOk returns a tuple with the CanMentionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMentionGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanMentionGroup(v GroupSettingValue)`

SetCanMentionGroup sets CanMentionGroup field to given value.

### HasCanMentionGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanMentionGroup() bool`

HasCanMentionGroup returns a boolean if a field has been set.

### GetCanRemoveMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanRemoveMembersGroup() GroupSettingValue`

GetCanRemoveMembersGroup returns the CanRemoveMembersGroup field if non-nil, zero value otherwise.

### GetCanRemoveMembersGroupOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetCanRemoveMembersGroupOk() (*GroupSettingValue, bool)`

GetCanRemoveMembersGroupOk returns a tuple with the CanRemoveMembersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemoveMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetCanRemoveMembersGroup(v GroupSettingValue)`

SetCanRemoveMembersGroup sets CanRemoveMembersGroup field to given value.

### HasCanRemoveMembersGroup

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasCanRemoveMembersGroup() bool`

HasCanRemoveMembersGroup returns a boolean if a field has been set.

### GetDeactivated

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.

### HasDeactivated

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf44Data) HasDeactivated() bool`

HasDeactivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


