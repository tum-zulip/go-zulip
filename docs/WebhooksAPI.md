# \WebhooksAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZulipOutgoingWebhooks**](WebhooksAPI.md#ZulipOutgoingWebhooks) | **Post** /zulip-outgoing-webhook | Outgoing webhooks



## ZulipOutgoingWebhooks

> ZulipOutgoingWebhooks200Response ZulipOutgoingWebhooks(ctx).Execute()

Outgoing webhooks



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
	resp, r, err := apiClient.ZulipOutgoingWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ZulipOutgoingWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZulipOutgoingWebhooks`: ZulipOutgoingWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ZulipOutgoingWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZulipOutgoingWebhooksRequest struct via the builder pattern


### Return type

[**ZulipOutgoingWebhooks200Response**](ZulipOutgoingWebhooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

