package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testScenario struct {
	Input    int
	Expected []int
	Message  string
}

func TestPrimeFactors(t *testing.T) {
	testScenarios := []testScenario{
		{
			Input:    1234567,
			Expected: []int{127, 9721},
			Message:  "1234567",
		},
		{
			Input:    333333,
			Expected: []int{3, 3, 7, 11, 13, 37},
			Message:  "333333",
		},
		{
			Input:    987653,
			Expected: []int{29, 34057},
			Message:  "987653",
		},
	}
	for _, test := range testScenarios {
		t.Run(test.Message, func(t *testing.T) {
			assert.Equal(t, test.Expected, PrimeFactors(test.Input))
		})
	}
}
