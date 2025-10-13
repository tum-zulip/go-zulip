# GetMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**Anchor** | Pointer to **int32** | The same &#x60;anchor&#x60; specified in the request (or the computed one, if &#x60;use_first_unread_anchor&#x60; is &#x60;true&#x60;).  Only present if &#x60;message_ids&#x60; is not provided.  | [optional] 
**FoundNewest** | Pointer to **bool** | Whether the server promises that the &#x60;messages&#x60; list includes the very newest messages matching the narrow (used by clients that paginate their requests to decide whether there may be more messages to fetch).  | [optional] 
**FoundOldest** | Pointer to **bool** | Whether the server promises that the &#x60;messages&#x60; list includes the very oldest messages matching the narrow (used by clients that paginate their requests to decide whether there may be more messages to fetch).  | [optional] 
**FoundAnchor** | Pointer to **bool** | Whether the anchor message is included in the response. If the message with the ID specified in the request does not exist, did not match the narrow, or was excluded via &#x60;\&quot;include_anchor\&quot;: false&#x60;, this will be false.  | [optional] 
**HistoryLimited** | Pointer to **bool** | Whether the message history was limited due to plan restrictions. This flag is set to &#x60;true&#x60; only when the oldest messages(&#x60;found_oldest&#x60;) matching the narrow is fetched.  | [optional] 
**Messages** | [**[]GetMessages200ResponseAllOfMessagesInner**](GetMessages200ResponseAllOfMessagesInner.md) | An array of &#x60;message&#x60; objects.  **Changes**: In Zulip 3.1 (feature level 26), the &#x60;sender_short_name&#x60; field was removed from message objects.  | 

## Methods

### NewGetMessages200Response

`func NewGetMessages200Response(result interface{}, msg interface{}, messages []GetMessages200ResponseAllOfMessagesInner, ) *GetMessages200Response`

NewGetMessages200Response instantiates a new GetMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessages200ResponseWithDefaults

`func NewGetMessages200ResponseWithDefaults() *GetMessages200Response`

NewGetMessages200ResponseWithDefaults instantiates a new GetMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetMessages200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMessages200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMessages200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetMessages200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetMessages200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetMessages200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetMessages200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetMessages200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetMessages200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetMessages200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetMessages200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetMessages200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetMessages200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetMessages200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetMessages200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetMessages200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetAnchor

`func (o *GetMessages200Response) GetAnchor() int32`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *GetMessages200Response) GetAnchorOk() (*int32, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *GetMessages200Response) SetAnchor(v int32)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *GetMessages200Response) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetFoundNewest

`func (o *GetMessages200Response) GetFoundNewest() bool`

GetFoundNewest returns the FoundNewest field if non-nil, zero value otherwise.

### GetFoundNewestOk

`func (o *GetMessages200Response) GetFoundNewestOk() (*bool, bool)`

GetFoundNewestOk returns a tuple with the FoundNewest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundNewest

`func (o *GetMessages200Response) SetFoundNewest(v bool)`

SetFoundNewest sets FoundNewest field to given value.

### HasFoundNewest

`func (o *GetMessages200Response) HasFoundNewest() bool`

HasFoundNewest returns a boolean if a field has been set.

### GetFoundOldest

`func (o *GetMessages200Response) GetFoundOldest() bool`

GetFoundOldest returns the FoundOldest field if non-nil, zero value otherwise.

### GetFoundOldestOk

`func (o *GetMessages200Response) GetFoundOldestOk() (*bool, bool)`

GetFoundOldestOk returns a tuple with the FoundOldest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundOldest

`func (o *GetMessages200Response) SetFoundOldest(v bool)`

SetFoundOldest sets FoundOldest field to given value.

### HasFoundOldest

`func (o *GetMessages200Response) HasFoundOldest() bool`

HasFoundOldest returns a boolean if a field has been set.

### GetFoundAnchor

`func (o *GetMessages200Response) GetFoundAnchor() bool`

GetFoundAnchor returns the FoundAnchor field if non-nil, zero value otherwise.

### GetFoundAnchorOk

`func (o *GetMessages200Response) GetFoundAnchorOk() (*bool, bool)`

GetFoundAnchorOk returns a tuple with the FoundAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAnchor

`func (o *GetMessages200Response) SetFoundAnchor(v bool)`

SetFoundAnchor sets FoundAnchor field to given value.

### HasFoundAnchor

`func (o *GetMessages200Response) HasFoundAnchor() bool`

HasFoundAnchor returns a boolean if a field has been set.

### GetHistoryLimited

`func (o *GetMessages200Response) GetHistoryLimited() bool`

GetHistoryLimited returns the HistoryLimited field if non-nil, zero value otherwise.

### GetHistoryLimitedOk

`func (o *GetMessages200Response) GetHistoryLimitedOk() (*bool, bool)`

GetHistoryLimitedOk returns a tuple with the HistoryLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryLimited

`func (o *GetMessages200Response) SetHistoryLimited(v bool)`

SetHistoryLimited sets HistoryLimited field to given value.

### HasHistoryLimited

`func (o *GetMessages200Response) HasHistoryLimited() bool`

HasHistoryLimited returns a boolean if a field has been set.

### GetMessages

`func (o *GetMessages200Response) GetMessages() []GetMessages200ResponseAllOfMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetMessages200Response) GetMessagesOk() (*[]GetMessages200ResponseAllOfMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetMessages200Response) SetMessages(v []GetMessages200ResponseAllOfMessagesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


