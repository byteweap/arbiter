package rule

import (
	"testing"
)

func TestDomain(t *testing.T) {
	tests := []struct {
		name    string
		domain  string
		wantErr bool
	}{
		{
			name:    "valid domain",
			domain:  "example.com",
			wantErr: false,
		},
		{
			name:    "invalid domain",
			domain:  "invalid-domain",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Domain()
			err := rule.Validate(tt.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Domain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
