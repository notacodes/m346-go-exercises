package main

import (
	"fmt"
)

func computeGrade(gotPoints float64, maxPoints float64) (float64, error) {
	if maxPoints < 0.0 || gotPoints < 0.0 {
		return 0.0, fmt.Errorf("maxPoints and gotPoints must be greater than 0")
	}
	if gotPoints > maxPoints {
		return 0.0, fmt.Errorf("gotPoints cannot be greater than maxPoints")
	}

	grade := (gotPoints/maxPoints)*5 + 1
	return grade, nil
}

func main() {
	grade, err := computeGrade(1.0, 5.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(grade)
}
