package main

import (
	"fmt"
	"image"
	"image/color"
	// "golang.org/x/tour/pic"
)

type Image struct {
	h uint
	w uint
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(i.h), int(i.w))
}

func (i Image) At(x, y int) color.Color {
	v := uint8((float64(x*y) / float64(i.h*i.w)) * 255.0)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	var x, y uint = 60, 20
	m := Image{x, y}
	if false {
		// pic.ShowImage(m)
	} else {
		var i, j uint = 0, 0
		for ; i < y; i++ {
			for j = 0; j < x; j++ {
				fmt.Println(m.At(int(j), int(i)))
			}
		}
	}
}
