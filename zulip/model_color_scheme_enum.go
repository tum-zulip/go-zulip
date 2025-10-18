package zulip

// Controls which [color theme] to use.
//   - 1 = Automatic
//   - 2 = Dark theme
//   - 3 = Light theme
//
// Automatic detection is implementing using the standard `prefers-color-scheme` media query.
//
// [color theme]: https://zulip.com/help/dark-theme
type ColorScheme int

const (
	ColorSchemeAutomatic ColorScheme = 1
	ColorSchemeDark      ColorScheme = 2
	ColorSchemeLight     ColorScheme = 3
)

var AllowedColorSchemeEnumValues = []ColorScheme{
	ColorSchemeAutomatic,
	ColorSchemeDark,
	ColorSchemeLight,
}

func NewColorSchemeFromValue(v int) (*ColorScheme, error) {
	ev := ColorScheme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Enum:    AllowedColorSchemeEnumValues,
			Value:   v,
			VarName: "ColorScheme",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ColorScheme) IsValid() bool {
	for _, existing := range AllowedColorSchemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
