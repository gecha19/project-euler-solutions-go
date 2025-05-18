package solutions

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
)

func FindNamesScores() {
	file, err := os.Open("names.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	csv := csv.NewReader(file)
	names, err := csv.Read()

	if err != nil {
		fmt.Println(err)
	}
	
	slices.Sort(names)
	totalScore := 0
	
	for pos, name := range names {
		score := 0

		for _, char := range name {
			score += int(char - 'A' + 1)
		}

		totalScore += score*(pos + 1)
	}

	fmt.Println(totalScore)
}
