package users

import "github.com/tum-zulip/go-zulip/zulip"

// GetAlertWordsResponse struct for GetAlertWordsResponse
type AlertWordsResponse struct {
	zulip.Response

	// An array of strings, where each string is an alert word (or phrase) configured by the user.
	AlertWords []string `json:"alert_words,omitempty"`
}

// CreateUserResponse struct for CreateUserResponse
type CreateUserResponse struct {
	zulip.Response

	// The Id assigned to the newly created user.
	//
	// **Changes**: New in Zulip 4.0 (feature level 30).
	UserID int64 `json:"user_id,omitempty"`
}

// GetUserGroupsResponse struct for GetUserGroupsResponse
type GetUserGroupsResponse struct {
	zulip.Response

	// A list of `user_group` objects.
	UserGroups []zulip.UserGroup `json:"user_groups,omitempty"`
}

// GetUserPresenceResponse struct for GetUserPresenceResponse
type GetUserPresenceResponse struct {
	zulip.Response

	// An object containing the presence details for the user.
	Presence map[string]zulip.UserPresence `json:"presence,omitempty"`
}

// GetUserStatusResponse struct for GetUserStatusResponse
type GetUserStatusResponse struct {
	zulip.Response

	// The status set by the user. Note that, if the user doesn't have a status currently set, then the returned dictionary will be empty as none of the keys listed below will be present.
	Status zulip.UserStatus `json:"status,omitempty"`
}

// GetUsersResponse struct for GetUsersResponse
type GetUsersResponse struct {
	zulip.Response

	// A list of `user` objects, each containing details about a user in the organization.
	Members []zulip.User `json:"members,omitempty"`
}

// GetUserGroupSubgroupsResponse struct for GetUserGroupSubgroupsResponse
type GetUserGroupSubgroupsResponse struct {
	zulip.Response
	Subgroups []int64 `json:"subgroups,omitempty"`
}

type GetUserGroupMembersResponse struct {
	zulip.Response
	// A list containing the user Ids of members of the user group.
	Members []int64 `json:"members,omitempty"`
}

type GetUserResponse struct {
	zulip.Response
	User zulip.User `json:"user,omitempty"`
}

type GetIsUserGroupMemberResponse struct {
	zulip.Response
	IsUserGroupMember bool `json:"is_user_group_member,omitempty"`
}

// GetAttachmentsResponse struct for GetAttachmentsResponse
type GetAttachmentsResponse struct {
	zulip.Response
	// A list of `attachment` objects, each containing details about a file uploaded by the user.
	Attachments []zulip.Attachment `json:"attachments,omitempty"`
	// The total size of all files uploaded by users in the organization, in bytes.
	UploadSpaceUsed *int64 `json:"upload_space_used,omitempty"`
}

// CreateUserGroupResponse struct for CreateUserGroupResponse
type CreateUserGroupResponse struct {
	zulip.Response
	// The unique Id of the created user group.
	GroupId int64 `json:"group_id"`
}

// UpdatePresenceResponse struct for UpdatePresenceResponse
type UpdatePresenceResponse struct {
	zulip.Response
	// The identifier for the latest user presence data returned in the `presences` data of the response.  If a value was passed for `last_update_id`, then this is guaranteed to be equal to or greater than that value. If it is the same value, then that indicates to the client that there were no updates to previously received user presence data.  The client should then pass this value as the `last_update_id` parameter when it next queries this endpoint, in order to only receive new user presence data and avoid redundantly fetching already known information.  This will be `-1` if no value was passed for [`last_update_id`] and no user presence data was returned by the server. This can happen, for example, if an organization has disabled presence.
	//
	// **Changes**: New in Zulip 9.0 (feature level 263).
	//
	// [`last_update_id`]: https://zulip.com/api/update-presence#parameter-last_update_id
	PresenceLastUpdateId int64 `json:"presence_last_update_id,omitempty"`
	// Only present if `ping_only` is `false`.  The time when the server fetched the `presences` data included in the response.
	ServerTimestamp *float32 `json:"server_timestamp,omitempty"`
	// Only present if `ping_only` is `false`.  A dictionary where each entry describes the presence details of a user in the Zulip organization. Entries can be in either the modern presence format or the legacy presence format.  These entries will be the modern presence format when the `last_updated_id` parameter is passed, or when the deprecated `slim_presence` parameter is `true`.  If the deprecated `slim_presence` parameter is `false` and the `last_updated_id` parameter is omitted, the entries will be in the legacy presence API format.  **Note**: The legacy presence format should only be used when interacting with old servers. It will be removed as soon as doing so is practical.
	Presences map[string]zulip.PresenceUpdateValue `json:"presences,omitempty"`
	// Legacy property for a part of the Zephyr mirroring system.  **Deprecated**. Clients should ignore this field.
	// Deprecated
	ZephyrMirrorActive *bool `json:"zephyr_mirror_active,omitempty"`
}

// GetOwnUserResponse struct for GetOwnUserResponse
type GetOwnUserResponse struct {
	zulip.Response

	zulip.User

	// The integer Id of the last message received by the requesting user's account.  **Deprecated**. We plan to remove this in favor of recommending using `GET /messages` with `"anchor": "newest"`.
	// Deprecated
	MaxMessageId *int64 `json:"max_message_id,omitempty"`
}
