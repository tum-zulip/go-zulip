package events

import "github.com/tum-zulip/go-zulip/zulip"

// UserGroupAddEvent Event sent to users in an organization when a [user group] is created.
//
// [user group]: https://zulip.com/help/user-groups
type UserGroupAddEvent struct {
	event

	Group zulip.UserGroup `json:"group,omitempty"`
}

// UserGroupRemoveEvent Event sent when a user group is deactivated but only to clients with `include_deactivated_groups` client capability set to `false`.
//
// **Changes**: Prior to Zulip 10.0 (feature level 294), this event was sent when a user group was deleted.
type UserGroupRemoveEvent struct {
	event

	// The ID of the group which has been deleted.
	GroupID int64 `json:"group_id,omitempty"`
}

// UserGroupUpdateEvent Event sent to all users in a Zulip organization when a property of a user group is changed.  For group deactivation, this event is only sent if `include_deactivated_groups` client capability is set to `true`.  This event is also sent when deactivating or reactivating a user for settings set to anonymous user groups which the user is direct member of. When deactivating the user, event is only sent to users who cannot access the deactivated user.
//
// **Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when deactivating or reactivating a user.  Prior to Zulip 10.0 (feature level 294), this event was sent to all clients when a user group was deactivated.
type UserGroupUpdateEvent struct {
	event

	// The ID of the user group whose details have changed.
	GroupID int64               `json:"group_id,omitempty"`
	Data    UserGroupUpdateData `json:"data,omitempty"`
}

// UserGroupAddMembersEvent Event sent to all users when users have been added to a user group.  This event is also sent when reactivating a user for all the user groups the reactivated user was a member of before being deactivated.
//
// **Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when reactivating a user.
type UserGroupMembersEvent struct {
	event

	// The ID of the user group with added/removed members.
	GroupID int64 `json:"group_id,omitempty"`
	// Array containing the IDs of the users who have been added/removed to the user group.
	UserIDs []int64 `json:"user_ids,omitempty"`
}

// UserGroupAddSubgroupsEvent Event sent to all users when subgroups have been added to a user group.
//
// **Changes**: New in Zulip 6.0 (feature level 127).
type UserGroupSubgroupsEvent struct {
	event

	// The ID of the user group whose details have changed.
	GroupID int64 `json:"group_id,omitempty"`
	// Array containing the IDs of the subgroups that have been added/removed to the user group.
	//
	// **Changes**: New in Zulip 6.0 (feature level 131). Previously, this was called `subgroup_ids`, but clients can ignore older events as this feature level predates subgroups being fully implemented.
	DirectSubgroupIDs []int64 `json:"direct_subgroup_ids,omitempty"`
}

// UserGroupUpdate Dictionary containing the changed details of the user group.
type UserGroupUpdateData struct {
	// The new name of the user group. Only present if the group's name changed.
	Name *string `json:"name,omitempty"`
	// The new description of the group. Only present if the description changed.
	Description           *string                  `json:"description,omitempty"`
	CanAddMembersGroup    *zulip.GroupSettingValue `json:"can_add_members_group,omitempty"`
	CanJoinGroup          *zulip.GroupSettingValue `json:"can_join_group,omitempty"`
	CanLeaveGroup         *zulip.GroupSettingValue `json:"can_leave_group,omitempty"`
	CanManageGroup        *zulip.GroupSettingValue `json:"can_manage_group,omitempty"`
	CanMentionGroup       *zulip.GroupSettingValue `json:"can_mention_group,omitempty"`
	CanRemoveMembersGroup *zulip.GroupSettingValue `json:"can_remove_members_group,omitempty"`
	// Whether the user group is deactivated. Deactivated groups cannot be used as a subgroup of another group or used for any other purpose.
	//
	// **Changes**: New in Zulip 10.0 (feature level 290).
	Deactivated *bool `json:"deactivated,omitempty"`
}
