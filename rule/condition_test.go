// Package rule provides tests for validation rules.
package rule

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockValidator is a simple validator implementation for testing.
// It allows customizing the validation behavior through a function.
type mockValidator[T any] struct {
	validateFunc func(T) error
}

// Validate implements the Rule interface by calling the custom validation function.
func (m *mockValidator[T]) Validate(value T) error {
	return m.validateFunc(value)
}

// TestConditionRule tests the AND and OR condition rules.
// It verifies that logical combinations of validation rules work correctly.
func TestConditionRule(t *testing.T) {
	tests := []struct {
		name      string
		rule      *ConditionRule[string]
		value     string
		wantError bool
	}{
		{
			name: "AND rule - all rules pass",
			rule: And(
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
			),
			value:     "test",
			wantError: false,
		},
		{
			name: "AND rule - one rule fails",
			rule: And(
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
				&mockValidator[string]{
					validateFunc: func(s string) error { return errors.New("validation failed") },
				},
			),
			value:     "test",
			wantError: true,
		},
		{
			name: "OR rule - all rules pass",
			rule: Or(
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
			),
			value:     "test",
			wantError: false,
		},
		{
			name: "OR rule - one rule passes",
			rule: Or(
				&mockValidator[string]{
					validateFunc: func(s string) error { return errors.New("validation failed") },
				},
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
			),
			value:     "test",
			wantError: false,
		},
		{
			name: "OR rule - all rules fail",
			rule: Or(
				&mockValidator[string]{
					validateFunc: func(s string) error { return errors.New("validation failed 1") },
				},
				&mockValidator[string]{
					validateFunc: func(s string) error { return errors.New("validation failed 2") },
				},
			),
			value:     "test",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestDependencyRule tests the dependency validation rule.
// It verifies that validation rules can be applied to dependent fields.
func TestDependencyRule(t *testing.T) {
	// User struct for testing field dependencies
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name      string
		rule      *DependencyRule[User, string]
		value     User
		wantError bool
	}{
		{
			name: "Valid field value passes validation",
			rule: Dependency(
				"Name",
				"",
				&mockValidator[string]{
					validateFunc: func(s string) error { return nil },
				},
				func(u User) string { return u.Name },
			),
			value: User{
				Name: "John",
				Age:  25,
			},
			wantError: false,
		},
		{
			name: "Invalid field value fails validation",
			rule: Dependency(
				"Name",
				"",
				&mockValidator[string]{
					validateFunc: func(s string) error { return errors.New("invalid name") },
				},
				func(u User) string { return u.Name },
			),
			value: User{
				Name: "John123",
				Age:  25,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestMutualExcludeRule tests the mutual exclusion validation rule.
// It verifies that values are mutually exclusive within their respective sets.
func TestMutualExcludeRule(t *testing.T) {
	tests := []struct {
		name      string
		rule      *MutualExcludeRule[string]
		value     string
		wantError bool
	}{
		{
			name: "Value exists in allowed set",
			rule: MutualExclude(
				[]string{"admin", "user", "guest"},
				[]string{"admin", "user", "guest"},
				func(a, b string) bool { return a == b },
			),
			value:     "admin",
			wantError: false,
		},
		{
			name: "Value does not exist in allowed set",
			rule: MutualExclude(
				[]string{"admin", "user", "guest"},
				[]string{"admin", "user", "guest"},
				func(a, b string) bool { return a == b },
			),
			value:     "invalid",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestCustomErrorMessage tests custom error messages for condition rules.
// It verifies that custom error messages are correctly set and returned.
func TestCustomErrorMessage(t *testing.T) {
	// Test AND rule with custom error message
	t.Run("AND rule custom error", func(t *testing.T) {
		rule := And(
			&mockValidator[string]{
				validateFunc: func(s string) error { return errors.New("first error") },
			},
			&mockValidator[string]{
				validateFunc: func(s string) error { return errors.New("second error") },
			},
		).Errf("Custom AND error")

		err := rule.Validate("test")
		assert.Equal(t, "Custom AND error", err.Error())
	})

	// Test OR rule with custom error message
	t.Run("OR rule custom error", func(t *testing.T) {
		rule := Or(
			&mockValidator[string]{
				validateFunc: func(s string) error { return errors.New("first error") },
			},
			&mockValidator[string]{
				validateFunc: func(s string) error { return errors.New("second error") },
			},
		).Errf("Custom OR error")

		err := rule.Validate("test")
		assert.Equal(t, "Custom OR error", err.Error())
	})

	// Test dependency rule with custom error message
	t.Run("Dependency rule custom error", func(t *testing.T) {
		type User struct {
			Name string
		}

		rule := Dependency(
			"Name",
			"",
			&mockValidator[string]{
				validateFunc: func(s string) error { return errors.New("validation error") },
			},
			func(u User) string { return u.Name },
		).Errf("Custom dependency error")

		err := rule.Validate(User{Name: "test"})
		assert.Equal(t, "Custom dependency error", err.Error())
	})

	// Test mutual exclude rule with custom error message
	t.Run("Mutual exclude rule custom error", func(t *testing.T) {
		rule := MutualExclude(
			[]string{"admin", "user"},
			[]string{"admin", "user"},
			func(a, b string) bool { return a == b },
		).Errf("Custom mutual exclude error")

		err := rule.Validate("invalid")
		assert.Equal(t, "Custom mutual exclude error", err.Error())
	})
}
