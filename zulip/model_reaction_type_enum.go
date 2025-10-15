package zulip

type ReactionType string

const (
	ReactionTypeUnicodeEmoji    ReactionType = "unicode_emoji"
	ReactionTypeRealmEmoji      ReactionType = "realm_emoji"
	ReactionTypeZulipExtraEmoji ReactionType = "zulip_extra_emoji"
	ReactionTypeEmpty           ReactionType = ""
)

var AllowedReactionTypeEnumValues = []ReactionType{
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
			Enum:    AllowedReactionTypeEnumValues,
			VarName: "ReactionType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReactionType) IsValid() bool {
	for _, existing := range AllowedReactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
