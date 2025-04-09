// Package rule provides a collection of validation rules for various data types.
// This file contains URL validation rules for validating web URLs.
package rule

import (
	"errors"
	"fmt"
	"net/url"
)

// ErrURL is returned when a string is not a valid URL.
// The URL must be properly formatted with a scheme (e.g., http://, https://).
var ErrURL = errors.New("invalid URL format")

// URLRule validates that a string is a valid URL.
// The rule uses url.ParseRequestURI to verify the URL format.
//
// Example:
//
//	rule := URL().Err("Please enter a valid URL")
//	err := rule.Validate("https://example.com")  // returns nil
//	err = rule.Validate("not-a-url")            // returns error
type URLRule struct {
	e error
}

// URL creates a new URL validation rule.
// The rule ensures that a string is a valid URL with a scheme.
//
// Example:
//
//	rule := URL()  // creates a rule that validates URLs
//	rule := URL().Err("Invalid URL format")  // with custom error message
func URL() *URLRule {
	return &URLRule{
		e: ErrURL,
	}
}

// Validate checks if the given string is a valid URL.
// Empty strings are considered valid (use Required() if needed).
// The validation uses url.ParseRequestURI to verify the URL format.
//
// Example:
//
//	rule := URL()
//	err := rule.Validate("https://example.com")  // returns nil
//	err = rule.Validate("http://localhost:8080") // returns nil
//	err = rule.Validate("not-a-url")            // returns error
//	err = rule.Validate("")                     // returns nil (empty string is valid)
func (r *URLRule) Validate(value string) error {
	if value == "" {
		return nil
	}
	_, err := url.ParseRequestURI(value)
	if err != nil {
		if r.e != nil {
			return r.e
		}
		return ErrURL
	}
	return nil
}

// Errf sets a custom error message for URL validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := URL().Errf("Please enter a valid URL including the scheme (e.g., https://)")
func (r *URLRule) Errf(format string, args ...any) *URLRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
