package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonZeroRule(t *testing.T) {
	err := NonZero[int]().Validate(1)
	assert.Nil(t, err)

	err = NonZero[int]().Validate(0)
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[int]().Validate(-1)
	assert.Nil(t, err)

	err = NonZero[float64]().Validate(1.0)
	assert.Nil(t, err)

	err = NonZero[float64]().Validate(0.0)
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[float64]().Validate(-1.0)
	assert.Nil(t, err)

	err = NonZero[string]().Validate("")
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[string]().Validate("hello")
	assert.Nil(t, err)

	err = NonZero[bool]().Validate(false)
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[bool]().Validate(true)
	assert.Nil(t, err)

	var p *int
	err = NonZero[*int]().Validate(p)
	assert.Equal(t, ErrNonZero, err)

	i := 1
	err = NonZero[*int]().Validate(&i)
	assert.Nil(t, err)

	err = NonZero[[]int]().Validate(nil)
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[[]int]().Validate([]int{})
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[[]int]().Validate([]int{1, 2, 3})
	assert.Nil(t, err)

	err = NonZero[map[string]int]().Validate(nil)
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[map[string]int]().Validate(map[string]int{})
	assert.Equal(t, ErrNonZero, err)

	err = NonZero[map[string]int]().Validate(map[string]int{"a": 1})
	assert.Nil(t, err)

	customErr := NonZero[int]().Errf("not zero").Validate(0)
	assert.Equal(t, "not zero", customErr.Error())
}

func BenchmarkNonZeroRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NonZero[int]().Validate(1)
	}
}
