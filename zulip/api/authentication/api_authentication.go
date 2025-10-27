// Package authentication provides API methods for Zulip authentication,
// including API key management and authentication-related operations.
package authentication

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"
)

type APIAuthentication interface {
	// DevFetchAPIKey Fetch an API key (development only)
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
	DevFetchAPIKey(ctx context.Context) DevFetchAPIKeyRequest

	// DevFetchAPIKeyExecute executes the request
	DevFetchAPIKeyExecute(r DevFetchAPIKeyRequest) (*APIKeyResponse, *http.Response, error)

	// FetchAPIKey Fetch an API key (production)
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
	// [reset your password]: https://zulip.com/help/change-your-password
	// [API keys]: https://zulip.com/api/api-keys
	// [Zulip development environment]: https://zulip.readthedocs.io/en/latest/development/overview.html
	// [the unauthenticated variant]: https://zulip.com/api/dev-fetch-api-key
	FetchAPIKey(ctx context.Context) FetchAPIKeyRequest

	// FetchAPIKeyExecute executes the request
	FetchAPIKeyExecute(r FetchAPIKeyRequest) (*APIKeyResponse, *http.Response, error)
}

type authenticationService struct {
	client clients.Client
}

func NewAuthenticationService(client clients.Client) APIAuthentication {
	return &authenticationService{client: client}
}

var _ APIAuthentication = (*authenticationService)(nil)

type DevFetchAPIKeyRequest struct {
	ctx        context.Context
	apiService APIAuthentication
	username   *string
}

// The email address for the user that owns the API key.
func (r DevFetchAPIKeyRequest) Username(username string) DevFetchAPIKeyRequest {
	r.username = &username
	return r
}

func (r DevFetchAPIKeyRequest) Execute() (*APIKeyResponse, *http.Response, error) {
	return r.apiService.DevFetchAPIKeyExecute(r)
}

// DevFetchAPIKey Fetch an API key (development only)
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
func (s *authenticationService) DevFetchAPIKey(ctx context.Context) DevFetchAPIKeyRequest {
	return DevFetchAPIKeyRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *authenticationService) DevFetchAPIKeyExecute(
	r DevFetchAPIKeyRequest,
) (*APIKeyResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &APIKeyResponse{}
		endpoint = "/dev_fetch_api_key"
	)
	if r.username == nil {
		return nil, nil, errors.New("username is required and must be specified")
	}

	headers["Content-Type"] = apiutils.ContentTypeFormURLEncoded
	headers["Accept"] = apiutils.ContentTypeJSON

	apiutils.AddParam(form, "username", r.username)
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}

type FetchAPIKeyRequest struct {
	ctx        context.Context
	apiService APIAuthentication
	username   *string
	password   *string
}

// The username to be used for authentication (typically, the email address, but depending on configuration, it could be an LDAP username).  See the `require_email_format_usernames` parameter documented in [GET /server_settings] for details.
//
// [GET /server_settings]: https://zulip.com/api/get-server-settings
func (r FetchAPIKeyRequest) Username(username string) FetchAPIKeyRequest {
	r.username = &username
	return r
}

// The user's Zulip password (or LDAP password, if LDAP authentication is in use).
func (r FetchAPIKeyRequest) Password(password string) FetchAPIKeyRequest {
	r.password = &password
	return r
}

func (r FetchAPIKeyRequest) Execute() (*APIKeyResponse, *http.Response, error) {
	return r.apiService.FetchAPIKeyExecute(r)
}

// FetchAPIKey Fetch an API key (production)
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
// [reset your password]: https://zulip.com/help/change-your-password
// [API keys]: https://zulip.com/api/api-keys
// [Zulip development environment]: https://zulip.readthedocs.io/en/latest/development/overview.html
// [the unauthenticated variant]: https://zulip.com/api/dev-fetch-api-key
func (s *authenticationService) FetchAPIKey(ctx context.Context) FetchAPIKeyRequest {
	return FetchAPIKeyRequest{
		apiService: s,
		ctx:        ctx,
	}
}

// Execute executes the request.
func (s *authenticationService) FetchAPIKeyExecute(r FetchAPIKeyRequest) (*APIKeyResponse, *http.Response, error) {
	var (
		method   = http.MethodPost
		headers  = make(map[string]string)
		query    = url.Values{}
		form     = url.Values{}
		response = &APIKeyResponse{}
		endpoint = "/fetch_api_key"
	)
	if r.username == nil {
		return nil, nil, errors.New("username is required and must be specified")
	}
	if r.password == nil {
		return nil, nil, errors.New("password is required and must be specified")
	}

	headers["Content-Type"] = apiutils.ContentTypeFormURLEncoded
	headers["Accept"] = apiutils.ContentTypeJSON

	apiutils.AddParam(form, "username", r.username)
	apiutils.AddParam(form, "password", r.password)
	req, err := apiutils.PrepareRequest(r.ctx, s.client, endpoint, method, headers, query, form, nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := s.client.CallAPI(r.ctx, endpoint, req, response)
	return response, httpResp, err
}
