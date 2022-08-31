package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsCycle(t *testing.T) {
	startNode := GetStatusGraph1()
	fatherSet := NewFathersVector[string]()
	resultNode := dfsCycle(&startNode, &startNode, fatherSet)
	assert.NotNil(t, resultNode)
}

func TestCycle(t *testing.T) {
	startNode := GetStatusGraph1()
	result := Cycle(&startNode)
	assert.Equal(t, len(result), 3)

	nodeJ := GetStatusNodeJ()
	result = Cycle(&nodeJ)
	assert.Equal(t, len(result), 3)
}
