package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testSet := New[int]()
	testSet.Add(1)
	testSet.Add(2)
	assert.Equal(t, 2, testSet.Size())
	assert.True(t, testSet.Add(10))
	assert.False(t, testSet.Add(1))
}

func TestRemove(t *testing.T) {
	testSet := New[int]()
	testSet.Add(1)
	testSet.Add(2)
	testSet.Remove(1)
	assert.Equal(t, 1, testSet.Size())
	assert.Equal(t, "{2}", testSet.String())
	assert.False(t, testSet.Remove(10))
	assert.True(t, testSet.Remove(2))
	assert.Equal(t, 0, testSet.Size())
}

func TestContains(t *testing.T) {
	testSet := New[int]()
	testSet.Add(1)
	assert.True(t, testSet.Contains(1))
	assert.False(t, testSet.Contains(42))
}
func TestString(t *testing.T) {
	testSet := New[int]()
	testSet.Add(1)
	testSet.Add(2)
	assert.Equal(t, len("{1,2}"), len(testSet.String()))

	testSet0 := New[string]()
	testSet0.Add("a")
	assert.Equal(t, "{a}", testSet0.String())

	testSet1 := New[string]()
	testSet1.Add("aaaaaaaaa")
	assert.Equal(t, "{aaaaaaaaa}", testSet1.String())

	testSet2 := New[int]()
	assert.Equal(t, "{}", testSet2.String())
}
