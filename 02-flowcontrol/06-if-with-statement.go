package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if n == 0 {
		return 1
	}

	// a statement is allowed
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// v not available outside of if-else scope

	return lim
}

func main() {
	fmt.Println(
		pow(2, 0, 10),
		pow(2, 1, 10),
		pow(2, 2, 10),
		pow(2, 3, 10),
	)
}
