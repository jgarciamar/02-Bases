package main

import "fmt"

func main() {
	var firstName string = "Mary"
	var lastName string = "Smith"
	var age int = 35
	married := false
	var salary float32 = 45857

	fmt.Printf("First Name: %s LastName: %s age: %d, Married: %t Salary: %.6f", firstName, lastName, age, married, salary)
}
