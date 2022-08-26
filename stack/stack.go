package stack

import "errors"

type data[T any] struct {
	element *T
	down    *data[T]
}

type Stack[T any] interface {
	Push(elem T) *data[T]
	Pop() *T
	Top() *T
	Size() int
}

func New[T any]() *data[T] {
	s := data[T]{}
	return &s
}

func (d *data[T]) Push(elem T) *data[T] {
	if d.element == nil {
		d.element = &elem
		d.down = nil
		return d
	}
	newStack := data[T]{
		element: &elem,
		down:    d,
	}
	return &newStack
}

func (d *data[T]) Pop() (T, error) {
	if d.element == nil {
		var null T
		return null, errors.New("stack is empty")
	}

	// last element
	element := d.element
	if d.down == nil {
		d.element = nil
		d.down = nil
		return *element, nil
	}

	d.element = d.down.element
	d.down = d.down.down
	return *element, nil
}

func (d *data[T]) Top() (T, error) {
	if d.element == nil {
		var null T
		return null, errors.New("stack is empty")
	}
	return *d.element, nil
}

func (d *data[T]) Size() int {
	if d.element == nil {
		return 0
	}
	size := 1
	probe := d
	for probe.down != nil {
		size++
		probe = probe.down
	}
	return size
}
