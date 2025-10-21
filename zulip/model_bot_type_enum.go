package zulip

// BotType types Zulip supports
//   - BotTypeGeneric = Generic bot
//   - BotTypeIncomingWebhook = Incoming webhook
//   - BotTypeOutgoingWebhook = Outgoing webhook
//   - BotTypeEmbedded = Embedded bot
//
// These can be used to identify the type of a bot user.
type BotType int

const (
	BotTypeGeneric         BotType = 1
	BotTypeIncomingWebhook BotType = 2
	BotTypeOutgoingWebhook BotType = 3
	BotTypeEmbedded        BotType = 4
)

var allowedBotTypeEnumValues = []BotType{
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
			Enum:    allowedBotTypeEnumValues,
			VarName: "BotType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BotType) IsValid() bool {
	for _, existing := range allowedBotTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
