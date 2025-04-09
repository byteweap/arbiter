package rule

import (
	"testing"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name    string
		rule    *UUIDRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: standard uuid",
			rule:    UUID(),
			value:   "123e4567-e89b-12d3-a456-426614174000",
			wantErr: false,
		},
		{
			name:    "valid: uppercase uuid",
			rule:    UUID(),
			value:   "123E4567-E89B-12D3-A456-426614174000",
			wantErr: false,
		},
		{
			name:    "valid: empty string",
			rule:    UUID(),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: wrong format",
			rule:    UUID(),
			value:   "123e4567-e89b-12d3-a456",
			wantErr: true,
		},
		{
			name:    "invalid: invalid characters",
			rule:    UUID(),
			value:   "123e4567-e89b-12d3-a456-42661417400g",
			wantErr: true,
		},
		{
			name:    "invalid: not a uuid",
			rule:    UUID(),
			value:   "not a uuid",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    UUID().Errf("custom error"),
			value:   "not a uuid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UUIDRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
