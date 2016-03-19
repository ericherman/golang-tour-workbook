package main

import (
	"fmt"
	"math"
)

type point2d struct {
	x, y float64
}

func (v *point2d) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *point2d) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &point2d{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
