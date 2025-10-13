# RealmDomainsRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | The domain to be removed.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf56

`func NewGetEvents200ResponseAllOfEventsInnerOneOf56() *RealmDomainsRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf56 instantiates a new RealmDomainsRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf56WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf56WithDefaults() *RealmDomainsRemoveEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf56WithDefaults instantiates a new RealmDomainsRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmDomainsRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmDomainsRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmDomainsRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmDomainsRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmDomainsRemoveEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmDomainsRemoveEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmDomainsRemoveEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmDomainsRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmDomainsRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmDomainsRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmDomainsRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmDomainsRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetDomain

`func (o *RealmDomainsRemoveEvent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RealmDomainsRemoveEvent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RealmDomainsRemoveEvent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RealmDomainsRemoveEvent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


