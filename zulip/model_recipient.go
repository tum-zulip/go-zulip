package zulip

// Recipients - A message's tentative target audience.
// For channel messages, the integer Id of the channel.
// For direct messages, a list containing integer user Ids.
type Recipient struct {
	Users   *[]int64
	Channel *int64
}

// UsersAsRecipient is a convenience function that returns a list of UserIds wrapped in Recipient
func UsersAsRecipient(v []int64) Recipient {
	return Recipient{
		Users: &v,
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
		Users: &[]int64{v},
	}
}

// ChannelAsReipient is a convenience function that returns ChannelId wrapped in Recipient
func ChannelAsReipient(v int64) Recipient {
	return Recipient{
		Channel: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Recipient) UnmarshalJSON(data []byte) error {
	return UnionUnmarshalJSON(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Recipient) MarshalJSON() ([]byte, error) {
	return UnionMarshalJSON(src)
}
