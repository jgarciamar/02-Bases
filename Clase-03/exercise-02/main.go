package main

import "fmt"

func CalcScoreAverage(scores ...int) float32 {
	var sum int
	for _, score := range scores {
		sum += score
	}

	return float32(sum) / float32(len(scores))

}
func main() {
	scoreAverage := CalcScoreAverage(4, 5, 3, 4, 3, 2, 2, 2, 3, 2, 3, 2, 4, 4, 4, 3)
	fmt.Printf("The score average is %.3f", scoreAverage)
}
