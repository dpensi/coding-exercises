package math

import (
	"github.com/dpensi/coding-exercises/list"
)

func Collatz(n int, visited *[]int) int {
	*visited = append(*visited, n)
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return Collatz(n/2, visited)
	} else {
		return Collatz(3*n+1, visited)
	}
}

func CachedCollatz(n int) []int {
	visited := make([]int, 0)
	cache := make(map[int]*[]int, 0)
	cachedColl(n, &visited, cache)
	return visited
}

func cachedColl(n int, visited *[]int, cache map[int]*[]int) {
	if val, ok := cache[n]; ok {
		*visited = append(*visited, *(val)...)
		return
	} else {
		*visited = append(*visited, n)
		cache[n] = visited
	}

	if n == 1 {
		return
	}
	if n%2 == 0 {
		cachedColl(n/2, visited, cache)
	} else {
		cachedColl(3*n+1, visited, cache)
	}

}

func CachedCollatzList(n int) list.List[int] {
	visited := list.New[int](nil)
	cache := make(map[int]list.List[int], 0)
	cachedCollLst(n, visited, cache)
	return visited
}

func cachedCollLst(n int, visited list.List[int], cache map[int]list.List[int]) {
	if val, ok := cache[n]; ok {
		visited.Append(val)
		return
	} else {
		visited.Append(list.New(&n))
		cache[n] = visited
	}

	if n == 1 {
		return
	}
	if n%2 == 0 {
		cachedCollLst(n/2, visited, cache)
	} else {
		cachedCollLst(3*n+1, visited, cache)
	}

}
