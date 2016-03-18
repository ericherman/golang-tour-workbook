package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = " "
	a[2] = "World"

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}
	fmt.Println()

	fmt.Printf("a: %v is of type %T\n", a, a)
}
