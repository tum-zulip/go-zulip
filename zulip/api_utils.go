package zulip

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// same as addParam, but does nothing if obj is nil
func addOptionalParam(values url.Values, key string, obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.IsNil() {
		return
	}
	addParam(values, key, obj)
}

func addOptionalJSONParam(form url.Values, key string, value any) error {
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

func addOptionalCSVParam(values url.Values, key string, obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.IsNil() {
		return
	}
	addParamImpl(values, key, obj, true)
}

// addParam adds the provided object to the request header or url query
// supporting deep object syntax
func addParam(values url.Values, key string, obj interface{}) {
	addParamImpl(values, key, obj, false)
}

// addParam adds the provided object to the request header or url query
// supporting deep object syntax
func addCSVParam(values url.Values, key string, obj interface{}) {
	addParamImpl(values, key, obj, true)
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
