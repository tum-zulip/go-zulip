# EmojiReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiName** | Pointer to **interface{}** |  | [optional] 
**EmojiCode** | Pointer to **interface{}** |  | [optional] 
**ReactionType** | Pointer to **interface{}** |  | [optional] 
**UserId** | Pointer to **int32** | The ID of the user who added the reaction.  **Changes**: New in Zulip 3.0 (feature level 2). The &#x60;user&#x60; object is deprecated and will be removed in the future.  In Zulip 10.0 (feature level 328), the deprecated &#x60;user&#x60; object was removed which contained the following properties: &#x60;id&#x60;, &#x60;email&#x60;, &#x60;full_name&#x60; and &#x60;is_mirror_dummy&#x60;.  | [optional] 

## Methods

### NewEmojiReaction

`func NewEmojiReaction() *EmojiReaction`

NewEmojiReaction instantiates a new EmojiReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiReactionWithDefaults

`func NewEmojiReactionWithDefaults() *EmojiReaction`

NewEmojiReactionWithDefaults instantiates a new EmojiReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiName

`func (o *EmojiReaction) GetEmojiName() interface{}`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *EmojiReaction) GetEmojiNameOk() (*interface{}, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *EmojiReaction) SetEmojiName(v interface{})`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *EmojiReaction) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *EmojiReaction) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *EmojiReaction) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetEmojiCode

`func (o *EmojiReaction) GetEmojiCode() interface{}`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *EmojiReaction) GetEmojiCodeOk() (*interface{}, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *EmojiReaction) SetEmojiCode(v interface{})`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *EmojiReaction) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### SetEmojiCodeNil

`func (o *EmojiReaction) SetEmojiCodeNil(b bool)`

 SetEmojiCodeNil sets the value for EmojiCode to be an explicit nil

### UnsetEmojiCode
`func (o *EmojiReaction) UnsetEmojiCode()`

UnsetEmojiCode ensures that no value is present for EmojiCode, not even an explicit nil
### GetReactionType

`func (o *EmojiReaction) GetReactionType() interface{}`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *EmojiReaction) GetReactionTypeOk() (*interface{}, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *EmojiReaction) SetReactionType(v interface{})`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *EmojiReaction) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### SetReactionTypeNil

`func (o *EmojiReaction) SetReactionTypeNil(b bool)`

 SetReactionTypeNil sets the value for ReactionType to be an explicit nil

### UnsetReactionType
`func (o *EmojiReaction) UnsetReactionType()`

UnsetReactionType ensures that no value is present for ReactionType, not even an explicit nil
### GetUserId

`func (o *EmojiReaction) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmojiReaction) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmojiReaction) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EmojiReaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


