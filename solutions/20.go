package solutions

import (
	"fmt"
	"math/big"
)

func FactorialDigitSum(number int64) {
	factorial := big.NewInt(1)

	for i := int64(2); i <= number; i += 1 {
		factorial.Mul(factorial, big.NewInt(i))
	}

	sum := 0

	for _, char := range factorial.String() {
		sum += int(char - '0')
	}

	fmt.Println(sum)
}
