package solutions

import (
	"fmt"
)

func FindCircularPrimes(number int) {
	count := 0
	seen := make([]bool, number)
	primes := make([]bool, number)

	for i := 2; i < number; i++ {
		primes[i] = true
	}

	for i := 2; i*i < number; i++ {
		if !primes[i] {
			continue
		}

		for j := i * i; j < number; j += i {
			primes[j] = false
		}
	}

	for i := 1; i < number; i++ {
		if i < 10 && primes[i] {
			count++
		} else if !seen[i] && primes[i] {
			length := 0
			tmp := i
			powerOfTen := 1

			for tmp != 0 {
				length++
				powerOfTen *= 10
				tmp /= 10
			}

			rotation := i

			for j := 1; j < length; j++ {
				rotation *= 10
				firstDigit := rotation / powerOfTen
				rotation += firstDigit - powerOfTen*firstDigit
				seen[rotation] = true

				if !primes[rotation] {
					break
				}

				if rotation == i {
					count++
					break
				}

				if j == length-1 {
					count += length
				}
			}
		}
	}

	fmt.Println(count)
}
