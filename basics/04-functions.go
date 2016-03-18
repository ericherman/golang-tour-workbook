package main

import "fmt"

// https://blog.golang.org/gos-declaration-syntax

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 23))
}
