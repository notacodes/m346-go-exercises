package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := b*b - 4*a*c
	if D > 0 {
		sqrtD := math.Sqrt(D)
		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		return []float64{x1, x2}
	} else if D == 0 {
		x := -b / (2 * a)
		return []float64{x}
	}
	return []float64{}
}

func main() {
	r1 := computeQuadraticFormula(3, 4, 1)
	fmt.Println(r1)
	r2 := computeQuadraticFormula(2, 4, 2)
	fmt.Println(r2)
	r3 := computeQuadraticFormula(3, 4, 2)
	fmt.Println(r3)
}
