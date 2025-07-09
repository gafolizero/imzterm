package imgToTensor

import (
	"image"
	"image/color"
)

func ImgToTensorRow(img image.Image) [][]color.Color {
	var pixels [][]color.Color
	size := img.Bounds().Size()

	for idx := range size.Y {
		var row []color.Color
		for jdx := range size.X {
			row = append(row, img.At(jdx, idx))
		}
		pixels = append(pixels, row)
	}
	return pixels
}

func ImgToTensorCol(img image.Image) [][]color.Color {
	var pixels [][]color.Color
	size := img.Bounds().Size()

	for idx := range size.X {
		var row []color.Color
		for jdx := range size.Y {
			row = append(row, img.At(idx, jdx))
		}
		pixels = append(pixels, row)
	}
	return pixels
}
