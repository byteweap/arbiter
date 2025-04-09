package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegativeRule(t *testing.T) {
	err := Negative[int]().Validate(1)
	assert.Equal(t, ErrNegative, err)

	err = Negative[int]().Validate(0)
	assert.Equal(t, ErrNegative, err)

	err = Negative[int]().Validate(-1)
	assert.Nil(t, err)

	err = Negative[float64]().Validate(1.0)
	assert.Equal(t, ErrNegative, err)

	err = Negative[float64]().Validate(0.0)
	assert.Equal(t, ErrNegative, err)

	err = Negative[float64]().Validate(-1.0)
	assert.Nil(t, err)

	customErr := Negative[int]().Errf("must be negative").Validate(1)
	assert.Equal(t, "must be negative", customErr.Error())
}

func BenchmarkNegativeRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Negative[int]().Validate(-1)
	}
}
