package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x, y int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}

func (i Image) At(x, y int) color.Color {

	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
