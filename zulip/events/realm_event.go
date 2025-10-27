package events

import (
	"github.com/tum-zulip/go-zulip/internal/union"
)

// RealmUpdateEvent The simpler of two possible event types sent to all users in a Zulip organization when the configuration of the organization (realm) has changed.  Often individual settings are migrated from this format to the [realm/update_dict] event format when additional realm settings are added whose values are coupled to each other in some way. The specific values supported by this event type are documented in the [realm/update_dict] documentation.  A correct client implementation should convert these events into the corresponding [realm/update_dict] event and then process that.
//
// **Changes**: Removed `extra_data` optional property in Zulip 10.0 (feature level 306). The `extra_data` used to include an `upload_quota` field when changed property was `plan_type`. The server now sends a standard `realm/update_dict` event for plan changes.
//
// [realm/update_dict]: https://zulip.com/api/get-events#realm-update_dict
type RealmUpdateEvent struct {
	event

	// The name of the property that was changed.
	Property string      `json:"property,omitempty"`
	Value    UpdateValue `json:"value,omitempty"`
}

// RealmDeactivatedEvent Event sent to all users in a Zulip organization when the organization (realm) is deactivated. Its main purpose is to flush active longpolling connections so clients can immediately show the organization as deactivated.  Clients cannot rely on receiving this event, because they will no longer be able to authenticate to the Zulip API due to the deactivation, and thus can miss it if they did not have an active longpolling connection at the moment of deactivation.  Correct handling of realm deactivations requires that clients parse authentication errors from GET /events; if that is done correctly, the client can ignore this event type and rely on its handling of the `GET /events` request it will do immediately after processing this batch of events.
type RealmDeactivatedEvent struct {
	event

	// The ID of the deactivated realm.
	RealmID int64 `json:"realm_id,omitempty"`
}

// RealmUpdateDictEvent The more general of two event types that may be used when sending an event to all users in a Zulip organization when the configuration of the organization (realm) has changed.  Unlike the simpler [realm/update] event format, this event type supports multiple properties being changed in a single event.  This event is also sent when deactivating or reactivating a user for settings set to anonymous user groups which the user is direct member of. When deactivating the user, event is only sent to users who cannot access the deactivated user.
//
// **Changes**: Starting with Zulip 10.0 (feature level 303), this event can also be sent when deactivating or reactivating a user.  In Zulip 7.0 (feature level 163), the realm setting `email_address_visibility` was removed. It was replaced by a [user setting] with a [realm user default], with the encoding of different values preserved. Clients can support all versions by supporting the current API and treating every user as having the realm's `email_address_visibility` value.
//
// [realm/update]: https://zulip.com/api/get-events#realm-update
// [realm user default]: https://zulip.com/api/update-realm-user-settings-defaults#parameter-email_address_visibility
// [user setting]: https://zulip.com/api/update-settings#parameter-email_address_visibility
type RealmUpdateDictEvent struct {
	event

	// Always `"default"`. Present for backwards-compatibility with older clients that predate the `update_dict` event style.  **Deprecated** and will be removed in a future release.
	// Deprecated
	Property           string             `json:"property,omitempty"`
	RealmConfiguration RealmConfiguration `json:"data,omitempty"`
}

// UpdateValue - New value of the changed setting.
type UpdateValue struct {
	Bool   *bool
	Int64  *int64
	String *string
}

// Unmarshal JSON data into one of the pointers in the struct.
func (u *UpdateValue) UnmarshalJSON(data []byte) error {
	return union.Unmarshal(data, u)
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (u UpdateValue) MarshalJSON() ([]byte, error) {
	return union.Marshal(u)
}
