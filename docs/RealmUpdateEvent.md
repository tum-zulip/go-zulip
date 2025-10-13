# RealmUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Property** | Pointer to **string** | The name of the property that was changed.  | [optional] 
**Value** | Pointer to [**EventEnvelopeOneOf63Value**](EventEnvelopeOneOf63Value.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf63

`func NewEventEnvelopeOneOf63() *RealmUpdateEvent`

NewEventEnvelopeOneOf63 instantiates a new RealmUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf63WithDefaults

`func NewEventEnvelopeOneOf63WithDefaults() *RealmUpdateEvent`

NewEventEnvelopeOneOf63WithDefaults instantiates a new RealmUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetProperty

`func (o *RealmUpdateEvent) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RealmUpdateEvent) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RealmUpdateEvent) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RealmUpdateEvent) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *RealmUpdateEvent) GetValue() EventEnvelopeOneOf63Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RealmUpdateEvent) GetValueOk() (*EventEnvelopeOneOf63Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RealmUpdateEvent) SetValue(v EventEnvelopeOneOf63Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *RealmUpdateEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


