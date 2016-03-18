package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var rnd int
	rnd = rand.Intn(10)
	fmt.Println("A call to rand.Intn(10): ", rnd)
	os.Exit(rnd)
}
