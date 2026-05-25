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

func TestNilRuleAllPointerTypes(t *testing.T) {
	assert.Nil(t, Nil.Validate((*int)(nil)))
	assert.Nil(t, Nil.Validate((*int8)(nil)))
	assert.Nil(t, Nil.Validate((*int16)(nil)))
	assert.Nil(t, Nil.Validate((*int32)(nil)))
	assert.Nil(t, Nil.Validate((*int64)(nil)))
	assert.Nil(t, Nil.Validate((*uint)(nil)))
	assert.Nil(t, Nil.Validate((*uint8)(nil)))
	assert.Nil(t, Nil.Validate((*uint16)(nil)))
	assert.Nil(t, Nil.Validate((*uint32)(nil)))
	assert.Nil(t, Nil.Validate((*uint64)(nil)))
	assert.Nil(t, Nil.Validate((*float32)(nil)))
	assert.Nil(t, Nil.Validate((*float64)(nil)))
	assert.Nil(t, Nil.Validate((*bool)(nil)))

	// value types (non-nil)
	assert.Error(t, Nil.Validate(0))
	assert.Error(t, Nil.Validate(int8(0)))
	assert.Error(t, Nil.Validate(int16(0)))
	assert.Error(t, Nil.Validate(int32(0)))
	assert.Error(t, Nil.Validate(int64(0)))
	assert.Error(t, Nil.Validate(uint(0)))
	assert.Error(t, Nil.Validate(uint8(0)))
	assert.Error(t, Nil.Validate(uint16(0)))
	assert.Error(t, Nil.Validate(uint32(0)))
	assert.Error(t, Nil.Validate(uint64(0)))
	assert.Error(t, Nil.Validate(float32(0)))
	assert.Error(t, Nil.Validate(float64(0)))
	assert.Error(t, Nil.Validate(false))
}

func TestNilRuleReflectBranches(t *testing.T) {
	var m map[string]int
	assert.Error(t, NotNil.Validate(m))
	assert.Nil(t, Nil.Validate(m))
	m2 := make(map[string]int)
	assert.Nil(t, NotNil.Validate(m2))
	assert.Error(t, Nil.Validate(m2))

	var ch chan int
	assert.Error(t, NotNil.Validate(ch))
	assert.Nil(t, Nil.Validate(ch))
	ch2 := make(chan int)
	assert.Nil(t, NotNil.Validate(ch2))
	assert.Error(t, Nil.Validate(ch2))

	var fn func()
	assert.Error(t, NotNil.Validate(fn))
	assert.Nil(t, Nil.Validate(fn))
	fn2 := func() {}
	assert.Nil(t, NotNil.Validate(fn2))
	assert.Error(t, Nil.Validate(fn2))

	assert.Error(t, Nil.Validate(struct{}{}))
	assert.Error(t, NotNil.Validate(struct{}{}))
}

func BenchmarkNilRule(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Nil.Validate(nil)
	}
}
