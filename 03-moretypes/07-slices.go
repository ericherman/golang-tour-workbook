package main

import "fmt"

// A slice points to an array of values and also includes a length.
// []T is a slice with elements of type T.
// len(primes) returns the length of slice primes.

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes ==", primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("primes[%d] == %d\n", i, primes[i])
	}
}
