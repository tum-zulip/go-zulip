package zulip

// Controls how animated images should be played in the message feed in the web/desktop application.
// - "always" - Always play the animated images in the message feed.
// - "on_hover" - Play the animated images on hover over them in the message feed.
// - "never" - Never play animated images in the message feed.
// **Changes**: New in Zulip 9.0 (feature level 275).
type WebAnimateImagePreviews string

const (
	WebAnimateImagePreviewsAlways  WebAnimateImagePreviews = "always"
	WebAnimateImagePreviewsOnHover WebAnimateImagePreviews = "on_hover"
	WebAnimateImagePreviewsNever   WebAnimateImagePreviews = "never"
)

var AllowedWebAnimateImagePreviewsEnumValues = []WebAnimateImagePreviews{
	WebAnimateImagePreviewsAlways,
	WebAnimateImagePreviewsOnHover,
	WebAnimateImagePreviewsNever,
}

func NewWebAnimateImagePreviewsFromValue(v string) (*WebAnimateImagePreviews, error) {
	ev := WebAnimateImagePreviews(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedWebAnimateImagePreviewsEnumValues,
			Value:   v,
			VarName: "WebAnimateImagePreviews",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebAnimateImagePreviews) IsValid() bool {
	for _, existing := range AllowedWebAnimateImagePreviewsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
