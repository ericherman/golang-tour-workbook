package main

import "fmt"

const hand = "World"

func main() {
	//Println adds a space between items
	fmt.Println("Hello", hand)

	const Truth = true
	fmt.Println("Golang is a bit lame?", Truth)
}
