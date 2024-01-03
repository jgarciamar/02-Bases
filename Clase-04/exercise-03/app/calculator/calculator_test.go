package calculator_test

import (
	"exercise-03/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {

	testCases := []struct {
		minutes  float64
		category rune
		expected float64
	}{
		{120, 'A', 6000},
		{90, 'B', 2250},
		{180, 'C', 3000},
	}

	for _, testCase := range testCases {
		assert.Equal(
			t,
			testCase.expected,
			calculator.CalculateSalary(testCase.minutes, testCase.category),
			"Expected value should be equal to calculate salary",
		)
	}
}
