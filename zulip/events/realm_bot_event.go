package events

import "github.com/tum-zulip/go-zulip/zulip"

// RealmBotAddEvent Event sent to users who can administer a newly created bot user. Clients will also receive a `realm_user` event that contains basic details (but not the API key).  The `realm_user` events are sufficient for clients that only need to interact with the bot; this `realm_bot` event type is relevant only for administering bots.  Only organization administrators and the user who owns the bot will receive this event.
type RealmBotAddEvent struct {
	event

	Bot zulip.Bot `json:"bot,omitempty"`
}

// RealmBotUpdateEvent Event sent to users who can administer a bot user when the bot is configured. Clients may also receive a `realm_user` event that for changes in public data about the bot (name, etc.).  The `realm_user` events are sufficient for clients that only need to interact with the bot; this `realm_bot` event type is relevant only for administering bots.  Only organization administrators and the user who owns the bot will receive this event.
type RealmBotUpdateEvent struct {
	event

	// Object containing details about the changed bot. It contains two properties: the user ID of the bot and the property to be changed. The changed property is one of the remaining properties listed below.
	Bot zulip.Bot `json:"bot,omitempty"`
}

// RealmBotDeleteEvent Event sent to all users when a bot has been deactivated. Note that this is very similar to the bot_remove event and one of them will be removed soon.
type RealmBotDeleteEvent struct {
	event

	Bot UserInfo `json:"bot,omitempty"`
}
