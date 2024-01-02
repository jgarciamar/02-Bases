package main

import "fmt"

func calculateTax(salary int) float32 {
	var tax float32
	if salary > 50000 {
		tax += 17
	}
	if salary > 150000 {
		tax += 10
	}
	return float32(salary) * tax
}

func main() {
	s1 := 56000
	s2 := 167000

	fmt.Printf("Tax1: %.6f, Tax2: %.6f", calculateTax(s1), calculateTax(s2))
}
