package models

// A string indicating the type of emoji.
//
// Each emoji `reaction_type` has an independent namespace for values of `emoji_code`.  Must be one of the following values:
//   - ReactionTypeRealmEmoji
//   - ReactionTypeUnicodeEmoji
//   - ReactionTypeZulipExtraEmoji
//   - ReactionTypeEmpty
type ReactionType string

const (
	// In this namespace, `emoji_code` will be the Id of the uploaded [custom emoji].
	//
	// [custom emoji]: https://zulip.com/help/custom-emoji
	ReactionTypeRealmEmoji      ReactionType = "realm_emoji"
	ReactionTypeUnicodeEmoji    ReactionType = "unicode_emoji"     // In this namespace, `emoji_code` will be a  dash-separated hex encoding of the sequence of Unicode codepoints that define this emoji in the Unicode specification.
	ReactionTypeZulipExtraEmoji ReactionType = "zulip_extra_emoji" // These are special emoji included with Zulip. In this namespace, `emoji_code` will be the name of the emoji (e.g. "zulip").
	ReactionTypeEmpty           ReactionType = ""                  // For users who set a status without selecting an emoji.
)

var allowedReactionTypeEnumValues = []ReactionType{
	ReactionTypeUnicodeEmoji,
	ReactionTypeRealmEmoji,
	ReactionTypeZulipExtraEmoji,
	ReactionTypeEmpty,
}

func NewReactionTypeFromValue(v string) (*ReactionType, error) {
	ev := ReactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedReactionTypeEnumValues,
			VarName: "ReactionType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReactionType) IsValid() bool {
	for _, existing := range allowedReactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
