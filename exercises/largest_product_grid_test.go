package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGrid(t *testing.T) {
	result := getGrid()

	assert.True(t, len(result) > 0)
}

func TestLargestProductGrid(t *testing.T) {
	result := LargestProductGrid()
	assert.True(t, result > 10)

}
