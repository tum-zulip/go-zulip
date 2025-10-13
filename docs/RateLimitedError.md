# RateLimitedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**RetryAfter** | Pointer to **int32** | How many seconds the client must wait before making additional requests.  | [optional] 

## Methods

### NewRateLimitedError

`func NewRateLimitedError(result interface{}, msg interface{}, code interface{}, ) *RateLimitedError`

NewRateLimitedError instantiates a new RateLimitedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitedErrorWithDefaults

`func NewRateLimitedErrorWithDefaults() *RateLimitedError`

NewRateLimitedErrorWithDefaults instantiates a new RateLimitedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *RateLimitedError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RateLimitedError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RateLimitedError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *RateLimitedError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RateLimitedError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *RateLimitedError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RateLimitedError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RateLimitedError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *RateLimitedError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *RateLimitedError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *RateLimitedError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RateLimitedError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RateLimitedError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *RateLimitedError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *RateLimitedError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetRetryAfter

`func (o *RateLimitedError) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *RateLimitedError) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *RateLimitedError) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *RateLimitedError) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


