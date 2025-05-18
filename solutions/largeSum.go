package solutions

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindLargeSum() {
	data, err := os.ReadFile("./large_numbers.txt")

	if err != nil {
		log.Println("Failed to read the file:", err)
	}

	largeNumbers := strings.Split(string(data), "\n")
	var sum int64

	for _, number := range largeNumbers {
		/* 10 digits + log10(100) = 10 + 2 = 12 */
		value, err := strconv.ParseInt(number[:12], 10, 64)

		if err != nil {
			fmt.Println("Failed to convert string to int64:", err)
		}

		sum += value
	}

	fmt.Println(strconv.FormatInt(sum, 10)[:10])
}
