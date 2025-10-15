package api

import (
	"encoding/json"
	"reflect"
)

type Union interface {
	// Union is a marker interface
}

func UnionMarshalJSON[T Union](v T) ([]byte, error) {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Struct {
		return nil, &json.UnsupportedTypeError{Type: typ}
	}

	val := reflect.ValueOf(v)
	match := 0
	fieldIndex := 0
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type().Kind() != reflect.Pointer {
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

func UnionUnmarshalJSON[T Union](data []byte, dst *T) error {
	var err error
	match := 0
	val := reflect.ValueOf(dst).Elem()
	typ := val.Type()
	// try to unmarshal data into each pointer in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type().Kind() != reflect.Pointer {
			return &json.UnsupportedTypeError{Type: typ}
		}
		err = newStrictDecoder(data).Decode(field.Addr().Interface())
		if err == nil {
			match++
		}
	}

	if match == 1 {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		field.Set(reflect.Zero(field.Type()))
	}

	if match > 1 { // more than 1 match
		return &json.UnsupportedValueError{Value: val, Str: "more than one field is set"}
	}

	// no match
	return &json.UnsupportedValueError{Value: val, Str: "no fields are set"}
}
