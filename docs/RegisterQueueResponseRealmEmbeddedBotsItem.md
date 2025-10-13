# RegisterQueueResponseRealmEmbeddedBotsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the bot.  | [optional] 
**Config** | Pointer to **map[string]string** | A dictionary of string key/value pairs, which describe the configuration for the bot. These are usually details like API keys, and are unique to the integration/bot. Can be an empty dictionary.  | [optional] 

## Methods

### NewRegisterQueueResponseRealmEmbeddedBotsItem

`func NewRegisterQueueResponseRealmEmbeddedBotsItem() *RegisterQueueResponseRealmEmbeddedBotsItem`

NewRegisterQueueResponseRealmEmbeddedBotsItem instantiates a new RegisterQueueResponseRealmEmbeddedBotsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseRealmEmbeddedBotsItemWithDefaults

`func NewRegisterQueueResponseRealmEmbeddedBotsItemWithDefaults() *RegisterQueueResponseRealmEmbeddedBotsItem`

NewRegisterQueueResponseRealmEmbeddedBotsItemWithDefaults instantiates a new RegisterQueueResponseRealmEmbeddedBotsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RegisterQueueResponseRealmEmbeddedBotsItem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


