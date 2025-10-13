# GetPresence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**ServerTimestamp** | Pointer to **float32** | The time when the server fetched the &#x60;presences&#x60; data included in the response.  | [optional] 
**Presences** | Pointer to [**map[string]map[string]LegacyPresenceFormat**](map.md) | A dictionary where each entry describes the presence details of a user in the Zulip organization.  | [optional] 

## Methods

### NewGetPresence200Response

`func NewGetPresence200Response(result interface{}, msg interface{}, ) *GetPresence200Response`

NewGetPresence200Response instantiates a new GetPresence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPresence200ResponseWithDefaults

`func NewGetPresence200ResponseWithDefaults() *GetPresence200Response`

NewGetPresence200ResponseWithDefaults instantiates a new GetPresence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetPresence200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetPresence200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetPresence200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetPresence200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetPresence200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetPresence200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetPresence200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetPresence200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetPresence200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetPresence200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetPresence200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetPresence200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetPresence200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetPresence200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetPresence200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetPresence200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetServerTimestamp

`func (o *GetPresence200Response) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *GetPresence200Response) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *GetPresence200Response) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *GetPresence200Response) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetPresences

`func (o *GetPresence200Response) GetPresences() map[string]map[string]LegacyPresenceFormat`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *GetPresence200Response) GetPresencesOk() (*map[string]map[string]LegacyPresenceFormat, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *GetPresence200Response) SetPresences(v map[string]map[string]LegacyPresenceFormat)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *GetPresence200Response) HasPresences() bool`

HasPresences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


