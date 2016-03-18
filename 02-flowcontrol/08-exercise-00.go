package main

import (
	"fmt"
	"math"
)

func GoodEnough(x, z float64) bool {
	er := math.Abs((z * z) - x)
	return (er < 0.0000000000001)
}

// newton's method
func Sqrt(x float64) float64 {
	z := x / 2

	var i int

	for i = 0; !GoodEnough(x, z) && i < 10; i++ {
		z = (z + x/z) / 2
	}

	if i < 10 {
		fmt.Printf("%v resolved in %v iterations\n", x, i)
	}

	return z
}

func main() {
	for i := 1; i < 100; i++ {
		mi := Sqrt(float64(i))
		ti := math.Sqrt(float64(i))
		diff := 100 * math.Abs(mi-ti)
		fmt.Printf("Sqrt(%v): %v, math.Sqrt(%v): %v (%v)\n",
			i, mi, i, ti, diff)
	}
}
