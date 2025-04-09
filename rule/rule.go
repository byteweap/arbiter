// Package rule provides a collection of validation rules for various data types.
// It offers a flexible and extensible way to validate data using a common interface.
package rule

// Rule is the core interface that all validation rules must implement.
// It provides a generic type parameter T to support validation of any data type.
//
// Example:
//
//	type CustomRule struct{}
//	func (r *CustomRule) Validate(value string) error {
//	    if value == "" {
//	        return errors.New("value cannot be empty")
//	    }
//	    return nil
//	}
type Rule[T any] interface {
	// Validate checks if the provided value meets the rule's criteria.
	// Returns nil if the value is valid, or an error describing the validation failure.
	//
	// Example:
	//
	//	rule := Required[string]()
	//	err := rule.Validate("hello")  // returns nil
	//	err = rule.Validate("")       // returns error
	Validate(value T) error
}
