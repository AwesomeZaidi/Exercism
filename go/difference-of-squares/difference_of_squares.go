package diffsquares

//SumOfSquares is the sum of squares of length numbers
func SumOfSquares(length int) int {
	var sqSums int
	for i := 0; i <= length; i++ {
		sqSums += i * i
	}
	return sqSums
}

//SquareOfSum sums up numbers in slice and then sqaures them
func SquareOfSum(length int) int {
	var sum int
	for i := 0; i <= length; i++ {
		sum += i
	}
	return sum * sum
}

//Difference gives the difference of squares
func Difference(length int) int {
	numOne := SumOfSquares(length)
	numTwo := SquareOfSum(length)

	var answer int
	answer = numTwo - numOne
	return answer
}
