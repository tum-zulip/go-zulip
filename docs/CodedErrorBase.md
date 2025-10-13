# CodedErrorBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Msg** | **string** |  | 
**Code** | **string** | A string that identifies the error.  | 

## Methods

### NewCodedErrorBase

`func NewCodedErrorBase(result string, msg string, code string, ) *CodedErrorBase`

NewCodedErrorBase instantiates a new CodedErrorBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodedErrorBaseWithDefaults

`func NewCodedErrorBaseWithDefaults() *CodedErrorBase`

NewCodedErrorBaseWithDefaults instantiates a new CodedErrorBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CodedErrorBase) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CodedErrorBase) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CodedErrorBase) SetResult(v string)`

SetResult sets Result field to given value.


### GetMsg

`func (o *CodedErrorBase) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CodedErrorBase) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CodedErrorBase) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCode

`func (o *CodedErrorBase) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CodedErrorBase) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CodedErrorBase) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


