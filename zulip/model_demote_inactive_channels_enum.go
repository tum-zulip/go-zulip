package zulip

// Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels) in the left sidebar.
// - 1 - Automatic
// - 2 - Always
// - 3 - Never
type DemoteInactiveChannels int

const (
	DemoteInactiveChannelsAutomatic DemoteInactiveChannels = 1
	DemoteInactiveChannelsAlways    DemoteInactiveChannels = 2
	DemoteInactiveChannelsNever     DemoteInactiveChannels = 3
)

var AllowedDemoteInactiveChannelsEnumValues = []DemoteInactiveChannels{
	DemoteInactiveChannelsAutomatic,
	DemoteInactiveChannelsAlways,
	DemoteInactiveChannelsNever,
}

func NewDemoteInactiveChannelsFromValue(v int) (*DemoteInactiveChannels, error) {
	ev := DemoteInactiveChannels(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedDemoteInactiveChannelsEnumValues,
			Value:   v,
			VarName: "DemoteInactiveChannels",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DemoteInactiveChannels) IsValid() bool {
	for _, existing := range AllowedDemoteInactiveChannelsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
