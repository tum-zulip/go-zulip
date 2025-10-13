# BasicBotBaseServicesInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The URL the outgoing webhook is configured to post to.  | [optional] 
**Token** | Pointer to **string** | A unique token that the third-party service can use to confirm that the request is indeed coming from Zulip.  | [optional] 
**Interface** | Pointer to **int32** | An integer indicating what format requests are posted in:  - 1 &#x3D; Zulip&#39;s native outgoing webhook format. - 2 &#x3D; Emulate the Slack outgoing webhook format.  | [optional] 

## Methods

### NewBasicBotBaseServicesInnerOneOf

`func NewBasicBotBaseServicesInnerOneOf() *BasicBotBaseServicesInnerOneOf`

NewBasicBotBaseServicesInnerOneOf instantiates a new BasicBotBaseServicesInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicBotBaseServicesInnerOneOfWithDefaults

`func NewBasicBotBaseServicesInnerOneOfWithDefaults() *BasicBotBaseServicesInnerOneOf`

NewBasicBotBaseServicesInnerOneOfWithDefaults instantiates a new BasicBotBaseServicesInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *BasicBotBaseServicesInnerOneOf) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *BasicBotBaseServicesInnerOneOf) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *BasicBotBaseServicesInnerOneOf) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *BasicBotBaseServicesInnerOneOf) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetToken

`func (o *BasicBotBaseServicesInnerOneOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BasicBotBaseServicesInnerOneOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BasicBotBaseServicesInnerOneOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BasicBotBaseServicesInnerOneOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetInterface

`func (o *BasicBotBaseServicesInnerOneOf) GetInterface() int32`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BasicBotBaseServicesInnerOneOf) GetInterfaceOk() (*int32, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BasicBotBaseServicesInnerOneOf) SetInterface(v int32)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *BasicBotBaseServicesInnerOneOf) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


