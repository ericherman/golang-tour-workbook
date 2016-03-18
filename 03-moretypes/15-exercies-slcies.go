package main

import "fmt"

// should return a slice of length dy,
// each element of which is a slice of dx 8-bit unsigned integers
func Pic(dx, dy int) [][]uint8 {
	rv := make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		rv[i] = make([]uint8, dx, dx)
		for j := 0; j < dx; j++ {
			rv[i][j] = uint8(i * j)
		}
	}
	return rv
}

func main() {
	x := 20
	y := 10
	aa := Pic(x, y)
	fmt.Println(aa)
}
