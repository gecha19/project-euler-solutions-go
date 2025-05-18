package solutions

import (
	"fmt"
	"slices"
)

func findDivisorSum(number int64) int64 {
	sum := int64(1)

	for i := int64(2); i*i <= number; i++ {
		if number%i == 0 {
			sum += i + (number / i)
		}
	}

	return sum
}

func FindAmicableNumbers() {
	sum := int64(0)
	seen := make([]int64, 0)

	for i := int64(1); i < 10000; i++ {
		if slices.Contains(seen, i) {
			continue
		}

		divisorSum := findDivisorSum(i)

		if divisorSum < 10000 && i == findDivisorSum(divisorSum) && i != divisorSum {
			seen = append(seen, divisorSum)
			sum += i + divisorSum
		}
	}

	fmt.Println(sum)
}
