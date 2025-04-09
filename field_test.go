// Package arbiter_test provides tests for the arbiter package.
// These tests verify the functionality of field validation rules.
package arbiter_test

import (
	"testing"

	"github.com/byteweap/arbiter"
	"github.com/byteweap/arbiter/rule"
)

// TestField tests the Field function and field validation rules.
// It verifies that field validation rules are correctly applied to struct fields.
func TestField(t *testing.T) {
	// Test struct with various field types
	type User struct {
		Username string
		Age      int
		Website  string
		Password string
		Chinese  string
	}

	// Test valid field values
	t.Run("valid fields", func(t *testing.T) {
		user := &User{
			Username: "johndoe",
			Age:      25,
			Website:  "example.com",
			Password: "StrongP@ssw0rd",
			Chinese:  "你好世界",
		}

		// Test username validation (half-width characters)
		err := arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username, rule.OnlyHalfWidth()),
		)
		if err != nil {
			t.Errorf("Expected no error for valid username, got %v", err)
		}

		// Test age validation (between 0 and 120)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
		)
		if err != nil {
			t.Errorf("Expected no error for valid age, got %v", err)
		}

		// Test website validation (valid domain)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Website, rule.Domain()),
		)
		if err != nil {
			t.Errorf("Expected no error for valid website, got %v", err)
		}

		// Test password validation (strong password)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Password, rule.PasswordStrength()),
		)
		if err != nil {
			t.Errorf("Expected no error for valid password, got %v", err)
		}

		// Test Chinese text validation
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Chinese, rule.OnlyChinese()),
		)
		if err != nil {
			t.Errorf("Expected no error for valid Chinese text, got %v", err)
		}

		// Test all fields together
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username, rule.OnlyHalfWidth()),
			arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
			arbiter.Field(&user.Website, rule.Domain()),
			arbiter.Field(&user.Password, rule.PasswordStrength()),
			arbiter.Field(&user.Chinese, rule.OnlyChinese()),
		)
		if err != nil {
			t.Errorf("Expected no error for valid user, got %v", err)
		}
	})

	// Test invalid field values
	t.Run("invalid fields", func(t *testing.T) {
		user := &User{
			Username: "Ｊｏｈｎ",    // Full-width characters
			Age:      -1,        // Negative age
			Website:  "invalid", // Invalid domain
			Password: "weak",    // Weak password
			Chinese:  "Hello",   // Non-Chinese characters
		}

		// Test username validation (should fail for full-width characters)
		err := arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username, rule.OnlyHalfWidth()),
		)
		if err == nil {
			t.Error("Expected error for full-width username, got nil")
		}

		// Test age validation (should fail for negative value)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
		)
		if err == nil {
			t.Error("Expected error for negative age, got nil")
		}

		// Test website validation (should fail for invalid domain)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Website, rule.Domain()),
		)
		if err == nil {
			t.Error("Expected error for invalid website domain, got nil")
		}

		// Test password validation (should fail for weak password)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Password, rule.PasswordStrength()),
		)
		if err == nil {
			t.Error("Expected error for weak password, got nil")
		}

		// Test Chinese text validation (should fail for non-Chinese characters)
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Chinese, rule.OnlyChinese()),
		)
		if err == nil {
			t.Error("Expected error for non-Chinese text, got nil")
		}

		// Test all fields together
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username, rule.OnlyHalfWidth()),
			arbiter.Field(&user.Age, rule.Min(0), rule.Max(120)),
			arbiter.Field(&user.Website, rule.Domain()),
			arbiter.Field(&user.Password, rule.PasswordStrength()),
			arbiter.Field(&user.Chinese, rule.OnlyChinese()),
		)
		if err == nil {
			t.Error("Expected error for invalid user, got nil")
		}
	})

	// Test nil pointer struct
	t.Run("nil pointer struct", func(t *testing.T) {
		var user *User = nil

		// Test validation of nil pointer struct
		err := arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username, rule.OnlyHalfWidth()),
		)
		if err == nil || err.Error() != "User cannot be nil" {
			t.Errorf("Expected error 'User cannot be nil', got %v", err)
		}
	})

	// Test multiple rules on a single field
	t.Run("multiple rules", func(t *testing.T) {
		user := &User{
			Username: "johndoe",
		}

		// Test multiple validation rules on username
		err := arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username,
				rule.OnlyHalfWidth(),
				rule.OnlyLowerCase(),
			),
		)
		if err != nil {
			t.Errorf("Expected no error for valid username with multiple rules, got %v", err)
		}

		// Test failing multiple rules
		user.Username = "JOHN123"
		err = arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Username,
				rule.OnlyHalfWidth(),
				rule.OnlyLowerCase(),
			),
		)
		if err == nil {
			t.Error("Expected error for uppercase username, got nil")
		}
	})

	// Test custom error messages
	t.Run("custom error messages", func(t *testing.T) {
		user := &User{
			Age: -1,
		}

		// Test validation with custom error message
		err := arbiter.ValidateStruct(user, "User cannot be nil",
			arbiter.Field(&user.Age,
				rule.Min[int](0).Errf("Age cannot be negative"),
			),
		)
		if err == nil || err.Error() != "Age cannot be negative" {
			t.Errorf("Expected custom error message, got %v", err)
		}
	})
}
