# UpdateMessageFlags200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Messages** | Pointer to **[]int32** | An array with the IDs of the modified messages.  | [optional] 
**IgnoredBecauseNotSubscribedChannels** | Pointer to **[]int32** | Only present if the flag is &#x60;read&#x60; and the operation is &#x60;remove&#x60;.  Zulip has an invariant that all unread messages must be in channels the user is subscribed to. This field will contain a list of the channels whose messages were skipped to mark as unread because the user is not subscribed to them.  **Changes**: New in Zulip 10.0 (feature level 355).  | [optional] 

## Methods

### NewUpdateMessageFlags200Response

`func NewUpdateMessageFlags200Response(result interface{}, msg interface{}, ) *UpdateMessageFlags200Response`

NewUpdateMessageFlags200Response instantiates a new UpdateMessageFlags200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageFlags200ResponseWithDefaults

`func NewUpdateMessageFlags200ResponseWithDefaults() *UpdateMessageFlags200Response`

NewUpdateMessageFlags200ResponseWithDefaults instantiates a new UpdateMessageFlags200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateMessageFlags200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateMessageFlags200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateMessageFlags200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UpdateMessageFlags200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdateMessageFlags200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UpdateMessageFlags200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateMessageFlags200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateMessageFlags200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UpdateMessageFlags200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateMessageFlags200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UpdateMessageFlags200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UpdateMessageFlags200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UpdateMessageFlags200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UpdateMessageFlags200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *UpdateMessageFlags200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *UpdateMessageFlags200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetMessages

`func (o *UpdateMessageFlags200Response) GetMessages() []int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *UpdateMessageFlags200Response) GetMessagesOk() (*[]int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *UpdateMessageFlags200Response) SetMessages(v []int32)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *UpdateMessageFlags200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlags200Response) GetIgnoredBecauseNotSubscribedChannels() []int32`

GetIgnoredBecauseNotSubscribedChannels returns the IgnoredBecauseNotSubscribedChannels field if non-nil, zero value otherwise.

### GetIgnoredBecauseNotSubscribedChannelsOk

`func (o *UpdateMessageFlags200Response) GetIgnoredBecauseNotSubscribedChannelsOk() (*[]int32, bool)`

GetIgnoredBecauseNotSubscribedChannelsOk returns a tuple with the IgnoredBecauseNotSubscribedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlags200Response) SetIgnoredBecauseNotSubscribedChannels(v []int32)`

SetIgnoredBecauseNotSubscribedChannels sets IgnoredBecauseNotSubscribedChannels field to given value.

### HasIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlags200Response) HasIgnoredBecauseNotSubscribedChannels() bool`

HasIgnoredBecauseNotSubscribedChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


