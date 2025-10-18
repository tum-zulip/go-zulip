package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ServerAndOrganizationsAPI interface {

	// AddCodePlayground Add a code playground
	//
	// Configure [code playgrounds] for the organization.
	//
	// *Changes**: New in Zulip 4.0 (feature level 49). A parameter encoding bug was
	// fixed in Zulip 4.0 (feature level 57).
	//
	// [code playgrounds]: https://zulip.com/help/code-blocks#code-playgrounds
	AddCodePlayground(ctx context.Context) AddCodePlaygroundRequest

	// AddCodePlaygroundExecute executes the request
	AddCodePlaygroundExecute(r AddCodePlaygroundRequest) (*AddCodePlaygroundResponse, *http.Response, error)

	// AddLinkifier Add a linkifier
	//
	// Configure [linkifiers],
	// regular expression patterns that are automatically linkified when they
	// appear in messages and topics.
	//
	// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
	AddLinkifier(ctx context.Context) AddLinkifierRequest

	// AddLinkifierExecute executes the request
	AddLinkifierExecute(r AddLinkifierRequest) (*AddLinkifierResponse, *http.Response, error)

	// CreateCustomProfileField Create a custom profile field
	//
	// [Create a custom profile field] in the user's organization.
	//
	// [Create a custom profile field]: https://zulip.com/help/custom-profile-fields#add-a-custom-profile-field
	CreateCustomProfileField(ctx context.Context) CreateCustomProfileFieldRequest

	// CreateCustomProfileFieldExecute executes the request
	CreateCustomProfileFieldExecute(r CreateCustomProfileFieldRequest) (*CreateCustomProfileFieldResponse, *http.Response, error)

	// DeactivateCustomEmoji Deactivate custom emoji
	//
	// [Deactivate a custom emoji] from
	// the user's organization.
	//
	// Users can only deactivate custom emoji that they added themselves except for
	// organization administrators, who can deactivate any custom emoji.
	//
	// Note that deactivated emoji will still be visible in old messages, reactions,
	// user statuses and channel descriptions.
	//
	// *Changes**: Before Zulip 8.0 (feature level 190), this endpoint returned an
	// HTTP status code of 400 when the emoji did not exist, instead of 404.
	//
	// [Deactivate a custom emoji]: https://zulip.com/help/custom-emoji#deactivate-custom-emoji
	DeactivateCustomEmoji(ctx context.Context, emojiName string) DeactivateCustomEmojiRequest

	// DeactivateCustomEmojiExecute executes the request
	DeactivateCustomEmojiExecute(r DeactivateCustomEmojiRequest) (*Response, *http.Response, error)

	// ExportRealm Create a data export
	//
	// Create a public or a standard [data export] of the organization.
	//
	// !!! warn ""
	//
	// *Note**: If you're the administrator of a self-hosted installation,
	// you may be looking for the documentation on [server data export and import] or [server backups].
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 304), only
	// public data exports could be created using this endpoint.
	//
	// New in Zulip 2.1.
	//
	// [data export]: https://zulip.com/help/export-your-organization#export-for-migrating-to-zulip-cloud-or-a-self-hosted-server
	// [server data export and import]: https://zulip.readthedocs.io/en/stable/production/export-and-import.html#data-export
	// [server backups]: https://zulip.readthedocs.io/en/stable/production/export-and-import.html#backups
	//
	ExportRealm(ctx context.Context) ExportRealmRequest

	// ExportRealmExecute executes the request
	ExportRealmExecute(r ExportRealmRequest) (*ExportRealmResponse, *http.Response, error)

	// GetCustomEmoji Get all custom emoji
	//
	// Get all the custom emoji in the user's organization.
	//
	GetCustomEmoji(ctx context.Context) GetCustomEmojiRequest

	// GetCustomEmojiExecute executes the request
	GetCustomEmojiExecute(r GetCustomEmojiRequest) (*GetCustomEmojiResponse, *http.Response, error)

	// GetCustomProfileFields Get all custom profile fields
	//
	// Get all the [custom profile fields]
	// configured for the user's organization.
	//
	// [custom profile fields]: https://zulip.com/help/custom-profile-fields
	GetCustomProfileFields(ctx context.Context) GetCustomProfileFieldsRequest

	// GetCustomProfileFieldsExecute executes the request
	GetCustomProfileFieldsExecute(r GetCustomProfileFieldsRequest) (*GetCustomProfileFieldsResponse, *http.Response, error)

	// GetLinkifiers Get linkifiers
	//
	// List all of an organization's configured
	// [linkifiers], regular
	// expression patterns that are automatically linkified when they appear
	// in messages and topics.
	//
	// *Changes**: New in Zulip 4.0 (feature level 54). On older versions,
	// a similar `GET /realm/filters` endpoint was available with each entry in
	// a `[pattern, url_format, id]` tuple format.
	//
	// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
	GetLinkifiers(ctx context.Context) GetLinkifiersRequest

	// GetLinkifiersExecute executes the request
	GetLinkifiersExecute(r GetLinkifiersRequest) (*GetLinkifiersResponse, *http.Response, error)

	// GetPresence Get presence of all users
	//
	// Get the presence information of all the users in an organization.
	//
	// If the `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level
	// setting is set to `true`, presence information of only accessible
	// users are returned.
	//
	// Complete Zulip apps are recommended to fetch presence
	// information when they post their own state using the [`POST
	// /presence`](https://zulip.com/api/update-presence) API endpoint.
	//
	GetPresence(ctx context.Context) GetPresenceRequest

	// GetPresenceExecute executes the request
	GetPresenceExecute(r GetPresenceRequest) (*GetPresenceResponse, *http.Response, error)

	// GetRealmExportConsents Get data export consent state
	//
	// Fetches which users have [consented]
	// for their private data to be exported by organization administrators.
	//
	// *Changes**: New in Zulip 10.0 (feature level 295).
	//
	// [consented]: https://zulip.com/help/export-your-organization#configure-whether-administrators-can-export-your-private-data
	GetRealmExportConsents(ctx context.Context) GetRealmExportConsentsRequest

	// GetRealmExportConsentsExecute executes the request
	GetRealmExportConsentsExecute(r GetRealmExportConsentsRequest) (*GetRealmExportConsentsResponse, *http.Response, error)

	// GetRealmExports Get all data exports
	//
	// Fetch all the public and standard [data exports]
	// of the organization.
	//
	// *Changes**: Prior to Zulip 10.0 (feature level 304), only
	// public data exports could be fetched using this endpoint.
	//
	// New in Zulip 2.1.
	//
	// [data exports]: https://zulip.com/help/export-your-organization#export-for-migrating-to-zulip-cloud-or-a-self-hosted-server
	//
	GetRealmExports(ctx context.Context) GetRealmExportsRequest

	// GetRealmExportsExecute executes the request
	GetRealmExportsExecute(r GetRealmExportsRequest) (*GetRealmExportsResponse, *http.Response, error)

	// GetServerSettings Get server settings
	//
	// Fetch global settings for a Zulip server.
	//
	// *Note:** this endpoint does not require any authentication at all, and you can use it to check:
	//
	//   - If this is a Zulip server, and if so, what version of Zulip it's running.
	//   - What a Zulip client (e.g. a mobile app or
	// [zulip-terminal]) needs to
	// know in order to display a login prompt for the server (e.g. what
	// authentication methods are available).
	//
	// [zulip-terminal]: https://github.com/zulip/zulip-terminal/
	GetServerSettings(ctx context.Context) GetServerSettingsRequest

	// GetServerSettingsExecute executes the request
	GetServerSettingsExecute(r GetServerSettingsRequest) (*GetServerSettingsResponse, *http.Response, error)

	// RemoveCodePlayground Remove a code playground
	//
	// Remove a [code playground] previously
	// configured for an organization.
	//
	// *Changes**: New in Zulip 4.0 (feature level 49).
	//
	// [code playground]: https://zulip.com/help/code-blocks#code-playgrounds
	RemoveCodePlayground(ctx context.Context, playgroundId int64) RemoveCodePlaygroundRequest

	// RemoveCodePlaygroundExecute executes the request
	RemoveCodePlaygroundExecute(r RemoveCodePlaygroundRequest) (*Response, *http.Response, error)

	// RemoveLinkifier Remove a linkifier
	//
	// Remove [linkifiers], regular
	// expression patterns that are automatically linkified when they appear
	// in messages and topics.
	//
	// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
	RemoveLinkifier(ctx context.Context, filterId int64) RemoveLinkifierRequest

	// RemoveLinkifierExecute executes the request
	RemoveLinkifierExecute(r RemoveLinkifierRequest) (*Response, *http.Response, error)

	// ReorderCustomProfileFields Reorder custom profile fields
	//
	// Reorder the custom profile fields in the user's organization.
	//
	// Custom profile fields are displayed in Zulip UI widgets in order; this
	// endpoint allows administrative settings UI to change the field ordering.
	//
	// This endpoint is used to implement the dragging feature described in the
	// [custom profile fields documentation].
	//
	// [custom profile fields documentation]: https://zulip.com/help/custom-profile-fields
	ReorderCustomProfileFields(ctx context.Context) ReorderCustomProfileFieldsRequest

	// ReorderCustomProfileFieldsExecute executes the request
	ReorderCustomProfileFieldsExecute(r ReorderCustomProfileFieldsRequest) (*Response, *http.Response, error)

	// ReorderLinkifiers Reorder linkifiers
	//
	// Change the order that the regular expression patterns in the organization's
	// [linkifiers] are matched in messages and topics.
	// Useful when defining linkifiers with overlapping patterns.
	//
	// *Changes**: New in Zulip 8.0 (feature level 202). Before this feature level,
	// linkifiers were always processed in order by Id, which meant users would
	// need to delete and recreate them to reorder the list of linkifiers.
	//
	// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
	ReorderLinkifiers(ctx context.Context) ReorderLinkifiersRequest

	// ReorderLinkifiersExecute executes the request
	ReorderLinkifiersExecute(r ReorderLinkifiersRequest) (*Response, *http.Response, error)

	// TestWelcomeBotCustomMessage Test welcome bot custom message
	//
	// Sends a test Welcome Bot custom message to the acting administrator.
	// This allows administrators to preview how the custom welcome message will
	// appear when received by new users upon joining the organization.
	//
	// *Changes**: New in Zulip 11.0 (feature level 416).
	//
	TestWelcomeBotCustomMessage(ctx context.Context) TestWelcomeBotCustomMessageRequest

	// TestWelcomeBotCustomMessageExecute executes the request
	TestWelcomeBotCustomMessageExecute(r TestWelcomeBotCustomMessageRequest) (*TestWelcomeBotCustomMessageResponse, *http.Response, error)

	// UpdateLinkifier Update a linkifier
	//
	// Update a [linkifier], regular
	// expression patterns that are automatically linkified when they appear
	// in messages and topics.
	//
	// *Changes**: New in Zulip 4.0 (feature level 57).
	//
	// [linkifier]: https://zulip.com/help/add-a-custom-linkifier
	UpdateLinkifier(ctx context.Context, filterId int64) UpdateLinkifierRequest

	// UpdateLinkifierExecute executes the request
	UpdateLinkifierExecute(r UpdateLinkifierRequest) (*Response, *http.Response, error)

	// UpdateRealmUserSettingsDefaults Update realm-level defaults of user settings
	//
	// Change the [default values of settings] for new users
	// joining the organization. Essentially all
	// [personal preference settings] are supported.
	//
	// This feature can be invaluable for customizing Zulip's default
	// settings for notifications or UI to be appropriate for how the
	// organization is using Zulip. (Note that this only supports
	// personal preference settings, like when to send push
	// notifications or what emoji set to use, not profile or
	// identity settings that naturally should be different for each user).
	//
	// Note that this endpoint cannot, at present, be used to modify
	// settings for existing users in any way.
	//
	// *Changes**: Removed `dense_mode` setting in Zulip 10.0 (feature level 364)
	// as we now have `web_font_size_px` and `web_line_height_percent`
	// settings for more control.
	//
	// New in Zulip 5.0 (feature level 96). If any parameters sent in the
	// request are not supported by this endpoint, an
	// [`ignored_parameters_unsupported`] array will
	// be returned in the JSON success response.
	//
	// [default values of settings]: https://zulip.com/help/configure-default-new-user-settings
	// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
	//
	// [personal preference settings]: https://zulip.com/api/update-settings
	UpdateRealmUserSettingsDefaults(ctx context.Context) UpdateRealmUserSettingsDefaultsRequest

	// UpdateRealmUserSettingsDefaultsExecute executes the request
	UpdateRealmUserSettingsDefaultsExecute(r UpdateRealmUserSettingsDefaultsRequest) (*Response, *http.Response, error)

	// UploadCustomEmoji Upload custom emoji
	//
	// This endpoint is used to upload a custom emoji for use in the user's
	// organization. Access to this endpoint depends on the
	// [organization's configuration].
	//
	// [organization's configuration]: https://zulip.com/help/custom-emoji#change-who-can-add-custom-emoji
	UploadCustomEmoji(ctx context.Context, emojiName string) UploadCustomEmojiRequest

	// UploadCustomEmojiExecute executes the request
	UploadCustomEmojiExecute(r UploadCustomEmojiRequest) (*Response, *http.Response, error)
}

