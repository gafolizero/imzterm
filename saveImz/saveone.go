package saveImz

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func SaveRectYX(filePath string, grid [][]color.Color) {
	ylen, xlen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)

	for y := range ylen {
		for x := range xlen {
			img.Set(x, y, grid[y][x])
		}
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal("cannot create output file", err)
	}

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatal("failed to encode image", err)
	}
}

func SaveRectXY(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)

	for y := range ylen {
		for x := range xlen {
			img.Set(x, y, grid[x][y])
		}
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal("cannot create output file", err)
	}

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatal("failed to encode image", err)
	}
}
