# GetInvites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Invites** | Pointer to [**[]Invite**](Invite.md) | An array of objects, each representing a single unexpired [invitation](zulip.com/help/invite-new-users.  | [optional] 

## Methods

### NewGetInvites200Response

`func NewGetInvites200Response(result interface{}, msg interface{}, ) *GetInvites200Response`

NewGetInvites200Response instantiates a new GetInvites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvites200ResponseWithDefaults

`func NewGetInvites200ResponseWithDefaults() *GetInvites200Response`

NewGetInvites200ResponseWithDefaults instantiates a new GetInvites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetInvites200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetInvites200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetInvites200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetInvites200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetInvites200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetInvites200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetInvites200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetInvites200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetInvites200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetInvites200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetInvites200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetInvites200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetInvites200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetInvites200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetInvites200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetInvites200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetInvites

`func (o *GetInvites200Response) GetInvites() []Invite`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *GetInvites200Response) GetInvitesOk() (*[]Invite, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *GetInvites200Response) SetInvites(v []Invite)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *GetInvites200Response) HasInvites() bool`

HasInvites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


