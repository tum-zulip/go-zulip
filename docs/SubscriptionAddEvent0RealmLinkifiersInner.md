# GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The [Python regular expression](https://docs.python.org/3/howto/regex.html) that represents the pattern that should be linkified by this linkifier.  | [optional] 
**UrlTemplate** | Pointer to **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template to be used for linkifying matches.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced &#x60;url_format&#x60;, which contained a URL format string.  | [optional] 
**Id** | Pointer to **int32** | The ID of the linkifier.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner

`func NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner() *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner`

NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInnerWithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInnerWithDefaults() *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner`

NewGetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInnerWithDefaults instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf50RealmLinkifiersInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


