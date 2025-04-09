// Package rule provides tests for validation rules.
package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestZeroRule tests the Zero validation rule with various data types.
// It verifies that the rule correctly identifies zero and non-zero values.
func TestZeroRule(t *testing.T) {
	// Test integers
	t.Run("integers", func(t *testing.T) {
		// Zero value should pass
		err := Zero[int]().Validate(0)
		assert.Nil(t, err)

		// Positive value should fail
		err = Zero[int]().Validate(1)
		assert.Equal(t, ErrZero, err)

		// Negative value should fail
		err = Zero[int]().Validate(-1)
		assert.Equal(t, ErrZero, err)
	})

	// Test floating-point numbers
	t.Run("floating-point", func(t *testing.T) {
		// Zero value should pass
		err := Zero[float64]().Validate(0.0)
		assert.Nil(t, err)

		// Positive value should fail
		err = Zero[float64]().Validate(1.0)
		assert.Equal(t, ErrZero, err)

		// Negative value should fail
		err = Zero[float64]().Validate(-1.0)
		assert.Equal(t, ErrZero, err)
	})

	// Test strings
	t.Run("strings", func(t *testing.T) {
		// Empty string should pass
		err := Zero[string]().Validate("")
		assert.Nil(t, err)

		// Non-empty string should fail
		err = Zero[string]().Validate("hello")
		assert.Equal(t, ErrZero, err)
	})

	// Test booleans
	t.Run("booleans", func(t *testing.T) {
		// False should pass
		err := Zero[bool]().Validate(false)
		assert.Nil(t, err)

		// True should fail
		err = Zero[bool]().Validate(true)
		assert.Equal(t, ErrZero, err)
	})

	// Test pointers
	t.Run("pointers", func(t *testing.T) {
		// Nil pointer should pass
		var p *int
		err := Zero[*int]().Validate(p)
		assert.Nil(t, err)

		// Non-nil pointer should fail
		i := 1
		err = Zero[*int]().Validate(&i)
		assert.Equal(t, ErrZero, err)
	})

	// Test slices
	t.Run("slices", func(t *testing.T) {
		// Nil slice should pass
		err := Zero[[]int]().Validate(nil)
		assert.Nil(t, err)

		// Empty slice should pass
		err = Zero[[]int]().Validate([]int{})
		assert.Nil(t, err)

		// Non-empty slice should fail
		err = Zero[[]int]().Validate([]int{1, 2, 3})
		assert.Equal(t, ErrZero, err)
	})

	// Test maps
	t.Run("maps", func(t *testing.T) {
		// Nil map should pass
		err := Zero[map[string]int]().Validate(nil)
		assert.Nil(t, err)

		// Empty map should pass
		err = Zero[map[string]int]().Validate(map[string]int{})
		assert.Nil(t, err)

		// Non-empty map should fail
		err = Zero[map[string]int]().Validate(map[string]int{"a": 1})
		assert.Equal(t, ErrZero, err)
	})

	// Test custom error messages
	t.Run("custom error messages", func(t *testing.T) {
		// Custom error message should be returned
		err := Zero[int]().Errf("Value must be zero").Validate(1)
		assert.Equal(t, "Value must be zero", err.Error())
	})
}

// BenchmarkZeroRule benchmarks the performance of the Zero validation rule.
// It measures the time taken to validate a zero integer value.
func BenchmarkZeroRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Zero[int]().Validate(0)
	}
}
