package zulip

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

const retryIndefinitely = -1

// Will not be returned if max-retires is set to zero
var ErrMaxRetriesReached = errors.New("hit max retires")

// Client manages communication with the Zulip REST API API v1.0.0
// In most cases there should be only one, shared, Client.
type simpleClient struct {
	cfg *ZulipRC

	defaultHeader   map[string]string
	userAgent       string
	httpClient      *http.Client
	maxRetries      int
	logger          *slog.Logger
	insecureWarning bool

	clientName string

	apiSuffix  string
	apiVersion string
}

var _ Client = &simpleClient{}

type Option func(*simpleClient)

// NewSimpleClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewSimpleClient(zuliprc *ZulipRC, options ...Option) (*simpleClient, error) {
	if zuliprc == nil {
		return nil, errors.New("invalid configuration: nil")
	}

	if zuliprc.APIKey == "" {
		return nil, ConfigMissingError{message: "api_key not specified"}
	}

	if zuliprc.Email == "" {
		return nil, ConfigMissingError{message: "email not specified"}
	}

	if zuliprc.Site == "" {
		return nil, ErrMissingURLError
	}

	if zuliprc.ClientCert == "" && zuliprc.ClientCertKey != "" {
		return nil, ConfigMissingError{message: fmt.Sprintf("client cert key '%s' specified, but no client cert public part provided", zuliprc.ClientCertKey)}
	}

	if zuliprc.ClientCert != "" {
		if _, err := os.Stat(zuliprc.ClientCert); err != nil {
			return nil, ConfigMissingError{message: fmt.Sprintf("client cert '%s' does not exist", zuliprc.ClientCert)}
		}
		if zuliprc.ClientCertKey != "" {
			if _, err := os.Stat(zuliprc.ClientCertKey); err != nil {
				return nil, ConfigMissingError{message: fmt.Sprintf("client cert key '%s' does not exist", zuliprc.ClientCertKey)}
			}
		}
	}

	if zuliprc.CertBundle != "" {
		if _, err := os.Stat(zuliprc.CertBundle); err != nil {
			return nil, ConfigMissingError{message: fmt.Sprintf("tls bundle '%s' does not exist", zuliprc.CertBundle)}
		}
	}

	client := &simpleClient{
		cfg:             zuliprc,
		defaultHeader:   make(map[string]string),
		clientName:      defaultClientName,
		maxRetries:      retryIndefinitely,
		insecureWarning: true,
		apiSuffix:       defaultAPISuffix,
		apiVersion:      defaultAPIVersion,
		logger:          slog.Default(),
	}

	for _, option := range options {
		option(client)
	}

	httpClient, userAgent, err := buildHTTPClient(zuliprc, client.logger, client.insecureWarning)
	if err != nil {
		return nil, err
	}

	client.httpClient = httpClient
	client.userAgent = userAgent

	return client, nil
}

func WithAPISuffix(suffix string) Option {
	return func(c *simpleClient) {
		c.apiSuffix = suffix
	}
}

func WithAPIVersion(version string) Option {
	return func(c *simpleClient) {
		c.apiVersion = version
	}
}

func WithDefaultHeader(key string, value string) Option {
	return func(c *simpleClient) {
		if c.defaultHeader == nil {
			c.defaultHeader = make(map[string]string)
		}
		c.defaultHeader[key] = value
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *simpleClient) {
		if logger == nil {
			c.logger = slog.Default()
			return
		}
		c.logger = logger
	}
}

func WithMaxRetries(maxRetries int) Option {
	return func(c *simpleClient) {
		c.maxRetries = maxRetries
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *simpleClient) {
		c.httpClient = httpClient
	}
}

func WithClientName(name string) Option {
	return func(c *simpleClient) {
		c.clientName = name
	}
}

func SkipWarnOnInsecureTLS() Option {
	return func(c *simpleClient) {
		c.insecureWarning = false
	}
}

func (c *simpleClient) ServerURL() (string, error) {
	if c.cfg.Site != "" {
		return fmt.Sprintf("%s/%s/%s", c.cfg.Site, c.apiSuffix, c.apiVersion), nil
	}
	return "", errors.New("base URL is not set")
}

