# BasicBot

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
**IsActive** | Pointer to **bool** | A boolean describing whether the user account has been deactivated.  **Changes**: New in Zulip 8.0 (feature level 222). Previously we sent &#x60;realm_user&#x60; event with &#x60;op&#x60; field set to &#x60;remove&#x60; when deactivating a bot and &#x60;realm_user&#x60; event with &#x60;op&#x60; field set to &#x60;add&#x60; when reactivating a bot.  | [optional] 

## Methods

### NewBasicBot

`func NewBasicBot() *BasicBot`

NewBasicBot instantiates a new BasicBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicBotWithDefaults

`func NewBasicBotWithDefaults() *BasicBot`

NewBasicBotWithDefaults instantiates a new BasicBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BasicBot) GetUserId() interface{}`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BasicBot) GetUserIdOk() (*interface{}, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BasicBot) SetUserId(v interface{})`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BasicBot) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BasicBot) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BasicBot) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetFullName

`func (o *BasicBot) GetFullName() interface{}`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *BasicBot) GetFullNameOk() (*interface{}, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *BasicBot) SetFullName(v interface{})`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *BasicBot) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *BasicBot) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *BasicBot) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetApiKey

`func (o *BasicBot) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *BasicBot) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *BasicBot) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *BasicBot) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *BasicBot) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *BasicBot) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetDefaultSendingStream

`func (o *BasicBot) GetDefaultSendingStream() interface{}`

GetDefaultSendingStream returns the DefaultSendingStream field if non-nil, zero value otherwise.

### GetDefaultSendingStreamOk

`func (o *BasicBot) GetDefaultSendingStreamOk() (*interface{}, bool)`

GetDefaultSendingStreamOk returns a tuple with the DefaultSendingStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSendingStream

`func (o *BasicBot) SetDefaultSendingStream(v interface{})`

SetDefaultSendingStream sets DefaultSendingStream field to given value.

### HasDefaultSendingStream

`func (o *BasicBot) HasDefaultSendingStream() bool`

HasDefaultSendingStream returns a boolean if a field has been set.

### SetDefaultSendingStreamNil

`func (o *BasicBot) SetDefaultSendingStreamNil(b bool)`

 SetDefaultSendingStreamNil sets the value for DefaultSendingStream to be an explicit nil

### UnsetDefaultSendingStream
`func (o *BasicBot) UnsetDefaultSendingStream()`

UnsetDefaultSendingStream ensures that no value is present for DefaultSendingStream, not even an explicit nil
### GetDefaultEventsRegisterStream

`func (o *BasicBot) GetDefaultEventsRegisterStream() interface{}`

GetDefaultEventsRegisterStream returns the DefaultEventsRegisterStream field if non-nil, zero value otherwise.

### GetDefaultEventsRegisterStreamOk

`func (o *BasicBot) GetDefaultEventsRegisterStreamOk() (*interface{}, bool)`

GetDefaultEventsRegisterStreamOk returns a tuple with the DefaultEventsRegisterStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEventsRegisterStream

`func (o *BasicBot) SetDefaultEventsRegisterStream(v interface{})`

SetDefaultEventsRegisterStream sets DefaultEventsRegisterStream field to given value.

### HasDefaultEventsRegisterStream

`func (o *BasicBot) HasDefaultEventsRegisterStream() bool`

HasDefaultEventsRegisterStream returns a boolean if a field has been set.

### SetDefaultEventsRegisterStreamNil

`func (o *BasicBot) SetDefaultEventsRegisterStreamNil(b bool)`

 SetDefaultEventsRegisterStreamNil sets the value for DefaultEventsRegisterStream to be an explicit nil

### UnsetDefaultEventsRegisterStream
`func (o *BasicBot) UnsetDefaultEventsRegisterStream()`

UnsetDefaultEventsRegisterStream ensures that no value is present for DefaultEventsRegisterStream, not even an explicit nil
### GetDefaultAllPublicStreams

`func (o *BasicBot) GetDefaultAllPublicStreams() interface{}`

GetDefaultAllPublicStreams returns the DefaultAllPublicStreams field if non-nil, zero value otherwise.

### GetDefaultAllPublicStreamsOk

`func (o *BasicBot) GetDefaultAllPublicStreamsOk() (*interface{}, bool)`

GetDefaultAllPublicStreamsOk returns a tuple with the DefaultAllPublicStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAllPublicStreams

`func (o *BasicBot) SetDefaultAllPublicStreams(v interface{})`

SetDefaultAllPublicStreams sets DefaultAllPublicStreams field to given value.

### HasDefaultAllPublicStreams

`func (o *BasicBot) HasDefaultAllPublicStreams() bool`

HasDefaultAllPublicStreams returns a boolean if a field has been set.

### SetDefaultAllPublicStreamsNil

`func (o *BasicBot) SetDefaultAllPublicStreamsNil(b bool)`

 SetDefaultAllPublicStreamsNil sets the value for DefaultAllPublicStreams to be an explicit nil

### UnsetDefaultAllPublicStreams
`func (o *BasicBot) UnsetDefaultAllPublicStreams()`

UnsetDefaultAllPublicStreams ensures that no value is present for DefaultAllPublicStreams, not even an explicit nil
### GetAvatarUrl

`func (o *BasicBot) GetAvatarUrl() interface{}`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *BasicBot) GetAvatarUrlOk() (*interface{}, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *BasicBot) SetAvatarUrl(v interface{})`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *BasicBot) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *BasicBot) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *BasicBot) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetOwnerId

`func (o *BasicBot) GetOwnerId() interface{}`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BasicBot) GetOwnerIdOk() (*interface{}, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BasicBot) SetOwnerId(v interface{})`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BasicBot) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *BasicBot) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *BasicBot) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetServices

`func (o *BasicBot) GetServices() interface{}`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BasicBot) GetServicesOk() (*interface{}, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BasicBot) SetServices(v interface{})`

SetServices sets Services field to given value.

### HasServices

`func (o *BasicBot) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *BasicBot) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *BasicBot) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetIsActive

`func (o *BasicBot) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BasicBot) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BasicBot) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BasicBot) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


