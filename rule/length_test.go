package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthRule(t *testing.T) {
	var err error
	err = Len[string](1, 10).Validate("哈")
	assert.Nil(t, err)

	str := "hello"
	err = Len[*string](1, 10).Validate(&str)
	assert.Nil(t, err)

	err = Len[[]int](1, 10).Validate([]int{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*int](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]int8](2, 3).Validate([]int8{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*int8](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]int16](2, 3).Validate([]int16{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*int16](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]int32](2, 3).Validate([]int32{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*int32](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]int64](2, 3).Validate([]int64{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*int64](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]uint](2, 3).Validate([]uint{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*uint](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]uint8](2, 3).Validate([]uint8{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*uint8](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]uint16](2, 3).Validate([]uint16{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*uint16](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]uint32](2, 3).Validate([]uint32{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*uint32](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]uint64](2, 3).Validate([]uint64{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*uint64](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]float32](2, 3).Validate([]float32{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*float32](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]float64](2, 3).Validate([]float64{1, 2, 3})
	assert.Nil(t, err)

	err = Len[[]*float64](0, 10).Validate(nil)
	assert.Nil(t, err)

	err = Len[[]bool](2, 3).Validate([]bool{true, false, true})
	assert.Nil(t, err)

	err = Len[[]*bool](0, 10).Validate(nil)
	assert.Nil(t, err)

	type Test struct {
		ID    int
		Name  string
		Age   int
		Email string
		Phone string
	}
	arr := []Test{
		{ID: 1, Name: "test", Age: 10, Email: "test@test.com", Phone: "1234567890"},
		{ID: 2, Name: "test", Age: 10, Email: "test@test.com", Phone: "1234567890"},
		{ID: 3, Name: "test", Age: 10, Email: "test@test.com", Phone: "1234567890"},
	}
	err = Len[[]Test](1, 10).Validate(arr)
	assert.Nil(t, err)

}

func BenchmarkLengthRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		Len[string](1, 10).Validate("哈")

		// str := "hello"
		// Len[*string](1, 10).Validate(&str)

		// Len[[]int](1, 10).Validate([]int{1, 2, 3})

		// Len[[]*int](0, 10).Validate(nil)

		// Len[[]int8](2, 3).Validate([]int8{1, 2, 3})

		// Len[[]*int8](0, 10).Validate(nil)

		// Len[[]int16](2, 3).Validate([]int16{1, 2, 3})

		// Len[[]*int16](0, 10).Validate(nil)

		// Len[[]int32](2, 3).Validate([]int32{1, 2, 3})

		// Len[[]*int32](0, 10).Validate(nil)

		// Len[[]int64](2, 3).Validate([]int64{1, 2, 3})

		// Len[[]*int64](0, 10).Validate(nil)

		// Len[[]uint](2, 3).Validate([]uint{1, 2, 3})

		// Len[[]*uint](0, 10).Validate(nil)

		// Len[[]uint8](2, 3).Validate([]uint8{1, 2, 3})

		// Len[[]*uint8](0, 10).Validate(nil)

		// Len[[]uint16](2, 3).Validate([]uint16{1, 2, 3})

		// Len[[]*uint16](0, 10).Validate(nil)

		// Len[[]uint32](2, 3).Validate([]uint32{1, 2, 3})

		// Len[[]*uint32](0, 10).Validate(nil)

		// Len[[]uint64](2, 3).Validate([]uint64{1, 2, 3})

		// Len[[]*uint64](0, 10).Validate(nil)

		// Len[[]float32](2, 3).Validate([]float32{1, 2, 3})

		// Len[[]*float32](0, 10).Validate(nil)

		// Len[[]float64](2, 3).Validate([]float64{1, 2, 3})

		// Len[[]*float64](0, 10).Validate(nil)

		// Len[[]bool](2, 3).Validate([]bool{true, false, true})

		// Len[[]*bool](0, 10).Validate(nil)
	}
}
