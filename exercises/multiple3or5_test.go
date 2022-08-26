package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiples3or5(t *testing.T) {
	expected := []int{3, 5, 6, 9}
	result := Multiples3or5(10)
	assert.Equal(t, result, expected)
}
