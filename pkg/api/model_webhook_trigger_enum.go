package api

// What aspect of the message triggered the outgoing webhook notification. Possible values include `direct_message` and `mention`.  **Changes**: In Zulip 8.0 (feature level 201), renamed the trigger `private_message` to `direct_message`.
type WebhookTrigger string

const (
	WebhookTriggerDirectMessage WebhookTrigger = "direct_message"
	WebhookTriggerMention       WebhookTrigger = "mention"
)

var AllowedWebhookTriggerEnumValues = []WebhookTrigger{
	WebhookTriggerDirectMessage,
	WebhookTriggerMention,
}

func NewWebhookTriggerFromValue(v string) (*WebhookTrigger, error) {
	ev := WebhookTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedWebhookTriggerEnumValues,
			Value:   v,
			VarName: "WebhookTrigger",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookTrigger) IsValid() bool {
	for _, existing := range AllowedWebhookTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
