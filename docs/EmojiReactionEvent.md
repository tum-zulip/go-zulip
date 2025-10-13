# EmojiReactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiName** | Pointer to **string** | Name of the emoji.  | [optional] 
**EmojiCode** | Pointer to **string** | A unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  | [optional] 
**ReactionType** | Pointer to **string** | A string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](zulip.com/help/custom-emoji.  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \&quot;zulip\&quot;).  | [optional] 
**UserId** | Pointer to **int32** | The ID of the user who added the reaction.  **Changes**: New in Zulip 3.0 (feature level 2). The &#x60;user&#x60; object is deprecated and will be removed in the future.  | [optional] 
**User** | Pointer to [**EmojiReactionEventAllOfUser**](EmojiReactionEventAllOfUser.md) |  | [optional] 

## Methods

### NewEmojiReactionEvent

`func NewEmojiReactionEvent() *EmojiReactionEvent`

NewEmojiReactionEvent instantiates a new EmojiReactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiReactionEventWithDefaults

`func NewEmojiReactionEventWithDefaults() *EmojiReactionEvent`

NewEmojiReactionEventWithDefaults instantiates a new EmojiReactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiName

`func (o *EmojiReactionEvent) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *EmojiReactionEvent) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *EmojiReactionEvent) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *EmojiReactionEvent) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetEmojiCode

`func (o *EmojiReactionEvent) GetEmojiCode() string`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *EmojiReactionEvent) GetEmojiCodeOk() (*string, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *EmojiReactionEvent) SetEmojiCode(v string)`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *EmojiReactionEvent) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### GetReactionType

`func (o *EmojiReactionEvent) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *EmojiReactionEvent) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *EmojiReactionEvent) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *EmojiReactionEvent) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetUserId

`func (o *EmojiReactionEvent) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmojiReactionEvent) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmojiReactionEvent) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EmojiReactionEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *EmojiReactionEvent) GetUser() EmojiReactionEventAllOfUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EmojiReactionEvent) GetUserOk() (*EmojiReactionEventAllOfUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EmojiReactionEvent) SetUser(v EmojiReactionEventAllOfUser)`

SetUser sets User field to given value.

### HasUser

`func (o *EmojiReactionEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


