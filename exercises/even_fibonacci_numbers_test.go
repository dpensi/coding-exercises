package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenFibonacciNumbers(t *testing.T) {
	expected := 2 + 8 + 34
	result := EvenFibonacciNumbers(89)
	assert.Equal(t, result, expected)
	assert.Equal(t, 4613732, EvenFibonacciNumbers(4000000))
}
