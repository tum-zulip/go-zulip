# NavigationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | **string** | A unique identifier for the view, used to determine navigation behavior when clicked.  Clients should use this value to navigate to the corresponding URL hash.  | 
**IsPinned** | **bool** | Determines whether the view appears directly in the sidebar or is hidden in the \&quot;More Views\&quot; menu.  - &#x60;true&#x60; - Pinned and visible in the sidebar. - &#x60;false&#x60; - Hidden and accessible via the \&quot;More Views\&quot; menu.  | 
**Name** | Pointer to **NullableString** | The user-facing name for custom navigation views. Omit this field for built-in views.  | [optional] 

## Methods

### NewNavigationView

`func NewNavigationView(fragment string, isPinned bool, ) *NavigationView`

NewNavigationView instantiates a new NavigationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationViewWithDefaults

`func NewNavigationViewWithDefaults() *NavigationView`

NewNavigationViewWithDefaults instantiates a new NavigationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *NavigationView) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *NavigationView) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *NavigationView) SetFragment(v string)`

SetFragment sets Fragment field to given value.


### GetIsPinned

`func (o *NavigationView) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *NavigationView) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *NavigationView) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.


### GetName

`func (o *NavigationView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NavigationView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NavigationView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NavigationView) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NavigationView) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NavigationView) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


