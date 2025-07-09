package rotate

import (
	"image"
	"image/color"

	"imz/imgToTensor"
)

func rotateMatrix(oldMat [][]color.Color) [][]color.Color {
	R := len(oldMat)
	C := len(oldMat[0])

	newMat := make([][]color.Color, C)

	for i := range C {
		newMat[i] = make([]color.Color, R)
	}

	for i := range C {
		R1 := len(oldMat)
		for j := range R {
			R1 = R1 - 1
			newMat[i][j] = oldMat[R1][i]
		}

	}

	return newMat
}

func RotateImg(img image.Image) [][]color.Color {
	grid := imgToTensor.ImgToTensorRow(img)
	rotatedGrid := rotateMatrix(grid)
	return rotatedGrid
}
