package zulip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type MobileAPI interface {

	/*
			E2eeTestNotify Send an E2EE test notification to mobile device(s)

			Trigger sending an end-to-end encrypted (E2EE) test push notification
		to the user's selected mobile device or all of their mobile devices.

		**Changes**: New in Zulip 11.0 (feature level 420).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return E2eeTestNotifyRequest
	*/
	E2eeTestNotify(ctx context.Context) E2eeTestNotifyRequest

	// E2eeTestNotifyExecute executes the request
	//  @return Response
	E2eeTestNotifyExecute(r E2eeTestNotifyRequest) (*Response, *http.Response, error)

	/*
			RegisterPushDevice Register E2EE push device

			Register a device to receive end-to-end encrypted mobile push notifications.

		**Changes**: New in Zulip 11.0 (feature level 406).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return RegisterPushDeviceRequest
	*/
	RegisterPushDevice(ctx context.Context) RegisterPushDeviceRequest

	// RegisterPushDeviceExecute executes the request
	//  @return Response
	RegisterPushDeviceExecute(r RegisterPushDeviceRequest) (*Response, *http.Response, error)

	/*
			TestNotify Send a test notification to mobile device(s)

			Trigger sending a test push notification to the user's
		selected mobile device or all of their mobile devices.

		**Changes**: Starting with Zulip 8.0 (feature level 234), test
		notifications sent via this endpoint use `test` rather than
		`test-by-device-token` in the `event` field. Also, as of this
		feature level, all mobile push notifications now include a
		`realm_name` field.

		New in Zulip 8.0 (feature level 217).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return TestNotifyRequest
	*/
	TestNotify(ctx context.Context) TestNotifyRequest

	// TestNotifyExecute executes the request
	//  @return Response
	TestNotifyExecute(r TestNotifyRequest) (*Response, *http.Response, error)
}

type E2eeTestNotifyRequest struct {
	ctx           context.Context
	ApiService    MobileAPI
	pushAccountId *int64
}

// The push account Id for the device to which to send the test notification.  If this parameter is not submitted, the E2EE test notification will be sent to all of the user&#39;s devices registered on the server.  A mobile client should pass this parameter, to avoid triggering a test notification for other clients.  See [&#x60;POST /mobile_push/register&#x60;](zulip.com/api/register-push-device) for details on push account Ids.
func (r E2eeTestNotifyRequest) PushAccountId(pushAccountId int64) E2eeTestNotifyRequest {
	r.pushAccountId = &pushAccountId
	return r
}

func (r E2eeTestNotifyRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.E2eeTestNotifyExecute(r)
}

/*
E2eeTestNotify Send an E2EE test notification to mobile device(s)

Trigger sending an end-to-end encrypted (E2EE) test push notification
to the user's selected mobile device or all of their mobile devices.

**Changes**: New in Zulip 11.0 (feature level 420).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return E2eeTestNotifyRequest
*/
func (c *simpleClient) E2eeTestNotify(ctx context.Context) E2eeTestNotifyRequest {
	return E2eeTestNotifyRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) E2eeTestNotifyExecute(r E2eeTestNotifyRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mobile_push/e2ee/test_notification"

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
	if r.pushAccountId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "push_account_id", r.pushAccountId, "", "")
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type RegisterPushDeviceRequest struct {
	ctx                       context.Context
	ApiService                MobileAPI
	tokenKind                 *string
	pushAccountId             *int64
	pushPublicKey             *string
	bouncerPublicKey          *string
	encryptedPushRegistration *string
}

// Whether the token was generated by FCM or APNs.
func (r RegisterPushDeviceRequest) TokenKind(tokenKind string) RegisterPushDeviceRequest {
	r.tokenKind = &tokenKind
	return r
}

// A random integer generated by the client that will be included in mobile push notifications along with encrypted payloads to identify pushes as being related to this registration. Must be unique in the client&#39;s table of accounts.
func (r RegisterPushDeviceRequest) PushAccountId(pushAccountId int64) RegisterPushDeviceRequest {
	r.pushAccountId = &pushAccountId
	return r
}

// Public part of the asymmetric key pair generated by the client using NaCl (the Networking and Cryptography Library), encoded using a RFC 4648 standard base64 encoder.  It is used to encrypt notifications for delivery to the client.
func (r RegisterPushDeviceRequest) PushPublicKey(pushPublicKey string) RegisterPushDeviceRequest {
	r.pushPublicKey = &pushPublicKey
	return r
}

