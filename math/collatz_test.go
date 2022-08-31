package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollatz(t *testing.T) {
	expected := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	visited := make([]int, 0, len(expected))
	_ = Collatz(13, &visited)
	assert.Equal(t, expected, visited)
}

func TestCachedCollatz(t *testing.T) {
	expected := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	visited := CachedCollatz(13)
	assert.Equal(t, expected, visited)
}

func TestCachedCollatzList(t *testing.T) {
	expected := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	visited := CachedCollatzList(13)
	i := 0
	for visited.HasNext() {
		assert.Equal(t, *visited.Data(), expected[i])
		visited = visited.Next()
		i++
	}
}

func BenchmarkCollatz(b *testing.B) {
	for i := 1; i < b.N; i++ {
		visited := make([]int, 0)
		_ = Collatz(i, &visited)
	}
}

func BenchmarkCachedCollatz(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_ = CachedCollatz(i)
	}
}

func BenchmarkCachedCollatzList(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_ = CachedCollatzList(i)
	}
}
