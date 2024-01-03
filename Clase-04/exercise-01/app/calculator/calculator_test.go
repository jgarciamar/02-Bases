package calculator_test

import (
	"exercise-01/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTex(t *testing.T) {

	t.Run("Case1: Salary under 50000", func(t *testing.T) {

		salary := 45000.0
		expectedResult := salary
		result := calculator.CalculateTax(salary)
		assert.Equal(t, expectedResult, result, "The salary is not equal to the expected Salary")

	})

	t.Run("Case 2: Salary over 50000", func(t *testing.T) {

		salary := 55000.0
		expectedResult := salary * 0.17
		result := calculator.CalculateTax(salary)

		assert.Equal(t, expectedResult, result, "The salary is not equal to the expected Salary")

	})

	t.Run("Case 3: Salary over 150000", func(t *testing.T) {

		salary := 160000.0

		expectedResult := salary * 0.27
		result := calculator.CalculateTax(salary)

		assert.Equal(t, expectedResult, result, "The salary is not equal to the expected Salary")
	})

}
