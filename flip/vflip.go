package flip

import (
	"image"
	"image/color"
)

func imgToTensorV(img image.Image) [][]color.Color {
	var pixels [][]color.Color

	size := img.Bounds().Size()

	for idx := range size.X {
		var col []color.Color
		for jdx := range size.Y {
			col = append(col, img.At(idx, jdx))
		}

		pixels = append(pixels, col)
	}

	return pixels
}

func VFlip(img image.Image) [][]color.Color {
	grid := imgToTensorV(img)
	flip(grid)
	//saveImz.SaveRectXY(filePath, grid)
	return grid
}