func buildHTTPClient(params *ZulipRC, logger *slog.Logger, warnOnInsecureTLS bool) (*http.Client, string, error) {
	if params == nil {
		return nil, "", errors.New("invalid configuration: nil")
	}

	if logger == nil {
		logger = slog.Default()
	}

	clientName := defaultClientName
	userAgent := buildUserAgent(clientName)

	var transport *http.Transport
	if def, ok := http.DefaultTransport.(*http.Transport); ok {
		transport = def.Clone()
	} else {
		transport = &http.Transport{}
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	} else {
		transport.TLSClientConfig = transport.TLSClientConfig.Clone()
	}

	if params.Insecure != nil && *params.Insecure {
		if warnOnInsecureTLS {
			logger.Warn("Insecure mode enabled. The server's SSL/TLS certificate will not be validated, making the HTTPS connection potentially insecure")
		}
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	if params.CertBundle != "" {
		bundlePath, err := normalizePath(params.CertBundle)
		if err != nil {
			return nil, "", err
		}

		data, err := os.ReadFile(bundlePath)
		if err != nil {
			return nil, "", err
		}

		var pool *x509.CertPool
		if transport.TLSClientConfig.RootCAs != nil {
			pool = transport.TLSClientConfig.RootCAs
		} else {
			pool, err = x509.SystemCertPool()
			if err != nil {
				pool = x509.NewCertPool()
			}
		}

		if ok := pool.AppendCertsFromPEM(data); !ok {
			return nil, "", fmt.Errorf("failed to parse certificate bundle %s", bundlePath)
		}
		transport.TLSClientConfig.RootCAs = pool
	}

	if params.ClientCert != "" {
		certPath, err := normalizePath(params.ClientCert)
		if err != nil {
			return nil, "", err
		}

		keyPath := params.ClientCertKey
		if keyPath == "" {
			keyPath = params.ClientCert
		}

		keyPath, err = normalizePath(keyPath)
		if err != nil {
			return nil, "", err
		}

		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to load client certificate: %w", err)
		}
		transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
	}

	baseClient := http.DefaultClient
	client := &http.Client{
		Timeout:       baseClient.Timeout,
		Transport:     &basicAuthRoundTripper{email: params.Email, apiKey: params.APIKey, userAgent: userAgent, base: transport},
		CheckRedirect: baseClient.CheckRedirect,
		Jar:           baseClient.Jar,
	}

	return client, userAgent, nil
}

func (c *simpleClient) callAPI(ctx context.Context, req *http.Request, responseModel responseModel) (httpResp *http.Response, err error) {
	retryForever := c.maxRetries == retryIndefinitely

	for i := 0; retryForever || i <= c.maxRetries; i++ {
		httpResp, err = c.doHTTPCallAndParseResponse(ctx, req, responseModel)
		if httpResp != nil && httpResp.StatusCode == http.StatusTooManyRequests {
			if apiErr, ok := err.(APIError); ok {
				if rateLimitedError, ok := apiErr.Model().(RateLimitedError); ok {
					slog.DebugContext(ctx, "hit API rate-limit", "retry-after", rateLimitedError.RetryAfter)
					time.Sleep(rateLimitedError.RetryAfter)
					continue
				}
			}
		}
		return
	}
	slog.DebugContext(ctx, "hit max retires", "max-retires", c.maxRetries)
	return httpResp, ErrMaxRetriesReached
}

func (c *simpleClient) doHTTPCallAndParseResponse(ctx context.Context, req *http.Request, responseModel responseModel) (*http.Response, error) {
	if responseModel == nil {
		return nil, fmt.Errorf("response model cannot be nil")
	}
	httpResp, err := c.doHTTPCall(ctx, req)
	if err != nil || httpResp == nil {
		return httpResp, err
	}

	body, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	httpResp.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return httpResp, err
	}

	if httpResp.StatusCode >= 300 {
		return httpResp, c.handleErrorResponse(ctx, httpResp)
	}

	err = c.decode(responseModel, body, httpResp.Header.Get("Content-Type"))
	if err != nil {
		newErr := &APIError{
			body:  body,
			error: err.Error(),
		}
		return httpResp, newErr
	}

	c.handleUnsupportedParameters(ctx, responseModel.getIgnoredParametersUnsupported())

	return httpResp, nil
}

