# UserUpdateEventEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user affected by this change.  | [optional] 
**NewEmail** | Pointer to **string** | The new value of &#x60;email&#x60; for the user. The client should update any data structures associated with this user to use this new value as the user&#39;s Zulip API email address.  | [optional] 

## Methods

### NewUserUpdateEventEmail

`func NewUserUpdateEventEmail() *UserUpdateEventEmail`

NewUserUpdateEventEmail instantiates a new UserUpdateEventEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventEmailWithDefaults

`func NewUserUpdateEventEmailWithDefaults() *UserUpdateEventEmail`

NewUserUpdateEventEmailWithDefaults instantiates a new UserUpdateEventEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventEmail) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventEmail) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventEmail) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventEmail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNewEmail

`func (o *UserUpdateEventEmail) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *UserUpdateEventEmail) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *UserUpdateEventEmail) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.

### HasNewEmail

`func (o *UserUpdateEventEmail) HasNewEmail() bool`

HasNewEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


