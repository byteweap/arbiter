// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"errors"
	"fmt"
	"regexp"
	"sync"
)

var (
	// ErrRegex is returned when a string does not match the specified regular expression.
	ErrRegex        = errors.New("does not match the regular expression")
	ErrIDCard       = errors.New("invalid ID card number format")
	ErrPassport     = errors.New("invalid passport number format")
	ErrBankCard     = errors.New("invalid bank card number format")
	ErrTaxNumber    = errors.New("invalid tax number format")
	ErrSocialCredit = errors.New("invalid social credit code format")
	ErrPhone        = errors.New("invalid phone number format")
	ErrEmail        = errors.New("invalid email format")

	// compiledRegexes is a map of compiled regular expressions.
	// It caches compiled regexes to avoid re-compiling the same pattern multiple times.
	compiledRegexes = make(map[string]*regexp.Regexp)
	regexMutex      sync.RWMutex
)

const (
	emailPattern        = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern        = `^\+?[1-9]\d{1,14}$`
	socialCreditPattern = `^[0-9A-HJ-NPQRTUWXY]{2}\d{6}[0-9A-HJ-NPQRTUWXY]{10}$`
	taxNumberPattern    = `^[A-Za-z0-9]{15,20}$`
	bankCardPattern     = `^\d{13,19}$`
	passportPattern     = `^[GEDSP][0-9]{8}$`
	idCardPattern       = `^[1-9]\d{5}(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]$`
)

// getCompiledRegex returns a compiled regular expression for the given pattern.
// It caches compiled regexes to avoid re-compiling the same pattern multiple times.
//
// Example:
//
//	re, err := getCompiledRegex("^[A-Z][a-z]+$")
func getCompiledRegex(pattern string) (*regexp.Regexp, error) {

	regexMutex.RLock()
	defer regexMutex.RUnlock()

	if re, ok := compiledRegexes[pattern]; ok {
		return re, nil
	}
	if re, ok := compiledRegexes[pattern]; ok {
		return re, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	compiledRegexes[pattern] = re
	return re, nil
}

// RegexRule is a validation rule that checks if a string matches a regular expression pattern.
// It can be used for custom pattern matching or with predefined patterns like email and phone.
//
// Example:
//
//	rule := Regex("^[A-Z][a-z]+$")
//	err := rule.Validate("Hello")  // returns nil
//	err = rule.Validate("hello")  // returns ErrRegex
type RegexRule struct {
	regex *regexp.Regexp
	e     error
}

// IsEmail returns a new RegexRule that validates email addresses.
// The rule checks for standard email format: user@domain.com
//
// Example:
//
//	rule := IsEmail()
//	err := rule.Validate("user@example.com")  // returns nil
//	err = rule.Validate("invalid-email")      // returns ErrRegex
func IsEmail() *RegexRule {
	regex, err := getCompiledRegex(emailPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrEmail}
}

// IsPhone returns a new RegexRule that validates phone numbers.
// The rule checks for international phone number format.
//
// Example:
//
//	rule := IsPhone()
//	err := rule.Validate("+8612345678901")  // returns nil
//	err = rule.Validate("123-456-7890")     // returns ErrRegex
func IsPhone() *RegexRule {
	regex, err := getCompiledRegex(phonePattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrPhone}
}

// IsSocialCredit returns a new RegexRule that validates social credit codes.
// The rule checks for the standard format of social credit codes.
//
// Example:
//
//	rule := IsSocialCredit()
//	err := rule.Validate("123456789012345678")  // returns nil
func IsSocialCredit() *RegexRule {
	regex, err := getCompiledRegex(socialCreditPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrSocialCredit}
}

// IsTaxNumber returns a new RegexRule that validates tax numbers.
// The rule checks for the standard format of tax numbers.
//
// Example:
//
//	rule := IsTaxNumber()
//	err := rule.Validate("123456789012345678")  // returns nil
func IsTaxNumber() *RegexRule {
	regex, err := getCompiledRegex(taxNumberPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrTaxNumber}
}

// IsBankCard returns a new RegexRule that validates bank card numbers.
// The rule checks for the standard format of bank card numbers.
//
// Example:
//
//	rule := IsBankCard()
//	err := rule.Validate("123456789012345678")  // returns nil
func IsBankCard() *RegexRule {
	regex, err := getCompiledRegex(bankCardPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrBankCard}
}

// IsPassport returns a new RegexRule that validates passport numbers.
// The rule checks for the standard format of passport numbers.
//
// Example:
//
//	rule := IsPassport()
//	err := rule.Validate("G12345678")  // returns nil
//	err = rule.Validate("12345678")  // returns ErrRegex
func IsPassport() *RegexRule {
	regex, err := getCompiledRegex(passportPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrPassport}
}

// IsIDCard returns a new RegexRule that validates ID card numbers.
// The rule checks for the standard format of ID card numbers.
//
// Example:
//
//	rule := IsIDCard()
//	err := rule.Validate("123456789012345678")  // returns nil
func IsIDCard() *RegexRule {
	regex, err := getCompiledRegex(idCardPattern)
	if err != nil {
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{regex: regex, e: ErrIDCard}
}

// Regex creates a new RegexRule with a custom regular expression pattern.
// If the pattern is invalid, the rule will always return an error.
//
// Example:
//
//	// Create a rule for usernames (letters and numbers only)
//	rule := Regex("^[a-zA-Z0-9]+$")
//	err := rule.Validate("user123")  // returns nil
//	err = rule.Validate("user@123")  // returns ErrRegex
func Regex(pattern string) *RegexRule {
	regex, err := getCompiledRegex(pattern)
	if err != nil {
		// If pattern compilation fails, return a rule that always returns an error
		return &RegexRule{
			regex: nil,
			e:     fmt.Errorf("invalid regular expression: %w", err),
		}
	}
	return &RegexRule{
		regex: regex,
		e:     ErrRegex,
	}
}

// Validate checks if the string matches the regular expression pattern.
// Returns nil if the string matches, or an error if it doesn't.
//
// Example:
//
//	rule := Regex("^[A-Z][a-z]+$")
//	if err := rule.Validate("Hello"); err != nil {
//	    // Handle validation error
//	}
func (r *RegexRule) Validate(value string) error {
	if r.regex == nil {
		return r.e
	}
	if !r.regex.MatchString(value) {
		if r.e != nil {
			return r.e
		}
		return ErrRegex
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Regex("^[A-Z][a-z]+$").Errf("Name must start with a capital letter")
//	err := rule.Validate("hello")  // returns error with message "Name must start with a capital letter"
func (r *RegexRule) Errf(format string, args ...any) *RegexRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
