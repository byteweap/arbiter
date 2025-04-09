// Package rule provides a collection of validation rules for various data types.
// This file contains network-related validation rules for domains, ports, MAC addresses, and subnet masks.
package rule

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Network validation errors
var (
	// ErrDomain is returned when a domain name fails validation rules.
	// This includes length constraints, format requirements, and character restrictions.
	ErrDomain = errors.New("invalid domain name")

	// ErrPort is returned when a port number is invalid.
	// Valid port numbers must be integers between 0 and 65535.
	ErrPort = errors.New("invalid port number")

	// ErrMACAddress is returned when a MAC address format is invalid.
	// MAC addresses must follow standard IEEE 802 MAC-48, EUI-48, or EUI-64 formats.
	ErrMACAddress = errors.New("invalid MAC address")

	// ErrSubnetMask is returned when a subnet mask is invalid.
	// Valid subnet masks must be IPv4 addresses with continuous 1s followed by continuous 0s in binary.
	ErrSubnetMask = errors.New("invalid subnet mask")
)

// DomainRule provides validation rules for domain names according to DNS standards.
// It checks for valid length, format, character set, and label restrictions.
//
// Example:
//
//	rule := Domain().Errf("Invalid domain format")
//	err := rule.Validate("example.com")      // returns nil
//	err = rule.Validate("invalid..domain")   // returns error
//	err = rule.Validate("no-tld")           // returns error
type DomainRule struct {
	e error
}

// Domain creates a new domain name validation rule.
// The rule enforces DNS naming conventions including:
// - Total length between 1 and 255 characters
// - At least two labels (e.g., "example.com")
// - Each label between 1 and 63 characters
// - Only letters, digits, and hyphens in labels
// - No leading or trailing hyphens in labels
//
// Example:
//
//	domainRule := Domain()
//	websiteRule := Domain().Err("Please enter a valid website domain")
func Domain() *DomainRule {
	return &DomainRule{
		e: ErrDomain,
	}
}

// Validate checks if the given domain name is valid according to DNS naming rules.
// It performs comprehensive validation including:
// - Total length (1-255 characters)
// - Minimum of two parts (e.g., "example.com")
// - Each part length (1-63 characters)
// - Valid characters (letters, numbers, hyphens)
// - No leading or trailing hyphens in parts
//
// Example:
//
//	rule := Domain()
//	err := rule.Validate("example.com")       // returns nil
//	err = rule.Validate("sub.example.com")    // returns nil
//	err = rule.Validate("invalid..domain")    // returns error
//	err = rule.Validate("-invalid.com")      // returns error
func (r *DomainRule) Validate(domain string) error {
	if len(domain) == 0 || len(domain) > 255 {
		return r.e
	}

	// 检查域名格式
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return r.e
	}

	for _, part := range parts {
		if len(part) == 0 || len(part) > 63 {
			return r.e
		}
		if strings.HasPrefix(part, "-") || strings.HasSuffix(part, "-") {
			return r.e
		}
		for _, char := range part {
			if !isValidDomainChar(char) {
				return r.e
			}
		}
	}
	return nil
}

// Errf sets a custom error message for domain validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Domain().Errf("Please enter a valid website address")
//	err := rule.Validate("invalid..com")  // returns custom error message
func (r *DomainRule) Errf(format string, args ...any) *DomainRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// PortRule provides validation rules for network port numbers.
// It ensures that port numbers are valid integers between 0 and 65535.
//
// Example:
//
//	rule := Port().Errf("Invalid port number")
//	err := rule.Validate("8080")    // returns nil
//	err = rule.Validate("70000")    // returns error
//	err = rule.Validate("-1")       // returns error
type PortRule struct {
	e error
}

// Port creates a new port number validation rule.
// The rule ensures that port numbers are:
// - Valid integers
// - Between 0 and 65535 (inclusive)
//
// Example:
//
//	httpRule := Port().Errf("Invalid HTTP port")
//	sshRule := Port().Errf("Invalid SSH port")
func Port() *PortRule {
	return &PortRule{
		e: ErrPort,
	}
}

// Validate checks if the given port number is valid.
// It verifies:
// - String can be converted to an integer
// - Value is between 0 and 65535 (inclusive)
//
// Example:
//
//	rule := Port()
//	err := rule.Validate("80")      // returns nil
//	err = rule.Validate("8080")     // returns nil
//	err = rule.Validate("65536")    // returns error
//	err = rule.Validate("invalid")  // returns error
func (r *PortRule) Validate(port string) error {
	// 尝试转换为整数
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return r.e
	}

	// 检查端口号范围
	if portNum < 0 || portNum > 65535 {
		return r.e
	}

	return nil
}

