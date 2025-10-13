# ReactionRemoveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiName** | Pointer to **interface{}** |  | [optional] 
**EmojiCode** | Pointer to **interface{}** |  | [optional] 
**ReactionType** | Pointer to **interface{}** |  | [optional] 
**UserId** | Pointer to **interface{}** |  | [optional] 
**User** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to [**EventEnvelopeOneOf20AllOfType**](EventEnvelopeOneOf20AllOfType.md) |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **int32** | The ID of the message from which the reaction was removed.  | [optional] 

## Methods

### NewEventEnvelopeOneOf20

`func NewEventEnvelopeOneOf20() *ReactionRemoveEvent`

NewEventEnvelopeOneOf20 instantiates a new ReactionRemoveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf20WithDefaults

`func NewEventEnvelopeOneOf20WithDefaults() *ReactionRemoveEvent`

NewEventEnvelopeOneOf20WithDefaults instantiates a new ReactionRemoveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiName

`func (o *ReactionRemoveEvent) GetEmojiName() interface{}`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *ReactionRemoveEvent) GetEmojiNameOk() (*interface{}, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *ReactionRemoveEvent) SetEmojiName(v interface{})`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *ReactionRemoveEvent) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *ReactionRemoveEvent) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *ReactionRemoveEvent) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetEmojiCode

`func (o *ReactionRemoveEvent) GetEmojiCode() interface{}`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *ReactionRemoveEvent) GetEmojiCodeOk() (*interface{}, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *ReactionRemoveEvent) SetEmojiCode(v interface{})`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *ReactionRemoveEvent) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### SetEmojiCodeNil

`func (o *ReactionRemoveEvent) SetEmojiCodeNil(b bool)`

 SetEmojiCodeNil sets the value for EmojiCode to be an explicit nil

### UnsetEmojiCode
`func (o *ReactionRemoveEvent) UnsetEmojiCode()`

UnsetEmojiCode ensures that no value is present for EmojiCode, not even an explicit nil
### GetReactionType

`func (o *ReactionRemoveEvent) GetReactionType() interface{}`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionRemoveEvent) GetReactionTypeOk() (*interface{}, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionRemoveEvent) SetReactionType(v interface{})`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *ReactionRemoveEvent) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### SetReactionTypeNil

`func (o *ReactionRemoveEvent) SetReactionTypeNil(b bool)`

 SetReactionTypeNil sets the value for ReactionType to be an explicit nil

### UnsetReactionType
`func (o *ReactionRemoveEvent) UnsetReactionType()`

UnsetReactionType ensures that no value is present for ReactionType, not even an explicit nil
### GetUserId

`func (o *ReactionRemoveEvent) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReactionRemoveEvent) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReactionRemoveEvent) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ReactionRemoveEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ReactionRemoveEvent) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ReactionRemoveEvent) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUser

`func (o *ReactionRemoveEvent) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReactionRemoveEvent) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReactionRemoveEvent) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *ReactionRemoveEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ReactionRemoveEvent) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ReactionRemoveEvent) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetId

`func (o *ReactionRemoveEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReactionRemoveEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReactionRemoveEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReactionRemoveEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReactionRemoveEvent) GetType() EventEnvelopeOneOf20AllOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionRemoveEvent) GetTypeOk() (*EventEnvelopeOneOf20AllOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionRemoveEvent) SetType(v EventEnvelopeOneOf20AllOfType)`

SetType sets Type field to given value.

### HasType

`func (o *ReactionRemoveEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ReactionRemoveEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ReactionRemoveEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ReactionRemoveEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ReactionRemoveEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetMessageId

`func (o *ReactionRemoveEvent) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ReactionRemoveEvent) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ReactionRemoveEvent) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ReactionRemoveEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


