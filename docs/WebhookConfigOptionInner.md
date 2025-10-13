# WebhookConfigOptionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A key for the configuration option.  | [optional] 
**Label** | Pointer to **string** | A human-readable label of the configuration option.  | [optional] 
**Validator** | Pointer to **string** | The name of the validator function for the configuration option.  | [optional] 

## Methods

### NewWebhookConfigOptionInner

`func NewWebhookConfigOptionInner() *WebhookConfigOptionInner`

NewWebhookConfigOptionInner instantiates a new WebhookConfigOptionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigOptionInnerWithDefaults

`func NewWebhookConfigOptionInnerWithDefaults() *WebhookConfigOptionInner`

NewWebhookConfigOptionInnerWithDefaults instantiates a new WebhookConfigOptionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebhookConfigOptionInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookConfigOptionInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookConfigOptionInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebhookConfigOptionInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *WebhookConfigOptionInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WebhookConfigOptionInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WebhookConfigOptionInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WebhookConfigOptionInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValidator

`func (o *WebhookConfigOptionInner) GetValidator() string`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *WebhookConfigOptionInner) GetValidatorOk() (*string, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *WebhookConfigOptionInner) SetValidator(v string)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *WebhookConfigOptionInner) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


