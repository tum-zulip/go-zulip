# UserUpdateEventFullName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of modified user.  | [optional] 
**FullName** | Pointer to **string** | The new full name for the user.  | [optional] 

## Methods

### NewUserUpdateEventFullName

`func NewUserUpdateEventFullName() *UserUpdateEventFullName`

NewUserUpdateEventFullName instantiates a new UserUpdateEventFullName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventFullNameWithDefaults

`func NewUserUpdateEventFullNameWithDefaults() *UserUpdateEventFullName`

NewUserUpdateEventFullNameWithDefaults instantiates a new UserUpdateEventFullName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventFullName) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventFullName) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventFullName) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventFullName) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFullName

`func (o *UserUpdateEventFullName) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserUpdateEventFullName) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserUpdateEventFullName) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserUpdateEventFullName) HasFullName() bool`

HasFullName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


