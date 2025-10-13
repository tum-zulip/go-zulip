# GetUserGroupSubgroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Subgroups** | Pointer to **[]int32** | A list containing the IDs of subgroups of the user group.  | [optional] 

## Methods

### NewGetUserGroupSubgroups200Response

`func NewGetUserGroupSubgroups200Response(result interface{}, msg interface{}, ) *GetUserGroupSubgroups200Response`

NewGetUserGroupSubgroups200Response instantiates a new GetUserGroupSubgroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGroupSubgroups200ResponseWithDefaults

`func NewGetUserGroupSubgroups200ResponseWithDefaults() *GetUserGroupSubgroups200Response`

NewGetUserGroupSubgroups200ResponseWithDefaults instantiates a new GetUserGroupSubgroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetUserGroupSubgroups200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUserGroupSubgroups200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUserGroupSubgroups200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetUserGroupSubgroups200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetUserGroupSubgroups200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetUserGroupSubgroups200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetUserGroupSubgroups200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetUserGroupSubgroups200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetUserGroupSubgroups200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetUserGroupSubgroups200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetUserGroupSubgroups200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetUserGroupSubgroups200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetUserGroupSubgroups200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetUserGroupSubgroups200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetUserGroupSubgroups200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetUserGroupSubgroups200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetSubgroups

`func (o *GetUserGroupSubgroups200Response) GetSubgroups() []int32`

GetSubgroups returns the Subgroups field if non-nil, zero value otherwise.

### GetSubgroupsOk

`func (o *GetUserGroupSubgroups200Response) GetSubgroupsOk() (*[]int32, bool)`

GetSubgroupsOk returns a tuple with the Subgroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgroups

`func (o *GetUserGroupSubgroups200Response) SetSubgroups(v []int32)`

SetSubgroups sets Subgroups field to given value.

### HasSubgroups

`func (o *GetUserGroupSubgroups200Response) HasSubgroups() bool`

HasSubgroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


