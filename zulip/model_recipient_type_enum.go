package zulip

import "encoding/json"

// RecipientType - Enum for RecipientType
//   - RecipientTypeEmpty = An unaddressed draft.
//   - RecipientTypeDirect = A direct message draft (one-on-one or group).
//   - RecipientTypePrivate = Legacy value, maps to RecipientTypeDirect.
//   - RecipientTypeChannel = A channel message draft.
//   - RecipientTypeStream = Legacy value, maps to RecipientTypeChannel.
//
// When unmarshaling, `RecipientTypeStream` will always be converted to `RecipientTypeChannel`, and `RecipientTypePrivate` to `RecipientTypeDirect`, but when marshaling, the values will be as-is, because some API endpoints still expect the legacy values.
type RecipientType string

const (
	RecipientTypeEmpty   RecipientType = ""        //  An unaddressed draft.
	RecipientTypeDirect  RecipientType = "direct"  //  A direct message draft (one-on-one or group).
	RecipientTypePrivate RecipientType = "private" //  Legacy value, maps to RecipientTypeDirect.
	RecipientTypeChannel RecipientType = "channel" //  A channel message draft.
	RecipientTypeStream  RecipientType = "stream"  //  Legacy value, maps to RecipientTypeChannel.
)

var allowedRecipientTypeEnumValues = []RecipientType{
	RecipientTypeEmpty,
	RecipientTypeDirect,
	RecipientTypePrivate,
	RecipientTypeChannel,
	RecipientTypeStream,
}

func (v RecipientType) ToLegacy() RecipientType {
	if v == RecipientTypeChannel {
		return RecipientTypeStream
	}
	return v
}

func NewRecipientTypeFromValue(v string) (*RecipientType, error) {
	ev := RecipientType(v)

	if ev.IsValid() {
		if ev == RecipientTypeStream {
			ev = RecipientTypeChannel
		}
		if ev == RecipientTypePrivate {
			ev = RecipientTypeDirect
		}
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedRecipientTypeEnumValues,
			VarName: "RecipientType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecipientType) IsValid() bool {
	for _, existing := range allowedRecipientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v RecipientType) IsDirectMessage() bool {
	return v == RecipientTypeDirect || v == RecipientTypePrivate
}

func (v RecipientType) IsChannelMessage() bool {
	return v == RecipientTypeChannel || v == RecipientTypeStream
}

func (o *RecipientType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	ev, err := NewRecipientTypeFromValue(val)
	if err != nil {
		return err
	}
	*o = *ev
	return nil
}
