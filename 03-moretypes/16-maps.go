package main

import "fmt"

type point2d struct {
	Lat, Long float64
}

var m map[string]point2d

func main() {
	m = make(map[string]point2d)
	m["Bell Labs"] = point2d{40.68433, -74.39967}
	m["north pole"] = point2d{90, 0}

	for key := range m {
		fmt.Printf("%v => %v\n", key, m[key])
	}
}
