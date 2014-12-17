package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for i, row := range p {
		for j := range row {
			p[j][i] = byte((i + j) / 2)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
