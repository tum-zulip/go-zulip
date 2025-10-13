# BasicBotBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** | The user ID of the bot.  | [optional] 
**FullName** | Pointer to **string** | The full name of the bot.  | [optional] 
**ApiKey** | Pointer to **string** | The API key of the bot which it uses to make API requests.  | [optional] 
**DefaultSendingStream** | Pointer to **NullableString** | The default sending channel of the bot. If &#x60;null&#x60;, the bot doesn&#39;t have a default sending channel.  | [optional] 
**DefaultEventsRegisterStream** | Pointer to **NullableString** | The default channel for which the bot receives events/register data. If &#x60;null&#x60;, the bot doesn&#39;t have such a default channel.  | [optional] 
**DefaultAllPublicStreams** | Pointer to **bool** | Whether the bot can send messages to all channels by default.  | [optional] 
**AvatarUrl** | Pointer to **string** | The URL of the bot&#39;s avatar.  | [optional] 
**OwnerId** | Pointer to **NullableInt32** | The user ID of the bot&#39;s owner.  If &#x60;null&#x60;, the bot has no owner.  | [optional] 
**Services** | Pointer to [**[]BasicBotBaseServicesInner**](BasicBotBaseServicesInner.md) | An array containing extra configuration fields only relevant for outgoing webhook bots and embedded bots. This is always a single-element array.  We consider this part of the Zulip API to be unstable; it is used only for UI elements for administering bots and is likely to change.  | [optional] 

## Methods

### NewBasicBotBase

`func NewBasicBotBase() *BasicBotBase`

NewBasicBotBase instantiates a new BasicBotBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicBotBaseWithDefaults

`func NewBasicBotBaseWithDefaults() *BasicBotBase`

NewBasicBotBaseWithDefaults instantiates a new BasicBotBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BasicBotBase) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BasicBotBase) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BasicBotBase) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BasicBotBase) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFullName

`func (o *BasicBotBase) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BasicBotBase) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BasicBotBase) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *BasicBotBase) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetApiKey

`func (o *BasicBotBase) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *BasicBotBase) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *BasicBotBase) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *BasicBotBase) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDefaultSendingStream

`func (o *BasicBotBase) GetDefaultSendingStream() string`

GetDefaultSendingStream returns the DefaultSendingStream field if non-nil, zero value otherwise.

### GetDefaultSendingStreamOk

`func (o *BasicBotBase) GetDefaultSendingStreamOk() (*string, bool)`

GetDefaultSendingStreamOk returns a tuple with the DefaultSendingStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSendingStream

`func (o *BasicBotBase) SetDefaultSendingStream(v string)`

SetDefaultSendingStream sets DefaultSendingStream field to given value.

### HasDefaultSendingStream

`func (o *BasicBotBase) HasDefaultSendingStream() bool`

HasDefaultSendingStream returns a boolean if a field has been set.

### SetDefaultSendingStreamNil

`func (o *BasicBotBase) SetDefaultSendingStreamNil(b bool)`

 SetDefaultSendingStreamNil sets the value for DefaultSendingStream to be an explicit nil

### UnsetDefaultSendingStream
`func (o *BasicBotBase) UnsetDefaultSendingStream()`

UnsetDefaultSendingStream ensures that no value is present for DefaultSendingStream, not even an explicit nil
### GetDefaultEventsRegisterStream

`func (o *BasicBotBase) GetDefaultEventsRegisterStream() string`

GetDefaultEventsRegisterStream returns the DefaultEventsRegisterStream field if non-nil, zero value otherwise.

### GetDefaultEventsRegisterStreamOk

`func (o *BasicBotBase) GetDefaultEventsRegisterStreamOk() (*string, bool)`

GetDefaultEventsRegisterStreamOk returns a tuple with the DefaultEventsRegisterStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEventsRegisterStream

`func (o *BasicBotBase) SetDefaultEventsRegisterStream(v string)`

SetDefaultEventsRegisterStream sets DefaultEventsRegisterStream field to given value.

### HasDefaultEventsRegisterStream

`func (o *BasicBotBase) HasDefaultEventsRegisterStream() bool`

HasDefaultEventsRegisterStream returns a boolean if a field has been set.

### SetDefaultEventsRegisterStreamNil

`func (o *BasicBotBase) SetDefaultEventsRegisterStreamNil(b bool)`

 SetDefaultEventsRegisterStreamNil sets the value for DefaultEventsRegisterStream to be an explicit nil

### UnsetDefaultEventsRegisterStream
`func (o *BasicBotBase) UnsetDefaultEventsRegisterStream()`

UnsetDefaultEventsRegisterStream ensures that no value is present for DefaultEventsRegisterStream, not even an explicit nil
### GetDefaultAllPublicStreams

`func (o *BasicBotBase) GetDefaultAllPublicStreams() bool`

GetDefaultAllPublicStreams returns the DefaultAllPublicStreams field if non-nil, zero value otherwise.

### GetDefaultAllPublicStreamsOk

`func (o *BasicBotBase) GetDefaultAllPublicStreamsOk() (*bool, bool)`

GetDefaultAllPublicStreamsOk returns a tuple with the DefaultAllPublicStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAllPublicStreams

`func (o *BasicBotBase) SetDefaultAllPublicStreams(v bool)`

SetDefaultAllPublicStreams sets DefaultAllPublicStreams field to given value.

### HasDefaultAllPublicStreams

`func (o *BasicBotBase) HasDefaultAllPublicStreams() bool`

HasDefaultAllPublicStreams returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *BasicBotBase) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *BasicBotBase) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *BasicBotBase) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *BasicBotBase) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetOwnerId

`func (o *BasicBotBase) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BasicBotBase) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BasicBotBase) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BasicBotBase) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *BasicBotBase) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *BasicBotBase) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetServices

`func (o *BasicBotBase) GetServices() []BasicBotBaseServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BasicBotBase) GetServicesOk() (*[]BasicBotBaseServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BasicBotBase) SetServices(v []BasicBotBaseServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *BasicBotBase) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


