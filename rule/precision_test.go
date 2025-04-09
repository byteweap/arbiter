package rule

import (
	"testing"
)

func TestPrecision(t *testing.T) {
	tests := []struct {
		name      string
		precision int
		value     float64
		wantErr   bool
	}{
		{
			name:      "valid precision",
			precision: 2,
			value:     3.14,
			wantErr:   false,
		},
		{
			name:      "exceed precision",
			precision: 2,
			value:     3.14159,
			wantErr:   true,
		},
		{
			name:      "integer value",
			precision: 2,
			value:     42,
			wantErr:   false,
		},
		{
			name:      "zero precision",
			precision: 0,
			value:     3.0,
			wantErr:   false,
		},
		{
			name:      "zero precision with decimal",
			precision: 0,
			value:     3.1,
			wantErr:   true,
		},
		{
			name:      "scientific notation",
			precision: 2,
			value:     1.23e-4,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Precision(tt.precision)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Precision() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFloat32Precision(t *testing.T) {
	tests := []struct {
		name      string
		precision int
		value     float32
		wantErr   bool
	}{
		{
			name:      "valid precision",
			precision: 2,
			value:     3.14,
			wantErr:   false,
		},
		{
			name:      "exceed precision",
			precision: 2,
			value:     3.14159,
			wantErr:   true,
		},
		{
			name:      "integer value",
			precision: 2,
			value:     42,
			wantErr:   false,
		},
		{
			name:      "zero precision",
			precision: 0,
			value:     3.0,
			wantErr:   false,
		},
		{
			name:      "zero precision with decimal",
			precision: 0,
			value:     3.1,
			wantErr:   true,
		},
		{
			name:      "scientific notation",
			precision: 2,
			value:     1.23e-4,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Float32Precision(tt.precision)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float32Precision() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrecisionCustomError(t *testing.T) {
	tests := []struct {
		name      string
		precision int
		value     float64
		wantErr   string
	}{
		{
			name:      "custom error message",
			precision: 2,
			value:     3.14159,
			wantErr:   ErrPrecision.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Precision(tt.precision).Errf(tt.wantErr)
			err := rule.Validate(tt.value)
			if err == nil || err.Error() != tt.wantErr {
				t.Errorf("Precision() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestFloat32PrecisionCustomError(t *testing.T) {
	tests := []struct {
		name      string
		precision int
		value     float32
		wantErr   string
	}{
		{
			name:      "custom error message",
			precision: 2,
			value:     3.14159,
			wantErr:   "小数位数不能超过2位",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Float32Precision(tt.precision).Errf(tt.wantErr)
			err := rule.Validate(tt.value)
			if err == nil || err.Error() != tt.wantErr {
				t.Errorf("Float32Precision() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