type AddCodePlaygroundRequest struct {
	ctx              context.Context
	ApiService       ServerAndOrganizationsAPI
	name             *string
	pygmentsLanguage *string
	urlTemplate      *string
}

// The user-visible display name of the playground which can be used to pick the target playground, especially when multiple playground options exist for that programming language.
func (r AddCodePlaygroundRequest) Name(name string) AddCodePlaygroundRequest {
	r.name = &name
	return r
}

// The name of the Pygments language lexer for that programming language.
func (r AddCodePlaygroundRequest) PygmentsLanguage(pygmentsLanguage string) AddCodePlaygroundRequest {
	r.pygmentsLanguage = &pygmentsLanguage
	return r
}

// The [RFC 6570] compliant URL template for the playground. The template should contain exactly one variable named `code`, which determines how the extracted code should be substituted in the playground URL.
//
//	**Changes**: New in Zulip 8.0 (feature level 196). This replaced the `url_prefix` parameter, which was used to construct URLs by just concatenating `url_prefix` and `code`.
//
// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
func (r AddCodePlaygroundRequest) UrlTemplate(urlTemplate string) AddCodePlaygroundRequest {
	r.urlTemplate = &urlTemplate
	return r
}

func (r AddCodePlaygroundRequest) Execute() (*AddCodePlaygroundResponse, *http.Response, error) {
	return r.ApiService.AddCodePlaygroundExecute(r)
}

// AddCodePlayground Add a code playground
//
// Configure [code playgrounds] for the organization.
//
// *Changes**: New in Zulip 4.0 (feature level 49). A parameter encoding bug was
// fixed in Zulip 4.0 (feature level 57).
//
// [code playgrounds]: https://zulip.com/help/code-blocks#code-playgrounds
func (c *simpleClient) AddCodePlayground(ctx context.Context) AddCodePlaygroundRequest {
	return AddCodePlaygroundRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) AddCodePlaygroundExecute(r AddCodePlaygroundRequest) (*AddCodePlaygroundResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddCodePlaygroundResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/playgrounds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.pygmentsLanguage == nil {
		return localVarReturnValue, nil, reportError("pygmentsLanguage is required and must be specified")
	}
	if r.urlTemplate == nil {
		return localVarReturnValue, nil, reportError("urlTemplate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "pygments_language", r.pygmentsLanguage, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "url_template", r.urlTemplate, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type AddLinkifierRequest struct {
	ctx         context.Context
	ApiService  ServerAndOrganizationsAPI
	pattern     *string
	urlTemplate *string
}

// The [Python regular expression] that should trigger the linkifier.
//
// [Python regular expression]: https://docs.python.org/3/howto/regex.html
func (r AddLinkifierRequest) Pattern(pattern string) AddLinkifierRequest {
	r.pattern = &pattern
	return r
}

// The [RFC 6570] compliant URL template used for the link. If you used named groups in `pattern`, you can insert their content here with `{name_of_group}`.
//
//	**Changes**: New in Zulip 7.0 (feature level 176). This replaced the `url_format_string` parameter, which was a format string in which named groups&#39; content could be inserted with `%(name_of_group)s`.
//
// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
func (r AddLinkifierRequest) UrlTemplate(urlTemplate string) AddLinkifierRequest {
	r.urlTemplate = &urlTemplate
	return r
}

func (r AddLinkifierRequest) Execute() (*AddLinkifierResponse, *http.Response, error) {
	return r.ApiService.AddLinkifierExecute(r)
}

// AddLinkifier Add a linkifier
//
// Configure [linkifiers],
// regular expression patterns that are automatically linkified when they
// appear in messages and topics.
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
func (c *simpleClient) AddLinkifier(ctx context.Context) AddLinkifierRequest {
	return AddLinkifierRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) AddLinkifierExecute(r AddLinkifierRequest) (*AddLinkifierResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AddLinkifierResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pattern == nil {
		return localVarReturnValue, nil, reportError("pattern is required and must be specified")
	}
	if r.urlTemplate == nil {
		return localVarReturnValue, nil, reportError("urlTemplate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pattern", r.pattern, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "url_template", r.urlTemplate, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateCustomProfileFieldRequest struct {
	ctx                     context.Context
	ApiService              ServerAndOrganizationsAPI
	fieldType               *int32
	name                    *string
	hint                    *string
	fieldData               *map[string]interface{}
	displayInProfileSummary *bool
	required                *bool
	editableByUser          *bool
}

// The field type can be any of the supported custom profile field types. See the [custom profile fields documentation] for more details on what each type means.  - **1**: Short text - **2**: Long text - **3**: List of options - **4**: Date picker - **5**: Link - **6**: Person picker - **7**: External account - **8**: Pronouns  **Changes**: Field type `8` added in Zulip 6.0 (feature level 151).
//
// [custom profile fields documentation]: https://zulip.com/help/custom-profile-fields
func (r CreateCustomProfileFieldRequest) FieldType(fieldType int32) CreateCustomProfileFieldRequest {
	r.fieldType = &fieldType
	return r
}

// The name of the custom profile field, which will appear both in user-facing settings UI for configuring custom profile fields and in UI displaying a user&#39;s profile.
func (r CreateCustomProfileFieldRequest) Name(name string) CreateCustomProfileFieldRequest {
	r.name = &name
	return r
}

// The help text to be displayed for the custom profile field in user-facing settings UI for configuring custom profile fields.
func (r CreateCustomProfileFieldRequest) Hint(hint string) CreateCustomProfileFieldRequest {
	r.hint = &hint
	return r
}

// Field types 3 (List of options) and 7 (External account) support storing additional configuration for the field type in the `field_data` attribute.  For field type 3 (List of options), this attribute is a JSON dictionary defining the choices and the order they will be displayed in the dropdown UI for individual users to select an option.  The interface for field type 7 is not yet stabilized.
func (r CreateCustomProfileFieldRequest) FieldData(fieldData map[string]interface{}) CreateCustomProfileFieldRequest {
	r.fieldData = &fieldData
	return r
}

// Whether clients should display this profile field in a summary section of a user&#39;s profile (or in a more easily accessible \\"small profile\\").  At most 2 profile fields may have this property be true in a given organization. The \\"Long text\\" [profile field types] profile field types cannot be selected to be displayed in profile summaries.  The \\"Person picker\\" profile field is also not supported, but that is likely to be temporary.  [profile field types]: https://zulip.com/help/custom-profile-fields#profile-field-types  **Changes**: New in Zulip 6.0 (feature level 146).
func (r CreateCustomProfileFieldRequest) DisplayInProfileSummary(displayInProfileSummary bool) CreateCustomProfileFieldRequest {
	r.displayInProfileSummary = &displayInProfileSummary
	return r
}

// Whether an organization administrator has configured this profile field as required.  Because the required property is mutable, clients cannot assume that a required custom profile field has a value. The Zulip web application displays a prominent banner to any user who has not set a value for a required field.
//
//	**Changes**: New in Zulip 9.0 (feature level 244).
func (r CreateCustomProfileFieldRequest) Required(required bool) CreateCustomProfileFieldRequest {
	r.required = &required
	return r
}

// Whether regular users can edit this profile field on their own account.  Note that organization administrators can edit custom profile fields for any user regardless of this setting.
//
//	**Changes**: New in Zulip 10.0 (feature level 296).
func (r CreateCustomProfileFieldRequest) EditableByUser(editableByUser bool) CreateCustomProfileFieldRequest {
	r.editableByUser = &editableByUser
	return r
}

func (r CreateCustomProfileFieldRequest) Execute() (*CreateCustomProfileFieldResponse, *http.Response, error) {
	return r.ApiService.CreateCustomProfileFieldExecute(r)
}

// CreateCustomProfileField Create a custom profile field
//
// [Create a custom profile field] in the user's organization.
//
// [Create a custom profile field]: https://zulip.com/help/custom-profile-fields#add-a-custom-profile-field
func (c *simpleClient) CreateCustomProfileField(ctx context.Context) CreateCustomProfileFieldRequest {
	return CreateCustomProfileFieldRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) CreateCustomProfileFieldExecute(r CreateCustomProfileFieldRequest) (*CreateCustomProfileFieldResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateCustomProfileFieldResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/profile_fields"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fieldType == nil {
		return localVarReturnValue, nil, reportError("fieldType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.hint != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hint", r.hint, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "field_type", r.fieldType, "form", "")
	if r.fieldData != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "field_data", r.fieldData, "form", "")
	}
	if r.displayInProfileSummary != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "display_in_profile_summary", r.displayInProfileSummary, "form", "")
	}
	if r.required != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "required", r.required, "form", "")
	}
	if r.editableByUser != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "editable_by_user", r.editableByUser, "form", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeactivateCustomEmojiRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
	emojiName  string
}

func (r DeactivateCustomEmojiRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeactivateCustomEmojiExecute(r)
}

// DeactivateCustomEmoji Deactivate custom emoji
//
// [Deactivate a custom emoji] from
// the user's organization.
//
// Users can only deactivate custom emoji that they added themselves except for
// organization administrators, who can deactivate any custom emoji.
//
// Note that deactivated emoji will still be visible in old messages, reactions,
// user statuses and channel descriptions.
//
// *Changes**: Before Zulip 8.0 (feature level 190), this endpoint returned an
// HTTP status code of 400 when the emoji did not exist, instead of 404.
//
// [Deactivate a custom emoji]: https://zulip.com/help/custom-emoji#deactivate-custom-emoji
func (c *simpleClient) DeactivateCustomEmoji(ctx context.Context, emojiName string) DeactivateCustomEmojiRequest {
	return DeactivateCustomEmojiRequest{
		ApiService: c,
		ctx:        ctx,
		emojiName:  emojiName,
	}
}

// Execute executes the request
func (c *simpleClient) DeactivateCustomEmojiExecute(r DeactivateCustomEmojiRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/emoji/{emoji_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"emoji_name"+"}", url.PathEscape(parameterValueToString(r.emojiName, "emojiName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ExportRealmRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
	exportType *int32
}

// Whether to create a public or a standard data export.  - 1 &#x3D; Public data export. - 2 &#x3D; Standard data export.  If not specified, defaults to 1.
//
//	**Changes**: New in Zulip 10.0 (feature level 304). Previously, all export requests were public data exports.
func (r ExportRealmRequest) ExportType(exportType int32) ExportRealmRequest {
	r.exportType = &exportType
	return r
}

func (r ExportRealmRequest) Execute() (*ExportRealmResponse, *http.Response, error) {
	return r.ApiService.ExportRealmExecute(r)
}

// ExportRealm Create a data export
//
// Create a public or a standard [data export] of the organization.
//
// !!! warn ""
//
// *Note**: If you're the administrator of a self-hosted installation,
// you may be looking for the documentation on [server data export and import] or [server backups].
//
// *Changes**: Prior to Zulip 10.0 (feature level 304), only
// public data exports could be created using this endpoint.
//
// New in Zulip 2.1.
//
// [server data export and import]: https://zulip.readthedocs.io/en/stable/production/export-and-import.html#data-export
// [server backups]: https://zulip.readthedocs.io/en/stable/production/export-and-import.html#backups
// [data export]: https://zulip.com/help/export-your-organization#export-for-migrating-to-zulip-cloud-or-a-self-hosted-server
func (c *simpleClient) ExportRealm(ctx context.Context) ExportRealmRequest {
	return ExportRealmRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) ExportRealmExecute(r ExportRealmRequest) (*ExportRealmResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportRealmResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/export/realm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.exportType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "export_type", r.exportType, "", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCustomEmojiRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetCustomEmojiRequest) Execute() (*GetCustomEmojiResponse, *http.Response, error) {
	return r.ApiService.GetCustomEmojiExecute(r)
}

// GetCustomEmoji Get all custom emoji
//
// Get all the custom emoji in the user's organization.
func (c *simpleClient) GetCustomEmoji(ctx context.Context) GetCustomEmojiRequest {
	return GetCustomEmojiRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetCustomEmojiExecute(r GetCustomEmojiRequest) (*GetCustomEmojiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCustomEmojiResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/emoji"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCustomProfileFieldsRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetCustomProfileFieldsRequest) Execute() (*GetCustomProfileFieldsResponse, *http.Response, error) {
	return r.ApiService.GetCustomProfileFieldsExecute(r)
}

// GetCustomProfileFields Get all custom profile fields
//
// Get all the [custom profile fields]
// configured for the user's organization.
//
// [custom profile fields]: https://zulip.com/help/custom-profile-fields
func (c *simpleClient) GetCustomProfileFields(ctx context.Context) GetCustomProfileFieldsRequest {
	return GetCustomProfileFieldsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetCustomProfileFieldsExecute(r GetCustomProfileFieldsRequest) (*GetCustomProfileFieldsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetCustomProfileFieldsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/profile_fields"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetLinkifiersRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetLinkifiersRequest) Execute() (*GetLinkifiersResponse, *http.Response, error) {
	return r.ApiService.GetLinkifiersExecute(r)
}

// GetLinkifiers Get linkifiers
//
// List all of an organization's configured
// [linkifiers], regular
// expression patterns that are automatically linkified when they appear
// in messages and topics.
//
// *Changes**: New in Zulip 4.0 (feature level 54). On older versions,
// a similar `GET /realm/filters` endpoint was available with each entry in
// a `[pattern, url_format, id]` tuple format.
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
func (c *simpleClient) GetLinkifiers(ctx context.Context) GetLinkifiersRequest {
	return GetLinkifiersRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetLinkifiersExecute(r GetLinkifiersRequest) (*GetLinkifiersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLinkifiersResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/linkifiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetPresenceRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetPresenceRequest) Execute() (*GetPresenceResponse, *http.Response, error) {
	return r.ApiService.GetPresenceExecute(r)
}

// GetPresence Get presence of all users
//
// Get the presence information of all the users in an organization.
//
// If the `CAN_ACCESS_ALL_USERS_GROUP_LIMITS_PRESENCE` server-level
// setting is set to `true`, presence information of only accessible
// users are returned.
//
// Complete Zulip apps are recommended to fetch presence
// information when they post their own state using the [`POST
// /presence`](https://zulip.com/api/update-presence) API endpoint.
func (c *simpleClient) GetPresence(ctx context.Context) GetPresenceRequest {
	return GetPresenceRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetPresenceExecute(r GetPresenceRequest) (*GetPresenceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPresenceResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/presence"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetRealmExportConsentsRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetRealmExportConsentsRequest) Execute() (*GetRealmExportConsentsResponse, *http.Response, error) {
	return r.ApiService.GetRealmExportConsentsExecute(r)
}

// GetRealmExportConsents Get data export consent state
//
// Fetches which users have [consented]
// for their private data to be exported by organization administrators.
//
// *Changes**: New in Zulip 10.0 (feature level 295).
//
// [consented]: https://zulip.com/help/export-your-organization#configure-whether-administrators-can-export-your-private-data
func (c *simpleClient) GetRealmExportConsents(ctx context.Context) GetRealmExportConsentsRequest {
	return GetRealmExportConsentsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetRealmExportConsentsExecute(r GetRealmExportConsentsRequest) (*GetRealmExportConsentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRealmExportConsentsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/export/realm/consents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetRealmExportsRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetRealmExportsRequest) Execute() (*GetRealmExportsResponse, *http.Response, error) {
	return r.ApiService.GetRealmExportsExecute(r)
}

// GetRealmExports Get all data exports
//
// Fetch all the public and standard [data exports]
// of the organization.
//
// *Changes**: Prior to Zulip 10.0 (feature level 304), only
// public data exports could be fetched using this endpoint.
//
// New in Zulip 2.1.
//
// [data exports]: https://zulip.com/help/export-your-organization#export-for-migrating-to-zulip-cloud-or-a-self-hosted-server
func (c *simpleClient) GetRealmExports(ctx context.Context) GetRealmExportsRequest {
	return GetRealmExportsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetRealmExportsExecute(r GetRealmExportsRequest) (*GetRealmExportsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRealmExportsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/export/realm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetServerSettingsRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
}

func (r GetServerSettingsRequest) Execute() (*GetServerSettingsResponse, *http.Response, error) {
	return r.ApiService.GetServerSettingsExecute(r)
}

// GetServerSettings Get server settings
//
// Fetch global settings for a Zulip server.
//
// *Note:** this endpoint does not require any authentication at all, and you can use it to check:
//
//   - If this is a Zulip server, and if so, what version of Zulip it's running.
//
//   - What a Zulip client (e.g. a mobile app or
//
// [zulip-terminal]) needs to
// know in order to display a login prompt for the server (e.g. what
// authentication methods are available).
//
// [zulip-terminal]: https://github.com/zulip/zulip-terminal/
func (c *simpleClient) GetServerSettings(ctx context.Context) GetServerSettingsRequest {
	return GetServerSettingsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) GetServerSettingsExecute(r GetServerSettingsRequest) (*GetServerSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetServerSettingsResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server_settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveCodePlaygroundRequest struct {
	ctx          context.Context
	ApiService   ServerAndOrganizationsAPI
	playgroundId int64
}

func (r RemoveCodePlaygroundRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveCodePlaygroundExecute(r)
}

// RemoveCodePlayground Remove a code playground
//
// Remove a [code playground] previously
// configured for an organization.
//
// *Changes**: New in Zulip 4.0 (feature level 49).
//
// [code playground]: https://zulip.com/help/code-blocks#code-playgrounds
func (c *simpleClient) RemoveCodePlayground(ctx context.Context, playgroundId int64) RemoveCodePlaygroundRequest {
	return RemoveCodePlaygroundRequest{
		ApiService:   c,
		ctx:          ctx,
		playgroundId: playgroundId,
	}
}

// Execute executes the request
func (c *simpleClient) RemoveCodePlaygroundExecute(r RemoveCodePlaygroundRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/playgrounds/{playground_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"playground_id"+"}", url.PathEscape(parameterValueToString(r.playgroundId, "playgroundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RemoveLinkifierRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
	filterId   int64
}

func (r RemoveLinkifierRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveLinkifierExecute(r)
}

// RemoveLinkifier Remove a linkifier
//
// Remove [linkifiers], regular
// expression patterns that are automatically linkified when they appear
// in messages and topics.
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
func (c *simpleClient) RemoveLinkifier(ctx context.Context, filterId int64) RemoveLinkifierRequest {
	return RemoveLinkifierRequest{
		ApiService: c,
		ctx:        ctx,
		filterId:   filterId,
	}
}

// Execute executes the request
func (c *simpleClient) RemoveLinkifierExecute(r RemoveLinkifierRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/filters/{filter_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"filter_id"+"}", url.PathEscape(parameterValueToString(r.filterId, "filterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ReorderCustomProfileFieldsRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
	order      *[]int64
}

// A list of the Ids of all the custom profile fields defined in this organization, in the desired new order.
func (r ReorderCustomProfileFieldsRequest) Order(order []int64) ReorderCustomProfileFieldsRequest {
	r.order = &order
	return r
}

func (r ReorderCustomProfileFieldsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.ReorderCustomProfileFieldsExecute(r)
}

// ReorderCustomProfileFields Reorder custom profile fields
//
// Reorder the custom profile fields in the user's organization.
//
// Custom profile fields are displayed in Zulip UI widgets in order; this
// endpoint allows administrative settings UI to change the field ordering.
//
// This endpoint is used to implement the dragging feature described in the
// [custom profile fields documentation].
//
// [custom profile fields documentation]: https://zulip.com/help/custom-profile-fields
func (c *simpleClient) ReorderCustomProfileFields(ctx context.Context) ReorderCustomProfileFieldsRequest {
	return ReorderCustomProfileFieldsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) ReorderCustomProfileFieldsExecute(r ReorderCustomProfileFieldsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/profile_fields"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.order == nil {
		return localVarReturnValue, nil, reportError("order is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "order", r.order, "form", "multi")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ReorderLinkifiersRequest struct {
	ctx                 context.Context
	ApiService          ServerAndOrganizationsAPI
	orderedLinkifierIds *[]int64
}

// A list of the Ids of all the linkifiers defined in this organization, in the desired new order.
func (r ReorderLinkifiersRequest) OrderedLinkifierIds(orderedLinkifierIds []int64) ReorderLinkifiersRequest {
	r.orderedLinkifierIds = &orderedLinkifierIds
	return r
}

func (r ReorderLinkifiersRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.ReorderLinkifiersExecute(r)
}

// ReorderLinkifiers Reorder linkifiers
//
// Change the order that the regular expression patterns in the organization's
// [linkifiers] are matched in messages and topics.
// Useful when defining linkifiers with overlapping patterns.
//
// *Changes**: New in Zulip 8.0 (feature level 202). Before this feature level,
// linkifiers were always processed in order by Id, which meant users would
// need to delete and recreate them to reorder the list of linkifiers.
//
// [linkifiers]: https://zulip.com/help/add-a-custom-linkifier
func (c *simpleClient) ReorderLinkifiers(ctx context.Context) ReorderLinkifiersRequest {
	return ReorderLinkifiersRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) ReorderLinkifiersExecute(r ReorderLinkifiersRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/linkifiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderedLinkifierIds == nil {
		return localVarReturnValue, nil, reportError("orderedLinkifierIds is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "ordered_linkifier_ids", r.orderedLinkifierIds, "form", "multi")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TestWelcomeBotCustomMessageRequest struct {
	ctx                      context.Context
	ApiService               ServerAndOrganizationsAPI
	welcomeMessageCustomText *string
}

// Custom message text, in Zulip Markdown format, to be used for this test message.  Maximum length is 8000 characters.
func (r TestWelcomeBotCustomMessageRequest) WelcomeMessageCustomText(welcomeMessageCustomText string) TestWelcomeBotCustomMessageRequest {
	r.welcomeMessageCustomText = &welcomeMessageCustomText
	return r
}

func (r TestWelcomeBotCustomMessageRequest) Execute() (*TestWelcomeBotCustomMessageResponse, *http.Response, error) {
	return r.ApiService.TestWelcomeBotCustomMessageExecute(r)
}

// TestWelcomeBotCustomMessage Test welcome bot custom message
//
// Sends a test Welcome Bot custom message to the acting administrator.
// This allows administrators to preview how the custom welcome message will
// appear when received by new users upon joining the organization.
//
// *Changes**: New in Zulip 11.0 (feature level 416).
func (c *simpleClient) TestWelcomeBotCustomMessage(ctx context.Context) TestWelcomeBotCustomMessageRequest {
	return TestWelcomeBotCustomMessageRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) TestWelcomeBotCustomMessageExecute(r TestWelcomeBotCustomMessageRequest) (*TestWelcomeBotCustomMessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TestWelcomeBotCustomMessageResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/test_welcome_bot_custom_message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.welcomeMessageCustomText == nil {
		return localVarReturnValue, nil, reportError("welcomeMessageCustomText is required and must be specified")
	}
	if strlen(*r.welcomeMessageCustomText) > 8000 {
		return localVarReturnValue, nil, reportError("welcomeMessageCustomText must have less than 8000 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "welcome_message_custom_text", r.welcomeMessageCustomText, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateLinkifierRequest struct {
	ctx         context.Context
	ApiService  ServerAndOrganizationsAPI
	filterId    int64
	pattern     *string
	urlTemplate *string
}

// The [Python regular expression] that should trigger the linkifier.
//
// [Python regular expression]: https://docs.python.org/3/howto/regex.html
func (r UpdateLinkifierRequest) Pattern(pattern string) UpdateLinkifierRequest {
	r.pattern = &pattern
	return r
}

// The [RFC 6570] compliant URL template used for the link. If you used named groups in `pattern`, you can insert their content here with `{name_of_group}`.
//
//	**Changes**: New in Zulip 7.0 (feature level 176). This replaced the `url_format_string` parameter, which was a format string in which named groups&#39; content could be inserted with `%(name_of_group)s`.
//
// [RFC 6570]: https://www.rfc-editor.org/rfc/rfc6570.html
func (r UpdateLinkifierRequest) UrlTemplate(urlTemplate string) UpdateLinkifierRequest {
	r.urlTemplate = &urlTemplate
	return r
}

func (r UpdateLinkifierRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateLinkifierExecute(r)
}

// UpdateLinkifier Update a linkifier
//
// Update a [linkifier], regular
// expression patterns that are automatically linkified when they appear
// in messages and topics.
//
// *Changes**: New in Zulip 4.0 (feature level 57).
//
// [linkifier]: https://zulip.com/help/add-a-custom-linkifier
func (c *simpleClient) UpdateLinkifier(ctx context.Context, filterId int64) UpdateLinkifierRequest {
	return UpdateLinkifierRequest{
		ApiService: c,
		ctx:        ctx,
		filterId:   filterId,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateLinkifierExecute(r UpdateLinkifierRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/filters/{filter_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"filter_id"+"}", url.PathEscape(parameterValueToString(r.filterId, "filterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pattern == nil {
		return localVarReturnValue, nil, reportError("pattern is required and must be specified")
	}
	if r.urlTemplate == nil {
		return localVarReturnValue, nil, reportError("urlTemplate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pattern", r.pattern, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "url_template", r.urlTemplate, "", "")
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateRealmUserSettingsDefaultsRequest struct {
	ctx                                            context.Context
	ApiService                                     ServerAndOrganizationsAPI
	starredMessageCounts                           *bool
	receivesTypingNotifications                    *bool
	webSuggestUpdateTimezone                       *bool
	fluidLayoutWidth                               *bool
	highContrastMode                               *bool
	webMarkReadOnScrollPolicy                      *int32
	webChannelDefaultView                          *int32
	webFontSizePx                                  *int32
	webLineHeightPercent                           *int32
	colorScheme                                    *ColorScheme
	enableDraftsSynchronization                    *bool
	translateEmoticons                             *bool
	displayEmojiReactionUsers                      *bool
	webHomeView                                    *string
	webEscapeNavigatesToHomeView                   *bool
	leftSideUserlist                               *bool
	emojiset                                       *string
	demoteInactiveChannels                         *DemoteInactiveChannels
	userListStyle                                  *int32
	webAnimateImagePreviews                        *string
	webChannelUnreadsCountDisplayPolicy            *int32
	hideAiFeatures                                 *bool
	webLeftSidebarShowChannelFolders               *bool
	webLeftSidebarUnreadsCountSummary              *bool
	enableChannelDesktopNotifications              *bool
	enableChannelEmailNotifications                *bool
	enableChannelPushNotifications                 *bool
	enableChannelAudibleNotifications              *bool
	notificationSound                              *string
	enableDesktopNotifications                     *bool
	enableSounds                                   *bool
	enableFollowedTopicDesktopNotifications        *bool
	enableFollowedTopicEmailNotifications          *bool
	enableFollowedTopicPushNotifications           *bool
	enableFollowedTopicAudibleNotifications        *bool
	emailNotificationsBatchingPeriodSeconds        *int32
	enableOfflineEmailNotifications                *bool
	enableOfflinePushNotifications                 *bool
	enableOnlinePushNotifications                  *bool
	enableDigestEmails                             *bool
	messageContentInEmailNotifications             *bool
	pmContentInDesktopNotifications                *bool
	wildcardMentionsNotify                         *bool
	enableFollowedTopicWildcardMentionsNotify      *bool
	desktopIconCountDisplay                        *int32
	realmNameInEmailNotificationsPolicy            *int32
	automaticallyFollowTopicsPolicy                *int32
	automaticallyUnmuteTopicsInMutedChannelsPolicy *int32
	automaticallyFollowTopicsWhereMentioned        *bool
	resolvedTopicNoticeAutoReadPolicy              *string
	presenceEnabled                                *bool
	enterSends                                     *bool
	twentyFourHourTime                             *bool
	sendPrivateTypingNotifications                 *bool
	sendChannelTypingNotifications                 *bool
	sendReadReceipts                               *bool
	emailAddressVisibility                         *int32
	webNavigateToSentMessage                       *bool
}

// Whether clients should display the [number of starred messages].
//
// [number of starred messages]: https://zulip.com/help/star-a-message#display-the-number-of-starred-messages
func (r UpdateRealmUserSettingsDefaultsRequest) StarredMessageCounts(starredMessageCounts bool) UpdateRealmUserSettingsDefaultsRequest {
	r.starredMessageCounts = &starredMessageCounts
	return r
}

// Whether the user is configured to receive typing notifications from other users. The server will only deliver typing notifications events to users who for whom this is enabled.
//
//	**Changes**: New in Zulip 9.0 (feature level 253). Previously, there were only options to disable sending typing notifications.
func (r UpdateRealmUserSettingsDefaultsRequest) ReceivesTypingNotifications(receivesTypingNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.receivesTypingNotifications = &receivesTypingNotifications
	return r
}

// Whether the user should be shown an alert, offering to update their [profile time zone], when the time displayed for the profile time zone differs from the current time displayed by the time zone configured on their device.
//
//	**Changes**: New in Zulip 10.0 (feature level 329).
//
// [profile time zone]: https://zulip.com/help/change-your-timezone
func (r UpdateRealmUserSettingsDefaultsRequest) WebSuggestUpdateTimezone(webSuggestUpdateTimezone bool) UpdateRealmUserSettingsDefaultsRequest {
	r.webSuggestUpdateTimezone = &webSuggestUpdateTimezone
	return r
}

// Whether to use the [maximum available screen width] for the web app&#39;s center panel (message feed, recent conversations) on wide screens.
//
// [maximum available screen width]: https://zulip.com/help/enable-full-width-display
func (r UpdateRealmUserSettingsDefaultsRequest) FluidLayoutWidth(fluidLayoutWidth bool) UpdateRealmUserSettingsDefaultsRequest {
	r.fluidLayoutWidth = &fluidLayoutWidth
	return r
}

// This setting is reserved for use to control variations in Zulip&#39;s design to help visually impaired users.
func (r UpdateRealmUserSettingsDefaultsRequest) HighContrastMode(highContrastMode bool) UpdateRealmUserSettingsDefaultsRequest {
	r.highContrastMode = &highContrastMode
	return r
}

// Whether or not to mark messages as read when the user scrolls through their feed.  - 1 - Always - 2 - Only in conversation views - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 175). Previously, there was no way for the user to configure this behavior on the web, and the Zulip web and desktop apps behaved like the \\"Always\\" setting when marking messages as read.
func (r UpdateRealmUserSettingsDefaultsRequest) WebMarkReadOnScrollPolicy(webMarkReadOnScrollPolicy int32) UpdateRealmUserSettingsDefaultsRequest {
	r.webMarkReadOnScrollPolicy = &webMarkReadOnScrollPolicy
	return r
}

// Web/desktop app setting controlling the default navigation behavior when clicking on a channel link.  - 1 - Top topic in the channel - 2 - Channel feed - 3 - List of topics - 4 - Top unread topic in channel  **Changes**: The \\"Top unread topic in channel\\" is new in Zulip 11.0 (feature level 401).  The \\"List of topics\\" option is new in Zulip 11.0 (feature level 383).  New in Zulip 9.0 (feature level 269). Previously, this was not configurable, and every user had the \\"Channel feed\\" behavior.
func (r UpdateRealmUserSettingsDefaultsRequest) WebChannelDefaultView(webChannelDefaultView int32) UpdateRealmUserSettingsDefaultsRequest {
	r.webChannelDefaultView = &webChannelDefaultView
	return r
}

// User-configured primary `font-size` for the web application, in pixels.
//
//	**Changes**: New in Zulip 9.0 (feature level 245). Previously, font size was only adjustable via browser zoom. Note that this setting was not fully implemented at this feature level.
func (r UpdateRealmUserSettingsDefaultsRequest) WebFontSizePx(webFontSizePx int32) UpdateRealmUserSettingsDefaultsRequest {
	r.webFontSizePx = &webFontSizePx
	return r
}

// User-configured primary `line-height` for the web application, in percent, so a value of 120 represents a `line-height` of 1.2.
//
//	**Changes**: New in Zulip 9.0 (feature level 245). Previously, line height was not user-configurable. Note that this setting was not fully implemented at this feature level.
func (r UpdateRealmUserSettingsDefaultsRequest) WebLineHeightPercent(webLineHeightPercent int32) UpdateRealmUserSettingsDefaultsRequest {
	r.webLineHeightPercent = &webLineHeightPercent
	return r
}

// Controls which [color theme] to use.  - 1 - Automatic - 2 - Dark theme - 3 - Light theme  Automatic detection is implementing using the standard `prefers-color-scheme` media query.
//
// [color theme]: https://zulip.com/help/dark-theme
func (r UpdateRealmUserSettingsDefaultsRequest) ColorScheme(colorScheme ColorScheme) UpdateRealmUserSettingsDefaultsRequest {
	r.colorScheme = &colorScheme
	return r
}

// A boolean parameter to control whether synchronizing drafts is enabled for the user. When synchronization is disabled, all drafts stored in the server will be automatically deleted from the server.  This does not do anything (like sending events) to delete local copies of drafts stored in clients.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableDraftsSynchronization(enableDraftsSynchronization bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableDraftsSynchronization = &enableDraftsSynchronization
	return r
}

// Whether to [translate emoticons to emoji] in messages the user sends.
//
// [translate emoticons to emoji]: https://zulip.com/help/configure-emoticon-translations
func (r UpdateRealmUserSettingsDefaultsRequest) TranslateEmoticons(translateEmoticons bool) UpdateRealmUserSettingsDefaultsRequest {
	r.translateEmoticons = &translateEmoticons
	return r
}

// Whether to display the names of reacting users on a message.  When enabled, clients should display the names of reacting users, rather than a count, for messages with few total reactions. The ideal cutoff may depend on the space available for displaying reactions; the official web application displays names when 3 or fewer total reactions are present with this setting enabled.
//
//	**Changes**: New in Zulip 6.0 (feature level 125).
func (r UpdateRealmUserSettingsDefaultsRequest) DisplayEmojiReactionUsers(displayEmojiReactionUsers bool) UpdateRealmUserSettingsDefaultsRequest {
	r.displayEmojiReactionUsers = &displayEmojiReactionUsers
	return r
}

// The [home view] used when opening a new Zulip web app window or hitting the `Esc` keyboard shortcut repeatedly.  - \\"recent_topics\\" - Recent conversations view - \\"inbox\\" - Inbox view - \\"all_messages\\" - Combined feed view  **Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `default_view`, which was new in Zulip 4.0 (feature level 42).
//
// [home view]: https://zulip.com/help/configure-home-view
func (r UpdateRealmUserSettingsDefaultsRequest) WebHomeView(webHomeView string) UpdateRealmUserSettingsDefaultsRequest {
	r.webHomeView = &webHomeView
	return r
}

// Whether the escape key navigates to the [configured home view].
//
//	**Changes**: New in Zulip 8.0 (feature level 219). Previously, this was called `escape_navigates_to_default_view`, which was new in Zulip 5.0 (feature level 107).
//
// [configured home view]: https://zulip.com/help/configure-home-view
func (r UpdateRealmUserSettingsDefaultsRequest) WebEscapeNavigatesToHomeView(webEscapeNavigatesToHomeView bool) UpdateRealmUserSettingsDefaultsRequest {
	r.webEscapeNavigatesToHomeView = &webEscapeNavigatesToHomeView
	return r
}

// Whether the users list on left sidebar in narrow windows.  This feature is not heavily used and is likely to be reworked.
func (r UpdateRealmUserSettingsDefaultsRequest) LeftSideUserlist(leftSideUserlist bool) UpdateRealmUserSettingsDefaultsRequest {
	r.leftSideUserlist = &leftSideUserlist
	return r
}

// The user&#39;s configured [emoji set], used to display emoji to the user everywhere they appear in the UI.  - \\"google\\" - Google - \\"twitter\\" - Twitter - \\"text\\" - Plain text - \\"google-blob\\" - Google blobs
//
// [emoji set]: https://zulip.com/help/emoji-and-emoticons#use-emoticons
func (r UpdateRealmUserSettingsDefaultsRequest) Emojiset(emojiset string) UpdateRealmUserSettingsDefaultsRequest {
	r.emojiset = &emojiset
	return r
}

// Whether to [hide inactive channels] in the left sidebar.  - 1 - Automatic - 2 - Always - 3 - Never
//
// [hide inactive channels]: https://zulip.com/help/manage-inactive-channels
func (r UpdateRealmUserSettingsDefaultsRequest) DemoteInactiveChannels(demoteInactiveChannels DemoteInactiveChannels) UpdateRealmUserSettingsDefaultsRequest {
	r.demoteInactiveChannels = &demoteInactiveChannels
	return r
}

// The style selected by the user for the right sidebar user list.  - 1 - Compact - 2 - With status - 3 - With avatar and status  **Changes**: New in Zulip 6.0 (feature level 141).
func (r UpdateRealmUserSettingsDefaultsRequest) UserListStyle(userListStyle int32) UpdateRealmUserSettingsDefaultsRequest {
	r.userListStyle = &userListStyle
	return r
}

// Controls how animated images should be played in the message feed in the web/desktop application.  - \\"always\\" - Always play the animated images in the message feed. - \\"on_hover\\" - Play the animated images on hover over them in the message feed. - \\"never\\" - Never play animated images in the message feed.
//
//	**Changes**: New in Zulip 9.0 (feature level 275). Previously, animated images always used to play in the message feed by default. This setting controls this behaviour.
func (r UpdateRealmUserSettingsDefaultsRequest) WebAnimateImagePreviews(webAnimateImagePreviews string) UpdateRealmUserSettingsDefaultsRequest {
	r.webAnimateImagePreviews = &webAnimateImagePreviews
	return r
}

// Configuration for which channels should be displayed with a numeric unread count in the left sidebar. Channels that do not have an unread count will have a simple dot indicator for whether there are any unread messages.  - 1 - All channels - 2 - Unmuted channels and topics - 3 - No channels  **Changes**: New in Zulip 8.0 (feature level 210).
func (r UpdateRealmUserSettingsDefaultsRequest) WebChannelUnreadsCountDisplayPolicy(webChannelUnreadsCountDisplayPolicy int32) UpdateRealmUserSettingsDefaultsRequest {
	r.webChannelUnreadsCountDisplayPolicy = &webChannelUnreadsCountDisplayPolicy
	return r
}

// Controls whether user wants AI features like topic summarization to be hidden in all Zulip clients.
//
//	**Changes**: New in Zulip 10.0 (feature level 350).
func (r UpdateRealmUserSettingsDefaultsRequest) HideAiFeatures(hideAiFeatures bool) UpdateRealmUserSettingsDefaultsRequest {
	r.hideAiFeatures = &hideAiFeatures
	return r
}

// Determines whether the web/desktop application&#39;s left sidebar displays any channel folders configured by the organization.
//
//	**Changes**: New in Zulip 11.0 (feature level 411).
func (r UpdateRealmUserSettingsDefaultsRequest) WebLeftSidebarShowChannelFolders(webLeftSidebarShowChannelFolders bool) UpdateRealmUserSettingsDefaultsRequest {
	r.webLeftSidebarShowChannelFolders = &webLeftSidebarShowChannelFolders
	return r
}

// Determines whether the web/desktop application&#39;s left sidebar displays the unread message count summary.
//
//	**Changes**: New in Zulip 11.0 (feature level 398).
func (r UpdateRealmUserSettingsDefaultsRequest) WebLeftSidebarUnreadsCountSummary(webLeftSidebarUnreadsCountSummary bool) UpdateRealmUserSettingsDefaultsRequest {
	r.webLeftSidebarUnreadsCountSummary = &webLeftSidebarUnreadsCountSummary
	return r
}

// Enable visual desktop notifications for channel messages.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableChannelDesktopNotifications(enableChannelDesktopNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableChannelDesktopNotifications = &enableChannelDesktopNotifications
	return r
}

// Enable email notifications for channel messages.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableChannelEmailNotifications(enableChannelEmailNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableChannelEmailNotifications = &enableChannelEmailNotifications
	return r
}

// Enable mobile notifications for channel messages.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableChannelPushNotifications(enableChannelPushNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableChannelPushNotifications = &enableChannelPushNotifications
	return r
}

// Enable audible desktop notifications for channel messages.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableChannelAudibleNotifications(enableChannelAudibleNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableChannelAudibleNotifications = &enableChannelAudibleNotifications
	return r
}

// Notification sound name.
func (r UpdateRealmUserSettingsDefaultsRequest) NotificationSound(notificationSound string) UpdateRealmUserSettingsDefaultsRequest {
	r.notificationSound = &notificationSound
	return r
}

// Enable visual desktop notifications for direct messages and @-mentions.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableDesktopNotifications(enableDesktopNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableDesktopNotifications = &enableDesktopNotifications
	return r
}

// Enable audible desktop notifications for direct messages and @-mentions.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableSounds(enableSounds bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableSounds = &enableSounds
	return r
}

// Enable visual desktop notifications for messages sent to followed topics.
//
//	**Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateRealmUserSettingsDefaultsRequest) EnableFollowedTopicDesktopNotifications(enableFollowedTopicDesktopNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableFollowedTopicDesktopNotifications = &enableFollowedTopicDesktopNotifications
	return r
}

// Enable email notifications for messages sent to followed topics.
//
//	**Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateRealmUserSettingsDefaultsRequest) EnableFollowedTopicEmailNotifications(enableFollowedTopicEmailNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableFollowedTopicEmailNotifications = &enableFollowedTopicEmailNotifications
	return r
}

// Enable push notifications for messages sent to followed topics.
//
//	**Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateRealmUserSettingsDefaultsRequest) EnableFollowedTopicPushNotifications(enableFollowedTopicPushNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableFollowedTopicPushNotifications = &enableFollowedTopicPushNotifications
	return r
}

// Enable audible desktop notifications for messages sent to followed topics.
//
//	**Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateRealmUserSettingsDefaultsRequest) EnableFollowedTopicAudibleNotifications(enableFollowedTopicAudibleNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableFollowedTopicAudibleNotifications = &enableFollowedTopicAudibleNotifications
	return r
}

// The duration (in seconds) for which the server should wait to batch email notifications before sending them.
func (r UpdateRealmUserSettingsDefaultsRequest) EmailNotificationsBatchingPeriodSeconds(emailNotificationsBatchingPeriodSeconds int32) UpdateRealmUserSettingsDefaultsRequest {
	r.emailNotificationsBatchingPeriodSeconds = &emailNotificationsBatchingPeriodSeconds
	return r
}

// Enable email notifications for direct messages and @-mentions received when the user is offline.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableOfflineEmailNotifications(enableOfflineEmailNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableOfflineEmailNotifications = &enableOfflineEmailNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is offline.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableOfflinePushNotifications(enableOfflinePushNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableOfflinePushNotifications = &enableOfflinePushNotifications
	return r
}

// Enable mobile notification for direct messages and @-mentions received when the user is online.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableOnlinePushNotifications(enableOnlinePushNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableOnlinePushNotifications = &enableOnlinePushNotifications
	return r
}

// Enable digest emails when the user is away.
func (r UpdateRealmUserSettingsDefaultsRequest) EnableDigestEmails(enableDigestEmails bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableDigestEmails = &enableDigestEmails
	return r
}

// Include the message&#39;s content in email notifications for new messages.
func (r UpdateRealmUserSettingsDefaultsRequest) MessageContentInEmailNotifications(messageContentInEmailNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.messageContentInEmailNotifications = &messageContentInEmailNotifications
	return r
}

// Include content of direct messages in desktop notifications.
func (r UpdateRealmUserSettingsDefaultsRequest) PmContentInDesktopNotifications(pmContentInDesktopNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.pmContentInDesktopNotifications = &pmContentInDesktopNotifications
	return r
}

// Whether wildcard mentions (E.g. @**all**) should send notifications like a personal mention.
func (r UpdateRealmUserSettingsDefaultsRequest) WildcardMentionsNotify(wildcardMentionsNotify bool) UpdateRealmUserSettingsDefaultsRequest {
	r.wildcardMentionsNotify = &wildcardMentionsNotify
	return r
}

// Whether wildcard mentions (e.g., @**all**) in messages sent to followed topics should send notifications like a personal mention.
//
//	**Changes**: New in Zulip 8.0 (feature level 189).
func (r UpdateRealmUserSettingsDefaultsRequest) EnableFollowedTopicWildcardMentionsNotify(enableFollowedTopicWildcardMentionsNotify bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enableFollowedTopicWildcardMentionsNotify = &enableFollowedTopicWildcardMentionsNotify
	return r
}

// Unread count badge (appears in desktop sidebar and browser tab)  - 1 - All unread messages - 2 - DMs, mentions, and followed topics - 3 - DMs and mentions - 4 - None  **Changes**: In Zulip 8.0 (feature level 227), added `DMs, mentions, and followed topics` option, renumbering the options to insert it in order.
func (r UpdateRealmUserSettingsDefaultsRequest) DesktopIconCountDisplay(desktopIconCountDisplay int32) UpdateRealmUserSettingsDefaultsRequest {
	r.desktopIconCountDisplay = &desktopIconCountDisplay
	return r
}

// Whether to [include organization name in subject of message notification emails].  - 1 - Automatic - 2 - Always - 3 - Never  **Changes**: New in Zulip 7.0 (feature level 168), replacing the previous `realm_name_in_notifications` boolean; `true` corresponded to `Always`, and `false` to `Never`.
//
// [include organization name in subject of message notification emails]: https://zulip.com/help/email-notifications#include-organization-name-in-subject-line
func (r UpdateRealmUserSettingsDefaultsRequest) RealmNameInEmailNotificationsPolicy(realmNameInEmailNotificationsPolicy int32) UpdateRealmUserSettingsDefaultsRequest {
	r.realmNameInEmailNotificationsPolicy = &realmNameInEmailNotificationsPolicy
	return r
}

// Which [topics to follow automatically].  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
//
// [topics to follow automatically]: https://zulip.com/help/mute-a-topic
func (r UpdateRealmUserSettingsDefaultsRequest) AutomaticallyFollowTopicsPolicy(automaticallyFollowTopicsPolicy int32) UpdateRealmUserSettingsDefaultsRequest {
	r.automaticallyFollowTopicsPolicy = &automaticallyFollowTopicsPolicy
	return r
}

// Which [topics to unmute automatically in muted channels].  - 1 - Topics the user participates in - 2 - Topics the user sends a message to - 3 - Topics the user starts - 4 - Never  **Changes**: New in Zulip 8.0 (feature level 214).
//
// [topics to unmute automatically in muted channels]: https://zulip.com/help/mute-a-topic
func (r UpdateRealmUserSettingsDefaultsRequest) AutomaticallyUnmuteTopicsInMutedChannelsPolicy(automaticallyUnmuteTopicsInMutedChannelsPolicy int32) UpdateRealmUserSettingsDefaultsRequest {
	r.automaticallyUnmuteTopicsInMutedChannelsPolicy = &automaticallyUnmuteTopicsInMutedChannelsPolicy
	return r
}

// Whether the server will automatically mark the user as following topics where the user is mentioned.
//
//	**Changes**: New in Zulip 8.0 (feature level 235).
func (r UpdateRealmUserSettingsDefaultsRequest) AutomaticallyFollowTopicsWhereMentioned(automaticallyFollowTopicsWhereMentioned bool) UpdateRealmUserSettingsDefaultsRequest {
	r.automaticallyFollowTopicsWhereMentioned = &automaticallyFollowTopicsWhereMentioned
	return r
}

// Controls whether the resolved-topic notices are marked as read.  - \\"always\\" - Always mark resolved-topic notices as read. - \\"except_followed\\" - Mark resolved-topic notices as read in topics not followed by the user. - \\"never\\" - Never mark resolved-topic notices as read.
//
//	**Changes**: New in Zulip 11.0 (feature level 385).
func (r UpdateRealmUserSettingsDefaultsRequest) ResolvedTopicNoticeAutoReadPolicy(resolvedTopicNoticeAutoReadPolicy string) UpdateRealmUserSettingsDefaultsRequest {
	r.resolvedTopicNoticeAutoReadPolicy = &resolvedTopicNoticeAutoReadPolicy
	return r
}

// Display the presence status to other users when online.
func (r UpdateRealmUserSettingsDefaultsRequest) PresenceEnabled(presenceEnabled bool) UpdateRealmUserSettingsDefaultsRequest {
	r.presenceEnabled = &presenceEnabled
	return r
}

// Whether pressing Enter in the compose box sends a message (or saves a message edit).
func (r UpdateRealmUserSettingsDefaultsRequest) EnterSends(enterSends bool) UpdateRealmUserSettingsDefaultsRequest {
	r.enterSends = &enterSends
	return r
}

// Whether time should be [displayed in 24-hour notation].
//
//	**Changes**: New in Zulip 5.0 (feature level 99). Previously, this default was edited using the `default_twenty_four_hour_time` parameter to the `PATCH /realm` endpoint.
//
// [displayed in 24-hour notation]: https://zulip.com/help/change-the-time-format
func (r UpdateRealmUserSettingsDefaultsRequest) TwentyFourHourTime(twentyFourHourTime bool) UpdateRealmUserSettingsDefaultsRequest {
	r.twentyFourHourTime = &twentyFourHourTime
	return r
}

// Whether [typing notifications] be sent when composing direct messages.
//
//	**Changes**: New in Zulip 5.0 (feature level 105).
//
// [typing notifications]: https://zulip.com/help/typing-notifications
func (r UpdateRealmUserSettingsDefaultsRequest) SendPrivateTypingNotifications(sendPrivateTypingNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.sendPrivateTypingNotifications = &sendPrivateTypingNotifications
	return r
}

// Whether [typing notifications] be sent when composing channel messages.
//
//	**Changes**: New in Zulip 5.0 (feature level 105).
//
// [typing notifications]: https://zulip.com/help/typing-notifications
func (r UpdateRealmUserSettingsDefaultsRequest) SendChannelTypingNotifications(sendChannelTypingNotifications bool) UpdateRealmUserSettingsDefaultsRequest {
	r.sendChannelTypingNotifications = &sendChannelTypingNotifications
	return r
}

// Whether other users are allowed to see whether you&#39;ve read messages.
//
//	**Changes**: New in Zulip 5.0 (feature level 105).
func (r UpdateRealmUserSettingsDefaultsRequest) SendReadReceipts(sendReadReceipts bool) UpdateRealmUserSettingsDefaultsRequest {
	r.sendReadReceipts = &sendReadReceipts
	return r
}

// The [policy] for [which other users] in this organization can see the user&#39;s real email address.  - 1 &#x3D; Everyone - 2 &#x3D; Members only - 3 &#x3D; Administrators only - 4 &#x3D; Nobody - 5 &#x3D; Moderators only  **Changes**: New in Zulip 7.0 (feature level 163), replacing the realm-level setting.
//
// [policy]: https://zulip.com/api/roles-and-permissions#permission-levels
// [which other users]: https://zulip.com/help/configure-email-visibility
func (r UpdateRealmUserSettingsDefaultsRequest) EmailAddressVisibility(emailAddressVisibility int32) UpdateRealmUserSettingsDefaultsRequest {
	r.emailAddressVisibility = &emailAddressVisibility
	return r
}

// Web/desktop app setting for whether the user&#39;s view should automatically go to the conversation where they sent a message.
//
//	**Changes**: New in Zulip 9.0 (feature level 268). Previously, this behavior was not configurable.
func (r UpdateRealmUserSettingsDefaultsRequest) WebNavigateToSentMessage(webNavigateToSentMessage bool) UpdateRealmUserSettingsDefaultsRequest {
	r.webNavigateToSentMessage = &webNavigateToSentMessage
	return r
}

func (r UpdateRealmUserSettingsDefaultsRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UpdateRealmUserSettingsDefaultsExecute(r)
}

// UpdateRealmUserSettingsDefaults Update realm-level defaults of user settings
//
// Change the [default values of settings] for new users
// joining the organization. Essentially all
// [personal preference settings] are supported.
//
// This feature can be invaluable for customizing Zulip's default
// settings for notifications or UI to be appropriate for how the
// organization is using Zulip. (Note that this only supports
// personal preference settings, like when to send push
// notifications or what emoji set to use, not profile or
// identity settings that naturally should be different for each user).
//
// Note that this endpoint cannot, at present, be used to modify
// settings for existing users in any way.
//
// *Changes**: Removed `dense_mode` setting in Zulip 10.0 (feature level 364)
// as we now have `web_font_size_px` and `web_line_height_percent`
// settings for more control.
//
// New in Zulip 5.0 (feature level 96). If any parameters sent in the
// request are not supported by this endpoint, an
// [`ignored_parameters_unsupported`] array will
// be returned in the JSON success response.
//
// [default values of settings]: https://zulip.com/help/configure-default-new-user-settings
// [`ignored_parameters_unsupported`]: https://zulip.com/api/rest-error-handling#ignored-parameters
// [personal preference settings]: https://zulip.com/api/update-settings
func (c *simpleClient) UpdateRealmUserSettingsDefaults(ctx context.Context) UpdateRealmUserSettingsDefaultsRequest {
	return UpdateRealmUserSettingsDefaultsRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
func (c *simpleClient) UpdateRealmUserSettingsDefaultsExecute(r UpdateRealmUserSettingsDefaultsRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/user_settings_defaults"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.starredMessageCounts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "starred_message_counts", r.starredMessageCounts, "form", "")
	}
	if r.receivesTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "receives_typing_notifications", r.receivesTypingNotifications, "form", "")
	}
	if r.webSuggestUpdateTimezone != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_suggest_update_timezone", r.webSuggestUpdateTimezone, "form", "")
	}
	if r.fluidLayoutWidth != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fluid_layout_width", r.fluidLayoutWidth, "form", "")
	}
	if r.highContrastMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "high_contrast_mode", r.highContrastMode, "form", "")
	}
	if r.webMarkReadOnScrollPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_mark_read_on_scroll_policy", r.webMarkReadOnScrollPolicy, "form", "")
	}
	if r.webChannelDefaultView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_channel_default_view", r.webChannelDefaultView, "form", "")
	}
	if r.webFontSizePx != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_font_size_px", r.webFontSizePx, "form", "")
	}
	if r.webLineHeightPercent != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_line_height_percent", r.webLineHeightPercent, "form", "")
	}
	if r.colorScheme != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color_scheme", r.colorScheme, "form", "")
	}
	if r.enableDraftsSynchronization != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_drafts_synchronization", r.enableDraftsSynchronization, "form", "")
	}
	if r.translateEmoticons != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "translate_emoticons", r.translateEmoticons, "form", "")
	}
	if r.displayEmojiReactionUsers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "display_emoji_reaction_users", r.displayEmojiReactionUsers, "form", "")
	}
	if r.webHomeView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_home_view", r.webHomeView, "", "")
	}
	if r.webEscapeNavigatesToHomeView != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_escape_navigates_to_home_view", r.webEscapeNavigatesToHomeView, "form", "")
	}
	if r.leftSideUserlist != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "left_side_userlist", r.leftSideUserlist, "form", "")
	}
	if r.emojiset != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "emojiset", r.emojiset, "", "")
	}
	if r.demoteInactiveChannels != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "demote_inactive_streams", r.demoteInactiveChannels, "form", "")
	}
	if r.userListStyle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_list_style", r.userListStyle, "form", "")
	}
	if r.webAnimateImagePreviews != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_animate_image_previews", r.webAnimateImagePreviews, "", "")
	}
	if r.webChannelUnreadsCountDisplayPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_stream_unreads_count_display_policy", r.webChannelUnreadsCountDisplayPolicy, "form", "")
	}
	if r.hideAiFeatures != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_ai_features", r.hideAiFeatures, "form", "")
	}
	if r.webLeftSidebarShowChannelFolders != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_left_sidebar_show_channel_folders", r.webLeftSidebarShowChannelFolders, "form", "")
	}
	if r.webLeftSidebarUnreadsCountSummary != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_left_sidebar_unreads_count_summary", r.webLeftSidebarUnreadsCountSummary, "form", "")
	}
	if r.enableChannelDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_desktop_notifications", r.enableChannelDesktopNotifications, "form", "")
	}
	if r.enableChannelEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_email_notifications", r.enableChannelEmailNotifications, "form", "")
	}
	if r.enableChannelPushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_push_notifications", r.enableChannelPushNotifications, "form", "")
	}
	if r.enableChannelAudibleNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_stream_audible_notifications", r.enableChannelAudibleNotifications, "form", "")
	}
	if r.notificationSound != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_sound", r.notificationSound, "", "")
	}
	if r.enableDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_desktop_notifications", r.enableDesktopNotifications, "form", "")
	}
	if r.enableSounds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_sounds", r.enableSounds, "form", "")
	}
	if r.enableFollowedTopicDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_desktop_notifications", r.enableFollowedTopicDesktopNotifications, "form", "")
	}
	if r.enableFollowedTopicEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_email_notifications", r.enableFollowedTopicEmailNotifications, "form", "")
	}
	if r.enableFollowedTopicPushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_push_notifications", r.enableFollowedTopicPushNotifications, "form", "")
	}
	if r.enableFollowedTopicAudibleNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_audible_notifications", r.enableFollowedTopicAudibleNotifications, "form", "")
	}
	if r.emailNotificationsBatchingPeriodSeconds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_notifications_batching_period_seconds", r.emailNotificationsBatchingPeriodSeconds, "form", "")
	}
	if r.enableOfflineEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_offline_email_notifications", r.enableOfflineEmailNotifications, "form", "")
	}
	if r.enableOfflinePushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_offline_push_notifications", r.enableOfflinePushNotifications, "form", "")
	}
	if r.enableOnlinePushNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_online_push_notifications", r.enableOnlinePushNotifications, "form", "")
	}
	if r.enableDigestEmails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_digest_emails", r.enableDigestEmails, "form", "")
	}
	if r.messageContentInEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_content_in_email_notifications", r.messageContentInEmailNotifications, "form", "")
	}
	if r.pmContentInDesktopNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pm_content_in_desktop_notifications", r.pmContentInDesktopNotifications, "form", "")
	}
	if r.wildcardMentionsNotify != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "wildcard_mentions_notify", r.wildcardMentionsNotify, "form", "")
	}
	if r.enableFollowedTopicWildcardMentionsNotify != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enable_followed_topic_wildcard_mentions_notify", r.enableFollowedTopicWildcardMentionsNotify, "form", "")
	}
	if r.desktopIconCountDisplay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "desktop_icon_count_display", r.desktopIconCountDisplay, "form", "")
	}
	if r.realmNameInEmailNotificationsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "realm_name_in_email_notifications_policy", r.realmNameInEmailNotificationsPolicy, "form", "")
	}
	if r.automaticallyFollowTopicsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_follow_topics_policy", r.automaticallyFollowTopicsPolicy, "form", "")
	}
	if r.automaticallyUnmuteTopicsInMutedChannelsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_unmute_topics_in_muted_streams_policy", r.automaticallyUnmuteTopicsInMutedChannelsPolicy, "form", "")
	}
	if r.automaticallyFollowTopicsWhereMentioned != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "automatically_follow_topics_where_mentioned", r.automaticallyFollowTopicsWhereMentioned, "form", "")
	}
	if r.resolvedTopicNoticeAutoReadPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "resolved_topic_notice_auto_read_policy", r.resolvedTopicNoticeAutoReadPolicy, "", "")
	}
	if r.presenceEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "presence_enabled", r.presenceEnabled, "form", "")
	}
	if r.enterSends != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "enter_sends", r.enterSends, "form", "")
	}
	if r.twentyFourHourTime != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "twenty_four_hour_time", r.twentyFourHourTime, "form", "")
	}
	if r.sendPrivateTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_private_typing_notifications", r.sendPrivateTypingNotifications, "form", "")
	}
	if r.sendChannelTypingNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_stream_typing_notifications", r.sendChannelTypingNotifications, "form", "")
	}
	if r.sendReadReceipts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "send_read_receipts", r.sendReadReceipts, "form", "")
	}
	if r.emailAddressVisibility != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_address_visibility", r.emailAddressVisibility, "form", "")
	}
	if r.webNavigateToSentMessage != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "web_navigate_to_sent_message", r.webNavigateToSentMessage, "form", "")
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UploadCustomEmojiRequest struct {
	ctx        context.Context
	ApiService ServerAndOrganizationsAPI
	emojiName  string
	filename   *os.File
}

