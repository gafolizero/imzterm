package grayscale

import (
	"image"
	"image/color"

	"imz/imgToTensor"
)

func grayScaleAvg(grid [][]color.Color) (grayImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])

	grayImg = make([][]color.Color, xlen)

	for idx := range len(grayImg) {
		grayImg[idx] = make([]color.Color, ylen)
	}

	for x := range xlen {
		for y := range ylen {
			pix := grid[x][y].(color.NRGBA)
			gray := uint8(float64(pix.R)/3 + float64(pix.G)/3 + float64(pix.B)/3)
			grayImg[x][y] = color.NRGBA{
				R: gray,
				G: gray,
				B: gray,
				A: pix.A,
			}
		}
	}

	return grayImg
}

func GrayImg(img image.Image) [][]color.Color {
	grid := imgToTensor.ImgToTensorCol(img)
	grayGrid := grayScaleAvg(grid)
	return grayGrid
}
