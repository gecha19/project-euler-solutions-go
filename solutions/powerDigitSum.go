package solutions

import (
	"fmt"
	"strconv"
	"math"
	"unicode"
)

// func pow(number float64, power float64) float64 {
// 	if power == 0 {
// 		return 1
// 	}

// 	if power == 1 {
// 		return number
// 	}

// 	return number * pow(number, power - 1)
// }

func FindPowerDigitSum(number float64, power float64) {
	sum := 0
	result := math.Pow(number, power)

	for _, char := range strconv.FormatFloat(result, 'f', 10, 64) {
		if unicode.IsDigit(char) {
			digit, err := strconv.ParseFloat(string(char), 64)

			if err != nil {
				fmt.Println("Could not parse string to int:", err)
			}

			sum += int(digit)
		} else {
			continue
		}
	}

	fmt.Println(sum)
}
