package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("OS: ")

	// breaks automatically, unless it ends with a "fallthrough" statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s", os)
	}
}
