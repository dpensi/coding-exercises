package math

import "math"

func SieveEratosthene(limit int) []int {
	sqrtLimit := getLimit(limit)
	result := []int{}
	sieve := getSieve(limit)
	for i := 2; i <= sqrtLimit; i++ {
		if sieve[i] > 0 {
			for j := i * i; j <= limit; j += i {
				if sieve[j] > 0 {
					sieve[j] = -sieve[j]
				}
			}
		}
	}
	for i := 2; i < len(sieve); i++ {
		if sieve[i] > 0 {
			result = append(result, sieve[i])
		}
	}

	return result
}

func IsPrime(n int) bool {
	limit := getLimit(n)
	for i := 2; i < limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getSieve(limit int) []int {
	sieve := make([]int, 0, limit)
	for i := 0; i <= limit; i++ {
		sieve = append(sieve, i)
	}
	return sieve
}

func getLimit(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n))))

}
