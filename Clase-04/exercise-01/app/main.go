package main

import (
	"exercise-01/calculator"
	"fmt"
)

func main() {
	s1 := 56000.0
	s2 := 167000.0

	fmt.Printf("Tax1: %.6f, Tax2: %.6f", calculator.CalculateTax(s1), calculator.CalculateTax(s2))

}
