package zulip

import (
	"github.com/tum-zulip/go-zulip/zulip/internal/union"
)

// Bot struct for Bot.
type Bot struct {
	// The user ID of the bot.
	UserID int64 `json:"user_id,omitempty"`
	// The full name of the bot.
	FullName string `json:"full_name,omitempty"`
	// The API key of the bot which it uses to make API requests.
	APIKey string `json:"api_key,omitempty"`
	// The default sending channel of the bot. If `null`, the bot doesn't have a default sending channel.
	DefaultSendingChannel *string `json:"default_sending_stream,omitempty"`
	// The default channel for which the bot receives events/register data. If `null`, the bot doesn't have such a default channel.
	DefaultEventsRegisterChannel *string `json:"default_events_register_stream,omitempty"`
	// Whether the bot can send messages to all channels by default.
	DefaultAllPublicChannels bool `json:"default_all_public_streams,omitempty"`
	// The URL of the bot's avatar.
	AvatarURL string `json:"avatar_url,omitempty"`
	// The user ID of the bot's owner.  If `null`, the bot has no owner.
	OwnerID *int64 `json:"owner_id,omitempty"`
	// An array containing extra configuration fields only relevant for outgoing webhook bots and embedded bots. This is always a single-element array.  We consider this part of the Zulip API to be unstable; it is used only for UI elements for administering bots and is likely to change.
	Services []BotData `json:"services,omitempty"`
	// A boolean describing whether the user account has been deactivated.
	IsActive bool `json:"is_active,omitempty"`
}

// BotData - Object with extra configuration details for the bot. The fields in the object depend on the type of bot.
type BotData struct {
	OutgoingWebhookBotData *OutgoingWebhookBotData
	EmbeddedBotData        *EmbeddedBotData
}

// OutgoingWebhookBotDataAsBotData is a convenience function that returns OutgoingWebhookBotData wrapped in BotData.
func OutgoingWebhookBotDataAsBotData(v *OutgoingWebhookBotData) BotData {
	return BotData{
		OutgoingWebhookBotData: v,
	}
}

// EmbeddedBotDataAsBotData is a convenience function that returns EmbeddedBotData wrapped in BotData.
func EmbeddedBotDataAsBotData(v *EmbeddedBotData) BotData {
	return BotData{
		EmbeddedBotData: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (b *BotData) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, b)
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (b BotData) MarshalJSON() ([]byte, error) {
	return union.Marshal(b)
}

// OutgoingWebhookBotData When the bot is an outgoing webhook.
type OutgoingWebhookBotData struct {
	// The URL the outgoing webhook is configured to post to.
	BaseURL string `json:"base_url,omitempty"`
	// A unique token that the third-party service can use to confirm that the request is indeed coming from Zulip.
	Token string `json:"token,omitempty"`
	// An integer indicating what format requests are posted in:
	//   - 1 = Zulip's native outgoing webhook format.
	//   - 2 = Emulate the Slack outgoing webhook format.
	Interface int `json:"interface,omitempty"`
}

// EmbeddedBotData When the bot is an embedded bot.
type EmbeddedBotData struct {
	// The name of the bot.
	ServiceName string `json:"service_name,omitempty"`
	// A dictionary of string key/value pairs, which describe the configuration for the bot. These are usually details like API keys, and are unique to the integration/bot. Can be an empty dictionary.
	ConfigData map[string]string `json:"config_data,omitempty"`
}
