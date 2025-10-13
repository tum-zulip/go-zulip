# RealmDeactivatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**RealmId** | Pointer to **int32** | The ID of the deactivated realm.  | [optional] 

## Methods

### NewEventEnvelopeOneOf64

`func NewEventEnvelopeOneOf64() *RealmDeactivatedEvent`

NewEventEnvelopeOneOf64 instantiates a new RealmDeactivatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf64WithDefaults

`func NewEventEnvelopeOneOf64WithDefaults() *RealmDeactivatedEvent`

NewEventEnvelopeOneOf64WithDefaults instantiates a new RealmDeactivatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmDeactivatedEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmDeactivatedEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmDeactivatedEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmDeactivatedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmDeactivatedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmDeactivatedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmDeactivatedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmDeactivatedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmDeactivatedEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmDeactivatedEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmDeactivatedEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmDeactivatedEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRealmId

`func (o *RealmDeactivatedEvent) GetRealmId() int32`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *RealmDeactivatedEvent) GetRealmIdOk() (*int32, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *RealmDeactivatedEvent) SetRealmId(v int32)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *RealmDeactivatedEvent) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


