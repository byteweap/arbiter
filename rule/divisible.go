package rule

import (
	"errors"
	"fmt"
	"math"
)

// ErrDivisibleBy is returned when a value is not divisible by the specified number
var ErrDivisibleBy = errors.New("value is not divisible by the specified number")

// DivisibleByRule represents a validation rule that checks if a number is divisible by a given divisor
// Example: DivisibleBy(2) will validate that a number is divisible by 2 (even numbers)
type DivisibleByRule struct {
	divisor float64
	e       error
}

// DivisibleBy creates a new rule that validates if a number is divisible by the given divisor
// Example: rule := DivisibleBy(3) // Creates a rule to check if numbers are divisible by 3
func DivisibleBy(divisor float64) *DivisibleByRule {
	return &DivisibleByRule{
		divisor: divisor,
		e:       ErrDivisibleBy,
	}
}

// Validate checks if the given value is divisible by the rule's divisor
// Returns nil if divisible, error otherwise
// Example:
//
//	rule := DivisibleBy(5)
//	err := rule.Validate(10) // Returns nil (10 is divisible by 5)
//	err := rule.Validate(7)  // Returns error (7 is not divisible by 5)
func (r *DivisibleByRule) Validate(value float64) error {
	if r.divisor == 0 {
		return errors.New("divisor cannot be zero")
	}

	remainder := math.Mod(value, r.divisor)

	if math.Abs(remainder) > 1e-10 {
		if r.e != nil {
			return r.e
		}
		return ErrDivisibleBy
	}
	return nil
}

// Errf sets a custom error message for the rule using a format string
// Example: rule.Errf("Number %v must be divisible by %v", value, divisor)
func (r *DivisibleByRule) Errf(format string, args ...any) *DivisibleByRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
