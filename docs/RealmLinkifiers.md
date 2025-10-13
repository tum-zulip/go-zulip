# RealmLinkifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The [Python regular expression](https://docs.python.org/3/howto/regex.html) pattern which represents the pattern that should be linkified on matching.  | [optional] 
**UrlTemplate** | Pointer to **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template with which the pattern matching string should be linkified.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced &#x60;url_format&#x60;, which contained a URL format string.  | [optional] 
**Id** | Pointer to **int32** | The ID of the linkifier.  | [optional] 

## Methods

### NewRegisterQueueResponseRealmLinkifiersItem

`func NewRegisterQueueResponseRealmLinkifiersItem() *RegisterQueueResponseRealmLinkifiersItem`

NewRegisterQueueResponseRealmLinkifiersItem instantiates a new RegisterQueueResponseRealmLinkifiersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseRealmLinkifiersItemWithDefaults

`func NewRegisterQueueResponseRealmLinkifiersItemWithDefaults() *RegisterQueueResponseRealmLinkifiersItem`

NewRegisterQueueResponseRealmLinkifiersItemWithDefaults instantiates a new RegisterQueueResponseRealmLinkifiersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *RegisterQueueResponseRealmLinkifiersItem) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *RegisterQueueResponseRealmLinkifiersItem) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *RegisterQueueResponseRealmLinkifiersItem) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *RegisterQueueResponseRealmLinkifiersItem) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetId

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterQueueResponseRealmLinkifiersItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterQueueResponseRealmLinkifiersItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterQueueResponseRealmLinkifiersItem) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