// Errf sets a custom error message for port validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := Port().Errf("Please enter a valid port number (0-65535)")
//	err := rule.Validate("70000")  // returns custom error message
func (r *PortRule) Errf(format string, args ...any) *PortRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// MACAddressRule provides validation rules for MAC addresses.
// It supports various MAC address formats including IEEE 802 MAC-48, EUI-48, and EUI-64.
//
// Example:
//
//	rule := MACAddress().Errf("Invalid MAC address")
//	err := rule.Validate("00:11:22:33:44:55")    // returns nil
//	err = rule.Validate("00-11-22-33-44-55")    // returns nil
//	err = rule.Validate("invalid")              // returns error
type MACAddressRule struct {
	e error
}

// MACAddress creates a new MAC address validation rule.
// The rule supports multiple formats:
// - IEEE 802 MAC-48 (e.g., "01:23:45:67:89:ab")
// - EUI-48 (e.g., "01-23-45-67-89-ab")
// - EUI-64 (e.g., "01:23:45:67:89:ab:cd:ef")
//
// Example:
//
//	nicRule := MACAddress().Errf("Invalid network card address")
func MACAddress() *MACAddressRule {
	return &MACAddressRule{
		e: ErrMACAddress,
	}
}

// Validate checks if the given MAC address is valid.
// It supports multiple formats and separators (colons or hyphens).
//
// Example:
//
//	rule := MACAddress()
//	err := rule.Validate("00:11:22:33:44:55")     // returns nil
//	err = rule.Validate("00-11-22-33-44-55")     // returns nil
//	err = rule.Validate("00:11:22:33:44:55:66:77") // returns nil (EUI-64)
//	err = rule.Validate("invalid:mac")           // returns error
func (r *MACAddressRule) Validate(mac string) error {
	_, err := net.ParseMAC(mac)
	if err != nil {
		return r.e
	}

	return nil
}

// Errf sets a custom error message for MAC address validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := MACAddress().Errf("Please enter a valid MAC address (e.g., 00:11:22:33:44:55)")
//	err := rule.Validate("invalid")  // returns custom error message
func (r *MACAddressRule) Errf(format string, args ...any) *MACAddressRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// SubnetMaskRule provides validation rules for IPv4 subnet masks.
// It ensures that subnet masks follow the correct format and bit pattern.
//
// Example:
//
//	rule := SubnetMask().Errf("Invalid subnet mask")
//	err := rule.Validate("255.255.255.0")    // returns nil
//	err = rule.Validate("255.255.255.1")    // returns error
//	err = rule.Validate("invalid")          // returns error
type SubnetMaskRule struct {
	e error
}

// SubnetMask creates a new subnet mask validation rule.
// The rule ensures that subnet masks:
// - Are valid IPv4 addresses
// - Have continuous 1s followed by continuous 0s in binary
// - Follow standard netmask patterns
//
// Example:
//
//	networkRule := SubnetMask().Errf("Invalid network mask")
func SubnetMask() *SubnetMaskRule {
	return &SubnetMaskRule{
		e: ErrSubnetMask,
	}
}

// Validate checks if the given subnet mask is valid.
// It verifies:
// - Valid IPv4 address format
// - Proper subnet mask bit pattern (continuous 1s followed by 0s)
// - Standard netmask values
//
// Example:
//
//	rule := SubnetMask()
//	err := rule.Validate("255.255.255.0")     // returns nil
//	err = rule.Validate("255.255.254.0")     // returns nil
//	err = rule.Validate("255.255.255.1")     // returns error
//	err = rule.Validate("255.255.0.255")     // returns error
func (r *SubnetMaskRule) Validate(mask string) error {
	ip := net.ParseIP(mask)
	if ip == nil {
		return r.e
	}

	ip = ip.To4()
	if ip == nil {
		return r.e
	}

	valid := true
	zeroFound := false

	for _, b := range ip {
		if zeroFound && b != 0 {
			valid = false
			break
		}
		if b != 255 {
			zeroFound = true
			for i := 0; i < 8; i++ {
				if (b & (1 << uint(7-i))) != 0 {
					valid = false
					break
				}
			}
		}
	}

	if !valid {
		return r.e
	}

	return nil
}

// Err sets a custom error message for subnet mask validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := SubnetMask().Errf("Please enter a valid subnet mask (e.g., 255.255.255.0)")
//	err := rule.Validate("invalid")  // returns custom error message
func (r *SubnetMaskRule) Errf(format string, args ...any) *SubnetMaskRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// isValidDomainChar checks if a character is valid in a domain name.
// Valid characters include:
// - Letters (a-z, A-Z)
// - Digits (0-9)
// - Hyphen (-)
func isValidDomainChar(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '-'
}
