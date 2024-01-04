package internal

import "errors"

func CalculateSalary(hours int, moneyPerHour int) (salary int, err error) {

	salary = hours * moneyPerHour
	tax := 0.0

	if salary > 150000 {
		tax = 0.10
		salary *= int(tax * float64(salary))
	}

	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	return salary, nil

}
