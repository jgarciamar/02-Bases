package main

import (
	"errors"
	"fmt"
)

var SalaryError error = errors.New("Error: salary is less than 10000")
var NegativeSalaryError error = errors.New("Negative salary? That must be an error, damn!")

func validateSalary(salary int) error {
	if salary < 0 {
		return NegativeSalaryError
	} else if salary <= 10000 {
		return SalaryError
	}

	return nil
}

func main() {
	salary := -1
	err := validateSalary(salary)
	if errors.Is(err, SalaryError) {
		fmt.Println(err.Error())
		return
	}
	if errors.Is(err, NegativeSalaryError) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Salary is valid")

}
