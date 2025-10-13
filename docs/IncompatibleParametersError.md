# IncompatibleParametersError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**Parameters** | Pointer to **string** | A string containing the parameters, separated by commas, that are incompatible.  | [optional] 

## Methods

### NewIncompatibleParametersError

`func NewIncompatibleParametersError(result interface{}, msg interface{}, code interface{}, ) *IncompatibleParametersError`

NewIncompatibleParametersError instantiates a new IncompatibleParametersError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncompatibleParametersErrorWithDefaults

`func NewIncompatibleParametersErrorWithDefaults() *IncompatibleParametersError`

NewIncompatibleParametersErrorWithDefaults instantiates a new IncompatibleParametersError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *IncompatibleParametersError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *IncompatibleParametersError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *IncompatibleParametersError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *IncompatibleParametersError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *IncompatibleParametersError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *IncompatibleParametersError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *IncompatibleParametersError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *IncompatibleParametersError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *IncompatibleParametersError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *IncompatibleParametersError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *IncompatibleParametersError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IncompatibleParametersError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IncompatibleParametersError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *IncompatibleParametersError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *IncompatibleParametersError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetParameters

`func (o *IncompatibleParametersError) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IncompatibleParametersError) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IncompatibleParametersError) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IncompatibleParametersError) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


