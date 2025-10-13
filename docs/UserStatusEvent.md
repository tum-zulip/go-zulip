# UserStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Away** | Pointer to **bool** | Whether the user has marked themself \&quot;away\&quot; with this status.  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, &#x60;away&#x60; is a legacy way to access the user&#39;s &#x60;presence_enabled&#x60; setting, with &#x60;away &#x3D; !presence_enabled&#x60;. To be removed in a future release.  | [optional] 
**StatusText** | Pointer to **string** | The text content of the status message.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting or writing a message.  | [optional] 
**EmojiName** | Pointer to **string** | The [emoji name](/api/update-status#parameter-emoji_name) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**EmojiCode** | Pointer to **string** | The [emoji code](/api/update-status#parameter-emoji_code) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**ReactionType** | Pointer to **string** | The [emoji type](/api/update-status#parameter-reaction_type) for the emoji the user selected for their new status.  This will be &#x60;\&quot;\&quot;&#x60; for users who set a status without selecting an emoji.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**UserId** | Pointer to **int32** | The ID of the user whose status changed.  | [optional] 

## Methods

### NewEventEnvelopeOneOf26

`func NewEventEnvelopeOneOf26() *UserStatusEvent`

NewEventEnvelopeOneOf26 instantiates a new UserStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf26WithDefaults

`func NewEventEnvelopeOneOf26WithDefaults() *UserStatusEvent`

NewEventEnvelopeOneOf26WithDefaults instantiates a new UserStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserStatusEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserStatusEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserStatusEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserStatusEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserStatusEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserStatusEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserStatusEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserStatusEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAway

`func (o *UserStatusEvent) GetAway() bool`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *UserStatusEvent) GetAwayOk() (*bool, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAway

`func (o *UserStatusEvent) SetAway(v bool)`

SetAway sets Away field to given value.

### HasAway

`func (o *UserStatusEvent) HasAway() bool`

HasAway returns a boolean if a field has been set.

### GetStatusText

`func (o *UserStatusEvent) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *UserStatusEvent) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *UserStatusEvent) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *UserStatusEvent) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetEmojiName

`func (o *UserStatusEvent) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *UserStatusEvent) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *UserStatusEvent) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *UserStatusEvent) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetEmojiCode

`func (o *UserStatusEvent) GetEmojiCode() string`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *UserStatusEvent) GetEmojiCodeOk() (*string, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *UserStatusEvent) SetEmojiCode(v string)`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *UserStatusEvent) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### GetReactionType

`func (o *UserStatusEvent) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *UserStatusEvent) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *UserStatusEvent) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *UserStatusEvent) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetUserId

`func (o *UserStatusEvent) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserStatusEvent) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserStatusEvent) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserStatusEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


