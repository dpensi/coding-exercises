package graph

import "github.com/dpensi/coding-exercises/hashset"

type times[T any] struct {
	start  int
	finish int
	Value  T
}

func (t *times[T]) Start() int {
	return t.start
}

func (t *times[T]) Finish() int {
	return t.finish
}

func Times[T comparable](start node[T]) hashset.HashSet[times[T]] {
	times := hashset.New[times[T]]()
	counter := 0
	dfsTimesRec(start, times, &counter)
	return times
}

func dfsTimesRec[T comparable](start node[T], currentTimes hashset.HashSet[times[T]], counter *int) {
	*counter++
	time := times[T]{
		start: *counter,
		Value: start.Value,
	}
	currentTimes.Add(&time)
	for _, adj := range start.Adjacents() {
		adjTime := times[T]{Value: adj.Value}
		if !currentTimes.Contains(adjTime) {
			dfsTimesRec(*adj, currentTimes, counter)
		}
	}
	time.finish = *counter
}
