package apiutils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/internal/clients"
	"github.com/tum-zulip/go-zulip/zulip/zuliprc"
)

var (
	queryplit    = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape = strings.NewReplacer("%5B", "[", "%5D", "]")
)

type FormFile struct {
	FileBytes    []byte
	FileName     string
	FormFileName string
}

func IdToString(id int64) string {
	return strconv.FormatInt(id, 10)
}

// prepareRequest build the request
func PrepareRequest(
	ctx context.Context,
	c clients.Client,
	endpoint string, method string,
	headerParams map[string]string,
	query url.Values,
	form url.Values,
	formFiles []FormFile) (localVarRequest *http.Request, err error) {

	basePath, err := c.ServerURL()
	if err != nil {
		return nil, zulip.NewAPIError(nil, err.Error(), nil)
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
					err = AddFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.FileBytes) > 0 && formFile.FileName != "" {
				headers := make(textproto.MIMEHeader)
				headers.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, formFile.FormFileName, filepath.Base(formFile.FileName)))
				if detected := http.DetectContentType(formFile.FileBytes); detected != "" {
					headers.Set("Content-Type", detected)
				}

				part, err := w.CreatePart(headers)
				if err != nil {
					return nil, err
				}
				if _, err = part.Write(formFile.FileBytes); err != nil {
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
	localVarRequest.Header.Add("User-Agent", c.GetUserAgent())

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(zuliprc.ContextBasicAuth).(zuliprc.BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

	}

	return localVarRequest, nil
}

// AddOptionalParam adds the provided object to url.Values, but does nothing if obj is nil.
// This is an internal helper for API request builders.
func AddOptionalParam(values url.Values, key string, obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.IsNil() {
		return
	}
	AddParam(values, key, obj)
}

// AddOptionalJSONParam adds the provided object as JSON to url.Values, but does nothing if value is nil.
// This is an internal helper for API request builders.
func AddOptionalJSONParam(form url.Values, key string, value any) error {
	v := reflect.ValueOf(value)
	if v.IsNil() {
		return nil
	}

	jsonBuf, err := json.Marshal(value)
	if err != nil {
		return err
	}
	form.Add(key, string(jsonBuf))
	return nil
}

// AddOptionalCSVParam adds the provided object to url.Values in CSV format, but does nothing if obj is nil.
// This is an internal helper for API request builders.
func AddOptionalCSVParam(values url.Values, key string, obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.IsNil() {
		return
	}
	addParamImpl(values, key, obj, true)
}

// AddParam adds the provided object to the request header or url query
// supporting deep object syntax.
// This is an internal helper for API request builders.
func AddParam(values url.Values, key string, obj interface{}) {
	addParamImpl(values, key, obj, false)
}

// AddCSVParam adds the provided object to the request header or url query in CSV format
// supporting deep object syntax.
// This is an internal helper for API request builders.
func AddCSVParam(values url.Values, key string, obj interface{}) {
	addParamImpl(values, key, obj, true)
}

// Add a file to the multipart request
func AddFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

func addParamImpl(values url.Values, key string, obj interface{}, csv bool) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(time.Time); ok {
				addParamImpl(values, key, t.Format(time.RFC3339Nano), csv)
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
				addParamImpl(values, key, arrayValue.Interface(), csv)
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
				addParamImpl(values, fmt.Sprintf("%s[%s]", key, k.String()), v.Interface(), csv)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			addParamImpl(values, key, v.Elem().Interface(), csv)
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

	if csv && values.Get(key) != "" {
		values.Set(key, values.Get(key)+","+value)
	} else {
		values.Add(key, value)
	}
}
