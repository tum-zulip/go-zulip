# UploadFile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Uri** | Pointer to **string** | The URL of the uploaded file. Alias of &#x60;url&#x60;.  **Changes**: Deprecated in Zulip 9.0 (feature level 272). The term \&quot;URI\&quot; is deprecated in [web standards](https://url.spec.whatwg.org/#goals).  | [optional] 
**Url** | Pointer to **string** | The URL of the uploaded file.  **Changes**: New in Zulip 9.0 (feature level 272). Previously, this property was only available under the legacy &#x60;uri&#x60; name.  | [optional] 
**Filename** | Pointer to **string** | The filename that Zulip stored the upload as. This usually differs from the basename of the URL when HTML escaping is required to generate a valid URL.  Clients generating a Markdown link to a newly uploaded file should do so by combining the &#x60;url&#x60; and &#x60;filename&#x60; fields in the response as follows: &#x60;[{filename}]({url})&#x60;, with care taken to clean &#x60;filename&#x60; of &#x60;[&#x60; and &#x60;]&#x60; characters that might break Markdown rendering.  **Changes**: New in Zulip 10.0 (feature level 285).  | [optional] 

## Methods

### NewUploadFile200Response

`func NewUploadFile200Response(result interface{}, msg interface{}, ) *UploadFile200Response`

NewUploadFile200Response instantiates a new UploadFile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFile200ResponseWithDefaults

`func NewUploadFile200ResponseWithDefaults() *UploadFile200Response`

NewUploadFile200ResponseWithDefaults instantiates a new UploadFile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UploadFile200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UploadFile200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UploadFile200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UploadFile200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UploadFile200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UploadFile200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UploadFile200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UploadFile200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UploadFile200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UploadFile200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UploadFile200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UploadFile200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UploadFile200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UploadFile200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *UploadFile200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *UploadFile200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetUri

`func (o *UploadFile200Response) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UploadFile200Response) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UploadFile200Response) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *UploadFile200Response) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUrl

`func (o *UploadFile200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFile200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFile200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadFile200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilename

`func (o *UploadFile200Response) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadFile200Response) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadFile200Response) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UploadFile200Response) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


