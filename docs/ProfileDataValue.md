# ProfileDataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | User&#39;s personal value for this custom profile field.  | [optional] 
**RenderedValue** | Pointer to **string** | The &#x60;value&#x60; rendered in HTML. Will only be present for custom profile field types that support Markdown rendering.  This user-generated HTML content should be rendered using the same CSS and client-side security protections as are used for message content.  See [Markdown message formatting](zulip.com/api/message-formatting for details on Zulip&#39;s HTML format.  | [optional] 

## Methods

### NewProfileDataValue

`func NewProfileDataValue() *ProfileDataValue`

NewProfileDataValue instantiates a new ProfileDataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileDataValueWithDefaults

`func NewProfileDataValueWithDefaults() *ProfileDataValue`

NewProfileDataValueWithDefaults instantiates a new ProfileDataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProfileDataValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileDataValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileDataValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProfileDataValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRenderedValue

`func (o *ProfileDataValue) GetRenderedValue() string`

GetRenderedValue returns the RenderedValue field if non-nil, zero value otherwise.

### GetRenderedValueOk

`func (o *ProfileDataValue) GetRenderedValueOk() (*string, bool)`

GetRenderedValueOk returns a tuple with the RenderedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedValue

`func (o *ProfileDataValue) SetRenderedValue(v string)`

SetRenderedValue sets RenderedValue field to given value.

### HasRenderedValue

`func (o *ProfileDataValue) HasRenderedValue() bool`

HasRenderedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


