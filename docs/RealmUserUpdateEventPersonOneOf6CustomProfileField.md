# GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the custom profile field which user updated.  | [optional] 
**Value** | Pointer to **NullableString** | User&#39;s personal value for this custom profile field, or &#x60;null&#x60; if unset.  | [optional] 
**RenderedValue** | Pointer to **string** | The &#x60;value&#x60; rendered in HTML. Will only be present for custom profile field types that support Markdown rendering.  This user-generated HTML content should be rendered using the same CSS and client-side security protections as are used for message content.  See [Markdown message formatting](/api/message-formatting) for details on Zulip&#39;s HTML format.  | [optional] 

## Methods

### NewPersonProfileFieldEvent

`func NewPersonProfileFieldEvent() *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField`

NewPersonProfileFieldEvent instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonProfileFieldEventWithDefaults

`func NewPersonProfileFieldEventWithDefaults() *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField`

NewPersonProfileFieldEventWithDefaults instantiates a new GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRenderedValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetRenderedValue() string`

GetRenderedValue returns the RenderedValue field if non-nil, zero value otherwise.

### GetRenderedValueOk

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) GetRenderedValueOk() (*string, bool)`

GetRenderedValueOk returns a tuple with the RenderedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) SetRenderedValue(v string)`

SetRenderedValue sets RenderedValue field to given value.

### HasRenderedValue

`func (o *GetEvents200ResponseAllOfEventsInnerOneOf4PersonOneOf6CustomProfileField) HasRenderedValue() bool`

HasRenderedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


