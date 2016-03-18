package main

import "fmt"

func main() {
	foo := 0

	// other than paren syntax, basically C
	for {
		foo += 1
		if foo%1024 == 1 {
			fmt.Printf("\r%v", foo)
		}
	}
}
