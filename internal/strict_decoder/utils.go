// Package strictdecoder provides utilities for strict JSON decoding.
package strictdecoder

import (
	"bytes"
	"encoding/json"
	"unicode/utf8"
)

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// A wrapper for strict JSON decoding.
func New(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}
