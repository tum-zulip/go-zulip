# WebhookUrlOptionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The parameter variable to encode the users input for this option in the integrations webhook URL.  | [optional] 
**Label** | Pointer to **string** | A human-readable label of the url option.  | [optional] 
**Validator** | Pointer to **string** | The name of the validator function for the configuration option.  | [optional] 

## Methods

### NewWebhookUrlOptionInner

`func NewWebhookUrlOptionInner() *WebhookUrlOptionInner`

NewWebhookUrlOptionInner instantiates a new WebhookUrlOptionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUrlOptionInnerWithDefaults

`func NewWebhookUrlOptionInnerWithDefaults() *WebhookUrlOptionInner`

NewWebhookUrlOptionInnerWithDefaults instantiates a new WebhookUrlOptionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebhookUrlOptionInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookUrlOptionInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookUrlOptionInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebhookUrlOptionInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *WebhookUrlOptionInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WebhookUrlOptionInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WebhookUrlOptionInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WebhookUrlOptionInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValidator

`func (o *WebhookUrlOptionInner) GetValidator() string`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *WebhookUrlOptionInner) GetValidatorOk() (*string, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *WebhookUrlOptionInner) SetValidator(v string)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *WebhookUrlOptionInner) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


