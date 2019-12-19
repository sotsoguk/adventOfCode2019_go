package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"strconv"
	"strings"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/imageutils"
	"github.com/adventOfCode2019_go/utils/intcode"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

const (
	EMPTY int = iota
	WALL
	BLOCK
	PADDLE
	BALL
)

func getColor(c complex64, grid map[complex64]int) int {
	if grid[c] == 1 {
		return 1
	} else {
		return 0
	}
}

func setColor(c complex64, grid map[complex64]int, color int) {
	grid[c] = color
}
func turnRight(d *complex64) {
	*d *= complex(0, -1)
}
func turnLeft(d *complex64) {
	*d *= complex(0, 1)
}
func doStep(c *complex64, d *complex64) {
	*c += *d
}
func partX(part2 bool, code []int64) (int64, [][]int) {

	grid := make(map[complex64]int)
	var currPos complex64 = complex(0, 0)
	var dir complex64 = complex(0, 1)
	var robot intcode.VM
	if part2 {
		setColor(currPos, grid, 1)
	}
	robot.LoadCode(code)
	robot.Reset()
	for robot.Mode != 99 {

		robot.LoadInput(int64(getColor(currPos, grid)))
		robot.RunCode()
		colorO := int(robot.Output[len(robot.Output)-2])
		cDir := robot.Output[len(robot.Output)-1]
		setColor(currPos, grid, colorO)
		if cDir == 1 {
			turnRight(&dir)
		} else {
			turnLeft(&dir)
		}
		doStep(&currPos, &dir)

	}
	return int64(len(grid)), imageutils.ConvertMapToGrid(grid)
}
func printScreen(screen [][]int, score int, round int) {
	rows, cols := len(screen), len(screen[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			tile := screen[y][x]
			switch tile {
			case EMPTY:
				fmt.Print(" ")
			case WALL:
				fmt.Print("X")
			case BLOCK:
				fmt.Print("#")
			case PADDLE:
				fmt.Print("-")
			case BALL:
				fmt.Print("*")

			}
		}
		fmt.Println()

	}
	fmt.Println("Round: ", round, "\nScore: ", score)
}
func parseScreen(output []int64) (ballX, ballY, paddleX, paddleY, score int) {
	tiles := len(output) / 3
	for i := 0; i < tiles; i++ {
		currTile := output[i*3 : i*3+3]
		currX := int(currTile[0])
		currY := int(currTile[1])
		tile := int(currTile[2])
		// minX = mathUtils.Min32(minX, currX)
		// minY = mathUtils.Min32(minY, currY)
		// maxX = mathUtils.Max32(maxX, currX)
		// maxY = mathUtils.Max32(maxY, currY)
		// screen[currY][currX] = tile
		if currX == -1 {
			score = tile
		} else {
			if tile == BALL {
				ballX = int(currX)
				ballY = int(currY)
			}
			if tile == PADDLE {
				paddleX = int(currX)
				paddleY = int(currY)
			}
		}

	}
	return
}
func makeScreen(output []int64) ([][]int, int) {
	score := 0
	screenDim := 44
	screen := make([][]int, screenDim)
	for i := 0; i < screenDim; i++ {
		screen[i] = make([]int, screenDim)
	}
	tiles := len(output) / 3
	for i := 0; i < tiles; i++ {
		currTile := output[i*3 : i*3+3]
		// if currTile[2] == 2 {
		// 	s1++
		// }
		currX := int(currTile[0])
		currY := int(currTile[1])
		tile := int(currTile[2])

		if currX > -1 {
			screen[currY][currX] = tile
		} else {
			score = tile
		}
	}
	return screen, score
}
func updateScreen(oldScreen [][]int, output []int64) ([][]int, int) {
	var score int
	tiles := len(output) / 3
	screenDim := 44
	screen := make([][]int, screenDim)
	for i := 0; i < screenDim; i++ {
		screen[i] = make([]int, screenDim)
	}
	copy(screen, oldScreen)
	for i := 0; i < tiles; i++ {
		currTile := output[i*3 : i*3+3]
		// if currTile[2] == 2 {
		// 	s1++
		// }
		currX := int(currTile[0])
		currY := int(currTile[1])
		tile := int(currTile[2])

		if currX > -1 {
			screen[currY][currX] = tile
		} else {
			score = tile
		}
	}
	return screen, score
}
func drawDot(img *image.Paletted, pixelSize int, c color.RGBA, px int, py int) {
	offsetX, offsetY := px*pixelSize, py*pixelSize
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {

			img.Set(x+offsetX, y+offsetY, c)

		}
	}
}
func drawBlock(img *image.Paletted, pixelSize int, c color.RGBA, px int, py int) {
	offsetX, offsetY := px*pixelSize, py*pixelSize
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {
			if (x+y)%3 == 0 {
				img.Set(x+offsetX, y+offsetY, c)
			}

		}
	}
}
func drawPaddle(img *image.Paletted, pixelSize int, c color.RGBA, px int, py int) {
	offsetX, offsetY := px*pixelSize, py*pixelSize
	vOffset := int(pixelSize / 4)
	for x := 0; x < pixelSize; x++ {
		for y := vOffset; y < pixelSize-vOffset; y++ {

			img.Set(x+offsetX, y+offsetY, c)

		}
	}
}
func drawBall(img *image.Paletted, pixelSize int, c color.RGBA, px int, py int) {

	offsetX, offsetY := px*pixelSize, py*pixelSize
	cx, cy := pixelSize/2, pixelSize/2
	r := pixelSize / 3
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {
			d := (x-cx)*(x-cx) + (y-cy)*(y-cy)
			if d <= r*r {

				img.Set(x+offsetX, y+offsetY, c)
			}
		}
	}
}
func ScreenToImage(screen [][]int, pixelSize int) *image.Paletted {
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var rows, cols int = 31, len(screen[0])
	var h, w int = pixelSize * rows, pixelSize * cols
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			elem := screen[r][c]
			switch elem {
			case WALL:
				drawDot(img, pixelSize, color.RGBA{255, 255, 255, 0xff}, c, r)
			case BALL:
				drawBall(img, pixelSize, color.RGBA{255, 255, 0, 0xff}, c, r)
			case BLOCK:
				drawBlock(img, pixelSize, color.RGBA{0, 255, 0, 0xff}, c, r)
			case PADDLE:
				drawPaddle(img, pixelSize, color.RGBA{0, 0, 255, 0xff}, c, r)
			case EMPTY:
				drawDot(img, pixelSize, color.RGBA{80, 80, 80, 0xff}, c, r)
			}
		}
	}
	return img
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())
	start := time.Now()

	const (
		year   = 2019
		day    = 13
		output = false
	)
	var (
		solution1, solution2 int64
	)

	// IO

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	lines := readAOC.ReadInput(filePath)
	ll := strings.Split(lines[0], ",")
	code := make([]int64, len(ll))
	for i := range ll {
		code[i], _ = strconv.ParseInt(ll[i], 10, 64)
	}
	var cabinet intcode.VM
	cabinet.LoadCode(code)
	cabinet.Reset()
	cabinet.RunCode()
	f1Screen, _ := makeScreen(cabinet.Output)
	// img1 := ScreenToImage(f1Screen, 20)
	// f, _ := os.Create("Day131.png")
	// png.Encode(f, img1)
	printScreen(f1Screen, 0, 0)
	fmt.Println(len(cabinet.Output), len(cabinet.Output)/3)
	tiles := len(cabinet.Output) / 3
	// s1 := 0
	var minX, maxX, minY, maxY int
	minX, minY = 10000, 10000
	screenDim := 44
	screen := make([][]int, screenDim)
	for i := 0; i < screenDim; i++ {
		screen[i] = make([]int, screenDim)
	}
	for i := 0; i < tiles; i++ {
		currTile := cabinet.Output[i*3 : i*3+3]
		// if currTile[2] == 2 {
		// 	s1++
		// }
		currX := int(currTile[0])
		currY := int(currTile[1])
		tile := int(currTile[2])
		minX = mathUtils.Min32(minX, currX)
		minY = mathUtils.Min32(minY, currY)
		maxX = mathUtils.Max32(maxX, currX)
		maxY = mathUtils.Max32(maxY, currY)
		screen[currY][currX] = tile

	}
	fmt.Println("STATS")
	fmt.Println(parseScreen(cabinet.Output))
	code[0] = 2
	cabinet.LoadCode(code)
	//Part 2
	var images []*image.Paletted
	var delays []int
	cabinet.Reset()
	cabinet.LoadInput(0)
	rnd := 0
	screen2 := make([][]int, screenDim)
	for cabinet.Mode != 99 {
		pngPath := fmt.Sprintf("%d/day%02d/img%05d.png", year, day, rnd)
		cabinet.RunCode()

		// fmt.Println("Round:", rnd, "Mode:", cabinet.Mode)
		// fmt.Println(parseScreen(cabinet.Output))
		bx, _, px, _, score := parseScreen(cabinet.Output)
		// printScreen(cabinet.Output, 0, 0)
		if rnd == 0 {
			screen2, _ = makeScreen(cabinet.Output)
		} else {
			screen2, _ = updateScreen(screen2, cabinet.Output)
		}
		img := ScreenToImage(screen2, 20)
		// f, err := os.Create(pngPath)
		// if err != nil {
		// 	fmt.Println("EOFRR")
		// 	break
		// }
		// png.Encode(f, img)
		// printScreen(screen2, 0, 0)
		images = append(images, img)
		delays = append(delays, 0)
		if cabinet.Mode == 99 || rnd > 150 {
			fmt.Println("FINAL:", score)
			f, err := os.Create("day13_2.gif")
			if err != nil {
				fmt.Println("ERROF F")
				break
			}
			defer f.Close()
			gif.EncodeAll(f, &gif.GIF{Image: images, Delay: delays})
			break
		}
		cabinet.ClearOuput()
		input := 0
		if bx < px {
			input = -1
		} else if bx > px {
			input = 1
		}
		cabinet.LoadInput(int64(input))
		rnd++
		fmt.Println(rnd, pngPath)
		// if rnd > 100 {
		// 	fmt.Println("FINAL:", score)
		// 	break
		// }
	}
	finalScreen, finalScore := makeScreen(cabinet.Output)
	printScreen(finalScreen, finalScore, rnd)
	// cabinet.RunCode()
	// fmt.Println("Mode:", cabinet.Mode)
	// fmt.Println(minX, maxX, minY, maxY)
	// fmt.Println(s1)
	// printScreen(screen, 12348, 10)

	// solution1, grid1 := partX(false, code)
	// solution2, grid2 := partX(true, code)
	// if output {
	// 	imageutils.PrintGrid(grid1)
	// }
	elapsed := time.Since(start)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v\n",
		header, len(lines), solution1, solution2, elapsed)
	// imageutils.PrintGrid(grid2)
	//imageutils.RenderGrid("day11_01.png", grid1, 10)
}
