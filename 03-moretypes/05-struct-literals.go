package main

import "fmt"

type point2d struct {
	x int
	y int
}

var (
	v1 = point2d{1, 2}  // has type point2d
	v2 = point2d{y: 2}  // x:0 is implicit
	v3 = point2d{}      // x:0 and y:0
	p  = &point2d{1, 2} // has type *point2d
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
