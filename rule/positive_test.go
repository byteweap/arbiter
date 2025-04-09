package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveRule(t *testing.T) {
	err := Positive[int]().Validate(-1)
	assert.Equal(t, ErrPositive, err)

	err = Positive[int]().Validate(0)
	assert.Equal(t, ErrPositive, err)

	err = Positive[int]().Validate(1)
	assert.Nil(t, err)

	err = Positive[float64]().Validate(-1.0)
	assert.Equal(t, ErrPositive, err)

	err = Positive[float64]().Validate(0.0)
	assert.Equal(t, ErrPositive, err)

	err = Positive[float64]().Validate(1.0)
	assert.Nil(t, err)

	customErr := Positive[int]().Errf("must be positive number").Validate(-1)
	assert.Equal(t, "must be positive number", customErr.Error())
}

func BenchmarkPositiveRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Positive[int]().Validate(1)
	}
}
