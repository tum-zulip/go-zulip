# GetLinkifiers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Linkifiers** | Pointer to [**[]GetLinkifiers200ResponseAllOfLinkifiersInner**](GetLinkifiers200ResponseAllOfLinkifiersInner.md) | An ordered array of objects, where each object describes a linkifier.  Clients should always process linkifiers in the order given; this is important if the realm has linkifiers with overlapping patterns. The order can be modified using [&#x60;PATCH /realm/linkifiers&#x60;](/api/reorder-linkifiers).  | [optional] 

## Methods

### NewGetLinkifiers200Response

`func NewGetLinkifiers200Response(result interface{}, msg interface{}, ) *GetLinkifiers200Response`

NewGetLinkifiers200Response instantiates a new GetLinkifiers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkifiers200ResponseWithDefaults

`func NewGetLinkifiers200ResponseWithDefaults() *GetLinkifiers200Response`

NewGetLinkifiers200ResponseWithDefaults instantiates a new GetLinkifiers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetLinkifiers200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetLinkifiers200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetLinkifiers200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetLinkifiers200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetLinkifiers200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetLinkifiers200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetLinkifiers200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetLinkifiers200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetLinkifiers200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetLinkifiers200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetLinkifiers200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetLinkifiers200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetLinkifiers200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetLinkifiers200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetLinkifiers200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetLinkifiers200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetLinkifiers

`func (o *GetLinkifiers200Response) GetLinkifiers() []GetLinkifiers200ResponseAllOfLinkifiersInner`

GetLinkifiers returns the Linkifiers field if non-nil, zero value otherwise.

### GetLinkifiersOk

`func (o *GetLinkifiers200Response) GetLinkifiersOk() (*[]GetLinkifiers200ResponseAllOfLinkifiersInner, bool)`

GetLinkifiersOk returns a tuple with the Linkifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkifiers

`func (o *GetLinkifiers200Response) SetLinkifiers(v []GetLinkifiers200ResponseAllOfLinkifiersInner)`

SetLinkifiers sets Linkifiers field to given value.

### HasLinkifiers

`func (o *GetLinkifiers200Response) HasLinkifiers() bool`

HasLinkifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


