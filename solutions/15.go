package solutions

import (
	"fmt"
)

func factorialOf(number float64) float64 {
	if number == 0 {
		return 1
	}

	if number == 1 {
		return 1
	}

	return number * factorialOf(number-1)
}

func LatticePaths(m float64, n float64) {
	/* each path consists of m downs and n rights */
	/* number of lattice paths is the number of combinations of m objects out of m + n objects */
	numerator := factorialOf(m + n)
	denominator := factorialOf(m) * factorialOf(n)
	fmt.Println(numerator / denominator)
}
