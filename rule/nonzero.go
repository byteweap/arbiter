// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating that a value is non-zero.
package rule

import (
	"errors"
	"fmt"
	"reflect"
)

// Error variable for non-zero validation
var (
	// ErrNonZero is returned when a value must be non-zero but is zero
	ErrNonZero = errors.New("value must be non-zero")
)

// NonZeroRule validates that a value is non-zero.
// This rule works with various types including numbers, strings, booleans,
// slices, maps, arrays, and structs.
//
// Example:
//
//	// Create a non-zero rule
//	rule := NonZero[int]()
//
//	// Validate numbers
//	err := rule.Validate(42)     // returns nil (non-zero)
//	err = rule.Validate(0)       // returns error (zero)
//
//	// Validate strings
//	err = rule.Validate("hello") // returns nil (non-empty)
//	err = rule.Validate("")      // returns error (empty)
//
//	// Validate slices
//	err = rule.Validate([]int{1, 2, 3}) // returns nil (non-empty)
//	err = rule.Validate([]int{})        // returns error (empty)
//
//	// Validate structs
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	err = rule.Validate(Person{Name: "John", Age: 30}) // returns nil (non-zero)
//	err = rule.Validate(Person{})                      // returns error (zero)
type NonZeroRule[T any] struct {
	e error
}

// NonZero creates a new non-zero validation rule.
// This function returns a rule that can be used to validate that a value
// is not equal to its type's zero value.
//
// Example:
//
//	// Create rules for different types
//	intRule := NonZero[int]()
//	strRule := NonZero[string]()
//	sliceRule := NonZero[[]int]()
//
//	// Use the rules
//	err := intRule.Validate(42)
//	err = strRule.Validate("hello")
//	err = sliceRule.Validate([]int{1, 2, 3})
func NonZero[T any]() *NonZeroRule[T] {
	return &NonZeroRule[T]{e: ErrNonZero}
}

// Validate checks if the value is non-zero.
// Returns nil if the value is non-zero, or an error otherwise.
//
// The validation rules for different types are:
// - Numbers: Must not be 0
// - Strings: Must not be empty
// - Booleans: Must be true
// - Slices/Maps: Must not be nil and must have length > 0
// - Arrays: Must have length > 0
// - Structs: At least one field must be non-zero
//
// Example:
//
//	rule := NonZero[int]()
//
//	// Numbers
//	err := rule.Validate(42)     // returns nil
//	err = rule.Validate(0)       // returns error
//
//	// Strings
//	err = rule.Validate("hello") // returns nil
//	err = rule.Validate("")      // returns error
//
//	// Slices
//	err = rule.Validate([]int{1}) // returns nil
//	err = rule.Validate([]int{})  // returns error
//
//	// Maps
//	err = rule.Validate(map[string]int{"a": 1}) // returns nil
//	err = rule.Validate(map[string]int{})       // returns error
//
//	// Structs
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	err = rule.Validate(Person{Name: "John"}) // returns nil
//	err = rule.Validate(Person{})             // returns error
func (r *NonZeroRule[T]) Validate(value T) error {

	// Get reflection value
	v := reflect.ValueOf(value)

	// Handle pointer types
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ErrNonZero
		}
		v = v.Elem()
	}

	// Check zero value based on type
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return ErrNonZero
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			return ErrNonZero
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return ErrNonZero
		}
	case reflect.String:
		if v.String() == "" {
			return ErrNonZero
		}
	case reflect.Bool:
		if !v.Bool() {
			return ErrNonZero
		}
	case reflect.Slice, reflect.Map:
		if v.IsNil() || v.Len() == 0 {
			return ErrNonZero
		}
	case reflect.Array:
		if v.Len() == 0 {
			return ErrNonZero
		}
	case reflect.Struct:
		// For structs, check if all fields are zero
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.IsZero() {
				return nil
			}
		}
		return r.e
	}

	return nil
}

// Errf sets a custom error message for non-zero validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := NonZero[int]().Errf("Age must be greater than 0")
//	err := rule.Validate(0)  // returns error with custom message
//
//	rule = NonZero[string]().Errf("Name is required")
//	err = rule.Validate("")  // returns error with custom message
func (r *NonZeroRule[T]) Errf(format string, args ...any) *NonZeroRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
