// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
)

// ErrNegative is returned when a value is not negative (less than zero).
var ErrNegative = errors.New("value must be negative")

// NegativeRule is a validation rule that checks if a value is negative (less than zero).
// It supports any ordered numeric type through generics.
//
// Example:
//
//	rule := Negative[int]()
//	err := rule.Validate(-42)   // returns nil
//	err = rule.Validate(0)      // returns ErrNegative
//	err = rule.Validate(1)      // returns ErrNegative
//
//	rule = Negative[float64]()
//	err = rule.Validate(-3.14)  // returns nil
//	err = rule.Validate(0.0)    // returns ErrNegative
type NegativeRule[T Ordered] struct {
	e error
}

// Negative creates a new negative number validation rule.
// The rule can be used with any type that implements the Ordered interface.
//
// Example:
//
//	// Create a rule for negative integers
//	temperatureRule := Negative[int]().Errf("Temperature must be negative")
//
//	// Create a rule for negative floating-point numbers
//	balanceRule := Negative[float64]().Errf("Balance must be negative")
func Negative[T Ordered]() *NegativeRule[T] {
	return &NegativeRule[T]{
		e: ErrNegative,
	}
}

// Validate checks if the value is negative (less than zero).
// Returns nil if the value is negative, or an error if it's not.
//
// Example:
//
//	rule := Negative[int]()
//	if err := rule.Validate(-42); err != nil {
//	    // Handle validation error
//	}
func (r *NegativeRule[T]) Validate(value T) error {
	var zero T
	if value >= zero {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v must be negative", value)
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Negative[int]().Errf("Value must be less than zero")
//	err := rule.Validate(0)  // returns error with message "Value must be less than zero"
func (r *NegativeRule[T]) Errf(format string, args ...any) *NegativeRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
