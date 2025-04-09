// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating that a number is odd.
package rule

import (
	"errors"
	"fmt"
)

// Error variable for odd validation
var (
	// ErrOdd is returned when a value must be odd but is even
	ErrOdd = errors.New("value must be odd")
)

// OddRule validates that a number is odd.
// This rule works with any integer type (int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64).
//
// Example:
//
//	// Create an odd rule
//	rule := Odd[int]()
//
//	// Validate numbers
//	err := rule.Validate(3)  // returns nil (odd)
//	err = rule.Validate(4)   // returns error (even)
//
//	// With different integer types
//	int8Rule := Odd[int8]()
//	err = int8Rule.Validate(5)  // returns nil (odd)
//	err = int8Rule.Validate(6)  // returns error (even)
//
//	uintRule := Odd[uint]()
//	err = uintRule.Validate(7)  // returns nil (odd)
//	err = uintRule.Validate(8)  // returns error (even)
type OddRule[T Integer] struct {
	e error
}

// Odd creates a new odd validation rule.
// This function returns a rule that can be used to validate that a number
// is odd (not divisible by 2).
//
// Example:
//
//	// Create rules for different integer types
//	intRule := Odd[int]()
//	int32Rule := Odd[int32]()
//	uint64Rule := Odd[uint64]()
//
//	// Use the rules
//	err := intRule.Validate(3)
//	err = int32Rule.Validate(5)
//	err = uint64Rule.Validate(7)
func Odd[T Integer]() *OddRule[T] {
	return &OddRule[T]{
		e: ErrOdd,
	}
}

// Validate checks if the value is odd.
// Returns nil if the value is odd, or an error otherwise.
//
// Example:
//
//	rule := Odd[int]()
//
//	// Valid odd numbers
//	err := rule.Validate(1)  // returns nil
//	err = rule.Validate(3)   // returns nil
//	err = rule.Validate(5)   // returns nil
//	err = rule.Validate(-1)  // returns nil
//	err = rule.Validate(-3)  // returns nil
//
//	// Invalid even numbers
//	err = rule.Validate(0)   // returns error
//	err = rule.Validate(2)   // returns error
//	err = rule.Validate(4)   // returns error
//	err = rule.Validate(-2)  // returns error
//	err = rule.Validate(-4)  // returns error
func (r *OddRule[T]) Validate(value T) error {
	// Since T is constrained to Integer, value is guaranteed to be an integer type
	// We can directly use the % operator
	if value%2 == 0 {
		if r.e != nil {
			return r.e
		}
		return ErrOdd
	}
	return nil
}

// Errf sets a custom error message for odd validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Odd[int]().Errf("The number must be odd")
//	err := rule.Validate(4)  // returns error with custom message
//
//	rule = Odd[int32]().Errf("Please enter an odd number")
//	err = rule.Validate(6)  // returns error with custom message
func (r *OddRule[T]) Errf(format string, args ...any) *OddRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
