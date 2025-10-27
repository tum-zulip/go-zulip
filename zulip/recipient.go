package zulip

import (
	"github.com/tum-zulip/go-zulip/zulip/internal/union"
)

// Recipient - A message's tentative target audience.
// For channel messages, the integer ID of the channel.
// For direct messages, a list containing integer user IDs.
//
//nolint:recvcheck // UnmarshalJSON requires pointer receiver, but for all other methods value receiver makes more sense
type Recipient struct {
	Users   []int64
	Channel *int64
}

func (r Recipient) asArray() []int64 {
	if r.Users != nil {
		return r.Users
	} else if r.Channel != nil {
		return []int64{*r.Channel}
	}
	return nil
}

// Get the RecipientType for the recipient. Either RecipientTypeChannel or RecipientTypeDirect or nil if the recipient is invalid.
func (r Recipient) RecipientType() *RecipientType {
	if r.Users == nil && r.Channel == nil {
		return nil
	}
	if r.Users != nil && r.Channel != nil {
		return nil
	}
	t := RecipientTypeChannel
	if r.Users != nil {
		t = RecipientTypeDirect
	}
	return &t
}

// UsersAsRecipient is a convenience function that returns a list of UserIDs wrapped in Recipient.
func UsersAsRecipient(v []int64) Recipient {
	return Recipient{
		Users: v,
	}
}

// UserAsRecipient is a convenience function that returns UserID wrapped in Recipient.
func UserAsRecipient(v int64) Recipient {
	return Recipient{
		Users: []int64{v},
	}
}

// ChannelAsRecipient is a convenience function that returns ChannelID wrapped in Recipient.
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

// Unmarshal JSON data into one of the pointers in the struct.
func (r *Recipient) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, r)
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (r Recipient) MarshalJSON() ([]byte, error) {
	return union.Marshal(r)
}
