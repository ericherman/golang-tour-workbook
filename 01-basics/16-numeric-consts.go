package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func times10Int(x int) int         { return x * 10 }
func div10Float(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Printf("Small is of type %T: %v\n", Small, Small)
	fmt.Println(times10Int(Small))
	fmt.Println(div10Float(Small))
	fmt.Println(div10Float(Big))

	// constant 1267650600228229401496703205376 overflows int
	// fmt.Printf("Big is of type %T: %v\n", Big, Big)
}
