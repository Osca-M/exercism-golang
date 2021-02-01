package diffsquares

// SquareOfSum takes in a integer n and returns the square of the first n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares takes in a integer n and returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference takes in integer n and finds the difference between the squares of the sum of the first n natural numbers
//and sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
