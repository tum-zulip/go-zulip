# RealmUserAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf13

`func NewEventEnvelopeOneOf13() *RealmUserAddEvent`

NewEventEnvelopeOneOf13 instantiates a new RealmUserAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf13WithDefaults

`func NewEventEnvelopeOneOf13WithDefaults() *RealmUserAddEvent`

NewEventEnvelopeOneOf13WithDefaults instantiates a new RealmUserAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUserAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUserAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUserAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUserAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUserAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUserAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUserAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUserAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUserAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUserAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUserAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUserAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPerson

`func (o *RealmUserAddEvent) GetPerson() User`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *RealmUserAddEvent) GetPersonOk() (*User, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *RealmUserAddEvent) SetPerson(v User)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *RealmUserAddEvent) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


