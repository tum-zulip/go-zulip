# UserSettingsUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | Name of the changed setting.  | [optional] 
**Value** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf1Setting**](GetEvents200ResponseAllOfEventsInnerOneOf1Setting.md) |  | [optional] 
**LanguageName** | Pointer to **string** | Present only if the setting to be changed is &#x60;default_language&#x60;. Contains the name of the new default language in English.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf3

`func NewGetEvents200ResponseAllOfEventsInnerOneOf3() *UserSettingsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf3 instantiates a new UserSettingsUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf3WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf3WithDefaults() *UserSettingsUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf3WithDefaults instantiates a new UserSettingsUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSettingsUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettingsUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettingsUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserSettingsUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserSettingsUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSettingsUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSettingsUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSettingsUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *UserSettingsUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UserSettingsUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UserSettingsUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UserSettingsUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *UserSettingsUpdateEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *UserSettingsUpdateEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *UserSettingsUpdateEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *UserSettingsUpdateEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *UserSettingsUpdateEvent) GetValue() GetEvents200ResponseAllOfEventsInnerOneOf1Setting`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserSettingsUpdateEvent) GetValueOk() (*GetEvents200ResponseAllOfEventsInnerOneOf1Setting, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserSettingsUpdateEvent) SetValue(v GetEvents200ResponseAllOfEventsInnerOneOf1Setting)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserSettingsUpdateEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLanguageName

`func (o *UserSettingsUpdateEvent) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *UserSettingsUpdateEvent) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *UserSettingsUpdateEvent) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *UserSettingsUpdateEvent) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


