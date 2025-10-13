# CheckMessagesMatchNarrow200ResponseAllOfMessagesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchContent** | Pointer to **string** | HTML content of a queried message that matches the narrow. If the narrow is a search narrow, &#x60;&lt;span class&#x3D;\&quot;highlight\&quot;&gt;&#x60; elements will be included, wrapping the matches for the search keywords.  | [optional] 
**MatchSubject** | Pointer to **string** | HTML-escaped topic of a queried message that matches the narrow. If the narrow is a search narrow, &#x60;&lt;span class&#x3D;\&quot;highlight\&quot;&gt;&#x60; elements will be included wrapping the matches for the search keywords.  | [optional] 

## Methods

### NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValue

`func NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValue() *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue`

NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValue instantiates a new CheckMessagesMatchNarrow200ResponseAllOfMessagesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValueWithDefaults

`func NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValueWithDefaults() *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue`

NewCheckMessagesMatchNarrow200ResponseAllOfMessagesValueWithDefaults instantiates a new CheckMessagesMatchNarrow200ResponseAllOfMessagesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchContent

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) GetMatchContent() string`

GetMatchContent returns the MatchContent field if non-nil, zero value otherwise.

### GetMatchContentOk

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) GetMatchContentOk() (*string, bool)`

GetMatchContentOk returns a tuple with the MatchContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchContent

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) SetMatchContent(v string)`

SetMatchContent sets MatchContent field to given value.

### HasMatchContent

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) HasMatchContent() bool`

HasMatchContent returns a boolean if a field has been set.

### GetMatchSubject

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) GetMatchSubject() string`

GetMatchSubject returns the MatchSubject field if non-nil, zero value otherwise.

### GetMatchSubjectOk

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) GetMatchSubjectOk() (*string, bool)`

GetMatchSubjectOk returns a tuple with the MatchSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSubject

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) SetMatchSubject(v string)`

SetMatchSubject sets MatchSubject field to given value.

### HasMatchSubject

`func (o *CheckMessagesMatchNarrow200ResponseAllOfMessagesValue) HasMatchSubject() bool`

HasMatchSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


