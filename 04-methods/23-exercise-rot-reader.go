package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case ((b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M')):
		return b + 13
	case ((b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z')):
		return b - 13
	default:
		return b
	}
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	i := 0
	for ; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return i, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}
