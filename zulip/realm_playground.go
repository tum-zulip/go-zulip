package zulip

// RealmPlayground Object containing details about a realm playground.
type RealmPlayground struct {
	// The unique ID for the realm playground.
	ID int64 `json:"id,omitempty"`
	// The user-visible display name of the playground. Clients should display this in UI for picking which playground to open a code block in, to differentiate between multiple configured playground options for a given pygments language.
	//
	// **Changes**: New in Zulip 4.0 (feature level 49).
	Name string `json:"name,omitempty"`
	// The name of the Pygments language lexer for that programming language.
	PygmentsLanguage string `json:"pygments_language,omitempty"`
	// The [RFC 6570] compliant URL template for the playground. The template contains exactly one variable named `code`, which determines how the extracted code should be substituted in the playground URL.
	//
	// **Changes**: New in Zulip 8.0 (feature level 196). This replaced the `url_prefix` parameter, which was used to construct URLs by just concatenating url_prefix and code.
	//
	// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
	URLTemplate string `json:"url_template,omitempty"`
}
