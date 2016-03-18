package main

import "fmt"

func main() {
	foo := 1

	// other than paren syntax, basically C
	for foo < 1000 {
		foo += foo
	}

	fmt.Println(foo)
}
