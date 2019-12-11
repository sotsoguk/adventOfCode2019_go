package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"sort"
	"time"

	readAOC "github.com/adventOfCode2019_go/utils"
	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type Slope struct {
	dx       int
	dy       int
	negative bool
}

type Coord struct {
	x int
	y int
}

func (p Coord) distance(q Coord) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))

}

func (s *Slope) shorten() {
	gcds := mathUtils.Abs32(mathUtils.Gcd(s.dx, s.dy))
	s.dx /= gcds
	s.dy /= gcds
}
func (s *Slope) toAngle() float64 {

	angle := math.Atan2(float64(s.dy), float64(s.dx))/(2*math.Pi)*360.0 + 90
	if angle < 0 {
		angle += 360
	}
	return angle
}

func makeGrid(input []string) [][]int {
	rows, cols := len(input), len(input)
	spaceMap := make([][]int, rows)
	for i := 0; i < rows; i++ {
		spaceMap[i] = make([]int, cols)
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 35 {
				spaceMap[y][x] = 1
			}
		}
	}
	return spaceMap
}

func printGrid(input [][]int) {
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

func countAsteroids(grid [][]int, posX, posY int) map[Slope][]Coord {
	rows, cols := len(grid), len(grid[0])
	dirs := make(map[Slope][]Coord)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 && !(posX == x && posY == y) {
				var dir Slope
				dir.dx = x - posX
				dir.dy = y - posY
				(&dir).shorten()
				newPos := Coord{x, y}
				dirs[dir] = append(dirs[dir], newPos)
			}
		}
	}
	return dirs
}

func part1(grid [][]int) (int, Coord) {
	maxAsteroids := 0
	var maxX, maxY int
	rows, cols := len(grid), len(grid[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 {
				currAsteroids := len(countAsteroids(grid, x, y))
				if currAsteroids > maxAsteroids {
					maxAsteroids = currAsteroids
					maxX = x
					maxY = y
				}
			}
		}
	}
	return maxAsteroids, Coord{maxX, maxY}
}

func part2(turret Coord, grid [][]int, target int) int {

	slopeAsteroidsMap := countAsteroids(grid, turret.x, turret.y)
	anglesAsteroidsMap := make(map[float64][]Coord)
	angles := make([]float64, 0)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()
		anglesAsteroidsMap[angle] = slopeAsteroidsMap[slope]
		sort.Slice(anglesAsteroidsMap[angle], func(p, q int) bool {
			return anglesAsteroidsMap[angle][p].distance(turret) < anglesAsteroidsMap[angle][q].distance(turret)
		})
		if len(anglesAsteroidsMap[angle]) >= 1 {
			angles = append(angles, angle)
		}
		sort.Float64s(angles)
	}
	i := 0
	for i < target {
		for _, a := range angles {
			if len(anglesAsteroidsMap[a]) > 0 {
				i++
				if i == target {
					fmt.Println(anglesAsteroidsMap[a][0])
					return anglesAsteroidsMap[a][0].x*100 + anglesAsteroidsMap[a][0].y
				}
				anglesAsteroidsMap[a] = anglesAsteroidsMap[a][1:]
			}
		}
	}
	return 0
}
func drawAsteroid(img *image.Paletted, pixelSize int, c color.RGBA, p Coord) {

	offsetX, offsetY := p.x*pixelSize, p.y*pixelSize
	cx, cy := pixelSize/2, pixelSize/2
	r := pixelSize / 4
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {
			d := (x-cx)*(x-cx) + (y-cy)*(y-cy)
			if d <= r*r {
				c.A = uint8(0xff - 0xff*d/r/r)
				img.Set(x+offsetX, y+offsetY, c)
			}
		}
	}
}
func drawAsteroidExploding(img *image.Paletted, pixelSize int, c color.RGBA, p Coord) {

	offsetX, offsetY := p.x*pixelSize, p.y*pixelSize
	cx, cy := pixelSize/2, pixelSize/2
	r := pixelSize/2 - 2
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {
			d := (x-cx)*(x-cx) + (y-cy)*(y-cy)
			if d <= r*r && !(x%3 == 0 && y%2 == 0) {
				c.A = uint8(0xff - 0xff*d/r/r)
				img.Set(x+offsetX, y+offsetY, c)
			}
		}
	}
}

