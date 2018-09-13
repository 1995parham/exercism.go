/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-07-2018
 * |
 * | File Name:     difference_of_squares.go
 * +===============================================
 */

package diffsquares

// SumOfSquares is 1 * 1 + 2 * 2 + ... + n * n
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSums is (1 + 2 + ... + n)^2
func SquareOfSums(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// Difference returns difference of (1 + 2 + ... + n)^2 - (1 * 1 + 2 * 2 + ... + n * n)
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
