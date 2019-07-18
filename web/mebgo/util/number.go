package util

import (
	"errors"
	"math"
	"reflect"
)

func NumberTypes() []reflect.Kind {
	return []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
	}
}

func IsNumber(value interface{}) bool {
	if nil == value {
		return false
	}

	valueType := reflect.TypeOf(value).Kind()

	for _, t := range NumberTypes() {
		if t == valueType {
			return true
		}
	}

	return false
}

func ToNumber(value interface{}) (float64, error) {
	if !IsNumber(value) {
		return 0, errors.New("Value is not number type")
	}

	switch i := value.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint16:
		return float64(i), nil
	case uint8:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int8:
		return float64(i), nil
	case int:
		return float64(i), nil
	default:
		return math.NaN(), errors.New("getFloat: unknown value is of incompatible type")
	}
}
