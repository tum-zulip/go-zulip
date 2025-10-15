package zulip

// Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels) in the left sidebar.
// - 1 - Automatic
// - 2 - Always
// - 3 - Never
type DemoteInactiveStreams int

const (
	DemoteInactiveStreamsAutomatic DemoteInactiveStreams = 1
	DemoteInactiveStreamsAlways    DemoteInactiveStreams = 2
	DemoteInactiveStreamsNever     DemoteInactiveStreams = 3
)

var AllowedDemoteInactiveStreamsEnumValues = []DemoteInactiveStreams{
	DemoteInactiveStreamsAutomatic,
	DemoteInactiveStreamsAlways,
	DemoteInactiveStreamsNever,
}

func NewDemoteInactiveStreamsFromValue(v int) (*DemoteInactiveStreams, error) {
	ev := DemoteInactiveStreams(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedDemoteInactiveStreamsEnumValues,
			Value:   v,
			VarName: "DemoteInactiveStreams",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DemoteInactiveStreams) IsValid() bool {
	for _, existing := range AllowedDemoteInactiveStreamsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
