// Package rule provides validation rules for various data types
package rule

import (
	"errors"
	"fmt"
)

// Common validation errors for condition rules
var (
	ErrCondition     = errors.New("condition validation failed")
	ErrDependency    = errors.New("dependency validation failed")
	ErrMutualExclude = errors.New("mutual exclude validation failed")
)

// ConditionRule represents a rule that combines multiple rules using logical operators.
// It can be used to create complex validation conditions using AND or OR operations.
//
// Example:
//
//	rule := And(Length(5), Contains("a"))
//	err := rule.Validate("abcde")  // returns nil
//	err = rule.Validate("bcde")   // returns ErrCondition
type ConditionRule[T any] struct {
	e        error
	operator string
	rules    []Rule[T]
}

// And creates a new condition rule that requires all sub-rules to pass.
// Returns a ConditionRule that validates using logical AND operation.
//
// Example:
//
//	rule := And(Length(5), Contains("a"))
//	err := rule.Validate("abcde")  // returns nil
func And[T any](rules ...Rule[T]) *ConditionRule[T] {
	return &ConditionRule[T]{
		e:        ErrCondition,
		operator: "AND",
		rules:    rules,
	}
}

// Or creates a new condition rule that requires at least one sub-rule to pass.
// Returns a ConditionRule that validates using logical OR operation.
//
// Example:
//
//	rule := Or(Length(5), Contains("a"))
//	err := rule.Validate("bcde")  // returns nil
func Or[T any](rules ...Rule[T]) *ConditionRule[T] {
	return &ConditionRule[T]{
		e:        ErrCondition,
		operator: "OR",
		rules:    rules,
	}
}

// Validate checks if the value satisfies the condition rule.
// For AND operation, all sub-rules must pass.
// For OR operation, at least one sub-rule must pass.
// Returns nil if validation passes, otherwise returns the error.
func (r *ConditionRule[T]) Validate(value T) error {
	var err error
	switch r.operator {
	case "AND":
		for _, rule := range r.rules {
			if err = rule.Validate(value); err != nil {
				return r.e
			}
		}
		return nil
	case "OR":
		for _, rule := range r.rules {
			if err = rule.Validate(value); err == nil {
				return nil
			}
		}
		return r.e
	default:
		return r.e
	}
}

// Errf sets a custom error message for the validation rule using a formatted string.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := And(Length(5), Contains("a")).Errf("Invalid value")
func (r *ConditionRule[T]) Errf(format string, args ...any) *ConditionRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// DependencyRule represents a rule that validates a field based on another field's value.
// It is used to create conditional validation rules where one field's validation
// depends on another field's value.
//
// Example:
//
//	rule := Dependency("age", "isAdult", GreaterThan(18), func(p Person) int { return p.Age })
//	err := rule.Validate(person)  // returns nil if age > 18
type DependencyRule[T any, D any] struct {
	e           error
	field       string
	dependency  string
	validator   Rule[D]
	valueGetter func(T) D
}

// Dependency creates a new dependency validation rule.
// field is the name of the field being validated.
// dependency is the name of the field it depends on.
// validator is the rule to validate the dependency value.
// valueGetter is a function to extract the dependency value from the main value.
//
// Example:
//
//	rule := Dependency("age", "isAdult", GreaterThan(18), func(p Person) int { return p.Age })
func Dependency[T any, D any](field, dependency string, validator Rule[D], valueGetter func(T) D) *DependencyRule[T, D] {
	return &DependencyRule[T, D]{
		e:           ErrDependency,
		field:       field,
		dependency:  dependency,
		validator:   validator,
		valueGetter: valueGetter,
	}
}

// Validate checks if the value satisfies the dependency rule.
// Returns nil if validation passes, otherwise returns the error.
func (r *DependencyRule[T, D]) Validate(value T) error {
	if r.validator == nil {
		return nil
	}

	dependencyValue := r.valueGetter(value)
	return r.validator.Validate(dependencyValue)
}

// Errf sets a custom error message for the validation rule using a formatted string.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Dependency("age", "isAdult", GreaterThan(18), func(p Person) int { return p.Age }).Errf("Age must be greater than 18")
func (r *DependencyRule[T, D]) Errf(format string, args ...any) *DependencyRule[T, D] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// MutualExcludeRule represents a rule that validates mutual exclusion between fields.
// It ensures that only one of the specified fields can have a value that matches
// the given values.
//
// Example:
//
//	rule := MutualExclude([]string{"type1", "type2"}, []string{"A", "B"}, func(a, b string) bool { return a == b })
//	err := rule.Validate("A")  // returns nil
type MutualExcludeRule[T any] struct {
	e       error
	fields  []string
	values  []T
	compare func(T, T) bool
}

// MutualExclude creates a new mutual exclusion validation rule.
// fields is the list of field names that are mutually exclusive.
// values is the list of values to check for mutual exclusion.
// compare is a function to compare two values for equality.
//
// Example:
//
//	rule := MutualExclude([]string{"type1", "type2"}, []string{"A", "B"}, func(a, b string) bool { return a == b })
func MutualExclude[T any](fields []string, values []T, compare func(T, T) bool) *MutualExcludeRule[T] {
	return &MutualExcludeRule[T]{
		e:       ErrMutualExclude,
		fields:  fields,
		values:  values,
		compare: compare,
	}
}

// Validate checks if the value satisfies the mutual exclusion rule.
// Returns nil if validation passes, otherwise returns the error.
func (r *MutualExcludeRule[T]) Validate(value T) error {
	if r.compare == nil {
		return nil
	}

	valid := false
	for _, v := range r.values {
		if r.compare(v, value) {
			valid = true
			break
		}
	}
	if !valid {
		return r.e
	}

	return nil
}

// Errf sets a custom error message for the validation rule using a formatted string.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := MutualExclude([]string{"type1", "type2"}, []string{"A", "B"}, func(a, b string) bool { return a == b }).Errf("Invalid type")
func (r *MutualExcludeRule[T]) Errf(format string, args ...any) *MutualExcludeRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
