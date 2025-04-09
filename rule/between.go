// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"fmt"
)

const (
	ErrBetweenFormat = "is not between %v and %v"
)

// BetweenRule is a validation rule that checks if a value falls within a specified range.
// It supports any ordered type (numbers) through generics.
//
// Example:
//
//	rule := Between[int](1, 10)
//	err := rule.Validate(5)   // returns nil
//	err = rule.Validate(15)   // returns ErrBetween
//
//	rule = Between[float64](0.0, 1.0)
//	err = rule.Validate(0.5)  // returns nil
//	err = rule.Validate(1.5)  // returns ErrBetween
type BetweenRule[T Ordered] struct {
	min T
	max T
	e   error
}

// Between creates a new range validation rule with the specified minimum and maximum values.
// The rule will validate that values fall within the inclusive range [min, max].
//
// Example:
//
//	// Create a rule for integers between 1 and 100
//	ageRule := Between[int](1, 100)
//
//	// Create a rule for floating-point numbers between 0 and 1
//	probabilityRule := Between[float64](0.0, 1.0)
func Between[T Ordered](min, max T) *BetweenRule[T] {
	return &BetweenRule[T]{
		min: min,
		max: max,
		e:   fmt.Errorf(ErrBetweenFormat, min, max),
	}
}

// Errf sets a custom error message for the validation rule using a formatted string.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Between[int](1, 10).Errf("Age must be between %d and %d", 1, 10)
//	err := rule.Validate(15)  // returns error with message "Age must be between 1 and 10"
func (r *BetweenRule[T]) Errf(format string, args ...any) *BetweenRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// Validate checks if the provided value falls within the rule's range.
// Returns nil if the value is valid, or an error if it's outside the range.
//
// Example:
//
//	rule := Between[int](1, 10)
//	if err := rule.Validate(5); err != nil {
//	    // Handle validation error
//	}
func (r *BetweenRule[T]) Validate(value T) error {
	if value < r.min || value > r.max {
		if r.e != nil {
			return r.e
		}
		return fmt.Errorf("value %v is not between %v and %v", value, r.min, r.max)
	}
	return nil
}
