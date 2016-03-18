package main

import "fmt"

// https://blog.golang.org/gos-declaration-syntax

func add(x int, y int) int {
	return x + y
}

// consecutive parameters of the same type,
// may omit the type from all but the last
func subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(subtract(add(42, 23), 13))
}
