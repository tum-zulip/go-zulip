package zulip

import (
	"context"
	"net/http"
	"net/url"
)

type AuthenticationAPI interface {

	// DevFetchApiKey Fetch an API key (development only)
	//
	// For easy testing of mobile apps and other clients and against Zulip
	// development servers, we support fetching a Zulip API key for any user
	// on the development server without authentication (so that they can
	// implement analogues of the one-click login process available for Zulip
	// development servers on the web).
	//
	// !!! warn ""
	//
	// *Note:** This endpoint is only available on Zulip development
	// servers; for obvious security reasons it will always return an error
	// in a Zulip production server.
	//
	//
	// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	// @return DevFetchApiKeyRequest
	DevFetchApiKey(ctx context.Context) DevFetchApiKeyRequest

	// DevFetchApiKeyExecute executes the request
	//  @return ApiKeyResponse
	DevFetchApiKeyExecute(r DevFetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error)

	// FetchApiKey Fetch an API key (production)
	//
	// This API endpoint is used by clients such as the Zulip mobile and
	// terminal apps to implement password-based authentication. Given the
	// user's Zulip login credentials, it returns a Zulip API key that the client
	// can use to make requests as the user.
	//
	// This endpoint is only useful for Zulip servers/organizations with
	// EmailAuthBackend or LDAPAuthBackend enabled.
	//
	// The Zulip mobile apps also support SSO/social authentication (GitHub
	// auth, Google auth, SAML, etc.) that does not use this endpoint. Instead,
	// the mobile apps reuse the web login flow passing the `mobile_flow_otp` in
	// a webview, and the credentials are returned to the app (encrypted) via a redirect
	// to a `zulip://` URL.
	//
	// !!! warn ""
	//
	// *Note:** If you signed up using passwordless authentication and
	// never had a password, you can [reset your password].
	//
	// See the [API keys] documentation for more details
	// on how to download an API key manually.
	//
	// In a [Zulip development environment], see also [the unauthenticated variant].
	//
	//
	// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	// @return FetchApiKeyRequest
	//
	// [reset your password]: https://zulip.com/help/change-your-password
	// [API keys]: https://zulip.com/api/api-keys
	// [Zulip development environment]: https://zulip.readthedocs.io/en/latest/development/overview.html
	// [the unauthenticated variant]: https://zulip.com/api/dev-fetch-api-key
	FetchApiKey(ctx context.Context) FetchApiKeyRequest

	// FetchApiKeyExecute executes the request
	//  @return ApiKeyResponse
	FetchApiKeyExecute(r FetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error)
}

type DevFetchApiKeyRequest struct {
	ctx        context.Context
	apiService AuthenticationAPI
	username   *string
}

// The email address for the user that owns the API key.
func (r DevFetchApiKeyRequest) Username(username string) DevFetchApiKeyRequest {
	r.username = &username
	return r
}

func (r DevFetchApiKeyRequest) Execute() (*ApiKeyResponse, *http.Response, error) {
	return r.apiService.DevFetchApiKeyExecute(r)
}

// DevFetchApiKey Fetch an API key (development only)
//
// For easy testing of mobile apps and other clients and against Zulip
// development servers, we support fetching a Zulip API key for any user
// on the development server without authentication (so that they can
// implement analogues of the one-click login process available for Zulip
// development servers on the web).
//
// !!! warn ""
//
// *Note:** This endpoint is only available on Zulip development
// servers; for obvious security reasons it will always return an error
// in a Zulip production server.
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return DevFetchApiKeyRequest
func (c *simpleClient) DevFetchApiKey(ctx context.Context) DevFetchApiKeyRequest {
	return DevFetchApiKeyRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiKeyResponse
func (c *simpleClient) DevFetchApiKeyExecute(r DevFetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &ApiKeyResponse{}
		endpoint = "/dev_fetch_api_key"
	)
	if r.username == nil {
		return nil, nil, reportError("username is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "username", r.username)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}

type FetchApiKeyRequest struct {
	ctx        context.Context
	apiService AuthenticationAPI
	username   *string
	password   *string
}

// The username to be used for authentication (typically, the email address, but depending on configuration, it could be an LDAP username).  See the `require_email_format_usernames` parameter documented in [GET /server_settings] for details.
//
// [GET /server_settings]: https://zulip.com/api/get-server-settings
func (r FetchApiKeyRequest) Username(username string) FetchApiKeyRequest {
	r.username = &username
	return r
}

// The user's Zulip password (or LDAP password, if LDAP authentication is in use).
func (r FetchApiKeyRequest) Password(password string) FetchApiKeyRequest {
	r.password = &password
	return r
}

func (r FetchApiKeyRequest) Execute() (*ApiKeyResponse, *http.Response, error) {
	return r.apiService.FetchApiKeyExecute(r)
}

// FetchApiKey Fetch an API key (production)
//
// This API endpoint is used by clients such as the Zulip mobile and
// terminal apps to implement password-based authentication. Given the
// user's Zulip login credentials, it returns a Zulip API key that the client
// can use to make requests as the user.
//
// This endpoint is only useful for Zulip servers/organizations with
// EmailAuthBackend or LDAPAuthBackend enabled.
//
// The Zulip mobile apps also support SSO/social authentication (GitHub
// auth, Google auth, SAML, etc.) that does not use this endpoint. Instead,
// the mobile apps reuse the web login flow passing the `mobile_flow_otp` in
// a webview, and the credentials are returned to the app (encrypted) via a redirect
// to a `zulip://` URL.
//
// !!! warn ""
//
// *Note:** If you signed up using passwordless authentication and
// never had a password, you can [reset your password].
//
// See the [API keys] documentation for more details
// on how to download an API key manually.
//
// In a [Zulip development environment], see also [the unauthenticated variant].
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return FetchApiKeyRequest
//
// [reset your password]: https://zulip.com/help/change-your-password
// [API keys]: https://zulip.com/api/api-keys
// [Zulip development environment]: https://zulip.readthedocs.io/en/latest/development/overview.html
// [the unauthenticated variant]: https://zulip.com/api/dev-fetch-api-key
func (c *simpleClient) FetchApiKey(ctx context.Context) FetchApiKeyRequest {
	return FetchApiKeyRequest{
		apiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiKeyResponse
func (c *simpleClient) FetchApiKeyExecute(r FetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &ApiKeyResponse{}
		endpoint = "/fetch_api_key"
	)
	if r.username == nil {
		return nil, nil, reportError("username is required and must be specified")
	}
	if r.password == nil {
		return nil, nil, reportError("password is required and must be specified")
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Accept"] = "application/json"

	addParam(form, "username", r.username)
	addParam(form, "password", r.password)
	req, err := c.prepareRequest(r.ctx, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := c.callAPI(r.ctx, req, response)
	return response, httpResp, err
}
