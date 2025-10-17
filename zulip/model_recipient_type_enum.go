package zulip

type RecipientType string

const (
	RecipientTypeEmpty   RecipientType = ""
	RecipientTypeDirect  RecipientType = "direct"
	RecipientTypePrivate RecipientType = "private"
	RecipientTypeChannel RecipientType = "channel"
	RecipientTypeStream  RecipientType = "stream" // Legacy value, maps to RecipientTypeChannel
)

var AllowedRecipientTypeEnumValues = []RecipientType{
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
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedRecipientTypeEnumValues,
			VarName: "RecipientType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecipientType) IsValid() bool {
	for _, existing := range AllowedRecipientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
