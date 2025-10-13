# AddAlertWords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**AlertWords** | Pointer to **[]string** | An array of strings, where each string is an alert word (or phrase) configured by the user.  | [optional] 

## Methods

### NewAddAlertWords200Response

`func NewAddAlertWords200Response(result interface{}, msg interface{}, ) *AddAlertWords200Response`

NewAddAlertWords200Response instantiates a new AddAlertWords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAlertWords200ResponseWithDefaults

`func NewAddAlertWords200ResponseWithDefaults() *AddAlertWords200Response`

NewAddAlertWords200ResponseWithDefaults instantiates a new AddAlertWords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AddAlertWords200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AddAlertWords200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AddAlertWords200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *AddAlertWords200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *AddAlertWords200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *AddAlertWords200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AddAlertWords200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AddAlertWords200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *AddAlertWords200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *AddAlertWords200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *AddAlertWords200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *AddAlertWords200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *AddAlertWords200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *AddAlertWords200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *AddAlertWords200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *AddAlertWords200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetAlertWords

`func (o *AddAlertWords200Response) GetAlertWords() []string`

GetAlertWords returns the AlertWords field if non-nil, zero value otherwise.

### GetAlertWordsOk

`func (o *AddAlertWords200Response) GetAlertWordsOk() (*[]string, bool)`

GetAlertWordsOk returns a tuple with the AlertWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertWords

`func (o *AddAlertWords200Response) SetAlertWords(v []string)`

SetAlertWords sets AlertWords field to given value.

### HasAlertWords

`func (o *AddAlertWords200Response) HasAlertWords() bool`

HasAlertWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


