package flip

import "image/color"

func flip(grid [][]color.Color) {
	for idx := range len(grid) {
		row := grid[idx]
		for jdx := range len(row) / 2 {
			kdx := len(row) - jdx - 1
			row[jdx], row[kdx] = row[kdx], row[jdx]
		}
	}
}
