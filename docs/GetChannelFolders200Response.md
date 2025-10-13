# GetChannelFolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**ChannelFolders** | Pointer to [**[]ChannelFolder**](ChannelFolder.md) | A list of channel folder objects.  | [optional] 

## Methods

### NewGetChannelFolders200Response

`func NewGetChannelFolders200Response(result interface{}, msg interface{}, ) *GetChannelFolders200Response`

NewGetChannelFolders200Response instantiates a new GetChannelFolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannelFolders200ResponseWithDefaults

`func NewGetChannelFolders200ResponseWithDefaults() *GetChannelFolders200Response`

NewGetChannelFolders200ResponseWithDefaults instantiates a new GetChannelFolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetChannelFolders200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetChannelFolders200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetChannelFolders200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetChannelFolders200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetChannelFolders200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetChannelFolders200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetChannelFolders200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetChannelFolders200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetChannelFolders200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetChannelFolders200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetChannelFolders200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetChannelFolders200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetChannelFolders200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetChannelFolders200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetChannelFolders200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetChannelFolders200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetChannelFolders

`func (o *GetChannelFolders200Response) GetChannelFolders() []ChannelFolder`

GetChannelFolders returns the ChannelFolders field if non-nil, zero value otherwise.

### GetChannelFoldersOk

`func (o *GetChannelFolders200Response) GetChannelFoldersOk() (*[]ChannelFolder, bool)`

GetChannelFoldersOk returns a tuple with the ChannelFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolders

`func (o *GetChannelFolders200Response) SetChannelFolders(v []ChannelFolder)`

SetChannelFolders sets ChannelFolders field to given value.

### HasChannelFolders

`func (o *GetChannelFolders200Response) HasChannelFolders() bool`

HasChannelFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


