// Package rule provides a collection of validation rules for various data types.
// This file contains rules for validating UUID (Universally Unique Identifier) strings.
package rule

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Error variable for UUID validation
var (
	// ErrUUID is returned when a string is not a valid UUID format
	ErrUUID = errors.New("invalid UUID format")
)

// UUIDRule validates that a string is a valid UUID.
// A UUID is a 128-bit number used to identify information in computer systems.
// The standard format is 8-4-4-4-12 hexadecimal digits (e.g., 123e4567-e89b-12d3-a456-426614174000).
//
// Example:
//
//	// Create a UUID rule
//	rule := UUID()
//
//	// Valid UUIDs
//	err := rule.Validate("123e4567-e89b-12d3-a456-426614174000")  // returns nil
//	err = rule.Validate("550e8400-e29b-41d4-a716-446655440000")   // returns nil
//	err = rule.Validate("")                                       // returns nil (empty string is valid)
//
//	// Invalid UUIDs
//	err = rule.Validate("123e4567-e89b-12d3-a456")               // returns error (incomplete)
//	err = rule.Validate("123e4567-e89b-12d3-a456-42661417400g")  // returns error (invalid character)
//	err = rule.Validate("123e4567-e89b-12d3-a456-42661417400")   // returns error (wrong length)
type UUIDRule struct {
	e error
}

// UUID creates a new UUID validation rule.
// This function returns a rule that can be used to validate that a string
// is a valid UUID in the standard format (8-4-4-4-12 hexadecimal digits).
//
// Example:
//
//	// Create a UUID rule
//	rule := UUID()
//
//	// Use the rule
//	err := rule.Validate("123e4567-e89b-12d3-a456-426614174000")  // returns nil
//	err = rule.Validate("invalid-uuid")                           // returns error
func UUID() *UUIDRule {
	return &UUIDRule{
		e: ErrUUID,
	}
}

// Validate checks if the string is a valid UUID.
// Returns nil if the string is a valid UUID or empty, or an error otherwise.
//
// A valid UUID must:
// - Be in the format 8-4-4-4-12 hexadecimal digits
// - Contain only hexadecimal characters (0-9, a-f)
// - Have hyphens in the correct positions
//
// Example:
//
//	rule := UUID()
//
//	// Valid UUIDs
//	err := rule.Validate("123e4567-e89b-12d3-a456-426614174000")  // returns nil
//	err = rule.Validate("550e8400-e29b-41d4-a716-446655440000")   // returns nil
//	err = rule.Validate("6ba7b810-9dad-11d1-80b4-00c04fd430c8")   // returns nil
//	err = rule.Validate("")                                       // returns nil
//
//	// Invalid UUIDs
//	err = rule.Validate("123e4567-e89b-12d3-a456")               // returns error
//	err = rule.Validate("123e4567-e89b-12d3-a456-42661417400g")  // returns error
//	err = rule.Validate("123e4567-e89b-12d3-a456-42661417400")   // returns error
//	err = rule.Validate("123e4567e89b12d3a456426614174000")      // returns error
func (r *UUIDRule) Validate(value string) error {
	if value == "" {
		return nil
	}

	// Regular expression for UUID format
	pattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !pattern.MatchString(strings.ToLower(value)) {
		if r.e != nil {
			return r.e
		}
		return ErrUUID
	}
	return nil
}

// Errf sets a custom error message for UUID validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := UUID().Err("The ID must be a valid UUID")
//	err := rule.Validate("invalid-uuid")  // returns error with custom message
//
//	rule = UUID().Errf("Please enter a valid UUID")
//	err = rule.Validate("123e4567-e89b-12d3-a456")  // returns error with custom message
func (r *UUIDRule) Errf(format string, args ...any) *UUIDRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
