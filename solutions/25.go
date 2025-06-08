package solutions

import (
	"fmt"
	"math/big"
)

func Find1000DigitFibonacciNumber() {
	previousTerm1 := big.NewInt(1)
	previousTerm2 := big.NewInt(1)
	currentTerm := big.NewInt(0)
	index := 2

	for  {
		currentTerm.Add(previousTerm1, previousTerm2)
		length := len(currentTerm.String())
		index ++

		if length >= 1000 {
			break
		}

		previousTerm2.Set(previousTerm1)
		previousTerm1.Set(currentTerm)
	}
	
	fmt.Println(index)
}
