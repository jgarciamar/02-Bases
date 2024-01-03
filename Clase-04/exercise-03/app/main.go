package main

import (
	"exercise-03/calculator"
	"fmt"
)

func main() {

	s1 := 400.0
	s2 := 250.0
	cat1 := 'C'
	cat2 := 'A'

	fmt.Printf("Salary1: %.3f", calculator.CalculateSalary(s1, cat1))
	fmt.Printf("Salary2: %.3f", calculator.CalculateSalary(s2, cat2))

}
