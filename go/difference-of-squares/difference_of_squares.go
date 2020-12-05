package diffsquares

func SquareOfSum(input int) (result int) {
	for i := 1; i <= input; i++ {
		result += i
	}

	return result * result
}

func SumOfSquares(input int) (result int) {
	for i := 1; i <= input; i++ {
		result += i * i
	}

	return result
}

func Difference(input int) int {
	squareOfSum := SquareOfSum(input)
	sumOfSquares := SumOfSquares(input)

	return squareOfSum - sumOfSquares
}
