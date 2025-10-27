package zulip

import (
	"github.com/tum-zulip/go-zulip/zulip/internal/union"
)

// SubscriptionData struct for SubscriptionData.
type SubscriptionData struct {
	// The unique ID of a channel.
	ChannelID int64 `json:"stream_id"`
	// One of the channel properties described below.
	//  - SubscriptionPropertyColor
	//  - SubscriptionPropertyIsMuted
	//  - SubscriptionPropertyPinToTop
	//  - SubscriptionPropertyDesktopNotifications
	//  - SubscriptionPropertyAudibleNotifications
	//  - SubscriptionPropertyPushNotifications
	//  - SubscriptionPropertyEmailNotifications
	//  - SubscriptionPropertyWildcardMentionsNotify
	//  - SubscriptionPropertyInHomeView
	Property SubscriptionProperty `json:"property"`
	// SubscriptionDataValue - The new value of the property being modified.  If the property is `"color"`, then `value` is a string representing the hex value of the user's display color for the channel. For all other above properties, `value` is a boolean.
	Value SubscriptionDataValue `json:"value"`
}

type SubscriptionDataValue struct {
	Bool   *bool
	String *string
}

// Unmarshal JSON data into one of the pointers in the struct.
func (s *SubscriptionDataValue) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, s)
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (s SubscriptionDataValue) MarshalJSON() ([]byte, error) {
	return union.Marshal(s)
}
