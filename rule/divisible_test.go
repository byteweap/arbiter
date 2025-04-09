package rule

import (
	"testing"
)

func TestDivisibleBy(t *testing.T) {
	tests := []struct {
		name    string
		rule    *DivisibleByRule
		value   float64
		wantErr bool
	}{
		{
			name:    "valid: integer divisible by integer",
			rule:    DivisibleBy(2),
			value:   10,
			wantErr: false,
		},
		{
			name:    "valid: float divisible by float",
			rule:    DivisibleBy(0.5),
			value:   1.5,
			wantErr: false,
		},
		{
			name:    "valid: zero divisible by any non-zero number",
			rule:    DivisibleBy(5),
			value:   0,
			wantErr: false,
		},
		{
			name:    "invalid: integer not divisible by integer",
			rule:    DivisibleBy(3),
			value:   10,
			wantErr: true,
		},
		{
			name:    "invalid: float not divisible by float",
			rule:    DivisibleBy(0.3),
			value:   1.0,
			wantErr: true,
		},
		{
			name:    "invalid: zero divisor",
			rule:    DivisibleBy(0),
			value:   10,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    DivisibleBy(3).Errf("custom error"),
			value:   10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("DivisibleByRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
