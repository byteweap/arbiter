package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInRule(t *testing.T) {

	// string
	err := In("a", "b", "c").Validate("a")
	assert.Nil(t, err)

	// int
	err = In(1, 2, 3).Validate(1)
	assert.Nil(t, err)

	// int8
	err = In(int8(1), int8(2), int8(3)).Validate(int8(1))
	assert.Nil(t, err)

	// int16
	err = In(int16(1), int16(2), int16(3)).Validate(int16(1))
	assert.Nil(t, err)

	// int32
	err = In(int32(1), int32(2), int32(3)).Validate(int32(1))
	assert.Nil(t, err)

	// int64
	err = In(int64(1), int64(2), int64(3)).Validate(int64(1))
	assert.Nil(t, err)

	// uint
	err = In(uint(1), uint(2), uint(3)).Validate(uint(1))
	assert.Nil(t, err)

	// uint8
	err = In(uint8(1), uint8(2), uint8(3)).Validate(uint8(1))
	assert.Nil(t, err)

	// uint16
	err = In(uint16(1), uint16(2), uint16(3)).Validate(uint16(1))
	assert.Nil(t, err)

	// uint32
	err = In(uint32(1), uint32(2), uint32(3)).Validate(uint32(1))
	assert.Nil(t, err)

	// uint64
	err = In(uint64(1), uint64(2), uint64(3)).Validate(uint64(1))
	assert.Nil(t, err)

	// float32
	err = In(float32(1), float32(2), float32(3)).Validate(float32(1))
	assert.Nil(t, err)

	// float64
	err = In(float64(1), float64(2), float64(3)).Validate(float64(1))
	assert.Nil(t, err)

	// bool
	err = In(true, false).Validate(true)
	assert.Nil(t, err)

	// string slice
	err = In([]string{"a", "b", "c"}...).Validate("a")
	assert.Nil(t, err)

	// int slice
	err = In([]int{1, 2, 3}...).Validate(1)
	assert.Nil(t, err)

	// int8 slice
	err = In([]int8{1, 2, 3}...).Validate(int8(1))
	assert.Nil(t, err)

	// int16 slice
	err = In([]int16{1, 2, 3}...).Validate(int16(1))
	assert.Nil(t, err)

	// int32 slice
	err = In([]int32{1, 2, 3}...).Validate(int32(1))
	assert.Nil(t, err)

	// int64 slice
	err = In([]int64{1, 2, 3}...).Validate(int64(1))
	assert.Nil(t, err)

	// uint slice
	err = In([]uint{1, 2, 3}...).Validate(uint(1))
	assert.Nil(t, err)

	// uint8 slice
	err = In([]uint8{1, 2, 3}...).Validate(uint8(1))
	assert.Nil(t, err)

	// uint16 slice
	err = In([]uint16{1, 2, 3}...).Validate(uint16(1))
	assert.Nil(t, err)

	// uint32 slice
	err = In([]uint32{1, 2, 3}...).Validate(uint32(1))
	assert.Nil(t, err)

	// uint64 slice
	err = In([]uint64{1, 2, 3}...).Validate(uint64(1))
	assert.Nil(t, err)

	// float32 slice
	err = In([]float32{1, 2, 3}...).Validate(float32(1))
	assert.Nil(t, err)

	// float64 slice
	err = In([]float64{1, 2, 3}...).Validate(float64(1))
	assert.Nil(t, err)

	// bool slice
	err = In([]bool{true, false}...).Validate(true)
	assert.Nil(t, err)

}

func TestNotInRule(t *testing.T) {

	// string
	err := NotIn("a", "b", "c").Validate("d")
	assert.Nil(t, err)

	// int
	err = NotIn(1, 2, 3).Validate(4)
	assert.Nil(t, err)

	// int8
	err = NotIn(int8(1), int8(2), int8(3)).Validate(int8(4))
	assert.Nil(t, err)

	// int16
	err = NotIn(int16(1), int16(2), int16(3)).Validate(int16(4))
	assert.Nil(t, err)

	// int32
	err = NotIn(int32(1), int32(2), int32(3)).Validate(int32(4))
	assert.Nil(t, err)

	// int64
	err = NotIn(int64(1), int64(2), int64(3)).Validate(int64(4))
	assert.Nil(t, err)

	// uint
	err = NotIn(uint(1), uint(2), uint(3)).Validate(uint(4))
	assert.Nil(t, err)

	// uint8
	err = NotIn(uint8(1), uint8(2), uint8(3)).Validate(uint8(4))
	assert.Nil(t, err)

	// uint16
	err = NotIn(uint16(1), uint16(2), uint16(3)).Validate(uint16(4))
	assert.Nil(t, err)

	// uint32
	err = NotIn(uint32(1), uint32(2), uint32(3)).Validate(uint32(4))
	assert.Nil(t, err)

	// uint64
	err = NotIn(uint64(1), uint64(2), uint64(3)).Validate(uint64(4))
	assert.Nil(t, err)

	// float32
	err = NotIn(float32(1), float32(2), float32(3)).Validate(float32(4))
	assert.Nil(t, err)

	// float64
	err = NotIn(float64(1), float64(2), float64(3)).Validate(float64(4))
	assert.Nil(t, err)

	// string slice
	err = NotIn([]string{"a", "b", "c"}...).Validate("d")
	assert.Nil(t, err)

	// int slice
	err = NotIn([]int{1, 2, 3}...).Validate(4)
	assert.Nil(t, err)

	// int8 slice
	err = NotIn([]int8{1, 2, 3}...).Validate(int8(4))
	assert.Nil(t, err)

	// int16 slice
	err = NotIn([]int16{1, 2, 3}...).Validate(int16(4))
	assert.Nil(t, err)

	// int32 slice
	err = NotIn([]int32{1, 2, 3}...).Validate(int32(4))
	assert.Nil(t, err)

	// int64 slice
	err = NotIn([]int64{1, 2, 3}...).Validate(int64(4))
	assert.Nil(t, err)

	// uint slice
	err = NotIn([]uint{1, 2, 3}...).Validate(uint(4))
	assert.Nil(t, err)

	// uint8 slice
	err = NotIn([]uint8{1, 2, 3}...).Validate(uint8(4))
	assert.Nil(t, err)

	// uint16 slice
	err = NotIn([]uint16{1, 2, 3}...).Validate(uint16(4))
	assert.Nil(t, err)

	// uint32 slice
	err = NotIn([]uint32{1, 2, 3}...).Validate(uint32(4))
	assert.Nil(t, err)

	// uint64 slice
	err = NotIn([]uint64{1, 2, 3}...).Validate(uint64(4))
	assert.Nil(t, err)

	// float32 slice
	err = NotIn([]float32{1, 2, 3}...).Validate(float32(4))
	assert.Nil(t, err)

	// float64 slice
	err = NotIn([]float64{1, 2, 3}...).Validate(float64(4))
	assert.Nil(t, err)

}

func BenchmarkInRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		In("a", "b", "c").Validate("a")
	}
}

func BenchmarkNotInRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NotIn("a", "b", "c").Validate("d")
	}
}
