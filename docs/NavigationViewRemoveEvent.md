# NavigationViewRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Fragment** | Pointer to **string** | The unique URL hash of the navigation view that was just deleted.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf74

`func NewGetEvents200ResponseAllOfEventsInnerOneOf74() *NavigationViewRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf74 instantiates a new NavigationViewRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf74WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf74WithDefaults() *NavigationViewRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf74WithDefaults instantiates a new NavigationViewRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NavigationViewRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NavigationViewRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NavigationViewRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NavigationViewRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NavigationViewRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NavigationViewRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NavigationViewRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NavigationViewRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *NavigationViewRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *NavigationViewRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *NavigationViewRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *NavigationViewRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetFragment

`func (o *NavigationViewRemoveEvent) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *NavigationViewRemoveEvent) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *NavigationViewRemoveEvent) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *NavigationViewRemoveEvent) HasFragment() bool`

HasFragment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


