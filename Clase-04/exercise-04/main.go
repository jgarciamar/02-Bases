package main

import (
	"exercise-04/calculator"
	"fmt"
)

func main() {

	op1 := "minimum"
	op2 := "maximum"
	op3 := "average"

	scores := []float64{4, 2, 3, 4, 4.5, 4, 3, 2.4, 1}

	minimum, err := calculator.GetOperation(op1)
	if err != "" {
		fmt.Println("Error1")
		return
	}
	maximum, _ := calculator.GetOperation(op2)
	if err != "" {
		fmt.Println("Error2")
		return
	}

	average, err := calculator.GetOperation(op3)
	if err != "" {
		fmt.Println("Error3")
		return
	}

	fmt.Printf("The minimum is %.4f\n", minimum(scores...))
	fmt.Printf("The maximum is %.4f\n", maximum(scores...))
	fmt.Printf("The average is %.4f\n", average(scores...))

}
