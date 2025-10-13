# GetScheduledMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**ScheduledMessages** | Pointer to [**[]ScheduledMessage**](ScheduledMessage.md) | Returns all of the current user&#39;s undelivered scheduled messages, ordered by &#x60;scheduled_delivery_timestamp&#x60; (ascending).  | [optional] 

## Methods

### NewGetScheduledMessages200Response

`func NewGetScheduledMessages200Response(result interface{}, msg interface{}, ) *GetScheduledMessages200Response`

NewGetScheduledMessages200Response instantiates a new GetScheduledMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetScheduledMessages200ResponseWithDefaults

`func NewGetScheduledMessages200ResponseWithDefaults() *GetScheduledMessages200Response`

NewGetScheduledMessages200ResponseWithDefaults instantiates a new GetScheduledMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetScheduledMessages200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetScheduledMessages200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetScheduledMessages200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetScheduledMessages200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetScheduledMessages200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetScheduledMessages200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetScheduledMessages200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetScheduledMessages200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetScheduledMessages200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetScheduledMessages200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetScheduledMessages200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetScheduledMessages200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetScheduledMessages200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetScheduledMessages200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetScheduledMessages200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetScheduledMessages200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetScheduledMessages

`func (o *GetScheduledMessages200Response) GetScheduledMessages() []ScheduledMessage`

GetScheduledMessages returns the ScheduledMessages field if non-nil, zero value otherwise.

### GetScheduledMessagesOk

`func (o *GetScheduledMessages200Response) GetScheduledMessagesOk() (*[]ScheduledMessage, bool)`

GetScheduledMessagesOk returns a tuple with the ScheduledMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMessages

`func (o *GetScheduledMessages200Response) SetScheduledMessages(v []ScheduledMessage)`

SetScheduledMessages sets ScheduledMessages field to given value.

### HasScheduledMessages

`func (o *GetScheduledMessages200Response) HasScheduledMessages() bool`

HasScheduledMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


