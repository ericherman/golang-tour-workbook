package main

import "fmt"

func main() {
	foo := 0

	// other than paren syntax, basically C
	for i := 0; i < 10; i++ {
		foo += i
	}

	fmt.Println(foo)
}
