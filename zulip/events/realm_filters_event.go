package events

// RealmFiltersEvent Legacy event type that is no longer sent to clients. Previously, sent to all users in a Zulip organization when the set of configured [linkifiers] for the organization was changed.
//
// **Changes**: Prior to Zulip 7.0 (feature level 176), this event type was sent to clients.  **Deprecated** in Zulip 4.0 (feature level 54), and replaced by the `realm_linkifiers` event type, which has a clearer name and format.
//
// # Deprecated
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
type RealmFiltersEvent struct {
	event
	// An array of tuples, where each tuple described a linkifier. The first element of the tuple was a string regex pattern which represented the pattern to be linkified on matching, for example `"#(?P<id>[123])"`. The second element was the URL format string that the pattern should be linkified with. A URL format string for the above example would be `"https://realm.com/my_realm_filter/%(id)s"`. And the third element was the ID of the realm filter.
	RealmFilters []interface{}
}
