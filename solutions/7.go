package solutions

import "fmt"

func NthPrime(N uint) {
	var count uint = 0
	var number uint = 2

	for {
		if isPrime(number) && count < N {
			count += 1
		}

		if count >= N {
			fmt.Printf("%v\n", number)
			break
		}

		number += 1
	}
}
