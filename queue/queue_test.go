package queue

import (
	"testing"

	assert "github.com/Tecu23/go-sda/internal/helpers"
)

func TestQueue(t *testing.T) {
	list := New[int]()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assert.Equal(t, *list.Deque(), 5)
	assert.Equal(t, list.length, 2)

	list.Enqueue(11)

	assert.Equal(t, *list.Deque(), 7)
	assert.Equal(t, *list.Deque(), 9)
	assert.Equal(t, *list.Peek(), 11)
	assert.Equal(t, *list.Deque(), 11)
	assert.Equal(t, list.Deque(), nil)
	assert.Equal(t, list.length, 0)

	list.Enqueue(69)
	assert.Equal(t, *list.Peek(), 69)
	assert.Equal(t, list.length, 1)
}
