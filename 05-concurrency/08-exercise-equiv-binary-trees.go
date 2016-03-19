package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, content chan int) {
	WalkInner(t, content)
	close(content)
}

func WalkInner(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkInner(t.Left, ch)
	ch <- t.Value
	WalkInner(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	go Walk(t1, chan1)
	chan2 := make(chan int)
	go Walk(t2, chan2)

	var j int
	var ok bool

	for i := range chan1 {
		j, ok = <-chan2
		if i != j {
			return false
		}
		if !ok {
			return false
		}
	}
	j, ok = <-chan2
	if ok {
		fmt.Println("chan 2 longer")
		return false
	}
	return true
}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Printf("Same(tree.New(1), tree.New(1)): %v\n",
		Same(tree.New(1), tree.New(1)))
	fmt.Printf("Same(tree.New(1), tree.New(2)): %v\n",
		Same(tree.New(1), tree.New(2)))
}
