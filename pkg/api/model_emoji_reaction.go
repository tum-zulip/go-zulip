package api

// EmojiReaction struct for EmojiReaction
type EmojiReaction struct {
	// Name of the emoji.
	EmojiName string `json:"emoji_name,omitempty"`
	// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
	EmojiCode string `json:"emoji_code,omitempty"`
	// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:  - `unicode_emoji` : In this namespace, `emoji_code` will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - `realm_emoji` : In this namespace, `emoji_code` will be the Id of   the uploaded [custom emoji](zulip.com/help/custom-emoji.  - `zulip_extra_emoji` : These are special emoji included with Zulip.   In this namespace, `emoji_code` will be the name of the emoji (e.g.   \"zulip\").
	ReactionType ReactionType `json:"reaction_type,omitempty"`
	// The Id of the user who added the reaction.  **Changes**: New in Zulip 3.0 (feature level 2). The `user` object is deprecated and will be removed in the future.  In Zulip 10.0 (feature level 328), the deprecated `user` object was removed which contained the following properties: `id`, `email`, `full_name` and `is_mirror_dummy`.
	UserId int64 `json:"user_id,omitempty"`
}
