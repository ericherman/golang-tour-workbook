package main

import "fmt"

func main() {
	i := 42
	f := 1.1
	c := 1 + 0.5i
	fmt.Printf("i is of type %T: %v\n", i, i)
	fmt.Printf("f is of type %T: %v\n", f, f)
	fmt.Printf("c is of type %T: %v\n", c, c)
}
