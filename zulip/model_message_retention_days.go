package zulip

const (
	MessageRetentionDaysRealmDefault = "realm_default"
	MessageRetentionDaysUnlimited    = "unlimited"
	// Deprecated: use MessageRetentionDaysUnlimited
	MessageRetentionDaysForever = "forever"
)

// MessageRetentionDays - Number of days that messages sent to this channel will be stored before being automatically deleted by the [message retention policy](zulip.com/help/message-retention-policy. Two special string format values are supported:  - `\"realm_default\"`: Return to the organization-level setting. - `\"unlimited\"`: Retain messages forever.  **Changes**: Prior to Zulip 5.0 (feature level 91), retaining messages forever was encoded using `\"forever\"` instead of `\"unlimited\"`.  New in Zulip 3.0 (feature level 17).
type MessageRetentionDays struct {
	Int32  *int32
	String *string
}

// int32AsMessageRetentionDays is a convenience function that returns int32 wrapped in MessageRetentionDays
func Int32AsMessageRetentionDays(v *int32) MessageRetentionDays {
	return MessageRetentionDays{
		Int32: v,
	}
}

// stringAsMessageRetentionDays is a convenience function that returns string wrapped in MessageRetentionDays
func StringAsMessageRetentionDays(v *string) MessageRetentionDays {
	return MessageRetentionDays{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageRetentionDays) UnmarshalJSON(data []byte) error {
	return unmarshalUnionType(data, dst)
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageRetentionDays) MarshalJSON() ([]byte, error) {
	return marshalUnionType(src)
}
