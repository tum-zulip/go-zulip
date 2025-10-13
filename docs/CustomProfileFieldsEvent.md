# CustomProfileFieldsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]CustomProfileField**](CustomProfileField.md) | An array of dictionaries where each dictionary contains details of a single new custom profile field for the Zulip organization.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf27

`func NewGetEvents200ResponseAllOfEventsInnerOneOf27() *CustomProfileFieldsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf27 instantiates a new CustomProfileFieldsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOf27WithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOf27WithDefaults() *CustomProfileFieldsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf27WithDefaults instantiates a new CustomProfileFieldsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomProfileFieldsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomProfileFieldsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomProfileFieldsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomProfileFieldsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomProfileFieldsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomProfileFieldsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomProfileFieldsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomProfileFieldsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFields

`func (o *CustomProfileFieldsEvent) GetFields() []CustomProfileField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomProfileFieldsEvent) GetFieldsOk() (*[]CustomProfileField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomProfileFieldsEvent) SetFields(v []CustomProfileField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomProfileFieldsEvent) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


