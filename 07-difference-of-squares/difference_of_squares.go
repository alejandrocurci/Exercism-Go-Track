package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the square of the first N natural numbers
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

// Difference calculates the difference between SquareOfSum(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
