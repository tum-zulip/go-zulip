package models

// Principals - A list of user Ids (preferred) or Zulip API email addresses of the users to be subscribed to or unsubscribed from the channels specified in the `subscriptions` parameter. If not provided, then the requesting user/bot is subscribed.
//
// **Changes**: The integer format is new in Zulip 3.0 (feature level 9).
type Principals struct {
	UserIds    *[]int64
	UserEmails *[]string
}

// []UserIdsAsPrincipals is a convenience function that returns []int64 wrapped in Principals
func UserIdsAsPrincipals(v []int64) Principals {
	return Principals{
		UserIds: &v,
	}
}

// []UserEmailsAsPrincipals is a convenience function that returns []string wrapped in Principals
func UserEmailsAsPrincipals(v []string) Principals {
	return Principals{
		UserEmails: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Principals) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Principals) MarshalJSON() ([]byte, error) {
	return marshalUnionType(src)
}
