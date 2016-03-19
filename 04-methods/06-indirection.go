package main

import "fmt"

type point2d struct {
	x, y float64
}

func (v *point2d) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *point2d, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := point2d{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &point2d{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
