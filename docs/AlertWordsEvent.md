# AlertWordsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event. Events appear in increasing order but may not be consecutive.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AlertWords** | Pointer to **[]string** | An array of strings, where each string is an alert word (or phrase) configured by the user.  | [optional] 

## Methods

### NewGetEvents200ResponseAllOfEventsInnerOneOf

`func NewGetEvents200ResponseAllOfEventsInnerOneOf() *AlertWordsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOf instantiates a new AlertWordsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseAllOfEventsInnerOneOfWithDefaults

`func NewGetEvents200ResponseAllOfEventsInnerOneOfWithDefaults() *AlertWordsEvent`

NewGetEvents200ResponseAllOfEventsInnerOneOfWithDefaults instantiates a new AlertWordsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertWordsEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertWordsEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertWordsEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertWordsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AlertWordsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertWordsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertWordsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AlertWordsEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlertWords

`func (o *AlertWordsEvent) GetAlertWords() []string`

GetAlertWords returns the AlertWords field if non-nil, zero value otherwise.

### GetAlertWordsOk

`func (o *AlertWordsEvent) GetAlertWordsOk() (*[]string, bool)`

GetAlertWordsOk returns a tuple with the AlertWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertWords

`func (o *AlertWordsEvent) SetAlertWords(v []string)`

SetAlertWords sets AlertWords field to given value.

### HasAlertWords

`func (o *AlertWordsEvent) HasAlertWords() bool`

HasAlertWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


