# RealmIncomingWebhookBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A machine-readable unique name identifying the integration, all-lower-case without spaces.  | [optional] 
**DisplayName** | Pointer to **string** | A human-readable display name identifying the integration that this bot implements, intended to be used in menus for selecting which integration to create.  **Changes**: New in Zulip 8.0 (feature level 207).  | [optional] 
**AllEventTypes** | Pointer to **[]string** | For incoming webhook integrations that support the Zulip server filtering incoming events, the list of event types supported by it.  A null value will be present if this incoming webhook integration doesn&#39;t support such filtering.  **Changes**: New in Zulip 8.0 (feature level 207).  | [optional] 
**ConfigOptions** | Pointer to [**[]WebhookConfigOptionInner**](WebhookConfigOptionInner.md) | An array of configuration options that can be set when creating a bot user for this incoming webhook integration.  This is an unstable API. Please discuss in chat.zulip.org before using it.  **Changes**: As of Zulip 11.0 (feature level 403), this object is reserved for integration-specific configuration options that can be set when creating a bot user. Previously, this object also included optional webhook URL parameters, which are now specified in the &#x60;url_options&#x60; object.  Before Zulip 10.0 (feature level 318), this field was named &#x60;config&#x60;, and was reserved for configuration data key-value pairs.  | [optional] 
**UrlOptions** | Pointer to [**[]WebhookUrlOptionInner**](WebhookUrlOptionInner.md) | An array of optional URL parameter options for the incoming webhook integration. In the web app, these are used when [generating a URL for an integration](/help/generate-integration-url).  This is an unstable API expected to be used only by the Zulip web app. Please discuss in chat.zulip.org before using it.  **Changes**: New in Zulip 11.0 (feature level 403). Previously, these optional URL parameter options were included in the &#x60;config_options&#x60; object.  | [optional] 

## Methods

### NewRegisterQueueResponseRealmIncomingWebhookBotsItem

`func NewRegisterQueueResponseRealmIncomingWebhookBotsItem() *RegisterQueueResponseRealmIncomingWebhookBotsItem`

NewRegisterQueueResponseRealmIncomingWebhookBotsItem instantiates a new RegisterQueueResponseRealmIncomingWebhookBotsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponseRealmIncomingWebhookBotsItemWithDefaults

`func NewRegisterQueueResponseRealmIncomingWebhookBotsItemWithDefaults() *RegisterQueueResponseRealmIncomingWebhookBotsItem`

NewRegisterQueueResponseRealmIncomingWebhookBotsItemWithDefaults instantiates a new RegisterQueueResponseRealmIncomingWebhookBotsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAllEventTypes

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetAllEventTypes() []string`

GetAllEventTypes returns the AllEventTypes field if non-nil, zero value otherwise.

### GetAllEventTypesOk

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetAllEventTypesOk() (*[]string, bool)`

GetAllEventTypesOk returns a tuple with the AllEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEventTypes

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetAllEventTypes(v []string)`

SetAllEventTypes sets AllEventTypes field to given value.

### HasAllEventTypes

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) HasAllEventTypes() bool`

HasAllEventTypes returns a boolean if a field has been set.

### SetAllEventTypesNil

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetAllEventTypesNil(b bool)`

 SetAllEventTypesNil sets the value for AllEventTypes to be an explicit nil

### UnsetAllEventTypes
`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) UnsetAllEventTypes()`

UnsetAllEventTypes ensures that no value is present for AllEventTypes, not even an explicit nil
### GetConfigOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetConfigOptions() []WebhookConfigOptionInner`

GetConfigOptions returns the ConfigOptions field if non-nil, zero value otherwise.

### GetConfigOptionsOk

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetConfigOptionsOk() (*[]WebhookConfigOptionInner, bool)`

GetConfigOptionsOk returns a tuple with the ConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetConfigOptions(v []WebhookConfigOptionInner)`

SetConfigOptions sets ConfigOptions field to given value.

### HasConfigOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) HasConfigOptions() bool`

HasConfigOptions returns a boolean if a field has been set.

### GetUrlOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetUrlOptions() []WebhookUrlOptionInner`

GetUrlOptions returns the UrlOptions field if non-nil, zero value otherwise.

### GetUrlOptionsOk

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) GetUrlOptionsOk() (*[]WebhookUrlOptionInner, bool)`

GetUrlOptionsOk returns a tuple with the UrlOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) SetUrlOptions(v []WebhookUrlOptionInner)`

SetUrlOptions sets UrlOptions field to given value.

### HasUrlOptions

`func (o *RegisterQueueResponseRealmIncomingWebhookBotsItem) HasUrlOptions() bool`

HasUrlOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


