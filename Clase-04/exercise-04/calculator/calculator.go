package calculator

type MathOperation func(...float64) float64

func GetMinimumScore(studentScores ...float64) (minimumScore float64) {

	minimumScore = studentScores[0]
	for _, v := range studentScores {
		if v <= minimumScore {
			minimumScore = v
		}
	}
	return
}

func GetMaximumScore(studentScores ...float64) (maximumScore float64) {
	maximumScore = studentScores[0]
	for _, v := range studentScores {
		if v >= maximumScore {
			maximumScore = v
		}
	}
	return
}

func GetAverageScore(studentScores ...float64) float64 {

	var sum float64

	for _, v := range studentScores {
		sum += v
	}

	return sum / float64(len(studentScores))

}

func GetOperation(operation string) (MathOperation, string) {

	switch operation {

	case "minimum":
		return GetMinimumScore, ""
	case "maximum":
		return GetMaximumScore, ""
	case "average":
		return GetAverageScore, ""
	default:
		return nil, "Operation not defined"
	}

}
