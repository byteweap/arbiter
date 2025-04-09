// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Common string validation errors
var (
	ErrStartsWith     = errors.New("string must start with the specified prefix")
	ErrEndsWith       = errors.New("string must end with the specified suffix")
	ErrChineseOnly    = errors.New("string must contain only Chinese characters")
	ErrFullWidthOnly  = errors.New("string must contain only full-width characters")
	ErrHalfWidthOnly  = errors.New("string must contain only half-width characters")
	ErrUpperCaseOnly  = errors.New("string must contain only uppercase letters")
	ErrLowerCaseOnly  = errors.New("string must contain only lowercase letters")
	ErrSpecialChars   = errors.New("string contains special characters")
	ErrNoSpecialChars = errors.New("string must not contain special characters")
	ErrContains       = errors.New("string must contain the specified substring")
	ErrNotContains    = errors.New("string must not contain the specified substring")
)

// StartWithRule validates that a string starts with a specific prefix.
//
// Example:
//
//	rule := StartWith("https://")
//	err := rule.Validate("https://example.com")  // returns nil
//	err = rule.Validate("http://example.com")    // returns ErrStartsWith
type StartWithRule struct {
	prefix string
	e      error
}

// StartWith creates a new prefix validation rule.
// The rule checks if a string starts with the specified prefix.
//
// Example:
//
//	// Create a rule for URLs starting with https://
//	urlRule := StartWith("https://").Err("URL must start with https://")
//
//	// Create a rule for file paths
//	pathRule := StartWith("/home/").Err("Path must be absolute")
func StartWith(prefix string) *StartWithRule {
	return &StartWithRule{
		prefix: prefix,
		e:      ErrStartsWith,
	}
}

