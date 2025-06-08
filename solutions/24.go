package solutions

import (
	"fmt"
	"slices"
)

func LexicographicPermutations() {
	candidates := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	factorials := make([]int, 10)

	for i, value := range candidates {
		factorials[i] = int(factorialOf(float64(value)))
	}

	remaining := 999999
	permutation := make([]int, 0)

	for j := 9; j >= 0; j-- {
		if len(candidates) == 1 {
			permutation = append(permutation, candidates[0])
			break
		}

		permutations := factorials[j]
		// permutations := int(factorialOf(float64(j)))
		index := int(remaining / permutations)
		permutation = append(permutation, candidates[index])
		remaining %= permutations
		candidates = slices.Delete(candidates, index, index+1)
	}

	fmt.Println(permutation)
}
