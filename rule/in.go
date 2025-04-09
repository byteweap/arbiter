// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating if a value is in or not in a list of values.
package rule

import (
	"errors"
	"fmt"
	"slices"
)

// Error variables for in/not in validation
var (
	// ErrIn is returned when a value must be in a list but is not found
	ErrIn = errors.New("must be in the list")
	// ErrNotIn is returned when a value must not be in a list but is found
	ErrNotIn = errors.New("must not be in the list")
)

// InRule validates if a value is in or not in a list of values.
// This rule works with any comparable type (strings, numbers, etc.).
//
// Example:
//
//	rule := In("apple", "banana", "orange")
//	err := rule.Validate("apple")   // returns nil (apple is in the list)
//	err = rule.Validate("grape")    // returns error (grape is not in the list)
//
//	rule = NotIn("apple", "banana")
//	err = rule.Validate("orange")   // returns nil (orange is not in the list)
//	err = rule.Validate("apple")    // returns error (apple is in the list)
type InRule[T InType] struct {
	values []T
	notIn  bool
	e      error
}

// In creates a new rule that validates a value must be in the specified list.
// The rule returns an error if the value is not found in the list.
//
// Example:
//
//	// Using variadic arguments
//	rule := In("red", "green", "blue")
//	err := rule.Validate("red")     // returns nil
//	err = rule.Validate("yellow")   // returns error
//
//	// Using a slice
//	colors := []string{"red", "green", "blue"}
//	rule = InSlice(colors)
//	err = rule.Validate("green")    // returns nil
func In[T InType](values ...T) *InRule[T] {
	return &InRule[T]{e: ErrIn, values: values, notIn: false}
}

// InSlice creates a new rule that validates a value must be in the specified slice.
// This is an alternative to In() when you already have a slice of values.
//
// Example:
//
//	colors := []string{"red", "green", "blue"}
//	rule := InSlice(colors)
//	err := rule.Validate("blue")    // returns nil
//	err = rule.Validate("yellow")   // returns error
func InSlice[T InType](values []T) *InRule[T] {
	return &InRule[T]{e: ErrIn, values: values, notIn: false}
}

// NotIn creates a new rule that validates a value must not be in the specified list.
// The rule returns an error if the value is found in the list.
//
// Example:
//
//	// Using variadic arguments
//	rule := NotIn("admin", "root")
//	err := rule.Validate("user")    // returns nil
//	err = rule.Validate("admin")    // returns error
//
//	// Using a slice
//	restricted := []string{"admin", "root"}
//	rule = NotInSlice(restricted)
//	err = rule.Validate("guest")    // returns nil
func NotIn[T InType](values ...T) *InRule[T] {
	return &InRule[T]{e: ErrNotIn, values: values, notIn: true}
}

// NotInSlice creates a new rule that validates a value must not be in the specified slice.
// This is an alternative to NotIn() when you already have a slice of values.
//
// Example:
//
//	restricted := []string{"admin", "root"}
//	rule := NotInSlice(restricted)
//	err := rule.Validate("user")    // returns nil
//	err = rule.Validate("admin")    // returns error
func NotInSlice[T InType](values []T) *InRule[T] {
	return &InRule[T]{e: ErrNotIn, values: values, notIn: true}
}

// Validate checks if the value is in or not in the list, depending on the rule type.
// Returns nil if the validation passes, or an error otherwise.
//
// Example:
//
//	// In rule
//	rule := In("apple", "banana")
//	err := rule.Validate("apple")   // returns nil
//	err = rule.Validate("orange")   // returns error
//
//	// NotIn rule
//	rule = NotIn("admin", "root")
//	err = rule.Validate("user")     // returns nil
//	err = rule.Validate("admin")    // returns error
func (r *InRule[T]) Validate(value T) error {
	isIn := slices.Contains(r.values, value)
	if isIn == r.notIn {
		if r.e != nil {
			return r.e
		}
		return ErrIn
	}
	return nil
}

// Errf sets a custom error message for validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := In("red", "green", "blue").Errf("Color must be one of: %v", []string{"red", "green", "blue"})
//	err := rule.Validate("yellow")  // returns error with custom message
//
//	rule = NotIn("admin", "root").Errf("Username cannot be %v", []string{"admin", "root"})
//	err = rule.Validate("admin")    // returns error with custom message
func (r *InRule[T]) Errf(format string, args ...any) *InRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
