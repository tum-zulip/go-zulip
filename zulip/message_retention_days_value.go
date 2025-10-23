package zulip

import "github.com/tum-zulip/go-zulip/zulip/internal/utils"

// MessageRetentionDaysValue - Number of days that messages sent to this channel will be stored before being automatically deleted by the [message retention policy]. Two special string format values are supported:  - `"realm_default"`: Return to the organization-level setting. - `"unlimited"`: Retain messages forever.
//
// **Changes**: Prior to Zulip 5.0 (feature level 91), retaining messages forever was encoded using `MessageRetentionDaysForever` instead of `MessageRetentionDaysUnlimited`.  New in Zulip 3.0 (feature level 17).
//
// [message retention policy]: https://zulip.com/help/message-retention-policy
type MessageRetentionDaysValue struct {
	Days          *int64
	RetentionDays *MessageRetentionDays
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageRetentionDaysValue) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageRetentionDaysValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalUnionType(src)
}
