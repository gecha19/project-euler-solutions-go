package solutions

import "fmt"

func SumOfMultiplesOf3And5() {
	count := 6
	sum := 8

	for count < 1000 {
		if count % 3 == 0 || count % 5 == 0 {
			sum += count
		}
		count += 1
	}

	fmt.Printf("%v\n", sum)
}
