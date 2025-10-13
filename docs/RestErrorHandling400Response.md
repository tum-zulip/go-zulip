# RestErrorHandling400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**VarName** | Pointer to **string** | It contains the information about the missing parameter.  | [optional] 
**Parameters** | Pointer to **string** | A string containing the parameters, separated by commas, that are incompatible.  | [optional] 

## Methods

### NewRestErrorHandling400Response

`func NewRestErrorHandling400Response(result interface{}, msg interface{}, code interface{}, ) *RestErrorHandling400Response`

NewRestErrorHandling400Response instantiates a new RestErrorHandling400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestErrorHandling400ResponseWithDefaults

`func NewRestErrorHandling400ResponseWithDefaults() *RestErrorHandling400Response`

NewRestErrorHandling400ResponseWithDefaults instantiates a new RestErrorHandling400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *RestErrorHandling400Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RestErrorHandling400Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RestErrorHandling400Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *RestErrorHandling400Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RestErrorHandling400Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *RestErrorHandling400Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RestErrorHandling400Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RestErrorHandling400Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *RestErrorHandling400Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *RestErrorHandling400Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *RestErrorHandling400Response) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RestErrorHandling400Response) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RestErrorHandling400Response) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *RestErrorHandling400Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *RestErrorHandling400Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetVarName

`func (o *RestErrorHandling400Response) GetVarName() string`

GetVarName returns the VarName field if non-nil, zero value otherwise.

### GetVarNameOk

`func (o *RestErrorHandling400Response) GetVarNameOk() (*string, bool)`

GetVarNameOk returns a tuple with the VarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarName

`func (o *RestErrorHandling400Response) SetVarName(v string)`

SetVarName sets VarName field to given value.

### HasVarName

`func (o *RestErrorHandling400Response) HasVarName() bool`

HasVarName returns a boolean if a field has been set.

### GetParameters

`func (o *RestErrorHandling400Response) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RestErrorHandling400Response) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RestErrorHandling400Response) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RestErrorHandling400Response) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


