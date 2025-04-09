// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating if a value is nil or not nil.
package rule

import (
	"errors"
	"fmt"
	"reflect"
)

// Error variables for nil validation
var (
	// ErrNil is returned when a value must be nil but is not
	ErrNil = errors.New("must be nil")
	// ErrNotNil is returned when a value must not be nil but is
	ErrNotNil = errors.New("must not be nil")
)

// Predefined rules for common nil validation scenarios
var (
	// Nil is a predefined rule that validates a value must be nil
	Nil = &NilRule[any]{e: ErrNil}
	// NotNil is a predefined rule that validates a value must not be nil
	NotNil = &NilRule[any]{e: ErrNotNil, not: true}
)

// NilRule validates if a value is nil or not nil.
// This rule works with various types including pointers, interfaces, maps, functions,
// arrays, channels, and slices.
//
// Example:
//
//	// Using predefined rules
//	err := Nil.Validate(nil)           // returns nil (value is nil)
//	err = Nil.Validate("not nil")      // returns error (value is not nil)
//
//	err = NotNil.Validate("not nil")   // returns nil (value is not nil)
//	err = NotNil.Validate(nil)         // returns error (value is nil)
//
//	// Using custom rules
//	var ptr *string
//	err = Nil.Validate(ptr)            // returns nil (pointer is nil)
//	err = NotNil.Validate(ptr)         // returns error (pointer is nil)
type NilRule[T any] struct {
	not bool
	e   error
}

// Validate checks if the value is nil or not nil, depending on the rule type.
// Returns nil if the validation passes, or an error otherwise.
//
// Example:
//
//	// Check if a pointer is nil
//	var ptr *string
//	err := Nil.Validate(ptr)  // returns nil (ptr is nil)
//
//	// Check if a slice is not nil
//	slice := []int{1, 2, 3}
//	err = NotNil.Validate(slice)  // returns nil (slice is not nil)
//
//	// Check if a map is nil
//	var m map[string]int
//	err = Nil.Validate(m)  // returns nil (m is nil)
//
//	// Check if a channel is not nil
//	ch := make(chan int)
//	err = NotNil.Validate(ch)  // returns nil (ch is not nil)
func (r *NilRule[T]) Validate(value T) error {
	var isNil bool
	switch v := any(value).(type) {
	case nil:
		isNil = true
	case string:
		isNil = false
	case *string:
		isNil = v == nil
	case *int:
		isNil = v == nil
	case int:
		isNil = false
	case *int8:
		isNil = v == nil
	case int8:
		isNil = false
	case *int16:
		isNil = v == nil
	case int16:
		isNil = false
	case *int32:
		isNil = v == nil
	case int32:
		isNil = false
	case *int64:
		isNil = v == nil
	case int64:
		isNil = false
	case *uint:
		isNil = v == nil
	case uint:
		isNil = false
	case *uint8:
		isNil = v == nil
	case uint8:
		isNil = false
	case *uint16:
		isNil = v == nil
	case uint16:
		isNil = false
	case *uint32:
		isNil = v == nil
	case uint32:
		isNil = false
	case *uint64:
		isNil = v == nil
	case uint64:
		isNil = false
	case *float32:
		isNil = v == nil
	case float32:
		isNil = false
	case *float64:
		isNil = v == nil
	case *bool:
		isNil = v == nil
	case bool:
		isNil = false
	default:
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Func, reflect.Array, reflect.Chan, reflect.Slice:
			isNil = rv.IsNil()
		default:
			return fmt.Errorf("unsupported type: %v", rv.Kind())
		}
	}
	if (r.not && isNil) || (!r.not && !isNil) {
		return r.e
	}
	return nil
}

// Errf sets a custom error message for nil validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := &NilRule[any]{not: true}.Errf("This field is required")
//	err := rule.Validate(nil)  // returns error with custom message
//
//	rule = &NilRule[any]{not: false}.Errf("This field must be empty")
//	err = rule.Validate("not empty")  // returns error with custom message
func (r *NilRule[T]) Errf(format string, args ...any) *NilRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
