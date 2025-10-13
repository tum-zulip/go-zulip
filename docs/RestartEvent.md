# RestartEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ZulipVersion** | Pointer to **string** | The Zulip version number, in the format where this appears in the [server_settings](/api/get-server-settings) and [register](/api/register-queue) responses.  **Changes**: New in Zulip 4.0 (feature level 59).  | [optional] 
**ZulipMergeBase** | Pointer to **string** | The Zulip merge base number, in the format where this appears in the [server_settings](/api/get-server-settings) and [register](/api/register-queue) responses.  **Changes**: New in Zulip 5.0 (feature level 88).  | [optional] 
**ZulipFeatureLevel** | Pointer to **int32** | The [Zulip feature level](/api/changelog) of the server after the restart.  Clients should use this to update their tracking of the server&#39;s capabilities, and may choose to refetch their state and create a new event queue when the API feature level has changed in a way that the client finds significant. Clients choosing to do so must implement a random delay strategy to spread such restarts over 5 or more minutes to avoid creating a synchronized thundering herd effect.  **Changes**: New in Zulip 4.0 (feature level 59).  | [optional] 
**ServerGeneration** | Pointer to **int32** | The timestamp at which the server started.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf65

`func NewGetEvents200ResponseAllOfEventsInnerOneOf65() *RestartEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf65 instantiates a new RestartEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf65WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf65WithDefaults() *RestartEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf65WithDefaults instantiates a new RestartEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestartEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestartEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestartEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RestartEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RestartEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestartEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestartEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestartEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZulipVersion

`func (o *RestartEvent) GetZulipVersion() string`

GetZulipVersion returns the ZulipVersion field if non-nil, zero value otherwise.

### GetZulipVersionOk

`func (o *RestartEvent) GetZulipVersionOk() (*string, bool)`

GetZulipVersionOk returns a tuple with the ZulipVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipVersion

`func (o *RestartEvent) SetZulipVersion(v string)`

SetZulipVersion sets ZulipVersion field to given value.

### HasZulipVersion

`func (o *RestartEvent) HasZulipVersion() bool`

HasZulipVersion returns a boolean if a field has been set.

### GetZulipMergeBase

`func (o *RestartEvent) GetZulipMergeBase() string`

GetZulipMergeBase returns the ZulipMergeBase field if non-nil, zero value otherwise.

### GetZulipMergeBaseOk

`func (o *RestartEvent) GetZulipMergeBaseOk() (*string, bool)`

GetZulipMergeBaseOk returns a tuple with the ZulipMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipMergeBase

`func (o *RestartEvent) SetZulipMergeBase(v string)`

SetZulipMergeBase sets ZulipMergeBase field to given value.

### HasZulipMergeBase

`func (o *RestartEvent) HasZulipMergeBase() bool`

HasZulipMergeBase returns a boolean if a field has been set.

### GetZulipFeatureLevel

`func (o *RestartEvent) GetZulipFeatureLevel() int32`

GetZulipFeatureLevel returns the ZulipFeatureLevel field if non-nil, zero value otherwise.

### GetZulipFeatureLevelOk

`func (o *RestartEvent) GetZulipFeatureLevelOk() (*int32, bool)`

GetZulipFeatureLevelOk returns a tuple with the ZulipFeatureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipFeatureLevel

`func (o *RestartEvent) SetZulipFeatureLevel(v int32)`

SetZulipFeatureLevel sets ZulipFeatureLevel field to given value.

### HasZulipFeatureLevel

`func (o *RestartEvent) HasZulipFeatureLevel() bool`

HasZulipFeatureLevel returns a boolean if a field has been set.

### GetServerGeneration

`func (o *RestartEvent) GetServerGeneration() int32`

GetServerGeneration returns the ServerGeneration field if non-nil, zero value otherwise.

### GetServerGenerationOk

`func (o *RestartEvent) GetServerGenerationOk() (*int32, bool)`

GetServerGenerationOk returns a tuple with the ServerGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGeneration

`func (o *RestartEvent) SetServerGeneration(v int32)`

SetServerGeneration sets ServerGeneration field to given value.

### HasServerGeneration

`func (o *RestartEvent) HasServerGeneration() bool`

HasServerGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


