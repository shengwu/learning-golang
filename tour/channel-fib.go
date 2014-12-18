package main

import (
	"fmt"
)

func fibonacci(n int) chan int {
	c := make(chan int, 10)

	// run anonymous function in a goroutine
	go func() {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()

	return c
}

func main() {
	for v := range fibonacci(10) {
		fmt.Println(v)
	}
}
