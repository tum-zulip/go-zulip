package zulip

// CustomProfileField Dictionary containing the details of a custom profile field configured for this organization.
type CustomProfileField struct {
	// The ID of the custom profile field. This will be referenced in the custom profile fields section of user objects.
	ID int64 `json:"id"`
	// An integer indicating the type of the custom profile field, which determines how it is configured and displayed to users.  See the [Custom profile fields] article for details on what each type means.
	//   - CustomFieldTypeShortText
	//   - CustomFieldTypeLongText
	//   - CustomFieldTypeListOfOptions
	//   - CustomFieldTypeDatePicker
	//   - CustomFieldTypeLink
	//   - CustomFieldTypePersonPicker
	//   - CustomFieldTypeExternalAccount
	//   - CustomFieldTypePronouns
	//
	// **Changes**: Field type `CustomFieldTypePronouns` added in Zulip 6.0 (feature level 151).
	//
	// [Custom profile fields]: https://zulip.com/help/custom-profile-fields#profile-field-types
	Type CustomFieldType `json:"type"`
	// Custom profile fields are displayed in both settings UI and UI showing users' profiles in increasing `order`.
	Order int32 `json:"order"`
	// The name of the custom profile field.
	Name string `json:"name"`
	// The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.
	Hint string `json:"hint"`
	// Field types CustomFieldTypeListOfOptions (List of options) and CustomFieldTypeExternalAccount (External account) support storing additional configuration for the field type in the `field_data` attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type CustomFieldTypeExternalAccount is not yet stabilized.
	FieldData *string `json:"field_data,omitempty"`
	// Whether the custom profile field, display or not on the user card.  Currently it's value not allowed to be `true` of `Long text` and `Person picker` [profile field types].  This field is only included when its value is `true`.
	//
	// **Changes**: New in Zulip 6.0 (feature level 146).
	//
	// [profile field types]: https://zulip.com/help/custom-profile-fields#profile-field-types
	DisplayInProfileSummary *bool `json:"display_in_profile_summary,omitempty"`
	// Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.
	//
	// **Changes**: New in Zulip 9.0 (feature level 244).
	Required bool `json:"required"`
	// Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.
	//
	// **Changes**: New in Zulip 10.0 (feature level 296).
	EditableByUser bool `json:"editable_by_user"`
}
