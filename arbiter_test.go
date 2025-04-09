// Package arbiter_test provides tests for the arbiter package.
// These tests verify the functionality of the validation functions.
package arbiter_test

import (
	"testing"

	"github.com/byteweap/arbiter"
	"github.com/byteweap/arbiter/rule"
)

// TestValidate tests the Validate function with various types and rules.
// It verifies that the function correctly applies validation rules to values.
func TestValidate(t *testing.T) {
	// Test string validation with domain rule
	t.Run("domain validation", func(t *testing.T) {
		// Valid domain
		err := arbiter.Validate("example.com", rule.Domain())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Invalid domain (missing TLD)
		err = arbiter.Validate("invalid", rule.Domain())
		if err == nil {
			t.Error("Expected error for invalid domain, got nil")
		}

		// Invalid domain (double dots)
		err = arbiter.Validate("invalid..com", rule.Domain())
		if err == nil {
			t.Error("Expected error for invalid domain, got nil")
		}
	})

	// Test number validation with min/max rules
	t.Run("number validation", func(t *testing.T) {
		// Valid number
		err := arbiter.Validate(42, rule.Min[int](0), rule.Max[int](100))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Invalid number (too small)
		err = arbiter.Validate(-1, rule.Min[int](0), rule.Max[int](100))
		if err == nil {
			t.Error("Expected error for number too small, got nil")
		}

		// Invalid number (too large)
		err = arbiter.Validate(101, rule.Min[int](0), rule.Max[int](100))
		if err == nil {
			t.Error("Expected error for number too large, got nil")
		}
	})

	// Test multiple rules
	t.Run("multiple rules", func(t *testing.T) {
		// Valid value (passes all rules)
		err := arbiter.Validate("example.com", rule.Domain(), rule.OnlyHalfWidth())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Invalid value (fails first rule)
		err = arbiter.Validate("invalid", rule.Domain(), rule.OnlyHalfWidth())
		if err == nil {
			t.Error("Expected error for invalid domain, got nil")
		}

		// Invalid value (fails second rule)
		err = arbiter.Validate("Ｈｅｌｌｏ", rule.Domain(), rule.OnlyHalfWidth())
		if err == nil {
			t.Error("Expected error for full-width characters, got nil")
		}
	})
}

// TestValidateStruct tests the ValidateStruct function with various structs and field rules.
// It verifies that the function correctly validates struct fields.
func TestValidateStruct(t *testing.T) {
	// Define a test struct
	type Person struct {
		Name     string
		Age      int
		Website  string
		Password string
	}

	// Test valid struct
	t.Run("valid struct", func(t *testing.T) {
		person := &Person{
			Name:     "John",
			Age:      30,
			Website:  "example.com",
			Password: "StrongP@ssw0rd",
		}

		err := arbiter.ValidateStruct(person, "Person cannot be nil",
			arbiter.Field(&person.Name, rule.OnlyHalfWidth()),
			arbiter.Field(&person.Age, rule.Min[int](0), rule.Max[int](120)),
			arbiter.Field(&person.Website, rule.Domain()),
			arbiter.Field(&person.Password, rule.PasswordStrength()),
		)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	// Test invalid struct (nil)
	t.Run("nil struct", func(t *testing.T) {
		var person *Person = nil

		err := arbiter.ValidateStruct(person, "Person cannot be nil",
			arbiter.Field(&person.Name, rule.OnlyHalfWidth()),
		)

		if err == nil {
			t.Error("Expected error for nil struct, got nil")
		}
	})

	// Test invalid struct (non-pointer)
	t.Run("non-pointer struct", func(t *testing.T) {
		person := Person{
			Name:     "John",
			Age:      30,
			Website:  "example.com",
			Password: "StrongP@ssw0rd",
		}

		err := arbiter.ValidateStruct(person, "Person cannot be nil",
			arbiter.Field(&person.Name, rule.OnlyHalfWidth()),
		)

		if err == nil {
			t.Error("Expected error for non-pointer struct, got nil")
		}
	})

	// Test invalid field
	t.Run("invalid field", func(t *testing.T) {
		person := &Person{
			Name:     "Ｊｏｈｎ", // Full-width characters
			Age:      30,
			Website:  "example.com",
			Password: "StrongP@ssw0rd",
		}

		err := arbiter.ValidateStruct(person, "Person cannot be nil",
			arbiter.Field(&person.Name, rule.OnlyHalfWidth()),
			arbiter.Field(&person.Age, rule.Min[int](0), rule.Max[int](120)),
			arbiter.Field(&person.Website, rule.Domain()),
			arbiter.Field(&person.Password, rule.PasswordStrength()),
		)

		if err == nil {
			t.Error("Expected error for full-width characters in name, got nil")
		}
	})
}
