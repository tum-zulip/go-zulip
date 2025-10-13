# GetUserStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to [**UserStatus**](UserStatus.md) | The status set by the user. Note that, if the user doesn&#39;t have a status currently set, then the returned dictionary will be empty as none of the keys listed below will be present.  | [optional] 

## Methods

### NewGetUserStatus200Response

`func NewGetUserStatus200Response(result interface{}, msg interface{}, ) *GetUserStatus200Response`

NewGetUserStatus200Response instantiates a new GetUserStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserStatus200ResponseWithDefaults

`func NewGetUserStatus200ResponseWithDefaults() *GetUserStatus200Response`

NewGetUserStatus200ResponseWithDefaults instantiates a new GetUserStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetUserStatus200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUserStatus200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUserStatus200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetUserStatus200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetUserStatus200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetUserStatus200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetUserStatus200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetUserStatus200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetUserStatus200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetUserStatus200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetUserStatus200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetUserStatus200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetUserStatus200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetUserStatus200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetUserStatus200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetUserStatus200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetStatus

`func (o *GetUserStatus200Response) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserStatus200Response) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserStatus200Response) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUserStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


