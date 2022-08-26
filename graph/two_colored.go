package graph

import (
	"errors"

	"github.com/dpensi/coding-exercises/stack"
)

type coloredNode[T any] struct {
	Value T
	color int
	arcs  []*coloredNode[T]
}

func NewColored[T any](value T) *coloredNode[T] {
	return &coloredNode[T]{Value: value}
}

func (n *coloredNode[T]) Grade() int {
	return len(n.arcs)
}

func (n *coloredNode[T]) AddArc(newNode *coloredNode[T]) {
	if n.arcs == nil {
		n.arcs = make([]*coloredNode[T], 0)
	}
	n.arcs = append(n.arcs, newNode)
}

func (n *coloredNode[T]) Adjacents() []*coloredNode[T] {
	return n.arcs
}

func TwoColoration[T comparable](start coloredNode[T]) error {
	start.color = 1
	nodesStack := stack.New[coloredNode[T]]()
	nodesStack = nodesStack.Push(start)

	for nodesStack.Size() > 0 {
		visiting, _ := nodesStack.Top()
		for _, adj := range visiting.Adjacents() {
			if visiting.color == adj.color {
				return errors.New("bipartition does not exist")
			} else if adj.color == 0 {

				if visiting.color == 1 {
					adj.color = 2
				} else {
					adj.color = 1
				}
				nodesStack = nodesStack.Push(*adj)
			} else {
				_, _ = nodesStack.Pop()
			}
		}
		if len(visiting.Adjacents()) == 0 {
			_, _ = nodesStack.Pop()
		}
	}

	return nil
}
