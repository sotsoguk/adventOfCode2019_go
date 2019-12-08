package main

import (
	"fmt"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	//"github.com/adventOfCode2019_go/utils/mathUtils"
)

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
		image[i] = 2
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
			if image[k] == 2 {
				image[k] = l[k]
			}
		}
	}

	solution1 = int64(product)
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nRuntime:\t%v\n",
		header, len(lines[0]), solution1, solution2, elapsed)
	fmt.Println()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pixel := image[y*w+x]
			if pixel == 1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
