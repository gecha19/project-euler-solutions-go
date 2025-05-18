package solutions

func FindSummationOfPrimes() uint {
	var sum uint = 2

	sieve := [2000000]bool{}

	for i := range 2000000 {
		sieve[i] = true
	}

	sieve[0], sieve[1] = false, false

	for i := 2; i < 2000000; i += 2 {
		sieve[i] = false
	}

	for i := 3; i < 2000000; i += 2 {
		for j := i * i; j < 2000000; j += 2 * i {
			sieve[j] = false
		}
	}

	for i := 3; i < 2000000; i += 2 {
		if sieve[i] {
			sum += uint(i)
		}
	}

	return sum
}
