// Package diffsquares provides method to find the difference
// between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

//Difference returns difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

//SquareOfSum returns the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

//SumOfSquares returns the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}
