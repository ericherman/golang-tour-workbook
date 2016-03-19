package main

import "fmt"

type I interface {
	M()
}

type T1 struct {
	S string
}

func (t *T1) M() {
	// guard against "panic: runtime error:
	// invalid memory address or nil pointer dereference"
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

type T2 struct {
	S string
}

func (t *T2) M() {
	fmt.Println(t.S)
}

func main() {
	var i I

	describe(i)

	var t *T1
	i = t
	describe(i)
	i.M()

	i = &T2{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
