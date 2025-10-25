package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unicode/utf8"
)

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
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

func UnmarshalStringEnum[T ~string](data []byte, obj *T, allowed []T) error {
	var u string
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	t := T(u)
	for _, valid := range allowed {
		if valid == t {
			*obj = valid
			return nil
		}
	}
	return &json.UnsupportedValueError{
		Value: reflect.ValueOf(u),
		Str:   "", // TODO
	}
}

func UnmarshalIntEnum[T ~int](data []byte, obj *T, allowed []T) error {
	var u int
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	t := T(u)
	for _, valid := range allowed {
		if valid == t {
			*obj = valid
			return nil
		}
	}
	return &json.UnsupportedValueError{
		Value: reflect.ValueOf(u),
		Str:   "", // TODO
	}
}
