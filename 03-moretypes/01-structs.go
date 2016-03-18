package main

import "fmt"

type point2d struct {
	x int
	y int
}

func main() {
	point := point2d{2, 3}
	point.x = 5
	fmt.Println(point)

	p := &point
	p.y = 7

	fmt.Println(point)
}
