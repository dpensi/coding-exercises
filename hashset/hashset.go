package hashset

import "github.com/dpensi/coding-exercises/utils"

type hashMap[T any] map[string]*T

type HashSet[T any] interface {
	Add(elem *T) bool
	Remove(elem T) bool
	Contains(elem T) bool
	Size() int
	Values() []T
}

func (m hashMap[T]) Add(elem *T) bool {
	hashed := utils.Hash(elem)
	if _, ok := m[hashed]; ok {
		return false
	}
	m[hashed] = elem
	return true
}

func (m hashMap[T]) Remove(elem T) bool {
	hashed := utils.Hash(elem)
	if _, ok := m[hashed]; ok {
		delete(m, hashed)
		return true
	}
	return false
}

func (m hashMap[T]) Contains(elem T) bool {
	hashed := utils.Hash(elem)
	_, ok := m[hashed]

	return ok
}

func (m hashMap[T]) Size() int {
	return len(m)
}

func (m hashMap[T]) Values() []T {
	list := make([]T, 0, len(m))
	for _, v := range m {
		list = append(list, *v)
	}
	return list
}

func New[T any]() HashSet[T] {
	return make(hashMap[T], 0)
}
