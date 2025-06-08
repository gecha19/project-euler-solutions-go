package solutions

import "fmt"

func findFactorials() [10]int {
	factorials := [10]int{}
	factorials[0] = 1

	for i := 1; i <= 9; i++ {
		factorials[i] = factorials[i-1] * i
	}

	return factorials
}

func DigitFactorials() {
	number := 10
	sum := 0
	factorials := findFactorials()

	/*
		→ Max sum of an n-digit number is n × 9!:
		• For an 7-digit number, 7 × 9! = 2,540,160 > 1,000,000 (smallest 7-digit number)
		• For an 8-digit number, 8 × 9! = 2,903,040 < 10,000,000 (smallest 8-digit number)
	*/
	for number <= 7*factorials[9] {
		sumOfFactorials := 0

		for i := number; i > 0; i /= 10 {
			sumOfFactorials += factorials[i%10]
		}

		if sumOfFactorials == number {
			sum += number
			fmt.Println(sum)
		}

		number++
	}
}
