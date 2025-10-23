package models

// EmojiReaction struct for EmojiReaction
type EmojiReaction struct {
	// Name of the emoji.
	EmojiName string `json:"emoji_name,omitempty"`
	// A unique identifier, defining the specific emoji codepoint requested, within the namespace of the `reaction_type`.
	EmojiCode string `json:"emoji_code,omitempty"`
	// A string indicating the type of emoji. Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:
	//   - ReactionTypeRealmEmoji
	//   - ReactionTypeUnicodeEmoji
	//   - ReactionTypeZulipExtraEmoji
	//   - ReactionTypeEmpty
	ReactionType ReactionType `json:"reaction_type,omitempty"`
	// The Id of the user who added the reaction.
	//
	// **Changes**: New in Zulip 3.0 (feature level 2). The `user` object is deprecated and will be removed in the future.  In Zulip 10.0 (feature level 328), the deprecated `user` object was removed which contained the following properties: `id`, `email`, `full_name` and `is_mirror_dummy`.
	UserId int64 `json:"user_id,omitempty"`
}
