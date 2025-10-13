# \DraftsAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrafts**](DraftsAPI.md#CreateDrafts) | **Post** /drafts | Create drafts
[**CreateSavedSnippet**](DraftsAPI.md#CreateSavedSnippet) | **Post** /saved_snippets | Create a saved snippet
[**DeleteDraft**](DraftsAPI.md#DeleteDraft) | **Delete** /drafts/{draft_id} | Delete a draft
[**DeleteSavedSnippet**](DraftsAPI.md#DeleteSavedSnippet) | **Delete** /saved_snippets/{saved_snippet_id} | Delete a saved snippet
[**EditDraft**](DraftsAPI.md#EditDraft) | **Patch** /drafts/{draft_id} | Edit a draft
[**EditSavedSnippet**](DraftsAPI.md#EditSavedSnippet) | **Patch** /saved_snippets/{saved_snippet_id} | Edit a saved snippet
[**GetDrafts**](DraftsAPI.md#GetDrafts) | **Get** /drafts | Get drafts
[**GetSavedSnippets**](DraftsAPI.md#GetSavedSnippets) | **Get** /saved_snippets | Get all saved snippets



## CreateDrafts

> CreateDrafts200Response CreateDrafts(ctx).Drafts(drafts).Execute()

Create drafts



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
	drafts := []openapiclient.Draft{*openapiclient.NewDraft("Type_example", []int32{int32(123)}, "Topic_example", "Content_example")} // []Draft | A JSON-encoded list of containing new draft objects.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateDrafts(context.Background()).Drafts(drafts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.CreateDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDrafts`: CreateDrafts200Response
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.CreateDrafts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDraftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **drafts** | [**[]Draft**](Draft.md) | A JSON-encoded list of containing new draft objects.  | 

### Return type

[**CreateDrafts200Response**](CreateDrafts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSavedSnippet

> CreateSavedSnippet200Response CreateSavedSnippet(ctx).Title(title).Content(content).Execute()

Create a saved snippet



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
	title := "title_example" // string | The title of the saved snippet. 
	content := "content_example" // string | The content of the saved snippet in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should insert this content into a message when using a saved snippet. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateSavedSnippet(context.Background()).Title(title).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.CreateSavedSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSavedSnippet`: CreateSavedSnippet200Response
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.CreateSavedSnippet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavedSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | The title of the saved snippet.  | 
 **content** | **string** | The content of the saved snippet in [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should insert this content into a message when using a saved snippet.  | 

### Return type

[**CreateSavedSnippet200Response**](CreateSavedSnippet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDraft

> JsonSuccess DeleteDraft(ctx, draftId).Execute()

Delete a draft



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
	draftId := int32(1) // int32 | The ID of the draft you want to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteDraft(context.Background(), draftId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.DeleteDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDraft`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.DeleteDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **int32** | The ID of the draft you want to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDraftRequest struct via the builder pattern


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


## DeleteSavedSnippet

> JsonSuccess DeleteSavedSnippet(ctx, savedSnippetId).Execute()

Delete a saved snippet



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
	savedSnippetId := int32(2) // int32 | The ID of the saved snippet to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteSavedSnippet(context.Background(), savedSnippetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.DeleteSavedSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSavedSnippet`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.DeleteSavedSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedSnippetId** | **int32** | The ID of the saved snippet to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedSnippetRequest struct via the builder pattern


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


## EditDraft

> JsonSuccess EditDraft(ctx, draftId).Draft(draft).Execute()

Edit a draft



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
	draftId := int32(2) // int32 | The ID of the draft to be edited. 
	draft := *openapiclient.NewDraft("Type_example", []int32{int32(123)}, "Topic_example", "Content_example") // Draft | A JSON-encoded object containing a replacement draft object for this ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EditDraft(context.Background(), draftId).Draft(draft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.EditDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDraft`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.EditDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **int32** | The ID of the draft to be edited.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **draft** | [**Draft**](Draft.md) | A JSON-encoded object containing a replacement draft object for this ID.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSavedSnippet

> JsonSuccess EditSavedSnippet(ctx, savedSnippetId).Title(title).Content(content).Execute()

Edit a saved snippet



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
	savedSnippetId := int32(3) // int32 | The ID of the saved snippet to edit. 
	title := "title_example" // string | The title of the saved snippet.  (optional)
	content := "content_example" // string | The content of the saved snippet in the original [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should insert this content into a message when using a saved snippet.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EditSavedSnippet(context.Background(), savedSnippetId).Title(title).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.EditSavedSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditSavedSnippet`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.EditSavedSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedSnippetId** | **int32** | The ID of the saved snippet to edit.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSavedSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | The title of the saved snippet.  | 
 **content** | **string** | The content of the saved snippet in the original [Zulip-flavored Markdown](zulip.com/help/format-your-message-using-markdown format.  Clients should insert this content into a message when using a saved snippet.  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrafts

> GetDrafts200Response GetDrafts(ctx).Execute()

Get drafts



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
	resp, r, err := apiClient.GetDrafts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.GetDrafts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrafts`: GetDrafts200Response
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.GetDrafts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftsRequest struct via the builder pattern


### Return type

[**GetDrafts200Response**](GetDrafts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedSnippets

> GetSavedSnippets200Response GetSavedSnippets(ctx).Execute()

Get all saved snippets



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
	resp, r, err := apiClient.GetSavedSnippets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftsAPI.GetSavedSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedSnippets`: GetSavedSnippets200Response
	fmt.Fprintf(os.Stdout, "Response from `DraftsAPI.GetSavedSnippets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedSnippetsRequest struct via the builder pattern


### Return type

[**GetSavedSnippets200Response**](GetSavedSnippets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

