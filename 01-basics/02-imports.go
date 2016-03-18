package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64
	f = math.Sqrt(7)
	// https://golang.org/pkg/fmt/
	fmt.Printf("Float: %g Int: %d Native: %v.\n", f, int(f), f)
}
