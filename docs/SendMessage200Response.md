# SendMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Id** | **int32** | The unique ID assigned to the sent message.  | 
**AutomaticNewVisibilityPolicy** | Pointer to **int32** | If the message&#39;s sender had configured their [visibility policy settings](/help/mute-a-topic) to potentially automatically follow or unmute topics when sending messages, and one of these policies did in fact change the user&#39;s visibility policy for the topic where this message was sent, the new value for that user&#39;s visibility policy for the recipient topic.  Only present if the sender&#39;s visibility was in fact changed.  The value can be either [unmuted or followed](/api/update-user-topic#parameter-visibility_policy).  Clients will also be notified about the change in policy via a &#x60;user_topic&#x60; event as usual. This field is intended to be used by clients to explicitly inform the user when a topic&#39;s visibility policy was changed automatically due to sending a message.  For example, the Zulip web application uses this field to decide whether to display a warning or notice suggesting to unmute the topic after sending a message to a muted channel. Such a notice would be confusing in the event that the act of sending the message had already resulted in the user automatically unmuting or following the topic in question.  **Changes**: New in Zulip 8.0 (feature level 218).  | [optional] 

## Methods

### NewSendMessage200Response

`func NewSendMessage200Response(result interface{}, msg interface{}, id int32, ) *SendMessage200Response`

NewSendMessage200Response instantiates a new SendMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessage200ResponseWithDefaults

`func NewSendMessage200ResponseWithDefaults() *SendMessage200Response`

NewSendMessage200ResponseWithDefaults instantiates a new SendMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SendMessage200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SendMessage200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SendMessage200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *SendMessage200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *SendMessage200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *SendMessage200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SendMessage200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SendMessage200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *SendMessage200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *SendMessage200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *SendMessage200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *SendMessage200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *SendMessage200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *SendMessage200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *SendMessage200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *SendMessage200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetId

`func (o *SendMessage200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendMessage200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendMessage200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetAutomaticNewVisibilityPolicy

`func (o *SendMessage200Response) GetAutomaticNewVisibilityPolicy() int32`

GetAutomaticNewVisibilityPolicy returns the AutomaticNewVisibilityPolicy field if non-nil, zero value otherwise.

### GetAutomaticNewVisibilityPolicyOk

`func (o *SendMessage200Response) GetAutomaticNewVisibilityPolicyOk() (*int32, bool)`

GetAutomaticNewVisibilityPolicyOk returns a tuple with the AutomaticNewVisibilityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticNewVisibilityPolicy

`func (o *SendMessage200Response) SetAutomaticNewVisibilityPolicy(v int32)`

SetAutomaticNewVisibilityPolicy sets AutomaticNewVisibilityPolicy field to given value.

### HasAutomaticNewVisibilityPolicy

`func (o *SendMessage200Response) HasAutomaticNewVisibilityPolicy() bool`

HasAutomaticNewVisibilityPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


