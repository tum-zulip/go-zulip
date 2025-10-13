# CheckMessagesMatchNarrow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Messages** | Pointer to [**map[string]NarrowMatch**](NarrowMatch.md) | A dictionary with a key for each queried message that matches the narrow, with message IDs as keys and search rendering data as values.  | [optional] 

## Methods

### NewCheckMessagesMatchNarrow200Response

`func NewCheckMessagesMatchNarrow200Response(result interface{}, msg interface{}, ) *CheckMessagesMatchNarrow200Response`

NewCheckMessagesMatchNarrow200Response instantiates a new CheckMessagesMatchNarrow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckMessagesMatchNarrow200ResponseWithDefaults

`func NewCheckMessagesMatchNarrow200ResponseWithDefaults() *CheckMessagesMatchNarrow200Response`

NewCheckMessagesMatchNarrow200ResponseWithDefaults instantiates a new CheckMessagesMatchNarrow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CheckMessagesMatchNarrow200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CheckMessagesMatchNarrow200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CheckMessagesMatchNarrow200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *CheckMessagesMatchNarrow200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *CheckMessagesMatchNarrow200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *CheckMessagesMatchNarrow200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CheckMessagesMatchNarrow200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CheckMessagesMatchNarrow200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *CheckMessagesMatchNarrow200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *CheckMessagesMatchNarrow200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *CheckMessagesMatchNarrow200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *CheckMessagesMatchNarrow200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *CheckMessagesMatchNarrow200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *CheckMessagesMatchNarrow200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *CheckMessagesMatchNarrow200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *CheckMessagesMatchNarrow200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetMessages

`func (o *CheckMessagesMatchNarrow200Response) GetMessages() map[string]NarrowMatch`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CheckMessagesMatchNarrow200Response) GetMessagesOk() (*map[string]NarrowMatch, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CheckMessagesMatchNarrow200Response) SetMessages(v map[string]NarrowMatch)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CheckMessagesMatchNarrow200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


