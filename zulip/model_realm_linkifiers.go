package zulip

// RealmLinkifiers struct for RealmLinkifiers
type RealmLinkifiers struct {
	// The [Python regular expression] pattern which represents the pattern that should be linkified on matching.
	//
	// [Python regular expression]: https://docs.python.org/3/howto/regex.html
	Pattern string `json:"pattern,omitempty"`
	// The [RFC 6570] compliant URL template with which the pattern matching string should be linkified.
	//
	// **Changes**: New in Zulip 7.0 (feature level 176). This replaced `url_format`, which contained a URL format string.
	//
	// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
	UrlTemplate string `json:"url_template,omitempty"`
	// The Id of the linkifier.
	Id int64 `json:"id,omitempty"`
}
