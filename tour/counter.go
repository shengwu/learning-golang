package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var counts map[string]int
	counts = make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word] += 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
