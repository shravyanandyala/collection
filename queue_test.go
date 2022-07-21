package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	collection "gojini.dev/template"
)

func TestQueue(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	// Create a queue
	q := collection.NewQueue()
	assert.NotNil(q)

	// Add to queue
	q.Add("hello")
	assert.Equal(1, q.Size())

	// Add again
	q.Add("world")
	assert.Equal(2, q.Size())

	// Peek at head of queue
	assert.Equal("hello", q.Peek())
	assert.Equal(2, q.Size())

	// Try to remove
	assert.Equal("hello", q.Remove())
	assert.Equal(1, q.Size())

	// Remove again
	assert.Equal("world", q.Remove())
	assert.True(q.IsEmpty())

	// Try to peek and remove from empty queue, should fail
	assert.Nil(q.Peek())
	assert.Nil(q.Remove())
}
