# EmojiReactionEventAllOfUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | ID of the user.  | [optional] 
**Email** | Pointer to **string** | Zulip API email of the user.  | [optional] 
**FullName** | Pointer to **string** | Full name of the user.  | [optional] 
**IsMirrorDummy** | Pointer to **bool** | Whether the user is a mirror dummy.  | [optional] 

## Methods

### NewEmojiReactionEventAllOfUser

`func NewEmojiReactionEventAllOfUser() *EmojiReactionEventAllOfUser`

NewEmojiReactionEventAllOfUser instantiates a new EmojiReactionEventAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiReactionEventAllOfUserWithDefaults

`func NewEmojiReactionEventAllOfUserWithDefaults() *EmojiReactionEventAllOfUser`

NewEmojiReactionEventAllOfUserWithDefaults instantiates a new EmojiReactionEventAllOfUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EmojiReactionEventAllOfUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmojiReactionEventAllOfUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmojiReactionEventAllOfUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EmojiReactionEventAllOfUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *EmojiReactionEventAllOfUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmojiReactionEventAllOfUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmojiReactionEventAllOfUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmojiReactionEventAllOfUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *EmojiReactionEventAllOfUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *EmojiReactionEventAllOfUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *EmojiReactionEventAllOfUser) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *EmojiReactionEventAllOfUser) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetIsMirrorDummy

`func (o *EmojiReactionEventAllOfUser) GetIsMirrorDummy() bool`

GetIsMirrorDummy returns the IsMirrorDummy field if non-nil, zero value otherwise.

### GetIsMirrorDummyOk

`func (o *EmojiReactionEventAllOfUser) GetIsMirrorDummyOk() (*bool, bool)`

GetIsMirrorDummyOk returns a tuple with the IsMirrorDummy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMirrorDummy

`func (o *EmojiReactionEventAllOfUser) SetIsMirrorDummy(v bool)`

SetIsMirrorDummy sets IsMirrorDummy field to given value.

### HasIsMirrorDummy

`func (o *EmojiReactionEventAllOfUser) HasIsMirrorDummy() bool`

HasIsMirrorDummy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


