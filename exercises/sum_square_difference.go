/*
The sum of the squares of the first ten natural numbers is,

	1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

	(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is

3025 - 385 = 2640

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

/*
	1^2 + 2^2 + ... + 10^2 = 1^2 + 2^2 + ... + (n-1)^2 + n^2
	10^2 + 9^2 + ... + 1^2 = n^2 + (n-1)^2 + ... + 1^2
	x	x	x	x	 x	  x		x	x	 x		x
	1 + 4 + 9 + 16 + 25 + 36 + 42 + 64 + 81 + 100
	101
	85
	73
	58
	61
*/

package exercises

import (
	"fmt"
	mymath "github.com/dpensi/coding-exercises/math"
	"math"
)

func SumSquareDifference(n int) int {
	squareSum := mymath.SumFirstIntegers(n)
	fmt.Printf("sum of first %d integers: %v\n", n, squareSum)
	squareSum = int(math.Pow(float64(squareSum), 2))
	fmt.Printf("square of the sum: %v\n", squareSum)
	sumSquare := 0.0
	for i := 0; i <= n; i++ {
		sumSquare += math.Pow(float64(i), 2)
	}
	fmt.Printf("sum of the squares: %v\n", sumSquare)

	return squareSum - int(sumSquare)
}
