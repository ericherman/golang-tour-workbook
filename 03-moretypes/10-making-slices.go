package main

import "fmt"

func main() {
	a := make([]int, 5) // len(a) == 5
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:2]
	for i:=0; i < len(c); i++ {
		c[i] = (i+1)*10
	}
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	fmt.Println("d[0] = 1")
	d[0] = 1

	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
