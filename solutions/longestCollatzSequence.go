package solutions

import "fmt"

// func collatzSequence(number int, sequence []int) []int {
// 	sequence = append(sequence, number)

// 	if number == 1 {
// 		return sequence
// 	}

// 	if number%2 == 0 {
// 		return collatzSequence(number/2, sequence)
// 	} else {
// 		return collatzSequence(3*number+1, sequence)
// 	}
// }

// func FindLongestCollatzSequence() {
// 	maxLength := 0
// 	startingPoint := 1

// 	for i := 1; i < 1000000; i++ {
// 		sequence := collatzSequence(i, []int{})

// 		currLength := len(sequence)

// 		if currLength > maxLength {
// 			maxLength = currLength
// 			startingPoint = i
// 		}
// 	}

// 	fmt.Println(startingPoint)
// }

func FindLongestCollatzSequence() {
	maxLength := 0
	seen := make([]int, 1000000)
	startingPoint := 1
	seen[0], seen[1] = 1, 1

	for i := 2; i < 1000000; i++ {
		length := 0
		n := i

		for n != 1 {
			if n%2 == 0 {
				n = n/2
				length ++
			} else {
				n = 3*n + 1
				length ++
			}

			if n < len(seen) && seen[n] != 0 {
				length += seen[n]
				n = 1
			}
		}
		
		if length > maxLength {
			maxLength = length
			startingPoint = i
		}

		seen[i] = length
	}

	fmt.Println(startingPoint)
}
