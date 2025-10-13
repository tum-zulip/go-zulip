# UpdateDisplaySettingsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** | The Zulip API email of the user.  | [optional] 
**SettingName** | Pointer to **string** | Name of the changed display setting.  | [optional] 
**Setting** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf1Setting**](GetEvents200ResponseAllOfEventsInnerOneOf1Setting.md) |  | [optional] 
**LanguageName** | Pointer to **string** | Present only if the setting to be changed is &#x60;default_language&#x60;. Contains the name of the new default language in English.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf1

`func NewGetEvents200ResponseAllOfEventsInnerOneOf1() *UpdateDisplaySettingsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf1 instantiates a new UpdateDisplaySettingsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf1WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf1WithDefaults() *UpdateDisplaySettingsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf1WithDefaults instantiates a new UpdateDisplaySettingsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateDisplaySettingsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDisplaySettingsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDisplaySettingsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDisplaySettingsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateDisplaySettingsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDisplaySettingsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDisplaySettingsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateDisplaySettingsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *UpdateDisplaySettingsEvent) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateDisplaySettingsEvent) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateDisplaySettingsEvent) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateDisplaySettingsEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSettingName

`func (o *UpdateDisplaySettingsEvent) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *UpdateDisplaySettingsEvent) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *UpdateDisplaySettingsEvent) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *UpdateDisplaySettingsEvent) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSetting

`func (o *UpdateDisplaySettingsEvent) GetSetting() GetEvents200ResponseAllOfEventsInnerOneOf1Setting`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UpdateDisplaySettingsEvent) GetSettingOk() (*GetEvents200ResponseAllOfEventsInnerOneOf1Setting, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UpdateDisplaySettingsEvent) SetSetting(v GetEvents200ResponseAllOfEventsInnerOneOf1Setting)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *UpdateDisplaySettingsEvent) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetLanguageName

`func (o *UpdateDisplaySettingsEvent) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *UpdateDisplaySettingsEvent) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *UpdateDisplaySettingsEvent) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *UpdateDisplaySettingsEvent) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


