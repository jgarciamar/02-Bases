package main

import (
	"exercise-05/internal"
	"fmt"
)

func main() {

	hoursWorked1 := 160
	moneyPerHour1 := 10000

	hoursWorked2 := 70
	moneyPerHour2 := 190000

	salary1, err := internal.CalculateSalary(hoursWorked1, moneyPerHour1)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(salary1)

	salary2, err := internal.CalculateSalary(hoursWorked2, moneyPerHour2)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(salary2)
}
