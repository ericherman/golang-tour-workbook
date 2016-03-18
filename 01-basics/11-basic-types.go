package main

import (
	"fmt"
	"math/cmplx"
)

// https://golang.org/ref/spec#Types
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32 (represents a Unicode code point)
// float32 float64
// complex64 complex128
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

// int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
// and 64 bits wide on 64-bit systems
//
// note this is different than C for arch i686 and arch x86_64:
//
// arch i686
//       sizeof(unsigned int):    4 bytes                4294967295 max
//      sizeof(unsigned long):    4 bytes                4294967295 max
// sizeof(unsigned long long):    8 bytes      18446744073709551615 max
//
// arch x86_64
//       sizeof(unsigned int):    4 bytes                4294967295 max
//      sizeof(unsigned long):    8 bytes      18446744073709551615 max
// sizeof(unsigned long long):    8 bytes      18446744073709551615 max
