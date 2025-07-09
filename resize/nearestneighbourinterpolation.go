package resize

import (
	"image"
	"image/color"
	"math"

	"imz/imgToTensor"
)

func resizeNNI(grid [][]color.Color, scale float64) (resizedGrid [][]color.Color) {
	newX := int(float64(len(grid)) * scale)
	newY := int(float64(len(grid[0])) * scale)

	resizedGrid = make([][]color.Color, newX)

	for idx := range len(resizedGrid) {
		resizedGrid[idx] = make([]color.Color, newY)
	}

	for x := range newX {
		for y := range newY {
			xp := int(math.Floor(float64(x) / scale))
			yp := int(math.Floor(float64(y) / scale))
			resizedGrid[x][y] = grid[xp][yp]
		}
	}

	return resizedGrid
}

func NNI(img image.Image, scale float64) [][]color.Color {
	grid := imgToTensor.ImgToTensorRow(img)
	resizedGrid := resizeNNI(grid, scale)
	return resizedGrid
}
