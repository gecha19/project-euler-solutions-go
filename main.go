package main

import (
	"fmt"
	"github.com/gecha19/project-euler-solutions-go/solutions"
	"time"
)

func main() {
	fmt.Println("Use this to run any of the solutions!")
	start := time.Now()
	solutions.PandigitalMultiples()
	timeElapsed := time.Since(start)
	fmt.Println(timeElapsed)
}
