package grains

import (
	"errors"
	"math"
)

// Total calculates the total amount of grains in the chessboard
func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		value, _ := Square(i)
		total += value
	}
	return total
}

// Square calculates the amount of grains in s specific square of the chessboard (between 1 and 64)
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square value must be between 1 and 64")
	}
	grains := math.Pow(2, float64(square-1))
	return uint64(grains), nil
}
