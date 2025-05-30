package solutions

import (
	"fmt"
	"math/big"
	"slices"
)

func FindDistinctPowers(start int64, end int64) {
	seen := make([]string, 0)

	for a := start; a <= end; a++ {
		for b := start; b <= end; b++ {
			z := big.NewInt(0).Exp(big.NewInt(a), big.NewInt(b), nil).String()

			if !slices.Contains(seen, z) {
				seen = append(seen, z)
			}
		}
	}

	fmt.Println(len(seen))
}