// Validate checks if the string starts with the specified prefix.
// Returns nil if the string has the prefix, or an error if it doesn't.
//
// Example:
//
//	rule := StartWith("https://")
//	if err := rule.Validate("https://example.com"); err != nil {
//	    // Handle validation error
//	}
func (r *StartWithRule) Validate(value string) error {
	if !strings.HasPrefix(value, r.prefix) {
		if r.e != nil {
			return r.e
		}
		return ErrStartsWith
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := StartWith("https://").Errf("URL must use HTTPS protocol")
//	err := rule.Validate("http://example.com")  // returns error with custom message
func (r *StartWithRule) Errf(format string, args ...any) *StartWithRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// EndWithRule validates that a string ends with a specific suffix.
//
// Example:
//
//	rule := EndWith(".go")
//	err := rule.Validate("main.go")  // returns nil
//	err = rule.Validate("main.py")   // returns ErrEndsWith
type EndWithRule struct {
	suffix string
	e      error
}

// EndWith creates a new suffix validation rule.
// The rule checks if a string ends with the specified suffix.
//
// Example:
//
//	// Create a rule for Go source files
//	goFileRule := EndWith(".go").Err("File must be a Go source file")
//
//	// Create a rule for domain names
//	domainRule := EndWith(".com").Err("Domain must end with .com")
func EndWith(suffix string) *EndWithRule {
	return &EndWithRule{
		suffix: suffix,
		e:      ErrEndsWith,
	}
}

// Validate checks if the string ends with the specified suffix.
// Returns nil if the string has the suffix, or an error if it doesn't.
//
// Example:
//
//	rule := EndWith(".go")
//	if err := rule.Validate("main.go"); err != nil {
//	    // Handle validation error
//	}
func (r *EndWithRule) Validate(value string) error {
	if !strings.HasSuffix(value, r.suffix) {
		if r.e != nil {
			return r.e
		}
		return ErrEndsWith
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := EndWith(".go").Errf("File must be a Go source file")
//	err := rule.Validate("main.py")  // returns error with custom message
func (r *EndWithRule) Errf(format string, args ...any) *EndWithRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// ChineseOnlyRule validates that a string contains only Chinese characters.
//
// Example:
//
//	rule := ChineseOnly()
//	err := rule.Validate("你好世界")  // returns nil
//	err = rule.Validate("Hello")   // returns ErrChineseOnly
type ChineseOnlyRule struct {
	e error
}

// ChineseOnly creates a new Chinese characters validation rule.
// The rule checks if a string contains only Chinese characters.
//
// Example:
//
//	// Create a rule for Chinese names
//	nameRule := ChineseOnly().Err("Name must contain only Chinese characters")
func OnlyChinese() *ChineseOnlyRule {
	return &ChineseOnlyRule{
		e: ErrChineseOnly,
	}
}

// Validate checks if the string contains only Chinese characters.
// Returns nil if all characters are Chinese, or an error if any character is not Chinese.
//
// Example:
//
//	rule := ChineseOnly()
//	if err := rule.Validate("你好世界"); err != nil {
//	    // Handle validation error
//	}
func (r *ChineseOnlyRule) Validate(value string) error {
	for _, char := range value {
		if !unicode.Is(unicode.Han, char) {
			if r.e != nil {
				return r.e
			}
			return ErrChineseOnly
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := ChineseOnly().Errf("Please enter Chinese characters only")
//	err := rule.Validate("Hello")  // returns error with custom message
func (r *ChineseOnlyRule) Errf(format string, args ...any) *ChineseOnlyRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// FullWidthRule validates that a string contains only full-width characters.
//
// Example:
//
//	rule := FullWidthOnly()
//	err := rule.Validate("Ｈｅｌｌｏ")  // returns nil
//	err = rule.Validate("Hello")    // returns ErrFullWidthOnly
type FullWidthRule struct {
	e error
}

// FullWidthOnly creates a new full-width characters validation rule.
// The rule checks if a string contains only full-width characters.
//
// Example:
//
//	// Create a rule for full-width text
//	textRule := FullWidthOnly().Err("Text must be in full-width format")
func OnlyFullWidth() *FullWidthRule {
	return &FullWidthRule{
		e: ErrFullWidthOnly,
	}
}

// Validate checks if the string contains only full-width characters.
// Returns nil if all characters are full-width, or an error if any character is not full-width.
//
// Example:
//
//	rule := FullWidthOnly()
//	if err := rule.Validate("Ｈｅｌｌｏ"); err != nil {
//	    // Handle validation error
//	}
func (r *FullWidthRule) Validate(value string) error {
	for _, char := range value {
		if char < 0xFF01 || char > 0xFF5E {
			if r.e != nil {
				return r.e
			}
			return ErrFullWidthOnly
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := FullWidthOnly().Errf("Please use full-width characters")
//	err := rule.Validate("Hello")  // returns error with custom message
func (r *FullWidthRule) Errf(format string, args ...any) *FullWidthRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// HalfWidthRule validates that a string contains only half-width characters.
//
// Example:
//
//	rule := HalfWidthOnly()
//	err := rule.Validate("Hello")     // returns nil
//	err = rule.Validate("Ｈｅｌｌｏ")  // returns ErrHalfWidthOnly
type HalfWidthRule struct {
	e error
}

// HalfWidthOnly creates a new half-width characters validation rule.
// The rule checks if a string contains only half-width characters.
//
// Example:
//
//	// Create a rule for ASCII text
//	textRule := HalfWidthOnly().Err("Text must be in half-width format")
func OnlyHalfWidth() *HalfWidthRule {
	return &HalfWidthRule{
		e: ErrHalfWidthOnly,
	}
}

// Validate checks if the string contains only half-width characters.
// Returns nil if all characters are half-width, or an error if any character is not half-width.
//
// Example:
//
//	rule := HalfWidthOnly()
//	if err := rule.Validate("Hello"); err != nil {
//	    // Handle validation error
//	}
func (r *HalfWidthRule) Validate(value string) error {
	for _, char := range value {
		if char > 0x7E {
			if r.e != nil {
				return r.e
			}
			return ErrHalfWidthOnly
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := HalfWidthOnly().Errf("Please use half-width characters")
//	err := rule.Validate("Ｈｅｌｌｏ")  // returns error with custom message
func (r *HalfWidthRule) Errf(format string, args ...any) *HalfWidthRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// UpperCaseRule validates that a string contains only uppercase letters.
//
// Example:
//
//	rule := OnlyUpperCase()
//	err := rule.Validate("HELLO")  // returns nil
//	err = rule.Validate("Hello")   // returns ErrUpperCaseOnly
type UpperCaseRule struct {
	e error
}

// OnlyUpperCase creates a new uppercase letters validation rule.
// The rule checks if a string contains only uppercase letters.
//
// Example:
//
//	// Create a rule for uppercase codes
//	codeRule := OnlyUpperCase().Err("Code must be in uppercase")
func OnlyUpperCase() *UpperCaseRule {
	return &UpperCaseRule{
		e: ErrUpperCaseOnly,
	}
}

// Validate checks if the string contains only uppercase letters.
// Returns nil if all letters are uppercase, or an error if any letter is not uppercase.
//
// Example:
//
//	rule := OnlyUpperCase()
//	if err := rule.Validate("HELLO"); err != nil {
//	    // Handle validation error
//	}
func (r *UpperCaseRule) Validate(value string) error {
	for _, char := range value {
		if !unicode.IsUpper(char) {
			if r.e != nil {
				return r.e
			}
			return ErrUpperCaseOnly
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := OnlyUpperCase().Errf("Please use uppercase letters only")
//	err := rule.Validate("Hello")  // returns error with custom message
func (r *UpperCaseRule) Errf(format string, args ...any) *UpperCaseRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// LowerCaseRule validates that a string contains only lowercase letters.
//
// Example:
//
//	rule := OnlyLowerCase()
//	err := rule.Validate("hello")  // returns nil
//	err = rule.Validate("Hello")   // returns ErrLowerCaseOnly
type LowerCaseRule struct {
	e error
}

// OnlyLowerCase creates a new lowercase letters validation rule.
// The rule checks if a string contains only lowercase letters.
//
// Example:
//
//	// Create a rule for lowercase usernames
//	usernameRule := OnlyLowerCase().Err("Username must be in lowercase")
func OnlyLowerCase() *LowerCaseRule {
	return &LowerCaseRule{
		e: ErrLowerCaseOnly,
	}
}

// Validate checks if the string contains only lowercase letters.
// Returns nil if all letters are lowercase, or an error if any letter is not lowercase.
//
// Example:
//
//	rule := OnlyLowerCase()
//	if err := rule.Validate("hello"); err != nil {
//	    // Handle validation error
//	}
func (r *LowerCaseRule) Validate(value string) error {
	for _, char := range value {
		if !unicode.IsLower(char) {
			if r.e != nil {
				return r.e
			}
			return ErrLowerCaseOnly
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := OnlyLowerCase().Errf("Please use lowercase letters only")
//	err := rule.Validate("Hello")  // returns error with custom message
func (r *LowerCaseRule) Errf(format string, args ...any) *LowerCaseRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// SpecialCharsRule validates the presence or absence of special characters in a string.
//
// Example:
//
//	rule := SpecialChars(false)
//	err := rule.Validate("Hello123")  // returns nil
//	err = rule.Validate("Hello@123")  // returns ErrSpecialChars
type SpecialCharsRule struct {
	allowSpecial bool
	e            error
}

// SpecialChars creates a new special characters validation rule.
// If allowSpecial is true, the rule requires special characters.
// If allowSpecial is false, the rule prohibits special characters.
//
// Example:
//
//	// Create a rule that allows special characters
//	passwordRule := SpecialChars(true).Err("Password must contain special characters")
//
//	// Create a rule that prohibits special characters
//	usernameRule := SpecialChars(false).Err("Username must not contain special characters")
func SpecialChars(allowSpecial bool) *SpecialCharsRule {
	var e error
	if allowSpecial {
		e = ErrNoSpecialChars
	} else {
		e = ErrSpecialChars
	}
	return &SpecialCharsRule{
		allowSpecial: allowSpecial,
		e:            e,
	}
}

// Validate checks if the string contains special characters according to the rule.
// Returns nil if the string satisfies the rule, or an error if it doesn't.
//
// Example:
//
//	rule := SpecialChars(false)
//	if err := rule.Validate("Hello123"); err != nil {
//	    // Handle validation error
//	}
func (r *SpecialCharsRule) Validate(value string) error {
	for _, char := range value {
		isSpecial := !unicode.IsLetter(char) && !unicode.IsNumber(char) && !unicode.IsSpace(char)
		if r.allowSpecial && isSpecial {
			if r.e != nil {
				return r.e
			}
			return ErrNoSpecialChars
		}
		if !r.allowSpecial && !isSpecial {
			if r.e != nil {
				return r.e
			}
			return ErrSpecialChars
		}
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := SpecialChars(false).Errf("Please use special characters only")
//	err := rule.Validate("Hello")  // returns error with custom message
func (r *SpecialCharsRule) Errf(format string, args ...any) *SpecialCharsRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// ContainsRule validates that a string contains a specific substring.
//
// Example:
//
//	rule := Contains("foo")
//	err := rule.Validate("foobar")  // returns nil
//	err = rule.Validate("bar")      // returns ErrContains
type ContainsRule struct {
	substring string
	e         error
}

// Contains creates a new substring validation rule.
// The rule checks if a string contains the specified substring.
//
// Example:
//
//	// Create a rule for required words
//	wordRule := Contains("important").Err("Text must contain the word 'important'")
func Contains(substring string) *ContainsRule {
	return &ContainsRule{
		substring: substring,
		e:         ErrContains,
	}
}

// Validate checks if the string contains the specified substring.
// Returns nil if the string contains the substring, or an error if it doesn't.
//
// Example:
//
//	rule := Contains("foo")
//	if err := rule.Validate("foobar"); err != nil {
//	    // Handle validation error
//	}
func (r *ContainsRule) Validate(value string) error {
	if !strings.Contains(value, r.substring) {
		if r.e != nil {
			return r.e
		}
		return ErrContains
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Contains("foo").Errf("Please include 'foo' in the string")
//	err := rule.Validate("bar")  // returns error with custom message
func (r *ContainsRule) Errf(format string, args ...any) *ContainsRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// NotContainsRule validates that a string does not contain a specific substring.
//
// Example:
//
//	rule := NotContains("foo")
//	err := rule.Validate("bar")      // returns nil
//	err = rule.Validate("foobar")    // returns ErrNotContains
type NotContainsRule struct {
	substring string
	e         error
}

// NotContains creates a new substring exclusion validation rule.
// The rule checks if a string does not contain the specified substring.
//
// Example:
//
//	// Create a rule for prohibited words
//	wordRule := NotContains("bad").Err("Text must not contain the word 'bad'")
func NotContains(substring string) *NotContainsRule {
	return &NotContainsRule{
		substring: substring,
		e:         ErrNotContains,
	}
}

// Validate checks if the string does not contain the specified substring.
// Returns nil if the string does not contain the substring, or an error if it does.
//
// Example:
//
//	rule := NotContains("foo")
//	if err := rule.Validate("bar"); err != nil {
//	    // Handle validation error
//	}
func (r *NotContainsRule) Validate(value string) error {
	if strings.Contains(value, r.substring) {
		if r.e != nil {
			return r.e
		}
		return ErrNotContains
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := NotContains("foo").Errf("Please exclude 'foo' from the string")
//	err := rule.Validate("bar")  // returns error with custom message
func (r *NotContainsRule) Errf(format string, args ...any) *NotContainsRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
