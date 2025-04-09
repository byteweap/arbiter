// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating minimum and maximum values.
package rule

import (
	"errors"
	"fmt"
)

// Error variables for min/max validation
var (
	// ErrMin is returned when a value is less than the minimum allowed value
	ErrMin = errors.New("value is less than minimum")
	// ErrMax is returned when a value is greater than the maximum allowed value
	ErrMax = errors.New("value is greater than maximum")
)

// MinRule validates that a value is greater than or equal to a minimum value.
// This rule works with any ordered type (numbers, strings, etc.).
//
// Example:
//
//	rule := Min[int](0)  // value must be >= 0
//	err := rule.Validate(5)   // returns nil (5 >= 0)
//	err = rule.Validate(-1)   // returns error (-1 < 0)
//
//	rule = Min[string]("apple")  // string must be >= "apple"
//	err = rule.Validate("banana")  // returns nil ("banana" >= "apple")
//	err = rule.Validate("ant")     // returns error ("ant" < "apple")
type MinRule[T Ordered] struct {
	min T
	e   error
}

// Min creates a new minimum value validation rule.
// The rule ensures that a value is greater than or equal to the specified minimum.
//
// Example:
//
//	// For numbers
//	rule := Min[int](0)  // value must be >= 0
//	err := rule.Validate(42)  // returns nil
//	err = rule.Validate(-5)   // returns error
//
//	// For strings
//	rule = Min[string]("a")  // string must be >= "a"
//	err = rule.Validate("hello")  // returns nil
//	err = rule.Validate("")       // returns error
func Min[T Ordered](min T) *MinRule[T] {
	return &MinRule[T]{min: min, e: ErrMin}
}

// Validate checks if the value is greater than or equal to the minimum.
// Returns nil if the value is valid, or an error otherwise.
//
// Example:
//
//	rule := Min[int](10)
//	err := rule.Validate(15)  // returns nil (15 >= 10)
//	err = rule.Validate(5)    // returns error (5 < 10)
//	err = rule.Validate(10)   // returns nil (10 >= 10)
func (r *MinRule[T]) Validate(value T) error {
	if value < r.min {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v is less than minimum %v", value, r.min)
	}
	return nil
}

// Errf sets a custom error message for minimum value validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Min[int](18).Errf("Age must be at least 18 years old")
//	err := rule.Validate(16)  // returns error with custom message
func (r *MinRule[T]) Errf(format string, args ...any) *MinRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// MaxRule validates that a value is less than or equal to a maximum value.
// This rule works with any ordered type (numbers, strings, etc.).
//
// Example:
//
//	rule := Max[int](100)  // value must be <= 100
//	err := rule.Validate(50)   // returns nil (50 <= 100)
//	err = rule.Validate(150)   // returns error (150 > 100)
//
//	rule = Max[string]("z")  // string must be <= "z"
//	err = rule.Validate("hello")  // returns nil ("hello" <= "z")
//	err = rule.Validate("zzz")    // returns error ("zzz" > "z")
type MaxRule[T Ordered] struct {
	max T
	e   error
}

// Max creates a new maximum value validation rule.
// The rule ensures that a value is less than or equal to the specified maximum.
//
// Example:
//
//	// For numbers
//	rule := Max[int](100)  // value must be <= 100
//	err := rule.Validate(50)  // returns nil
//	err = rule.Validate(150)  // returns error
//
//	// For strings
//	rule = Max[string]("z")  // string must be <= "z"
//	err = rule.Validate("hello")  // returns nil
//	err = rule.Validate("zzz")    // returns error
func Max[T Ordered](max T) *MaxRule[T] {
	return &MaxRule[T]{max: max, e: ErrMax}
}

// Validate checks if the value is less than or equal to the maximum.
// Returns nil if the value is valid, or an error otherwise.
//
// Example:
//
//	rule := Max[int](100)
//	err := rule.Validate(50)   // returns nil (50 <= 100)
//	err = rule.Validate(150)   // returns error (150 > 100)
//	err = rule.Validate(100)   // returns nil (100 <= 100)
func (r *MaxRule[T]) Validate(value T) error {
	if value > r.max {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v is greater than maximum %v", value, r.max)
	}
	return nil
}

// Errf sets a custom error message for maximum value validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Max[int](100).Errf("Score cannot exceed 100 points")
//	err := rule.Validate(150)  // returns error with custom message
func (r *MaxRule[T]) Errf(format string, args ...any) *MaxRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
