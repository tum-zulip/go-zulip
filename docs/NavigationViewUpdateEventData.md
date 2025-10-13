# NavigationViewUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Fragment** | Pointer to **string** | The unique URL hash of the navigation view being updated.  | [optional] 
**Data** | Pointer to [**EventEnvelopeOneOf73Data**](EventEnvelopeOneOf73Data.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf73

`func NewEventEnvelopeOneOf73() *NavigationViewUpdateEvent`

NewEventEnvelopeOneOf73 instantiates a new NavigationViewUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf73WithDefaults

`func NewEventEnvelopeOneOf73WithDefaults() *NavigationViewUpdateEvent`

NewEventEnvelopeOneOf73WithDefaults instantiates a new NavigationViewUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NavigationViewUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NavigationViewUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NavigationViewUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NavigationViewUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NavigationViewUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NavigationViewUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NavigationViewUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NavigationViewUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *NavigationViewUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *NavigationViewUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *NavigationViewUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *NavigationViewUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetFragment

`func (o *NavigationViewUpdateEvent) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *NavigationViewUpdateEvent) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *NavigationViewUpdateEvent) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *NavigationViewUpdateEvent) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetData

`func (o *NavigationViewUpdateEvent) GetData() EventEnvelopeOneOf73Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NavigationViewUpdateEvent) GetDataOk() (*EventEnvelopeOneOf73Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NavigationViewUpdateEvent) SetData(v EventEnvelopeOneOf73Data)`

SetData sets Data field to given value.

### HasData

`func (o *NavigationViewUpdateEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


