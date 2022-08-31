package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	assert.True(t, IsPrime(13))
	assert.True(t, IsPrime(2))
	assert.True(t, IsPrime(3))
	assert.True(t, IsPrime(17))
	assert.True(t, IsPrime(7))

	assert.False(t, IsPrime(6))
	assert.False(t, IsPrime(15))
	assert.False(t, IsPrime(17*2))
}

func TestSieveEratosthene(t *testing.T) {
	expected := []int{2, 3, 5}
	input := 5
	result := SieveEratosthene(input)
	assert.Equal(t, result, expected)

	expected = []int{2, 3, 5, 7, 11, 13}
	input = 13
	result = SieveEratosthene(input)
	assert.Equal(t, result, expected)
}
