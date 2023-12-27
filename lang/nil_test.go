package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNil(t *testing.T) {
	assert.True(t, true, IsNil(nil))
	var (
		m  map[string]string
		s  []int
		c  chan int
		p  *int
		i  int
		st struct{}
	)

	assert.True(t, true, IsNil(m))
	assert.True(t, true, IsNil(s))
	assert.True(t, true, IsNil(c))
	assert.True(t, true, IsNil(p))
	assert.True(t, true, IsNil(&i))
	assert.True(t, true, IsNil(&st))
}
