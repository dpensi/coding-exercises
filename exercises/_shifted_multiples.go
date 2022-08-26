/*
For a positive integer n, let s(n) be the integer obtained by shifting the leftmost digit of the decimal representation of n to the rightmost position.
For example, s(142857) = 428571 and s(10) = 1 .

For a positive rational number r, we define N(r) as the smallest positive integer n such that s(n) = r * n .
If no such integer exists, then N(r) is defined as zero.
For example, N(3) = 142857 , N(1/10) = 10 and N(2) = 0.

Let
be the sum of where ranges over all ordered pairs of coprime positive integers not exceeding .
For example,

.

Find
. Give your answer modulo . 

package shifted-multiples

func S(n int) int {
	fmt.Println("hello world")
} 
