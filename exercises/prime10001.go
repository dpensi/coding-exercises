/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package exercises

func PrimeTh(nth int) int {
	count := 2
	currentPrime := 0
	primeCount := 0
	for primeCount != nth {
		if isPrime(count) {
			primeCount++
			currentPrime = count
		}
		count++
	}
	return currentPrime
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
