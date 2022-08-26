/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package exercises

func EvenFibonacciNumbers(n int) int {
	if n <= 2 {
		return 2
	}
	fibonacciList := make([]int, 2, n)
	fibonacciList[0] = 1
	fibonacciList[1] = 2
	i := 2
	sum := 2
	for {
		next := fibonacciList[i-1] + fibonacciList[i-2]
		if next > n {
			break
		}

		fibonacciList = append(fibonacciList, next)

		if fibonacciList[i]%2 == 0 {
			sum += fibonacciList[i]
		}
		i++
	}
	return sum
}
