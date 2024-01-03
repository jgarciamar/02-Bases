package calculator_test

import (
	"exercise-02/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalStudentScoreAvg(t *testing.T) {

	t.Run("Should fail: An element is under 0 or over 5 ", func(t *testing.T) {

		scores := []float64{2, 4, 5, 6, 7, 8, 1, 2, 9, 2, 3}
		_, err := calculator.CalcStudentScoreAvg(scores)
		assert.NotEqual(t, err, "", "Err shoudlnt be empty")

	})

	t.Run("Should suceeed: Average should be calculated correctly", func(t *testing.T) {

		scores := []float64{4, 3, 3, 2, 4.5}
		average, _ := calculator.CalcStudentScoreAvg(scores)
		expectedAverage := float64(4+3+3+2+4.5) / float64(len(scores))

		assert.Equal(t, expectedAverage, average, "Expected should be equal to actual")

	})

}
