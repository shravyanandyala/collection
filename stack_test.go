package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	collection "gojini.dev/template"
)

func TestStack(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Create a stack
	s := collection.NewStack()
	assert.NotNil(s)

	// Push to stack
	s.Push("hello")
	assert.Equal(1, s.Size())

	// Push again
	s.Push("world")
	assert.Equal(2, s.Size())

	// Peek at top of stack
	assert.Equal("world", s.Peek())
	assert.Equal(2, s.Size())

	// Try to pop
	assert.Equal("world", s.Pop())
	assert.Equal(1, s.Size())

	// Pop again
	assert.Equal("hello", s.Pop())
	assert.True(s.IsEmpty())

	// Try to peek and remove from empty queue, should fail
	assert.Nil(s.Peek())
	assert.Nil(s.Pop())
}
