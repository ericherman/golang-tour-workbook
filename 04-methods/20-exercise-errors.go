package main

import (
	"fmt"
	"math"
)

func GoodEnough(x, z float64) bool {
	er := math.Abs((z * z) - x)
	return (er < 0.0000000000001)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Complex numbers not supported. (%v)", float64(e))
}

// newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x)
	}
	z := x / 2

	var i int

	for i = 0; !GoodEnough(x, z) && i < 10; i++ {
		z = (z + x/z) / 2
	}

	if i < 10 {
		fmt.Printf("%v resolved in %v iterations\n", x, i)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
