package diffsquares

// SumOfSquares is 1 * 1 + 2 * 2 + ... + n * n.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// SquareOfSum is (1 + 2 + ... + n)^2.
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// Difference returns difference of (1 + 2 + ... + n)^2 - (1 * 1 + 2 * 2 + ... + n * n).
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
