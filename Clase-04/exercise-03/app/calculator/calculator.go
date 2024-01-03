package calculator

func CalculateSalary(minutes float64, category rune) (salary float64) {

	salaryMap := map[rune]float64{'C': 1000, 'B': 1500, 'A': 3000}
	hoursWorked := float64(minutes) / 60.0

	return salaryMap[category] * hoursWorked

}
