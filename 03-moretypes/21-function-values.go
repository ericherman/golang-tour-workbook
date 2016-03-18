package main

import (
	"fmt"
	"math"
)

func compute_3_5(fn func(float64, float64) float64) float64 {
	return fn(3, 5)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute_3_5(hypot))
	fmt.Println(compute_3_5(math.Pow))
}
