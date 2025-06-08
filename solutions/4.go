package solutions

import (
	"fmt"
	"strconv"
)

func isPalindrome(number int) bool {
	var original string = strconv.FormatInt(int64(number), 10)
	var length int = len(original)

	for i := range original {
		if original[i] != original[length-i-1] {
			return false
		}
	}

	return true
}

func FindLargestPalindromeProduct() {
	var i int = 999
	var largestPalindrome int = 0
	var number1 int
	var number2 int

	for i > 99 {
		var j int = i

		for j > 99 {
			var product int = i * j

			if product <= largestPalindrome {
				break
			}

			if isPalindrome(product) && largestPalindrome < product {
				largestPalindrome = product
				number1 = i
				number2 = j
			}

			j -= 1
		}

		i -= 1
	}

	fmt.Printf("%v = %v x %v\n", largestPalindrome, number1, number2)
}
