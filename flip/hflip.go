package flip

import (
	"image"
	"image/color"

	"imz/imgToTensor"
)

func HFlip(img image.Image) [][]color.Color {
	grid := imgToTensor.ImgToTensorRow(img)
	flip(grid)
	return grid
}
