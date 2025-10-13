# UserUpdateEventBotOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The ID of the user/bot whose owner has changed.  | [optional] 
**BotOwnerId** | Pointer to **int32** | The user ID of the new bot owner.  | [optional] 

## Methods

### NewUserUpdateEventBotOwner

`func NewUserUpdateEventBotOwner() *UserUpdateEventBotOwner`

NewUserUpdateEventBotOwner instantiates a new UserUpdateEventBotOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateEventBotOwnerWithDefaults

`func NewUserUpdateEventBotOwnerWithDefaults() *UserUpdateEventBotOwner`

NewUserUpdateEventBotOwnerWithDefaults instantiates a new UserUpdateEventBotOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserUpdateEventBotOwner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserUpdateEventBotOwner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserUpdateEventBotOwner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserUpdateEventBotOwner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetBotOwnerId

`func (o *UserUpdateEventBotOwner) GetBotOwnerId() int32`

GetBotOwnerId returns the BotOwnerId field if non-nil, zero value otherwise.

### GetBotOwnerIdOk

`func (o *UserUpdateEventBotOwner) GetBotOwnerIdOk() (*int32, bool)`

GetBotOwnerIdOk returns a tuple with the BotOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotOwnerId

`func (o *UserUpdateEventBotOwner) SetBotOwnerId(v int32)`

SetBotOwnerId sets BotOwnerId field to given value.

### HasBotOwnerId

`func (o *UserUpdateEventBotOwner) HasBotOwnerId() bool`

HasBotOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


