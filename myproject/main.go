package main

import (
	"fmt"
	"myproject/cal"
)

func main() {
	a := 2
	b := 4

	results := cal.Calculations(a, b)

	fmt.Println("results:", results)
}
