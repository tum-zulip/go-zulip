# UpdateMessageFlagsForNarrow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **[]string** | An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling](zulip.com/api/rest-error-handling#ignored-parameters documentation for details on this and its change history.  | [optional] 
**ProcessedCount** | **int32** | The number of messages that were within the update range (at most &#x60;num_before + 1 + num_after&#x60;).  | 
**UpdatedCount** | **int32** | The number of messages where the flag&#39;s value was changed (at most &#x60;processed_count&#x60;).  | 
**FirstProcessedId** | **NullableInt32** | The ID of the oldest message within the update range, or &#x60;null&#x60; if the range was empty.  | 
**LastProcessedId** | **NullableInt32** | The ID of the newest message within the update range, or &#x60;null&#x60; if the range was empty.  | 
**FoundOldest** | **bool** | Whether the update range reached backward far enough to include very oldest message matching the narrow (used by clients doing a bulk update to decide whether to issue another request anchored at &#x60;first_processed_id&#x60;).  | 
**FoundNewest** | **bool** | Whether the update range reached forward far enough to include very oldest message matching the narrow (used by clients doing a bulk update to decide whether to issue another request anchored at &#x60;last_processed_id&#x60;).  | 
**IgnoredBecauseNotSubscribedChannels** | Pointer to **[]int32** | Only present if the flag is &#x60;read&#x60; and the operation is &#x60;remove&#x60;.  Zulip has an invariant that all unread messages must be in channels the user is subscribed to. This field will contain a list of the channels whose messages were skipped to mark as unread because the user is not subscribed to them.  **Changes**: New in Zulip 10.0 (feature level 355).  | [optional] 

## Methods

### NewUpdateMessageFlagsForNarrow200Response

`func NewUpdateMessageFlagsForNarrow200Response(result interface{}, msg interface{}, processedCount int32, updatedCount int32, firstProcessedId NullableInt32, lastProcessedId NullableInt32, foundOldest bool, foundNewest bool, ) *UpdateMessageFlagsForNarrow200Response`

NewUpdateMessageFlagsForNarrow200Response instantiates a new UpdateMessageFlagsForNarrow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageFlagsForNarrow200ResponseWithDefaults

`func NewUpdateMessageFlagsForNarrow200ResponseWithDefaults() *UpdateMessageFlagsForNarrow200Response`

NewUpdateMessageFlagsForNarrow200ResponseWithDefaults instantiates a new UpdateMessageFlagsForNarrow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateMessageFlagsForNarrow200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateMessageFlagsForNarrow200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *UpdateMessageFlagsForNarrow200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdateMessageFlagsForNarrow200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *UpdateMessageFlagsForNarrow200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateMessageFlagsForNarrow200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *UpdateMessageFlagsForNarrow200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateMessageFlagsForNarrow200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *UpdateMessageFlagsForNarrow200Response) GetIgnoredParametersUnsupported() []string`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetIgnoredParametersUnsupportedOk() (*[]string, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *UpdateMessageFlagsForNarrow200Response) SetIgnoredParametersUnsupported(v []string)`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *UpdateMessageFlagsForNarrow200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### GetProcessedCount

`func (o *UpdateMessageFlagsForNarrow200Response) GetProcessedCount() int32`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetProcessedCountOk() (*int32, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *UpdateMessageFlagsForNarrow200Response) SetProcessedCount(v int32)`

SetProcessedCount sets ProcessedCount field to given value.


### GetUpdatedCount

`func (o *UpdateMessageFlagsForNarrow200Response) GetUpdatedCount() int32`

GetUpdatedCount returns the UpdatedCount field if non-nil, zero value otherwise.

### GetUpdatedCountOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetUpdatedCountOk() (*int32, bool)`

GetUpdatedCountOk returns a tuple with the UpdatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCount

`func (o *UpdateMessageFlagsForNarrow200Response) SetUpdatedCount(v int32)`

SetUpdatedCount sets UpdatedCount field to given value.


### GetFirstProcessedId

`func (o *UpdateMessageFlagsForNarrow200Response) GetFirstProcessedId() int32`

GetFirstProcessedId returns the FirstProcessedId field if non-nil, zero value otherwise.

### GetFirstProcessedIdOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetFirstProcessedIdOk() (*int32, bool)`

GetFirstProcessedIdOk returns a tuple with the FirstProcessedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstProcessedId

`func (o *UpdateMessageFlagsForNarrow200Response) SetFirstProcessedId(v int32)`

SetFirstProcessedId sets FirstProcessedId field to given value.


### SetFirstProcessedIdNil

`func (o *UpdateMessageFlagsForNarrow200Response) SetFirstProcessedIdNil(b bool)`

 SetFirstProcessedIdNil sets the value for FirstProcessedId to be an explicit nil

### UnsetFirstProcessedId
`func (o *UpdateMessageFlagsForNarrow200Response) UnsetFirstProcessedId()`

UnsetFirstProcessedId ensures that no value is present for FirstProcessedId, not even an explicit nil
### GetLastProcessedId

`func (o *UpdateMessageFlagsForNarrow200Response) GetLastProcessedId() int32`

GetLastProcessedId returns the LastProcessedId field if non-nil, zero value otherwise.

### GetLastProcessedIdOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetLastProcessedIdOk() (*int32, bool)`

GetLastProcessedIdOk returns a tuple with the LastProcessedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedId

`func (o *UpdateMessageFlagsForNarrow200Response) SetLastProcessedId(v int32)`

SetLastProcessedId sets LastProcessedId field to given value.


### SetLastProcessedIdNil

`func (o *UpdateMessageFlagsForNarrow200Response) SetLastProcessedIdNil(b bool)`

 SetLastProcessedIdNil sets the value for LastProcessedId to be an explicit nil

### UnsetLastProcessedId
`func (o *UpdateMessageFlagsForNarrow200Response) UnsetLastProcessedId()`

UnsetLastProcessedId ensures that no value is present for LastProcessedId, not even an explicit nil
### GetFoundOldest

`func (o *UpdateMessageFlagsForNarrow200Response) GetFoundOldest() bool`

GetFoundOldest returns the FoundOldest field if non-nil, zero value otherwise.

### GetFoundOldestOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetFoundOldestOk() (*bool, bool)`

GetFoundOldestOk returns a tuple with the FoundOldest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundOldest

`func (o *UpdateMessageFlagsForNarrow200Response) SetFoundOldest(v bool)`

SetFoundOldest sets FoundOldest field to given value.


### GetFoundNewest

`func (o *UpdateMessageFlagsForNarrow200Response) GetFoundNewest() bool`

GetFoundNewest returns the FoundNewest field if non-nil, zero value otherwise.

### GetFoundNewestOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetFoundNewestOk() (*bool, bool)`

GetFoundNewestOk returns a tuple with the FoundNewest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundNewest

`func (o *UpdateMessageFlagsForNarrow200Response) SetFoundNewest(v bool)`

SetFoundNewest sets FoundNewest field to given value.


### GetIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlagsForNarrow200Response) GetIgnoredBecauseNotSubscribedChannels() []int32`

GetIgnoredBecauseNotSubscribedChannels returns the IgnoredBecauseNotSubscribedChannels field if non-nil, zero value otherwise.

### GetIgnoredBecauseNotSubscribedChannelsOk

`func (o *UpdateMessageFlagsForNarrow200Response) GetIgnoredBecauseNotSubscribedChannelsOk() (*[]int32, bool)`

GetIgnoredBecauseNotSubscribedChannelsOk returns a tuple with the IgnoredBecauseNotSubscribedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlagsForNarrow200Response) SetIgnoredBecauseNotSubscribedChannels(v []int32)`

SetIgnoredBecauseNotSubscribedChannels sets IgnoredBecauseNotSubscribedChannels field to given value.

### HasIgnoredBecauseNotSubscribedChannels

`func (o *UpdateMessageFlagsForNarrow200Response) HasIgnoredBecauseNotSubscribedChannels() bool`

HasIgnoredBecauseNotSubscribedChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