// callAPI do the request.
func (c *simpleClient) doHTTPCall(ctx context.Context, request *http.Request) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	debug := c.logger.Enabled(ctx, slog.LevelDebug)

	if debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		c.logger.DebugContext(ctx, "HTTP Request", "dump", string(dump))
	}

	resp, err := c.httpClient.Do(request)

	if debug && resp != nil {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.logger.ErrorContext(ctx, "failed to dump http response to string", "error", err)
		}
		c.logger.DebugContext(ctx, "HTTP Response", "dump", string(dump))
	}

	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *simpleClient) GetZulipRC() *ZulipRC {
	return c.cfg
}

type formFile struct {
	fileBytes    []byte
	fileName     string
	formFileName string
}

// prepareRequest build the request
func (c *simpleClient) prepareRequest(
	ctx context.Context,
	endpoint string, method string,
	headerParams map[string]string,
	query url.Values,
	form url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	basePath, err := c.ServerURL()
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	path := basePath + endpoint

	var body *bytes.Buffer

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(form) > 0 || (len(formFiles) > 0) {
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range form {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				headers := make(textproto.MIMEHeader)
				headers.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, formFile.formFileName, filepath.Base(formFile.fileName)))
				if detected := http.DetectContentType(formFile.fileBytes); detected != "" {
					headers.Set("Content-Type", detected)
				}

				part, err := w.CreatePart(headers)
				if err != nil {
					return nil, err
				}
				if _, err = part.Write(formFile.fileBytes); err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(form) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(form.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	q := url.Query()
	for k, v := range query {
		for _, iv := range v {
			q.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryplit.ReplaceAllStringFunc(q.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.userAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

	}

	for header, value := range c.defaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *simpleClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

func (c *simpleClient) handleUnsupportedParameters(ctx context.Context, parameters []string) {
	if len(parameters) > 0 {
		c.logger.WarnContext(ctx, "Unsupported parameters were ignored", "parameters", parameters)
	}
}

func tryUnmarshalErrorModel[T any](data []byte) (*T, error) {
	var model T
	dec := newStrictDecoder(data)
	err := dec.Decode(&model)
	if err != nil {
		return nil, err
	}
	if reflect.ValueOf(model).IsZero() {
		return nil, errors.New("no data")
	}
	return &model, nil
}

func (c *simpleClient) handleErrorResponse(ctx context.Context, resp *http.Response) error {
	var model interface{}
	var err error

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIError{
			body:  body,
			error: fmt.Errorf("failed to read response body: %w", err).Error(),
			model: nil,
		}
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		model, err = tryUnmarshalErrorModel[RateLimitedError](body)
	case http.StatusBadRequest:
		model, err = tryUnmarshalErrorModel[BadEventQueueIdError](body)
		if err == nil {
			break
		}

		model, err = tryUnmarshalErrorModel[IncompatibleParametersError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[MissingArgumentError](body)
		if err == nil {
			break
		}
	case http.StatusNotFound:
		model, err = tryUnmarshalErrorModel[NonExistingChannelIdError](body)
		if err == nil {
			break
		}

		model, err = tryUnmarshalErrorModel[InvitationFailedError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[NonExistingChannelNameError](body)
		if err == nil {
			break
		}
	}

	if model == nil {
		model, err = tryUnmarshalErrorModel[CodedError](body)
	}

	if err != nil {
		slog.WarnContext(ctx, "API returned an unknown error response", "status", resp.StatusCode, "body", string(body), "model", model)
		return &APIError{
			body:  body,
			error: fmt.Sprintf("status %d: %s", resp.StatusCode, string(body)),
			model: nil,
		}
	}

	return &APIError{
		body:  body,
		error: fmt.Sprintf("status %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode)),
		model: model,
	}
}
