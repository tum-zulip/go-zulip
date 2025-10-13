# ZulipOutgoingWebhooks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotEmail** | Pointer to **string** | Email of the bot user.  | [optional] 
**BotFullName** | Pointer to **string** | The full name of the bot user.  | [optional] 
**Data** | Pointer to **string** | The message content, in raw [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format (not rendered to HTML).  | [optional] 
**Trigger** | Pointer to **string** | What aspect of the message triggered the outgoing webhook notification. Possible values include &#x60;direct_message&#x60; and &#x60;mention&#x60;.  **Changes**: In Zulip 8.0 (feature level 201), renamed the trigger &#x60;private_message&#x60; to &#x60;direct_message&#x60;.  | [optional] 
**Token** | Pointer to **string** | A string of alphanumeric characters that can be used to authenticate the webhook request (each bot user uses a fixed token). You can get the token used by a given outgoing webhook bot in the &#x60;zuliprc&#x60; file downloaded when creating the bot.  | [optional] 
**Message** | Pointer to [**ZulipOutgoingWebhooks200ResponseMessage**](ZulipOutgoingWebhooks200ResponseMessage.md) |  | [optional] 

## Methods

### NewZulipOutgoingWebhooks200Response

`func NewZulipOutgoingWebhooks200Response() *ZulipOutgoingWebhooks200Response`

NewZulipOutgoingWebhooks200Response instantiates a new ZulipOutgoingWebhooks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZulipOutgoingWebhooks200ResponseWithDefaults

`func NewZulipOutgoingWebhooks200ResponseWithDefaults() *ZulipOutgoingWebhooks200Response`

NewZulipOutgoingWebhooks200ResponseWithDefaults instantiates a new ZulipOutgoingWebhooks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotEmail

`func (o *ZulipOutgoingWebhooks200Response) GetBotEmail() string`

GetBotEmail returns the BotEmail field if non-nil, zero value otherwise.

### GetBotEmailOk

`func (o *ZulipOutgoingWebhooks200Response) GetBotEmailOk() (*string, bool)`

GetBotEmailOk returns a tuple with the BotEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEmail

`func (o *ZulipOutgoingWebhooks200Response) SetBotEmail(v string)`

SetBotEmail sets BotEmail field to given value.

### HasBotEmail

`func (o *ZulipOutgoingWebhooks200Response) HasBotEmail() bool`

HasBotEmail returns a boolean if a field has been set.

### GetBotFullName

`func (o *ZulipOutgoingWebhooks200Response) GetBotFullName() string`

GetBotFullName returns the BotFullName field if non-nil, zero value otherwise.

### GetBotFullNameOk

`func (o *ZulipOutgoingWebhooks200Response) GetBotFullNameOk() (*string, bool)`

GetBotFullNameOk returns a tuple with the BotFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotFullName

`func (o *ZulipOutgoingWebhooks200Response) SetBotFullName(v string)`

SetBotFullName sets BotFullName field to given value.

### HasBotFullName

`func (o *ZulipOutgoingWebhooks200Response) HasBotFullName() bool`

HasBotFullName returns a boolean if a field has been set.

### GetData

`func (o *ZulipOutgoingWebhooks200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ZulipOutgoingWebhooks200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ZulipOutgoingWebhooks200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ZulipOutgoingWebhooks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTrigger

`func (o *ZulipOutgoingWebhooks200Response) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ZulipOutgoingWebhooks200Response) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ZulipOutgoingWebhooks200Response) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ZulipOutgoingWebhooks200Response) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetToken

`func (o *ZulipOutgoingWebhooks200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ZulipOutgoingWebhooks200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ZulipOutgoingWebhooks200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ZulipOutgoingWebhooks200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetMessage

`func (o *ZulipOutgoingWebhooks200Response) GetMessage() ZulipOutgoingWebhooks200ResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ZulipOutgoingWebhooks200Response) GetMessageOk() (*ZulipOutgoingWebhooks200ResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ZulipOutgoingWebhooks200Response) SetMessage(v ZulipOutgoingWebhooks200ResponseMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ZulipOutgoingWebhooks200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


