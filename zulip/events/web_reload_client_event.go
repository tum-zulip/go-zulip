package events

// WebReloadClientEvent An event which signals the official Zulip web/desktop app to update, by reloading the page and fetching a new queue; this will generally follow a `restart` event. Clients which do not obtain their code from the server (e.g. mobile and terminal clients, which store their code locally) should ignore this event.  Clients choosing to reload the application must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.
//
// **Changes**: New in Zulip 9.0 (feature level 240).
type WebReloadClientEvent struct {
	event

	// Whether the client should fetch a new event queue immediately, rather than using a backoff strategy to avoid thundering herds. A Zulip development server uses this parameter to reload clients immediately.
	Immediate bool `json:"immediate,omitempty"`
}
