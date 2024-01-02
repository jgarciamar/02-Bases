package main

import "fmt"

func MinToHours(minutes int) int {

	return minutes / 60
}

func calculateSalary(minutes int, category rune) (salary int) {

	salaryMap := map[rune]int{'C': 1000, 'B': 1500, 'A': 3000}
	hoursWorked := MinToHours(minutes)

	return salaryMap[category] * hoursWorked

}

func main() {

	s1 := 400
	s2 := 250
	cat1 := 'C'
	cat2 := 'A'

	fmt.Printf("Salary1: %d", calculateSalary(s1, cat1))
	fmt.Printf("Salary2: %d", calculateSalary(s2, cat2))

}
