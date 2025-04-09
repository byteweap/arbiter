package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredRule(t *testing.T) {
	var err error
	// string
	err = Required[string]().Validate("")
	assert.Equal(t, err, ErrRequired)

	// int
	err = Required[int]().Validate(0)
	assert.Equal(t, err, ErrRequired)

	// int8
	err = Required[int8]().Validate(int8(0))
	assert.Equal(t, err, ErrRequired)

	// int16
	err = Required[int16]().Validate(int16(0))
	assert.Equal(t, err, ErrRequired)

	// int32
	err = Required[int32]().Validate(int32(0))
	assert.Equal(t, err, ErrRequired)

	// int64
	err = Required[int64]().Validate(int64(0))
	assert.Equal(t, err, ErrRequired)

	// uint
	err = Required[uint]().Validate(uint(0))
	assert.Equal(t, err, ErrRequired)

	// uint8
	err = Required[uint8]().Validate(uint8(0))
	assert.Equal(t, err, ErrRequired)

	// uint16
	err = Required[uint16]().Validate(uint16(0))
	assert.Equal(t, err, ErrRequired)

	// uint32
	err = Required[uint32]().Validate(uint32(0))
	assert.Equal(t, err, ErrRequired)

	// uint64
	err = Required[uint64]().Validate(uint64(0))
	assert.Equal(t, err, ErrRequired)

	// float32
	err = Required[float32]().Validate(float32(0.0))
	assert.Equal(t, err, ErrRequired)

	// float64
	err = Required[float64]().Validate(float64(0.0))
	assert.Equal(t, err, ErrRequired)

	// nil pointer
	err = Required[*int]().Validate(nil)
	assert.Equal(t, err, ErrRequired)
}

func BenchmarkRequiredRule(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		// string
		Required[string]().Validate("")
		// int
		Required[int]().Validate(0)
		// int8
		Required[int8]().Validate(int8(0))
		// int16
		Required[int16]().Validate(int16(0))

		// int32
		Required[int32]().Validate(int32(0))

		// int64
		Required[int64]().Validate(int64(0))

		// uint
		Required[uint]().Validate(uint(0))
		// uint8
		Required[uint8]().Validate(uint8(0))
		// uint16
		Required[uint16]().Validate(uint16(0))

		// uint32
		Required[uint32]().Validate(uint32(0))

		// uint64
		Required[uint64]().Validate(uint64(0))

		// float32
		Required[float32]().Validate(float32(0.0))

		// float64
		Required[float64]().Validate(float64(0.0))

		// pointer
		Required[*int]().Validate(nil)

		// nil pointer
		Required[*int]().Validate(nil)
	}
}
