package solutions

import "fmt"

// func findTriangleNumber(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	if n == 2 {
// 		return 3
// 	}

// 	return n + findTriangleNumber(n - 1)
// }

func FindHighlyDivisibleTriangularNumber() {
	currTriangleNumber := 1
	position := 1

	for {
		count := 2
		position ++
		currTriangleNumber += position

		for i := 2; i*i <= currTriangleNumber; i ++ {
			if currTriangleNumber%i == 0 {
				count += 2
			}
		}

		if count > 500 {
			fmt.Println(currTriangleNumber)
			break
		}
	}
}
