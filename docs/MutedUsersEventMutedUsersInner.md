# MutedUsersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MutedUsers** | Pointer to [**[]EventEnvelopeOneOf33MutedUsersInner**](EventEnvelopeOneOf33MutedUsersInner.md) | A list of dictionaries where each dictionary describes a muted user.  | [optional] 

## Methods

### NewEventEnvelopeOneOf33

`func NewEventEnvelopeOneOf33() *MutedUsersEvent`

NewEventEnvelopeOneOf33 instantiates a new MutedUsersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf33WithDefaults

`func NewEventEnvelopeOneOf33WithDefaults() *MutedUsersEvent`

NewEventEnvelopeOneOf33WithDefaults instantiates a new MutedUsersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MutedUsersEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutedUsersEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutedUsersEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MutedUsersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MutedUsersEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MutedUsersEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MutedUsersEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MutedUsersEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMutedUsers

`func (o *MutedUsersEvent) GetMutedUsers() []EventEnvelopeOneOf33MutedUsersInner`

GetMutedUsers returns the MutedUsers field if non-nil, zero value otherwise.

### GetMutedUsersOk

`func (o *MutedUsersEvent) GetMutedUsersOk() (*[]EventEnvelopeOneOf33MutedUsersInner, bool)`

GetMutedUsersOk returns a tuple with the MutedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUsers

`func (o *MutedUsersEvent) SetMutedUsers(v []EventEnvelopeOneOf33MutedUsersInner)`

SetMutedUsers sets MutedUsers field to given value.

### HasMutedUsers

`func (o *MutedUsersEvent) HasMutedUsers() bool`

HasMutedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


