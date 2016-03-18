package main

import "fmt"
import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) {
		m[f]++
	}
	return m
}

func main() {
	s := "the chair is against the wall"
	fmt.Println(s)
	fmt.Println(WordCount(s))
}
