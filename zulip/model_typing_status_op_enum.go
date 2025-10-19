package zulip

// TypingStatusOp represents whether the user has started or stopped typing.
//   - TypingStatusOpStart = start
//   - TypingStatusOpStop = stop
type TypingStatusOp string

const (
	TypingStatusOpStart TypingStatusOp = "start"
	TypingStatusOpStop  TypingStatusOp = "stop"
)

// All allowed values of TypingStatusOp enum
var AllowedTypingStatusOpEnumValues = []TypingStatusOp{
	TypingStatusOpStart,
	TypingStatusOpStop,
}

// NewTypingStatusOpFromValue returns a pointer to a valid TypingStatusOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypingStatusOpFromValue(v string) (*TypingStatusOp, error) {
	ev := TypingStatusOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedTypingStatusOpEnumValues,
			VarName: "TypingStatusOp",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypingStatusOp) IsValid() bool {
	for _, existing := range AllowedTypingStatusOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
