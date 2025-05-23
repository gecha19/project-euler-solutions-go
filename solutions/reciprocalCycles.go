package solutions

import (
	"fmt"
	"slices"
)

func LongestRecurringCycle() {
	longestRecurringCycle := 0
	maxDenominator := 1

	for d := 2; d < 1000; d++ {
		remainder := 1
		seen := make([]int, 0)

		for {
			currRemainder := (remainder*10)%d

				if slices.Contains(seen, currRemainder) {
					break
				} else {
					seen = append(seen, currRemainder)
				}
			
			remainder = currRemainder
		}

		currentRecurringCycle := len(seen)

		if currentRecurringCycle > longestRecurringCycle {
			longestRecurringCycle = currentRecurringCycle
			maxDenominator = d
		}
	}

	fmt.Println(maxDenominator)
}
