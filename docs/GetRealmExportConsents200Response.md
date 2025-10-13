# GetRealmExportConsents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**ExportConsents** | Pointer to [**[]GetRealmExportConsents200ResponseAllOfExportConsentsInner**](GetRealmExportConsents200ResponseAllOfExportConsentsInner.md) | An array of objects where each object contains a user ID and whether the user has consented for their private data to be exported.  | [optional] 

## Methods

### NewGetRealmExportConsents200Response

`func NewGetRealmExportConsents200Response(result interface{}, msg interface{}, ) *GetRealmExportConsents200Response`

NewGetRealmExportConsents200Response instantiates a new GetRealmExportConsents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRealmExportConsents200ResponseWithDefaults

`func NewGetRealmExportConsents200ResponseWithDefaults() *GetRealmExportConsents200Response`

NewGetRealmExportConsents200ResponseWithDefaults instantiates a new GetRealmExportConsents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetRealmExportConsents200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRealmExportConsents200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRealmExportConsents200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetRealmExportConsents200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetRealmExportConsents200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetRealmExportConsents200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRealmExportConsents200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRealmExportConsents200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetRealmExportConsents200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetRealmExportConsents200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetRealmExportConsents200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetRealmExportConsents200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetRealmExportConsents200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetRealmExportConsents200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetRealmExportConsents200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetRealmExportConsents200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetExportConsents

`func (o *GetRealmExportConsents200Response) GetExportConsents() []GetRealmExportConsents200ResponseAllOfExportConsentsInner`

GetExportConsents returns the ExportConsents field if non-nil, zero value otherwise.

### GetExportConsentsOk

`func (o *GetRealmExportConsents200Response) GetExportConsentsOk() (*[]GetRealmExportConsents200ResponseAllOfExportConsentsInner, bool)`

GetExportConsentsOk returns a tuple with the ExportConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportConsents

`func (o *GetRealmExportConsents200Response) SetExportConsents(v []GetRealmExportConsents200ResponseAllOfExportConsentsInner)`

SetExportConsents sets ExportConsents field to given value.

### HasExportConsents

`func (o *GetRealmExportConsents200Response) HasExportConsents() bool`

HasExportConsents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


