package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleRule(t *testing.T) {
	var err error

	err = MultipleOf(2).Validate(2)
	assert.Nil(t, err)

	err = MultipleOf(2).Validate(3)
	assert.Equal(t, fmt.Errorf(ErrMultipleFormat, 2), err)
}
