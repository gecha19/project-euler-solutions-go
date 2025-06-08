package solutions

/*
a + b + c = 1000 => c = 1000 - b - a
a < b < c => a < b < 1000 - b - a => a < b < 500 - a/2
a < b < c and 1000/3 = 333.33 => a < 333, since a is a natural number
*/

func SpecialPythagoreanTriplet() int {
	for a := range 333 {
		for b := a; b < 500-a/2; b++ {
			c := 1000 - b - a
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}

	return 0
}
