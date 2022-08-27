package graph

import "github.com/dpensi/coding-exercises/utils"

type statusNode[T any] struct {
	Value  T
	status int // 0 = not visited, -1 visit started, 1 visit finished
	arcs   []*statusNode[T]
}

func NewStatusNode[T any](value T) *statusNode[T] {
	return &statusNode[T]{Value: value}
}

func (n *statusNode[T]) Grade() int {
	return len(n.arcs)
}

func (n *statusNode[T]) AddArc(newNode *statusNode[T]) {
	if n.arcs == nil {
		n.arcs = make([]*statusNode[T], 0)
	}
	n.arcs = append(n.arcs, newNode)
}

func (n *statusNode[T]) Adjacents() []*statusNode[T] {
	return n.arcs
}

func (n *statusNode[T]) Status() int {
	return n.status
}

func (n *statusNode[T]) SetStatus(s int) {
	s = utils.Clamp(s, -1, 1)
	n.status = s
}

type fathersVector[T any] map[string]*statusNode[T]

type FathersSet[T any] interface {
	AddParent(*statusNode[T], *statusNode[T])
	GetParent(*statusNode[T]) *statusNode[T]
}

func NewFathersVector[T any]() FathersSet[T] {
	return &fathersVector[T]{}
}

func (fv fathersVector[T]) AddParent(parent *statusNode[T], son *statusNode[T]) {
	fv[utils.Hash(*son)] = parent
}

func (fv fathersVector[T]) GetParent(son *statusNode[T]) *statusNode[T] {
	key := utils.Hash(*son)
	value, ok := fv[key]
	if ok {
		return value
	}
	return nil
}

func dfsCycle[T comparable](son, parent *statusNode[T], fathers FathersSet[T]) *statusNode[T] {
	parent.SetStatus(-1)
	fathers.AddParent(parent, son)
	previous := parent
	for _, adj := range son.Adjacents() {
		if adj.Value == previous.Value {
			continue
		}
		adjParent := fathers.GetParent(adj)
		if adjParent == nil || adjParent.Status() == 0 {
			cycle := dfsCycle(adj, son, fathers)
			if cycle != nil { // cycle already found
				fathers.GetParent(son).SetStatus(
					-fathers.GetParent(son).Status())
				return cycle
			}
		} else if adjParent != nil && adj.Status() < 0 {
			fathers.GetParent(adj).SetStatus(0)
			fathers.GetParent(son).SetStatus(1)
			son.SetStatus(1)
			return son
		}
	}
	parent.SetStatus(1)
	return nil
}

func Cycle[T comparable](n *statusNode[T]) []statusNode[T] {
	fatherSet := NewFathersVector[T]()
	cycleNode := dfsCycle(n, n, fatherSet)
	result := []statusNode[T]{*cycleNode}

	for cycleNode.Status() > 0 {
		cycleNode = fatherSet.GetParent(cycleNode)
		result = append(result, *cycleNode)
	}

	return result
}
