package grayscale

import "image/color"

func grayScaleITU(grid [][]color.Color) (grayImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])

	grayImg = make([][]color.Color, xlen)

	for idx := range len(grayImg) {
		grayImg[idx] = make([]color.Color, ylen)
	}

	for x := range xlen {
		for y := range ylen {
			pix := grid[x][y].(color.NRGBA)
			gray := uint8(float64(pix.R)*0.2126 + float64(pix.G)*0.7152 + float64(pix.B)*0.0722)
			grayImg[x][y] = color.NRGBA{gray, gray, gray, pix.A}
		}
	}

	return grayImg
}
