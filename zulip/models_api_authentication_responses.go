package zulip

type ApiKeyResponse struct {
	Response
	// The API key that can be used to authenticate as the requested user.
	ApiKey string `json:"api_key"`
	// The email address of the user who owns the API key.
	Email string `json:"email"`
	// The unique Id of the user who owns the API key.
	//
	// **Changes**: New in Zulip 7.0 (feature level 171).
	UserId int64 `json:"user_id,omitempty"`
}
