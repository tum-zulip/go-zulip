# NarrowMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchContent** | Pointer to **string** | HTML content of a queried message that matches the narrow. If the narrow is a search narrow, &#x60;&lt;span class&#x3D;\&quot;highlight\&quot;&gt;&#x60; elements will be included, wrapping the matches for the search keywords.  | [optional] 
**MatchSubject** | Pointer to **string** | HTML-escaped topic of a queried message that matches the narrow. If the narrow is a search narrow, &#x60;&lt;span class&#x3D;\&quot;highlight\&quot;&gt;&#x60; elements will be included wrapping the matches for the search keywords.  | [optional] 

## Methods

### NewNarrowMatch

`func NewNarrowMatch() *NarrowMatch`

NewNarrowMatch instantiates a new NarrowMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNarrowMatchWithDefaults

`func NewNarrowMatchWithDefaults() *NarrowMatch`

NewNarrowMatchWithDefaults instantiates a new NarrowMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchContent

`func (o *NarrowMatch) GetMatchContent() string`

GetMatchContent returns the MatchContent field if non-nil, zero value otherwise.

### GetMatchContentOk

`func (o *NarrowMatch) GetMatchContentOk() (*string, bool)`

GetMatchContentOk returns a tuple with the MatchContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchContent

`func (o *NarrowMatch) SetMatchContent(v string)`

SetMatchContent sets MatchContent field to given value.

### HasMatchContent

`func (o *NarrowMatch) HasMatchContent() bool`

HasMatchContent returns a boolean if a field has been set.

### GetMatchSubject

`func (o *NarrowMatch) GetMatchSubject() string`

GetMatchSubject returns the MatchSubject field if non-nil, zero value otherwise.

### GetMatchSubjectOk

`func (o *NarrowMatch) GetMatchSubjectOk() (*string, bool)`

GetMatchSubjectOk returns a tuple with the MatchSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSubject

`func (o *NarrowMatch) SetMatchSubject(v string)`

SetMatchSubject sets MatchSubject field to given value.

### HasMatchSubject

`func (o *NarrowMatch) HasMatchSubject() bool`

HasMatchSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


