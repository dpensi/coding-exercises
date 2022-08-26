package graph

type node[T any] struct {
	Value T
	arcs  []*node[T]
}

type Node[T comparable] interface {
	Grade() int
	AddArc(node[T])
	Adjacents() []*node[T]
}

func New[T comparable](value T) *node[T] {
	return &node[T]{Value: value}
}

func (n *node[T]) Grade() int {
	return len(n.arcs)
}

func (n *node[T]) AddArc(newNode *node[T]) {
	if n.arcs == nil {
		n.arcs = make([]*node[T], 0)
	}
	n.arcs = append(n.arcs, newNode)
}

func (n *node[T]) Adjacents() []*node[T] {
	return n.arcs
}
