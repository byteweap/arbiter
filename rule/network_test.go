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

func TestDomainErrf(t *testing.T) {
	rule := Domain().Errf("custom domain error")
	err := rule.Validate("invalid-domain")
	if err == nil || err.Error() != "custom domain error" {
		t.Errorf("Domain().Errf() error = %v, want custom domain error", err)
	}
}

func TestPortRule(t *testing.T) {
	tests := []struct {
		name    string
		port    string
		wantErr bool
	}{
		{name: "valid port", port: "8080", wantErr: false},
		{name: "port 0", port: "0", wantErr: false},
		{name: "port 65535", port: "65535", wantErr: false},
		{name: "negative", port: "-1", wantErr: true},
		{name: "too large", port: "65536", wantErr: true},
		{name: "non-numeric", port: "abc", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Port()
			err := rule.Validate(tt.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("Port().Validate(%q) error = %v, wantErr %v", tt.port, err, tt.wantErr)
			}
		})
	}
}

func TestPortRuleErrf(t *testing.T) {
	rule := Port().Errf("custom port error")
	err := rule.Validate("99999")
	if err == nil || err.Error() != "custom port error" {
		t.Errorf("Port().Errf() error = %v, want custom port error", err)
	}
}

func TestMACAddressRule(t *testing.T) {
	tests := []struct {
		name    string
		mac     string
		wantErr bool
	}{
		{name: "valid colon format", mac: "00:11:22:33:44:55", wantErr: false},
		{name: "valid hyphen format", mac: "00-11-22-33-44-55", wantErr: false},
		{name: "invalid", mac: "invalid-mac", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := MACAddress()
			err := rule.Validate(tt.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("MACAddress().Validate(%q) error = %v, wantErr %v", tt.mac, err, tt.wantErr)
			}
		})
	}
}

func TestMACAddressRuleErrf(t *testing.T) {
	rule := MACAddress().Errf("custom mac error")
	err := rule.Validate("invalid-mac")
	if err == nil || err.Error() != "custom mac error" {
		t.Errorf("MACAddress().Errf() error = %v, want custom mac error", err)
	}
}

func TestSubnetMaskRule(t *testing.T) {
	tests := []struct {
		name    string
		mask    string
		wantErr bool
	}{
		{name: "valid /0", mask: "0.0.0.0", wantErr: false},
		{name: "valid /8", mask: "255.0.0.0", wantErr: false},
		{name: "valid /16", mask: "255.255.0.0", wantErr: false},
		{name: "valid /24", mask: "255.255.255.0", wantErr: false},
		{name: "valid /20", mask: "255.255.240.0", wantErr: false},
		{name: "valid /10", mask: "255.192.0.0", wantErr: false},
		{name: "non-contiguous", mask: "255.255.255.1", wantErr: true},
		{name: "non-contiguous 2", mask: "255.255.0.255", wantErr: true},
		{name: "not an ip", mask: "not-an-ip", wantErr: true},
		{name: "ipv6", mask: "::1", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := SubnetMask()
			err := rule.Validate(tt.mask)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubnetMask().Validate(%q) error = %v, wantErr %v", tt.mask, err, tt.wantErr)
			}
		})
	}
}

func TestSubnetMaskRuleErrf(t *testing.T) {
	rule := SubnetMask().Errf("custom mask error")
	err := rule.Validate("255.255.255.1")
	if err == nil || err.Error() != "custom mask error" {
		t.Errorf("SubnetMask().Errf() error = %v, want custom mask error", err)
	}
}
