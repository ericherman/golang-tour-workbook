package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := point2d{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *point2d implements Abser

	// In the following line, v is a point2d (not *point2d)
	// and does NOT implement Abser.
	// a = v // this would not compile

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type point2d struct {
	x, y float64
}

func (v *point2d) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
