# BasicBotBaseServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The URL the outgoing webhook is configured to post to.  | [optional] 
**Token** | Pointer to **string** | A unique token that the third-party service can use to confirm that the request is indeed coming from Zulip.  | [optional] 
**Interface** | Pointer to **int32** | An integer indicating what format requests are posted in:  - 1 &#x3D; Zulip&#39;s native outgoing webhook format. - 2 &#x3D; Emulate the Slack outgoing webhook format.  | [optional] 
**ServiceName** | Pointer to **string** | The name of the bot.  | [optional] 
**ConfigData** | Pointer to **map[string]string** | A dictionary of string key/value pairs, which describe the configuration for the bot. These are usually details like API keys, and are unique to the integration/bot. Can be an empty dictionary.  | [optional] 

## Methods

### NewBasicBotBaseServicesInner

`func NewBasicBotBaseServicesInner() *BasicBotBaseServicesInner`

NewBasicBotBaseServicesInner instantiates a new BasicBotBaseServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicBotBaseServicesInnerWithDefaults

`func NewBasicBotBaseServicesInnerWithDefaults() *BasicBotBaseServicesInner`

NewBasicBotBaseServicesInnerWithDefaults instantiates a new BasicBotBaseServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *BasicBotBaseServicesInner) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *BasicBotBaseServicesInner) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *BasicBotBaseServicesInner) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *BasicBotBaseServicesInner) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetToken

`func (o *BasicBotBaseServicesInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BasicBotBaseServicesInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BasicBotBaseServicesInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BasicBotBaseServicesInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetInterface

`func (o *BasicBotBaseServicesInner) GetInterface() int32`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BasicBotBaseServicesInner) GetInterfaceOk() (*int32, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BasicBotBaseServicesInner) SetInterface(v int32)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *BasicBotBaseServicesInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetServiceName

`func (o *BasicBotBaseServicesInner) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BasicBotBaseServicesInner) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BasicBotBaseServicesInner) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BasicBotBaseServicesInner) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetConfigData

`func (o *BasicBotBaseServicesInner) GetConfigData() map[string]string`

GetConfigData returns the ConfigData field if non-nil, zero value otherwise.

### GetConfigDataOk

`func (o *BasicBotBaseServicesInner) GetConfigDataOk() (*map[string]string, bool)`

GetConfigDataOk returns a tuple with the ConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigData

`func (o *BasicBotBaseServicesInner) SetConfigData(v map[string]string)`

SetConfigData sets ConfigData field to given value.

### HasConfigData

`func (o *BasicBotBaseServicesInner) HasConfigData() bool`

HasConfigData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


