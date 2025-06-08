package solutions

import (
	"fmt"
)

func isDigitCancellingFraction(numerator, denominator int) bool {
	a := numerator % 10
	b := (numerator / 10) % 10
	c := denominator % 10
	d := (denominator / 10) % 10

	if c == 0 || d == 0 {
		return false
	}

	if a == c {
		return b*denominator == d*numerator
	} else if a == d {
		return b*denominator == c*numerator
	} else if b == c {
		return a*denominator == d*numerator
	} else if b == d {
		return a*denominator == c*numerator
	} else {
		return false
	}
}

func findGCF(num1, num2 int) int {
	minNumber := num1
	maxNumber := num2

	if num2 < num1 {
		minNumber = num2
		maxNumber = num1
	}

	if maxNumber%minNumber == 0 {
		return minNumber
	}

	for minNumber != 0 {
		maxNumber = minNumber
		minNumber = maxNumber & minNumber
	}

	return maxNumber
}

func DigitCancellingFractions() {
	numerator := 1
	denominator := 1

	for i := 10; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if isDigitCancellingFraction(i, j) {
				numerator *= i
				denominator *= j
			}
		}
	}

	gcf := findGCF(numerator, denominator)
	fmt.Println(denominator / gcf)
}
