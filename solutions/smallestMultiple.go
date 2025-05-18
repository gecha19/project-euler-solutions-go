package solutions

import (
	"fmt"
	"math"
)

func FindSmallestMultiple(start uint, end uint) {
	var i uint = start
	var smallestMultiple uint = 1

	for i <= end {
		if isPrime(i) {
			var n uint = 1

			for uint(math.Pow(float64(i), float64(n))) <= end {
				n += 1
			}
			
			smallestMultiple *= uint(math.Pow(float64(i), float64(n - 1)))
		}

		i += 1
	}

	fmt.Printf("%v\n", smallestMultiple)
}
