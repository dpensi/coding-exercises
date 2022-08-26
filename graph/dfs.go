package graph

import (
	"github.com/dpensi/coding-exercises/hashset"
	"github.com/dpensi/coding-exercises/stack"
)

func Dfs[T any](start node[T], operation func(n node[T])) []node[T] {
	orderedVisitedNodes := make([]node[T], 0)
	visitedNodes := hashset.New[node[T]]()
	nodesStack := stack.New[node[T]]()
	_ = visitedNodes.Add(&start)
	orderedVisitedNodes = append(orderedVisitedNodes, start)
	nodesStack = nodesStack.Push(start)
	operation(start)
	for nodesStack.Size() > 0 {
		visiting, _ := nodesStack.Top()
		for _, adj := range visiting.Adjacents() {
			if !visitedNodes.Contains(*adj) {
				orderedVisitedNodes = append(orderedVisitedNodes, *adj) // not really O(1)!
				visitedNodes.Add(adj)
				nodesStack = nodesStack.Push(*adj)
				operation(*adj)
			} else {
				_, _ = nodesStack.Pop()
			}
		}
	}

	return orderedVisitedNodes
}

func DfsRecursive[T any](start node[T], operation func(n node[T])) []node[T] {
	orderedVisitedNodes := make([]node[T], 0)
	visitedNodes := hashset.New[node[T]]()
	dfsRec(start, operation, &orderedVisitedNodes, visitedNodes)
	return orderedVisitedNodes

}

func dfsRec[T any](
	start node[T],
	operation func(n node[T]),
	orderedVisitedNodes *[]node[T],
	visitedNodes hashset.HashSet[node[T]],
) {

	*orderedVisitedNodes = append(*orderedVisitedNodes, start)
	_ = visitedNodes.Add(&start)
	operation(start)
	for _, adj := range start.Adjacents() {
		if !visitedNodes.Contains(*adj) {
			dfsRec(*adj, operation, orderedVisitedNodes, visitedNodes)
		}
	}
}
