# CreateUserGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**GroupId** | **int32** | The unique ID of the created user group.  | 

## Methods

### NewCreateUserGroup200Response

`func NewCreateUserGroup200Response(result interface{}, msg interface{}, groupId int32, ) *CreateUserGroup200Response`

NewCreateUserGroup200Response instantiates a new CreateUserGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserGroup200ResponseWithDefaults

`func NewCreateUserGroup200ResponseWithDefaults() *CreateUserGroup200Response`

NewCreateUserGroup200ResponseWithDefaults instantiates a new CreateUserGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateUserGroup200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateUserGroup200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateUserGroup200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *CreateUserGroup200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CreateUserGroup200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *CreateUserGroup200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateUserGroup200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateUserGroup200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *CreateUserGroup200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *CreateUserGroup200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *CreateUserGroup200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *CreateUserGroup200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *CreateUserGroup200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *CreateUserGroup200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *CreateUserGroup200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *CreateUserGroup200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetGroupId

`func (o *CreateUserGroup200Response) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateUserGroup200Response) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateUserGroup200Response) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


