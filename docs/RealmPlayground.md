# RealmPlayground

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for the realm playground.  | [optional] 
**Name** | Pointer to **string** | The user-visible display name of the playground. Clients should display this in UI for picking which playground to open a code block in, to differentiate between multiple configured playground options for a given pygments language.  **Changes**: New in Zulip 4.0 (feature level 49).  | [optional] 
**PygmentsLanguage** | Pointer to **string** | The name of the Pygments language lexer for that programming language.  | [optional] 
**UrlTemplate** | Pointer to **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template for the playground. The template contains exactly one variable named &#x60;code&#x60;, which determines how the extracted code should be substituted in the playground URL.  **Changes**: New in Zulip 8.0 (feature level 196). This replaced the &#x60;url_prefix&#x60; parameter, which was used to construct URLs by just concatenating url_prefix and code.  | [optional] 

## Methods

### NewRealmPlayground

`func NewRealmPlayground() *RealmPlayground`

NewRealmPlayground instantiates a new RealmPlayground object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmPlaygroundWithDefaults

`func NewRealmPlaygroundWithDefaults() *RealmPlayground`

NewRealmPlaygroundWithDefaults instantiates a new RealmPlayground object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmPlayground) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmPlayground) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmPlayground) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmPlayground) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RealmPlayground) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmPlayground) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmPlayground) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealmPlayground) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPygmentsLanguage

`func (o *RealmPlayground) GetPygmentsLanguage() string`

GetPygmentsLanguage returns the PygmentsLanguage field if non-nil, zero value otherwise.

### GetPygmentsLanguageOk

`func (o *RealmPlayground) GetPygmentsLanguageOk() (*string, bool)`

GetPygmentsLanguageOk returns a tuple with the PygmentsLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPygmentsLanguage

`func (o *RealmPlayground) SetPygmentsLanguage(v string)`

SetPygmentsLanguage sets PygmentsLanguage field to given value.

### HasPygmentsLanguage

`func (o *RealmPlayground) HasPygmentsLanguage() bool`

HasPygmentsLanguage returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *RealmPlayground) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *RealmPlayground) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *RealmPlayground) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *RealmPlayground) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


