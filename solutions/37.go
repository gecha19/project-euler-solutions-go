package solutions

import "fmt"

func TruncatablePrimes() {
	//number := 3797
	primes := make([]bool, 1000000)
	sum := 0

	for i := 2; i < 1000000; i++ {
		primes[i] = true
	}

	for i := 2; i*i < 1000000; i++ {
		if !primes[i] {
			continue
		}

		for j := i * i; j < 1000000; j += i {
			primes[j] = false
		}
	}

	for number := 1; number < 1000000; number++ {
		if primes[number] {
			numOfDigits := 0

			for i := number; i > 0; i /= 10 {
				numOfDigits++
			}

			for j := 1; j < numOfDigits; j++ {
				power := 1

				for k := 0; k < j; k++ {
					power *= 10
				}

				// From right to left
				rightToLeft := number % power

				// From left to right
				leftToRight := number / power

				if !primes[rightToLeft] || !primes[leftToRight] {
					break
				}

				if j == numOfDigits-1 {
					sum += number
				}

			}
		}
	}

	fmt.Println(sum)
}
