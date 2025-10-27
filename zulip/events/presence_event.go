package events

import "github.com/tum-zulip/go-zulip/zulip"

// PresenceEvent Event sent to all users in an organization when a user comes back online after being offline for a while.  In addition to handling these events, a client that wants to maintain presence data must poll the [main presence endpoint]. Most updates to presence data, refreshing the timestamps of users who are already online, do not appear in the event queue. This design is an optimization by allowing those updates to be batched up, because there is no urgency in the information that an already-online user is still online.  These events are provided because when a user transitions from offline to online, that is information the client may want to show promptly in the UI to avoid showing a confusing state (for example, if the newly-online user sends a message or otherwise demonstrates they're online).  If the client supports the `simplified_presence_events` [client capability], the `simplified_presence_events` client capability did not exist. Therefore, all events were in the legacy format, and did not include the `presences` field.  Prior to Zulip 8.0 (feature level 228), this event was sent to all users in the organization.
//
// **Changes**: Prior to Zulip 11.0 (feature level 419
//
// [main presence endpoint]: https://zulip.com/api/get-presence
// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities, these events will include the `presences` field, which provides the modified user's presence data in the modern format. Clients are strongly encouraged to implement this client capability, as legacy format support will be removed in a future release.  If the `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level setting is set to `true`, then the event is only sent to users who can access the user who came back online.
type PresenceEvent struct {
	event

	// Only present for clients that support the `simplified_presence_events` [client capability] for the modified user(s). Clients should support updating multiple users in a single event.
	//
	// **Changes**: New in Zulip 11.0 (feature level 419).
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities.  A dictionary mapping user IDs to the presence data (modern format
	Presences map[string]zulip.ModernPresenceFormat `json:"presences,omitempty"`

	// todo: custom unmarshal
	// Not present for clients that support the `simplified_presence_events` [client capability].
	//
	// [client capability]: https://zulip.com/api/register-queue#parameter-client_capabilities
	Deprecated *PresenceEventDeprecated
}

type PresenceEventDeprecated struct {
	// The ID of the modified user.
	UserID int64 `json:"user_id,omitempty"`
	// The Zulip API email of the user.  **Deprecated**: This field will be removed in a future release as it is redundant with the `user_id`.
	// Deprecated
	Email string `json:"email,omitempty"`
	// The timestamp of when the Zulip server received the user's presence as a UNIX timestamp.
	ServerTimestamp float32 `json:"server_timestamp,omitempty"`
	// Object containing the presence data (legacy format) of of the modified user.
	// PresenceFormatDeprecated `{client_name}`: Object containing the details of the user's presence.
	//
	// **Changes**: Starting with Zulip 7.0 (feature level 178), this will always be `"website"` as the server no longer stores which client submitted presence updates.  Previously, the object key was the client's platform name, for example `website` or `ZulipDesktop`.
	Presence map[string]interface{} `json:"presence,omitempty"`
}
