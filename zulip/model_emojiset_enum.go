package zulip

// The user's configured [emoji set](https://zulip.com/help/emoji-and-emoticons#use-emoticons), used to display emoji to the user everywhere they appear in the UI.
// - "google" - Google modern
// - "google-blob" - Google classic
// - "twitter" - Twitter
// - "text" - Plain text
type Emojiset string

const (
	EmojisetGoogle     Emojiset = "google"
	EmojisetGoogleBlob Emojiset = "google-blob"
	EmojisetTwitter    Emojiset = "twitter"
	EmojisetText       Emojiset = "text"
)

var AllowedEmojisetEnumValues = []Emojiset{
	EmojisetGoogle,
	EmojisetGoogleBlob,
	EmojisetTwitter,
	EmojisetText,
}

func NewEmojisetFromValue(v string) (*Emojiset, error) {
	ev := Emojiset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedEmojisetEnumValues,
			Value:   v,
			VarName: "Emojiset",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Emojiset) IsValid() bool {
	for _, existing := range AllowedEmojisetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
