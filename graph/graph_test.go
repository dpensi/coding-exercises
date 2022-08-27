package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printNodeValue[T comparable](n any) {
	if print, ok := n.(node[T]); ok {
		fmt.Println(print)
		return
	}
	panic("node conversion error")
}

func TestBuild(t *testing.T) {
	testNode := node[int]{
		Value: 1,
		arcs:  make([]*node[int], 0),
	}

	assert.Equal(t, testNode.Value, 1)
	assert.Equal(t, testNode.Grade(), 0)
}

func TestAddArc(t *testing.T) {
	node0 := node[int]{
		Value: 0,
		arcs:  make([]*node[int], 0),
	}
	node1 := node[int]{
		Value: 1,
		arcs:  make([]*node[int], 0),
	}
	node2 := node[int]{
		Value: 2,
		arcs:  make([]*node[int], 0),
	}

	node0.AddArc(&node1)
	node0.AddArc(&node2)
	assert.Equal(t, node0.Grade(), 2)
}

func TestAdjacents(t *testing.T) {
	node0 := node[int]{
		Value: 0,
		arcs:  make([]*node[int], 0),
	}
	node1 := node[int]{
		Value: 1,
		arcs:  make([]*node[int], 0),
	}
	node2 := node[int]{
		Value: 2,
		arcs:  make([]*node[int], 0),
	}

	node0.AddArc(&node1)
	node0.AddArc(&node2)

	assert.Equal(t, len(node0.Adjacents()), 2)
	assert.Equal(t, node0.Adjacents()[0].Value, 1)
	assert.Equal(t, node0.Adjacents()[1].Value, 2)
}

func TestDfs(t *testing.T) {
	nodeA := node[string]{Value: "a"}
	nodeB := node[string]{Value: "b"}
	nodeC := node[string]{Value: "c"}
	nodeD := node[string]{Value: "d"}
	nodeE := node[string]{Value: "e"}
	nodeF := node[string]{Value: "f"}
	nodeG := node[string]{Value: "g"}
	nodeH := node[string]{Value: "h"}
	nodeI := node[string]{Value: "i"}

	nodeA.AddArc(&nodeB)
	nodeA.AddArc(&nodeC)
	nodeA.AddArc(&nodeD)
	assert.Equal(t, nodeA.Grade(), 3)
	nodeB.AddArc(&nodeA)
	nodeB.AddArc(&nodeC)
	nodeB.AddArc(&nodeG)
	nodeC.AddArc(&nodeB)
	nodeC.AddArc(&nodeA)
	nodeC.AddArc(&nodeG)
	nodeD.AddArc(&nodeA)
	nodeD.AddArc(&nodeF)
	nodeD.AddArc(&nodeI)
	nodeE.AddArc(&nodeH)
	nodeE.AddArc(&nodeD)
	nodeE.AddArc(&nodeI)
	nodeF.AddArc(&nodeI)
	nodeF.AddArc(&nodeD)
	assert.Equal(t, nodeF.Grade(), 2)
	nodeG.AddArc(&nodeB)
	nodeG.AddArc(&nodeC)
	nodeH.AddArc(&nodeE)
	assert.Equal(t, nodeH.Grade(), 1)
	assert.Equal(t, len(nodeH.Adjacents()), 1)
	nodeI.AddArc(&nodeD)
	nodeI.AddArc(&nodeF)
	nodeI.AddArc(&nodeE)

	visited := Dfs(nodeA, printNodeValue[string])
	assert.Equal(t, len(visited), 9)

	visitedRec := DfsRecursive(nodeA, printNodeValue[string])
	assert.Equal(t, len(visitedRec), 9)
}

func TestTwoColoration(t *testing.T) {
	nodeA := NewColored("a")
	nodeB := NewColored("b")
	nodeC := NewColored("c")
	nodeA.AddArc(nodeB)
	nodeB.AddArc(nodeC)

	assert.Nil(t, TwoColoration(*nodeA))

	node1 := NewColored(1)
	node2 := NewColored(2)
	node3 := NewColored(3)
	node1.AddArc(node2)
	node2.AddArc(node3)
	node3.AddArc(node1)
	assert.Error(t, TwoColoration(*node1))

	node1 = NewColored(1)
	node2 = NewColored(2)
	node3 = NewColored(3)
	node4 := NewColored(4)
	node1.AddArc(node2)
	node2.AddArc(node3)
	node3.AddArc(node4)
	node4.AddArc(node1)
	assert.NoError(t, TwoColoration(*node1))
}
