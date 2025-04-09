package rule

import (
	"testing"
)

func TestPasswordStrength(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid password",
			value:   "StrongP@ssw0rd",
			wantErr: false,
		},
		{
			name:    "too short",
			value:   "Abc123!",
			wantErr: true,
		},
		{
			name:    "too long",
			value:   "ThisIsAVeryLongPasswordThatExceedsTheMaximumLength123!@#",
			wantErr: true,
		},
		{
			name:    "no uppercase",
			value:   "password123!",
			wantErr: true,
		},
		{
			name:    "no lowercase",
			value:   "PASSWORD123!",
			wantErr: true,
		},
		{
			name:    "no numbers",
			value:   "Password!@#",
			wantErr: true,
		},
		{
			name:    "no special chars",
			value:   "Password123",
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PasswordStrength()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordStrength().Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPasswordComplex(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid complex password",
			value:   "ComplexP@ssw0rd123",
			wantErr: false,
		},
		{
			name:    "too short",
			value:   "Abc123!",
			wantErr: true,
		},
		{
			name:    "insufficient char types",
			value:   "abcdefghijkl",
			wantErr: true,
		},
		{
			name:    "too many repeated chars",
			value:   "aaaabbbbcccc",
			wantErr: true,
		},
		{
			name:    "contains forbidden pattern",
			value:   "mypassword123!",
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := PasswordComplex()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordComplex().Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXSS(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid text",
			value:   "Hello, World!",
			wantErr: false,
		},
		{
			name:    "script tag",
			value:   "<script>alert('xss')</script>",
			wantErr: true,
		},
		{
			name:    "javascript protocol",
			value:   "javascript:alert('xss')",
			wantErr: true,
		},
		{
			name:    "event handler",
			value:   "onclick=alert('xss')",
			wantErr: true,
		},
		{
			name:    "iframe",
			value:   "<iframe src='malicious.html'></iframe>",
			wantErr: true,
		},
		{
			name:    "img tag with event",
			value:   "<img src='x' onerror='alert(1)'>",
			wantErr: true,
		},
		{
			name:    "style tag",
			value:   "<style>body{background:red}</style>",
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := XSS()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("XSS().Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLInjection(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid input",
			value:   "John Doe",
			wantErr: false,
		},
		{
			name:    "select statement",
			value:   "SELECT * FROM users",
			wantErr: true,
		},
		{
			name:    "union attack",
			value:   "1 UNION SELECT * FROM users",
			wantErr: true,
		},
		{
			name:    "comment attack",
			value:   "admin'--",
			wantErr: true,
		},
		{
			name:    "or condition",
			value:   "1 OR 1=1",
			wantErr: true,
		},
		{
			name:    "time-based attack",
			value:   "1; WAITFOR DELAY '0:0:5'",
			wantErr: true,
		},
		{
			name:    "benchmark attack",
			value:   "1; BENCHMARK(10000000,MD5('A'))",
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := SQLInjection()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLInjection().Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomErrorMessages(t *testing.T) {
	tests := []struct {
		name    string
		rule    func() interface{ Validate(string) error }
		value   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "password strength custom error",
			rule: func() interface{ Validate(string) error } {
				return PasswordStrength().Errf("密码强度不足")
			},
			value:   "weak",
			wantErr: true,
			errMsg:  "密码强度不足",
		},
		{
			name: "password complex custom error",
			rule: func() interface{ Validate(string) error } {
				return PasswordComplex().Errf("密码复杂度不足")
			},
			value:   "simple",
			wantErr: true,
			errMsg:  "密码复杂度不足",
		},
		{
			name: "xss custom error",
			rule: func() interface{ Validate(string) error } {
				return XSS().Errf("检测到XSS攻击")
			},
			value:   "<script>alert(1)</script>",
			wantErr: true,
			errMsg:  "检测到XSS攻击",
		},
		{
			name: "sql injection custom error",
			rule: func() interface{ Validate(string) error } {
				return SQLInjection().Errf("检测到SQL注入")
			},
			value:   "SELECT * FROM users",
			wantErr: true,
			errMsg:  "检测到SQL注入",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := tt.rule()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("Validate() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestRuleConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		rule    func() interface{ Validate(string) error }
		value   string
		wantErr bool
	}{
		{
			name: "custom password strength",
			rule: func() interface{ Validate(string) error } {
				return PasswordStrength().
					MinLength(6).
					MaxLength(10).
					RequireUpper(false).
					RequireLower(true).
					RequireNumber(true).
					RequireSpecial(false)
			},
			value:   "pass123",
			wantErr: false,
		},
		{
			name: "custom password complex",
			rule: func() interface{ Validate(string) error } {
				return PasswordComplex().
					MinLength(8).
					MinCharTypes(2).
					MaxRepeatedChars(4).
					AddForbiddenPattern("test")
			},
			value:   "abcd1234",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := tt.rule()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
