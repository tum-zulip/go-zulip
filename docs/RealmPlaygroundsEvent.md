# RealmPlaygroundsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RealmPlaygrounds** | Pointer to [**[]RealmPlayground**](RealmPlayground.md) | An array of dictionaries where each dictionary contains data about a single playground entry.  | [optional] 

## Methods

### NewEventEnvelopeOneOf52

`func NewEventEnvelopeOneOf52() *RealmPlaygroundsEvent`

NewEventEnvelopeOneOf52 instantiates a new RealmPlaygroundsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf52WithDefaults

`func NewEventEnvelopeOneOf52WithDefaults() *RealmPlaygroundsEvent`

NewEventEnvelopeOneOf52WithDefaults instantiates a new RealmPlaygroundsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmPlaygroundsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmPlaygroundsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmPlaygroundsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmPlaygroundsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmPlaygroundsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmPlaygroundsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmPlaygroundsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmPlaygroundsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRealmPlaygrounds

`func (o *RealmPlaygroundsEvent) GetRealmPlaygrounds() []RealmPlayground`

GetRealmPlaygrounds returns the RealmPlaygrounds field if non-nil, zero value otherwise.

### GetRealmPlaygroundsOk

`func (o *RealmPlaygroundsEvent) GetRealmPlaygroundsOk() (*[]RealmPlayground, bool)`

GetRealmPlaygroundsOk returns a tuple with the RealmPlaygrounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPlaygrounds

`func (o *RealmPlaygroundsEvent) SetRealmPlaygrounds(v []RealmPlayground)`

SetRealmPlaygrounds sets RealmPlaygrounds field to given value.

### HasRealmPlaygrounds

`func (o *RealmPlaygroundsEvent) HasRealmPlaygrounds() bool`

HasRealmPlaygrounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


