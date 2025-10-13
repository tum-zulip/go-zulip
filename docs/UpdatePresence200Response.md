# UpdatePresence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**PresenceLastUpdateId** | Pointer to **int32** | The identifier for the latest user presence data returned in the &#x60;presences&#x60; data of the response.  If a value was passed for &#x60;last_update_id&#x60;, then this is guaranteed to be equal to or greater than that value. If it is the same value, then that indicates to the client that there were no updates to previously received user presence data.  The client should then pass this value as the &#x60;last_update_id&#x60; parameter when it next queries this endpoint, in order to only receive new user presence data and avoid redundantly fetching already known information.  This will be &#x60;-1&#x60; if no value was passed for [&#x60;last_update_id&#x60;](#parameter-last_update_id) and no user presence data was returned by the server. This can happen, for example, if an organization has disabled presence.  **Changes**: New in Zulip 9.0 (feature level 263).  | [optional] 
**ServerTimestamp** | Pointer to **float32** | Only present if &#x60;ping_only&#x60; is &#x60;false&#x60;.  The time when the server fetched the &#x60;presences&#x60; data included in the response.  | [optional] 
**Presences** | Pointer to [**map[string]PresenceUpdateValue**](PresenceUpdateValue.md) | Only present if &#x60;ping_only&#x60; is &#x60;false&#x60;.  A dictionary where each entry describes the presence details of a user in the Zulip organization. Entries can be in either the modern presence format or the legacy presence format.  These entries will be the modern presence format when the &#x60;last_updated_id&#x60; parameter is passed, or when the deprecated &#x60;slim_presence&#x60; parameter is &#x60;true&#x60;.  If the deprecated &#x60;slim_presence&#x60; parameter is &#x60;false&#x60; and the &#x60;last_updated_id&#x60; parameter is omitted, the entries will be in the legacy presence API format.  **Note**: The legacy presence format should only be used when interacting with old servers. It will be removed as soon as doing so is practical.  | [optional] 
**ZephyrMirrorActive** | Pointer to **bool** | Legacy property for a part of the Zephyr mirroring system.  **Deprecated**. Clients should ignore this field.  | [optional] 

## Methods

### NewUpdatePresence200Response

`func NewUpdatePresence200Response(result interface{}, msg interface{}, ) *UpdatePresence200Response`

NewUpdatePresence200Response instantiates a new UpdatePresence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePresence200ResponseWithDefaults

`func NewUpdatePresence200ResponseWithDefaults() *UpdatePresence200Response`

NewUpdatePresence200ResponseWithDefaults instantiates a new UpdatePresence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdatePresence200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdatePresence200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdatePresence200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UpdatePresence200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdatePresence200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UpdatePresence200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdatePresence200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdatePresence200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UpdatePresence200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdatePresence200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UpdatePresence200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UpdatePresence200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UpdatePresence200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UpdatePresence200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *UpdatePresence200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *UpdatePresence200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetPresenceLastUpdateId

`func (o *UpdatePresence200Response) GetPresenceLastUpdateId() int32`

GetPresenceLastUpdateId returns the PresenceLastUpdateId field if non-nil, zero value otherwise.

### GetPresenceLastUpdateIdOk

`func (o *UpdatePresence200Response) GetPresenceLastUpdateIdOk() (*int32, bool)`

GetPresenceLastUpdateIdOk returns a tuple with the PresenceLastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceLastUpdateId

`func (o *UpdatePresence200Response) SetPresenceLastUpdateId(v int32)`

SetPresenceLastUpdateId sets PresenceLastUpdateId field to given value.

### HasPresenceLastUpdateId

`func (o *UpdatePresence200Response) HasPresenceLastUpdateId() bool`

HasPresenceLastUpdateId returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *UpdatePresence200Response) GetServerTimestamp() float32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *UpdatePresence200Response) GetServerTimestampOk() (*float32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *UpdatePresence200Response) SetServerTimestamp(v float32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *UpdatePresence200Response) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetPresences

`func (o *UpdatePresence200Response) GetPresences() map[string]PresenceUpdateValue`

GetPresences returns the Presences field if non-nil, zero value otherwise.

### GetPresencesOk

`func (o *UpdatePresence200Response) GetPresencesOk() (*map[string]PresenceUpdateValue, bool)`

GetPresencesOk returns a tuple with the Presences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresences

`func (o *UpdatePresence200Response) SetPresences(v map[string]PresenceUpdateValue)`

SetPresences sets Presences field to given value.

### HasPresences

`func (o *UpdatePresence200Response) HasPresences() bool`

HasPresences returns a boolean if a field has been set.

### GetZephyrMirrorActive

`func (o *UpdatePresence200Response) GetZephyrMirrorActive() bool`

GetZephyrMirrorActive returns the ZephyrMirrorActive field if non-nil, zero value otherwise.

### GetZephyrMirrorActiveOk

`func (o *UpdatePresence200Response) GetZephyrMirrorActiveOk() (*bool, bool)`

GetZephyrMirrorActiveOk returns a tuple with the ZephyrMirrorActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZephyrMirrorActive

`func (o *UpdatePresence200Response) SetZephyrMirrorActive(v bool)`

SetZephyrMirrorActive sets ZephyrMirrorActive field to given value.

### HasZephyrMirrorActive

`func (o *UpdatePresence200Response) HasZephyrMirrorActive() bool`

HasZephyrMirrorActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


