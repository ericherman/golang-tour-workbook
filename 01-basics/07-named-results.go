package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// "naked" return returns x, y named above
	return
}

func main() {
	fmt.Println(split(123))
}
