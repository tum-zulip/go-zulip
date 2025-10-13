# BadEventQueueIdError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**QueueId** | Pointer to **string** | The string that identifies the invalid event queue.  | [optional] 

## Methods

### NewBadEventQueueIdError

`func NewBadEventQueueIdError(result interface{}, msg interface{}, code interface{}, ) *BadEventQueueIdError`

NewBadEventQueueIdError instantiates a new BadEventQueueIdError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadEventQueueIdErrorWithDefaults

`func NewBadEventQueueIdErrorWithDefaults() *BadEventQueueIdError`

NewBadEventQueueIdErrorWithDefaults instantiates a new BadEventQueueIdError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *BadEventQueueIdError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BadEventQueueIdError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BadEventQueueIdError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *BadEventQueueIdError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *BadEventQueueIdError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *BadEventQueueIdError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *BadEventQueueIdError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *BadEventQueueIdError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *BadEventQueueIdError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *BadEventQueueIdError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *BadEventQueueIdError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadEventQueueIdError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadEventQueueIdError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *BadEventQueueIdError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BadEventQueueIdError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetQueueId

`func (o *BadEventQueueIdError) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *BadEventQueueIdError) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *BadEventQueueIdError) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *BadEventQueueIdError) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


