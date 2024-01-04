package main

import "fmt"

type SalaryError struct {
}

func (e *SalaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

func main() {
	salary := 2

	if salary < 150.000 {
		err := &SalaryError{}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")

}
