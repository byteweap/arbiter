package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetweenRule(t *testing.T) {
	err := Between(3, 10).Validate(2)
	assert.Equal(t, fmt.Errorf(ErrBetweenFormat, 3, 10), err)

	err = Between(3, 10).Validate(5)
	assert.Nil(t, err)

	err = Between(3, 10).Validate(11)
	assert.Equal(t, fmt.Errorf(ErrBetweenFormat, 3, 10), err)

	err = Between(3.0, 10.0).Validate(2.5)
	assert.Equal(t, fmt.Errorf(ErrBetweenFormat, 3.0, 10.0), err)

	err = Between(3.0, 10.0).Validate(5.5)
	assert.Nil(t, err)

	err = Between(3.0, 10.0).Validate(10.5)
	assert.Equal(t, fmt.Errorf(ErrBetweenFormat, 3.0, 10.0), err)

	customErr := Between(3, 10).Errf("invalid range").Validate(2)
	assert.Equal(t, "invalid range", customErr.Error())
}

func BenchmarkBetweenRule(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Between(3, 10).Validate(5)
	}
}
