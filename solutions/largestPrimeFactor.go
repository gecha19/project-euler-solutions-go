package solutions

import "fmt"

func isPrime(number uint) bool {
	var count uint = 3

	if number == 2 {
		return true
	}

	if number % 2 == 0 || number < 2 {
		return false
	}

	for count*count <= number {
		if number % count == 0 {
			return false
		}
		count += 2
	}

	return true
}

func LargestPrimeFactor(number uint) {
	var largestPrimeFactor uint = 1
	var count uint = 2

	for count*count <= number {
		if isPrime(count) && number % count == 0 {
			largestPrimeFactor = count
		}
		count += 1
	}

	fmt.Printf("%v\n", largestPrimeFactor)
}
