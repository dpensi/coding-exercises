package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsCycle(t *testing.T) {
	startNode := GetStatusGraph1()
	fatherSet := NewFathersVector[string]()
	resultNode := dfsCycle(&startNode, &startNode, fatherSet)
	fmt.Println(resultNode)
}

func TestCycle(t *testing.T) {
	startNode := GetStatusGraph1()
	result := Cycle(&startNode)
	assert.Equal(t, len(result), 3)
}
