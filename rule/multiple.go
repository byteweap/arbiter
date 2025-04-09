// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating if a number is a multiple of another number.
package rule

import "fmt"

// ErrMultiple is returned when a value is not a multiple of the specified base number
const ErrMultipleFormat = "is not a multiple of %v"

// MultipleRule validates that a number is a multiple of a specified base number.
// This rule is useful for validating values that must be divisible by a specific number.
//
// Example:
//
//	rule := MultipleOf(2)  // value must be divisible by 2
//	err := rule.Validate(4)   // returns nil (4 is divisible by 2)
//	err = rule.Validate(5)    // returns error (5 is not divisible by 2)
//
//	rule = MultipleOf(3)  // value must be divisible by 3
//	err = rule.Validate(6)   // returns nil (6 is divisible by 3)
//	err = rule.Validate(7)   // returns error (7 is not divisible by 3)
type MultipleRule struct {
	base int // base multiple
	e    error
}

// MultipleOf creates a new multiple validation rule.
// The rule ensures that a value is divisible by the specified base number.
//
// Example:
//
//	// Check if a number is even (divisible by 2)
//	rule := MultipleOf(2)
//	err := rule.Validate(10)  // returns nil (10 is divisible by 2)
//	err = rule.Validate(11)   // returns error (11 is not divisible by 2)
//
//	// Check if a number is divisible by 5
//	rule = MultipleOf(5)
//	err = rule.Validate(15)   // returns nil (15 is divisible by 5)
//	err = rule.Validate(16)   // returns error (16 is not divisible by 5)
func MultipleOf(base int) *MultipleRule {
	return &MultipleRule{base: base, e: fmt.Errorf(ErrMultipleFormat, base)}
}

// Validate checks if the value is divisible by the base number.
// Returns nil if the value is a multiple of the base, or an error otherwise.
//
// Example:
//
//	rule := MultipleOf(3)
//	err := rule.Validate(6)   // returns nil (6 is divisible by 3)
//	err = rule.Validate(7)    // returns error (7 is not divisible by 3)
//	err = rule.Validate(0)    // returns nil (0 is divisible by any number)
func (r *MultipleRule) Validate(value int) error {
	if value%r.base != 0 {
		return r.e
	}
	return nil
}

// Errf sets a custom error message for multiple validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := MultipleOf(2).Errf("Value must be an even number")
//	err := rule.Validate(5)  // returns error with custom message
//
//	rule = MultipleOf(5).Errf("Amount must be in multiples of 5")
//	err = rule.Validate(7)  // returns error with custom message
func (r *MultipleRule) Errf(format string, args ...any) *MultipleRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
