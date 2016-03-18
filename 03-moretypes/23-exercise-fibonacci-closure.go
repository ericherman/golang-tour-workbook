package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	b1, b2 := -1, 1
	return func() int {
		v := b1 + b2
		b1 = b2
		b2 = v
		return v
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
