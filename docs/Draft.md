# Draft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of the draft. It will only used whenever the drafts are fetched. This field should not be specified when the draft is being created or edited.  | [optional] 
**Type** | **string** | The type of the draft. Either unaddressed (empty string), &#x60;\&quot;stream\&quot;&#x60;, or &#x60;\&quot;private\&quot;&#x60; (for one-on-one and group direct messages).  | 
**To** | **[]int32** | An array of the tentative target audience IDs. For channel messages, this should contain exactly 1 ID, the ID of the target channel. For direct messages, this should be an array of target user IDs. For unaddressed drafts, this is ignored, and clients should send an empty array.  | 
**Topic** | **string** | For channel message drafts, the tentative topic name. For direct or unaddressed messages, this will be ignored and should ideally be the empty string. Should not contain null bytes.  | 
**Content** | **string** | The body of the draft. Should not contain null bytes.  | 
**Timestamp** | Pointer to **int32** | A Unix timestamp (seconds only) representing when the draft was last edited. When creating a draft, this key need not be present and it will be filled in automatically by the server.  | [optional] 

## Methods

### NewDraft

`func NewDraft(type_ string, to []int32, topic string, content string, ) *Draft`

NewDraft instantiates a new Draft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftWithDefaults

`func NewDraftWithDefaults() *Draft`

NewDraftWithDefaults instantiates a new Draft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Draft) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Draft) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Draft) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Draft) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Draft) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Draft) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Draft) SetType(v string)`

SetType sets Type field to given value.


### GetTo

`func (o *Draft) GetTo() []int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Draft) GetToOk() (*[]int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Draft) SetTo(v []int32)`

SetTo sets To field to given value.


### GetTopic

`func (o *Draft) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Draft) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Draft) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetContent

`func (o *Draft) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Draft) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Draft) SetContent(v string)`

SetContent sets Content field to given value.


### GetTimestamp

`func (o *Draft) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Draft) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Draft) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Draft) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


