package models

// Recipients - A message's tentative target audience.
// For channel messages, the integer Id of the channel.
// For direct messages, a list containing integer user Ids.
type Recipient struct {
	Users   []int64
	Channel *int64
}

func (o Recipient) asArray() []int64 {
	if o.Users != nil {
		return o.Users
	} else if o.Channel != nil {
		return []int64{*o.Channel}
	}
	return nil
}

// UsersAsRecipient is a convenience function that returns a list of UserIds wrapped in Recipient
func UsersAsRecipient(v []int64) Recipient {
	return Recipient{
		Users: v,
	}
}

// UsersEmailsAsRecipient is a convenience function that returns a list of UserEmails wrapped in Recipient
func UsersEmailsAsRecipient(v []string) Recipient {
	return Recipient{
		Users:   nil,
		Channel: nil,
	}
}

// UserAsRecipient is a convenience function that returns UserId wrapped in Recipient
func UserAsRecipient(v int64) Recipient {
	return Recipient{
		Users: []int64{v},
	}
}

// ChannelAsRecipient is a convenience function that returns ChannelId wrapped in Recipient
func ChannelAsRecipient(v int64) Recipient {
	return Recipient{
		Channel: &v,
	}
}

func (r Recipient) Equals(other Recipient) bool {
	if (r.Channel == nil) != (other.Channel == nil) {
		return false
	}
	if (r.Users == nil) != (other.Users == nil) {
		return false
	}

	if r.Channel != nil {
		if *r.Channel != *other.Channel {
			return false
		}
	}
	if r.Users != nil {
		if len(r.Users) != len(other.Users) {
			return false
		}
		for i, v := range r.Users {
			if v != other.Users[i] {
				return false
			}
		}
		return true
	}
	return true
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Recipient) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Recipient) MarshalJSON() ([]byte, error) {
	return marshalUnionType(src)
}
