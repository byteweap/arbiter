// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
)

// ErrPositive is returned when a value is not positive (greater than zero).
var ErrPositive = errors.New("value must be positive")

// PositiveRule is a validation rule that checks if a value is positive (greater than zero).
// It supports any ordered numeric type through generics.
//
// Example:
//
//	rule := Positive[int]()
//	err := rule.Validate(42)   // returns nil
//	err = rule.Validate(0)     // returns ErrPositive
//	err = rule.Validate(-1)    // returns ErrPositive
//
//	rule = Positive[float64]()
//	err = rule.Validate(3.14)  // returns nil
//	err = rule.Validate(0.0)   // returns ErrPositive
type PositiveRule[T Ordered] struct {
	e error
}

// Positive creates a new positive number validation rule.
// The rule can be used with any type that implements the Ordered interface.
//
// Example:
//
//	// Create a rule for positive integers
//	ageRule := Positive[int]().Err("Age must be positive")
//
//	// Create a rule for positive floating-point numbers
//	priceRule := Positive[float64]().Err("Price must be positive")
func Positive[T Ordered]() *PositiveRule[T] {
	return &PositiveRule[T]{
		e: ErrPositive,
	}
}

// Validate checks if the value is positive (greater than zero).
// Returns nil if the value is positive, or an error if it's not.
//
// Example:
//
//	rule := Positive[int]()
//	if err := rule.Validate(42); err != nil {
//	    // Handle validation error
//	}
func (r *PositiveRule[T]) Validate(value T) error {
	var zero T
	if value <= zero {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v must be positive", value)
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Positive[int]().Errf("Value must be greater than zero")
//	err := rule.Validate(0)  // returns error with message "Value must be greater than zero"
func (r *PositiveRule[T]) Errf(format string, args ...any) *PositiveRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
