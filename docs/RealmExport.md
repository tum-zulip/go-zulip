# RealmExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the data export.  | [optional] 
**ActingUserId** | Pointer to **int32** | The ID of the user who created the data export.  | [optional] 
**ExportTime** | Pointer to **float32** | The UNIX timestamp of when the data export was started.  | [optional] 
**DeletedTimestamp** | Pointer to **NullableFloat32** | The UNIX timestamp of when the data export was deleted.  Will be &#x60;null&#x60; if the data export has not been deleted.  | [optional] 
**FailedTimestamp** | Pointer to **NullableFloat32** | The UNIX timestamp of when the data export failed.  Will be &#x60;null&#x60; if the data export succeeded, or if it&#39;s still being generated.  | [optional] 
**ExportUrl** | Pointer to **NullableString** | The URL to download the generated data export.  Will be &#x60;null&#x60; if the data export failed, or if it&#39;s still being generated.  | [optional] 
**Pending** | Pointer to **bool** | Whether the data export is pending, which indicates it is still being generated, or if it succeeded, failed or was deleted before being generated.  Depending on the size of the organization, it can take anywhere from seconds to an hour to generate the data export.  | [optional] 
**ExportType** | Pointer to **int32** | Whether the data export is a public or a standard data export.  - 1 &#x3D; Public data export. - 2 &#x3D; Standard data export.  **Changes**: New in Zulip 10.0 (feature level 304). Previously, the export type was not included in these objects because only public data exports could be created or listed via the API or UI.  | [optional] 

## Methods

### NewRealmExport

`func NewRealmExport() *RealmExport`

NewRealmExport instantiates a new RealmExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmExportWithDefaults

`func NewRealmExportWithDefaults() *RealmExport`

NewRealmExportWithDefaults instantiates a new RealmExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmExport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmExport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmExport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RealmExport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActingUserId

`func (o *RealmExport) GetActingUserId() int32`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *RealmExport) GetActingUserIdOk() (*int32, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *RealmExport) SetActingUserId(v int32)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *RealmExport) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.

### GetExportTime

`func (o *RealmExport) GetExportTime() float32`

GetExportTime returns the ExportTime field if non-nil, zero value otherwise.

### GetExportTimeOk

`func (o *RealmExport) GetExportTimeOk() (*float32, bool)`

GetExportTimeOk returns a tuple with the ExportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTime

`func (o *RealmExport) SetExportTime(v float32)`

SetExportTime sets ExportTime field to given value.

### HasExportTime

`func (o *RealmExport) HasExportTime() bool`

HasExportTime returns a boolean if a field has been set.

### GetDeletedTimestamp

`func (o *RealmExport) GetDeletedTimestamp() float32`

GetDeletedTimestamp returns the DeletedTimestamp field if non-nil, zero value otherwise.

### GetDeletedTimestampOk

`func (o *RealmExport) GetDeletedTimestampOk() (*float32, bool)`

GetDeletedTimestampOk returns a tuple with the DeletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTimestamp

`func (o *RealmExport) SetDeletedTimestamp(v float32)`

SetDeletedTimestamp sets DeletedTimestamp field to given value.

### HasDeletedTimestamp

`func (o *RealmExport) HasDeletedTimestamp() bool`

HasDeletedTimestamp returns a boolean if a field has been set.

### SetDeletedTimestampNil

`func (o *RealmExport) SetDeletedTimestampNil(b bool)`

 SetDeletedTimestampNil sets the value for DeletedTimestamp to be an explicit nil

### UnsetDeletedTimestamp
`func (o *RealmExport) UnsetDeletedTimestamp()`

UnsetDeletedTimestamp ensures that no value is present for DeletedTimestamp, not even an explicit nil
### GetFailedTimestamp

`func (o *RealmExport) GetFailedTimestamp() float32`

GetFailedTimestamp returns the FailedTimestamp field if non-nil, zero value otherwise.

### GetFailedTimestampOk

`func (o *RealmExport) GetFailedTimestampOk() (*float32, bool)`

GetFailedTimestampOk returns a tuple with the FailedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTimestamp

`func (o *RealmExport) SetFailedTimestamp(v float32)`

SetFailedTimestamp sets FailedTimestamp field to given value.

### HasFailedTimestamp

`func (o *RealmExport) HasFailedTimestamp() bool`

HasFailedTimestamp returns a boolean if a field has been set.

### SetFailedTimestampNil

`func (o *RealmExport) SetFailedTimestampNil(b bool)`

 SetFailedTimestampNil sets the value for FailedTimestamp to be an explicit nil

### UnsetFailedTimestamp
`func (o *RealmExport) UnsetFailedTimestamp()`

UnsetFailedTimestamp ensures that no value is present for FailedTimestamp, not even an explicit nil
### GetExportUrl

`func (o *RealmExport) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *RealmExport) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *RealmExport) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.

### HasExportUrl

`func (o *RealmExport) HasExportUrl() bool`

HasExportUrl returns a boolean if a field has been set.

### SetExportUrlNil

`func (o *RealmExport) SetExportUrlNil(b bool)`

 SetExportUrlNil sets the value for ExportUrl to be an explicit nil

### UnsetExportUrl
`func (o *RealmExport) UnsetExportUrl()`

UnsetExportUrl ensures that no value is present for ExportUrl, not even an explicit nil
### GetPending

`func (o *RealmExport) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *RealmExport) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *RealmExport) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *RealmExport) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetExportType

`func (o *RealmExport) GetExportType() int32`

GetExportType returns the ExportType field if non-nil, zero value otherwise.

### GetExportTypeOk

`func (o *RealmExport) GetExportTypeOk() (*int32, bool)`

GetExportTypeOk returns a tuple with the ExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportType

`func (o *RealmExport) SetExportType(v int32)`

SetExportType sets ExportType field to given value.

### HasExportType

`func (o *RealmExport) HasExportType() bool`

HasExportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


