package solutions

import "fmt"

var numbersToLetters = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func NumberLetterCounts() {
	sum := 0

	for i := 1; i <= 1000; i++ {
		if i <= 99 {
			number, ok := numbersToLetters[i]

			if !ok {
				sum += len(numbersToLetters[i-i%10]) + len(numbersToLetters[i%10])
			} else {
				sum += len(number)
			}
		} else if i <= 999 {
			if i%100 == 0 {
				sum += len(numbersToLetters[i/100]) + len("hundred")
			} else {
				number, ok := numbersToLetters[i%100]

				if !ok {
					sum += len(numbersToLetters[i/100]) + len("hundred") + len("and") + len(numbersToLetters[i%100-i%100%10]) + len(numbersToLetters[i%100%10])
				} else {
					sum += len(numbersToLetters[i/100]) + len("hundred") + len("and") + len(number)
				}
			}
		} else {
			sum += len(numbersToLetters[1]) + len("thousand")
		}
	}

	fmt.Println(sum)
}
