package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// fmt.Println(scores)

	// totalScores := scores[0] + scores[1] + scores[2] + scores[3] + scores[4] + scores[5] + scores[6] + scores[7]
	// fmt.Println(totalScores / len(scores))

	var total int

	for _, score := range scores {
		total = total + score

	}

	length := len(scores)
	average := float64(total) / float64(length)
	fmt.Println(average)

	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
			fmt.Println(goodScores)
		}

	}

}
