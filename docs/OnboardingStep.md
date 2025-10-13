# OnboardingStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the onboarding step. Valid value is &#x60;\&quot;one_time_notice\&quot;&#x60;.  **Changes**: Removed type &#x60;\&quot;hotspot\&quot;&#x60; in Zulip 9.0 (feature level 259).  New in Zulip 8.0 (feature level 233).  | [optional] 
**Name** | Pointer to **string** | The name of the onboarding step.  | [optional] 

## Methods

### NewOnboardingStep

`func NewOnboardingStep() *OnboardingStep`

NewOnboardingStep instantiates a new OnboardingStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingStepWithDefaults

`func NewOnboardingStepWithDefaults() *OnboardingStep`

NewOnboardingStepWithDefaults instantiates a new OnboardingStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnboardingStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingStep) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnboardingStep) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *OnboardingStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardingStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardingStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnboardingStep) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


