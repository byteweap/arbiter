package rule

import (
	"testing"
)

func TestEven(t *testing.T) {
	tests := []struct {
		name    string
		rule    *EvenRule[int]
		value   int
		wantErr bool
	}{
		{
			name:    "valid: even number",
			rule:    Even[int](),
			value:   2,
			wantErr: false,
		},
		{
			name:    "valid: even number",
			rule:    Even[int](),
			value:   -2,
			wantErr: false,
		},
		{
			name:    "invalid: odd number",
			rule:    Even[int](),
			value:   1,
			wantErr: true,
		},
		{
			name:    "invalid: odd number",
			rule:    Even[int](),
			value:   -1,
			wantErr: true,
		},
		{
			name:    "valid: zero",
			rule:    Even[int](),
			value:   0,
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    Even[int]().Errf("custom error"),
			value:   1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvenRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
