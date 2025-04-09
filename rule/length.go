// Package rule provides a collection of validation rules for various data types.
package rule

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// ErrLength is returned when a value's length is outside the specified range.
const ErrLengthFormat = "length is not between %v and %v"

// LengthRule validates that a value's length falls within a specified range.
// It supports strings (counting Unicode characters), slices, arrays, and maps.
//
// Example:
//
//	rule := Len[string](5, 10)
//	err := rule.Validate("hello")     // returns nil
//	err = rule.Validate("hi")         // returns ErrLength
//
//	rule = Len[[]int](2, 4)
//	err = rule.Validate([]int{1, 2})  // returns nil
//	err = rule.Validate([]int{1})     // returns ErrLength
type LengthRule[T any] struct {
	min int
	max int
	e   error
}

// Len creates a new length validation rule with the specified minimum and maximum lengths.
// The rule can be used with any type that has a measurable length (strings, slices, arrays, maps).
//
// Example:
//
//	// Create a rule for usernames (5-20 characters)
//	usernameRule := Len[string](5, 20).Err("Username must be 5-20 characters long")
//
//	// Create a rule for arrays (1-5 elements)
//	arrayRule := Len[[]int](1, 5).Err("Array must have 1-5 elements")
func Len[T any](min int, max int) *LengthRule[T] {
	return &LengthRule[T]{min: min, max: max, e: fmt.Errorf(ErrLengthFormat, min, max)}
}

// Validate checks if the value's length falls within the specified range.
// For strings, it counts Unicode characters (runes).
// For slices and arrays, it counts elements.
// For maps, it counts key-value pairs.
// Returns nil if the length is valid, or an error if it's outside the range.
//
// Example:
//
//	rule := Len[string](5, 10)
//	if err := rule.Validate("hello"); err != nil {
//	    // Handle validation error
//	}
//
//	rule = Len[[]int](2, 4)
//	if err := rule.Validate([]int{1, 2, 3}); err != nil {
//	    // Handle validation error
//	}
func (r *LengthRule[T]) Validate(value T) error {
	var (
		length int
		err    error
	)
	switch v := any(value).(type) {
	case string:
		length = utf8.RuneCountInString(v) // Count Unicode characters
	case *string:
		length = Ternary(v == nil, 0, utf8.RuneCountInString(*v))
	case []int:
		length = len(v)
	case []*int:
		length = Ternary(v == nil, 0, len(v))
	case []int8:
		length = len(v)
	case []*int8:
		length = Ternary(v == nil, 0, len(v))
	case []int16:
		length = len(v)
	case []*int16:
		length = Ternary(v == nil, 0, len(v))
	case []int32:
		length = len(v)
	case []*int32:
		length = Ternary(v == nil, 0, len(v))
	case []int64:
		length = len(v)
	case []*int64:
		length = Ternary(v == nil, 0, len(v))
	case []uint:
		length = len(v)
	case []*uint:
		length = Ternary(v == nil, 0, len(v))
	case []uint8:
		length = len(v)
	case []*uint8:
		length = Ternary(v == nil, 0, len(v))
	case []uint16:
		length = len(v)
	case []*uint16:
		length = Ternary(v == nil, 0, len(v))
	case []uint32:
		length = len(v)
	case []*uint32:
		length = Ternary(v == nil, 0, len(v))
	case []uint64:
		length = len(v)
	case []*uint64:
		length = Ternary(v == nil, 0, len(v))
	case []float32:
		length = len(v)
	case []*float32:
		length = Ternary(v == nil, 0, len(v))
	case []float64:
		length = len(v)
	case []*float64:
		length = Ternary(v == nil, 0, len(v))
	case []bool:
		length = len(v)
	case []*bool:
		length = Ternary(v == nil, 0, len(v))
	case []any:
		length = len(v)
	case []*any:
		length = Ternary(v == nil, 0, len(v))
	case []struct{}:
		length = len(v)
	case []*struct{}:
		length = Ternary(v == nil, 0, len(v))
	default:
		// Use reflection for other types
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String, reflect.Map, reflect.Slice, reflect.Array:
			length = val.Len()
		default:
			err = fmt.Errorf("cannot get length of %v", val.Kind())
		}
	}
	if err != nil {
		return err
	}
	if length < r.min || length > r.max {
		return r.e
	}
	return nil
}

// Errf sets a custom error message for the validation rule.
// Returns the rule instance for method chaining.
//
// Example:
//
//	rule := Len[string](5, 10).Errf("Text must be between 5 and 10 characters")
//	err := rule.Validate("hi")  // returns error with custom message
func (r *LengthRule[T]) Errf(format string, args ...any) *LengthRule[T] {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
