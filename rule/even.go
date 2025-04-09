// Package rule provides a collection of validation rules for various data types.
// This file contains the even number validation rule.
package rule

import (
	"errors"
	"fmt"
)

// Error returned when a value is not an even number
var (
	ErrEven = errors.New("value must be even")
)

// EvenRule validates that a number is even (divisible by 2).
// This rule works with any integer type (int, int8, int16, int32, int64, uint, etc.).
//
// Example:
//
//	rule := Even[int]()
//	err := rule.Validate(2)   // returns nil (2 is even)
//	err = rule.Validate(3)    // returns error (3 is odd)
//	err = rule.Validate(0)    // returns nil (0 is even)
type EvenRule[T Integer] struct {
	e error
}

// Even creates a new even number validation rule.
// The rule ensures that a number is divisible by 2.
//
// Example:
//
//	// For int type
//	rule := Even[int]()
//	err := rule.Validate(42)  // returns nil
//
//	// For int64 type
//	rule64 := Even[int64]()
//	err = rule64.Validate(100)  // returns nil
func Even[T Integer]() *EvenRule[T] {
	return &EvenRule[T]{
		e: ErrEven,
	}
}

// Validate checks if the given number is even (divisible by 2).
// Returns nil if the number is even, or an error otherwise.
//
// Example:
//
//	rule := Even[int]()
//	err := rule.Validate(4)    // returns nil (4 is even)
//	err = rule.Validate(5)     // returns error (5 is odd)
//	err = rule.Validate(-2)    // returns nil (-2 is even)
func (r *EvenRule[T]) Validate(value T) error {
	// Since T is constrained to Integer, value is guaranteed to be an integer type
	// We can directly use the % operator
	if value%2 != 0 {
		if r.e != nil {
			return r.e
		}
		return ErrEven
	}
	return nil
}

// Errf sets a custom error message for even number validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Even[int]().Errf("Please enter an even number")
//	err := rule.Validate(3)  // returns error with custom message
func (r *EvenRule[T]) Errf(format string, args ...any) *EvenRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
