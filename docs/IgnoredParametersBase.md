# IgnoredParametersBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Msg** | **string** |  | 
**IgnoredParametersUnsupported** | Pointer to **[]string** | An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling](/api/rest-error-handling#ignored-parameters) documentation for details on this and its change history.  | [optional] 

## Methods

### NewIgnoredParametersBase

`func NewIgnoredParametersBase(result string, msg string, ) *IgnoredParametersBase`

NewIgnoredParametersBase instantiates a new IgnoredParametersBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgnoredParametersBaseWithDefaults

`func NewIgnoredParametersBaseWithDefaults() *IgnoredParametersBase`

NewIgnoredParametersBaseWithDefaults instantiates a new IgnoredParametersBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *IgnoredParametersBase) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *IgnoredParametersBase) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *IgnoredParametersBase) SetResult(v string)`

SetResult sets Result field to given value.


### GetMsg

`func (o *IgnoredParametersBase) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *IgnoredParametersBase) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *IgnoredParametersBase) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetIgnoredParametersUnsupported

`func (o *IgnoredParametersBase) GetIgnoredParametersUnsupported() []string`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *IgnoredParametersBase) GetIgnoredParametersUnsupportedOk() (*[]string, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *IgnoredParametersBase) SetIgnoredParametersUnsupported(v []string)`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *IgnoredParametersBase) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


