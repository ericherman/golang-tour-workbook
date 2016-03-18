package main

import (
	"fmt"
	"math"
)

type point2d struct {
	x, y float64
}

func (v point2d) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := point2d{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}
