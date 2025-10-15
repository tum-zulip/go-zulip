package api

// GetAlertWordsResponse struct for GetAlertWordsResponse
type AlertWordsResponse struct {
	Response

	// An array of strings, where each string is an alert word (or phrase) configured by the user.
	AlertWords []string `json:"alert_words,omitempty"`
}

// CreateUserResponse struct for CreateUserResponse
type CreateUserResponse struct {
	Response

	// The Id assigned to the newly created user.  **Changes**: New in Zulip 4.0 (feature level 30).
	UserId int64 `json:"user_id,omitempty"`
}

// GetUserGroupsResponse struct for GetUserGroupsResponse
type GetUserGroupsResponse struct {
	Response

	// A list of `user_group` objects.
	UserGroups []UserGroup `json:"user_groups,omitempty"`
}

// GetUserPresenceResponse struct for GetUserPresenceResponse
type GetUserPresenceResponse struct {
	Response

	// An object containing the presence details for the user.
	Presence map[string]UserPresence `json:"presence,omitempty"`
}

// GetUserStatusResponse struct for GetUserStatusResponse
type GetUserStatusResponse struct {
	Response

	// The status set by the user. Note that, if the user doesn't have a status currently set, then the returned dictionary will be empty as none of the keys listed below will be present.
	Status UserStatus `json:"status,omitempty"`
}

// GetUsersResponse struct for GetUsersResponse
type GetUsersResponse struct {
	Response

	// A list of `user` objects, each containing details about a user in the organization.
	Members []User `json:"members,omitempty"`
}
