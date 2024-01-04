package main

import (
	"errors"
	"fmt"
)

type SalaryError struct{}

func (e *SalaryError) Error() string {
	return "Error: Salary is lesss than 10000"
}

func validateSalary(salary int) error {
	if salary <= 10000 {
		return &SalaryError{}
	}
	return nil
}

func main() {
	salary := 10
	err := validateSalary(salary)
	if errors.Is(err, &SalaryError{}) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Salary is valid")

}
