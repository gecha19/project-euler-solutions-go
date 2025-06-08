package solutions

import "fmt"

func EvenFibonacciTerms(number int) {
	previous_term := 1
	current_term := 2
	sum := 0

	for current_term < number {
		if current_term%2 == 0 {
			sum += current_term
		}
		tmp := current_term
		current_term += previous_term
		previous_term = tmp
	}

	fmt.Printf("%v\n", sum)
}
