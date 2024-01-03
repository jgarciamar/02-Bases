package FoodCalculator_test

import (
	"exercise-05/FoodCalculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalCalculator(t *testing.T) {

	testCases := []struct {
		animal   string
		amount   int
		expected float64
	}{
		{"dog", 10, 10 * 250},
		{"cat", 2, 2 * 5},
		{"spider", 9, 9 * 150},
		{"hamster", 8, 8 * 250},
	}

	for _, testCase := range testCases {
		animalFunction, _ := FoodCalculator.Animal(testCase.animal)
		assert.Equal(t, testCase.expected, animalFunction(testCase.amount), "The amount calculated of Food is not the one expected")
	}

}
