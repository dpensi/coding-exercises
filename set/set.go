package set

import "fmt"

type data[T comparable] map[T]struct{}

type Set[T comparable] interface {
	Add(elem T) bool
	Remove(elem T) bool
	Contains(elem T) bool
	String() string
	Size() int
}

func New[T comparable]() data[T] {
	emptyData := make(data[T])
	return emptyData
}

func (d data[T]) Add(elem T) bool {
	if _, ok := d[elem]; ok {
		// element already present
		return false
	}
	d[elem] = struct{}{}
	return true
}

func (d data[T]) Remove(elem T) bool {
	if _, ok := d[elem]; ok {
		delete(d, elem)
		return true
	}
	return false
}

func (d data[T]) Contains(elem T) bool {
	_, ok := d[elem]
	return ok
}
func (d data[T]) String() string {

	setToString := "{"
	for k, _ := range d {
		setToString += fmt.Sprintf("%v,", k)
	}
	if len(setToString) < 2 {
		return "{}"
	}
	return setToString[0:len(setToString)-1] + "}"
}

func (d data[T]) Size() int {
	return len(d)
}
