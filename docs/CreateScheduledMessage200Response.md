# CreateScheduledMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**ScheduledMessageId** | Pointer to **int32** | The unique ID of the scheduled message.  This is different from the unique ID that the message will have after it is sent.  | [optional] 

## Methods

### NewCreateScheduledMessage200Response

`func NewCreateScheduledMessage200Response(result interface{}, msg interface{}, ) *CreateScheduledMessage200Response`

NewCreateScheduledMessage200Response instantiates a new CreateScheduledMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduledMessage200ResponseWithDefaults

`func NewCreateScheduledMessage200ResponseWithDefaults() *CreateScheduledMessage200Response`

NewCreateScheduledMessage200ResponseWithDefaults instantiates a new CreateScheduledMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateScheduledMessage200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateScheduledMessage200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateScheduledMessage200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *CreateScheduledMessage200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CreateScheduledMessage200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *CreateScheduledMessage200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateScheduledMessage200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateScheduledMessage200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *CreateScheduledMessage200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *CreateScheduledMessage200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *CreateScheduledMessage200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *CreateScheduledMessage200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *CreateScheduledMessage200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *CreateScheduledMessage200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *CreateScheduledMessage200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *CreateScheduledMessage200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetScheduledMessageId

`func (o *CreateScheduledMessage200Response) GetScheduledMessageId() int32`

GetScheduledMessageId returns the ScheduledMessageId field if non-nil, zero value otherwise.

### GetScheduledMessageIdOk

`func (o *CreateScheduledMessage200Response) GetScheduledMessageIdOk() (*int32, bool)`

GetScheduledMessageIdOk returns a tuple with the ScheduledMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessageId

`func (o *CreateScheduledMessage200Response) SetScheduledMessageId(v int32)`

SetScheduledMessageId sets ScheduledMessageId field to given value.

### HasScheduledMessageId

`func (o *CreateScheduledMessage200Response) HasScheduledMessageId() bool`

HasScheduledMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