func (r UploadCustomEmojiRequest) Filename(filename *os.File) UploadCustomEmojiRequest {
	r.filename = filename
	return r
}

func (r UploadCustomEmojiRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.UploadCustomEmojiExecute(r)
}

// UploadCustomEmoji Upload custom emoji
//
// This endpoint is used to upload a custom emoji for use in the user's
// organization. Access to this endpoint depends on the
// [organization's configuration].
//
// [organization's configuration]: https://zulip.com/help/custom-emoji#change-who-can-add-custom-emoji
func (c *simpleClient) UploadCustomEmoji(ctx context.Context, emojiName string) UploadCustomEmojiRequest {
	return UploadCustomEmojiRequest{
		ApiService: c,
		ctx:        ctx,
		emojiName:  emojiName,
	}
}

// Execute executes the request
func (c *simpleClient) UploadCustomEmojiExecute(r UploadCustomEmojiRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/realm/emoji/{emoji_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"emoji_name"+"}", url.PathEscape(parameterValueToString(r.emojiName, "emojiName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var filenameLocalVarFormFileName string
	var filenameLocalVarFileName string
	var filenameLocalVarFileBytes []byte

	filenameLocalVarFormFileName = "filename"
	filenameLocalVarFile := r.filename

	if filenameLocalVarFile != nil {
		fbs, _ := io.ReadAll(filenameLocalVarFile)

		filenameLocalVarFileBytes = fbs
		filenameLocalVarFileName = filenameLocalVarFile.Name()
		filenameLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: filenameLocalVarFileBytes, fileName: filenameLocalVarFileName, formFileName: filenameLocalVarFormFileName})
	}
	req, err := c.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := c.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, c.handleErrorResponse(r.ctx, localVarHTTPResponse)
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}
