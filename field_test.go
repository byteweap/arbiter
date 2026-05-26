// Package arbiter_test provides tests for the arbiter package.
// These tests verify the functionality of field validation rules.
package arbiter_test

import (
	"testing"

	"github.com/byteweap/arbiter"
	"github.com/byteweap/arbiter/rule"
)

type testUser struct {
	Username string
	Age      int
	Website  string
	Password string
	Chinese  string
}

func TestFieldValid(t *testing.T) {
	user := &testUser{
		Username: "johndoe",
		Age:      25,
		Website:  "example.com",
		Password: "StrongP@ssw0rd",
		Chinese:  "你好世界",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid username, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
	)
	if err != nil {
		t.Errorf("Expected no error for valid age, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Website, rule.Domain()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid website, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Password, rule.PasswordStrength()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid password, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid Chinese text, got %v", err)
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
		arbiter.Field(&user.Website, rule.Domain()),
		arbiter.Field(&user.Password, rule.PasswordStrength()),
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err != nil {
		t.Errorf("Expected no error for valid user, got %v", err)
	}
}

func TestFieldInvalid(t *testing.T) {
	user := &testUser{
		Username: "Ｊｏｈｎ",
		Age:      -1,
		Website:  "invalid",
		Password: "weak",
		Chinese:  "Hello",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
	)
	if err == nil {
		t.Error("Expected error for full-width username, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age, rule.Min[int](0), rule.Max[int](120)),
	)
	if err == nil {
		t.Error("Expected error for negative age, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Website, rule.Domain()),
	)
	if err == nil {
		t.Error("Expected error for invalid website domain, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Password, rule.PasswordStrength()),
	)
	if err == nil {
		t.Error("Expected error for weak password, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err == nil {
		t.Error("Expected error for non-Chinese text, got nil")
	}

	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username, rule.HalfWidthOnly()),
		arbiter.Field(&user.Age, rule.Min(0), rule.Max(120)),
		arbiter.Field(&user.Website, rule.Domain()),
		arbiter.Field(&user.Password, rule.PasswordStrength()),
		arbiter.Field(&user.Chinese, rule.ChineseOnly()),
	)
	if err == nil {
		t.Error("Expected error for invalid user, got nil")
	}
}

func TestFieldNilPointer(t *testing.T) {
	err := arbiter.ValidateStruct(nil, "User cannot be nil")
	if err == nil || err.Error() != "User cannot be nil" {
		t.Errorf("Expected error 'User cannot be nil', got %v", err)
	}
}

func TestFieldMultipleRules(t *testing.T) {
	user := &testUser{
		Username: "johndoe",
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username,
			rule.HalfWidthOnly(),
			rule.LowerCaseOnly(),
		),
	)
	if err != nil {
		t.Errorf("Expected no error for valid username with multiple rules, got %v", err)
	}

	user.Username = "JOHN123"
	err = arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Username,
			rule.HalfWidthOnly(),
			rule.LowerCaseOnly(),
		),
	)
	if err == nil {
		t.Error("Expected error for uppercase username, got nil")
	}
}

func TestFieldCustomError(t *testing.T) {
	user := &testUser{
		Age: -1,
	}

	err := arbiter.ValidateStruct(user, "User cannot be nil",
		arbiter.Field(&user.Age,
			rule.Min[int](0).Errf("Age cannot be negative"),
		),
	)
	if err == nil || err.Error() != "Age cannot be negative" {
		t.Errorf("Expected custom error message, got %v", err)
	}
}
