package hashset

import (
	"testing"

	"github.com/dpensi/coding-exercises/utils"
	"github.com/stretchr/testify/assert"
)

var (
	zero = 0
	one  = 1
	two  = 2
	a    = "a"
	b    = "b"
	c    = "c"
)

func TestNew(t *testing.T) {
	testSet := hashMap[int]{}
	val, ok := testSet["dummy"]
	assert.False(t, ok)
	assert.True(t, nil == val)
}

func TestHash(t *testing.T) {
	testSet := hashMap[int]{}
	integer := 1
	testSet["ciao"] = &integer
	hashed := utils.Hash(testSet)
	assert.Equal(t, len(hashed), 64)

	a := utils.Hash("a")
	b := utils.Hash("b")
	assert.NotEqual(t, a, b)
}

func TestAdd(t *testing.T) {
	testSet := New[int]()
	zero := 0
	one := 1
	assert.True(t, testSet.Add(&zero))
	assert.True(t, testSet.Add(&one))
	assert.False(t, testSet.Add(&one))
	assert.Equal(t, testSet.Size(), 2)
}

func TestRemove(t *testing.T) {
	testSet := New[int]()
	assert.True(t, testSet.Add(&zero))
	assert.True(t, testSet.Add(&one))
	assert.True(t, testSet.Add(&two))
	assert.Equal(t, testSet.Size(), 3)

	assert.True(t, testSet.Remove(2))
	assert.True(t, testSet.Remove(1))
	assert.True(t, testSet.Remove(0))
	assert.False(t, testSet.Remove(0))
	assert.Equal(t, testSet.Size(), 0)
}

func TestContains(t *testing.T) {
	testSet := New[string]()
	assert.True(t, testSet.Add(&a))
	assert.True(t, testSet.Add(&b))
	assert.True(t, testSet.Add(&c))

	assert.True(t, testSet.Contains("a"))
	assert.False(t, testSet.Contains("aa"))
}
