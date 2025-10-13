# \AuthenticationAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevFetchApiKey**](AuthenticationAPI.md#DevFetchApiKey) | **Post** /dev_fetch_api_key | Fetch an API key (development only)
[**FetchApiKey**](AuthenticationAPI.md#FetchApiKey) | **Post** /fetch_api_key | Fetch an API key (production)



## DevFetchApiKey

> ApiKeyResponse DevFetchApiKey(ctx).Username(username).Execute()

Fetch an API key (development only)



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
	username := "username_example" // string | The email address for the user that owns the API key. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevFetchApiKey(context.Background()).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DevFetchApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevFetchApiKey`: ApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.DevFetchApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevFetchApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The email address for the user that owns the API key.  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchApiKey

> ApiKeyResponse FetchApiKey(ctx).Username(username).Password(password).Execute()

Fetch an API key (production)



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
	username := "username_example" // string | The username to be used for authentication (typically, the email address, but depending on configuration, it could be an LDAP username).  See the `require_email_format_usernames` parameter documented in [GET /server_settings](zulip.com/api/get-server-settings for details. 
	password := "password_example" // string | The user's Zulip password (or LDAP password, if LDAP authentication is in use). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FetchApiKey(context.Background()).Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.FetchApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchApiKey`: ApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.FetchApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | The username to be used for authentication (typically, the email address, but depending on configuration, it could be an LDAP username).  See the &#x60;require_email_format_usernames&#x60; parameter documented in [GET /server_settings](zulip.com/api/get-server-settings for details.  | 
 **password** | **string** | The user&#39;s Zulip password (or LDAP password, if LDAP authentication is in use).  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

