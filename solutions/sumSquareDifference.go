package solutions

import (
	"fmt"
	"math"
)

func FindSumSquareDifference(number uint) {
	var sumOfSquares uint = (number*(number + 1) * (2*number + 1)) / 6
	var squareOfTheSum uint = uint(math.Pow(float64((number*(number + 1))/2), float64(2)))
	var difference = squareOfTheSum - sumOfSquares
	fmt.Printf("%v\n", difference)
}
