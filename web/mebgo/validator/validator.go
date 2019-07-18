package validator

import (
	"fmt"
	"reflect"

	"github.com/musenwill/experimentgo/web/mebgo/util"
)

type FieldValidator func(data map[string]interface{}, field string) error

type FormValidator interface {
	GetValidators() map[string][]FieldValidator
}

// Apply to any field that required, exists and not nil (can be empty, eg "", [], {})
// If value not exists or nil, then use the specified default value.
// If don't want specify default value, then set nil as default.
func Required(defaultValue interface{}) FieldValidator {
	return FieldValidator(func(data map[string]interface{}, field string) error {
		value, exists := data[field]
		if nil == value || !exists {
			if nil == defaultValue {
				return fmt.Errorf("Field %s is required.\n", field)
			} else {
				data[field] = defaultValue
			}
		}
		return nil
	})
}

// Apply to limit the minimum length of string field
func MinLen(minimumLen int) FieldValidator {
	return FieldValidator(func(data map[string]interface{}, field string) error {
		value, exists := data[field]
		if nil == value || !exists {
			return nil
		}

		if reflect.TypeOf(value).Kind() != reflect.String {
			return fmt.Errorf("MinLen applied to a none string field %s.\n", field)
		}

		if len(value.(string)) < minimumLen {
			return fmt.Errorf("%s length should >= %d\n", field, minimumLen)
		}

		return nil
	})
}

// Apply to limit the max length of string field
func MaxLen(maxLen int) FieldValidator {
	return FieldValidator(func(data map[string]interface{}, field string) error {
		value, exists := data[field]
		if nil == value || !exists {
			return nil
		}

		if reflect.TypeOf(value).Kind() != reflect.String {
			return fmt.Errorf("MaxLen applied to a none string field %s.\n", field)
		}

		if len(value.(string)) > maxLen {
			return fmt.Errorf("%s length should <= %d\n", field, maxLen)
		}

		return nil
	})
}

// Apply to limit the minimum value of number field
func Min(minValue int) FieldValidator {
	return FieldValidator(func(data map[string]interface{}, field string) error {
		value, exists := data[field]
		if nil == value || !exists {
			return nil
		}

		val, err := util.ToNumber(value)
		if err != nil {
			return fmt.Errorf("Min applied to a none number field %s.\n", field)
		}

		if val < float64(minValue) {
			return fmt.Errorf("%s should >= %d\n", field, minValue)
		}

		return nil
	})
}

// Apply to limit the max value of number field
func Max(maxValue int) FieldValidator {
	return FieldValidator(func(data map[string]interface{}, field string) error {
		value, exists := data[field]
		if nil == value || !exists {
			return nil
		}

		val, err := util.ToNumber(value)
		if err != nil {
			return fmt.Errorf("Max applied to a none number field %s.\n", field)
		}

		if val > float64(maxValue) {
			return fmt.Errorf("%s should <= %d\n", field, maxValue)
		}

		return nil
	})
}
