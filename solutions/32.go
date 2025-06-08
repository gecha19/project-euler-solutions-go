package solutions

import (
	"fmt"
)

func isPandigitalTriple(a, b, c int) bool {
	digits := [10]bool{true}

	for i := a; i > 0; i /= 10 {
		digit := i % 10

		if digits[digit] || digit == 0 {
			return false
		} else {
			digits[digit] = true
		}
	}

	for i := b; i > 0; i /= 10 {
		digit := i % 10

		if digits[digit] || digit == 0 {
			return false
		} else {
			digits[digit] = true
		}
	}

	for i := c; i > 0; i /= 10 {
		digit := i % 10

		if digits[digit] || digit == 0 {
			return false
		} else {
			digits[digit] = true
		}
	}

	return true
}

func PandigitalProducts() {
	seenProducts := make(map[int]bool)

	// 1 digit ⨯ 4 digits = 4 digits
	for multiplicand := 1; multiplicand <= 9; multiplicand++ {
		for multiplier := 1000; multiplier <= 9999; multiplier++ {
			product := multiplicand * multiplier

			if seenProducts[product] || product > 9999 {
				continue
			}

			if isPandigitalTriple(multiplicand, multiplier, product) {
				seenProducts[product] = true
			}
		}
	}

	// 2 digits ⨯ 3 digits = 4 digits
	for multiplicand := 10; multiplicand <= 99; multiplicand++ {
		for multiplier := 100; multiplier <= 999; multiplier++ {
			product := multiplicand * multiplier

			if seenProducts[product] || product > 9999 {
				continue
			}

			if isPandigitalTriple(multiplicand, multiplier, product) {
				seenProducts[product] = true
			}
		}
	}

	sum := 0

	for number := range seenProducts {
		sum += number
	}

	fmt.Println(sum)
}
