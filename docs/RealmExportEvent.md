# RealmExportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Exports** | Pointer to [**[]RealmExport**](RealmExport.md) | An array of dictionaries where each dictionary contains details about a data export of the organization.  **Changes**: Prior to Zulip 10.0 (feature level 304), &#x60;export_type&#x60; parameter was not present as only public data export was supported via API.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf57

`func NewGetEvents200ResponseAllOfEventsInnerOneOf57() *RealmExportEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf57 instantiates a new RealmExportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf57WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf57WithDefaults() *RealmExportEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf57WithDefaults instantiates a new RealmExportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmExportEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmExportEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmExportEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmExportEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RealmExportEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmExportEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmExportEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmExportEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExports

`func (o *RealmExportEvent) GetExports() []RealmExport`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *RealmExportEvent) GetExportsOk() (*[]RealmExport, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *RealmExportEvent) SetExports(v []RealmExport)`

SetExports sets Exports field to given value.

### HasExports

`func (o *RealmExportEvent) HasExports() bool`

HasExports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


