# ChannelFolderAddEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**ChannelFolder** | Pointer to [**ChannelFolder**](ChannelFolder.md) |  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf83

`func NewGetEvents200ResponseAllOfEventsInnerOneOf83() *ChannelFolderAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf83 instantiates a new ChannelFolderAddEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf83WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf83WithDefaults() *ChannelFolderAddEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf83WithDefaults instantiates a new ChannelFolderAddEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelFolderAddEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFolderAddEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFolderAddEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelFolderAddEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ChannelFolderAddEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChannelFolderAddEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChannelFolderAddEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChannelFolderAddEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOp

`func (o *ChannelFolderAddEvent) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ChannelFolderAddEvent) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ChannelFolderAddEvent) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ChannelFolderAddEvent) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetChannelFolder

`func (o *ChannelFolderAddEvent) GetChannelFolder() ChannelFolder`

GetChannelFolder returns the ChannelFolder field if non-nil, zero value otherwise.

### GetChannelFolderOk

`func (o *ChannelFolderAddEvent) GetChannelFolderOk() (*ChannelFolder, bool)`

GetChannelFolderOk returns a tuple with the ChannelFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFolder

`func (o *ChannelFolderAddEvent) SetChannelFolder(v ChannelFolder)`

SetChannelFolder sets ChannelFolder field to given value.

### HasChannelFolder

`func (o *ChannelFolderAddEvent) HasChannelFolder() bool`

HasChannelFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


