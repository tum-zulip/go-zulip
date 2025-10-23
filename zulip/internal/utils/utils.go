package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// cacheExpires helper function to determine remaining time before repeating a request.
func cacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
}

func BuildUserAgent(clientName string) string {
	vendor, version := detectPlatform()
	return fmt.Sprintf("%s (%s; %s)", clientName, vendor, version)
}

func detectPlatform() (string, string) {
	switch runtime.GOOS {
	case "linux":
		if name, version, err := readOSRelease(); err == nil {
			return name, version
		}
		return "Linux", runtime.GOARCH
	case "windows":
		return "Windows", runtime.GOARCH
	case "darwin":
		return "Darwin", runtime.GOARCH
	default:
		return runtime.GOOS, runtime.GOARCH
	}
}

func readOSRelease() (string, string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", "", err
	}

	var name, version string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") && name == "" {
			name = trimOSReleaseValue(line[5:])
		}
		if strings.HasPrefix(line, "VERSION_Id=") && version == "" {
			version = trimOSReleaseValue(line[11:])
		}
	}

	if name == "" {
		name = "Linux"
	}
	if version == "" {
		version = runtime.GOARCH
	}

	return name, version, nil
}

func trimOSReleaseValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "\"")
	value = strings.TrimSuffix(value, "\"")
	return value
}

// A wrapper for strict JSON decoding
func NewStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}
