// Package arbiter provides validation functionality for various data types.
// It includes functions for validating individual values and structs with multiple fields.
package arbiter

import (
	"errors"
	"reflect"

	"github.com/byteweap/arbiter/rule"
)

// Validate applies a series of validation rules to a value.
// It returns nil if all rules pass, or the first error encountered.
//
// Example:
//
//	// Validate a string with multiple rules
//	err := Validate("hello",
//	    rule.Len[string](3, 10), // length between 3 and 10
//	    rule.Required[string](), // must not be empty
//	)
//
//	// Validate a number with multiple rules
//	err := Validate(42,
//	    rule.Min(0),                  // greater than or equal to 0
//	    rule.Max(100),                // less than or equal to 100
//	    rule.Even(),                  // must be even
//	)
//
//	// Validate a slice with multiple rules
//	err := Validate([]int{1, 2, 3},
//	    rule.Len[[]int](1, 5), // length between 1 and 5
//	    rule.NonZero[[]int](), // must not be empty
//	)
func Validate[T any](value T, rules ...rule.Rule[T]) error {
	for _, r := range rules {
		if err := r.Validate(value); err != nil {
			return err
		}
	}
	return nil
}

// ValidateAll applies a series of validation rules to a value and returns a list of errors.
// It returns nil if all rules pass.
//
// Example:
//
//	errs := ValidateAll("hello",
//	    rule.Len[string](3, 10), // length between 3 and 10
//	    rule.Required[string](), // must not be empty
//	)
func ValidateAll[T any](value T, rules ...rule.Rule[T]) []error {
	var errs []error
	for _, r := range rules {
		if err := r.Validate(value); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// ValidateStruct validates a struct by applying rules to its fields.
// The value parameter must be a pointer to a struct.
// The nilErr parameter is the error message to use if the struct is nil.
// The fields parameter is a list of field rules to apply.
//
// Example:
//
//	type Person struct {
//	    Name  string
//	    Age   int
//	    Email string
//	}
//
//	person := &Person{
//	    Name:  "John",
//	    Age:   30,
//	    Email: "john@example.com",
//	}
//
//	err := ValidateStruct(person, "Person cannot be nil",
//	    Field(&person.Name,
//	        rule.Len[string](2, 50),
//	        rule.Required[string]().Errf("Name is required"),
//	    ),
//	    Field(&person.Age,
//	        rule.Min(0),
//	        rule.Max(120),
//	    ),
//	    Field(&person.Email,
//	        rule.Required[string](),
//	        rule.IsEmail().Errf("Invalid email"),
//	    ),
//	)
func ValidateStruct(value any, nilErr string, fields ...IFieldRule) error {
	if value == nil {
		if nilErr != "" {
			return errors.New(nilErr)
		}
		return errors.New("value cannot be nil")
	}
	// value is must be a pointer
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || (!v.IsNil() && v.Elem().Kind() != reflect.Struct) {
		return errors.New("value must be a pointer")
	}
	// value is must not nil
	if err := Validate(value, rule.NotNil.Errf(nilErr)); err != nil {
		return err
	}
	// validate fields
	for _, field := range fields {
		if err := field.validate(); err != nil {
			return err
		}
	}
	return nil
}
