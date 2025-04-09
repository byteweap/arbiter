// Package rule provides tests for validation rules.
package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMinRule tests the Min validation rule with various data types.
// It verifies that values are correctly validated against minimum thresholds.
func TestMinRule(t *testing.T) {
	// Test integer validation
	t.Run("integer validation", func(t *testing.T) {
		// Value below minimum should fail
		err := Min(10).Validate(5)
		assert.Equal(t, ErrMin, err)

		// Value above minimum should pass
		err = Min(10).Validate(15)
		assert.Nil(t, err)

		// Value equal to minimum should pass
		err = Min(10).Validate(10)
		assert.Nil(t, err)
	})

	// Test floating-point validation
	t.Run("floating-point validation", func(t *testing.T) {
		// Value below minimum should fail
		err := Min(10.5).Validate(5.5)
		assert.Equal(t, ErrMin, err)

		// Value above minimum should pass
		err = Min(10.5).Validate(15.5)
		assert.Nil(t, err)

		// Value equal to minimum should pass
		err = Min(10.5).Validate(10.5)
		assert.Nil(t, err)
	})

	// Test custom error messages
	t.Run("custom error messages", func(t *testing.T) {
		err := Min(10).Errf("Value must be at least 10").Validate(5)
		assert.Equal(t, "Value must be at least 10", err.Error())
	})
}

// TestMaxRule tests the Max validation rule with various data types.
// It verifies that values are correctly validated against maximum thresholds.
func TestMaxRule(t *testing.T) {
	// Test integer validation
	t.Run("integer validation", func(t *testing.T) {
		// Value above maximum should fail
		err := Max(10).Validate(15)
		assert.Equal(t, ErrMax, err)

		// Value below maximum should pass
		err = Max(10).Validate(5)
		assert.Nil(t, err)

		// Value equal to maximum should pass
		err = Max(10).Validate(10)
		assert.Nil(t, err)
	})

	// Test floating-point validation
	t.Run("floating-point validation", func(t *testing.T) {
		// Value above maximum should fail
		err := Max(10.5).Validate(15.5)
		assert.Equal(t, ErrMax, err)

		// Value below maximum should pass
		err = Max(10.5).Validate(5.5)
		assert.Nil(t, err)

		// Value equal to maximum should pass
		err = Max(10.5).Validate(10.5)
		assert.Nil(t, err)
	})

	// Test custom error messages
	t.Run("custom error messages", func(t *testing.T) {
		err := Max(10).Errf("Value must not exceed 10").Validate(15)
		assert.Equal(t, "Value must not exceed 10", err.Error())
	})
}

// TestMaxWithTableDriven tests the Max validation rule using table-driven tests.
// It provides comprehensive test cases for integer validation.
func TestMaxWithTableDriven(t *testing.T) {
	tests := []struct {
		name    string
		rule    *MaxRule[int]
		value   int
		wantErr bool
	}{
		{
			name:    "value below maximum",
			rule:    Max(10),
			value:   5,
			wantErr: false,
		},
		{
			name:    "value equals maximum",
			rule:    Max(10),
			value:   10,
			wantErr: false,
		},
		{
			name:    "value above maximum",
			rule:    Max(10),
			value:   15,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Max(10).Errf("Value must not exceed 10"),
			value:   15,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestMaxFloatWithTableDriven tests the Max validation rule using table-driven tests.
// It provides comprehensive test cases for floating-point validation.
func TestMaxFloatWithTableDriven(t *testing.T) {
	tests := []struct {
		name    string
		rule    *MaxRule[float64]
		value   float64
		wantErr bool
	}{
		{
			name:    "value below maximum",
			rule:    Max(10.5),
			value:   5.5,
			wantErr: false,
		},
		{
			name:    "value equals maximum",
			rule:    Max(10.5),
			value:   10.5,
			wantErr: false,
		},
		{
			name:    "value above maximum",
			rule:    Max(10.5),
			value:   15.5,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Max(10.5).Errf("Value must not exceed 10.5"),
			value:   15.5,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// BenchmarkMinRule benchmarks the performance of the Min validation rule.
// It measures the time taken to validate an integer value against a minimum threshold.
func BenchmarkMinRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Min(10).Validate(5)
	}
}

// BenchmarkMaxRule benchmarks the performance of the Max validation rule.
// It measures the time taken to validate an integer value against a maximum threshold.
func BenchmarkMaxRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Max(10).Validate(15)
	}
}
