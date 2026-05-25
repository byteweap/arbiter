package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilRule(t *testing.T) {
	var err error

	err = NotNil.Validate(1)
	assert.Equal(t, nil, err)

	err = NotNil.Validate("11111")
	assert.Equal(t, nil, err)

	var v *string
	err = Nil.Validate(v)
	assert.Equal(t, nil, err)

}

func TestNilRuleErrf(t *testing.T) {
	err := Nil.Errf("custom nil error").Validate("not nil")
	if err == nil || err.Error() != "custom nil error" {
		t.Errorf("Nil.Errf() error = %v, want custom nil error", err)
	}

	var v *string
	err = NotNil.Errf("custom notnil error").Validate(v)
	if err == nil || err.Error() != "custom notnil error" {
		t.Errorf("NotNil.Errf() error = %v, want custom notnil error", err)
	}
}

func BenchmarkNilRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Nil.Validate(nil)
	}
}
