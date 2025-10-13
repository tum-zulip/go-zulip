# UpdateMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**DetachedUploads** | Pointer to [**[]Attachment**](Attachment.md) | Details on all files uploaded by the acting user whose only references were removed when editing this message. Clients should ask the acting user if they wish to delete the uploaded files returned in this response, which might otherwise remain visible only in message edit history.  Note that [access to message edit history][edit-history-access] is configurable; this detail may be important in presenting the question clearly to users.  New in Zulip 10.0 (feature level 285).  [edit-history-access]: /help/restrict-message-edit-history-access  | [optional] 

## Methods

### NewUpdateMessage200Response

`func NewUpdateMessage200Response(result interface{}, msg interface{}, ) *UpdateMessage200Response`

NewUpdateMessage200Response instantiates a new UpdateMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessage200ResponseWithDefaults

`func NewUpdateMessage200ResponseWithDefaults() *UpdateMessage200Response`

NewUpdateMessage200ResponseWithDefaults instantiates a new UpdateMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateMessage200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateMessage200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateMessage200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UpdateMessage200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdateMessage200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UpdateMessage200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateMessage200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateMessage200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UpdateMessage200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateMessage200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UpdateMessage200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UpdateMessage200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UpdateMessage200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UpdateMessage200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *UpdateMessage200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *UpdateMessage200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetDetachedUploads

`func (o *UpdateMessage200Response) GetDetachedUploads() []Attachment`

GetDetachedUploads returns the DetachedUploads field if non-nil, zero value otherwise.

### GetDetachedUploadsOk

`func (o *UpdateMessage200Response) GetDetachedUploadsOk() (*[]Attachment, bool)`

GetDetachedUploadsOk returns a tuple with the DetachedUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachedUploads

`func (o *UpdateMessage200Response) SetDetachedUploads(v []Attachment)`

SetDetachedUploads sets DetachedUploads field to given value.

### HasDetachedUploads

`func (o *UpdateMessage200Response) HasDetachedUploads() bool`

HasDetachedUploads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


