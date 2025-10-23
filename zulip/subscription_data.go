package zulip

import "github.com/tum-zulip/go-zulip/zulip/internal/utils"

// SubscriptionData struct for SubscriptionData
type SubscriptionData struct {
	// The unique Id of a channel.
	ChannelId int64 `json:"stream_id"`
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

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionDataValue) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionDataValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalUnionType(src)
}
