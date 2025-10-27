package events

// HasZoomTokenEvent Event sent to a user's clients when the user completes the OAuth flow for the [Zoom integration]. Clients need to know whether initiating Zoom OAuth is required before creating a Zoom call.
//
// [Zoom integration]: https://zulip.com/help/configure-call-provider
type HasZoomTokenEvent struct {
	event

	// A boolean specifying whether the user has zoom token or not.
	Value bool `json:"value,omitempty"`
}
