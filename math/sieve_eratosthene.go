package math

import "math"

func SieveEratosthene(limit int) []int {
	sqrtLimit := int(math.Sqrt(float64(limit)))
	result := []int{}
	sieve := getSieve(limit)
	for i := 2; i < sqrtLimit; i++ {
		if sieve[i] > 0 {
			k := 1
			for j := i*i; j < limit; j = (i*i) + (k*1) {
				sieve[j] = -sieve[j]
				k ++
			}
		}
	}
	for _, v := range sieve {
		if v > 0 {result = append(result, v); }
	}
	
	return result
}

func IsPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func getSieve(limit int) []int {
	sieve := make([]int, 0, limit)
	for i := 2 ; i < limit; i++ {
		sieve = append(sieve, i)
	}
	return sieve	
}
