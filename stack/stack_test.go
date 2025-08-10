package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := New[int]()
	for i := range 5 {
		s.Push(i)
	}

	for i := 4; i >= 0; i-- {
		val, err := s.Top()
		assert.Nil(t, err)
		assert.Equal(t, i, *val)
		s.Pop()
	}
	_, err := s.Top()
	assert.ErrorIs(t, err, StackEmpty)

	assert.True(t,s.IsEmpty())
}
