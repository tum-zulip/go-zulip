# RealmDomainsChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**RealmDomain** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain**](GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf55

`func NewGetEvents200ResponseAllOfEventsInnerOneOf55() *RealmDomainsChangeEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf55 instantiates a new RealmDomainsChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf55WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf55WithDefaults() *RealmDomainsChangeEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf55WithDefaults instantiates a new RealmDomainsChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmDomainsChangeEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmDomainsChangeEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmDomainsChangeEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmDomainsChangeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmDomainsChangeEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmDomainsChangeEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmDomainsChangeEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmDomainsChangeEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmDomainsChangeEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmDomainsChangeEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmDomainsChangeEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmDomainsChangeEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRealmDomain

`func (o *RealmDomainsChangeEvent) GetRealmDomain() GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain`

GetRealmDomain returns the RealmDomain field if non-nil, zero value otherwise.

### GetRealmDomainOk

`func (o *RealmDomainsChangeEvent) GetRealmDomainOk() (*GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain, bool)`

GetRealmDomainOk returns a tuple with the RealmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDomain

`func (o *RealmDomainsChangeEvent) SetRealmDomain(v GetEvents200ResponseAllOfEventsInnerOneOf55RealmDomain)`

SetRealmDomain sets RealmDomain field to given value.

### HasRealmDomain

`func (o *RealmDomainsChangeEvent) HasRealmDomain() bool`

HasRealmDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


