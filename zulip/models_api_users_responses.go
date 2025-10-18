package zulip

// GetUserGroupSubgroupsResponse struct for GetUserGroupSubgroupsResponse
type GetUserGroupSubgroupsResponse struct {
	Response
	Subgroups []int64 `json:"subgroups,omitempty"`
}

type GetUserGroupMembersResponse struct {
	Response
	// A list containing the user Ids of members of the user group.
	Members []int64 `json:"members,omitempty"`
}

type GetUserResponse struct {
	Response
	User User `json:"user,omitempty"`
}

type GetIsUserGroupMemberResponse struct {
	Response
	IsUserGroupMember bool `json:"is_user_group_member,omitempty"`
}

// GetAttachmentsResponse struct for GetAttachmentsResponse
type GetAttachmentsResponse struct {
	Response
	// A list of `attachment` objects, each containing details about a file uploaded by the user.
	Attachments []Attachment `json:"attachments,omitempty"`
	// The total size of all files uploaded by users in the organization, in bytes.
	UploadSpaceUsed *int64 `json:"upload_space_used,omitempty"`
}

// CreateUserGroupResponse struct for CreateUserGroupResponse
type CreateUserGroupResponse struct {
	Response
	// The unique Id of the created user group.
	GroupId int64 `json:"group_id"`
}

// UpdatePresenceResponse struct for UpdatePresenceResponse
type UpdatePresenceResponse struct {
	Response
	// The identifier for the latest user presence data returned in the `presences` data of the response.  If a value was passed for `last_update_id`, then this is guaranteed to be equal to or greater than that value. If it is the same value, then that indicates to the client that there were no updates to previously received user presence data.  The client should then pass this value as the `last_update_id` parameter when it next queries this endpoint, in order to only receive new user presence data and avoid redundantly fetching already known information.  This will be `-1` if no value was passed for [`last_update_id`](#parameter-last_update_id) and no user presence data was returned by the server. This can happen, for example, if an organization has disabled presence.  **Changes**: New in Zulip 9.0 (feature level 263).
	PresenceLastUpdateId int64 `json:"presence_last_update_id,omitempty"`
	// Only present if `ping_only` is `false`.  The time when the server fetched the `presences` data included in the response.
	ServerTimestamp *float32 `json:"server_timestamp,omitempty"`
	// Only present if `ping_only` is `false`.  A dictionary where each entry describes the presence details of a user in the Zulip organization. Entries can be in either the modern presence format or the legacy presence format.  These entries will be the modern presence format when the `last_updated_id` parameter is passed, or when the deprecated `slim_presence` parameter is `true`.  If the deprecated `slim_presence` parameter is `false` and the `last_updated_id` parameter is omitted, the entries will be in the legacy presence API format.  **Note**: The legacy presence format should only be used when interacting with old servers. It will be removed as soon as doing so is practical.
	Presences map[string]PresenceUpdateValue `json:"presences,omitempty"`
	// Legacy property for a part of the Zephyr mirroring system.  **Deprecated**. Clients should ignore this field.
	// Deprecated
	ZephyrMirrorActive *bool `json:"zephyr_mirror_active,omitempty"`
}

// GetOwnUserResponse struct for GetOwnUserResponse
type GetOwnUserResponse struct {
	Response

	User

	// The integer Id of the last message received by the requesting user's account.  **Deprecated**. We plan to remove this in favor of recommending using `GET /messages` with `\"anchor\": \"newest\"`.
	// Deprecated
	MaxMessageId *int64 `json:"max_message_id,omitempty"`
}
