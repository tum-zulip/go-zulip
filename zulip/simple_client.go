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
	"strconv"
	"strings"
	"time"
)

// Client manages communication with the Zulip REST API API v1.0.0
// In most cases there should be only one, shared, Client.
type simpleClient struct {
	cfg *ZulipRC

	defaultHeader map[string]string
	userAgent     string
	httpClient    *http.Client
	logger        *slog.Logger

	clientName    string
	retryOnErrors bool

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
		cfg:           zuliprc,
		defaultHeader: make(map[string]string),
		clientName:    defaultClientName,
		retryOnErrors: true,
		apiSuffix:     defaultAPISuffix,
		apiVersion:    defaultAPIVersion,
		logger:        slog.Default(),
	}

	for _, option := range options {
		option(client)
	}

	httpClient, userAgent, err := buildHTTPClient(zuliprc, client.logger)
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

func (c *simpleClient) ServerURL() (string, error) {
	if c.cfg.Site != "" {
		return fmt.Sprintf("%s/%s/%s", c.cfg.Site, c.apiSuffix, c.apiVersion), nil
	}
	return "", errors.New("base URL is not set")
}

func buildHTTPClient(params *ZulipRC, logger *slog.Logger) (*http.Client, string, error) {
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
		logger.Warn("Insecure mode enabled. The server's SSL/TLS certificate will not be validated, making the HTTPS connection potentially insecure")
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

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		if actualObj, ok := obj.(interface{ GetActualInstanceValue() interface{} }); ok {
			return fmt.Sprintf("%v", actualObj.GetActualInstanceValue())
		}

		return fmt.Sprintf("%v", obj)
	}
	var param, ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			if marshaler, ok := obj.(json.Marshaler); ok {
				if data, err := marshaler.MarshalJSON(); err == nil {
					value = string(data)
					break
				}
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			if v.Type().Elem().Kind() == reflect.Uint8 {
				value = string(v.Bytes())
				break
			}
			if data, err := json.Marshal(obj); err == nil {
				value = string(data)
				break
			}
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				var keyPrefixForCollectionType = keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			if data, err := json.Marshal(obj); err == nil {
				value = string(data)
				break
			}
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			if marshaler, ok := obj.(json.Marshaler); ok {
				if data, err := marshaler.MarshalJSON(); err == nil {
					value = string(data)
					break
				}
			}
			if data, err := json.Marshal(obj); err == nil {
				value = string(data)
				break
			}
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
		break
	case map[string]string:
		valuesMap[keyPrefix] = value
		break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *simpleClient) callAPI(ctx context.Context, request *http.Request) (*http.Response, error) {
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
	if err != nil {
		return resp, err
	}

	if debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
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
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
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

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
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
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
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
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
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
