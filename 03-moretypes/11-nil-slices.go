package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("val s: %v, len(s): %v, cap(s): %v\n", s, len(s), cap(s))
	fmt.Println("s == nil:", s == nil)
}
