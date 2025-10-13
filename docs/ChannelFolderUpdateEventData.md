# ChannelFolderUpdateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ChannelFolderId** | Pointer to **float32** | ID of the updated channel folder.  | [optional] 
**Data** | Pointer to [**GetEvents200ResponseAllOfEventsInnerOneOf84Data**](GetEvents200ResponseAllOfEventsInnerOneOf84Data.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf84

`func NewGetEvents200ResponseAllOfEventsInnerOneOf84() *ChannelFolderUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf84 instantiates a new ChannelFolderUpdateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf84WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf84WithDefaults() *ChannelFolderUpdateEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf84WithDefaults instantiates a new ChannelFolderUpdateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelFolderUpdateEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFolderUpdateEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFolderUpdateEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelFolderUpdateEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ChannelFolderUpdateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelFolderUpdateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelFolderUpdateEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelFolderUpdateEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ChannelFolderUpdateEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ChannelFolderUpdateEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ChannelFolderUpdateEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ChannelFolderUpdateEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetChannelFolderId

`func (o *ChannelFolderUpdateEvent) GetChannelFolderId() float32`

GetChannelFolderId returns the ChannelFolderId field if non-nil, zero value otherwise.

### GetChannelFolderIdOk

`func (o *ChannelFolderUpdateEvent) GetChannelFolderIdOk() (*float32, bool)`

GetChannelFolderIdOk returns a tuple with the ChannelFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolderId

`func (o *ChannelFolderUpdateEvent) SetChannelFolderId(v float32)`

SetChannelFolderId sets ChannelFolderId field to given value.

### HasChannelFolderId

`func (o *ChannelFolderUpdateEvent) HasChannelFolderId() bool`

HasChannelFolderId returns a boolean if a field has been set.

### GetData

`func (o *ChannelFolderUpdateEvent) GetData() GetEvents200ResponseAllOfEventsInnerOneOf84Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChannelFolderUpdateEvent) GetDataOk() (*GetEvents200ResponseAllOfEventsInnerOneOf84Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChannelFolderUpdateEvent) SetData(v GetEvents200ResponseAllOfEventsInnerOneOf84Data)`

SetData sets Data field to given value.

### HasData

`func (o *ChannelFolderUpdateEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


