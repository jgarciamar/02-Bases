package calculator

func CalcStudentScoreAvg(scores []float64) (float64, string) {
	var sum float64

	for _, score := range scores {

		if score < 0 || score > 5 {
			return 0, "Theres one element under zero"
		}

		sum += score
	}

	return float64(sum) / float64(len(scores)), ""
}
