package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// ignore the index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// ignore the value
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}
