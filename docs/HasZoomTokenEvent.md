# HasZoomTokenEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **bool** | A boolean specifying whether the user has zoom token or not.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf11

`func NewGetEvents200ResponseAllOfEventsInnerOneOf11() *HasZoomTokenEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf11 instantiates a new HasZoomTokenEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf11WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf11WithDefaults() *HasZoomTokenEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf11WithDefaults instantiates a new HasZoomTokenEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HasZoomTokenEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HasZoomTokenEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HasZoomTokenEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HasZoomTokenEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *HasZoomTokenEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HasZoomTokenEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HasZoomTokenEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HasZoomTokenEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *HasZoomTokenEvent) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HasZoomTokenEvent) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HasZoomTokenEvent) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *HasZoomTokenEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


