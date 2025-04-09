package rule

import (
	"testing"
)

func TestIP(t *testing.T) {
	tests := []struct {
		name    string
		rule    *IPRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: ipv4",
			rule:    IP(),
			value:   "192.168.1.1",
			wantErr: false,
		},
		{
			name:    "valid: ipv6",
			rule:    IP(),
			value:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: false,
		},
		{
			name:    "valid: empty string",
			rule:    IP(),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: invalid ip",
			rule:    IP(),
			value:   "256.256.256.256",
			wantErr: true,
		},
		{
			name:    "invalid: not an ip",
			rule:    IP(),
			value:   "not an ip",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    IP().Errf("custom error"),
			value:   "not an ip",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPv4(t *testing.T) {
	tests := []struct {
		name    string
		rule    *IPv4Rule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: ipv4",
			rule:    IPv4(),
			value:   "192.168.1.1",
			wantErr: false,
		},
		{
			name:    "valid: localhost",
			rule:    IPv4(),
			value:   "127.0.0.1",
			wantErr: false,
		},
		{
			name:    "valid: empty string",
			rule:    IPv4(),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: ipv6",
			rule:    IPv4(),
			value:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: true,
		},
		{
			name:    "invalid: invalid ipv4",
			rule:    IPv4(),
			value:   "256.256.256.256",
			wantErr: true,
		},
		{
			name:    "invalid: not an ip",
			rule:    IPv4(),
			value:   "not an ip",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    IPv4().Errf("custom error"),
			value:   "not an ip",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPv4Rule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPv6(t *testing.T) {
	tests := []struct {
		name    string
		rule    *IPv6Rule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: ipv6",
			rule:    IPv6(),
			value:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: false,
		},
		{
			name:    "valid: compressed ipv6",
			rule:    IPv6(),
			value:   "2001:db8::2:1",
			wantErr: false,
		},
		{
			name:    "valid: empty string",
			rule:    IPv6(),
			value:   "",
			wantErr: false,
		},
		{
			name:    "invalid: ipv4",
			rule:    IPv6(),
			value:   "192.168.1.1",
			wantErr: true,
		},
		{
			name:    "invalid: invalid ipv6",
			rule:    IPv6(),
			value:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334:extra",
			wantErr: true,
		},
		{
			name:    "invalid: not an ip",
			rule:    IPv6(),
			value:   "not an ip",
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    IPv6().Errf("custom error"),
			value:   "not an ip",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPv6Rule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
