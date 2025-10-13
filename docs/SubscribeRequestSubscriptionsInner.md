# SubscribeRequestSubscriptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the channel.  Clients should use the &#x60;max_stream_name_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum channel name length.  | 
**Description** | Pointer to **string** | The [description](/help/change-the-channel-description) to use for a new channel being created, in text/markdown format.  See the help center article on [message formatting](/help/format-your-message-using-markdown) for details on Zulip-flavored Markdown.  Clients should use the &#x60;max_stream_description_length&#x60; returned by the [&#x60;POST /register&#x60;](/api/register-queue) endpoint to determine the maximum channel description length.  | [optional] 

## Methods

### NewSubscribeRequestSubscriptionsInner

`func NewSubscribeRequestSubscriptionsInner(name string, ) *SubscribeRequestSubscriptionsInner`

NewSubscribeRequestSubscriptionsInner instantiates a new SubscribeRequestSubscriptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeRequestSubscriptionsInnerWithDefaults

`func NewSubscribeRequestSubscriptionsInnerWithDefaults() *SubscribeRequestSubscriptionsInner`

NewSubscribeRequestSubscriptionsInnerWithDefaults instantiates a new SubscribeRequestSubscriptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscribeRequestSubscriptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscribeRequestSubscriptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscribeRequestSubscriptionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscribeRequestSubscriptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscribeRequestSubscriptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscribeRequestSubscriptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscribeRequestSubscriptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


