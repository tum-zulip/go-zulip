# ExternalAuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique, table, machine-readable name for the authentication method, intended to be used by clients with special behavior for specific authentication methods to correctly identify the method.  | [optional] 
**DisplayName** | Pointer to **string** | Display name of the authentication method, to be used in all buttons for the authentication method.  | [optional] 
**DisplayIcon** | Pointer to **NullableString** | URL for an image to be displayed as an icon in all buttons for the external authentication method.  When &#x60;null&#x60;, no icon should be displayed.  | [optional] 
**LoginUrl** | Pointer to **string** | URL to be used to initiate authentication using this method.  | [optional] 
**SignupUrl** | Pointer to **string** | URL to be used to initiate account registration using this method.  | [optional] 

## Methods

### NewExternalAuthMethod

`func NewExternalAuthMethod() *ExternalAuthMethod`

NewExternalAuthMethod instantiates a new ExternalAuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAuthMethodWithDefaults

`func NewExternalAuthMethodWithDefaults() *ExternalAuthMethod`

NewExternalAuthMethodWithDefaults instantiates a new ExternalAuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalAuthMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalAuthMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalAuthMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalAuthMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ExternalAuthMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ExternalAuthMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ExternalAuthMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ExternalAuthMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayIcon

`func (o *ExternalAuthMethod) GetDisplayIcon() string`

GetDisplayIcon returns the DisplayIcon field if non-nil, zero value otherwise.

### GetDisplayIconOk

`func (o *ExternalAuthMethod) GetDisplayIconOk() (*string, bool)`

GetDisplayIconOk returns a tuple with the DisplayIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIcon

`func (o *ExternalAuthMethod) SetDisplayIcon(v string)`

SetDisplayIcon sets DisplayIcon field to given value.

### HasDisplayIcon

`func (o *ExternalAuthMethod) HasDisplayIcon() bool`

HasDisplayIcon returns a boolean if a field has been set.

### SetDisplayIconNil

`func (o *ExternalAuthMethod) SetDisplayIconNil(b bool)`

 SetDisplayIconNil sets the value for DisplayIcon to be an explicit nil

### UnsetDisplayIcon
`func (o *ExternalAuthMethod) UnsetDisplayIcon()`

UnsetDisplayIcon ensures that no value is present for DisplayIcon, not even an explicit nil
### GetLoginUrl

`func (o *ExternalAuthMethod) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *ExternalAuthMethod) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *ExternalAuthMethod) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *ExternalAuthMethod) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetSignupUrl

`func (o *ExternalAuthMethod) GetSignupUrl() string`

GetSignupUrl returns the SignupUrl field if non-nil, zero value otherwise.

### GetSignupUrlOk

`func (o *ExternalAuthMethod) GetSignupUrlOk() (*string, bool)`

GetSignupUrlOk returns a tuple with the SignupUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUrl

`func (o *ExternalAuthMethod) SetSignupUrl(v string)`

SetSignupUrl sets SignupUrl field to given value.

### HasSignupUrl

`func (o *ExternalAuthMethod) HasSignupUrl() bool`

HasSignupUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


