package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(3, 4))
	fmt.Println(computeHypotenuse(5, 12))
	fmt.Println(computeHypotenuse(1.5, 2))

	ss1 := ShortSides{a: 3, b: 4}
	fmt.Println(ss1.Hypotenuse())

	ss2 := ShortSides{a: 5, b: 12}
	fmt.Println(ss2.Hypotenuse())

	ss3 := ShortSides{a: 1.5, b: 2}
	fmt.Println(ss3.Hypotenuse())
}
