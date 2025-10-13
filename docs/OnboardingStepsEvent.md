# OnboardingStepsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**OnboardingSteps** | Pointer to [**[]OnboardingStep**](OnboardingStep.md) | An array of dictionaries where each dictionary contains details about a single onboarding step.  **Changes**: Before Zulip 8.0 (feature level 233), this array was named &#x60;hotspots&#x60;. Prior to this feature level, one-time notice onboarding steps were not supported, and the &#x60;type&#x60; field in these objects did not exist as all onboarding steps were implicitly hotspots.  | [optional] 

## Methods

### NewEventEnvelopeOneOf35

`func NewEventEnvelopeOneOf35() *OnboardingStepsEvent`

NewEventEnvelopeOneOf35 instantiates a new OnboardingStepsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf35WithDefaults

`func NewEventEnvelopeOneOf35WithDefaults() *OnboardingStepsEvent`

NewEventEnvelopeOneOf35WithDefaults instantiates a new OnboardingStepsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnboardingStepsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingStepsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingStepsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OnboardingStepsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *OnboardingStepsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingStepsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingStepsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnboardingStepsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnboardingSteps

`func (o *OnboardingStepsEvent) GetOnboardingSteps() []OnboardingStep`

GetOnboardingSteps returns the OnboardingSteps field if non-nil, zero value otherwise.

### GetOnboardingStepsOk

`func (o *OnboardingStepsEvent) GetOnboardingStepsOk() (*[]OnboardingStep, bool)`

GetOnboardingStepsOk returns a tuple with the OnboardingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingSteps

`func (o *OnboardingStepsEvent) SetOnboardingSteps(v []OnboardingStep)`

SetOnboardingSteps sets OnboardingSteps field to given value.

### HasOnboardingSteps

`func (o *OnboardingStepsEvent) HasOnboardingSteps() bool`

HasOnboardingSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


