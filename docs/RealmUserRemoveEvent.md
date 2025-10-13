# RealmUserRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**EventEnvelopeOneOf14Person**](EventEnvelopeOneOf14Person.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf14

`func NewEventEnvelopeOneOf14() *RealmUserRemoveEvent`

NewEventEnvelopeOneOf14 instantiates a new RealmUserRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf14WithDefaults

`func NewEventEnvelopeOneOf14WithDefaults() *RealmUserRemoveEvent`

NewEventEnvelopeOneOf14WithDefaults instantiates a new RealmUserRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUserRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUserRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUserRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUserRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUserRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUserRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUserRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUserRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUserRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUserRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUserRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUserRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPerson

`func (o *RealmUserRemoveEvent) GetPerson() EventEnvelopeOneOf14Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *RealmUserRemoveEvent) GetPersonOk() (*EventEnvelopeOneOf14Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *RealmUserRemoveEvent) SetPerson(v EventEnvelopeOneOf14Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *RealmUserRemoveEvent) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


