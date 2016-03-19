package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = byte('A')
	}
	return i, nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 8, 8)
	r.Read(b)
	fmt.Printf("bytes: %v\n", b)
}
