// Package union provides utilities for marshaling and unmarshaling union-like types in Go.
package union

import (
	"encoding/json"
	"reflect"

	"gopkg.in/validator.v2"

	strictdecoder "github.com/tum-zulip/go-zulip/zulip/internal/strict_decoder"
)

type Union interface {
	// Union is a marker interface
}

func Marshal[T Union](v T) ([]byte, error) {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Struct {
		return nil, &json.UnsupportedTypeError{Type: typ}
	}

	val := reflect.ValueOf(v)
	match := 0
	fieldIndex := 0
	for i := range val.NumField() {
		field := val.Field(i)
		if field.Type().Kind() != reflect.Pointer && field.Type().Kind() != reflect.Slice &&
			field.Type().Kind() != reflect.Map {
			return nil, &json.UnsupportedTypeError{Type: typ}
		}
		if !field.IsNil() {
			match++
			fieldIndex = i
		}
	}
	if match > 1 { // more than 1 match
		return nil, &json.UnsupportedValueError{Value: val, Str: "more than one field is set"}
	} else if match == 1 {
		return json.Marshal(val.Field(fieldIndex).Interface())
	}

	return []byte("{}"), nil
}

func Unmarshal[T Union](data []byte, dst *T) error {
	var err error
	match := 0
	val := reflect.ValueOf(dst).Elem()
	typ := val.Type()
	// try to unmarshal data into each pointer in the struct
	for i := range val.NumField() {
		field := val.Field(i)
		if field.Type().Kind() != reflect.Pointer && field.Type().Kind() != reflect.Slice &&
			field.Type().Kind() != reflect.Map {
			return &json.UnsupportedTypeError{Type: typ}
		}
		err = strictdecoder.New(data).Decode(field.Addr().Interface())
		if err == nil {
			jsonField, _ := json.Marshal(field.Interface())
			if string(jsonField) != "{}" { // at least one field is set
				if err = validator.Validate(field.Interface()); err == nil {
					match++
					continue
				}
			}
		}
		field.Set(reflect.Zero(field.Type()))
	}

	if match == 1 {
		return nil
	}

	for i := range val.NumField() {
		field := val.Field(i)
		field.Set(reflect.Zero(field.Type()))
	}

	if match > 1 { // more than 1 match
		return &json.UnsupportedValueError{Value: val, Str: "more than one field is set"}
	}

	// no match
	return &json.UnsupportedValueError{Value: val, Str: "no fields are set"}
}
