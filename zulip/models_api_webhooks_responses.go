package zulip

// ZulipOutgoingWebhooksResponse This is an example of the JSON payload that the Zulip server will `POST` to your server:
type ZulipOutgoingWebhooksResponse struct {
	// Email of the bot user.
	BotEmail string `json:"bot_email,omitempty"`
	// The full name of the bot user.
	BotFullName string `json:"bot_full_name,omitempty"`
	// The message content, in raw [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown) format (not rendered to HTML).
	Data string `json:"data,omitempty"`
	// What aspect of the message triggered the outgoing webhook notification. Possible values include `direct_message` and `mention`.  **Changes**: In Zulip 8.0 (feature level 201), renamed the trigger `private_message` to `direct_message`.
	Trigger WebhookTrigger `json:"trigger,omitempty"`
	// A string of alphanumeric characters that can be used to authenticate the webhook request (each bot user uses a fixed token). You can get the token used by a given outgoing webhook bot in the `zuliprc` file downloaded when creating the bot.
	Token string `json:"token,omitempty"`
	// Details on the message that triggered the outgoing webhook, matching the format used by GET /messages responses.
	Message *Message `json:"message,omitempty"`
}
