// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
)

// ErrRequired is returned when a required value is empty or zero.
var ErrRequired = errors.New("required")

// RequiredRule is a validation rule that checks if a value is required (non-empty/non-zero).
// It supports both value types and their pointer variants.
//
// Example:
//
//	rule := Required[string]()
//	err := rule.Validate("hello")  // returns nil
//	err = rule.Validate("")       // returns ErrRequired
//
//	rule = Required[int]()
//	err = rule.Validate(42)       // returns nil
//	err = rule.Validate(0)        // returns ErrRequired
type RequiredRule[T RequiredType] struct {
	e error
}

// Required creates a new required validation rule.
// The rule can be used with any type that implements the RequiredType interface.
//
// Example:
//
//	// Create a rule for required string
//	nameRule := Required[string]().Err("Name is required")
//
//	// Create a rule for required integer
//	ageRule := Required[int]().Err("Age is required")
func Required[T RequiredType]() *RequiredRule[T] {
	return &RequiredRule[T]{}
}

// Validate checks if the value is required (non-empty/non-zero).
// For strings, it checks if the string is non-empty.
// For numbers, it checks if the value is non-zero.
// For pointers, it checks if the pointer is non-nil and the value is non-empty/non-zero.
//
// Example:
//
//	rule := Required[string]()
//	if err := rule.Validate(""); err != nil {
//	    // Handle validation error
//	}
func (r *RequiredRule[T]) Validate(value T) error {
	var ok bool
	switch v := any(value).(type) {
	case string:
		ok = v != ""
	case *string:
		ok = v != nil && *v != ""
	case int:
		ok = v != 0
	case *int:
		ok = v != nil && *v != 0
	case int8:
		ok = v != 0
	case *int8:
		ok = v != nil && *v != 0
	case int16:
		ok = v != 0
	case *int16:
		ok = v != nil && *v != 0
	case int32:
		ok = v != 0
	case *int32:
		ok = v != nil && *v != 0
	case int64:
		ok = v != 0
	case *int64:
		ok = v != nil && *v != 0
	case uint:
		ok = v != 0
	case *uint:
		ok = v != nil && *v != 0
	case uint8:
		ok = v != 0
	case *uint8:
		ok = v != nil && *v != 0
	case uint16:
		ok = v != 0
	case *uint16:
		ok = v != nil && *v != 0
	case uint32:
		ok = v != 0
	case *uint32:
		ok = v != nil && *v != 0
	case uint64:
		ok = v != 0
	case *uint64:
		ok = v != nil && *v != 0
	case float32:
		ok = v != 0
	case *float32:
		ok = v != nil && *v != 0
	case float64:
		ok = v != 0
	case *float64:
		ok = v != nil && *v != 0
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	if !ok {
		if r.e != nil {
			return r.e
		}
		return ErrRequired
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
// If msg is empty, the default error message 'required' will be used.
//
// Example:
//
//	rule := Required[string]().Errf("This field is required")
//	err := rule.Validate("")  // returns error with message "This field is required"
func (r *RequiredRule[T]) Errf(format string, args ...any) *RequiredRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
