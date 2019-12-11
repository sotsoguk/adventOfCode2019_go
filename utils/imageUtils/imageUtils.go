package imageutils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/adventOfCode2019_go/utils/mathUtils"
)

func PrintGrid(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func ConvertMapToGrid(g map[complex64]int) [][]int {
	// find limits of map

	var minX, maxX, minY, maxY int
	for k := range g {
		currX := int(real(k))
		currY := int(imag(k))
		minX = mathUtils.Min32(currX, minX)
		maxX = mathUtils.Max32(currX, maxX)
		minY = mathUtils.Min32(currY, minY)
		maxY = mathUtils.Max32(currY, maxY)
	}
	// fmt.Println(minX, maxX, minY, maxY)
	w := maxX - minX + 1
	h := maxY - minY + 1
	result := make([][]int, h)
	for r := 0; r < h; r++ {
		result[r] = make([]int, w)
	}
	for k, v := range g {
		currX := int(real(k))
		currY := int(imag(k))
		imgX := currX - minX
		imgY := maxY - currY
		result[imgY][imgX] = v
	}
	return result
}

func RenderGrid(filename string, grid [][]int, pixelSize int) {
	rows, cols := len(grid), len(grid[0])
	h, w := rows*pixelSize, cols*pixelSize
	rect := image.Rect(0, 0, w, h)
	img := image.NewNRGBA(rect)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for y := r * pixelSize; y < (r+1)*pixelSize; y++ {
				for x := c * pixelSize; x < (c+1)*pixelSize; x++ {
					pixel := grid[r][c]
					switch pixel {
					case 0:
						img.Set(x, y, color.RGBA{30, 30, 30, 0xff})
					case 1:
						//img.Set(x, y, color.White)
						img.Set(x, y, color.RGBA{240, 240, 240, 0xff})
					}
				}
			}
		}
	}

	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, img)
}
