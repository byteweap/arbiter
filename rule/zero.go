// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrZero is returned when a value is not zero.
var (
	ErrZero = errors.New("value must be zero")
)

// ZeroRule is a validation rule that checks if a value is zero.
// It supports all basic types and complex types through the Zeroable interface.
//
// Example:
//
//	rule := Zero[int]()
//	err := rule.Validate(0)     // returns nil
//	err = rule.Validate(42)     // returns ErrZero
//
//	rule = Zero[string]()
//	err = rule.Validate("")     // returns nil
//	err = rule.Validate("hello") // returns ErrZero
type ZeroRule[T Zeroable] struct {
	e error
}

// Zero creates a new zero value validation rule.
// The rule can be used with any type that implements the Zeroable interface.
//
// Example:
//
//	// Create a rule for zero integers
//	countRule := Zero[int]().Err("Count must be zero")
//
//	// Create a rule for empty strings
//	nameRule := Zero[string]().Err("Name must be empty")
func Zero[T Zeroable]() *ZeroRule[T] {
	return &ZeroRule[T]{
		e: ErrZero,
	}
}

// Validate checks if the value is zero.
// For numbers, it checks if the value is 0.
// For strings, it checks if the string is empty.
// For booleans, it checks if the value is false.
// For pointers, it checks if the pointer is nil.
// For other types, it uses reflection to determine if the value is zero.
//
// Example:
//
//	rule := Zero[int]()
//	if err := rule.Validate(0); err != nil {
//	    // Handle validation error
//	}
func (r *ZeroRule[T]) Validate(value T) error {
	if !isZero(value) {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v must be zero", value)
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Zero[int]().Errf("Value must be zero")
//	err := rule.Validate(42)  // returns error with message "Value must be zero"
func (r *ZeroRule[T]) Errf(format string, args ...any) *ZeroRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// isZero determines if a value is considered zero.
// It handles various types differently:
// - Numbers: checks if the value is 0
// - Strings: checks if the string is empty
// - Booleans: checks if the value is false
// - Pointers: checks if the pointer is nil
// - Arrays/Slices/Maps: checks if the length is 0
// - Structs: checks if all fields are zero
// - Other types: uses reflection to compare with zero value
//
// Example:
//
//	isZero(0)           // returns true
//	isZero("")          // returns true
//	isZero(false)       // returns true
//	isZero(nil)         // returns true
//	isZero([]int{})     // returns true
//	isZero(struct{}{})  // returns true
func isZero[T Zeroable](value T) bool {
	// Handle basic types
	switch v := any(value).(type) {
	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case uint8:
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return !v
	case complex64:
		return v == 0
	case complex128:
		return v == 0
	case nil:
		return true
	default:
		// Use reflection for other types
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr, reflect.Interface:
			return rv.IsNil()
		case reflect.String:
			return rv.Len() == 0
		case reflect.Array, reflect.Slice, reflect.Map:
			return rv.Len() == 0
		case reflect.Struct:
			// For structs, check if all fields are zero
			for i := 0; i < rv.NumField(); i++ {
				if !isZero(rv.Field(i).Interface()) {
					return false
				}
			}
			return true
		default:
			// For other types, try to compare with zero value
			zero := reflect.Zero(rv.Type()).Interface()
			return reflect.DeepEqual(value, zero)
		}
	}
}