func drawTurret(img *image.Paletted, pixelSize int, c color.RGBA, p Coord) {
	offsetX, offsetY := p.x*pixelSize, p.y*pixelSize
	for x := 0; x < pixelSize; x++ {
		for y := 0; y < pixelSize; y++ {

			img.Set(x+offsetX, y+offsetY, c)

		}
	}
}
func part2Animation(turret Coord, grid [][]int, target int, filePath string) int {

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	var pixelSize int = 20
	var h, w int = pixelSize * len(grid), pixelSize * len(grid[0])
	var images []*image.Paletted
	var delays []int
	f, err := os.Create("test.gif")
	if err != nil {
		fmt.Println(err)
		fmt.Println("=!=!")
		return -1
	}
	defer f.Close()
	defer func() {
		gif.EncodeAll(f, &gif.GIF{Image: images, Delay: delays})
	}()
	slopeAsteroidsMap := countAsteroids(grid, turret.x, turret.y)
	anglesAsteroidsMap := make(map[float64][]Coord)
	angles := make([]float64, 0)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()
		anglesAsteroidsMap[angle] = slopeAsteroidsMap[slope]
		sort.Slice(anglesAsteroidsMap[angle], func(p, q int) bool {
			return anglesAsteroidsMap[angle][p].distance(turret) < anglesAsteroidsMap[angle][q].distance(turret)
		})
		if len(anglesAsteroidsMap[angle]) >= 1 {
			angles = append(angles, angle)
		}
		sort.Float64s(angles)
	}

	i := 0

	for i < target {
		found := false
		for _, a := range angles {
			var elim Coord
			if len(anglesAsteroidsMap[a]) > 0 {
				i++
				found = true
				if i == target {
					fmt.Println(anglesAsteroidsMap[a][0])
					return anglesAsteroidsMap[a][0].x*100 + anglesAsteroidsMap[a][0].y
				}
				elim = anglesAsteroidsMap[a][0]
				anglesAsteroidsMap[a] = anglesAsteroidsMap[a][1:]
			}
			if !(elim.x == 0 && elim.y == 0) {
				img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
				images = append(images, img)
				delays = append(delays, 0)
				drawAsteroidExploding(img, pixelSize, color.RGBA{255, 100, 100, 100}, elim)
				for _, asts := range anglesAsteroidsMap {
					for _, p := range asts {
						// for x := (p.x) * pixelSize; x < (p.x+1)*pixelSize; x++ {
						// 	for y := (p.y) * pixelSize; y < (p.y+1)*pixelSize; y++ {
						// 		img.Set(x, y, color.RGBA{0xff, 0x00, 0x00, 0x00})
						// 	}
						// }
						drawAsteroid(img, pixelSize, color.RGBA{211, 211, 211, 0xff}, p)
					}
				}
				drawTurret(img, pixelSize, color.RGBA{0, 120, 180, 0xff}, turret)
				// image setup
			}
		}
		if !found {
			break
		}
	}
	//f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return -1
	// }
	return 0
}
func main() {
	// Debug path
	// lines := readAOC.ReadInput("../../2019/inputs/input09_2019.txt")
	// fmt.Println(os.Getwd())

	const (
		year   = 2019
		day    = 10
		target = 200
	)
	var (
		solution1, solution2 int
	)

	filePath := fmt.Sprintf("%d/inputs/input%02d_%d.txt", year, day, year)
	header := fmt.Sprintf("AoC %d - Day %02d\n-----------------\n", year, day)
	// filePath := fmt.Sprintf("%d/inputs/input10_3.txt", year)
	lines := readAOC.ReadInput(filePath)
	start := time.Now()
	spaceMap := makeGrid(lines)
	solution1, maxPos := part1(spaceMap)
	//solution2 = part2(maxPos, spaceMap, target)
	elapsed := time.Since(start)
	// fp := "2019/day10/day10.gif"
	fp := "day10.gif"
	part2Animation(maxPos, spaceMap, 500, fp)
	fmt.Printf("%sLength of Input (lines):\t%v\n\nSolution:\nPart1:\t%v\nPart2:\t%v\nTime:\t%v",
		header, len(lines), solution1, solution2, elapsed)

}
