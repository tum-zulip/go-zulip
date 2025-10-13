# RealmDomainsAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**RealmDomain** | Pointer to [**RealmDomain**](RealmDomain.md) |  | [optional] 

## Methods

### NewEventEnvelopeOneOf54

`func NewEventEnvelopeOneOf54() *RealmDomainsAddEvent`

NewEventEnvelopeOneOf54 instantiates a new RealmDomainsAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf54WithDefaults

`func NewEventEnvelopeOneOf54WithDefaults() *RealmDomainsAddEvent`

NewEventEnvelopeOneOf54WithDefaults instantiates a new RealmDomainsAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmDomainsAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmDomainsAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmDomainsAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmDomainsAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmDomainsAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmDomainsAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmDomainsAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmDomainsAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmDomainsAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmDomainsAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmDomainsAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmDomainsAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRealmDomain

`func (o *RealmDomainsAddEvent) GetRealmDomain() RealmDomain`

GetRealmDomain returns the RealmDomain field if non-nil, zero value otherwise.

### GetRealmDomainOk

`func (o *RealmDomainsAddEvent) GetRealmDomainOk() (*RealmDomain, bool)`

GetRealmDomainOk returns a tuple with the RealmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDomain

`func (o *RealmDomainsAddEvent) SetRealmDomain(v RealmDomain)`

SetRealmDomain sets RealmDomain field to given value.

### HasRealmDomain

`func (o *RealmDomainsAddEvent) HasRealmDomain() bool`

HasRealmDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


