/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

/*
11 12 13 14 15 16 17 18 19 20
*/
package exercises

func SmallestMultipleBruteForce(n int) int {
	smallestMultiple := 1
	divisible := false
	for !divisible {
		divisible = true
		for i := 1; i <= n; i++ {
			divisible = divisible && (smallestMultiple%i == 0)
			if !divisible {
				smallestMultiple++
				break
			}
		}
	}
	return smallestMultiple
}

func SmallestMultiple20() int { //WRONG
	return 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20
}
