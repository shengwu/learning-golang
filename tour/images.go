package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
	//"fmt"
	//"math/rand"
)

type Image struct{}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x ^ y),
		uint8(x + y)}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
	//fmt.Printf("%#v", m)
}
