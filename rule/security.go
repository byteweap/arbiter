// Package rule provides a collection of validation rules for various data types.
// This file contains security-related validation rules for passwords, XSS protection, and SQL injection prevention.
package rule

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Security validation errors
var (
	// ErrPasswordStrength is returned when a password does not meet the strength requirements.
	// This includes length, character types, and other password policy requirements.
	ErrPasswordStrength = errors.New("password does not meet strength requirements")

	// ErrPasswordComplex is returned when a password does not meet the complexity requirements.
	// This includes minimum length, character variety, and forbidden patterns.
	ErrPasswordComplex = errors.New("password does not meet complexity requirements")

	// ErrXSS is returned when input contains potential XSS (Cross-Site Scripting) attack patterns.
	// This helps prevent malicious script injection in web applications.
	ErrXSS = errors.New("input contains potential XSS attack")

	// ErrSQLInjection is returned when input contains potential SQL injection attack patterns.
	// This helps prevent malicious SQL query manipulation in database operations.
	ErrSQLInjection = errors.New("input contains potential SQL injection")
)

// PasswordStrengthRule validates that a password meets strength requirements.
// The rule checks for minimum and maximum length, and required character types.
//
// Example:
//
//	rule := PasswordStrength().
//		MinLength(10).
//		MaxLength(30).
//		RequireUpper(true).
//		RequireLower(true).
//		RequireNumber(true).
//		RequireSpecial(true).
//		Err("Password must be 10-30 characters with uppercase, lowercase, numbers, and special characters")
//	err := rule.Validate("StrongP@ssw0rd")  // returns nil
//	err = rule.Validate("weak")             // returns error
type PasswordStrengthRule struct {
	e              error
	minLength      int
	maxLength      int
	requireUpper   bool
	requireLower   bool
	requireNumber  bool
	requireSpecial bool
}

// PasswordStrength creates a new password strength validation rule.
// By default, the rule requires:
// - Minimum length: 8 characters
// - Maximum length: 32 characters
// - Uppercase letters: required
// - Lowercase letters: required
// - Numbers: required
// - Special characters: required
//
// Example:
//
//	rule := PasswordStrength()  // use default settings
//	rule := PasswordStrength().
//		MinLength(12).
//		MaxLength(24).
//		RequireSpecial(false)  // customize settings
func PasswordStrength() *PasswordStrengthRule {
	return &PasswordStrengthRule{
		e:              ErrPasswordStrength,
		minLength:      8,
		maxLength:      32,
		requireUpper:   true,
		requireLower:   true,
		requireNumber:  true,
		requireSpecial: true,
	}
}

