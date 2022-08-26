package math

func PrimeFactors(n int) []int {
	result := make([]int, 0)
	mod, div := 0, 0
	for i := 2; i <= n; {
		div = n / i
		mod = n % i
		for mod == 0 {
			result = append(result, i)
			n = div
			div = n / i
			mod = n % i
		}
		i++
	}
	return result
}