// Which of the bouncer&#39;s public keys the client used to encrypt the &#x60;PushRegistration&#x60; dictionary.  When the bouncer rotates the key, a new asymmetric key pair is created, and the new public key is baked into a new client release. Because the bouncer routinely rotates key, this field clarifies which public key the client is using.
func (r RegisterPushDeviceRequest) BouncerPublicKey(bouncerPublicKey string) RegisterPushDeviceRequest {
	r.bouncerPublicKey = &bouncerPublicKey
	return r
}

// Ciphertext generated by encrypting a &#x60;PushRegistration&#x60; dictionary using the &#x60;bouncer_public_key&#x60;, encoded using a RFC 4648 standard base64 encoder.  The &#x60;PushRegistration&#x60; dictionary contains the fields &#x60;token&#x60;, &#x60;token_kind&#x60;, &#x60;timestamp&#x60;, and (for iOS devices) &#x60;ios_app_id&#x60;. The dictionary is JSON-encoded before encryption.
func (r RegisterPushDeviceRequest) EncryptedPushRegistration(encryptedPushRegistration string) RegisterPushDeviceRequest {
	r.encryptedPushRegistration = &encryptedPushRegistration
	return r
}

func (r RegisterPushDeviceRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RegisterPushDeviceExecute(r)
}

/*
RegisterPushDevice Register E2EE push device

Register a device to receive end-to-end encrypted mobile push notifications.

**Changes**: New in Zulip 11.0 (feature level 406).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegisterPushDeviceRequest
*/
func (c *simpleClient) RegisterPushDevice(ctx context.Context) RegisterPushDeviceRequest {
	return RegisterPushDeviceRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) RegisterPushDeviceExecute(r RegisterPushDeviceRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mobile_push/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tokenKind == nil {
		return localVarReturnValue, nil, reportError("tokenKind is required and must be specified")
	}
	if r.pushAccountId == nil {
		return localVarReturnValue, nil, reportError("pushAccountId is required and must be specified")
	}
	if r.pushPublicKey == nil {
		return localVarReturnValue, nil, reportError("pushPublicKey is required and must be specified")
	}
	if r.bouncerPublicKey == nil {
		return localVarReturnValue, nil, reportError("bouncerPublicKey is required and must be specified")
	}
	if r.encryptedPushRegistration == nil {
		return localVarReturnValue, nil, reportError("encryptedPushRegistration is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "token_kind", r.tokenKind, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "push_account_id", r.pushAccountId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "push_public_key", r.pushPublicKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "bouncer_public_key", r.bouncerPublicKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "encrypted_push_registration", r.encryptedPushRegistration, "", "")
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TestNotifyRequest struct {
	ctx        context.Context
	ApiService MobileAPI
	token      *string
}

// The push token for the device to which to send the test notification.  If this parameter is not submitted, the test notification will be sent to all of the user&#39;s devices registered on the server.  A mobile client should pass this parameter, to avoid triggering a test notification for other clients.
func (r TestNotifyRequest) Token(token string) TestNotifyRequest {
	r.token = &token
	return r
}

func (r TestNotifyRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.TestNotifyExecute(r)
}

/*
TestNotify Send a test notification to mobile device(s)

Trigger sending a test push notification to the user's
selected mobile device or all of their mobile devices.

**Changes**: Starting with Zulip 8.0 (feature level 234), test
notifications sent via this endpoint use `test` rather than
`test-by-device-token` in the `event` field. Also, as of this
feature level, all mobile push notifications now include a
`realm_name` field.

New in Zulip 8.0 (feature level 217).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TestNotifyRequest
*/
func (c *simpleClient) TestNotify(ctx context.Context) TestNotifyRequest {
	return TestNotifyRequest{
		ApiService: c,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Response
func (c *simpleClient) TestNotifyExecute(r TestNotifyRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Response
	)

	localBasePath, err := c.ServerURL()
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mobile_push/test_notification"

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
	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "token", r.token, "", "")
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CodedError
			err = c.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = c.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	c.handleUnsupportedParameters(r.ctx, localVarReturnValue.IgnoredParametersUnsupported)

	return localVarReturnValue, localVarHTTPResponse, nil
}
