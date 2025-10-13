# ReactionAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiName** | Pointer to **interface{}** |  | [optional] 
**EmojiCode** | Pointer to **interface{}** |  | [optional] 
**ReactionType** | Pointer to **interface{}** |  | [optional] 
**UserId** | Pointer to **interface{}** |  | [optional] 
**User** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf19AllOfType**](GetEvents200ResponseAllOfEventsInnerOneOf19AllOfType.md) |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **int32** | The ID of the message to which a reaction was added.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf19

`func NewGetEvents200ResponseAllOfEventsInnerOneOf19() *ReactionAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf19 instantiates a new ReactionAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf19WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf19WithDefaults() *ReactionAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf19WithDefaults instantiates a new ReactionAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiName

`func (o *ReactionAddEvent) GetEmojiName() interface{}`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *ReactionAddEvent) GetEmojiNameOk() (*interface{}, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *ReactionAddEvent) SetEmojiName(v interface{})`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *ReactionAddEvent) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *ReactionAddEvent) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *ReactionAddEvent) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetEmojiCode

`func (o *ReactionAddEvent) GetEmojiCode() interface{}`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *ReactionAddEvent) GetEmojiCodeOk() (*interface{}, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *ReactionAddEvent) SetEmojiCode(v interface{})`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *ReactionAddEvent) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### SetEmojiCodeNil

`func (o *ReactionAddEvent) SetEmojiCodeNil(b bool)`

 SetEmojiCodeNil sets the value for EmojiCode to be an explicit nil

### UnsetEmojiCode
`func (o *ReactionAddEvent) UnsetEmojiCode()`

UnsetEmojiCode ensures that no value is present for EmojiCode, not even an explicit nil
### GetReactionType

`func (o *ReactionAddEvent) GetReactionType() interface{}`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionAddEvent) GetReactionTypeOk() (*interface{}, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionAddEvent) SetReactionType(v interface{})`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *ReactionAddEvent) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### SetReactionTypeNil

`func (o *ReactionAddEvent) SetReactionTypeNil(b bool)`

 SetReactionTypeNil sets the value for ReactionType to be an explicit nil

### UnsetReactionType
`func (o *ReactionAddEvent) UnsetReactionType()`

UnsetReactionType ensures that no value is present for ReactionType, not even an explicit nil
### GetUserId

`func (o *ReactionAddEvent) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReactionAddEvent) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReactionAddEvent) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ReactionAddEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ReactionAddEvent) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ReactionAddEvent) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUser

`func (o *ReactionAddEvent) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReactionAddEvent) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReactionAddEvent) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *ReactionAddEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ReactionAddEvent) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ReactionAddEvent) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetId

`func (o *ReactionAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReactionAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReactionAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReactionAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReactionAddEvent) GetType() GetEvents200ResponseAllOfEventsInnerOneOf19AllOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionAddEvent) GetTypeOk() (*GetEvents200ResponseAllOfEventsInnerOneOf19AllOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionAddEvent) SetType(v GetEvents200ResponseAllOfEventsInnerOneOf19AllOfType)`

SetType sets Type field to given value.

### HasType

`func (o *ReactionAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ReactionAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ReactionAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ReactionAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ReactionAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetMessageId

`func (o *ReactionAddEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ReactionAddEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ReactionAddEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ReactionAddEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


