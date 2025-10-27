package events

import "github.com/tum-zulip/go-zulip/zulip"

// CustomProfileFieldsEvent Event sent to all users in a Zulip organization when new custom profile field types are configured for that Zulip organization.
type CustomProfileFieldsEvent struct {
	event

	// An array of dictionaries where each dictionary contains details of a single new custom profile field for the Zulip organization.
	Fields []zulip.CustomProfileField `json:"fields,omitempty"`
}
