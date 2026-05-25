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

	assert.Equal(t, ErrNonZero, NonZero[[0]int]().Validate([0]int{}))
	assert.Nil(t, NonZero[[3]int]().Validate([3]int{1, 2, 3}))

	assert.Equal(t, "bool err", NonZero[bool]().Errf("bool err").Validate(false).Error())
	assert.Equal(t, "slice err", NonZero[[]int]().Errf("slice err").Validate(nil).Error())
	assert.Equal(t, "map err", NonZero[map[string]int]().Errf("map err").Validate(nil).Error())
}

func TestNonZeroStruct(t *testing.T) {
	type Test struct {
		A int
		B string
	}
	assert.Nil(t, NonZero[Test]().Validate(Test{A: 1}))
	assert.Equal(t, ErrNonZero, NonZero[Test]().Validate(Test{}))
}

func BenchmarkNonZeroRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NonZero[int]().Validate(1)
	}
}
