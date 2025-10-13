# MissingArgumentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**VarName** | Pointer to **string** | It contains the information about the missing parameter.  | [optional] 

## Methods

### NewMissingArgumentError

`func NewMissingArgumentError(result interface{}, msg interface{}, code interface{}, ) *MissingArgumentError`

NewMissingArgumentError instantiates a new MissingArgumentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissingArgumentErrorWithDefaults

`func NewMissingArgumentErrorWithDefaults() *MissingArgumentError`

NewMissingArgumentErrorWithDefaults instantiates a new MissingArgumentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *MissingArgumentError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MissingArgumentError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MissingArgumentError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *MissingArgumentError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *MissingArgumentError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *MissingArgumentError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *MissingArgumentError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *MissingArgumentError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *MissingArgumentError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *MissingArgumentError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *MissingArgumentError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MissingArgumentError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MissingArgumentError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *MissingArgumentError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MissingArgumentError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetVarName

`func (o *MissingArgumentError) GetVarName() string`

GetVarName returns the VarName field if non-nil, zero value otherwise.

### GetVarNameOk

`func (o *MissingArgumentError) GetVarNameOk() (*string, bool)`

GetVarNameOk returns a tuple with the VarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarName

`func (o *MissingArgumentError) SetVarName(v string)`

SetVarName sets VarName field to given value.

### HasVarName

`func (o *MissingArgumentError) HasVarName() bool`

HasVarName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


