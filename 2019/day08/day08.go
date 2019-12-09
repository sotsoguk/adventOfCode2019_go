package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	//"github.com/adventOfCode2019_go/utils/mathUtils"
)

const (
	pixelBlack = iota
	pixelWhite
	pixelTransparent
)

func printSIF(image []int, width int, height int) {
	fmt.Println()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := image[y*width+x]
			if pixel == pixelWhite {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func renderImage(filename string, imgData []int, width int, height int, pixelSize int) {
	rect := image.Rect(0, 0, width*pixelSize, height*pixelSize)
	img := image.NewRGBA(rect)
	// blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{14, 184, 14, 0xff}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			for y := r * pixelSize; y < (r+1)*pixelSize; y++ {
				for x := c * pixelSize; x < (c+1)*pixelSize; x++ {
					pixel := imgData[r*width+c]
					switch pixel {
					case pixelBlack:
						img.Set(x, y, color.Black)
					case pixelWhite:
						img.Set(x, y, green)
					}
				}
			}
		}
	}
	f, _ := os.Create(filename)
	png.Encode(f, img)
}
func main() {
	// Debug path
	// fmt.Println(os.Getwd())
	// filePath = fmt.Sprintf("../../%d/inputs/input%02d_%d.txt", year, day, year)

	start := time.Now()
	const (
		year = 2019
		day  = 8
		w    = 25
		h    = 6
	)

	var (
		solution1, solution2 int64
	)
	// IO
	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	imagePath := fmt.Sprintf("%d/day%02d/day%02d.png", year, day, day)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)

	imageData := make([]int, len(lines[0]))
	for i, c := range lines[0] {
		imageData[i] = int(c - '0')
	}

	layers := len(imageData) / (w * h)
	minZeros := 200
	product := 0
	image := make([]int, w*h)
	for i := range image {
		image[i] = pixelTransparent
	}
	for i := 0; i < layers; i++ {
		l := imageData[i*150 : (i+1)*150]
		cnt := make(map[int]int)
		for _, p := range l {
			cnt[p]++
		}
		if cnt[0] < minZeros {
			product = cnt[1] * cnt[2]
			minZeros = cnt[0]
		}
	}
	for i := 0; i < layers; i++ {
		l := imageData[i*150 : (i+1)*150]
		for k := range l {
			if image[k] == pixelTransparent {
				image[k] = l[k]
			}
		}
	}

	solution1 = int64(product)
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nRuntime:\t%v\n",
		header, len(lines[0]), solution1, solution2, elapsed)

	printSIF(image, w, h)
	renderImage(imagePath, image, w, h, 20)
}
