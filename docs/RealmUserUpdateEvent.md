# RealmUserUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf4Person**](GetEvents200ResponseAllOfEventsInnerOneOf4Person.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf4

`func NewGetEvents200ResponseAllOfEventsInnerOneOf4() *RealmUserUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf4 instantiates a new RealmUserUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf4WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf4WithDefaults() *RealmUserUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf4WithDefaults instantiates a new RealmUserUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmUserUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmUserUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmUserUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmUserUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmUserUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmUserUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmUserUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmUserUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *RealmUserUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RealmUserUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RealmUserUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RealmUserUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPerson

`func (o *RealmUserUpdateEvent) GetPerson() GetEvents200ResponseAllOfEventsInnerOneOf4Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *RealmUserUpdateEvent) GetPersonOk() (*GetEvents200ResponseAllOfEventsInnerOneOf4Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *RealmUserUpdateEvent) SetPerson(v GetEvents200ResponseAllOfEventsInnerOneOf4Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *RealmUserUpdateEvent) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


