package zulip

type UserStatus struct {
	// If present, the user has marked themself \"away\".
	//
	// **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, `away` is a legacy way to access the user's `presence_enabled` setting, with `away = !presence_enabled`. To be removed in a future release.
	// Deprecated
	Away bool `json:"away,omitempty"`
	// If present, the text content of the user's status message.
	StatusText string `json:"status_text,omitempty"`
	// If present, the name for the emoji to associate with the user's status.
	//
	// **Changes**: New in Zulip 5.0 (feature level 86).
	EmojiName string `json:"emoji_name,omitempty"`
	// If present, a unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
	//
	// **Changes**: New in Zulip 5.0 (feature level 86).
	EmojiCode string `json:"emoji_code,omitempty"`
	// If present, a string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the Id of   the uploaded [custom emoji].  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \"zulip\").
	//
	// **Changes**: New in Zulip 5.0 (feature level 86).
	//
	// [custom emoji]: https://zulip.com/help/custom-emoji
	ReactionType ReactionType `json:"reaction_type,omitempty"`
}
