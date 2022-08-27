package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func PrimeThTest(t *testing.T) {
	input := 6
	expected := 13
	actual := PrimeTh(input)
	assert.Equal(t, expected, actual)
}
