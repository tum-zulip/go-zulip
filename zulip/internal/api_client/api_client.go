package api_client

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
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	. "github.com/tum-zulip/go-zulip/zulip/internal/apiutils"
	. "github.com/tum-zulip/go-zulip/zulip/internal/utils"
	. "github.com/tum-zulip/go-zulip/zulip/models"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

// Will not be returned if max-retires is set to zero
var ErrMaxRetriesReached = errors.New("hit max retires")

const (
	defaultClientName = "go-zulip/1.0.0"
	defaultAPISuffix  = "api"
	defaultAPIVersion = "v1"
)

const retryIndefinitely = -1

// Client manages communication with the Zulip REST API API v1.0.0
// In most cases there should be only one, shared, Client.
type APIClient struct {
	Cfg *zuliprc.ZulipRC

	UserAgent       string
	HttpClient      *http.Client
	MaxRetries      int
	Logger          *slog.Logger
	InsecureWarning bool

	ClientName string

	ApiSuffix  string
	ApiVersion string
}

// NewSimpleClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(zuliprc *zuliprc.ZulipRC) (*APIClient, error) {

	client := &APIClient{
		Cfg:             zuliprc,
		ClientName:      defaultClientName,
		MaxRetries:      retryIndefinitely,
		InsecureWarning: true,
		ApiSuffix:       defaultAPISuffix,
		ApiVersion:      defaultAPIVersion,
		Logger:          slog.Default(),
	}

	// for _, option := range options {
	// 	option(client)
	// }

	httpClient, userAgent, err := buildHTTPClient(zuliprc, client.Logger, client.InsecureWarning)
	if err != nil {
		return nil, err
	}

	client.HttpClient = httpClient
	client.UserAgent = userAgent

	return client, nil
}

func (c *APIClient) ServerURL() (string, error) {
	if c.Cfg.Site != "" {
		return fmt.Sprintf("%s/%s/%s", c.Cfg.Site, c.ApiSuffix, c.ApiVersion), nil
	}
	return "", errors.New("base URL is not set")
}

func buildHTTPClient(params *zuliprc.ZulipRC, logger *slog.Logger, warnOnInsecureTLS bool) (*http.Client, string, error) {
	if params == nil {
		return nil, "", errors.New("invalid configuration: nil")
	}

	if logger == nil {
		logger = slog.Default()
	}

	clientName := defaultClientName
	userAgent := BuildUserAgent(clientName)

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

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetZulipRC() *zuliprc.ZulipRC {
	return c.Cfg
}

func (c *APIClient) CallAPI(ctx context.Context, req *http.Request, responseModel ResponseModel) (httpResp *http.Response, err error) {
	retryForever := c.MaxRetries == retryIndefinitely

	for i := 0; retryForever || i <= c.MaxRetries; i++ {
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
	slog.DebugContext(ctx, "hit max retires", "max-retires", c.MaxRetries)
	return httpResp, ErrMaxRetriesReached
}

func (c *APIClient) doHTTPCallAndParseResponse(ctx context.Context, req *http.Request, responseModel ResponseModel) (*http.Response, error) {
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
		return httpResp, NewAPIError(body, err.Error(), nil)
	}

	c.handleUnsupportedParameters(ctx, responseModel.GetIgnoredParametersUnsupported())

	return httpResp, nil
}

// callAPI do the request.
func (c *APIClient) doHTTPCall(ctx context.Context, request *http.Request) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	debug := c.Logger.Enabled(ctx, slog.LevelDebug)

	if debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		c.Logger.DebugContext(ctx, "HTTP Request", "dump", string(dump))
	}

	resp, err := c.HttpClient.Do(request)

	if debug && resp != nil {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			c.Logger.ErrorContext(ctx, "failed to dump http response to string", "error", err)
		}
		c.Logger.DebugContext(ctx, "HTTP Response", "dump", string(dump))
	}

	return resp, err
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
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

func (c *APIClient) handleUnsupportedParameters(ctx context.Context, parameters []string) {
	if len(parameters) > 0 {
		c.Logger.WarnContext(ctx, "Unsupported parameters were ignored", "parameters", parameters)
	}
}

func tryUnmarshalErrorModel[T any](data []byte) (*T, error) {
	var model T
	dec := NewStrictDecoder(data)
	err := dec.Decode(&model)
	if err != nil {
		return nil, err
	}
	if reflect.ValueOf(model).IsZero() {
		return nil, errors.New("no data")
	}
	return &model, nil
}

func (c *APIClient) handleErrorResponse(ctx context.Context, resp *http.Response) error {
	var model interface{}
	var err error

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewAPIError(nil, fmt.Errorf("failed to read response body: %w", err).Error(), nil)
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
		return NewAPIError(body, fmt.Sprintf("status %d: %s", resp.StatusCode, string(body)), nil)
	}

	return NewAPIError(body, fmt.Sprintf("status %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode)), model)
}

func normalizePath(pathStr string) (string, error) {
	expanded := os.ExpandEnv(pathStr)
	if strings.HasPrefix(expanded, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		expanded = filepath.Join(home, strings.TrimPrefix(expanded, "~"))
	}
	return filepath.Abs(expanded)
}

type basicAuthRoundTripper struct {
	email     string
	apiKey    string
	userAgent string
	base      http.RoundTripper
}

func (t *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.SetBasicAuth(t.email, t.apiKey)
	if t.userAgent != "" {
		clone.Header.Set("User-Agent", t.userAgent)
	}
	return t.base.RoundTrip(clone)
}
