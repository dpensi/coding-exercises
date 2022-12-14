/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package exercises

import (
	mymath "github.com/dpensi/coding-exercises/math"
)

func LongestCollatzSequence(n int) int {
	maxSequenceLen := 0
	result := 0
	var visited []int
	for i := 1; i < n; i++ {
		//visited := make([]int, 0)
		//_ = mymath.Collatz(i, &visited)
		visited = mymath.CachedCollatz(i)
		if len(visited) > maxSequenceLen {
			maxSequenceLen = len(visited)
			result = i
		}
	}
	return result
}
