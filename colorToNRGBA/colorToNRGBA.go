package colorToNRGBA

import (
	"image"
	"image/color"
)

func ColorToNRGBA(pixels [][]color.Color) *image.NRGBA {
	y := len(pixels)

	if y == 0 {
		return nil
	}

	x := len(pixels[0])

	img := image.NewNRGBA(image.Rect(0, 0, x, y))

	for idx := 0; idx < y; idx++ {
		for jdx := 0; jdx < x; jdx++ {
			r, g, b, a := pixels[idx][x].RGBA()
			img.SetNRGBA(jdx, idx, color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return img
}
