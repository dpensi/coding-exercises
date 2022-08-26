package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	testStack := New[int]()
	assert.Equal(t, 0, testStack.Size())
	assert.Nil(t, testStack.element)
	assert.Nil(t, testStack.down)
	testStack = testStack.Push(1)
	testStack = testStack.Push(2)
	assert.Equal(t, 2, testStack.Size())
}

func TestPop(t *testing.T) {
	testStack := New[int]()
	testStack = testStack.Push(1)
	testStack = testStack.Push(2)
	testStack = testStack.Push(3)
	assert.Equal(t, 3, testStack.Size())

	element, err := testStack.Pop()
	assert.Equal(t, 3, element)

	element, err = testStack.Pop()
	assert.Equal(t, 2, element)

	assert.Equal(t, 1, testStack.Size())

	element, err = testStack.Pop()
	assert.Equal(t, 1, element)

	assert.Equal(t, 0, testStack.Size())

	element, err = testStack.Pop()
	assert.Error(t, err)
}

func TestTop(t *testing.T) {
	testStack := New[string]()
	testStack = testStack.Push("aaa")
	testStack = testStack.Push("bbb")
	testStack = testStack.Push("ccc")

	element, err := testStack.Top()
	assert.Equal(t, "ccc", element)

	element, err = testStack.Pop()
	assert.Equal(t, "ccc", element)

	element, err = testStack.Top()
	assert.Equal(t, "bbb", element)

	element, err = testStack.Pop()
	assert.Equal(t, "bbb", element)

	element, err = testStack.Top()
	assert.Equal(t, "aaa", element)

	element, err = testStack.Pop()
	assert.Equal(t, "aaa", element)

	assert.Nil(t, err)

	element, err = testStack.Top()
	assert.Error(t, err)
}
