package main

import (
	"fmt"
)

// stuff inits to zero-like values
func main() {
	var b bool
	var s string
	var i int
	var f float32
	var c complex64

	fmt.Printf("%v %q %v %v %v\n", b, s, i, f, c)
}
