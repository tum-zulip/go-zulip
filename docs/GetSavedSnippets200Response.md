# GetSavedSnippets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**SavedSnippets** | Pointer to [**[]SavedSnippet**](SavedSnippet.md) | An array of dictionaries containing data on all of the current user&#39;s saved snippets.  | [optional] 

## Methods

### NewGetSavedSnippets200Response

`func NewGetSavedSnippets200Response(result interface{}, msg interface{}, ) *GetSavedSnippets200Response`

NewGetSavedSnippets200Response instantiates a new GetSavedSnippets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSavedSnippets200ResponseWithDefaults

`func NewGetSavedSnippets200ResponseWithDefaults() *GetSavedSnippets200Response`

NewGetSavedSnippets200ResponseWithDefaults instantiates a new GetSavedSnippets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetSavedSnippets200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSavedSnippets200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSavedSnippets200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetSavedSnippets200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetSavedSnippets200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetSavedSnippets200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetSavedSnippets200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetSavedSnippets200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetSavedSnippets200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetSavedSnippets200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetSavedSnippets200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetSavedSnippets200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetSavedSnippets200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetSavedSnippets200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetSavedSnippets200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetSavedSnippets200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetSavedSnippets

`func (o *GetSavedSnippets200Response) GetSavedSnippets() []SavedSnippet`

GetSavedSnippets returns the SavedSnippets field if non-nil, zero value otherwise.

### GetSavedSnippetsOk

`func (o *GetSavedSnippets200Response) GetSavedSnippetsOk() (*[]SavedSnippet, bool)`

GetSavedSnippetsOk returns a tuple with the SavedSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedSnippets

`func (o *GetSavedSnippets200Response) SetSavedSnippets(v []SavedSnippet)`

SetSavedSnippets sets SavedSnippets field to given value.

### HasSavedSnippets

`func (o *GetSavedSnippets200Response) HasSavedSnippets() bool`

HasSavedSnippets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


