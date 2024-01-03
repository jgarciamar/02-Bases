package calculator_test

import (
	"exercise-04/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinimumScore(t *testing.T) {

	testArray := []float64{4.9, 4.3, 3.2, 3.0, 3.1, 2.5, 1.8}
	expectedMinimum := 1.8
	expectedMaximum := 4.9
	expectedAverage := 3.2571428571429

	minimumFunction, _ := calculator.GetOperation("minimum")
	maximumFunction, _ := calculator.GetOperation("maximum")
	averageFunction, _ := calculator.GetOperation("average")

	minimum := minimumFunction(testArray...)
	maximum := maximumFunction(testArray...)
	average := averageFunction(testArray...)

	assert.Equal(t, expectedMinimum, minimum, "Expected minimum is not equal to the calculated minimum")

	assert.Equal(t, expectedMaximum, maximum, "Expected maximum is not equal to the calculated maximum")

	assert.InDelta(t, expectedAverage, average, 0.00001, "Expected average is not equal to the calculated average")

}
