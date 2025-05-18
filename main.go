package main

import (
	"fmt"
	"time"
	"github.com/gecha19/project-euler-solutions-go/solutions"
)

func main() {
	fmt.Println("Use this to run any of the solutions!")
	start := time.Now()
	solutions.FindNonAbundantsSum(28123)
	timeElapsed := time.Since(start)
	fmt.Println(timeElapsed)
}
