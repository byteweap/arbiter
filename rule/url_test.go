package rule

import (
	"testing"
)

func TestURL(t *testing.T) {
	tests := []struct {
		name    string
		rule    *URLRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: http url",
			rule:    URL(),
			value:   "http://example.com",
			wantErr: false,
		},
		{
			name:    "valid: https url",
			rule:    URL(),
			value:   "https://example.com",
			wantErr: false,
		},
		{
			name:    "valid: url with path",
			rule:    URL(),
			value:   "https://example.com/path",
			wantErr: false,
		},
		{
			name:    "valid: url with query",
			rule:    URL(),
			value:   "https://example.com?query=value",
			wantErr: false,
		},
		{
			name:    "valid: empty string",
			rule:    URL(),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: no scheme",
			rule:    URL(),
			value:   "example.com",
			wantErr: true,
		},
		{
			name:    "invalid: invalid url",
			rule:    URL(),
			value:   "not a url",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    URL().Errf("custom error"),
			value:   "not a url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
