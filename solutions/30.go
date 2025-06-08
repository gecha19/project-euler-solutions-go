package solutions

import "fmt"

func DigitFifthPowers() {
	currentNumber := 2
	sum := 0
	sieve := make([]int, 10)

	for i := 0; i <= 9; i++ {
		sieve[i] = i * i * i * i * i
	}

	// We want the maximum sum to be able to reach the smallest n-digit number, e.g., n×9⁵ ≥ 10ⁿ⁻¹
	// max sum of digits of a 6-digit number: 6×9⁵=354,294 > 100,000, and 7×9⁵=354,294 < 1,000,000
	for currentNumber < 354294 {
		firstDigit := currentNumber / 100000
		secondDigit := (currentNumber % 100000) / 10000
		thirdDigit := (currentNumber % 10000) / 1000
		fourthDigit := (currentNumber % 1000) / 100
		fifthDigit := (currentNumber % 100) / 10
		sixthDigit := currentNumber % 10

		sumOfDigits := sieve[firstDigit] + sieve[secondDigit] + sieve[thirdDigit] + sieve[fourthDigit] + sieve[fifthDigit] + sieve[sixthDigit]

		if sumOfDigits == currentNumber {
			sum += currentNumber
		}

		currentNumber += 1
	}

	fmt.Println(sum)
}
