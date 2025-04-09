package rule

import (
	"errors"
	"testing"
)

func TestStringStartsWith(t *testing.T) {
	tests := []struct {
		name    string
		rule    *StartWithRule
		value   string
		wantErr bool
	}{
		{
			name:    "valid: string starts with prefix",
			rule:    StartWith("hello"),
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "valid: string equals prefix",
			rule:    StartWith("hello"),
			value:   "hello",
			wantErr: false,
		},
		{
			name:    "invalid: string does not start with prefix",
			rule:    StartWith("hello"),
			value:   "world hello",
			wantErr: true,
		},
		{
			name:    "invalid: empty string",
			rule:    StartWith("hello"),
			value:   "",
			wantErr: true,
		},
		{
			name:    "invalid: empty prefix",
			rule:    StartWith(""),
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "custom error message",
			rule:    StartWith("hello").Errf("custom error"),
			value:   "world hello",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartsWithRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		suffix   string
		expected error
	}{
		{
			name:     "valid suffix",
			value:    "hello world",
			suffix:   "world",
			expected: nil,
		},
		{
			name:     "exact match",
			value:    "world",
			suffix:   "world",
			expected: nil,
		},
		{
			name:     "invalid suffix",
			value:    "hello world",
			suffix:   "hello",
			expected: ErrEndsWith,
		},
		{
			name:     "empty string",
			value:    "",
			suffix:   "world",
			expected: ErrEndsWith,
		},
		{
			name:     "empty suffix",
			value:    "hello world",
			suffix:   "",
			expected: nil,
		},
		{
			name:     "custom error message",
			value:    "hello world",
			suffix:   "hello",
			expected: errors.New("custom error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := EndWith(tt.suffix)
			if tt.expected != nil && tt.expected.Error() == "custom error" {
				rule.Errf("custom error")
			}

			err := rule.Validate(tt.value)
			if err != tt.expected {
				if err == nil || tt.expected == nil || err.Error() != tt.expected.Error() {
					t.Errorf("EndsWith(%q).Validate(%q) = %v, want %v", tt.suffix, tt.value, err, tt.expected)
				}
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		value   string
		wantErr bool
	}{
		{
			name:    "valid prefix",
			prefix:  "hello",
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			prefix:  "world",
			value:   "hello world",
			wantErr: true,
		},
		{
			name:    "empty string",
			prefix:  "hello",
			value:   "",
			wantErr: true,
		},
		{
			name:    "empty prefix",
			prefix:  "",
			value:   "hello world",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := StartWith(tt.prefix)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartWith() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name    string
		suffix  string
		value   string
		wantErr bool
	}{
		{
			name:    "valid suffix",
			suffix:  "world",
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "invalid suffix",
			suffix:  "hello",
			value:   "hello world",
			wantErr: true,
		},
		{
			name:    "empty string",
			suffix:  "world",
			value:   "",
			wantErr: true,
		},
		{
			name:    "empty suffix",
			suffix:  "",
			value:   "hello world",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := EndWith(tt.suffix)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("EndsWith() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChineseOnly(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid chinese",
			value:   "你好世界",
			wantErr: false,
		},
		{
			name:    "mixed with english",
			value:   "你好world",
			wantErr: true,
		},
		{
			name:    "mixed with numbers",
			value:   "你好123",
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
			rule := OnlyChinese()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChineseOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFullWidthOnly(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid full width",
			value:   "ＨＥＬＬＯ",
			wantErr: false,
		},
		{
			name:    "mixed with half width",
			value:   "ＨＥＬＬＯ world",
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
			rule := OnlyFullWidth()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("FullWidthOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHalfWidthOnly(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid half width",
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "mixed with full width",
			value:   "hello ＷＯＲＬＤ",
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
			rule := OnlyHalfWidth()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HalfWidthOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpperCaseOnly(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid uppercase",
			value:   "HELLO WORLD",
			wantErr: false,
		},
		{
			name:    "mixed with lowercase",
			value:   "HELLO world",
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
			rule := OnlyUpperCase()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpperCaseOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLowerCaseOnly(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "valid lowercase",
			value:   "hello world",
			wantErr: false,
		},
		{
			name:    "mixed with uppercase",
			value:   "hello WORLD",
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
			rule := OnlyLowerCase()
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("LowerCaseOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecialChars(t *testing.T) {
	tests := []struct {
		name         string
		allowSpecial bool
		value        string
		wantErr      bool
	}{
		{
			name:         "allow special chars - valid",
			allowSpecial: true,
			value:        "hello@world",
			wantErr:      false,
		},
		{
			name:         "allow special chars - invalid",
			allowSpecial: true,
			value:        "hello world",
			wantErr:      true,
		},
		{
			name:         "disallow special chars - valid",
			allowSpecial: false,
			value:        "hello world",
			wantErr:      false,
		},
		{
			name:         "disallow special chars - invalid",
			allowSpecial: false,
			value:        "hello@world",
			wantErr:      true,
		},
		{
			name:         "empty string",
			allowSpecial: true,
			value:        "",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := SpecialChars(tt.allowSpecial)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpecialChars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name      string
		substring string
		value     string
		wantErr   bool
	}{
		{
			name:      "valid contains",
			substring: "world",
			value:     "hello world",
			wantErr:   false,
		},
		{
			name:      "invalid contains",
			substring: "test",
			value:     "hello world",
			wantErr:   true,
		},
		{
			name:      "empty string",
			substring: "world",
			value:     "",
			wantErr:   true,
		},
		{
			name:      "empty substring",
			substring: "",
			value:     "hello world",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := Contains(tt.substring)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotContains(t *testing.T) {
	tests := []struct {
		name      string
		substring string
		value     string
		wantErr   bool
	}{
		{
			name:      "valid not contains",
			substring: "test",
			value:     "hello world",
			wantErr:   false,
		},
		{
			name:      "invalid not contains",
			substring: "world",
			value:     "hello world",
			wantErr:   true,
		},
		{
			name:      "empty string",
			substring: "world",
			value:     "",
			wantErr:   false,
		},
		{
			name:      "empty substring",
			substring: "",
			value:     "hello world",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := NotContains(tt.substring)
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotContains() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
