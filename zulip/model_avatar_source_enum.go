package zulip

// The avatar data source type for the current user.
// - "G" - Gravatar
// - "U" - Uploaded by user
type AvatarSource string

const (
	AvatarSourceGravatar AvatarSource = "G"
	AvatarSourceUploaded AvatarSource = "U"
)

var AllowedAvatarSourceEnumValues = []AvatarSource{
	AvatarSourceGravatar,
	AvatarSourceUploaded,
}

func NewAvatarSourceFromValue(v string) (*AvatarSource, error) {
	ev := AvatarSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    AllowedAvatarSourceEnumValues,
			VarName: "AvatarSource",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AvatarSource) IsValid() bool {
	for _, existing := range AllowedAvatarSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
