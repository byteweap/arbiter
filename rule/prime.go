// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating that a number is prime.
package rule

import (
	"errors"
	"fmt"
	"math"
)

// Error variable for prime validation
var (
	// ErrPrime is returned when a value must be prime but is not
	ErrPrime = errors.New("value is not a prime number")
)

// PrimeRule validates that a number is prime.
// A prime number is a natural number greater than 1 that is not a product of two smaller natural numbers.
//
// Example:
//
//	// Create a prime rule
//	rule := Prime()
//
//	// Valid prime numbers
//	err := rule.Validate(2)  // returns nil (2 is prime)
//	err = rule.Validate(3)   // returns nil (3 is prime)
//	err = rule.Validate(5)   // returns nil (5 is prime)
//	err = rule.Validate(7)   // returns nil (7 is prime)
//	err = rule.Validate(11)  // returns nil (11 is prime)
//
//	// Invalid non-prime numbers
//	err = rule.Validate(1)   // returns error (1 is not prime)
//	err = rule.Validate(4)   // returns error (4 is not prime)
//	err = rule.Validate(6)   // returns error (6 is not prime)
//	err = rule.Validate(8)   // returns error (8 is not prime)
//	err = rule.Validate(9)   // returns error (9 is not prime)
type PrimeRule struct {
	e error
}

// Prime creates a new prime validation rule.
// This function returns a rule that can be used to validate that a number
// is prime (greater than 1 and divisible only by 1 and itself).
//
// Example:
//
//	// Create a prime rule
//	rule := Prime()
//
//	// Use the rule
//	err := rule.Validate(2)  // returns nil
//	err = rule.Validate(3)   // returns nil
//	err = rule.Validate(4)   // returns error
func Prime() *PrimeRule {
	return &PrimeRule{
		e: ErrPrime,
	}
}

// Validate checks if the value is a prime number.
// Returns nil if the value is prime, or an error otherwise.
//
// A number is considered prime if:
// - It is greater than 1
// - It is not divisible by any number between 2 and its square root
//
// Example:
//
//	rule := Prime()
//
//	// Valid prime numbers
//	err := rule.Validate(2)   // returns nil
//	err = rule.Validate(3)    // returns nil
//	err = rule.Validate(5)    // returns nil
//	err = rule.Validate(7)    // returns nil
//	err = rule.Validate(11)   // returns nil
//	err = rule.Validate(13)   // returns nil
//	err = rule.Validate(17)   // returns nil
//	err = rule.Validate(19)   // returns nil
//	err = rule.Validate(23)   // returns nil
//	err = rule.Validate(29)   // returns nil
//
//	// Invalid non-prime numbers
//	err = rule.Validate(1)    // returns error
//	err = rule.Validate(4)    // returns error
//	err = rule.Validate(6)    // returns error
//	err = rule.Validate(8)    // returns error
//	err = rule.Validate(9)    // returns error
//	err = rule.Validate(10)   // returns error
//	err = rule.Validate(12)   // returns error
//	err = rule.Validate(14)   // returns error
//	err = rule.Validate(15)   // returns error
//	err = rule.Validate(16)   // returns error
func (r *PrimeRule) Validate(value int) error {
	if value <= 1 {
		if r.e != nil {
			return r.e
		}
		return ErrPrime
	}

	// For numbers greater than 1, check if they are divisible by any number
	// between 2 and the square root of the value
	sqrt := int(math.Sqrt(float64(value)))
	for i := 2; i <= sqrt; i++ {
		if value%i == 0 {
			if r.e != nil {
				return r.e
			}
			return ErrPrime
		}
	}
	return nil
}

// Errf sets a custom error message for prime validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Prime().Errf("The number must be prime")
//	err := rule.Validate(4)  // returns error with custom message
//
//	rule = Prime().Errf("Please enter a prime number")
//	err = rule.Validate(6)  // returns error with custom message
func (r *PrimeRule) Errf(format string, args ...any) *PrimeRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
