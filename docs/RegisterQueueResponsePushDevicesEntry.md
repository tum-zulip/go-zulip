# RegisterQueueResponsePushDevicesEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The push account&#39;s registration status. Either &#x60;\&quot;active\&quot;&#x60;, &#x60;\&quot;pending\&quot;&#x60;, or &#x60;\&quot;failed\&quot;&#x60;.  | [optional] 
**ErrorCode** | Pointer to **NullableString** | If the status is &#x60;\&quot;failed\&quot;&#x60;, a [Zulip API error code](/api/rest-error-handling) indicating the type of failure that occurred.  The following error codes have recommended client behavior:  - &#x60;\&quot;INVALID_BOUNCER_PUBLIC_KEY\&quot;&#x60; - Inform the user to update app. - &#x60;\&quot;REQUEST_EXPIRED&#x60; - Retry with a fresh payload.  | [optional] 

## Methods

### NewRegisterQueueResponsePushDevicesEntry

`func NewRegisterQueueResponsePushDevicesEntry() *RegisterQueueResponsePushDevicesEntry`

NewRegisterQueueResponsePushDevicesEntry instantiates a new RegisterQueueResponsePushDevicesEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterQueueResponsePushDevicesEntryWithDefaults

`func NewRegisterQueueResponsePushDevicesEntryWithDefaults() *RegisterQueueResponsePushDevicesEntry`

NewRegisterQueueResponsePushDevicesEntryWithDefaults instantiates a new RegisterQueueResponsePushDevicesEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RegisterQueueResponsePushDevicesEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterQueueResponsePushDevicesEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterQueueResponsePushDevicesEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisterQueueResponsePushDevicesEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *RegisterQueueResponsePushDevicesEntry) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RegisterQueueResponsePushDevicesEntry) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RegisterQueueResponsePushDevicesEntry) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RegisterQueueResponsePushDevicesEntry) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *RegisterQueueResponsePushDevicesEntry) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *RegisterQueueResponsePushDevicesEntry) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


