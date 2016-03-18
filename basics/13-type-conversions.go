package main

import (
	"fmt"
	"math"
)

func main() {
	// no implicit conversions
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)

	var x, y int = 3, i / 10
	var ftmp float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ftmp)
	fmt.Println(x, y, z)
}
