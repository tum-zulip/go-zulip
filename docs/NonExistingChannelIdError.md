# NonExistingChannelIdError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**StreamId** | Pointer to **int32** | The channel ID that could not be found.  | [optional] 

## Methods

### NewNonExistingChannelIdError

`func NewNonExistingChannelIdError(result interface{}, msg interface{}, code interface{}, ) *NonExistingChannelIdError`

NewNonExistingChannelIdError instantiates a new NonExistingChannelIdError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonExistingChannelIdErrorWithDefaults

`func NewNonExistingChannelIdErrorWithDefaults() *NonExistingChannelIdError`

NewNonExistingChannelIdErrorWithDefaults instantiates a new NonExistingChannelIdError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *NonExistingChannelIdError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *NonExistingChannelIdError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *NonExistingChannelIdError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *NonExistingChannelIdError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *NonExistingChannelIdError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *NonExistingChannelIdError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *NonExistingChannelIdError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *NonExistingChannelIdError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *NonExistingChannelIdError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *NonExistingChannelIdError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *NonExistingChannelIdError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NonExistingChannelIdError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NonExistingChannelIdError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *NonExistingChannelIdError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *NonExistingChannelIdError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetStreamId

`func (o *NonExistingChannelIdError) GetStreamId() int32`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *NonExistingChannelIdError) GetStreamIdOk() (*int32, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *NonExistingChannelIdError) SetStreamId(v int32)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *NonExistingChannelIdError) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


