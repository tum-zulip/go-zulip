# UserUpdateEventAvatar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user who is affected by this change.  | [optional] 
**AvatarUrl** | Pointer to **string** | The URL of the new avatar for the user.  | [optional] 
**AvatarSource** | Pointer to **string** | The new avatar data source type for the user.  Value values are &#x60;G&#x60; (gravatar) and &#x60;U&#x60; (uploaded by user).  | [optional] 
**AvatarUrlMedium** | Pointer to **string** | The new medium-size avatar URL for user.  | [optional] 
**AvatarVersion** | Pointer to **int32** | The version number for the user&#39;s avatar. This is useful for cache-busting.  | [optional] 

## Methods

### NewUserUpdateEventAvatar

`func NewUserUpdateEventAvatar() *UserUpdateEventAvatar`

NewUserUpdateEventAvatar instantiates a new UserUpdateEventAvatar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventAvatarWithDefaults

`func NewUserUpdateEventAvatarWithDefaults() *UserUpdateEventAvatar`

NewUserUpdateEventAvatarWithDefaults instantiates a new UserUpdateEventAvatar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventAvatar) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventAvatar) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventAvatar) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventAvatar) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserUpdateEventAvatar) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserUpdateEventAvatar) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserUpdateEventAvatar) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserUpdateEventAvatar) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetAvatarSource

`func (o *UserUpdateEventAvatar) GetAvatarSource() string`

GetAvatarSource returns the AvatarSource field if non-nil, zero value otherwise.

### GetAvatarSourceOk

`func (o *UserUpdateEventAvatar) GetAvatarSourceOk() (*string, bool)`

GetAvatarSourceOk returns a tuple with the AvatarSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSource

`func (o *UserUpdateEventAvatar) SetAvatarSource(v string)`

SetAvatarSource sets AvatarSource field to given value.

### HasAvatarSource

`func (o *UserUpdateEventAvatar) HasAvatarSource() bool`

HasAvatarSource returns a boolean if a field has been set.

### GetAvatarUrlMedium

`func (o *UserUpdateEventAvatar) GetAvatarUrlMedium() string`

GetAvatarUrlMedium returns the AvatarUrlMedium field if non-nil, zero value otherwise.

### GetAvatarUrlMediumOk

`func (o *UserUpdateEventAvatar) GetAvatarUrlMediumOk() (*string, bool)`

GetAvatarUrlMediumOk returns a tuple with the AvatarUrlMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrlMedium

`func (o *UserUpdateEventAvatar) SetAvatarUrlMedium(v string)`

SetAvatarUrlMedium sets AvatarUrlMedium field to given value.

### HasAvatarUrlMedium

`func (o *UserUpdateEventAvatar) HasAvatarUrlMedium() bool`

HasAvatarUrlMedium returns a boolean if a field has been set.

### GetAvatarVersion

`func (o *UserUpdateEventAvatar) GetAvatarVersion() int32`

GetAvatarVersion returns the AvatarVersion field if non-nil, zero value otherwise.

### GetAvatarVersionOk

`func (o *UserUpdateEventAvatar) GetAvatarVersionOk() (*int32, bool)`

GetAvatarVersionOk returns a tuple with the AvatarVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarVersion

`func (o *UserUpdateEventAvatar) SetAvatarVersion(v int32)`

SetAvatarVersion sets AvatarVersion field to given value.

### HasAvatarVersion

`func (o *UserUpdateEventAvatar) HasAvatarVersion() bool`

HasAvatarVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


