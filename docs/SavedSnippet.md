# SavedSnippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of the saved snippet.  | [optional] 
**Title** | Pointer to **string** | The title of the saved snippet.  | [optional] 
**Content** | Pointer to **string** | The content of the saved snippet in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should insert this content into a message when using a saved snippet.  | [optional] 
**DateCreated** | Pointer to **int32** | The UNIX timestamp for when the saved snippet was created, in UTC seconds.  | [optional] 

## Methods

### NewSavedSnippet

`func NewSavedSnippet() *SavedSnippet`

NewSavedSnippet instantiates a new SavedSnippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSnippetWithDefaults

`func NewSavedSnippetWithDefaults() *SavedSnippet`

NewSavedSnippetWithDefaults instantiates a new SavedSnippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedSnippet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedSnippet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedSnippet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SavedSnippet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *SavedSnippet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SavedSnippet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SavedSnippet) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SavedSnippet) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *SavedSnippet) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SavedSnippet) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SavedSnippet) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SavedSnippet) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDateCreated

`func (o *SavedSnippet) GetDateCreated() int32`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SavedSnippet) GetDateCreatedOk() (*int32, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SavedSnippet) SetDateCreated(v int32)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SavedSnippet) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


