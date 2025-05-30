package solutions

import "fmt"

func FindNumberSpiralDiagonals(number int) {
	sum := 0
	start := 1
	spiral := 3

	for spiral <= number {
		for i := start; i < spiral*spiral; i += spiral - 1 {
			sum += i
		}

		start = spiral * spiral
		spiral += 2
	}

	sum += start
	fmt.Println(sum)
}
