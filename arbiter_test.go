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

// TestValidateAll tests the ValidateAll function with various types and rules.
// It verifies that the function collects all validation errors instead of stopping at the first one.
func TestValidateAll(t *testing.T) {
	// Test all rules pass
	t.Run("all rules pass", func(t *testing.T) {
		errs := arbiter.ValidateAll("hello", rule.Len[string](3, 10))
		if len(errs) != 0 {
			t.Errorf("Expected no errors, got %d: %v", len(errs), errs)
		}
	})

	// Test single rule fails
	t.Run("single rule fails", func(t *testing.T) {
		errs := arbiter.ValidateAll("hi", rule.Len[string](3, 10))
		if len(errs) != 1 {
			t.Errorf("Expected 1 error, got %d", len(errs))
		}
	})

	// Test no rules
	t.Run("no rules", func(t *testing.T) {
		errs := arbiter.ValidateAll("hello")
		if len(errs) != 0 {
			t.Errorf("Expected no errors, got %d", len(errs))
		}
	})
}

// TestValidateAllMultipleErrors tests that ValidateAll collects multiple errors.
func TestValidateAllMultipleErrors(t *testing.T) {
	errs := arbiter.ValidateAll(151, rule.Min[int](0), rule.Max[int](100), rule.Even[int]())
	if len(errs) != 2 {
		t.Errorf("Expected 2 errors, got %d: %v", len(errs), errs)
	}
}

// TestValidateAllCustomError tests ValidateAll with custom error messages.
func TestValidateAllCustomError(t *testing.T) {
	errs := arbiter.ValidateAll("ab", rule.Len[string](3, 10).Errf("too short"))
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "too short" {
		t.Errorf("Expected 'too short', got '%s'", errs[0].Error())
	}
}

// TestValidateAllSlice tests ValidateAll with slice validation.
func TestValidateAllSlice(t *testing.T) {
	// Valid slice
	errs := arbiter.ValidateAll([]int{1, 2, 3}, rule.Len[[]int](2, 5))
	if len(errs) != 0 {
		t.Errorf("Expected no errors for valid slice, got %d", len(errs))
	}

	// Empty slice
	errs = arbiter.ValidateAll([]int{}, rule.Len[[]int](2, 5))
	if len(errs) != 1 {
		t.Errorf("Expected 1 error for empty slice, got %d", len(errs))
	}
}

// TestValidateAllVsValidate compares ValidateAll and Validate behavior.
func TestValidateAllVsValidate(t *testing.T) {
	// Validate stops at first error
	err := arbiter.Validate(151, rule.Min[int](0), rule.Max[int](100), rule.Even[int]())
	if err == nil {
		t.Error("Expected error from Validate")
	}

	// ValidateAll collects all errors
	errs := arbiter.ValidateAll(151, rule.Min[int](0), rule.Max[int](100), rule.Even[int]())
	if len(errs) != 2 {
		t.Errorf("Expected 2 errors from ValidateAll, got %d: %v", len(errs), errs)
	}
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
		err := arbiter.ValidateStruct(nil, "Person cannot be nil")
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
