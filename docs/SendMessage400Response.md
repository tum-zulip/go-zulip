# SendMessage400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**Stream** | Pointer to **string** | The name of the channel that could not be found.  | [optional] 

## Methods

### NewSendMessage400Response

`func NewSendMessage400Response(result interface{}, msg interface{}, code interface{}, ) *SendMessage400Response`

NewSendMessage400Response instantiates a new SendMessage400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessage400ResponseWithDefaults

`func NewSendMessage400ResponseWithDefaults() *SendMessage400Response`

NewSendMessage400ResponseWithDefaults instantiates a new SendMessage400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SendMessage400Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SendMessage400Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SendMessage400Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *SendMessage400Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *SendMessage400Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *SendMessage400Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SendMessage400Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SendMessage400Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *SendMessage400Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *SendMessage400Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *SendMessage400Response) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SendMessage400Response) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SendMessage400Response) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *SendMessage400Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SendMessage400Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetStream

`func (o *SendMessage400Response) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *SendMessage400Response) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *SendMessage400Response) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *SendMessage400Response) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


