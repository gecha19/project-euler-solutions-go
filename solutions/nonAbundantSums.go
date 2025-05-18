package solutions

import (
	"fmt"
)

func isAbundant(number int64) bool {
	sumOfDivisors := int64(1)

	for j := int64(2); j*j <= number; j++ {
		if number%j == 0 {
			if j*j != number {
				sumOfDivisors += j + number/j
			} else {
				sumOfDivisors += j
			}
		}
	}

	return sumOfDivisors > number
}

func FindNonAbundantsSum(limit int64) {
	abundants := make([]int64, 0)
	isAbundantsSum := make([]bool, limit + 1)
	sum := 0

	for i := int64(1); i <= limit; i++ {
		if isAbundant(i) {
			abundants = append(abundants, i)
		}
	}

	for _, i := range abundants {
		for _, j := range abundants {
			if i + j < limit + 1 {
				isAbundantsSum[i + j] = true
			}
		}
	}

	for i := range isAbundantsSum {
		if !isAbundantsSum[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
