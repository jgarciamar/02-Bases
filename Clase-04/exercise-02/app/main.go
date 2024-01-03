package main

import (
	"exercise-02/calculator"
	"fmt"
)

func main() {
	scores := []float64{4, 5, 3, 4, 3, 2, 2, 2, 3, 2, 3, 2, 4, 4, 4, 3}

	scoreAverage, _ := calculator.CalcStudentScoreAvg(scores)

	fmt.Printf("The score average is %.3f", scoreAverage)
}
