package zulip

import (
	"encoding/json"
	"time"

	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

// UserGroup struct for UserGroup
type UserGroup struct {
	// The human-readable description of the user group.
	Description string `json:"description,omitempty"`
	// The user group's integer Id.
	Id int64 `json:"id,omitempty"`
	// The UNIX timestamp for when the user group was created, in UTC seconds.  A `null` value means the user group has no recorded date, which is often because the group predates the metadata being tracked starting in Zulip 8.0, or because it was created via a data import tool or [management command].
	//
	// **Changes**: New in Zulip 10.0 (feature level 292).
	//
	// [management command]: https://zulip.readthedocs.io/en/latest/production/management-commands.html
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The Id of the user who created this user group.  A `null` value means the user group has no recorded creator, which is often because the group predates the metadata being tracked starting in Zulip 8.0, or because it was created via a data import tool or [management command].
	//
	// **Changes**: New in Zulip 10.0 (feature level 292).
	//
	// [management command]: https://zulip.readthedocs.io/en/latest/production/management-commands.html
	CreatorId *int64 `json:"creator_id,omitempty"`
	// The integer user Ids of the user group's members, which are guaranteed to be non-deactivated users in the organization.
	//
	// **Changes**: Prior to Zulip 10.0 (feature level 303), this list also included deactivated users who were members of the user group before being deactivated.
	Members []int64 `json:"members,omitempty"`
	// The integer user group Ids of the direct subgroups.
	//
	// **Changes**: New in Zulip 6.0 (feature level 131). Introduced in feature level 127 as `subgroups`, but clients can ignore older events as this feature level predates subgroups being fully implemented.
	DirectSubgroupIds []int64 `json:"direct_subgroup_ids,omitempty"`
	// User group name.
	Name string `json:"name,omitempty"`
	// Whether the user group is a system group which cannot be modified by users.
	//
	// **Changes**: New in Zulip 5.0 (feature level 93).
	IsSystemGroup         bool              `json:"is_system_group,omitempty"`
	CanAddMembersGroup    GroupSettingValue `json:"can_add_members_group,omitempty"`
	CanJoinGroup          GroupSettingValue `json:"can_join_group,omitempty"`
	CanLeaveGroup         GroupSettingValue `json:"can_leave_group,omitempty"`
	CanManageGroup        GroupSettingValue `json:"can_manage_group,omitempty"`
	CanMentionGroup       GroupSettingValue `json:"can_mention_group,omitempty"`
	CanRemoveMembersGroup GroupSettingValue `json:"can_remove_members_group,omitempty"`
	// Whether the user group is deactivated. Deactivated groups cannot be used as a subgroup of another group or used for any other purpose.
	//
	// **Changes**: New in Zulip 10.0 (feature level 290).
	Deactivated bool `json:"deactivated,omitempty"`
}

type usergroupJSON struct {
	Description           string            `json:"description,omitempty"`
	Id                    int64             `json:"id,omitempty"`
	DateCreated           *int64            `json:"date_created,omitempty"`
	CreatorId             *int64            `json:"creator_id,omitempty"`
	Members               []int64           `json:"members,omitempty"`
	DirectSubgroupIds     []int64           `json:"direct_subgroup_ids,omitempty"`
	Name                  string            `json:"name,omitempty"`
	IsSystemGroup         bool              `json:"is_system_group,omitempty"`
	CanAddMembersGroup    GroupSettingValue `json:"can_add_members_group,omitempty"`
	CanJoinGroup          GroupSettingValue `json:"can_join_group,omitempty"`
	CanLeaveGroup         GroupSettingValue `json:"can_leave_group,omitempty"`
	CanManageGroup        GroupSettingValue `json:"can_manage_group,omitempty"`
	CanMentionGroup       GroupSettingValue `json:"can_mention_group,omitempty"`
	CanRemoveMembersGroup GroupSettingValue `json:"can_remove_members_group,omitempty"`
	Deactivated           bool              `json:"deactivated,omitempty"`
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	var aux usergroupJSON
	aux.Description = o.Description
	aux.Id = o.Id
	if o.DateCreated != nil {
		t := o.DateCreated.Unix()
		aux.DateCreated = &t
	}
	aux.CreatorId = o.CreatorId
	aux.Members = o.Members
	aux.DirectSubgroupIds = o.DirectSubgroupIds
	aux.Name = o.Name
	aux.IsSystemGroup = o.IsSystemGroup
	aux.CanAddMembersGroup = o.CanAddMembersGroup
	aux.CanJoinGroup = o.CanJoinGroup
	aux.CanLeaveGroup = o.CanLeaveGroup
	aux.CanManageGroup = o.CanManageGroup
	aux.CanMentionGroup = o.CanMentionGroup
	aux.CanRemoveMembersGroup = o.CanRemoveMembersGroup
	aux.Deactivated = o.Deactivated

	return json.Marshal(&aux)
}

func (o *UserGroup) UnmarshalJSON(data []byte) error {
	var aux usergroupJSON
	dec := utils.NewStrictDecoder(data)
	if err := dec.Decode(&aux); err != nil {
		return err
	}

	o.Description = aux.Description
	o.Id = aux.Id
	if aux.DateCreated != nil {
		t := time.Unix(*aux.DateCreated, 0).UTC()
		o.DateCreated = &t
	} else {
		o.DateCreated = nil
	}
	o.CreatorId = aux.CreatorId
	o.Members = aux.Members
	o.DirectSubgroupIds = aux.DirectSubgroupIds
	o.Name = aux.Name
	o.IsSystemGroup = aux.IsSystemGroup
	o.CanAddMembersGroup = aux.CanAddMembersGroup
	o.CanJoinGroup = aux.CanJoinGroup
	o.CanLeaveGroup = aux.CanLeaveGroup
	o.CanManageGroup = aux.CanManageGroup
	o.CanMentionGroup = aux.CanMentionGroup
	o.CanRemoveMembersGroup = aux.CanRemoveMembersGroup
	o.Deactivated = aux.Deactivated

	return nil
}
