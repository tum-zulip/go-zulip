package api

type RecipientType string

const (
	RecipientTypeDirect  RecipientType = "direct"
	RecipientTypeStream  RecipientType = "stream"
	RecipientTypePrivate RecipientType = "private"
	RecipientTypeChannel RecipientType = "channel"
)

var AllowedRecipientTypeEnumValues = []RecipientType{
	RecipientTypeDirect,
	RecipientTypeStream,
	RecipientTypePrivate,
	RecipientTypeChannel,
}

func NewRecipientTypeFromValue(v string) (*RecipientType, error) {
	ev := RecipientType(v)
	if ev.IsValid() {
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
