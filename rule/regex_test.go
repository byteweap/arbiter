package rule

import (
	"testing"
)

func TestRegex(t *testing.T) {
	tests := []struct {
		name    string
		rule    *RegexRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: matches pattern",
			rule:    Regex(`^[a-z]+$`),
			value:   "abc",
			wantErr: false,
		},
		{
			name:    "invalid: does not match pattern",
			rule:    Regex(`^[a-z]+$`),
			value:   "123",
			wantErr: true,
		},
		{
			name:    "valid: empty string",
			rule:    Regex(`^[a-z]+$`),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: invalid regex pattern",
			rule:    Regex(`[invalid`),
			value:   "test",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Regex(`^[a-z]+$`).Errf("custom error"),
			value:   "123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegexRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
