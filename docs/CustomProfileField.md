# CustomProfileField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the custom profile field. This will be referenced in the custom profile fields section of user objects.  | 
**Type** | **int32** | An integer indicating the type of the custom profile field, which determines how it is configured and displayed to users.  See the [Custom profile fields](/help/custom-profile-fields#profile-field-types) article for details on what each type means.  - **1**: Short text - **2**: Long text - **3**: List of options - **4**: Date picker - **5**: Link - **6**: Person picker - **7**: External account - **8**: Pronouns  **Changes**: Field type &#x60;8&#x60; added in Zulip 6.0 (feature level 151).  | 
**Order** | **int32** | Custom profile fields are displayed in both settings UI and UI showing users&#39; profiles in increasing &#x60;order&#x60;.  | 
**Name** | **string** | The name of the custom profile field.  | 
**Hint** | **string** | The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.  | 
**FieldData** | Pointer to **string** | Field types 3 (List of options) and 7 (External account) support storing additional configuration for the field type in the &#x60;field_data&#x60; attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type 7 is not yet stabilized.  | [optional] 
**DisplayInProfileSummary** | Pointer to **bool** | Whether the custom profile field, display or not on the user card.  Currently it&#39;s value not allowed to be &#x60;true&#x60; of &#x60;Long text&#x60; and &#x60;Person picker&#x60; [profile field types](/help/custom-profile-fields#profile-field-types).  This field is only included when its value is &#x60;true&#x60;.  **Changes**: New in Zulip 6.0 (feature level 146).  | [optional] [default to false]
**Required** | **bool** | Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.  **Changes**: New in Zulip 9.0 (feature level 244).  | 
**EditableByUser** | **bool** | Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.  **Changes**: New in Zulip 10.0 (feature level 296).  | [default to true]

## Methods

### NewCustomProfileField

`func NewCustomProfileField(id int32, type_ int32, order int32, name string, hint string, required bool, editableByUser bool, ) *CustomProfileField`

NewCustomProfileField instantiates a new CustomProfileField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomProfileFieldWithDefaults

`func NewCustomProfileFieldWithDefaults() *CustomProfileField`

NewCustomProfileFieldWithDefaults instantiates a new CustomProfileField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomProfileField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomProfileField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomProfileField) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *CustomProfileField) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomProfileField) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomProfileField) SetType(v int32)`

SetType sets Type field to given value.


### GetOrder

`func (o *CustomProfileField) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CustomProfileField) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CustomProfileField) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetName

`func (o *CustomProfileField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomProfileField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomProfileField) SetName(v string)`

SetName sets Name field to given value.


### GetHint

`func (o *CustomProfileField) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CustomProfileField) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CustomProfileField) SetHint(v string)`

SetHint sets Hint field to given value.


### GetFieldData

`func (o *CustomProfileField) GetFieldData() string`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *CustomProfileField) GetFieldDataOk() (*string, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *CustomProfileField) SetFieldData(v string)`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *CustomProfileField) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### GetDisplayInProfileSummary

`func (o *CustomProfileField) GetDisplayInProfileSummary() bool`

GetDisplayInProfileSummary returns the DisplayInProfileSummary field if non-nil, zero value otherwise.

### GetDisplayInProfileSummaryOk

`func (o *CustomProfileField) GetDisplayInProfileSummaryOk() (*bool, bool)`

GetDisplayInProfileSummaryOk returns a tuple with the DisplayInProfileSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInProfileSummary

`func (o *CustomProfileField) SetDisplayInProfileSummary(v bool)`

SetDisplayInProfileSummary sets DisplayInProfileSummary field to given value.

### HasDisplayInProfileSummary

`func (o *CustomProfileField) HasDisplayInProfileSummary() bool`

HasDisplayInProfileSummary returns a boolean if a field has been set.

### GetRequired

`func (o *CustomProfileField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomProfileField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomProfileField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetEditableByUser

`func (o *CustomProfileField) GetEditableByUser() bool`

GetEditableByUser returns the EditableByUser field if non-nil, zero value otherwise.

### GetEditableByUserOk

`func (o *CustomProfileField) GetEditableByUserOk() (*bool, bool)`

GetEditableByUserOk returns a tuple with the EditableByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByUser

`func (o *CustomProfileField) SetEditableByUser(v bool)`

SetEditableByUser sets EditableByUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


