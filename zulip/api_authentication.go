package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type AuthenticationAPI interface {

	/*
			DevFetchApiKey Fetch an API key (development only)

			For easy testing of mobile apps and other clients and against Zulip
		development servers, we support fetching a Zulip API key for any user
		on the development server without authentication (so that they can
		implement analogues of the one-click login process available for Zulip
		development servers on the web).

		!!! warn ""

		    **Note:** This endpoint is only available on Zulip development
		    servers; for obvious security reasons it will always return an error
		    in a Zulip production server.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return DevFetchApiKeyRequest
	*/
	DevFetchApiKey(ctx context.Context) DevFetchApiKeyRequest

	// DevFetchApiKeyExecute executes the request
	//  @return ApiKeyResponse
	DevFetchApiKeyExecute(r DevFetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error)

	/*
			FetchApiKey Fetch an API key (production)

			This API endpoint is used by clients such as the Zulip mobile and
		terminal apps to implement password-based authentication. Given the
		user's Zulip login credentials, it returns a Zulip API key that the client
		can use to make requests as the user.

		This endpoint is only useful for Zulip servers/organizations with
		EmailAuthBackend or LDAPAuthBackend enabled.

		The Zulip mobile apps also support SSO/social authentication (GitHub
		auth, Google auth, SAML, etc.) that does not use this endpoint. Instead,
		the mobile apps reuse the web login flow passing the `mobile_flow_otp` in
		a webview, and the credentials are returned to the app (encrypted) via a redirect
		to a `zulip://` URL.

		!!! warn ""

		    **Note:** If you signed up using passwordless authentication and
		    never had a password, you can [reset your password](zulip.com/help/change-your-password.

		See the [API keys](zulip.com/api/api-keys) documentation for more details
		on how to download an API key manually.

		In a [Zulip development environment](https://zulip.readthedocs.io/en/latest/development/overview.html),
		see also [the unauthenticated variant](zulip.com/api/dev-fetch-api-key.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return FetchApiKeyRequest
	*/
	FetchApiKey(ctx context.Context) FetchApiKeyRequest

	// FetchApiKeyExecute executes the request
	//  @return ApiKeyResponse
	FetchApiKeyExecute(r FetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error)
}

type DevFetchApiKeyRequest struct {
	ctx        context.Context
	ApiService AuthenticationAPI
	username   *string
}

// The email address for the user that owns the API key.
func (r DevFetchApiKeyRequest) Username(username string) DevFetchApiKeyRequest {
	r.username = &username
	return r
}

func (r DevFetchApiKeyRequest) Execute() (*ApiKeyResponse, *http.Response, error) {
	return r.ApiService.DevFetchApiKeyExecute(r)
}

/*
DevFetchApiKey Fetch an API key (development only)

For easy testing of mobile apps and other clients and against Zulip
development servers, we support fetching a Zulip API key for any user
on the development server without authentication (so that they can
implement analogues of the one-click login process available for Zulip
development servers on the web).

!!! warn ""

	   **Note:** This endpoint is only available on Zulip development
	   servers; for obvious security reasons it will always return an error
	   in a Zulip production server.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DevFetchApiKeyRequest
*/
func (c *simpleClient) DevFetchApiKey(ctx context.Context) DevFetchApiKeyRequest {
	return DevFetchApiKeyRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiKeyResponse
func (c *simpleClient) DevFetchApiKeyExecute(r DevFetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiKeyResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dev_fetch_api_key"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.username == nil {
		return localVarReturnValue, nil, reportError("username is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "username", r.username, "", "")
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

type FetchApiKeyRequest struct {
	ctx        context.Context
	ApiService AuthenticationAPI
	username   *string
	password   *string
}

// The username to be used for authentication (typically, the email address, but depending on configuration, it could be an LDAP username).  See the &#x60;require_email_format_usernames&#x60; parameter documented in [GET /server_settings](zulip.com/api/get-server-settings) for details.
func (r FetchApiKeyRequest) Username(username string) FetchApiKeyRequest {
	r.username = &username
	return r
}

// The user&#39;s Zulip password (or LDAP password, if LDAP authentication is in use).
func (r FetchApiKeyRequest) Password(password string) FetchApiKeyRequest {
	r.password = &password
	return r
}

func (r FetchApiKeyRequest) Execute() (*ApiKeyResponse, *http.Response, error) {
	return r.ApiService.FetchApiKeyExecute(r)
}

/*
FetchApiKey Fetch an API key (production)

This API endpoint is used by clients such as the Zulip mobile and
terminal apps to implement password-based authentication. Given the
user's Zulip login credentials, it returns a Zulip API key that the client
can use to make requests as the user.

This endpoint is only useful for Zulip servers/organizations with
EmailAuthBackend or LDAPAuthBackend enabled.

The Zulip mobile apps also support SSO/social authentication (GitHub
auth, Google auth, SAML, etc.) that does not use this endpoint. Instead,
the mobile apps reuse the web login flow passing the `mobile_flow_otp` in
a webview, and the credentials are returned to the app (encrypted) via a redirect
to a `zulip://` URL.

!!! warn ""

	**Note:** If you signed up using passwordless authentication and
	never had a password, you can [reset your password](zulip.com/help/change-your-password.

See the [API keys](zulip.com/api/api-keys) documentation for more details
on how to download an API key manually.

In a [Zulip development environment](https://zulip.readthedocs.io/en/latest/development/overview.html),
see also [the unauthenticated variant](zulip.com/api/dev-fetch-api-key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FetchApiKeyRequest
*/
func (c *simpleClient) FetchApiKey(ctx context.Context) FetchApiKeyRequest {
	return FetchApiKeyRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiKeyResponse
func (c *simpleClient) FetchApiKeyExecute(r FetchApiKeyRequest) (*ApiKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiKeyResponse
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &APIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fetch_api_key"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.username == nil {
		return localVarReturnValue, nil, reportError("username is required and must be specified")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "username", r.username, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
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
