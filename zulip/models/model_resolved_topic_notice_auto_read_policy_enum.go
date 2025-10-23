package models

// Controls whether the resolved-topic notices are marked as read.
//   - ResolvedTopicNoticeAutoReadPolicyAlways = Always mark resolved-topic notices as read.
//   - ResolvedTopicNoticeAutoReadPolicyExceptFollowed = Mark resolved-topic notices as read in topics not followed by the user.
//   - ResolvedTopicNoticeAutoReadPolicyNever = Never mark resolved-topic notices as read.
//
// **Changes**: New in Zulip 11.0 (feature level 385).
type ResolvedTopicNoticeAutoReadPolicy string

const (
	ResolvedTopicNoticeAutoReadPolicyAlways         ResolvedTopicNoticeAutoReadPolicy = "always"
	ResolvedTopicNoticeAutoReadPolicyExceptFollowed ResolvedTopicNoticeAutoReadPolicy = "except_followed"
	ResolvedTopicNoticeAutoReadPolicyNever          ResolvedTopicNoticeAutoReadPolicy = "never"
)

var allowedResolvedTopicNoticeAutoReadPolicyEnumValues = []ResolvedTopicNoticeAutoReadPolicy{
	ResolvedTopicNoticeAutoReadPolicyAlways,
	ResolvedTopicNoticeAutoReadPolicyExceptFollowed,
	ResolvedTopicNoticeAutoReadPolicyNever,
}

func NewResolvedTopicNoticeAutoReadPolicyFromValue(v string) (*ResolvedTopicNoticeAutoReadPolicy, error) {
	ev := ResolvedTopicNoticeAutoReadPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedResolvedTopicNoticeAutoReadPolicyEnumValues,
			VarName: "ResolvedTopicNoticeAutoReadPolicy",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResolvedTopicNoticeAutoReadPolicy) IsValid() bool {
	for _, existing := range allowedResolvedTopicNoticeAutoReadPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
