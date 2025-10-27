// Package enum provides utilities for unmarshaling enum-like types in Go.
package enum

import (
	"encoding/json"
	"reflect"
)

func UnmarshalString[T ~string](data []byte, obj *T, allowed []T) error {
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

func UnmarshalInt[T ~int](data []byte, obj *T, allowed []T) error {
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
