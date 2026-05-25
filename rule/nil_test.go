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
	assert.Error(t, err)
	assert.Equal(t, "custom nil error", err.Error())

	var v *string
	err = NotNil.Errf("custom notnil error").Validate(v)
	assert.Error(t, err)
	assert.Equal(t, "custom notnil error", err.Error())
}

func BenchmarkNilRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Nil.Validate(nil)
	}
}
