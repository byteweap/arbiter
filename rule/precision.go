// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating the precision of floating-point numbers.
package rule

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Error variable for precision validation
var (
	// ErrPrecision is returned when a number's decimal places exceed the specified precision
	ErrPrecision = errors.New("number precision exceeds the specified limit")
)

// PrecisionRule validates that a float64 number's decimal places do not exceed
// a specified precision. This rule ensures that floating-point numbers maintain
// a consistent level of precision.
//
// Example:
//
//	// Create a precision rule that allows up to 2 decimal places
//	rule := Precision(2)
//
//	// Valid numbers (2 or fewer decimal places)
//	err := rule.Validate(3.14)    // returns nil
//	err = rule.Validate(3.1)      // returns nil
//	err = rule.Validate(3.0)      // returns nil
//	err = rule.Validate(3)        // returns nil
//
//	// Invalid numbers (more than 2 decimal places)
//	err = rule.Validate(3.141)    // returns error
//	err = rule.Validate(3.14159)  // returns error
type PrecisionRule struct {
	precision int
	e         error
}

// Precision creates a new precision validation rule for float64 values.
// The precision parameter specifies the maximum number of decimal places allowed.
//
// Example:
//
//	// Create rules with different precision requirements
//	twoDecimals := Precision(2)    // allows up to 2 decimal places
//	threeDecimals := Precision(3)  // allows up to 3 decimal places
//	noDecimals := Precision(0)     // allows only whole numbers
//
//	// Use the rules
//	err := twoDecimals.Validate(3.14)
//	err = threeDecimals.Validate(3.141)
//	err = noDecimals.Validate(3.0)
func Precision(precision int) *PrecisionRule {
	return &PrecisionRule{
		precision: precision,
		e:         ErrPrecision,
	}
}

// Validate checks if the float64 value's decimal places do not exceed
// the specified precision. Returns nil if the precision is valid,
// or an error otherwise.
//
// Example:
//
//	rule := Precision(2)
//
//	// Valid cases
//	err := rule.Validate(3.14)    // returns nil
//	err = rule.Validate(3.1)      // returns nil
//	err = rule.Validate(3.0)      // returns nil
//	err = rule.Validate(3)        // returns nil
//	err = rule.Validate(-3.14)    // returns nil
//
//	// Invalid cases
//	err = rule.Validate(3.141)    // returns error
//	err = rule.Validate(3.14159)  // returns error
//	err = rule.Validate(-3.141)   // returns error
func (r *PrecisionRule) Validate(value float64) error {
	// Convert float to string using scientific notation to avoid precision loss
	str := strconv.FormatFloat(value, 'e', -1, 64)

	// Split into integer and decimal parts
	parts := strings.Split(str, ".")
	if len(parts) != 2 {
		// If no decimal part, it's an integer, so validation passes
		return nil
	}

	// Get the length of the decimal part
	decimalLen := len(parts[1])

	// If decimal length exceeds the specified precision, return error
	if decimalLen > r.precision {
		return r.e
	}

	return nil
}

// Errf sets a custom error message for precision validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Precision(2).Errf("Price must have at most 2 decimal places")
//	err := rule.Validate(3.141)  // returns error with custom message
//
//	rule = Precision(1).Errf("Amount cannot have more than 1 decimal place")
//	err = rule.Validate(3.14)  // returns error with custom message
func (r *PrecisionRule) Errf(format string, args ...any) *PrecisionRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// Float32PrecisionRule validates that a float32 number's decimal places do not exceed
// a specified precision. This rule ensures that 32-bit floating-point numbers maintain
// a consistent level of precision.
//
// Example:
//
//	// Create a precision rule that allows up to 2 decimal places
//	rule := Float32Precision(2)
//
//	// Valid numbers (2 or fewer decimal places)
//	err := rule.Validate(3.14)    // returns nil
//	err = rule.Validate(3.1)      // returns nil
//	err = rule.Validate(3.0)      // returns nil
//	err = rule.Validate(3)        // returns nil
//
//	// Invalid numbers (more than 2 decimal places)
//	err = rule.Validate(3.141)    // returns error
//	err = rule.Validate(3.14159)  // returns error
type Float32PrecisionRule struct {
	precision int
	e         error
}

// Float32Precision creates a new precision validation rule for float32 values.
// The precision parameter specifies the maximum number of decimal places allowed.
//
// Example:
//
//	// Create rules with different precision requirements
//	twoDecimals := Float32Precision(2)    // allows up to 2 decimal places
//	threeDecimals := Float32Precision(3)  // allows up to 3 decimal places
//	noDecimals := Float32Precision(0)     // allows only whole numbers
//
//	// Use the rules
//	err := twoDecimals.Validate(3.14)
//	err = threeDecimals.Validate(3.141)
//	err = noDecimals.Validate(3.0)
func Float32Precision(precision int) *Float32PrecisionRule {
	return &Float32PrecisionRule{
		precision: precision,
		e:         ErrPrecision,
	}
}

// Validate checks if the float32 value's decimal places do not exceed
// the specified precision. Returns nil if the precision is valid,
// or an error otherwise.
//
// Example:
//
//	rule := Float32Precision(2)
//
//	// Valid cases
//	err := rule.Validate(3.14)    // returns nil
//	err = rule.Validate(3.1)      // returns nil
//	err = rule.Validate(3.0)      // returns nil
//	err = rule.Validate(3)        // returns nil
//	err = rule.Validate(-3.14)    // returns nil
//
//	// Invalid cases
//	err = rule.Validate(3.141)    // returns error
//	err = rule.Validate(3.14159)  // returns error
//	err = rule.Validate(-3.141)   // returns error
func (r *Float32PrecisionRule) Validate(value float32) error {
	// Convert float32 to string using scientific notation to avoid precision loss
	str := strconv.FormatFloat(float64(value), 'e', -1, 32)

	// Split into integer and decimal parts
	parts := strings.Split(str, ".")
	if len(parts) != 2 {
		// If no decimal part, it's an integer, so validation passes
		return nil
	}

	// Get the length of the decimal part
	decimalLen := len(parts[1])

	// If decimal length exceeds the specified precision, return error
	if decimalLen > r.precision {
		return r.e
	}

	return nil
}

// Errf sets a custom error message for precision validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Float32Precision(2).Errf("Price must have at most 2 decimal places")
//	err := rule.Validate(3.141)  // returns error with custom message
//
//	rule = Float32Precision(1).Errf("Amount cannot have more than 1 decimal place")
//	err = rule.Validate(3.14)  // returns error with custom message
func (r *Float32PrecisionRule) Errf(format string, args ...any) *Float32PrecisionRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
