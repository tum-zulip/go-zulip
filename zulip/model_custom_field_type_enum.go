package zulip

// See the [Custom profile fields] article for details on what each type means.
//   - CustomFieldTypeShortText = Short text
//   - CustomFieldTypeLongText = Long text
//   - CustomFieldTypeListOfOptions = List of options
//   - CustomFieldTypeDatePicker = Date picker
//   - CustomFieldTypeLink = Link
//   - CustomFieldTypePersonPicker = Person picker
//   - CustomFieldTypeExternalAccount = External account
//   - CustomFieldTypePronouns = Pronouns
type CustomFieldType int

const (
	CustomFieldTypeShortText       CustomFieldType = 1 // Short text
	CustomFieldTypeLongText        CustomFieldType = 2 // Long text
	CustomFieldTypeListOfOptions   CustomFieldType = 3 // List of options
	CustomFieldTypeDatePicker      CustomFieldType = 4 // Date picker
	CustomFieldTypeLink            CustomFieldType = 5 // Link
	CustomFieldTypePersonPicker    CustomFieldType = 6 // Person picker
	CustomFieldTypeExternalAccount CustomFieldType = 7 // External account
	CustomFieldTypePronouns        CustomFieldType = 8 // Pronouns
)

var allowedCustomFieldTypes = []CustomFieldType{
	CustomFieldTypeShortText,
	CustomFieldTypeLongText,
	CustomFieldTypeListOfOptions,
	CustomFieldTypeDatePicker,
	CustomFieldTypeLink,
	CustomFieldTypePersonPicker,
	CustomFieldTypeExternalAccount,
	CustomFieldTypePronouns,
}

func NewCustomFieldTypeFromValue(v int) (*CustomFieldType, error) {
	ev := CustomFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, &ErrInvalidEnumValue{
			Value:   v,
			Enum:    allowedCustomFieldTypes,
			VarName: "CustomFieldType",
		}
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldType) IsValid() bool {
	for _, existing := range allowedCustomFieldTypes {
		if existing == v {
			return true
		}
	}
	return false
}
