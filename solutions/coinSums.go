package solutions

import "fmt"

func FindCoinSums() {
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
	const target = 200
	ways := make(map[int]int, target+1)
	// Only one way to not make any; do not choose any
	ways[0] = 1

	for _, coin := range coins {
		for s := coin; s <= target; s++ {
			ways[s] += ways[s-coin]
		}
	}

	fmt.Println(ways[200])
}
