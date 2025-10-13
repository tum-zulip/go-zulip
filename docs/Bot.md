# Bot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **interface{}** |  | [optional] 
**FullName** | Pointer to **interface{}** |  | [optional] 
**ApiKey** | Pointer to **interface{}** |  | [optional] 
**DefaultSendingStream** | Pointer to **interface{}** |  | [optional] 
**DefaultEventsRegisterStream** | Pointer to **interface{}** |  | [optional] 
**DefaultAllPublicStreams** | Pointer to **interface{}** |  | [optional] 
**AvatarUrl** | Pointer to **interface{}** |  | [optional] 
**OwnerId** | Pointer to **interface{}** |  | [optional] 
**Services** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **string** | The email of the bot.  | [optional] 
**BotType** | Pointer to **NullableInt32** | An integer describing the type of bot:  - &#x60;1&#x60; for a &#x60;Generic&#x60; bot. - &#x60;2&#x60; for an &#x60;Incoming webhook&#x60; bot. - &#x60;3&#x60; for an &#x60;Outgoing webhook&#x60; bot. - &#x60;4&#x60; for an &#x60;Embedded&#x60; bot.  | [optional] 
**IsActive** | Pointer to **bool** | A boolean describing whether the user account has been deactivated.  | [optional] 

## Methods

### NewBot

`func NewBot() *Bot`

NewBot instantiates a new Bot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotWithDefaults

`func NewBotWithDefaults() *Bot`

NewBotWithDefaults instantiates a new Bot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *Bot) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Bot) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Bot) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Bot) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *Bot) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *Bot) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetFullName

`func (o *Bot) GetFullName() interface{}`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Bot) GetFullNameOk() (*interface{}, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Bot) SetFullName(v interface{})`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Bot) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *Bot) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *Bot) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetApiKey

`func (o *Bot) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Bot) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Bot) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Bot) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *Bot) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *Bot) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetDefaultSendingStream

`func (o *Bot) GetDefaultSendingStream() interface{}`

GetDefaultSendingStream returns the DefaultSendingStream field if non-nil, zero value otherwise.

### GetDefaultSendingStreamOk

`func (o *Bot) GetDefaultSendingStreamOk() (*interface{}, bool)`

GetDefaultSendingStreamOk returns a tuple with the DefaultSendingStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSendingStream

`func (o *Bot) SetDefaultSendingStream(v interface{})`

SetDefaultSendingStream sets DefaultSendingStream field to given value.

### HasDefaultSendingStream

`func (o *Bot) HasDefaultSendingStream() bool`

HasDefaultSendingStream returns a boolean if a field has been set.

### SetDefaultSendingStreamNil

`func (o *Bot) SetDefaultSendingStreamNil(b bool)`

 SetDefaultSendingStreamNil sets the value for DefaultSendingStream to be an explicit nil

### UnsetDefaultSendingStream
`func (o *Bot) UnsetDefaultSendingStream()`

UnsetDefaultSendingStream ensures that no value is present for DefaultSendingStream, not even an explicit nil
### GetDefaultEventsRegisterStream

`func (o *Bot) GetDefaultEventsRegisterStream() interface{}`

GetDefaultEventsRegisterStream returns the DefaultEventsRegisterStream field if non-nil, zero value otherwise.

### GetDefaultEventsRegisterStreamOk

`func (o *Bot) GetDefaultEventsRegisterStreamOk() (*interface{}, bool)`

GetDefaultEventsRegisterStreamOk returns a tuple with the DefaultEventsRegisterStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEventsRegisterStream

`func (o *Bot) SetDefaultEventsRegisterStream(v interface{})`

SetDefaultEventsRegisterStream sets DefaultEventsRegisterStream field to given value.

### HasDefaultEventsRegisterStream

`func (o *Bot) HasDefaultEventsRegisterStream() bool`

HasDefaultEventsRegisterStream returns a boolean if a field has been set.

### SetDefaultEventsRegisterStreamNil

`func (o *Bot) SetDefaultEventsRegisterStreamNil(b bool)`

 SetDefaultEventsRegisterStreamNil sets the value for DefaultEventsRegisterStream to be an explicit nil

### UnsetDefaultEventsRegisterStream
`func (o *Bot) UnsetDefaultEventsRegisterStream()`

UnsetDefaultEventsRegisterStream ensures that no value is present for DefaultEventsRegisterStream, not even an explicit nil
### GetDefaultAllPublicStreams

`func (o *Bot) GetDefaultAllPublicStreams() interface{}`

GetDefaultAllPublicStreams returns the DefaultAllPublicStreams field if non-nil, zero value otherwise.

### GetDefaultAllPublicStreamsOk

`func (o *Bot) GetDefaultAllPublicStreamsOk() (*interface{}, bool)`

GetDefaultAllPublicStreamsOk returns a tuple with the DefaultAllPublicStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAllPublicStreams

`func (o *Bot) SetDefaultAllPublicStreams(v interface{})`

SetDefaultAllPublicStreams sets DefaultAllPublicStreams field to given value.

### HasDefaultAllPublicStreams

`func (o *Bot) HasDefaultAllPublicStreams() bool`

HasDefaultAllPublicStreams returns a boolean if a field has been set.

### SetDefaultAllPublicStreamsNil

`func (o *Bot) SetDefaultAllPublicStreamsNil(b bool)`

 SetDefaultAllPublicStreamsNil sets the value for DefaultAllPublicStreams to be an explicit nil

### UnsetDefaultAllPublicStreams
`func (o *Bot) UnsetDefaultAllPublicStreams()`

UnsetDefaultAllPublicStreams ensures that no value is present for DefaultAllPublicStreams, not even an explicit nil
### GetAvatarUrl

`func (o *Bot) GetAvatarUrl() interface{}`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Bot) GetAvatarUrlOk() (*interface{}, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Bot) SetAvatarUrl(v interface{})`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Bot) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *Bot) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *Bot) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetOwnerId

`func (o *Bot) GetOwnerId() interface{}`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Bot) GetOwnerIdOk() (*interface{}, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Bot) SetOwnerId(v interface{})`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Bot) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *Bot) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *Bot) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetServices

`func (o *Bot) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Bot) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Bot) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *Bot) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *Bot) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *Bot) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetEmail

`func (o *Bot) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Bot) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Bot) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Bot) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBotType

`func (o *Bot) GetBotType() int32`

GetBotType returns the BotType field if non-nil, zero value otherwise.

### GetBotTypeOk

`func (o *Bot) GetBotTypeOk() (*int32, bool)`

GetBotTypeOk returns a tuple with the BotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotType

`func (o *Bot) SetBotType(v int32)`

SetBotType sets BotType field to given value.

### HasBotType

`func (o *Bot) HasBotType() bool`

HasBotType returns a boolean if a field has been set.

### SetBotTypeNil

`func (o *Bot) SetBotTypeNil(b bool)`

 SetBotTypeNil sets the value for BotType to be an explicit nil

### UnsetBotType
`func (o *Bot) UnsetBotType()`

UnsetBotType ensures that no value is present for BotType, not even an explicit nil
### GetIsActive

`func (o *Bot) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Bot) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Bot) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Bot) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


