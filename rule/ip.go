// Package rule provides a collection of validation rules for various data types.
// This file contains IP address validation rules for general IP, IPv4, and IPv6 addresses.
package rule

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var (
	// ErrIP is returned when a string is not a valid IP address.
	// This error is used for general IP address validation.
	ErrIP = errors.New("invalid IP address format")

	// ErrIPv4 is returned when a string is not a valid IPv4 address.
	// The address must be in the format x.x.x.x where x is a number between 0 and 255.
	ErrIPv4 = errors.New("invalid IPv4 address format")

	// ErrIPv6 is returned when a string is not a valid IPv6 address.
	// The address must be in the standard IPv6 format with hexadecimal numbers.
	ErrIPv6 = errors.New("invalid IPv6 address format")
)

// IPRule validates that a string is a valid IP address (either IPv4 or IPv6).
// The rule uses net.ParseIP to verify the IP address format.
//
// Example:
//
//	rule := IP().Err("Please enter a valid IP address")
//	err := rule.Validate("192.168.1.1")     // returns nil
//	err = rule.Validate("2001:db8::1")     // returns nil
//	err = rule.Validate("invalid")         // returns error
type IPRule struct {
	e error
}

// IP creates a new IP address validation rule.
// The rule ensures that a string is a valid IP address.
//
// Example:
//
//	rule := IP()  // creates a rule that validates any IP address
//	rule := IP().Err("Invalid IP address format")  // with custom error message
func IP() *IPRule {
	return &IPRule{
		e: ErrIP,
	}
}

// Validate checks if the given string is a valid IP address.
// Empty strings are considered valid (use Required() if needed).
// The validation uses net.ParseIP to verify the IP address format.
//
// Example:
//
//	rule := IP()
//	err := rule.Validate("192.168.1.1")     // returns nil
//	err = rule.Validate("2001:db8::1")     // returns nil
//	err = rule.Validate("invalid")         // returns error
//	err = rule.Validate("")                // returns nil (empty string is valid)
func (r *IPRule) Validate(value string) error {
	if value == "" {
		return nil
	}
	if net.ParseIP(value) == nil {
		if r.e != nil {
			return r.e
		}
		return ErrIP
	}
	return nil
}

// Errf sets a custom error message for IP validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := IP().Errf("Please enter a valid IP address")
func (r *IPRule) Errf(format string, args ...any) *IPRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// IPv4Rule validates that a string is a valid IPv4 address.
// The rule ensures the address is in the format x.x.x.x where x is a number between 0 and 255.
//
// Example:
//
//	rule := IPv4().Err("Please enter a valid IPv4 address")
//	err := rule.Validate("192.168.1.1")  // returns nil
//	err = rule.Validate("2001:db8::1")  // returns error (not IPv4)
//	err = rule.Validate("invalid")      // returns error
type IPv4Rule struct {
	e error
}

// IPv4 creates a new IPv4 address validation rule.
// The rule ensures that a string is a valid IPv4 address.
//
// Example:
//
//	rule := IPv4()  // creates a rule that validates IPv4 addresses
//	rule := IPv4().Err("Invalid IPv4 address format")  // with custom error message
func IPv4() *IPv4Rule {
	return &IPv4Rule{
		e: ErrIPv4,
	}
}

// Validate checks if the given string is a valid IPv4 address.
// Empty strings are considered valid (use Required() if needed).
// The validation ensures the address contains dots and is a valid IP.
//
// Example:
//
//	rule := IPv4()
//	err := rule.Validate("192.168.1.1")  // returns nil
//	err = rule.Validate("2001:db8::1")  // returns error (not IPv4)
//	err = rule.Validate("invalid")      // returns error
//	err = rule.Validate("")             // returns nil (empty string is valid)
func (r *IPv4Rule) Validate(value string) error {
	if value == "" {
		return nil
	}
	ip := net.ParseIP(value)
	if ip == nil || !strings.Contains(value, ".") {
		if r.e != nil {
			return r.e
		}
		return ErrIPv4
	}
	return nil
}

// Errf sets a custom error message for IPv4 validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := IPv4().Errf("Please enter a valid IPv4 address")
func (r *IPv4Rule) Errf(format string, args ...any) *IPv4Rule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// IPv6Rule validates that a string is a valid IPv6 address.
// The rule ensures the address is in the standard IPv6 format with hexadecimal numbers.
//
// Example:
//
//	rule := IPv6().Err("Please enter a valid IPv6 address")
//	err := rule.Validate("2001:db8::1")     // returns nil
//	err = rule.Validate("192.168.1.1")     // returns error (not IPv6)
//	err = rule.Validate("invalid")         // returns error
type IPv6Rule struct {
	e error
}

// IPv6 creates a new IPv6 address validation rule.
// The rule ensures that a string is a valid IPv6 address.
//
// Example:
//
//	rule := IPv6()  // creates a rule that validates IPv6 addresses
//	rule := IPv6().Err("Invalid IPv6 address format")  // with custom error message
func IPv6() *IPv6Rule {
	return &IPv6Rule{
		e: ErrIPv6,
	}
}

// Validate checks if the given string is a valid IPv6 address.
// Empty strings are considered valid (use Required() if needed).
// The validation ensures the address contains colons and is a valid IP.
//
// Example:
//
//	rule := IPv6()
//	err := rule.Validate("2001:db8::1")     // returns nil
//	err = rule.Validate("192.168.1.1")     // returns error (not IPv6)
//	err = rule.Validate("invalid")         // returns error
//	err = rule.Validate("")                // returns nil (empty string is valid)
func (r *IPv6Rule) Validate(value string) error {
	if value == "" {
		return nil
	}
	ip := net.ParseIP(value)
	if ip == nil || !strings.Contains(value, ":") {
		if r.e != nil {
			return r.e
		}
		return ErrIPv6
	}
	return nil
}

// Errf sets a custom error message for IPv6 validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := IPv6().Errf("Please enter a valid IPv6 address")
func (r *IPv6Rule) Errf(format string, args ...any) *IPv6Rule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
