package authentication

import (
	"github.com/tum-zulip/go-zulip/zulip"
)

type ApiKeyResponse struct {
	zulip.Response

	// The API key that can be used to authenticate as the requested user.
	ApiKey string `json:"api_key"`
	// The email address of the user who owns the API key.
	Email string `json:"email"`
	// The unique ID of the user who owns the API key.
	//
	// **Changes**: New in Zulip 7.0 (feature level 171).
	UserID int64 `json:"user_id,omitempty"`
}