// Validate checks if the given password meets the strength requirements.
// Empty strings are considered valid (use Required() if needed).
//
// Example:
//
//	rule := PasswordStrength()
//	err := rule.Validate("StrongP@ssw0rd")  // returns nil
//	err = rule.Validate("weak")             // returns error
//	err = rule.Validate("")                 // returns nil (empty string is valid)
func (r *PasswordStrengthRule) Validate(value string) error {
	if value == "" {
		return nil
	}

	length := len(value)
	if length < r.minLength || length > r.maxLength {
		if r.e != nil {
			return r.e
		}
		return ErrPasswordStrength
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if (r.requireUpper && !hasUpper) ||
		(r.requireLower && !hasLower) ||
		(r.requireNumber && !hasNumber) ||
		(r.requireSpecial && !hasSpecial) {
		if r.e != nil {
			return r.e
		}
		return ErrPasswordStrength
	}

	return nil
}

// MinLength sets the minimum required length for the password.
//
// Example:
//
//	rule := PasswordStrength().MinLength(12)  // password must be at least 12 characters
func (r *PasswordStrengthRule) MinLength(length int) *PasswordStrengthRule {
	r.minLength = length
	return r
}

// MaxLength sets the maximum allowed length for the password.
//
// Example:
//
//	rule := PasswordStrength().MaxLength(24)  // password must be no more than 24 characters
func (r *PasswordStrengthRule) MaxLength(length int) *PasswordStrengthRule {
	r.maxLength = length
	return r
}

// RequireUpper sets whether uppercase letters are required in the password.
//
// Example:
//
//	rule := PasswordStrength().RequireUpper(true)  // password must contain at least one uppercase letter
func (r *PasswordStrengthRule) RequireUpper(require bool) *PasswordStrengthRule {
	r.requireUpper = require
	return r
}

// RequireLower sets whether lowercase letters are required in the password.
//
// Example:
//
//	rule := PasswordStrength().RequireLower(true)  // password must contain at least one lowercase letter
func (r *PasswordStrengthRule) RequireLower(require bool) *PasswordStrengthRule {
	r.requireLower = require
	return r
}

// RequireNumber sets whether numbers are required in the password.
//
// Example:
//
//	rule := PasswordStrength().RequireNumber(true)  // password must contain at least one number
func (r *PasswordStrengthRule) RequireNumber(require bool) *PasswordStrengthRule {
	r.requireNumber = require
	return r
}

// RequireSpecial sets whether special characters are required in the password.
//
// Example:
//
//	rule := PasswordStrength().RequireSpecial(true)  // password must contain at least one special character
func (r *PasswordStrengthRule) RequireSpecial(require bool) *PasswordStrengthRule {
	r.requireSpecial = require
	return r
}

// Errf sets a custom error message for password strength validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := PasswordStrength().Errf("Password must be at least 8 characters with mixed case, numbers, and special characters")
func (r *PasswordStrengthRule) Errf(format string, args ...any) *PasswordStrengthRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// PasswordComplexRule validates that a password meets complexity requirements.
// The rule checks for minimum length, character variety, repeated characters, and forbidden patterns.
//
// Example:
//
//	rule := PasswordComplex().
//		MinLength(12).
//		MinCharTypes(3).
//		MaxRepeatedChars(2).
//		AddForbiddenPattern("password").
//		Err("Password must be complex and not contain common words")
//	err := rule.Validate("C0mpl3x!P@ss")  // returns nil
//	err = rule.Validate("password123")     // returns error
type PasswordComplexRule struct {
	e                 error
	minLength         int
	minCharTypes      int
	maxRepeatedChars  int
	forbiddenPatterns []string
}

// PasswordComplex creates a new password complexity validation rule.
// By default, the rule requires:
// - Minimum length: 12 characters
// - Minimum character types: 3 (uppercase, lowercase, numbers, special)
// - Maximum repeated characters: 3
// - Forbidden patterns: "password", "123456", "qwerty", "admin"
//
// Example:
//
//	rule := PasswordComplex()  // use default settings
//	rule := PasswordComplex().
//		MinLength(16).
//		MinCharTypes(4).
//		MaxRepeatedChars(2)  // customize settings
func PasswordComplex() *PasswordComplexRule {
	return &PasswordComplexRule{
		e:                ErrPasswordComplex,
		minLength:        12,
		minCharTypes:     3,
		maxRepeatedChars: 3,
		forbiddenPatterns: []string{
			"password",
			"123456",
			"qwerty",
			"admin",
		},
	}
}

// Validate checks if the given password meets the complexity requirements.
// Empty strings are considered valid (use Required() if needed).
//
// Example:
//
//	rule := PasswordComplex()
//	err := rule.Validate("C0mpl3x!P@ss")  // returns nil
//	err = rule.Validate("password123")     // returns error
//	err = rule.Validate("")                // returns nil (empty string is valid)
func (r *PasswordComplexRule) Validate(value string) error {
	if value == "" {
		return nil
	}

	if len(value) < r.minLength {
		if r.e != nil {
			return r.e
		}
		return ErrPasswordComplex
	}

	// Check character type count
	var charTypes int
	if regexp.MustCompile(`[A-Z]`).MatchString(value) {
		charTypes++
	}
	if regexp.MustCompile(`[a-z]`).MatchString(value) {
		charTypes++
	}
	if regexp.MustCompile(`[0-9]`).MatchString(value) {
		charTypes++
	}
	if regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(value) {
		charTypes++
	}

	if charTypes < r.minCharTypes {
		if r.e != nil {
			return r.e
		}
		return ErrPasswordComplex
	}

	// Check repeated characters
	for i := 0; i <= len(value)-r.maxRepeatedChars; i++ {
		if strings.Count(value, value[i:i+1]) > r.maxRepeatedChars {
			if r.e != nil {
				return r.e
			}
			return ErrPasswordComplex
		}
	}

	// Check forbidden patterns
	valueLower := strings.ToLower(value)
	for _, pattern := range r.forbiddenPatterns {
		if strings.Contains(valueLower, strings.ToLower(pattern)) {
			if r.e != nil {
				return r.e
			}
			return ErrPasswordComplex
		}
	}

	return nil
}

// MinLength sets the minimum required length for the password.
//
// Example:
//
//	rule := PasswordComplex().MinLength(16)  // password must be at least 16 characters
func (r *PasswordComplexRule) MinLength(length int) *PasswordComplexRule {
	r.minLength = length
	return r
}

// MinCharTypes sets the minimum number of character types required in the password.
// Character types are: uppercase letters, lowercase letters, numbers, and special characters.
//
// Example:
//
//	rule := PasswordComplex().MinCharTypes(4)  // password must contain all four character types
func (r *PasswordComplexRule) MinCharTypes(count int) *PasswordComplexRule {
	r.minCharTypes = count
	return r
}

// MaxRepeatedChars sets the maximum number of times a character can be repeated in the password.
//
// Example:
//
//	rule := PasswordComplex().MaxRepeatedChars(2)  // no character can appear more than twice
func (r *PasswordComplexRule) MaxRepeatedChars(count int) *PasswordComplexRule {
	r.maxRepeatedChars = count
	return r
}

// AddForbiddenPattern adds a pattern to the list of forbidden patterns in the password.
//
// Example:
//
//	rule := PasswordComplex().AddForbiddenPattern("qwerty")  // password cannot contain "qwerty"
func (r *PasswordComplexRule) AddForbiddenPattern(pattern string) *PasswordComplexRule {
	r.forbiddenPatterns = append(r.forbiddenPatterns, pattern)
	return r
}

// Errf sets a custom error message for password complexity validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := PasswordComplex().Errf("Password must be complex and not contain common words")
func (r *PasswordComplexRule) Errf(format string, args ...any) *PasswordComplexRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// XSSRule validates that input does not contain potential XSS (Cross-Site Scripting) attack patterns.
// This helps prevent malicious script injection in web applications.
//
// Example:
//
//	rule := XSS().Err("Input contains potentially dangerous content")
//	err := rule.Validate("Hello, world!")  // returns nil
//	err = rule.Validate("<script>alert('XSS')</script>")  // returns error
type XSSRule struct {
	e error
}

// XSS creates a new XSS protection validation rule.
// The rule checks for common XSS attack patterns in the input.
//
// Example:
//
//	rule := XSS()  // creates a rule that checks for XSS patterns
//	rule := XSS().Err("Input contains potentially dangerous content")  // with custom error message
func XSS() *XSSRule {
	return &XSSRule{
		e: ErrXSS,
	}
}

// Validate checks if the given input contains potential XSS attack patterns.
// Empty strings are considered valid (use Required() if needed).
//
// Example:
//
//	rule := XSS()
//	err := rule.Validate("Hello, world!")  // returns nil
//	err = rule.Validate("<script>alert('XSS')</script>")  // returns error
//	err = rule.Validate("")  // returns nil (empty string is valid)
func (r *XSSRule) Validate(value string) error {
	if value == "" {
		return nil
	}

	// Check for common XSS attack patterns
	patterns := []string{
		`<script[^>]*>.*?</script>`,
		`javascript:`,
		`vbscript:`,
		`onload=`,
		`onerror=`,
		`onclick=`,
		`onmouseover=`,
		`eval\(.*\)`,
		`expression\(.*\)`,
		`<iframe[^>]*>`,
		`<img[^>]*>`,
		`<embed[^>]*>`,
		`<object[^>]*>`,
		`<style[^>]*>.*?</style>`,
	}

	for _, pattern := range patterns {
		if regexp.MustCompile(`(?i)` + pattern).MatchString(value) {
			if r.e != nil {
				return r.e
			}
			return ErrXSS
		}
	}

	return nil
}

// Errf sets a custom error message for XSS validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := XSS().Errf("Input contains potentially dangerous content")
func (r *XSSRule) Errf(format string, args ...any) *XSSRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// SQLInjectionRule validates that input does not contain potential SQL injection attack patterns.
// This helps prevent malicious SQL query manipulation in database operations.
//
// Example:
//
//	rule := SQLInjection().Errf("Input contains potentially dangerous SQL content")
//	err := rule.Validate("John Doe")  // returns nil
//	err = rule.Validate("'; DROP TABLE users; --")  // returns error
type SQLInjectionRule struct {
	e error
}

// SQLInjection creates a new SQL injection protection validation rule.
// The rule checks for common SQL injection attack patterns in the input.
//
// Example:
//
//	rule := SQLInjection()  // creates a rule that checks for SQL injection patterns
//	rule := SQLInjection().Err("Input contains potentially dangerous SQL content")  // with custom error message
func SQLInjection() *SQLInjectionRule {
	return &SQLInjectionRule{
		e: ErrSQLInjection,
	}
}

// Validate checks if the given input contains potential SQL injection attack patterns.
// Empty strings are considered valid (use Required() if needed).
//
// Example:
//
//	rule := SQLInjection()
//	err := rule.Validate("John Doe")  // returns nil
//	err = rule.Validate("'; DROP TABLE users; --")  // returns error
//	err = rule.Validate("")  // returns nil (empty string is valid)
func (r *SQLInjectionRule) Validate(value string) error {
	if value == "" {
		return nil
	}

	// Check for common SQL injection attack patterns
	patterns := []string{
		`(?i)(select|insert|update|delete|drop|union|exec|execute)\s+`,
		`(?i)(\s+and\s+|\s+or\s+)[\d'"]`,
		`(?i)(\s+xor\s+|\s+nand\s+|\s+not\s+)[\d'"]`,
		`(?i)(\s+like\s+|\s+between\s+|\s+in\s+)[\d'"]`,
		`(?i)(\s+is\s+null|\s+is\s+not\s+null)`,
		`(?i)(--|\#|\*|;)$`,
		`(?i)'(\s*)(union|select|or|and)`,
		`(?i)\/\*.*\*\/`,
		`(?i)waitfor\s+delay\s+`,
		`(?i)benchmark\(.*\)`,
		`(?i)sleep\(.*\)`,
	}

	valueLower := strings.ToLower(value)
	for _, pattern := range patterns {
		if regexp.MustCompile(pattern).MatchString(valueLower) {
			if r.e != nil {
				return r.e
			}
			return ErrSQLInjection
		}
	}

	return nil
}

// Errf sets a custom error message for SQL injection validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := SQLInjection().Errf("Input contains potentially dangerous SQL content")
func (r *SQLInjectionRule) Errf(format string, args ...any) *SQLInjectionRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
