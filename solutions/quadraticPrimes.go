package solutions

import (
	"fmt"
	"math"
)

func FindQuadraticPrimes() {
	// For n = 0, b should be a prime => we can skip all non-prime numbers
	bValues := make([]int, 0)

	for i := 2; i <= 1000; i++ {
		if isPrime(uint(i)) {
			bValues = append(bValues, i, -i)
		}
	}

	primesCount := 0
	coefficient1 := 0
	coefficient2 := 0

	for a := -999; a < 1000; a++ {
		for _, b := range bValues {
			n := 0
			count := 0

			for {
				integer := uint(math.Pow(float64(n), 2.0) + float64(a*n) + float64(b))

				if isPrime(integer) {
					count++
					n += 1
				} else {
					break
				}
			}

			if count > primesCount {
				primesCount = count
				coefficient1 = a
				coefficient2 = b
			}
		}
	}

	product := coefficient1 * coefficient2
	fmt.Println(product)
}
