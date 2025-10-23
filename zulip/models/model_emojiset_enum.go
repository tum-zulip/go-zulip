package models

// The user's configured [emoji set], used to display emoji to the user everywhere they appear in the UI.
//   - EmojisetGoogle = Google modern
//   - EmojisetGoogleBlob = Google classic
//   - EmojisetTwitter = Twitter
//   - EmojisetText = Plain text
//
// [emoji set]: https://zulip.com/help/emoji-and-emoticons#use-emoticons
type Emojiset string

const (
	EmojisetGoogle     Emojiset = "google"
	EmojisetGoogleBlob Emojiset = "google-blob"
	EmojisetTwitter    Emojiset = "twitter"
	EmojisetText       Emojiset = "text"
)

var allowedEmojisetEnumValues = []Emojiset{
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
			Enum:    allowedEmojisetEnumValues,
			Value:   v,
			VarName: "Emojiset",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Emojiset) IsValid() bool {
	for _, existing := range allowedEmojisetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
