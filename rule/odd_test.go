package rule

import (
	"testing"
)

func TestOdd(t *testing.T) {
	tests := []struct {
		name    string
		rule    *OddRule[int]
		value   int
		wantErr bool
	}{
		{
			name:    "valid: odd number",
			rule:    Odd[int](),
			value:   1,
			wantErr: false,
		},
		{
			name:    "valid: odd number",
			rule:    Odd[int](),
			value:   -1,
			wantErr: false,
		},
		{
			name:    "invalid: even number",
			rule:    Odd[int](),
			value:   2,
			wantErr: true,
		},
		{
			name:    "invalid: even number",
			rule:    Odd[int](),
			value:   -2,
			wantErr: true,
		},
		{
			name:    "invalid: zero",
			rule:    Odd[int](),
			value:   0,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Odd[int]().Errf("custom error"),
			value:   2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("OddRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
