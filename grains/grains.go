package grains

import (
	"errors"
	"math"
)

// Square determines how many grains are on each square
// this function takes in an integer and returns a uint64 or error
func Square(n int) (uint64, error) {

	// edge case: if n <= 0 or > 64, return an error
	if n <= 0 || n > 64 {
		err := errors.New("invalid number")
		return 0, err
	}

	// edge case: if n is 1 return 1
	// if n == 1 {
	// 	return 1, nil
	// }
	var sum float64
	// convert float
	sum = math.Exp2(float64(n - 1))

	return uint64(sum), nil

}

// Total function  return a uint64 int by going through each square and finding the Square val,
// then appending that value to the total sum which we return.
// Total number of grains on the board
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		a, _ := Square(i)
		sum += a
	}
	return sum
}

// 64 squares on a chessboard
// square 1 starts with initial value, every following square mutliplies init value by 2

// Write code that shows:

// how many grains were on each square, and
// the total number of grains
