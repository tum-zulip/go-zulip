# RealmDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The new allowed domain.  | [optional] 
**AllowSubdomains** | Pointer to **bool** | Whether subdomains are allowed for this domain.  | [optional] 

## Methods

### NewRealmDomain

`func NewRealmDomain() *RealmDomain`

NewRealmDomain instantiates a new RealmDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmDomainWithDefaults

`func NewRealmDomainWithDefaults() *RealmDomain`

NewRealmDomainWithDefaults instantiates a new RealmDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RealmDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RealmDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RealmDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RealmDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAllowSubdomains

`func (o *RealmDomain) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *RealmDomain) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *RealmDomain) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.

### HasAllowSubdomains

`func (o *RealmDomain) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


