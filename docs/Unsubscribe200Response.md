# Unsubscribe200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**NotRemoved** | Pointer to **[]string** | A list of the names of channels that the user is already unsubscribed from, and hence doesn&#39;t need to be unsubscribed.  | [optional] 
**Removed** | Pointer to **[]string** | A list of the names of channels which were unsubscribed from as a result of the query.  | [optional] 

## Methods

### NewUnsubscribe200Response

`func NewUnsubscribe200Response(result interface{}, msg interface{}, ) *Unsubscribe200Response`

NewUnsubscribe200Response instantiates a new Unsubscribe200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscribe200ResponseWithDefaults

`func NewUnsubscribe200ResponseWithDefaults() *Unsubscribe200Response`

NewUnsubscribe200ResponseWithDefaults instantiates a new Unsubscribe200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *Unsubscribe200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Unsubscribe200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Unsubscribe200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *Unsubscribe200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *Unsubscribe200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *Unsubscribe200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *Unsubscribe200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *Unsubscribe200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *Unsubscribe200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *Unsubscribe200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *Unsubscribe200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *Unsubscribe200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *Unsubscribe200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *Unsubscribe200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *Unsubscribe200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *Unsubscribe200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetNotRemoved

`func (o *Unsubscribe200Response) GetNotRemoved() []string`

GetNotRemoved returns the NotRemoved field if non-nil, zero value otherwise.

### GetNotRemovedOk

`func (o *Unsubscribe200Response) GetNotRemovedOk() (*[]string, bool)`

GetNotRemovedOk returns a tuple with the NotRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotRemoved

`func (o *Unsubscribe200Response) SetNotRemoved(v []string)`

SetNotRemoved sets NotRemoved field to given value.

### HasNotRemoved

`func (o *Unsubscribe200Response) HasNotRemoved() bool`

HasNotRemoved returns a boolean if a field has been set.

### GetRemoved

`func (o *Unsubscribe200Response) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *Unsubscribe200Response) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *Unsubscribe200Response) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *Unsubscribe200Response) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


