package zulip

type BotType int

const (
	BotTypeGeneric         BotType = 1
	BotTypeIncomingWebhook BotType = 2
	BotTypeOutgoingWebhook BotType = 3
	BotTypeEmbedded        BotType = 4
)

var AllowedBotTypeEnumValues = []BotType{
	BotTypeGeneric,
	BotTypeIncomingWebhook,
	BotTypeOutgoingWebhook,
	BotTypeEmbedded,
}

func NewBotTypeFromValue(v int) (*BotType, error) {
	ev := BotType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedBotTypeEnumValues,
			VarName: "BotType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BotType) IsValid() bool {
	for _, existing := range AllowedBotTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
