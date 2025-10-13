# GetMessageHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**MessageHistory** | Pointer to [**[]Snapshot**](Snapshot.md) | A chronologically sorted, oldest to newest, array of &#x60;snapshot&#x60; objects, each one with the values of the message after the edit.  | [optional] 

## Methods

### NewGetMessageHistory200Response

`func NewGetMessageHistory200Response(result interface{}, msg interface{}, ) *GetMessageHistory200Response`

NewGetMessageHistory200Response instantiates a new GetMessageHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageHistory200ResponseWithDefaults

`func NewGetMessageHistory200ResponseWithDefaults() *GetMessageHistory200Response`

NewGetMessageHistory200ResponseWithDefaults instantiates a new GetMessageHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetMessageHistory200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMessageHistory200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMessageHistory200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetMessageHistory200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetMessageHistory200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetMessageHistory200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetMessageHistory200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetMessageHistory200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetMessageHistory200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetMessageHistory200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetMessageHistory200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetMessageHistory200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetMessageHistory200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetMessageHistory200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetMessageHistory200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetMessageHistory200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetMessageHistory

`func (o *GetMessageHistory200Response) GetMessageHistory() []Snapshot`

GetMessageHistory returns the MessageHistory field if non-nil, zero value otherwise.

### GetMessageHistoryOk

`func (o *GetMessageHistory200Response) GetMessageHistoryOk() (*[]Snapshot, bool)`

GetMessageHistoryOk returns a tuple with the MessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHistory

`func (o *GetMessageHistory200Response) SetMessageHistory(v []Snapshot)`

SetMessageHistory sets MessageHistory field to given value.

### HasMessageHistory

`func (o *GetMessageHistory200Response) HasMessageHistory() bool`

HasMessageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


