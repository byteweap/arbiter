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

func BenchmarkNilRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Nil.Validate(nil)
	}
}
