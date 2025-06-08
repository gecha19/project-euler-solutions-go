package solutions

import (
	"fmt"
	"strconv"
)

func isBinaryPalindrome(number int) bool {
	binary := ""
	reversedBinary := ""

	for number > 0 {
		remainder := number % 2
		binary = strconv.Itoa(remainder) + binary
		reversedBinary = reversedBinary + strconv.Itoa(remainder)
		number /= 2
	}

	return reversedBinary == binary
}

func DoubleBasePalindromes() {
	sum := 0

	// Generate 1-digit and 2-digit palindromes
	for i := 1; i <= 9; i++ {
		oneDigit := i
		twoDigits := 11 * i

		if isBinaryPalindrome(oneDigit) {
			sum += oneDigit
		}

		if isBinaryPalindrome(twoDigits) {
			sum += twoDigits
		}
	}

	// Generate 3-digit and 4-digit palindromes
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			threeDigits := 101*i + 10*j
			fourDigits := 1001*i + 110*j

			if isBinaryPalindrome(threeDigits) {
				sum += threeDigits
			}

			if isBinaryPalindrome(fourDigits) {
				sum += fourDigits
			}
		}
	}

	// Generate 5-digit and 6-digit palindromes
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				fiveDigits := 10001*i + 1010*j + 100*k
				sixDigits := 100001*i + 10010*j + 1100*k

				if isBinaryPalindrome(fiveDigits) {
					sum += fiveDigits
				}

				if isBinaryPalindrome(sixDigits) {
					sum += sixDigits
				}
			}
		}
	}

	fmt.Println(sum)
}
