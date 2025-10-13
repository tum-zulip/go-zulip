# \NavigationViewsAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNavigationView**](NavigationViewsAPI.md#AddNavigationView) | **Post** /navigation_views | Add a navigation view
[**EditNavigationView**](NavigationViewsAPI.md#EditNavigationView) | **Patch** /navigation_views/{fragment} | Update the navigation view
[**GetNavigationViews**](NavigationViewsAPI.md#GetNavigationViews) | **Get** /navigation_views | Get all navigation views
[**RemoveNavigationView**](NavigationViewsAPI.md#RemoveNavigationView) | **Delete** /navigation_views/{fragment} | Remove a navigation view



## AddNavigationView

> AddNavigationView200Response AddNavigationView(ctx).Fragment(fragment).IsPinned(isPinned).Name(name).Execute()

Add a navigation view



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	fragment := "fragment_example" // string | A unique identifier for the view, used to determine navigation behavior when clicked.  Clients should use this value to navigate to the corresponding URL hash. 
	isPinned := true // bool | Determines whether the view appears directly in the sidebar or is hidden in the \\\"More Views\\\" menu.  - `true` - Pinned and visible in the sidebar. - `false` - Hidden and accessible via the \\\"More Views\\\" menu. 
	name := "name_example" // string | The user-facing name for custom navigation views. Omit this field for built-in views.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddNavigationView(context.Background()).Fragment(fragment).IsPinned(isPinned).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NavigationViewsAPI.AddNavigationView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNavigationView`: AddNavigationView200Response
	fmt.Fprintf(os.Stdout, "Response from `NavigationViewsAPI.AddNavigationView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNavigationViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fragment** | **string** | A unique identifier for the view, used to determine navigation behavior when clicked.  Clients should use this value to navigate to the corresponding URL hash.  | 
 **isPinned** | **bool** | Determines whether the view appears directly in the sidebar or is hidden in the \\\&quot;More Views\\\&quot; menu.  - &#x60;true&#x60; - Pinned and visible in the sidebar. - &#x60;false&#x60; - Hidden and accessible via the \\\&quot;More Views\\\&quot; menu.  | 
 **name** | **string** | The user-facing name for custom navigation views. Omit this field for built-in views.  | 

### Return type

[**AddNavigationView200Response**](AddNavigationView200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditNavigationView

> AddNavigationView200Response EditNavigationView(ctx, fragment).IsPinned(isPinned).Name(name).Execute()

Update the navigation view



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	fragment := "narrow/is/alerted" // string | The unique URL hash of the navigation view to be updated.  This also serves as the identifier for the navigation view. 
	isPinned := true // bool | Determines whether the view is pinned (true) or hidden in the menu (false).  (optional)
	name := "name_example" // string | The user-facing name for custom navigation views. Omit this field for built-in views.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EditNavigationView(context.Background(), fragment).IsPinned(isPinned).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NavigationViewsAPI.EditNavigationView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditNavigationView`: AddNavigationView200Response
	fmt.Fprintf(os.Stdout, "Response from `NavigationViewsAPI.EditNavigationView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fragment** | **string** | The unique URL hash of the navigation view to be updated.  This also serves as the identifier for the navigation view.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNavigationViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isPinned** | **bool** | Determines whether the view is pinned (true) or hidden in the menu (false).  | 
 **name** | **string** | The user-facing name for custom navigation views. Omit this field for built-in views.  | 

### Return type

[**AddNavigationView200Response**](AddNavigationView200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNavigationViews

> GetNavigationViews200Response GetNavigationViews(ctx).Execute()

Get all navigation views



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetNavigationViews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NavigationViewsAPI.GetNavigationViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNavigationViews`: GetNavigationViews200Response
	fmt.Fprintf(os.Stdout, "Response from `NavigationViewsAPI.GetNavigationViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNavigationViewsRequest struct via the builder pattern


### Return type

[**GetNavigationViews200Response**](GetNavigationViews200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNavigationView

> JsonSuccess RemoveNavigationView(ctx, fragment).Execute()

Remove a navigation view



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tum-zulip/go-zulip/go-zulip"
)

func main() {
	fragment := "narrow/is/alerted" // string | The unique URL hash of the navigation view to be removed.  This also serves as the identifier for the navigation view. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveNavigationView(context.Background(), fragment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NavigationViewsAPI.RemoveNavigationView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNavigationView`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `NavigationViewsAPI.RemoveNavigationView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fragment** | **string** | The unique URL hash of the navigation view to be removed.  This also serves as the identifier for the navigation view.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNavigationViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

