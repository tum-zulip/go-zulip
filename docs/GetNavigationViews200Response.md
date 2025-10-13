# GetNavigationViews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **[]string** | An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling](/api/rest-error-handling#ignored-parameters) documentation for details on this and its change history.  | [optional] 
**NavigationViews** | Pointer to [**[]NavigationView**](NavigationView.md) | An array of dictionaries containing the user&#39;s navigation views.  | [optional] 

## Methods

### NewGetNavigationViews200Response

`func NewGetNavigationViews200Response(result interface{}, msg interface{}, ) *GetNavigationViews200Response`

NewGetNavigationViews200Response instantiates a new GetNavigationViews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNavigationViews200ResponseWithDefaults

`func NewGetNavigationViews200ResponseWithDefaults() *GetNavigationViews200Response`

NewGetNavigationViews200ResponseWithDefaults instantiates a new GetNavigationViews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetNavigationViews200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNavigationViews200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNavigationViews200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetNavigationViews200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetNavigationViews200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetNavigationViews200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetNavigationViews200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetNavigationViews200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetNavigationViews200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetNavigationViews200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetNavigationViews200Response) GetIgnoredParametersUnsupported() []string`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetNavigationViews200Response) GetIgnoredParametersUnsupportedOk() (*[]string, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetNavigationViews200Response) SetIgnoredParametersUnsupported(v []string)`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetNavigationViews200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### GetNavigationViews

`func (o *GetNavigationViews200Response) GetNavigationViews() []NavigationView`

GetNavigationViews returns the NavigationViews field if non-nil, zero value otherwise.

### GetNavigationViewsOk

`func (o *GetNavigationViews200Response) GetNavigationViewsOk() (*[]NavigationView, bool)`

GetNavigationViewsOk returns a tuple with the NavigationViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationViews

`func (o *GetNavigationViews200Response) SetNavigationViews(v []NavigationView)`

SetNavigationViews sets NavigationViews field to given value.

### HasNavigationViews

`func (o *GetNavigationViews200Response) HasNavigationViews() bool`

HasNavigationViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


