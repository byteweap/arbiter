package rule

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIsEmail(t *testing.T) {
	assert.Nil(t, IsEmail().Validate("user@example.com"))
	assert.Error(t, IsEmail().Validate("not-an-email"))
}

func TestIsPhone(t *testing.T) {
	assert.Nil(t, IsPhone().Validate("+8613800138000"))
	assert.Error(t, IsPhone().Validate("abc"))
}

func TestIsSocialCredit(t *testing.T) {
	assert.Nil(t, IsSocialCredit().Validate("913101157984053414"))
	assert.Error(t, IsSocialCredit().Validate("123"))
}

func TestIsTaxNumber(t *testing.T) {
	assert.Nil(t, IsTaxNumber().Validate("123456789012345"))
	assert.Error(t, IsTaxNumber().Validate("abc"))
}

func TestIsBankCard(t *testing.T) {
	assert.Nil(t, IsBankCard().Validate("6222021234567890"))
	assert.Error(t, IsBankCard().Validate("123"))
}

func TestIsPassport(t *testing.T) {
	assert.Nil(t, IsPassport().Validate("E12345678"))
	assert.Error(t, IsPassport().Validate("12345678"))
}

func TestIsIDCard(t *testing.T) {
	assert.Nil(t, IsIDCard().Validate("11010119900307777X"))
	assert.Error(t, IsIDCard().Validate("123"))
}

func TestRegexFallback(t *testing.T) {
	re := regexp.MustCompile(`^[a-z]+$`)
	err := (&RegexRule{regex: re}).Validate("123")
	assert.Error(t, err)
}
