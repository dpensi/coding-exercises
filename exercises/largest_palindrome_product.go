/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package exercises

import (
	"fmt"
)

type result struct {
	Num1       int
	Num2       int
	Palindrome int
}

func LargestPalindromeProduct() result {
	max := 999
	palindromes := []result{}
	for i := max; i > 100; i-- {
		for k := max; k > 100; k-- {
			mul := i * k
			mulStr := fmt.Sprintf("%v", mul)
			if mulStr == reverse(mulStr) {
				found := result{
					Num1:       i,
					Num2:       k,
					Palindrome: mul,
				}

				palindromes = append(palindromes, found)
			}
		}
	}
	return palMax(palindromes)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func palMax(pals []result) result {
	max := result{}
	for _, res := range pals {
		if res.Palindrome > max.Palindrome {
			max = res
		}
	}
	return max
}
