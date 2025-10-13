# PushDeviceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PushAccountId** | Pointer to **string** | The push account ID for this client registration.  See [&#x60;POST /mobile_push/register&#x60;](zulip.com/api/register-push-device for details on push account IDs.  | [optional] 
**Status** | Pointer to **string** | The updated registration status. Will be &#x60;\&quot;active\&quot;&#x60;, &#x60;\&quot;failed\&quot;&#x60;, or &#x60;\&quot;pending\&quot;&#x60;.  | [optional] 
**ErrorCode** | Pointer to **NullableString** | If the status is &#x60;\&quot;failed\&quot;&#x60;, a [Zulip API error code](zulip.com/api/rest-error-handling indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - &#x60;\&quot;INVALID_BOUNCER_PUBLIC_KEY\&quot;&#x60; - Inform the user to update app. - &#x60;\&quot;REQUEST_EXPIRED&#x60; - Retry with a fresh payload.   If the status is \&quot;failed\&quot;, an error code explaining the failure.  | [optional] 

## Methods

### NewEventEnvelopeOneOf24

`func NewEventEnvelopeOneOf24() *PushDeviceEvent`

NewEventEnvelopeOneOf24 instantiates a new PushDeviceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEnvelopeOneOf24WithDefaults

`func NewEventEnvelopeOneOf24WithDefaults() *PushDeviceEvent`

NewEventEnvelopeOneOf24WithDefaults instantiates a new PushDeviceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushDeviceEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushDeviceEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushDeviceEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PushDeviceEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PushDeviceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PushDeviceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PushDeviceEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PushDeviceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPushAccountId

`func (o *PushDeviceEvent) GetPushAccountId() string`

GetPushAccountId returns the PushAccountId field if non-nil, zero value otherwise.

### GetPushAccountIdOk

`func (o *PushDeviceEvent) GetPushAccountIdOk() (*string, bool)`

GetPushAccountIdOk returns a tuple with the PushAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushAccountId

`func (o *PushDeviceEvent) SetPushAccountId(v string)`

SetPushAccountId sets PushAccountId field to given value.

### HasPushAccountId

`func (o *PushDeviceEvent) HasPushAccountId() bool`

HasPushAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *PushDeviceEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PushDeviceEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PushDeviceEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PushDeviceEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *PushDeviceEvent) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PushDeviceEvent) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PushDeviceEvent) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PushDeviceEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *PushDeviceEvent) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *PushDeviceEvent) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


