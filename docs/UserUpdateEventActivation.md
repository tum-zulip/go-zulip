# UserUpdateEventActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user affected by this change.  | [optional] 
**IsActive** | Pointer to **bool** | A boolean describing whether the user account has been deactivated.  | [optional] 

## Methods

### NewUserUpdateEventActivation

`func NewUserUpdateEventActivation() *UserUpdateEventActivation`

NewUserUpdateEventActivation instantiates a new UserUpdateEventActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventActivationWithDefaults

`func NewUserUpdateEventActivationWithDefaults() *UserUpdateEventActivation`

NewUserUpdateEventActivationWithDefaults instantiates a new UserUpdateEventActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventActivation) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventActivation) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventActivation) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventActivation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsActive

`func (o *UserUpdateEventActivation) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserUpdateEventActivation) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserUpdateEventActivation) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserUpdateEventActivation) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


