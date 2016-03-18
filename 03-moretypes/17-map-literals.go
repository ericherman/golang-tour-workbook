package main

import "fmt"

type point2d struct {
	Lat, Long float64
}

var m map[string]point2d

func main() {
	m = map[string]point2d{
		"Bell Labs":  {40.68433, -74.39967},
		"north pole": {90, 0},
	}

	for key := range m {
		fmt.Printf("%v => %v\n", key, m[key])
	}
}
