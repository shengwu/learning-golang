package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	walk(t, ch, quit)
	close(ch)
}

func walk(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// no fallthrough by default
	case <-quit:
		return
	}
	ch <- t.Value
	walk(t.Right, ch, quit)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	quit := make(chan int)
	c1, c2 := make(chan int), make(chan int)

	// avoid goroutine leak
	defer close(quit)

	go Walk(t1, c1, quit)
	go Walk(t2, c2, quit)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {

			// if both are closed, return true
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Printf("Same tree: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Different trees: %v\n", Same(tree.New(1), tree.New(2)))

	var ch chan int
	ch = make(chan int)
	go Walk(tree.New(1), ch, nil)
	for i := range ch {
		fmt.Println(i)
	}
}
