# UserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Away** | Pointer to **bool** | If present, the user has marked themself \&quot;away\&quot;.  **Changes**: Deprecated in Zulip 6.0 (feature level 148); starting with that feature level, &#x60;away&#x60; is a legacy way to access the user&#39;s &#x60;presence_enabled&#x60; setting, with &#x60;away &#x3D; !presence_enabled&#x60;. To be removed in a future release.  | [optional] 
**StatusText** | Pointer to **string** | If present, the text content of the user&#39;s status message.  | [optional] 
**EmojiName** | Pointer to **string** | If present, the name for the emoji to associate with the user&#39;s status.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**EmojiCode** | Pointer to **string** | If present, a unique identifier, defining the specific emoji codepoint requested, within the namespace of the &#x60;reaction_type&#x60;.  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 
**ReactionType** | Pointer to **string** | If present, a string indicating the type of emoji. Each emoji &#x60;reaction_type&#x60; has an independent namespace for values of &#x60;emoji_code&#x60;.  Must be one of the following values:  - &#x60;unicode_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be a   dash-separated hex encoding of the sequence of Unicode codepoints   that define this emoji in the Unicode specification.  - &#x60;realm_emoji&#x60; : In this namespace, &#x60;emoji_code&#x60; will be the ID of   the uploaded [custom emoji](/help/custom-emoji).  - &#x60;zulip_extra_emoji&#x60; : These are special emoji included with Zulip.   In this namespace, &#x60;emoji_code&#x60; will be the name of the emoji (e.g.   \&quot;zulip\&quot;).  **Changes**: New in Zulip 5.0 (feature level 86).  | [optional] 

## Methods

### NewUserStatus

`func NewUserStatus() *UserStatus`

NewUserStatus instantiates a new UserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusWithDefaults

`func NewUserStatusWithDefaults() *UserStatus`

NewUserStatusWithDefaults instantiates a new UserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAway

`func (o *UserStatus) GetAway() bool`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *UserStatus) GetAwayOk() (*bool, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAway

`func (o *UserStatus) SetAway(v bool)`

SetAway sets Away field to given value.

### HasAway

`func (o *UserStatus) HasAway() bool`

HasAway returns a boolean if a field has been set.

### GetStatusText

`func (o *UserStatus) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *UserStatus) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *UserStatus) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *UserStatus) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetEmojiName

`func (o *UserStatus) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *UserStatus) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *UserStatus) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *UserStatus) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetEmojiCode

`func (o *UserStatus) GetEmojiCode() string`

GetEmojiCode returns the EmojiCode field if non-nil, zero value otherwise.

### GetEmojiCodeOk

`func (o *UserStatus) GetEmojiCodeOk() (*string, bool)`

GetEmojiCodeOk returns a tuple with the EmojiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiCode

`func (o *UserStatus) SetEmojiCode(v string)`

SetEmojiCode sets EmojiCode field to given value.

### HasEmojiCode

`func (o *UserStatus) HasEmojiCode() bool`

HasEmojiCode returns a boolean if a field has been set.

### GetReactionType

`func (o *UserStatus) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *UserStatus) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *UserStatus) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *UserStatus) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


