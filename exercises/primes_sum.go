/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package exercises

import (
	mymath "github.com/dpensi/coding-exercises/math"
)


func PrimesSum(n int) (sum int){
	primes := mymath.SieveEratosthene(n)
	for _,v := range primes {
		sum += v
	}
	return
} 
