# \ServerAndOrganizationsAPI

All URIs are relative to *https://example.zulipchat.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCodePlayground**](ServerAndOrganizationsAPI.md#AddCodePlayground) | **Post** /realm/playgrounds | Add a code playground
[**AddLinkifier**](ServerAndOrganizationsAPI.md#AddLinkifier) | **Post** /realm/filters | Add a linkifier
[**CreateCustomProfileField**](ServerAndOrganizationsAPI.md#CreateCustomProfileField) | **Post** /realm/profile_fields | Create a custom profile field
[**DeactivateCustomEmoji**](ServerAndOrganizationsAPI.md#DeactivateCustomEmoji) | **Delete** /realm/emoji/{emoji_name} | Deactivate custom emoji
[**ExportRealm**](ServerAndOrganizationsAPI.md#ExportRealm) | **Post** /export/realm | Create a data export
[**GetCustomEmoji**](ServerAndOrganizationsAPI.md#GetCustomEmoji) | **Get** /realm/emoji | Get all custom emoji
[**GetCustomProfileFields**](ServerAndOrganizationsAPI.md#GetCustomProfileFields) | **Get** /realm/profile_fields | Get all custom profile fields
[**GetLinkifiers**](ServerAndOrganizationsAPI.md#GetLinkifiers) | **Get** /realm/linkifiers | Get linkifiers
[**GetPresence**](ServerAndOrganizationsAPI.md#GetPresence) | **Get** /realm/presence | Get presence of all users
[**GetRealmExportConsents**](ServerAndOrganizationsAPI.md#GetRealmExportConsents) | **Get** /export/realm/consents | Get data export consent state
[**GetRealmExports**](ServerAndOrganizationsAPI.md#GetRealmExports) | **Get** /export/realm | Get all data exports
[**GetServerSettings**](ServerAndOrganizationsAPI.md#GetServerSettings) | **Get** /server_settings | Get server settings
[**RemoveCodePlayground**](ServerAndOrganizationsAPI.md#RemoveCodePlayground) | **Delete** /realm/playgrounds/{playground_id} | Remove a code playground
[**RemoveLinkifier**](ServerAndOrganizationsAPI.md#RemoveLinkifier) | **Delete** /realm/filters/{filter_id} | Remove a linkifier
[**ReorderCustomProfileFields**](ServerAndOrganizationsAPI.md#ReorderCustomProfileFields) | **Patch** /realm/profile_fields | Reorder custom profile fields
[**ReorderLinkifiers**](ServerAndOrganizationsAPI.md#ReorderLinkifiers) | **Patch** /realm/linkifiers | Reorder linkifiers
[**TestWelcomeBotCustomMessage**](ServerAndOrganizationsAPI.md#TestWelcomeBotCustomMessage) | **Post** /realm/test_welcome_bot_custom_message | Test welcome bot custom message
[**UpdateLinkifier**](ServerAndOrganizationsAPI.md#UpdateLinkifier) | **Patch** /realm/filters/{filter_id} | Update a linkifier
[**UpdateRealmUserSettingsDefaults**](ServerAndOrganizationsAPI.md#UpdateRealmUserSettingsDefaults) | **Patch** /realm/user_settings_defaults | Update realm-level defaults of user settings
[**UploadCustomEmoji**](ServerAndOrganizationsAPI.md#UploadCustomEmoji) | **Post** /realm/emoji/{emoji_name} | Upload custom emoji



## AddCodePlayground

> AddCodePlayground200Response AddCodePlayground(ctx).Name(name).PygmentsLanguage(pygmentsLanguage).UrlTemplate(urlTemplate).Execute()

Add a code playground



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
	name := "name_example" // string | The user-visible display name of the playground which can be used to pick the target playground, especially when multiple playground options exist for that programming language. 
	pygmentsLanguage := "pygmentsLanguage_example" // string | The name of the Pygments language lexer for that programming language. 
	urlTemplate := "urlTemplate_example" // string | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template for the playground. The template should contain exactly one variable named `code`, which determines how the extracted code should be substituted in the playground URL.  **Changes**: New in Zulip 8.0 (feature level 196). This replaced the `url_prefix` parameter, which was used to construct URLs by just concatenating `url_prefix` and `code`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddCodePlayground(context.Background()).Name(name).PygmentsLanguage(pygmentsLanguage).UrlTemplate(urlTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.AddCodePlayground``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCodePlayground`: AddCodePlayground200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.AddCodePlayground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCodePlaygroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The user-visible display name of the playground which can be used to pick the target playground, especially when multiple playground options exist for that programming language.  | 
 **pygmentsLanguage** | **string** | The name of the Pygments language lexer for that programming language.  | 
 **urlTemplate** | **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template for the playground. The template should contain exactly one variable named &#x60;code&#x60;, which determines how the extracted code should be substituted in the playground URL.  **Changes**: New in Zulip 8.0 (feature level 196). This replaced the &#x60;url_prefix&#x60; parameter, which was used to construct URLs by just concatenating &#x60;url_prefix&#x60; and &#x60;code&#x60;.  | 

### Return type

[**AddCodePlayground200Response**](AddCodePlayground200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLinkifier

> AddLinkifier200Response AddLinkifier(ctx).Pattern(pattern).UrlTemplate(urlTemplate).Execute()

Add a linkifier



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
	pattern := "pattern_example" // string | The [Python regular expression](https://docs.python.org/3/howto/regex.html) that should trigger the linkifier. 
	urlTemplate := "urlTemplate_example" // string | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template used for the link. If you used named groups in `pattern`, you can insert their content here with `{name_of_group}`.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced the `url_format_string` parameter, which was a format string in which named groups' content could be inserted with `%(name_of_group)s`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddLinkifier(context.Background()).Pattern(pattern).UrlTemplate(urlTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.AddLinkifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLinkifier`: AddLinkifier200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.AddLinkifier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLinkifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pattern** | **string** | The [Python regular expression](https://docs.python.org/3/howto/regex.html) that should trigger the linkifier.  | 
 **urlTemplate** | **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template used for the link. If you used named groups in &#x60;pattern&#x60;, you can insert their content here with &#x60;{name_of_group}&#x60;.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced the &#x60;url_format_string&#x60; parameter, which was a format string in which named groups&#39; content could be inserted with &#x60;%(name_of_group)s&#x60;.  | 

### Return type

[**AddLinkifier200Response**](AddLinkifier200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomProfileField

> CreateCustomProfileField200Response CreateCustomProfileField(ctx).FieldType(fieldType).Name(name).Hint(hint).FieldData(fieldData).DisplayInProfileSummary(displayInProfileSummary).Required(required).EditableByUser(editableByUser).Execute()

Create a custom profile field



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
	fieldType := int32(56) // int32 | The field type can be any of the supported custom profile field types. See the [custom profile fields documentation](zulip.com/help/custom-profile-fields for more details on what each type means.  - **1**: Short text - **2**: Long text - **3**: List of options - **4**: Date picker - **5**: Link - **6**: Person picker - **7**: External account - **8**: Pronouns  **Changes**: Field type `8` added in Zulip 6.0 (feature level 151). 
	name := "name_example" // string | The name of the custom profile field, which will appear both in user-facing settings UI for configuring custom profile fields and in UI displaying a user's profile.  (optional)
	hint := "hint_example" // string | The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.  (optional)
	fieldData := map[string]interface{}{ ... } // map[string]interface{} | Field types 3 (List of options) and 7 (External account) support storing additional configuration for the field type in the `field_data` attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type 7 is not yet stabilized.  (optional)
	displayInProfileSummary := true // bool | Whether clients should display this profile field in a summary section of a user's profile (or in a more easily accessible \\\"small profile\\\").  At most 2 profile fields may have this property be true in a given organization. The \\\"Long text\\\" [profile field types][profile-field-types] profile field types cannot be selected to be displayed in profile summaries.  The \\\"Person picker\\\" profile field is also not supported, but that is likely to be temporary.  [profile-field-types]: /help/custom-profile-fields#profile-field-types  **Changes**: New in Zulip 6.0 (feature level 146).  (optional)
	required := true // bool | Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.  **Changes**: New in Zulip 9.0 (feature level 244).  (optional)
	editableByUser := true // bool | Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.  **Changes**: New in Zulip 10.0 (feature level 296).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateCustomProfileField(context.Background()).FieldType(fieldType).Name(name).Hint(hint).FieldData(fieldData).DisplayInProfileSummary(displayInProfileSummary).Required(required).EditableByUser(editableByUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.CreateCustomProfileField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomProfileField`: CreateCustomProfileField200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.CreateCustomProfileField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomProfileFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldType** | **int32** | The field type can be any of the supported custom profile field types. See the [custom profile fields documentation](zulip.com/help/custom-profile-fields for more details on what each type means.  - **1**: Short text - **2**: Long text - **3**: List of options - **4**: Date picker - **5**: Link - **6**: Person picker - **7**: External account - **8**: Pronouns  **Changes**: Field type &#x60;8&#x60; added in Zulip 6.0 (feature level 151).  | 
 **name** | **string** | The name of the custom profile field, which will appear both in user-facing settings UI for configuring custom profile fields and in UI displaying a user&#39;s profile.  | 
 **hint** | **string** | The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.  | 
 **fieldData** | [**map[string]interface{}**](map[string]interface{}.md) | Field types 3 (List of options) and 7 (External account) support storing additional configuration for the field type in the &#x60;field_data&#x60; attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type 7 is not yet stabilized.  | 
 **displayInProfileSummary** | **bool** | Whether clients should display this profile field in a summary section of a user&#39;s profile (or in a more easily accessible \\\&quot;small profile\\\&quot;).  At most 2 profile fields may have this property be true in a given organization. The \\\&quot;Long text\\\&quot; [profile field types][profile-field-types] profile field types cannot be selected to be displayed in profile summaries.  The \\\&quot;Person picker\\\&quot; profile field is also not supported, but that is likely to be temporary.  [profile-field-types]: /help/custom-profile-fields#profile-field-types  **Changes**: New in Zulip 6.0 (feature level 146).  | 
 **required** | **bool** | Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.  **Changes**: New in Zulip 9.0 (feature level 244).  | 
 **editableByUser** | **bool** | Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.  **Changes**: New in Zulip 10.0 (feature level 296).  | 

### Return type

[**CreateCustomProfileField200Response**](CreateCustomProfileField200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateCustomEmoji

> JsonSuccess DeactivateCustomEmoji(ctx, emojiName).Execute()

Deactivate custom emoji



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
	emojiName := "green_tick" // string | The name of the custom emoji to deactivate. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeactivateCustomEmoji(context.Background(), emojiName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.DeactivateCustomEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateCustomEmoji`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.DeactivateCustomEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emojiName** | **string** | The name of the custom emoji to deactivate.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateCustomEmojiRequest struct via the builder pattern


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


## ExportRealm

> ExportRealm200Response ExportRealm(ctx).ExportType(exportType).Execute()

Create a data export



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
	exportType := int32(56) // int32 | Whether to create a public or a standard data export.  - 1 = Public data export. - 2 = Standard data export.  If not specified, defaults to 1.  **Changes**: New in Zulip 10.0 (feature level 304). Previously, all export requests were public data exports.  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportRealm(context.Background()).ExportType(exportType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.ExportRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportRealm`: ExportRealm200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.ExportRealm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportType** | **int32** | Whether to create a public or a standard data export.  - 1 &#x3D; Public data export. - 2 &#x3D; Standard data export.  If not specified, defaults to 1.  **Changes**: New in Zulip 10.0 (feature level 304). Previously, all export requests were public data exports.  | [default to 1]

### Return type

[**ExportRealm200Response**](ExportRealm200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEmoji

> GetCustomEmoji200Response GetCustomEmoji(ctx).Execute()

Get all custom emoji



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
	resp, r, err := apiClient.GetCustomEmoji(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetCustomEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomEmoji`: GetCustomEmoji200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetCustomEmoji`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEmojiRequest struct via the builder pattern


### Return type

[**GetCustomEmoji200Response**](GetCustomEmoji200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomProfileFields

> GetCustomProfileFields200Response GetCustomProfileFields(ctx).Execute()

Get all custom profile fields



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
	resp, r, err := apiClient.GetCustomProfileFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetCustomProfileFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomProfileFields`: GetCustomProfileFields200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetCustomProfileFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomProfileFieldsRequest struct via the builder pattern


### Return type

[**GetCustomProfileFields200Response**](GetCustomProfileFields200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkifiers

> GetLinkifiers200Response GetLinkifiers(ctx).Execute()

Get linkifiers



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
	resp, r, err := apiClient.GetLinkifiers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetLinkifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkifiers`: GetLinkifiers200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetLinkifiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkifiersRequest struct via the builder pattern


### Return type

[**GetLinkifiers200Response**](GetLinkifiers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresence

> GetPresence200Response GetPresence(ctx).Execute()

Get presence of all users



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
	resp, r, err := apiClient.GetPresence(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetPresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresence`: GetPresence200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetPresence`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresenceRequest struct via the builder pattern


### Return type

[**GetPresence200Response**](GetPresence200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealmExportConsents

> GetRealmExportConsents200Response GetRealmExportConsents(ctx).Execute()

Get data export consent state



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
	resp, r, err := apiClient.GetRealmExportConsents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetRealmExportConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealmExportConsents`: GetRealmExportConsents200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetRealmExportConsents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmExportConsentsRequest struct via the builder pattern


### Return type

[**GetRealmExportConsents200Response**](GetRealmExportConsents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealmExports

> GetRealmExports200Response GetRealmExports(ctx).Execute()

Get all data exports



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
	resp, r, err := apiClient.GetRealmExports(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetRealmExports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealmExports`: GetRealmExports200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetRealmExports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmExportsRequest struct via the builder pattern


### Return type

[**GetRealmExports200Response**](GetRealmExports200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerSettings

> GetServerSettings200Response GetServerSettings(ctx).Execute()

Get server settings



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
	resp, r, err := apiClient.GetServerSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.GetServerSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerSettings`: GetServerSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.GetServerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerSettingsRequest struct via the builder pattern


### Return type

[**GetServerSettings200Response**](GetServerSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCodePlayground

> JsonSuccess RemoveCodePlayground(ctx, playgroundId).Execute()

Remove a code playground



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
	playgroundId := int32(1) // int32 | The ID of the playground that you want to remove. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveCodePlayground(context.Background(), playgroundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.RemoveCodePlayground``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCodePlayground`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.RemoveCodePlayground`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playgroundId** | **int32** | The ID of the playground that you want to remove.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCodePlaygroundRequest struct via the builder pattern


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


## RemoveLinkifier

> JsonSuccess RemoveLinkifier(ctx, filterId).Execute()

Remove a linkifier



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
	filterId := int32(43) // int32 | The ID of the linkifier that you want to remove. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoveLinkifier(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.RemoveLinkifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLinkifier`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.RemoveLinkifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32** | The ID of the linkifier that you want to remove.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLinkifierRequest struct via the builder pattern


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


## ReorderCustomProfileFields

> JsonSuccess ReorderCustomProfileFields(ctx).Order(order).Execute()

Reorder custom profile fields



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
	order := []int32{int32(123)} // []int32 | A list of the IDs of all the custom profile fields defined in this organization, in the desired new order. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReorderCustomProfileFields(context.Background()).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.ReorderCustomProfileFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderCustomProfileFields`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.ReorderCustomProfileFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderCustomProfileFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **[]int32** | A list of the IDs of all the custom profile fields defined in this organization, in the desired new order.  | 

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


## ReorderLinkifiers

> JsonSuccess ReorderLinkifiers(ctx).OrderedLinkifierIds(orderedLinkifierIds).Execute()

Reorder linkifiers



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
	orderedLinkifierIds := []int32{int32(123)} // []int32 | A list of the IDs of all the linkifiers defined in this organization, in the desired new order. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReorderLinkifiers(context.Background()).OrderedLinkifierIds(orderedLinkifierIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.ReorderLinkifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderLinkifiers`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.ReorderLinkifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderLinkifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderedLinkifierIds** | **[]int32** | A list of the IDs of all the linkifiers defined in this organization, in the desired new order.  | 

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


## TestWelcomeBotCustomMessage

> TestWelcomeBotCustomMessage200Response TestWelcomeBotCustomMessage(ctx).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()

Test welcome bot custom message



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
	welcomeMessageCustomText := "welcomeMessageCustomText_example" // string | Custom message text, in Zulip Markdown format, to be used for this test message.  Maximum length is 8000 characters. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestWelcomeBotCustomMessage(context.Background()).WelcomeMessageCustomText(welcomeMessageCustomText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.TestWelcomeBotCustomMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWelcomeBotCustomMessage`: TestWelcomeBotCustomMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.TestWelcomeBotCustomMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestWelcomeBotCustomMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **welcomeMessageCustomText** | **string** | Custom message text, in Zulip Markdown format, to be used for this test message.  Maximum length is 8000 characters.  | 

### Return type

[**TestWelcomeBotCustomMessage200Response**](TestWelcomeBotCustomMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkifier

> JsonSuccess UpdateLinkifier(ctx, filterId).Pattern(pattern).UrlTemplate(urlTemplate).Execute()

Update a linkifier



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
	filterId := int32(5) // int32 | The ID of the linkifier that you want to update. 
	pattern := "pattern_example" // string | The [Python regular expression](https://docs.python.org/3/howto/regex.html) that should trigger the linkifier. 
	urlTemplate := "urlTemplate_example" // string | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template used for the link. If you used named groups in `pattern`, you can insert their content here with `{name_of_group}`.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced the `url_format_string` parameter, which was a format string in which named groups' content could be inserted with `%(name_of_group)s`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateLinkifier(context.Background(), filterId).Pattern(pattern).UrlTemplate(urlTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.UpdateLinkifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLinkifier`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.UpdateLinkifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32** | The ID of the linkifier that you want to update.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pattern** | **string** | The [Python regular expression](https://docs.python.org/3/howto/regex.html) that should trigger the linkifier.  | 
 **urlTemplate** | **string** | The [RFC 6570](https://www.rfc-editor.org/rfc/rfc6570.html) compliant URL template used for the link. If you used named groups in &#x60;pattern&#x60;, you can insert their content here with &#x60;{name_of_group}&#x60;.  **Changes**: New in Zulip 7.0 (feature level 176). This replaced the &#x60;url_format_string&#x60; parameter, which was a format string in which named groups&#39; content could be inserted with &#x60;%(name_of_group)s&#x60;.  | 

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


## UpdateRealmUserSettingsDefaults

> IgnoredParametersSuccess UpdateRealmUserSettingsDefaults(ctx).StarredMessageCounts(starredMessageCounts).ReceivesTypingNotifications(receivesTypingNotifications).WebSuggestUpdateTimezone(webSuggestUpdateTimezone).FluidLayoutWidth(fluidLayoutWidth).HighContrastMode(highContrastMode).WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy).WebChannelDefaultView(webChannelDefaultView).WebFontSizePx(webFontSizePx).WebLineHeightPercent(webLineHeightPercent).ColorScheme(colorScheme).EnableDraftsSynchronization(enableDraftsSynchronization).TranslateEmoticons(translateEmoticons).DisplayEmojiReactionUsers(displayEmojiReactionUsers).WebHomeView(webHomeView).WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView).LeftSideUserlist(leftSideUserlist).Emojiset(emojiset).DemoteInactiveStreams(demoteInactiveStreams).UserListStyle(userListStyle).WebAnimateImagePreviews(webAnimateImagePreviews).WebStreamUnreadsCountDisplayPolicy(webStreamUnreadsCountDisplayPolicy).HideAiFeatures(hideAiFeatures).WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders).WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary).EnableStreamDesktopNotifications(enableStreamDesktopNotifications).EnableStreamEmailNotifications(enableStreamEmailNotifications).EnableStreamPushNotifications(enableStreamPushNotifications).EnableStreamAudibleNotifications(enableStreamAudibleNotifications).NotificationSound(notificationSound).EnableDesktopNotifications(enableDesktopNotifications).EnableSounds(enableSounds).EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications).EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications).EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications).EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications).EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds).EnableOfflineEmailNotifications(enableOfflineEmailNotifications).EnableOfflinePushNotifications(enableOfflinePushNotifications).EnableOnlinePushNotifications(enableOnlinePushNotifications).EnableDigestEmails(enableDigestEmails).MessageContentInEmailNotifications(messageContentInEmailNotifications).PmContentInDesktopNotifications(pmContentInDesktopNotifications).WildcardMentionsNotify(wildcardMentionsNotify).EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify).DesktopIconCountDisplay(desktopIconCountDisplay).RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy).AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy).AutomaticallyUnmuteTopicsInMutedStreamsPolicy(automaticallyUnmuteTopicsInMutedStreamsPolicy).AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned).ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy).PresenceEnabled(presenceEnabled).EnterSends(enterSends).TwentyFourHourTime(twentyFourHourTime).SendPrivateTypingNotifications(sendPrivateTypingNotifications).SendStreamTypingNotifications(sendStreamTypingNotifications).SendReadReceipts(sendReadReceipts).EmailAddressVisibility(emailAddressVisibility).WebNavigateToSentMessage(webNavigateToSentMessage).Execute()

Update realm-level defaults of user settings



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
	starredMessageCounts := true // bool | Whether clients should display the [number of starred messages](zulip.com/help/star-a-message#display-the-number-of-starred-messages.  (optional)
	receivesTypingNotifications := true // bool | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  (optional)
	webSuggestUpdateTimezone := true // bool | Whether the user should be shown an alert, offering to update their [profile time zone](zulip.com/help/change-your-timezone, when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).  (optional)
	fluidLayoutWidth := true // bool | Whether to use the [maximum available screen width](zulip.com/help/enable-full-width-display for the web app's center panel (message feed, recent conversations) on wide screens.  (optional)
	highContrastMode := true // bool | This setting is reserved for use to control variations in Zulip's design to help visually impaired users.  (optional)
	webMarkReadOnScrollPolicy := int32(56) // int32 | Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\\"Always\\\" setting when marking messages as read.  (optional)
	webChannelDefaultView := int32(56) // int32 | Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\\"Top unread topic in channel\\\" is new in Zulip 11.0 (feature level 401).  The \\\"List of topics\\\" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\\"Channel feed\\\" behavior.  (optional)
	webFontSizePx := int32(56) // int32 | User-configured primary `font-size` for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.  (optional)
	webLineHeightPercent := int32(56) // int32 | User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.  (optional)
	colorScheme := int32(56) // int32 | Controls which [color theme](zulip.com/help/dark-theme to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard `prefers-color-scheme` media query.  (optional)
	enableDraftsSynchronization := true // bool | A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  (optional)
	translateEmoticons := true // bool | Whether to [translate emoticons to emoji](zulip.com/help/configure-emoticon-translations in messages the user sends.  (optional)
	displayEmojiReactionUsers := true // bool | Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).  (optional)
	webHomeView := "webHomeView_example" // string | The [home view](zulip.com/help/configure-home-view used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.  - \\\"recent_topics\\\" - Recent conversations view - \\\"inbox\\\" - Inbox view - \\\"all_messages\\\" - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).  (optional)
	webEscapeNavigatesToHomeView := true // bool | Whether the escape key navigates to the [configured home view](zulip.com/help/configure-home-view.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).  (optional)
	leftSideUserlist := true // bool | Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  (optional)
	emojiset := "emojiset_example" // string | The user's configured [emoji set](zulip.com/help/emoji-and-emoticons#use-emoticons, used to display emoji to the user everywhere they appear in the UI.  - \\\"google\\\" - Google - \\\"twitter\\\" - Twitter - \\\"text\\\" - Plain text - \\\"google-blob\\\" - Google blobs  (optional)
	demoteInactiveStreams := int32(56) // int32 | Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  (optional)
	userListStyle := int32(56) // int32 | The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).  (optional)
	webAnimateImagePreviews := "webAnimateImagePreviews_example" // string | Controls how animated images should be played in the message feed in the web/desktop application.  - \\\"always\\\" - Always play the animated images in the message feed. - \\\"on_hover\\\" - Play the animated images on hover over them in the message feed. - \\\"never\\\" - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275). Previously, animated images always used to play in the message feed by default. This setting controls this behaviour.  (optional)
	webStreamUnreadsCountDisplayPolicy := int32(56) // int32 | Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).  (optional)
	hideAiFeatures := true // bool | Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).  (optional)
	webLeftSidebarShowChannelFolders := true // bool | Determines whether the web/desktop application's left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).  (optional)
	webLeftSidebarUnreadsCountSummary := true // bool | Determines whether the web/desktop application's left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).  (optional)
	enableStreamDesktopNotifications := true // bool | Enable visual desktop notifications for channel messages.  (optional)
	enableStreamEmailNotifications := true // bool | Enable email notifications for channel messages.  (optional)
	enableStreamPushNotifications := true // bool | Enable mobile notifications for channel messages.  (optional)
	enableStreamAudibleNotifications := true // bool | Enable audible desktop notifications for channel messages.  (optional)
	notificationSound := "notificationSound_example" // string | Notification sound name.  (optional)
	enableDesktopNotifications := true // bool | Enable visual desktop notifications for direct messages and @-mentions.  (optional)
	enableSounds := true // bool | Enable audible desktop notifications for direct messages and @-mentions.  (optional)
	enableFollowedTopicDesktopNotifications := true // bool | Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicEmailNotifications := true // bool | Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicPushNotifications := true // bool | Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	enableFollowedTopicAudibleNotifications := true // bool | Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	emailNotificationsBatchingPeriodSeconds := int32(56) // int32 | The duration (in seconds) for which the server should wait to batch email notifications before sending them.  (optional)
	enableOfflineEmailNotifications := true // bool | Enable email notifications for direct messages and @-mentions received when the user is offline.  (optional)
	enableOfflinePushNotifications := true // bool | Enable mobile notification for direct messages and @-mentions received when the user is offline.  (optional)
	enableOnlinePushNotifications := true // bool | Enable mobile notification for direct messages and @-mentions received when the user is online.  (optional)
	enableDigestEmails := true // bool | Enable digest emails when the user is away.  (optional)
	messageContentInEmailNotifications := true // bool | Include the message's content in email notifications for new messages.  (optional)
	pmContentInDesktopNotifications := true // bool | Include content of direct messages in desktop notifications.  (optional)
	wildcardMentionsNotify := true // bool | Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  (optional)
	enableFollowedTopicWildcardMentionsNotify := true // bool | Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).  (optional)
	desktopIconCountDisplay := int32(56) // int32 | Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.  (optional)
	realmNameInEmailNotificationsPolicy := int32(56) // int32 | Whether to [include organization name in subject of message notification emails](zulip.com/help/email-notifications#include-organization-name-in-subject-line.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.  (optional)
	automaticallyFollowTopicsPolicy := int32(56) // int32 | Which [topics to follow automatically](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  (optional)
	automaticallyUnmuteTopicsInMutedStreamsPolicy := int32(56) // int32 | Which [topics to unmute automatically in muted channels](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  (optional)
	automaticallyFollowTopicsWhereMentioned := true // bool | Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).  (optional)
	resolvedTopicNoticeAutoReadPolicy := "resolvedTopicNoticeAutoReadPolicy_example" // string | Controls whether the resolved-topic notices are marked as read.  - \\\"always\\\" - Always mark resolved-topic notices as read. - \\\"except_followed\\\" - Mark resolved-topic notices as read in topics not followed by the user. - \\\"never\\\" - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).  (optional)
	presenceEnabled := true // bool | Display the presence status to other users when online.  (optional)
	enterSends := true // bool | Whether pressing Enter in the compose box sends a message (or saves a message edit).  (optional)
	twentyFourHourTime := true // bool | Whether time should be [displayed in 24-hour notation](zulip.com/help/change-the-time-format.  **Changes**: New in Zulip 5.0 (feature level 99). Previously, this default was edited using the `default_twenty_four_hour_time` parameter to the `PATCH /realm` endpoint.  (optional)
	sendPrivateTypingNotifications := true // bool | Whether [typing notifications](zulip.com/help/typing-notifications be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	sendStreamTypingNotifications := true // bool | Whether [typing notifications](zulip.com/help/typing-notifications be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	sendReadReceipts := true // bool | Whether other users are allowed to see whether you've read messages.  **Changes**: New in Zulip 5.0 (feature level 105).  (optional)
	emailAddressVisibility := int32(56) // int32 | The [policy][permission-level] for [which other users][help-email-visibility] in this organization can see the user's real email address.  - 1 = Everyone - 2 = Members only - 3 = Administrators only - 4 = Nobody - 5 = Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility  (optional)
	webNavigateToSentMessage := true // bool | Web/desktop app setting for whether the user's view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateRealmUserSettingsDefaults(context.Background()).StarredMessageCounts(starredMessageCounts).ReceivesTypingNotifications(receivesTypingNotifications).WebSuggestUpdateTimezone(webSuggestUpdateTimezone).FluidLayoutWidth(fluidLayoutWidth).HighContrastMode(highContrastMode).WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy).WebChannelDefaultView(webChannelDefaultView).WebFontSizePx(webFontSizePx).WebLineHeightPercent(webLineHeightPercent).ColorScheme(colorScheme).EnableDraftsSynchronization(enableDraftsSynchronization).TranslateEmoticons(translateEmoticons).DisplayEmojiReactionUsers(displayEmojiReactionUsers).WebHomeView(webHomeView).WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView).LeftSideUserlist(leftSideUserlist).Emojiset(emojiset).DemoteInactiveStreams(demoteInactiveStreams).UserListStyle(userListStyle).WebAnimateImagePreviews(webAnimateImagePreviews).WebStreamUnreadsCountDisplayPolicy(webStreamUnreadsCountDisplayPolicy).HideAiFeatures(hideAiFeatures).WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders).WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary).EnableStreamDesktopNotifications(enableStreamDesktopNotifications).EnableStreamEmailNotifications(enableStreamEmailNotifications).EnableStreamPushNotifications(enableStreamPushNotifications).EnableStreamAudibleNotifications(enableStreamAudibleNotifications).NotificationSound(notificationSound).EnableDesktopNotifications(enableDesktopNotifications).EnableSounds(enableSounds).EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications).EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications).EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications).EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications).EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds).EnableOfflineEmailNotifications(enableOfflineEmailNotifications).EnableOfflinePushNotifications(enableOfflinePushNotifications).EnableOnlinePushNotifications(enableOnlinePushNotifications).EnableDigestEmails(enableDigestEmails).MessageContentInEmailNotifications(messageContentInEmailNotifications).PmContentInDesktopNotifications(pmContentInDesktopNotifications).WildcardMentionsNotify(wildcardMentionsNotify).EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify).DesktopIconCountDisplay(desktopIconCountDisplay).RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy).AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy).AutomaticallyUnmuteTopicsInMutedStreamsPolicy(automaticallyUnmuteTopicsInMutedStreamsPolicy).AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned).ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy).PresenceEnabled(presenceEnabled).EnterSends(enterSends).TwentyFourHourTime(twentyFourHourTime).SendPrivateTypingNotifications(sendPrivateTypingNotifications).SendStreamTypingNotifications(sendStreamTypingNotifications).SendReadReceipts(sendReadReceipts).EmailAddressVisibility(emailAddressVisibility).WebNavigateToSentMessage(webNavigateToSentMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.UpdateRealmUserSettingsDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRealmUserSettingsDefaults`: IgnoredParametersSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.UpdateRealmUserSettingsDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRealmUserSettingsDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **starredMessageCounts** | **bool** | Whether clients should display the [number of starred messages](zulip.com/help/star-a-message#display-the-number-of-starred-messages.  | 
 **receivesTypingNotifications** | **bool** | Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.  **Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.  | 
 **webSuggestUpdateTimezone** | **bool** | Whether the user should be shown an alert, offering to update their [profile time zone](zulip.com/help/change-your-timezone, when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.  **Changes**: New in Zulip 10.0 (feature level 329).  | 
 **fluidLayoutWidth** | **bool** | Whether to use the [maximum available screen width](zulip.com/help/enable-full-width-display for the web app&#39;s center panel (message feed, recent conversations) on wide screens.  | 
 **highContrastMode** | **bool** | This setting is reserved for use to control variations in Zulip&#39;s design to help visually impaired users.  | 
 **webMarkReadOnScrollPolicy** | **int32** | Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\\&quot;Always\\\&quot; setting when marking messages as read.  | 
 **webChannelDefaultView** | **int32** | Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\\&quot;Top unread topic in channel\\\&quot; is new in Zulip 11.0 (feature level 401).  The \\\&quot;List of topics\\\&quot; option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\\&quot;Channel feed\\\&quot; behavior.  | 
 **webFontSizePx** | **int32** | User-configured primary &#x60;font-size&#x60; for the web application, in pixels.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.  | 
 **webLineHeightPercent** | **int32** | User-configured primary &#x60;line-height&#x60; for the web application, in percent, so a value of 120 represents a &#x60;line-height&#x60; of 1.2.  **Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.  | 
 **colorScheme** | **int32** | Controls which [color theme](zulip.com/help/dark-theme to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard &#x60;prefers-color-scheme&#x60; media query.  | 
 **enableDraftsSynchronization** | **bool** | A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.  | 
 **translateEmoticons** | **bool** | Whether to [translate emoticons to emoji](zulip.com/help/configure-emoticon-translations in messages the user sends.  | 
 **displayEmojiReactionUsers** | **bool** | Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.  **Changes**: New in Zulip 6.0 (feature level 125).  | 
 **webHomeView** | **string** | The [home view](zulip.com/help/configure-home-view used when opening a new Zulip web app window or hitting the &#x60;Esc&#x60; keyboard shortcut repeatedly.  - \\\&quot;recent_topics\\\&quot; - Recent conversations view - \\\&quot;inbox\\\&quot; - Inbox view - \\\&quot;all_messages\\\&quot; - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;default_view&#x60;, which was new in Zulip 4.0 (feature level 42).  | 
 **webEscapeNavigatesToHomeView** | **bool** | Whether the escape key navigates to the [configured home view](zulip.com/help/configure-home-view.  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called &#x60;escape_navigates_to_default_view&#x60;, which was new in Zulip 5.0 (feature level 107).  | 
 **leftSideUserlist** | **bool** | Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.  | 
 **emojiset** | **string** | The user&#39;s configured [emoji set](zulip.com/help/emoji-and-emoticons#use-emoticons, used to display emoji to the user everywhere they appear in the UI.  - \\\&quot;google\\\&quot; - Google - \\\&quot;twitter\\\&quot; - Twitter - \\\&quot;text\\\&quot; - Plain text - \\\&quot;google-blob\\\&quot; - Google blobs  | 
 **demoteInactiveStreams** | **int32** | Whether to [hide inactive channels](zulip.com/help/manage-inactive-channels in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never  | 
 **userListStyle** | **int32** | The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).  | 
 **webAnimateImagePreviews** | **string** | Controls how animated images should be played in the message feed in the web/desktop application.  - \\\&quot;always\\\&quot; - Always play the animated images in the message feed. - \\\&quot;on_hover\\\&quot; - Play the animated images on hover over them in the message feed. - \\\&quot;never\\\&quot; - Never play animated images in the message feed.  **Changes**: New in Zulip 9.0 (feature level 275). Previously, animated images always used to play in the message feed by default. This setting controls this behaviour.  | 
 **webStreamUnreadsCountDisplayPolicy** | **int32** | Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).  | 
 **hideAiFeatures** | **bool** | Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.  **Changes**: New in Zulip 10.0 (feature level 350).  | 
 **webLeftSidebarShowChannelFolders** | **bool** | Determines whether the web/desktop application&#39;s left sidebar displays any channel folders configured by the organization.  **Changes**: New in Zulip 11.0 (feature level 411).  | 
 **webLeftSidebarUnreadsCountSummary** | **bool** | Determines whether the web/desktop application&#39;s left sidebar displays the unread message count summary.  **Changes**: New in Zulip 11.0 (feature level 398).  | 
 **enableStreamDesktopNotifications** | **bool** | Enable visual desktop notifications for channel messages.  | 
 **enableStreamEmailNotifications** | **bool** | Enable email notifications for channel messages.  | 
 **enableStreamPushNotifications** | **bool** | Enable mobile notifications for channel messages.  | 
 **enableStreamAudibleNotifications** | **bool** | Enable audible desktop notifications for channel messages.  | 
 **notificationSound** | **string** | Notification sound name.  | 
 **enableDesktopNotifications** | **bool** | Enable visual desktop notifications for direct messages and @-mentions.  | 
 **enableSounds** | **bool** | Enable audible desktop notifications for direct messages and @-mentions.  | 
 **enableFollowedTopicDesktopNotifications** | **bool** | Enable visual desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicEmailNotifications** | **bool** | Enable email notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicPushNotifications** | **bool** | Enable push notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **enableFollowedTopicAudibleNotifications** | **bool** | Enable audible desktop notifications for messages sent to followed topics.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **emailNotificationsBatchingPeriodSeconds** | **int32** | The duration (in seconds) for which the server should wait to batch email notifications before sending them.  | 
 **enableOfflineEmailNotifications** | **bool** | Enable email notifications for direct messages and @-mentions received when the user is offline.  | 
 **enableOfflinePushNotifications** | **bool** | Enable mobile notification for direct messages and @-mentions received when the user is offline.  | 
 **enableOnlinePushNotifications** | **bool** | Enable mobile notification for direct messages and @-mentions received when the user is online.  | 
 **enableDigestEmails** | **bool** | Enable digest emails when the user is away.  | 
 **messageContentInEmailNotifications** | **bool** | Include the message&#39;s content in email notifications for new messages.  | 
 **pmContentInDesktopNotifications** | **bool** | Include content of direct messages in desktop notifications.  | 
 **wildcardMentionsNotify** | **bool** | Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.  | 
 **enableFollowedTopicWildcardMentionsNotify** | **bool** | Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.  **Changes**: New in Zulip 8.0 (feature level 189).  | 
 **desktopIconCountDisplay** | **int32** | Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added &#x60;DMs, mentions, and followed topics&#x60; option, renumbering the options to insert it in order.  | 
 **realmNameInEmailNotificationsPolicy** | **int32** | Whether to [include organization name in subject of message notification emails](zulip.com/help/email-notifications#include-organization-name-in-subject-line.  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous &#x60;realm_name_in_notifications&#x60; boolean; &#x60;true&#x60; corresponded to &#x60;Always&#x60;, and &#x60;false&#x60; to &#x60;Never&#x60;.  | 
 **automaticallyFollowTopicsPolicy** | **int32** | Which [topics to follow automatically](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | 
 **automaticallyUnmuteTopicsInMutedStreamsPolicy** | **int32** | Which [topics to unmute automatically in muted channels](zulip.com/help/mute-a-topic.  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).  | 
 **automaticallyFollowTopicsWhereMentioned** | **bool** | Whether the server will automatically mark the user as following topics where the user is mentioned.  **Changes**: New in Zulip 8.0 (feature level 235).  | 
 **resolvedTopicNoticeAutoReadPolicy** | **string** | Controls whether the resolved-topic notices are marked as read.  - \\\&quot;always\\\&quot; - Always mark resolved-topic notices as read. - \\\&quot;except_followed\\\&quot; - Mark resolved-topic notices as read in topics not followed by the user. - \\\&quot;never\\\&quot; - Never mark resolved-topic notices as read.  **Changes**: New in Zulip 11.0 (feature level 385).  | 
 **presenceEnabled** | **bool** | Display the presence status to other users when online.  | 
 **enterSends** | **bool** | Whether pressing Enter in the compose box sends a message (or saves a message edit).  | 
 **twentyFourHourTime** | **bool** | Whether time should be [displayed in 24-hour notation](zulip.com/help/change-the-time-format.  **Changes**: New in Zulip 5.0 (feature level 99). Previously, this default was edited using the &#x60;default_twenty_four_hour_time&#x60; parameter to the &#x60;PATCH /realm&#x60; endpoint.  | 
 **sendPrivateTypingNotifications** | **bool** | Whether [typing notifications](zulip.com/help/typing-notifications be sent when composing direct messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **sendStreamTypingNotifications** | **bool** | Whether [typing notifications](zulip.com/help/typing-notifications be sent when composing channel messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **sendReadReceipts** | **bool** | Whether other users are allowed to see whether you&#39;ve read messages.  **Changes**: New in Zulip 5.0 (feature level 105).  | 
 **emailAddressVisibility** | **int32** | The [policy][permission-level] for [which other users][help-email-visibility] in this organization can see the user&#39;s real email address.  - 1 &#x3D; Everyone - 2 &#x3D; Members only - 3 &#x3D; Administrators only - 4 &#x3D; Nobody - 5 &#x3D; Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.  [permission-level]: /api/roles-and-permissions#permission-levels [help-email-visibility]: /help/configure-email-visibility  | 
 **webNavigateToSentMessage** | **bool** | Web/desktop app setting for whether the user&#39;s view should automatically go to the conversation where they sent a message.  **Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.  | 

### Return type

[**IgnoredParametersSuccess**](IgnoredParametersSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCustomEmoji

> JsonSuccess UploadCustomEmoji(ctx, emojiName).Filename(filename).Execute()

Upload custom emoji



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
	emojiName := "smile" // string | The name that should be associated with the uploaded emoji image/gif. The emoji name can only contain letters, numbers, dashes, and spaces. Upper and lower case letters are treated the same, and underscores (\\_) are treated the same as spaces (consistent with how the Zulip UI handles emoji). 
	filename := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadCustomEmoji(context.Background(), emojiName).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerAndOrganizationsAPI.UploadCustomEmoji``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCustomEmoji`: JsonSuccess
	fmt.Fprintf(os.Stdout, "Response from `ServerAndOrganizationsAPI.UploadCustomEmoji`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emojiName** | **string** | The name that should be associated with the uploaded emoji image/gif. The emoji name can only contain letters, numbers, dashes, and spaces. Upper and lower case letters are treated the same, and underscores (\\_) are treated the same as spaces (consistent with how the Zulip UI handles emoji).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCustomEmojiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | ***os.File** |  | 

### Return type

[**JsonSuccess**](JsonSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

