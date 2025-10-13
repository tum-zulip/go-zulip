# NavigationViewAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**NavigationView** | Pointer to [**NavigationView**](NavigationView.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf72

`func NewGetEvents200ResponseAllOfEventsInnerOneOf72() *NavigationViewAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf72 instantiates a new NavigationViewAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf72WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf72WithDefaults() *NavigationViewAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf72WithDefaults instantiates a new NavigationViewAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NavigationViewAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NavigationViewAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NavigationViewAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NavigationViewAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NavigationViewAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NavigationViewAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NavigationViewAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NavigationViewAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *NavigationViewAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *NavigationViewAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *NavigationViewAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *NavigationViewAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetNavigationView

`func (o *NavigationViewAddEvent) GetNavigationView() NavigationView`

GetNavigationView returns the NavigationView field if non-nil, zero value otherwise.

### GetNavigationViewOk

`func (o *NavigationViewAddEvent) GetNavigationViewOk() (*NavigationView, bool)`

GetNavigationViewOk returns a tuple with the NavigationView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationView

`func (o *NavigationViewAddEvent) SetNavigationView(v NavigationView)`

SetNavigationView sets NavigationView field to given value.

### HasNavigationView

`func (o *NavigationViewAddEvent) HasNavigationView() bool`

HasNavigationView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


