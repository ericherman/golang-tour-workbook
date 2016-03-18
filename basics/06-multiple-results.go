package main

import "fmt"

// return two strings
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("cold", "ice")
	fmt.Println(a, b)
}
