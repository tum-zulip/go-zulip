# UserUpdateEventRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user affected by this change.  | [optional] 
**Role** | Pointer to **int32** | The new [role](zulip.com/api/roles-and-permissions of the user.  | [optional] 

## Methods

### NewUserUpdateEventRole

`func NewUserUpdateEventRole() *UserUpdateEventRole`

NewUserUpdateEventRole instantiates a new UserUpdateEventRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventRoleWithDefaults

`func NewUserUpdateEventRoleWithDefaults() *UserUpdateEventRole`

NewUserUpdateEventRoleWithDefaults instantiates a new UserUpdateEventRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventRole) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventRole) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventRole) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventRole) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRole

`func (o *UserUpdateEventRole) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserUpdateEventRole) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserUpdateEventRole) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserUpdateEventRole) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


