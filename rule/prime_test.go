package rule

import "testing"

func TestPrime(t *testing.T) {
	tests := []struct {
		name    string
		rule    *PrimeRule
		value   int
		wantErr bool
	}{
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   2,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   3,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   5,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   7,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   11,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   13,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   17,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   19,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   23,
			wantErr: false,
		},
		{
			name:    "valid prime numbers",
			rule:    Prime(),
			value:   29,
			wantErr: false,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   4,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   6,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   8,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   9,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   10,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   12,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   14,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   15,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   16,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   18,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   20,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   21,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   22,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   24,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   25,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   26,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   27,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   28,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   30,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   1,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   0,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -1,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -2,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -3,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -4,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -5,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -6,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -7,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -8,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -9,
			wantErr: true,
		},
		{
			name:    "non-prime numbers",
			rule:    Prime(),
			value:   -10,
			wantErr: true,
		},
		{
			name:    "custom error message",
			rule:    Prime().Errf("custom error"),
			value:   4,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrimeRule.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
