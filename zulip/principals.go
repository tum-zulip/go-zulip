package zulip

import (
	"github.com/tum-zulip/go-zulip/zulip/internal/union"
)

// Principals - A list of user IDs (preferred) or Zulip API email addresses of the users to be subscribed to or unsubscribed from the channels specified in the `subscriptions` parameter. If not provided, then the requesting user/bot is subscribed.
//
// **Changes**: The integer format is new in Zulip 3.0 (feature level 9).
type Principals struct {
	UserIDs    *[]int64
	UserEmails *[]string
}

// []UserIDsAsPrincipals is a convenience function that returns []int64 wrapped in Principals.
func UserIDsAsPrincipals(v ...int64) Principals {
	return Principals{
		UserIDs: &v,
	}
}

// []UserEmailsAsPrincipals is a convenience function that returns []string wrapped in Principals.
func UserEmailsAsPrincipals(v ...string) Principals {
	return Principals{
		UserEmails: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (p *Principals) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, p)
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (p Principals) MarshalJSON() ([]byte, error) {
	return union.Marshal(p)
}
