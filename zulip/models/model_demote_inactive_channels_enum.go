package models

// Whether to [hide inactive channels] in the left sidebar.
//   - DemoteInactiveChannelsAutomatic = Automatic
//   - DemoteInactiveChannelsAlways = Always
//   - DemoteInactiveChannelsNever = Never
//
// [hide inactive channels]: https://zulip.com/help/manage-inactive-channels
type DemoteInactiveChannels int

const (
	DemoteInactiveChannelsAutomatic DemoteInactiveChannels = 1
	DemoteInactiveChannelsAlways    DemoteInactiveChannels = 2
	DemoteInactiveChannelsNever     DemoteInactiveChannels = 3
)

var allowedDemoteInactiveChannelsEnumValues = []DemoteInactiveChannels{
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
			Enum:    allowedDemoteInactiveChannelsEnumValues,
			Value:   v,
			VarName: "DemoteInactiveChannels",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DemoteInactiveChannels) IsValid() bool {
	for _, existing := range allowedDemoteInactiveChannelsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
